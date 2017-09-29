package melee

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

type MenuState int

const (
	CHARACTER_SELECT MenuState = 0
	STAGE_SELECT               = 1
	IN_GAME                    = 2
	POSTGAME_SCORES            = 4
)

type DolphinTuple struct {
	label  string
	player []byte
}

type GameState struct {
	DolphinInstance            *Dolphin
	Frame                      int
	Stage                      Stage
	MenuState                  MenuState
	StageCursorX, StageCursorY float32
	Ready                      bool
	ProcessingTime             time.Time
	FrameTimestamp             time.Time
	Players                    [8]PlayerState
	Memory                     chan DolphinTuple
	Locations                  map[int]int
	Socket                     *net.Conn
	ReadingSocket              bool
}

func (g *GameState) ReadSocket() {
	c := g.Socket

	buf := make([]byte, 9096)
	g.ReadingSocket = true

	for g.ReadingSocket {
		n, err := (*c).Read(buf[:])
		if err != nil {
			return
		}

		s := strings.Split(string(buf[0:n]), "\n")
		g.Memory <- DolphinTuple{s[0], []byte(s[1])}
	}
}

func (g *GameState) BindSocket() {
	p := filepath.Join(g.DolphinInstance.MemoryPath, "MemoryWatcher")
	c, err := net.Dial("unixgram", p)

	if err != nil {
		syscall.Unlink(p)

		fmt.Println(err)
		fmt.Println("Could not connect to existing socket. Creating new one. ")

		// socket does not exist, open it
		c, err = net.ListenUnixgram("unixgram", &net.UnixAddr{p, "unixgram"})

		if err != nil {
			panic(err)
		}
	}
	c.SetDeadline(time.Now().Add(time.Second * 5))
	g.Socket = &c
}

func NewGameState(d *Dolphin) GameState {
	state := GameState{
		DolphinInstance: d,
		Stage:           FINAL_DESTINATION,
		MenuState:       CHARACTER_SELECT,
		Memory:          make(chan DolphinTuple),
	}
	state.BindSocket()

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//log.Fatal(err)
	//}
	//f, err := os.Create(filepath.Join(dir, "locations.dat"))
	//defer f.Close()

	state.Players[0] = NewPlayerState()
	state.Players[1] = NewPlayerState()
	state.Players[2] = NewPlayerState()
	state.Players[3] = NewPlayerState()
	state.Players[4] = NewPlayerState()
	state.Players[5] = NewPlayerState()
	state.Players[6] = NewPlayerState()
	state.Players[7] = NewPlayerState()

	return state
}

func (g *GameState) Step() {
	//g.ProcessingTime = time.Now().Sub(g.FrameTimestamp)
	// for mem, _ := range ?
	if !g.ReadingSocket {
		g.ReadSocket()
	}

	g.Update()
}

func (g *GameState) FixFrameIndexing() {
	//for i, p := range g.Players {

	//}
}
func (g *GameState) FixIASA() {
	for _, p := range g.Players {
		if p.Action < NEUTRAL_ATTACK_1 || p.Action > DAIR {
			p.IASA = false
		}
	}
}

func (g *GameState) Update() bool {
	go func() {
		m := <-g.Memory

		//label := Locations
		fmt.Println(m, "\n")
	}()

	return true
}
