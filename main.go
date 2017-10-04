package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"runtime"

	. "./melee"
	ui "github.com/gizak/termui"
)

var d Dolphin

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	initialize()

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		d.StopLoop()
		d.GameState.FrameWriter.Close()
		ui.StopLoop()
	})
	ui.Handle("/timer/1s", func(e ui.Event) {
		d.CUI.Draw()
	})
	go ui.Loop()

	run()
}

func initialize() {
	d = NewDolphin()

	if d.DolphinPath == "" {
		exists := false

		for !exists {
			reader := bufio.NewReader(os.Stdin)
			//fmt.Print("Enter Dolphin Path: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(strings.TrimSuffix(text, "\n"), " ")
			exe_name := "Dolphin"

			if runtime.GOOS == "windows" {
				exe_name += ".exe"
			} else {
				exe_name += ".app"
			}
			//user_exists, _ := FilepathExists(filepath.Join(text, "User"))
			//exists = user_exists || exists
			exists, _ = FilepathExists(text)

			if !exists {
				//fmt.Println("\nInvalid Dolphin path")
			}
			finalPath := filepath.Join(text, "User")
			d.SetPath(finalPath)
		}
	}

	if !d.Initialize() {
		log.Fatalln()
		return
	}
}

func run() {
	defer (*d.GameState.Socket).Close()

	newFrame := make(chan bool)

	for d.RUNNING {
		d.GameState.Update(newFrame)
		<-newFrame

		speed, err := d.GameState.Players[1].GetFloat(SPEED_ANIMATION)
		if err != nil {
			//continue
		}

		action, err := d.GameState.Players[1].GetAction()
		if err != nil {
			//continue
		}

		l_cancellable := (action == UAIR_LANDING || action == BAIR_LANDING ||
			action == NAIR_LANDING || action == DAIR_LANDING || action ==
			FAIR_LANDING)

		if speed >= 3.0 && l_cancellable {
			fmt.Println("L-Canceled", action)
		}
	}
}
