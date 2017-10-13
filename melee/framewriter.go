package melee

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

type Frame struct {
	Players []Player `json:"players"`
	empty   bool     `json:"empty"`
}

type BasicFrame struct {
	// no discrete use now for FrameNumber but may be useful later to
	// always store this.  Violates YAGNI for good reason
	FrameNumber uint32 `json:"frame_number"`
	Action      string `json:"action"`
	Message     string `json:"message"`
}

// Data actually written to disk. For the most part this should be used
// similar to a C union
type DiskFrame struct {
	BasicFrame    BasicFrame
	DetailedFrame Frame
}

func NewFrame(players PlayerContainer) Frame {
	f := Frame{
		players[0 : PLAYER4+1], false,
	}
	return f
}

func (f *Frame) Empty() bool {
	return f.empty
}

const FRAME_BUFFER_SIZE = 30 // hold 0.5 seconds of frames in memory
type FrameBuffer struct {
	Frames *sll.List
}

func NewFrameBuffer() FrameBuffer {
	return FrameBuffer{sll.New()}
}

func (fb *FrameBuffer) Insert(f Frame) {
	fb.Frames.Prepend(f)

	if fb.Frames.Size() > FRAME_BUFFER_SIZE {
		fb.Frames.Remove(FRAME_BUFFER_SIZE)
	}
}

// second return parameter is true if index is withing bounds of the list and
// it is not empty
func (fb *FrameBuffer) Get(i int) (Frame, bool) {
	f, success := fb.Frames.Get(i)
	return f.(Frame), success
}

// this method should only be called AFTER the most recent frame is inserted
func (fb *FrameBuffer) GetPrevious() (Frame, bool) {
	return fb.Get(1)
}

func (fb *FrameBuffer) GetCurrent() (Frame, bool) {
	return fb.Get(0)
}

/* Since this can be a bit confusing to understand, here's an explanation:
* GetRange(6, 10) would return a slice of size 4 equivalent to the following:
*
* frames[0] = fb.Get(6)
* frames[1] = fb.Get(7)
* frames[2] = fb.Get(8)
* frames[3] = fb.Get(9)
*
* Meaning that the indexing of the returned slice is the same as the list(e.g.
* lower index = newer frame)
 */
func (fb *FrameBuffer) GetRange(start, end int) []Frame {
	frames := make([]Frame, end-start)

	for i := start; i < end; i++ {
		f, exists := fb.Frames.Get(i)
		if exists {
			frames[i-start] = f.(Frame)
		} else {
			frames[i-start] = Frame{empty: true}
		}
	}

	return frames
}

type FrameWriter struct {
	StatsFile   *os.File
	Writer      *bufio.Writer
	Mutex       sync.Mutex
	FrameBuffer *FrameBuffer
	Match       Match
}

//TODO: dynamic/changeable opponent ports
type Match struct {
	SelfCharacter     string `json:"self_character"`
	OpponentCharacter string `json:"opponent_character"`
	SelfPort          int    `json:"self_port"`

	// Port Number 1-4
	Winner int `json:"winner"`

	Stage string    `json:"stage"`
	Time  time.Time `json:"time"`

	DiskFrames []DiskFrame `json:"disk_frames"`
}

func NewFrameWriter() *FrameWriter {
	fb := NewFrameBuffer()

	fw := FrameWriter{
		Mutex:       sync.Mutex{},
		FrameBuffer: &fb,
		Match:       Match{DiskFrames: make([]DiskFrame, 1)},
	}

	return &fw
}

func NewStatFileName() string {
	//u, _ := uuid.NewV4()
	u := time.Now().Format("2006-01-02T15:04:05.999999-07:00")
	name := "./stats/" + u + ".json"
	return name
}

func (fw *FrameWriter) LogFrame(f Frame) {
	fw.FrameBuffer.Insert(f)

	validator := NewFrameValidator(fw.FrameBuffer)
	_, exists := fw.FrameBuffer.GetCurrent()

	if exists {
		validator.CheckAll()
	}
}

func (fw *FrameWriter) Write(data DiskFrame) {
	fw.Mutex.Lock()
	fw.Match.DiskFrames = append(fw.Match.DiskFrames, data)
	fw.Mutex.Unlock()
}

func (fw *FrameWriter) Flush() {
	fw.Mutex.Lock()

	if len(fw.Match.DiskFrames) > 1 {
		fw.GenerateSummaryText()

		name := NewStatFileName()
		f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalln(err)
		}

		fw.StatsFile = f
		fw.Writer = bufio.NewWriter(f)

		fw.Match.Time = time.Now()
		encoded, _ := json.Marshal(fw.Match)
		fw.Writer.Write(encoded)
		fw.Match.DiskFrames = make([]DiskFrame, 1)
	}

	if fw.Writer != nil && fw.StatsFile != nil {
		fw.Writer.Flush()
		fw.StatsFile.Close()
	}

	fw.Mutex.Unlock()
}

func (fw *FrameWriter) GenerateSummaryText() (text string) {
	var l_pass, l_miss int

	CUI.ClearLog()

	for _, data := range fw.Match.DiskFrames {
		if data.BasicFrame.Message == "L_CANCEL_PASS" {
			l_pass++
		} else if data.BasicFrame.Message == "L_CANCEL_MISS" {
			l_miss++
		}
	}

	missed_percent := float32(l_miss) / float32(l_miss+l_pass) * 100.0
	text += "Missed " + strconv.Itoa(l_miss) + "/" +
		strconv.Itoa(l_miss+l_pass) + "(" + fmt.Sprintf("%.2f", missed_percent) +
		"%)" + " L-Cancels.\n"

	CUI.LogText(text)

	return text
}

func (fw *FrameWriter) Close() {
	fw.Flush()
	fw.StatsFile.Close()
}
