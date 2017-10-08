package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"runtime"

	melee "./melee"
)

func main() {
	initialize()
	run()
}

func initialize() {
	if melee.Dolphin.DolphinPath == "" {
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
			exists, _ = melee.FilepathExists(text)

			if !exists {
				//fmt.Println("\nInvalid Dolphin path")
			}
			finalPath := filepath.Join(text, "User")
			melee.Dolphin.SetPath(finalPath)
		}
	}

	err := melee.Init()
	if err != nil {
		log.Fatalln()
	}
}

func run() {
	done := make(chan bool)
	melee.GameState.Update()

	<-done

	(*melee.GameState.Socket).Close()
}
