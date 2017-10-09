package main

import (
	"log"

	melee "./melee"
	ui "github.com/gizak/termui"
)

func main() {
	err := melee.Init()

	if err != nil {
		log.Fatalln(err)
	}

	melee.GameState.Update()

	ui.Loop()
}
