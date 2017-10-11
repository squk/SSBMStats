package melee

import (
	"errors"
	"log"
	"os"
	"time"
)

var Dolphin *DolphinManager
var CUI *ConsoleUI
var GameState *GameStateManager
var FWriter *FrameWriter

func init() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetFlags(log.Lshortfile)
	log.SetOutput(f)

	Dolphin = NewDolphinManager()
	CUI = NewConsoleUI()
	GameState = NewGameStateManager()
	FWriter = NewFrameWriter()

	go func() {
		for Dolphin.RUNNING {
			time.Sleep(200 * time.Millisecond)
			CUI.Draw()
		}
	}()
}

func Init() (err error) {
	//if Dolphin.DolphinPath == "" {
	//exists := false

	//for !exists {
	//reader := bufio.NewReader(os.Stdin)
	////fmt.Print("Enter Dolphin Path: ")
	//text, _ := reader.ReadString('\n')
	//text = strings.TrimSuffix(strings.TrimSuffix(text, "\n"), " ")
	//exe_name := "Dolphin"

	//if runtime.GOOS == "windows" {
	//exe_name += ".exe"
	//} else {
	//exe_name += ".app"
	//}
	////user_exists, _ := FilepathExists(filepath.Join(text, "User"))
	////exists = user_exists || exists
	//exists, _ = FilepathExists(text)

	//if !exists {
	////fmt.Println("\nInvalid Dolphin path")
	//}
	//finalPath := filepath.Join(text, "User")
	//Dolphin.SetPath(finalPath)
	//}
	//}

	if !Dolphin.Init() {
		err = errors.New("DolphinManager failed to initialize")
	}
	return err
}
