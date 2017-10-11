package melee

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	ui "github.com/gizak/termui"
)

type ConsoleUI struct {
	LogEntries         []string
	Draws              uint64
	CurrentX, CurrentY int
}

func NewConsoleUI() *ConsoleUI {
	err := ui.Init()
	if err != nil {
		log.Fatalln(err)
	}

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		(*GameState.Socket).Close()
		FWriter.Close()
		Dolphin.StopLoop()
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/<left>", func(ui.Event) {
		Dolphin.DecreasePort()
	})
	ui.Handle("/sys/kbd/<right>", func(ui.Event) {
		Dolphin.IncreasePort()
	})

	return &ConsoleUI{[]string{" ", " ", " "}, 0, 0, 0}
}

func (c *ConsoleUI) Draw() {
	c.DrawFrame()

	c.Draws++
	c.CurrentX = 0
	c.CurrentY = 0
}

func (c *ConsoleUI) AdjustX(x int) {
	c.CurrentX += x
}

func (c *ConsoleUI) AdjustY(y int) {
	c.CurrentY += y
}

var character = make([]string, 4)

var FRAME_W int = 9
var FRAME_H int = 3

func (c *ConsoleUI) DrawFrame() {
	frame := ui.NewPar(strconv.FormatUint(uint64(GameState.FrameNumber), 10))

	frame.X = 0
	frame.Y = c.CurrentY

	frame.Height = FRAME_H
	c.AdjustY(FRAME_H)
	frame.Width = FRAME_W
	c.AdjustX(FRAME_W)

	frame.TextFgColor = ui.ColorWhite
	frame.BorderLabel = "FRAME"
	frame.BorderFg = ui.ColorRed
	frame.BorderLabelFg = ui.ColorCyan
	ui.Render(frame)

	c.DrawStage()
}

var STAGE_W int = 20
var STAGE_H int = 3

func (c *ConsoleUI) DrawStage() {
	stage := ui.NewPar(GetStageName(GameState.Stage))

	stage.X = c.CurrentX
	c.AdjustY(-1 * FRAME_H) // want at same Y position as Frame window
	stage.Y = c.CurrentY

	stage.Width = STAGE_W
	c.AdjustX(STAGE_W)
	stage.Height = STAGE_H
	c.AdjustY(STAGE_H)

	stage.TextFgColor = ui.ColorWhite
	stage.BorderLabel = "STAGE"
	stage.BorderFg = ui.ColorRed
	stage.BorderLabelFg = ui.ColorCyan
	ui.Render(stage)

	c.DrawMenuState()
}

var MENU_STATE_W int = 16
var MENU_STATE_H int = 3

func (c *ConsoleUI) DrawMenuState() {
	menustate := ui.NewPar(GetMenuStateName(GameState.MenuState))

	menustate.X = c.CurrentX
	c.AdjustY(-1 * STAGE_H) // want at same Y position as Stage window
	menustate.Y = c.CurrentY

	menustate.Width = MENU_STATE_W
	c.AdjustX(MENU_STATE_W)
	menustate.Height = MENU_STATE_H
	c.AdjustY(MENU_STATE_H)

	menustate.TextFgColor = ui.ColorWhite
	menustate.BorderLabel = "MENU STATE"
	menustate.BorderFg = ui.ColorRed
	menustate.BorderLabelFg = ui.ColorCyan
	ui.Render(menustate)

	c.DrawAPM()
}

var APM_W int = 10
var APM_H int = 3

func (c *ConsoleUI) DrawAPM() {
	APM := ui.NewPar(fmt.Sprintf("%d", GameState.CalculateAPM()))

	APM.X = c.CurrentX
	c.AdjustY(-1 * STAGE_H) // want at same Y position as Stage window
	APM.Y = c.CurrentY

	APM.Width = APM_W
	c.AdjustX(APM_W)
	APM.Height = APM_H
	c.AdjustY(APM_H)

	APM.TextFgColor = ui.ColorWhite
	APM.BorderLabel = "APM"
	APM.BorderFg = ui.ColorRed
	APM.BorderLabelFg = ui.ColorCyan
	ui.Render(APM)

	c.DrawPlayerTable()
}

var PLAYER_TABLE_W int = 95
var PLAYER_TABLE_H int = 10

func (c *ConsoleUI) DrawPlayerTable() {
	c.CurrentX = 0
	percent := make([]string, 4)
	stock := make([]string, 4)
	//action := make([]Action, 4)

	for i := 1; i < 5; i++ {
		char, _ := GameState.Players[i].GetCharacter()
		if char != 0xFF {
			character[i-1] = GameState.Players[i].GetCharacterString()
		}
		s, _ := GameState.Players[i].GetUint(STOCK)
		p, _ := GameState.Players[i].GetUint(PERCENT)

		stock[i-1] = strconv.FormatUint(uint64(s), 10)
		percent[i-1] = strconv.FormatUint(uint64(p), 10)
		//action[i-1], _ = GameState.Players[i].GetAction()
	}

	rows1 := [][]string{
		[]string{" ", "Player 1", "Player 2", "Player 3", "Player 4"},
		[]string{"[Character](fg-green)"},
		[]string{"[Stocks](fg-green)"},
		[]string{"[Percent](fg-green)"},
		//[]string{"Action"},
	}
	rows1[0][Dolphin.SelfPort] = "[" + rows1[0][Dolphin.SelfPort] + "](fg-red)"

	rows1[1] = append(rows1[1][0:1], character...)
	rows1[2] = append(rows1[2][0:1], stock...)
	rows1[3] = append(rows1[3][0:1], percent...)

	table1 := ui.NewTable()
	table1.Rows = rows1
	table1.FgColor = ui.ColorWhite
	table1.BgColor = ui.ColorDefault
	table1.Analysis()
	table1.FgColors[0] = ui.ColorGreen

	table1.X = 0
	table1.Y = c.CurrentY

	table1.Height = PLAYER_TABLE_H
	c.AdjustY(PLAYER_TABLE_H)

	table1.Width = PLAYER_TABLE_W
	c.AdjustX(PLAYER_TABLE_W)

	table1.SetSize()

	ui.Render(table1)

	c.DrawLog()
}

var LOG_W = 56
var LOG_H = 6

func (c *ConsoleUI) DrawLog() {
	var logText string
	for i := len(c.LogEntries) - 1; i >= 0; i-- {
		logText += c.LogEntries[i] + "\n"
	}

	log := ui.NewPar(logText)
	log.Height = LOG_H
	log.Y = c.CurrentY
	c.AdjustY(LOG_H)
	log.Width = LOG_W
	log.TextFgColor = ui.ColorWhite
	log.BorderLabel = "MATCH SUMMARY"
	log.BorderFg = ui.ColorCyan
	log.BorderLabelFg = ui.ColorGreen

	ui.Render(log)
	c.DrawHelp()
}

var HELP_W int = 56
var HELP_H int = 2

func (c *ConsoleUI) DrawHelp() {
	c.InsertLineBreaks(1)

	par_text := []string{
		"PRESS [q](fg-red) TO QUIT",
		"USE [← →](fg-red) TO CHANGE YOUR PORT. ",
	}
	HELP_H = 2 + len(par_text)
	p := ui.NewPar(strings.Join(par_text, "\n"))
	p.Y = c.CurrentY
	c.AdjustY(HELP_H)
	p.Height = HELP_H
	p.Width = HELP_W
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "HELP"
	p.BorderFg = ui.ColorCyan

	ui.Render(p)
}

func (c *ConsoleUI) LogText(str string) {
	c.LogEntries = append(c.LogEntries, str)
}

func (c *ConsoleUI) ClearLog() {
	c.LogEntries = make([]string, 10)
}

func (c *ConsoleUI) InsertLineBreaks(size int) {
	c.AdjustY(size)
}
