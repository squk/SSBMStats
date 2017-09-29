package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"runtime"

	. "./melee"
	"github.com/gosuri/uilive"
)

var d Dolphin
var writer *uilive.Writer

func main() {
	writer = uilive.New()
	writer.Start()

	run()
}

func run() {
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
		//fmt.Println("Initialization failed")
		return
	}

	defer (*d.GameState.Socket).Close()

	newFrame := make(chan bool)

	d.GameState.Update(newFrame)

	for d.RUNNING {
		fmt.Fprintln(writer, d.GameState.Players[1].Values[CHARACTER])
		<-newFrame
	}
}
