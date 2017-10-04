package melee

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type Frame struct {
	MenuState MenuState
	Stage     Stage
	Players   []Player
}

func NewFrame(players PlayerContainer, m MenuState, s Stage) Frame {
	f := Frame{
		m, s, players[PLAYER1 : PLAYER4+1],
	}
	return f
}

type FrameWriter struct {
	StatsFile *os.File
	Writer    *bufio.Writer
	Encoder   *json.Encoder
	Mutex     sync.Mutex
}

const JSON_DIR = "./statistics.json"

func NewFrameWriter() FrameWriter {
	f, err := os.OpenFile(JSON_DIR, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	enc := json.NewEncoder(w)
	fw := FrameWriter{f, w, enc, sync.Mutex{}}

	return fw
}

func (fw *FrameWriter) WriteFrame(f Frame) {
	if f.MenuState == IN_GAME {
		err := fw.Encoder.Encode(f)
		if err != nil {
			panic(err)
		}

		fw.Writer.Flush()
	}
}

func (fw *FrameWriter) Close() {
	fw.Writer.Flush()
	fw.StatsFile.Close()
}
