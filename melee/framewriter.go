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

	"gopkg.in/fatih/set.v0"
)

/*
* FrameWriter handles only the WRITING of Matches and Frames to disk.
 */
type FrameWriter struct {
	StatsFile   *os.File
	Writer      *bufio.Writer
	Mutex       sync.Mutex
	FrameBuffer *FrameBuffer
	Match       Match
}

type Match struct {
	SelfCharacter      string   `json:"self_character"`
	OpponentCharacters []string `json:"opponent_characters"`

	SelfPort      Port     `json:"self_port"`
	OpponentPorts PortList `json:"opponent_ports"`

	// Port Number 1-4
	Winner int `json:"winner"`

	Stage string    `json:"stage"`
	Time  time.Time `json:"time"`

	DiskFrames  []DiskFrame `json:"disk_frames"`
	ActiveFlags *set.Set    `json:"active_flags"`
}

func NewFrameWriter() *FrameWriter {
	fb := NewFrameBuffer()

	fw := FrameWriter{
		Mutex:       sync.Mutex{},
		FrameBuffer: &fb,
		Match: Match{
			DiskFrames:  make([]DiskFrame, 1),
			ActiveFlags: set.New(),
		},
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
		fw.Match.Time = time.Now()

		fw.GenerateSummaryText()

		name := NewStatFileName()
		f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalln(err)
		}

		fw.StatsFile = f
		fw.Writer = bufio.NewWriter(f)

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

func (fw *FrameWriter) AddFlag(f Flag) {
	fw.Match.ActiveFlags.Add(f)
}

func (fw *FrameWriter) RemoveFlag(f Flag) {
	fw.Match.ActiveFlags.Remove(f)
}

func (fw *FrameWriter) CheckFlag(f Flag) bool {
	return fw.Match.ActiveFlags.Has(f)
}

func (fw *FrameWriter) GetSelfPort() Port {
	return fw.Match.SelfPort
}

func (fw *FrameWriter) GetOpponentPorts() PortList {
	return fw.Match.OpponentPorts
}
