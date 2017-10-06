package melee

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"sync"
)

type FrameWriter struct {
	StatsFile   *os.File
	Writer      *bufio.Writer
	Encoder     *json.Encoder
	Mutex       sync.Mutex
	FrameBuffer *FrameBuffer
}

const JSON_DIR = "./statistics.json"

func NewFrameWriter() FrameWriter {
	f, err := os.OpenFile(JSON_DIR, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	enc := json.NewEncoder(w)
	fb := NewFrameBuffer()

	fw := FrameWriter{f, w, enc, sync.Mutex{}, &fb}

	return fw
}

func (fw *FrameWriter) LogFrame(f Frame) {
	fw.FrameBuffer.Insert(f)
	fw.Write(f)
}

func (fw *FrameWriter) Write(f Frame) {
	if fw.ValidateWrite() {
		err := fw.Encoder.Encode(f)
		if err != nil {
			log.Fatalln(err)
		}

		fw.Writer.Flush()
	}
}

func (fw *FrameWriter) Close() {
	fw.Writer.Flush()
	fw.StatsFile.Close()
}

func (fw *FrameWriter) ValidateWrite() bool {
	validator := NewFrameValidator(fw.FrameBuffer)
	current_frame, exists := fw.FrameBuffer.GetCurrent()

	if exists {
		if current_frame.MenuState == IN_GAME {
			if validator.AnyPassed() {
				return true
			}
		}
	}

	return false
}
