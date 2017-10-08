package melee

import (
	"strconv"
	"sync"

	ui "github.com/gizak/termui"
)

type ConsoleUI struct {
	LogEntries         []string
	Draws              uint64
	CurrentX, CurrentY int
}

var once sync.Once

func NewConsoleUI() *ConsoleUI {
	return &ConsoleUI{[]string{" ", " ", " "}, 0, 0, 0}
}

func (c *ConsoleUI) Draw() {
	c.DrawLog()

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

var LOG_W = 50
var LOG_H = 6

func (c *ConsoleUI) DrawLog() {
	var logText string
	for i := len(c.LogEntries) - 1; i >= 0; i-- {
		logText += c.LogEntries[i] + "\n"
	}

	log := ui.NewPar(logText)
	log.Height = LOG_H
	c.AdjustY(LOG_H)

	log.Width = LOG_W
	log.TextFgColor = ui.ColorWhite
	log.BorderLabel = "Log"
	log.BorderFg = ui.ColorRed
	log.BorderLabelFg = ui.ColorCyan

	ui.Render(log)
	c.DrawFrame()
}

var FRAME_W int = 9
var FRAME_H int = 3

func (c *ConsoleUI) DrawFrame() {
	frame := ui.NewPar(strconv.FormatUint(uint64(GameState.FrameNumber), 10))

	frame.X = 0
	frame.Y = c.CurrentY

	frame.Height = FRAME_H
	c.AdjustY(FRAME_H)
	frame.Width = FRAME_W

	frame.TextFgColor = ui.ColorWhite
	frame.BorderLabel = "Frame"
	frame.BorderFg = ui.ColorRed
	frame.BorderLabelFg = ui.ColorCyan
	ui.Render(frame)

	c.DrawStage()
}

var STAGE_W int = 21
var STAGE_H int = 3

func (c *ConsoleUI) DrawStage() {
	stage := ui.NewPar(GetStageName(GameState.Stage))

	stage.X = FRAME_W
	c.AdjustY(-1 * FRAME_H) // want at same Y position as Frame window
	stage.Y = c.CurrentY

	stage.Width = STAGE_W
	stage.Height = STAGE_H
	c.AdjustY(STAGE_H)

	stage.TextFgColor = ui.ColorWhite
	stage.BorderLabel = "Stage"
	stage.BorderFg = ui.ColorRed
	stage.BorderLabelFg = ui.ColorCyan
	ui.Render(stage)

	c.DrawMenuState()
}

var MENU_STATE_W int = 20
var MENU_STATE_H int = 3

func (c *ConsoleUI) DrawMenuState() {
	menustate := ui.NewPar(GetMenuStateName(GameState.MenuState))

	menustate.X = FRAME_W + STAGE_W
	c.AdjustY(-1 * STAGE_H) // want at same Y position as Stage window
	menustate.Y = c.CurrentY

	menustate.Width = MENU_STATE_W
	menustate.Height = MENU_STATE_H
	c.AdjustY(MENU_STATE_H)

	menustate.TextFgColor = ui.ColorWhite
	menustate.BorderLabel = "Menu State"
	menustate.BorderFg = ui.ColorRed
	menustate.BorderLabelFg = ui.ColorCyan
	ui.Render(menustate)

	c.DrawPlayerTable()
}

var PLAYER_TABLE_W int = 95
var PLAYER_TABLE_H int = 10

func (c *ConsoleUI) DrawPlayerTable() {
	c.InsertLineBreaks(2)

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
		[]string{"Character"},
		[]string{"Stocks"},
		[]string{"Percent"},
		//[]string{"Action"},
	}

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
