package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"runtime"

	. "./melee"
	ui "github.com/gizak/termui"
)

var d Dolphin

func main() {
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

	initialize()

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		d.StopLoop()
		d.GameState.FrameWriter.Close()
		f.Close()
		ui.StopLoop()
	})
	go func() {
		for d.RUNNING {
			time.Sleep(200 * time.Millisecond)

			d.CUI.Draw()
		}
	}()
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
	}
}

func run() {
	done := make(chan bool)

	go func() {
		for d.RUNNING {
			d.GameState.Update()
		}
		done <- true
	}()

	<-done

	(*d.GameState.Socket).Close()
}
