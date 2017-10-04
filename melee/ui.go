package melee

import (
	"strconv"
	"sync"

	ui "github.com/gizak/termui"
)

type CUI struct {
	LogEntries      []string
	DolphinInstance *Dolphin
	Draws           uint64
}

var instance *CUI
var once sync.Once

func NewCUI() *CUI {
	return &CUI{[]string{" ", " ", " "}, nil, 0}
}

func (c *CUI) Draw() {
	c.DrawFrame()
	c.DrawLog()
	c.DrawPlayerTable()
}

func (c *CUI) DrawPlayerTable() {
	character := make([]string, 4)
	percent := make([]string, 4)
	stock := make([]string, 4)
	//action := make([]Action, 4)

	for i := 1; i < 5; i++ {
		character[i-1] = c.DolphinInstance.GameState.Players[i].GetCharacterString()
		s, _ := c.DolphinInstance.GameState.Players[i].GetUint(STOCK)
		p, _ := c.DolphinInstance.GameState.Players[i].GetUint(PERCENT)

		stock[i-1] = strconv.FormatUint(uint64(s), 10)
		percent[i-1] = strconv.FormatUint(uint64(p), 10)
		//action[i-1], _ = c.GameState.Players[i].GetAction()
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
	table1.Y = 10
	table1.X = 0
	table1.Height = 10
	table1.Width = 95
	table1.SetSize()

	ui.Render(table1)
}

func (c *CUI) DrawLog() {
	var logText string
	for i := len(c.LogEntries) - 1; i >= 0; i-- {
		logText += c.LogEntries[i] + "\n"
	}

	log := ui.NewPar(logText)
	log.Height = 6
	log.Width = 50
	log.TextFgColor = ui.ColorWhite
	log.BorderLabel = "Log"
	log.BorderFg = ui.ColorCyan

	ui.Render(log)
}

func (c *CUI) DrawFrame() {
	frame := ui.NewPar(strconv.FormatUint(uint64(c.DolphinInstance.GameState.FrameNumber), 10))
	c.Draws++

	frame.Height = 6
	frame.Width = 10
	frame.X = 55
	frame.TextFgColor = ui.ColorWhite
	frame.BorderLabel = "Frame"
	frame.BorderFg = ui.ColorRed
	ui.Render(frame)
}

func (c *CUI) LogText(str string) {
	c.LogEntries = append(c.LogEntries, str)
}

func (c *CUI) ClearLog() {
	c.LogEntries = make([]string, 10)
}
