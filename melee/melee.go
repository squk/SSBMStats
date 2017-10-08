package melee

import (
	"errors"
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui"
)

var Dolphin *DolphinManager
var GameState *GameStateManager
var FWriter *FrameWriter
var CUI *ConsoleUI

func init() {
	Dolphin = NewDolphinManager()
	GameState = NewGameStateManager()
	CUI = NewConsoleUI()
	FWriter = NewFrameWriter()

	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetFlags(log.Lshortfile)
	log.SetOutput(f)

	err = ui.Init()
	if err != nil {
		panic(err)
	}

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		Dolphin.StopLoop()
		FWriter.Close()
		f.Close()
		ui.StopLoop()
	})
	go func() {
		for Dolphin.RUNNING {
			time.Sleep(200 * time.Millisecond)

			CUI.Draw()
		}
	}()
	go ui.Loop()

}

func Init() (err error) {
	if !Dolphin.Init() {
		err = errors.New("DolphinManager failed to initialize")
	}
	return err
}
