package melee

import (
	"log"
	"os"
	"time"

	"github.com/ctnieves/golphin"
)

var Dolphin *golphin.Golphin
var CUI *ConsoleUI
var GameState *GameStateManager
var FWriter *FrameWriter

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Init() (err error) {
	// setup log
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)

	log.SetFlags(log.Lshortfile)
	log.SetOutput(f)

	// setup Golphin and memory locations
	Dolphin = golphin.New()
	err = Dolphin.SetPath("/Users/christian/Desktop/FM/Dolphin.app/Contents/Resources/User")
	check(err)

	err = Dolphin.Init()
	check(err)

	for _, a := range Locations {
		Dolphin.Subscribe(a)
	}
	Dolphin.WriteLocations()

	CUI = NewConsoleUI()
	GameState = NewGameStateManager()
	FWriter = NewFrameWriter()

	go func() {
		for Dolphin.Looping {
			time.Sleep(200 * time.Millisecond)
			CUI.Draw()
		}
	}()

	return err
}
