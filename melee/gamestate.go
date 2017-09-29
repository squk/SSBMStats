package melee

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"net"
	"path/filepath"
	"strings"
	"syscall"
)

type MenuState int

const (
	CHARACTER_SELECT MenuState = 0
	STAGE_SELECT               = 1
	IN_GAME                    = 2
	POSTGAME_SCORES            = 4
)

type DolphinTuple struct {
	Address string
	Value   []byte
}

type GameState struct {
	DolphinInstance            *Dolphin
	Frame                      int
	Stage                      Stage
	MenuState                  MenuState
	StageCursorX, StageCursorY float32
	Ready                      bool
	// Players are indexed 1-8(inclusive). Player 0 is used to store non-player
	// values. Slightly un-intuitive but it allows for the great MemoryMap and
	// Player structure designs. Also allows for verbose yet concise printing
	// of each interface because any player that has not been
	// played/initialized has a map size of 0.
	Players      [9]Player
	MemoryUpdate chan DolphinTuple
	Socket       *net.UnixConn

	SocketBuffer []byte
	MemoryMap    MemoryMap
}

func NewGameState(d *Dolphin) GameState {
	state := GameState{
		DolphinInstance: d,
		Stage:           FINAL_DESTINATION,
		MenuState:       CHARACTER_SELECT,
		MemoryUpdate:    make(chan DolphinTuple),
	}
	state.MemoryMap = GetMemoryMap()
	state.SocketBuffer = make([]byte, 9096)
	state.BindSocket()

	for i := 0; i < 9; i++ {
		state.Players[i] = NewPlayer()
	}

	return state
}

func (g *GameState) ReadSocket() {
	c := g.Socket
	buf := g.SocketBuffer

	for g.DolphinInstance.RUNNING {
		n, err := (*c).Read(buf[:])
		if err != nil {
			// TODO: Log this.
			panic(err)
		}

		s := strings.Split(string(buf[0:n]), "\n")
		padded := fmt.Sprintf("%08s", strings.Replace(s[1], "\x00", "", -1))
		decoded, err := hex.DecodeString(padded)

		if err != nil {
			log.Fatal(err)
		} else {
			g.MemoryUpdate <- DolphinTuple{s[0], decoded}
		}
	}
}

func (g *GameState) BindSocket() {
	p := filepath.Join(g.DolphinInstance.MemoryPath, "MemoryWatcher")

	syscall.Unlink(p)
	c, err := net.ListenUnixgram("unixgram", &net.UnixAddr{p, "unixgram"})

	if err != nil {
		panic(err)
	}

	g.Socket = c
}

func (g *GameState) Update(newFrame chan<- bool) {
	go g.ReadSocket()

	go func() {
		for g.DolphinInstance.RUNNING {
			m := <-g.MemoryUpdate

			if g.MemoryMap[m.Address].StateID == FRAME {
				newFrame <- true
			}
			playerIndex := g.MemoryMap[m.Address].PlayerIndex
			attribute := g.MemoryMap[m.Address].StateID
			g.AssignPlayerValue(playerIndex, attribute, m.Value)

			//fmt.Println(g.Players[playerIndex])
		}
	}()
}

func (g *GameState) AssignPlayerValue(index int, state StateID, value []byte) {
	littleEndianInt := binary.LittleEndian.Uint32(value)
	bigEndianInt := binary.BigEndian.Uint32(value)
	floatVal := math.Float32frombits(littleEndianInt)

	floatIDs := []StateID{
		PERCENT, X, Y, CURSOR_X, CURSOR_Y, ACTION_FRAME, HITLAG_FRAMES_LEFT,
		HITSTUN_FRAMES_LEFT, HITBOX_1_X, HITBOX_1_Y, HITBOX_2_X, HITBOX_2_Y,
		HITBOX_3_X, HITBOX_3_Y, HITBOX_4_X, HITBOX_4_Y, HITBOX_1_SIZE,
		HITBOX_2_SIZE, HITBOX_3_SIZE, HITBOX_4_SIZE,
	}

	littleEndianIntIDs := []StateID{
		STOCK, ACTION, CHARACTER, INVULNERABLE, CHARGING_SMASH, ON_GROUND,
		COIN_DOWN, HITBOX_1_STATUS, HITBOX_2_STATUS, HITBOX_3_STATUS,
		HITBOX_4_STATUS, IASA, TRANSFORMED, ISZELDA,
	}

	bigEndianIntIDs := []StateID{
		READY_TO_START, CONTROLLER_STATUS,
	}

	if state == FRAME {
		g.Players[index].Values[state] = littleEndianInt
	} else if state == STAGE {
		val := Stage((littleEndianInt >> 16) & 0x000000ff)
		g.Players[index].Values[state] = val
		g.Stage = val
	} else if state == MENU_STATE {
		val := MenuState(littleEndianInt & 0x000000ff)
		g.Players[index].Values[state] = val
		g.MenuState = val
	} else if StateSliceContains(floatIDs, state) {
		g.Players[index].Values[state] = floatVal
	} else if StateSliceContains(littleEndianIntIDs, state) {
		g.Players[index].Values[state] = littleEndianInt
	} else if StateSliceContains(bigEndianIntIDs, state) {
		g.Players[index].Values[state] = bigEndianInt
	}
}

func StateSliceContains(s []StateID, e StateID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
