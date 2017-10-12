package main

import (
	"log"

	melee "github.com/ctnieves/hopelessstats/melee"
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
