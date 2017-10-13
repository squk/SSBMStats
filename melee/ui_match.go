package melee

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	ui "github.com/gizak/termui"
)

var page = 0
var selected_data Match

func (c *ConsoleUI) DrawMatchView() {
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		c.UIState = DEFAULT
		ui.Clear()
	})

	ui.Handle("/sys/kbd/<left>", func(ui.Event) {
		if page > 0 {
			page--
		}
	})
	ui.Handle("/sys/kbd/<right>", func(ui.Event) {
		page++
	})

	ui.Handle("/sys/kbd/0", func(ui.Event) { c.SelectMatch(0) })
	ui.Handle("/sys/kbd/1", func(ui.Event) { c.SelectMatch(1) })
	ui.Handle("/sys/kbd/2", func(ui.Event) { c.SelectMatch(2) })
	ui.Handle("/sys/kbd/3", func(ui.Event) { c.SelectMatch(3) })
	ui.Handle("/sys/kbd/4", func(ui.Event) { c.SelectMatch(4) })
	ui.Handle("/sys/kbd/5", func(ui.Event) { c.SelectMatch(5) })
	ui.Handle("/sys/kbd/6", func(ui.Event) { c.SelectMatch(6) })
	ui.Handle("/sys/kbd/7", func(ui.Event) { c.SelectMatch(7) })
	ui.Handle("/sys/kbd/8", func(ui.Event) { c.SelectMatch(8) })
	ui.Handle("/sys/kbd/9", func(ui.Event) { c.SelectMatch(9) })

	c.DrawMatchSelector()
}

var SELECTOR_W int = 56
var SELECTOR_H int = 12

func (c *ConsoleUI) DrawMatchSelector() {
	files, err := ioutil.ReadDir("./stats")
	if err != nil {
		log.Fatalln(err)
	}
	var matches []string

	if page*10+10 > len(files) {
		page = 0
	}

	for i, f := range files {
		if i >= page*10 && i < (page*10)+10 {
			selector := "([" + strconv.Itoa(i-page*10) + "](fg-yellow)) " + f.Name()
			matches = append(matches, selector)
		}
	}

	ls := ui.NewList()
	ls.Items = matches
	ls.Width = SELECTOR_W
	ls.Height = SELECTOR_H
	ls.Y = 0
	c.AdjustY(SELECTOR_H)

	ui.Render(ls)
	c.DrawMatchData()
}

var MATCH_DATA_W int = 56
var MATCH_DATA_H int = 20

func (c *ConsoleUI) DrawMatchData() {
	stage := ui.NewPar(selected_data.Stage)
	stage.X = 0
	stage.Y = c.CurrentY

	stage.Height = 3
	stage.Width = 20

	stage.TextFgColor = ui.ColorWhite
	stage.BorderLabel = "STAGE"
	stage.BorderFg = ui.ColorRed
	stage.BorderLabelFg = ui.ColorCyan
	ui.Render(stage)

	character := ui.NewPar(selected_data.SelfCharacter + " vs " +
		strings.Join(selected_data.OpponentCharacters, ", "))
	character.X = 20
	character.Y = c.CurrentY

	character.Height = 3
	character.Width = 30

	character.TextFgColor = ui.ColorWhite
	character.BorderLabel = "CHARACTERS"
	character.BorderFg = ui.ColorRed
	character.BorderLabelFg = ui.ColorCyan
	ui.Render(character)

	c.AdjustY(3)

	c.DrawMatchHelpBox()
}

var HELP_MATCH_W int = 56
var HELP_MATCH_H int = 2

func (c *ConsoleUI) DrawMatchHelpBox() {
	c.InsertLineBreaks(1)

	par_text := []string{
		"PRESS [q](fg-red) TO GO BACK.",
		"USE [← →](fg-red) TO CHANGE PAGE. ",
	}
	HELP_MATCH_H = 2 + len(par_text)
	p := ui.NewPar(strings.Join(par_text, "\n"))
	p.Y = c.CurrentY
	c.AdjustY(HELP_MATCH_H)
	p.Height = HELP_MATCH_H
	p.Width = HELP_MATCH_W
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "HELP"
	p.BorderFg = ui.ColorCyan

	ui.Render(p)
}

func (c *ConsoleUI) SelectMatch(match int) {
	log.Println("par:", match)
	log.Println("calc", (page*10)+match)

	files, err := ioutil.ReadDir("./stats")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := ioutil.ReadFile("./stats/" + files[(page*10)+match].Name())
	if err != nil {
		log.Fatalln(err)
	}

	data := Match{}
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Fatalln(err)
	}
	selected_data = data
}
