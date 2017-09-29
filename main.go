package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"runtime"

	. "./melee"
)

func main() {
	d := NewDolphin()

	if d.DolphinPath == "" {
		exists := false

		for !exists {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter Dolphin Path: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(strings.TrimSuffix(text, "\n"), " ")
			exe_name := "Dolphin"

			if runtime.GOOS == "windows" {
				exe_name += ".exe"
			} else {
				exe_name += ".app"
			}
			exe_exists, _ := FilepathExists(filepath.Join(text, exe_name))
			user_exists, _ := FilepathExists(filepath.Join(text, "User"))

			exists = exe_exists || user_exists || exists

			if !exists {
				fmt.Println("\nInvalid Dolphin path")
			}
			exists = d.SetPath(filepath.Join(text, "User"))
		}
	}

	if !d.Initialize() {
		fmt.Println("Initialization failed")
		return
	}

	defer (*d.GameState.Socket).Close()

	running := true
	for running {
		d.GameState.Step()

		//if d.GameState.ProcessingTime*1000 > 12 {
		//fmt.Println("WARNING: Last frame took ", d.GameState.ProcessingTime*1000, "ms to process.")
		//}

		if d.GameState.MenuState == IN_GAME {
			//framedata.recordframe(gamestate)
			fmt.Println("In Game")
		} else if d.GameState.MenuState == CHARACTER_SELECT {
			//fmt.Println("Character Select")
		} else if d.GameState.MenuState == POSTGAME_SCORES {
			fmt.Println("Postgame")
		} else if d.GameState.MenuState == STAGE_SELECT {
			fmt.Println("Stage Select")
		}
		//log.logframe(d.GameState)
		//log.writeframe()
	}
}
