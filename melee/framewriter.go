package melee

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
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
}

const JSON_DIR = "./statistics.json"

func NewFrameWriter() FrameWriter {
	f, err := os.OpenFile(JSON_DIR, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}

	w := bufio.NewWriter(f)
	enc := json.NewEncoder(w)
	fw := FrameWriter{f, w, enc}

	return fw
}

func (fw *FrameWriter) WriteFrame(f Frame) {
	if f.MenuState == IN_GAME {
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
