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
	APMStore     [120]int
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
		log.Fatalln(err)
		return
	}

	s := strings.Split(string(buf[0:n]), "\n")
	padded := fmt.Sprintf("%08s", strings.Replace(s[1], "\x00", "", -1))
	decoded, err := hex.DecodeString(padded)

	if err != nil {
		c.Close()
		log.Fatalln(err)
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
		log.Fatalln(err)
	}

	g.Socket = c
}

var apmIndex int = 0

func (g *GameStateManager) CalculateAPM() int {
	sum := 0
	for _, n := range g.APMStore {
		sum += n
	}
	// 120 frames = 2 seconds.
	return sum
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
				g.OnFrame()
			}
		}
	}()
}

func (g *GameStateManager) OnFrame() {
	//if apmIndex >= 120 {
	////log.Println(g.APMStore)
	//apmIndex = 0
	//}
	//if g.Players[Dolphin.SelfPort].GetController() != g.Players[Dolphin.SelfPort].GetControllerPrevious() {
	//g.APMStore[apmIndex] = g.Players[Dolphin.SelfPort].GetController().Count()
	//}
	//log.Printf("%x %x\n", g.Players[Dolphin.SelfPort].GetController(), g.Players[Dolphin.SelfPort].GetControllerPrevious())
	//apmIndex++

	player_container := g.Players

	frame := NewFrame(player_container)
	if g.MenuState == IN_GAME {
		go FWriter.LogFrame(frame)
	}
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
		READY_TO_START, CONTROLLER_STATUS, ACTION, CHARACTER, CONTROLLER_DATA,
	}

	if state == FRAME {
		g.FrameNumber = bigEndianInt
	} else if state == STAGE {
		g.SetStage(Stage((littleEndianInt >> 8) & 0xFF))
	} else if state == MENU_STATE {
		g.SetMenuState(MenuState(bigEndianInt & 0xFF))
	} else if state == STOCK {
		if index > 0 && index <= 4 {
			g.Players[index].SetUint(STOCK, littleEndianInt)
		}
	} else if state == CHARACTER {
		g.Players[index].SetCharacter(bigEndianInt >> 24)

		if index == Dolphin.SelfPort {
			FWriter.Match.SelfCharacter = g.Players[index].GetCharacterString()
		}
	} else if state == ACTION {
		g.Players[index].SetAction(bigEndianInt)
	} else if state == PERCENT {
		g.Players[index].SetUint(PERCENT, uint32(value[1]))
	} else if state == CONTROLLER_DATA {
		g.Players[index].SetController(bigEndianInt)
	} else if state == CONTROLLER_DATA_PREVIOUS {
		g.Players[index].SetControllerPrevious(bigEndianInt)
	} else if StateSliceContains(floatIDs, state) {
		g.Players[index].SetFloat(state, floatVal)
	} else if StateSliceContains(littleEndianIntIDs, state) {
		g.Players[index].SetUint(state, littleEndianInt)
	} else if StateSliceContains(bigEndianIntIDs, state) {
		g.Players[index].SetUint(state, bigEndianInt)
	}
}

func (g *GameStateManager) SetMenuState(state MenuState) {
	if state == CHARACTER_SELECT || state == POSTGAME_SCORES {
		FWriter.Flush()
	}
	g.MenuState = state
}

func (g *GameStateManager) SetStage(state Stage) {
	g.Stage = state
	FWriter.Match.Stage = GetStageName(state)
}

func StateSliceContains(s []StateID, f StateID) bool {
	for _, state := range s {
		if state == f {
			return true
		}
	}
	return false
}
