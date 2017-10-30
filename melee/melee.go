package melee

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ctnieves/golphin"
)

var Dolphin *golphin.Golphin
var CUI *ConsoleUI
var GameState *MeleeState
var FWriter *FrameWriter

func Init() (err error) {
	// setup log
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}

	log.SetFlags(log.Lshortfile)
	log.SetOutput(f)

	// setup Golphin and memory locations
	Dolphin = golphin.New()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = Dolphin.Init()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for k, _ := range GetMemoryMap() {
		Dolphin.Subscribe(k)
	}
	Dolphin.CommitLocations()

	CUI = NewConsoleUI()
	GameState = NewMeleeState()
	FWriter = NewFrameWriter()

	go func() {
		for Dolphin.Looping {
			time.Sleep(200 * time.Millisecond)
			CUI.Draw()
		}
	}()

	return err
}
