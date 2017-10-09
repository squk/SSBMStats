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
	"sync"
	"syscall"
)

type DolphinTuple struct {
	Address string
	Value   []byte
}

type GameStateManager struct {
	Players     PlayerContainer
	FrameNumber uint32
	Stage       Stage
	MenuState   MenuState
	StageCursorX,
	StageCursorY float32
	Ready bool

	Socket       *net.UnixConn
	SocketBuffer []byte
	SocketMutex  sync.Mutex

	MemoryMap    MemoryMap
	MemoryUpdate chan DolphinTuple
}

func NewGameStateManager() *GameStateManager {
	state := GameStateManager{
		Stage:        FINAL_DESTINATION,
		MenuState:    IN_GAME,
		SocketBuffer: make([]byte, 9096),
		SocketMutex:  sync.Mutex{},
		MemoryMap:    GetMemoryMap(),
		MemoryUpdate: make(chan DolphinTuple),
	}

	state.BindSocket()

	for i := 0; i < 9; i++ {
		state.Players[i] = NewPlayer()
	}

	return &state
}

func (g *GameStateManager) ReadSocket() {
	c := g.Socket
	buf := g.SocketBuffer

	g.SocketMutex.Lock()
	n, err := (*c).Read(buf[:])
	g.SocketMutex.Unlock()

	if err != nil {
		c.Close()
		panic(err)
		return
	}

	s := strings.Split(string(buf[0:n]), "\n")
	padded := fmt.Sprintf("%08s", strings.Replace(s[1], "\x00", "", -1))
	decoded, err := hex.DecodeString(padded)

	if err != nil {
		c.Close()
		panic(err)
		return
	} else {
		g.MemoryUpdate <- DolphinTuple{s[0], decoded}
	}
}

func (g *GameStateManager) BindSocket() {
	p := filepath.Join(Dolphin.MemoryPath, "MemoryWatcher")

	syscall.Unlink(p)
	c, err := net.ListenUnixgram("unixgram", &net.UnixAddr{p, "unixgram"})

	if err != nil {
		panic(err)
	}

	g.Socket = c
}

func (g *GameStateManager) LogFrame() {
	pc := g.Players
	frame := NewFrame(pc, g.MenuState, g.Stage)
	go FWriter.LogFrame(frame)
}

func (g *GameStateManager) Update() {
	go func() {
		for Dolphin.RUNNING {
			go g.ReadSocket()

			m := <-g.MemoryUpdate
			playerIndex := g.MemoryMap[m.Address].PlayerIndex
			state := g.MemoryMap[m.Address].StateID

			g.AssignPlayerValues(playerIndex, state, m.Value)

			if state == FRAME {
				g.LogFrame()
			}

			log.Println("update")
		}
	}()
}

func (g *GameStateManager) AssignPlayerValues(index int, state StateID, value []byte) {
	littleEndianInt := binary.LittleEndian.Uint32(value)
	bigEndianInt := binary.BigEndian.Uint32(value)
	floatVal := math.Float32frombits(binary.BigEndian.Uint32(value))

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
		g.FrameNumber = bigEndianInt
	} else if state == STAGE {
		val := Stage((littleEndianInt >> 8) & 0xFF)
		g.Stage = val
	} else if state == MENU_STATE {
		val := MenuState(bigEndianInt & 0xFF)
		g.MenuState = val
	} else if state == STOCK {
		if index > 0 && index <= 4 {
			g.Players[index].SetUint(STOCK, littleEndianInt)
		}
	} else if state == CHARACTER {
		g.Players[index].SetCharacter(state, Character(bigEndianInt>>24))
	} else if state == ACTION {
		g.Players[index].SetAction(state, Action(bigEndianInt))
		//log.Printf("0x%08X  0x%08X\n", bigEndianInt, littleEndianInt)
	} else if StateSliceContains(floatIDs, state) {
		if state == PERCENT {
			g.Players[index].SetUint(PERCENT, uint32(value[1]))
		} else {
			g.Players[index].SetFloat(state, floatVal)
		}
	} else if StateSliceContains(littleEndianIntIDs, state) {
		g.Players[index].SetUint(state, littleEndianInt)
	} else if StateSliceContains(bigEndianIntIDs, state) {
		g.Players[index].SetUint(state, bigEndianInt)
	}
}

func StateSliceContains(s []StateID, f StateID) bool {
	for _, state := range s {
		if state == f {
			return true
		}
	}
	return false
}
