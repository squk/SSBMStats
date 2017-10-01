package melee

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
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
	Frame                      uint32
	Stage                      Stage
	MenuState                  MenuState
	StageCursorX, StageCursorY float32
	Ready                      bool
	// Players are indexed 1-8(inclusive). Player 0 is used to store non-player
	// values. Slightly un-intuitive but it allows for the great MemoryMap and
	// Player structure designs. Also allows for verbose yet concise printing
	// of each interface because any player that has not been
	// played/initialized has a map size of 0.
	Players [9]Player

	MemoryUpdate chan DolphinTuple
	Socket       *net.UnixConn

	SocketBuffer []byte
	MemoryMap    MemoryMap
	//ActionData   []ActionData
}

func NewGameState(d *Dolphin) GameState {
	state := GameState{
		DolphinInstance: d,
		Stage:           FINAL_DESTINATION,
		MenuState:       CHARACTER_SELECT,
		MemoryUpdate:    make(chan DolphinTuple),
		MemoryMap:       GetMemoryMap(),
		//ActionData:      GetActionData(),
		SocketBuffer: make([]byte, 9096),
	}

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
			//TODO: Implement logging!
			continue
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

			playerIndex := g.MemoryMap[m.Address].PlayerIndex
			state := g.MemoryMap[m.Address].StateID
			g.AssignPlayerValues(playerIndex, state, m.Value)

			if g.MemoryMap[m.Address].StateID == FRAME {
				newFrame <- true
				g.FixFrameIndexing()
			}
		}
	}()
}

func (g *GameState) AssignPlayerValues(index int, state StateID, value []byte) {
	littleEndianInt := binary.LittleEndian.Uint32(value)
	bigEndianInt := binary.BigEndian.Uint32(value)
	floatVal := math.Float32frombits(bigEndianInt)

	floatIDs := []StateID{
		PERCENT, X, Y, CURSOR_X, CURSOR_Y, ACTION_FRAME, HITLAG_FRAMES_LEFT,
		HITSTUN_FRAMES_LEFT, HITBOX_1_X, HITBOX_1_Y, HITBOX_2_X, HITBOX_2_Y,
		HITBOX_3_X, HITBOX_3_Y, HITBOX_4_X, HITBOX_4_Y, HITBOX_1_SIZE,
		HITBOX_2_SIZE, HITBOX_3_SIZE, HITBOX_4_SIZE, SPEED_ANIMATION,
	}

	littleEndianIntIDs := []StateID{
		STOCK, INVULNERABLE, CHARGING_SMASH, ON_GROUND,
		COIN_DOWN, HITBOX_1_STATUS, HITBOX_2_STATUS, HITBOX_3_STATUS,
		HITBOX_4_STATUS, IASA, TRANSFORMED, IS_ZELDA,
	}

	bigEndianIntIDs := []StateID{
		READY_TO_START, CONTROLLER_STATUS, ACTION, CHARACTER,
	}

	if state == FRAME {
		g.Frame = littleEndianInt
		//g.Players[index].SetUint(state, littleEndianInt)
	} else if state == STAGE {
		val := Stage((littleEndianInt >> 16) & 0x000000ff)
		//g.Players[index].SetStage(state, val)
		g.Stage = val
	} else if state == MENU_STATE {
		val := MenuState(littleEndianInt & 0x000000ff)
		//g.Players[index].SetMenuState(state, val)
		g.MenuState = val
	} else if state == CHARACTER {
		g.Players[index].SetCharacter(state, Character(bigEndianInt>>24))
	} else if state == ACTION {
		g.Players[index].SetAction(state, Action(bigEndianInt))
	} else if StateSliceContains(floatIDs, state) {
		g.Players[index].SetFloat(state, floatVal)
	} else if StateSliceContains(littleEndianIntIDs, state) {
		g.Players[index].SetUint(state, littleEndianInt)
	} else if StateSliceContains(bigEndianIntIDs, state) {
		g.Players[index].SetUint(state, bigEndianInt)
	}
}

// normalizes all Player structs ACTION_FRAME to be indexed from 1
func (g *GameState) FixFrameIndexing() {
}

func StateSliceContains(s []StateID, f StateID) bool {
	for _, state := range s {
		if state == f {
			return true
		}
	}
	return false
}
