package melee

import (
	"encoding/binary"
	"math"
)

type MeleeState struct {
	Players     PlayerContainer
	FrameNumber uint32
	Stage       Stage

	MenuState     MenuState
	LastMenuState MenuState

	StageCursorX,
	StageCursorY float32
	Ready bool

	MemoryMap MemoryMap
	APMStore  [120]int

	SelfPort      int
	OpponentPorts []int
}

func NewMeleeState() *MeleeState {
	state := MeleeState{
		Stage:     FINAL_DESTINATION,
		MenuState: IN_GAME,
		MemoryMap: GetMemoryMap(),
		SelfPort:  1,
	}

	for i := 0; i < 9; i++ {
		state.Players[i] = NewPlayer()
	}

	return &state
}

func (g *MeleeState) Update() {
	go func() {
		for Dolphin.Looping {
			go Dolphin.ReadSocket()

			m := <-Dolphin.MemoryUpdate

			playerIndex := g.MemoryMap[m.Address].PlayerIndex
			state := g.MemoryMap[m.Address].StateID

			g.AssignPlayerValues(playerIndex, state, m.Value)

			if state == FRAME {
				g.OnFrame()
			}
		}
	}()
}

func (g *MeleeState) AssignPlayerValues(index int, state StateID, value []byte) {
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
		READY_TO_START, CONTROLLER_STATUS, PLAYER_ACTION, CHARACTER, CONTROLLER_DATA,
	}

	if state == FRAME {
		g.FrameNumber = bigEndianInt
		g.Players[0].SetUint(FRAME, bigEndianInt)
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
		g.OnCharacter()
	} else if state == PLAYER_ACTION {
		g.Players[index].SetPlayerAction(bigEndianInt)
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

func (g *MeleeState) SetMenuState(state MenuState) {
	if state == CHARACTER_SELECT || state == POSTGAME_SCORES {
		FWriter.Flush()
	}
	g.MenuState = state
}

func (g *MeleeState) SetStage(state Stage) {
	g.Stage = state
	FWriter.Match.Stage = STAGE_NAMES[state]
}

func StateSliceContains(s []StateID, f StateID) bool {
	for _, state := range s {
		if state == f {
			return true
		}
	}
	return false
}

func (g *MeleeState) IncreasePort() {
	g.SelfPort++

	if g.SelfPort > 4 {
		g.SelfPort = 1
	}
}

func (g *MeleeState) DecreasePort() {
	g.SelfPort--

	if g.SelfPort < 1 {
		g.SelfPort = 4
	}
}

func (g *MeleeState) StopLoop() {
	Dolphin.Looping = false
}
