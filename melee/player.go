package melee

import (
	"errors"
	"sync"
)

// indices for PlayerContainer
const (
	STAGE_INFO int = iota
	PLAYER1
	PLAYER2
	PLAYER3
	PLAYER4
	PLAYER1_TRANSFORM
	PLAYER2_TRANSFORM
	PLAYER3_TRANSFORM
	PLAYER4_TRANSFORM
)

// Players are indexed 1-8(inclusive). Player 0 is used to store non-player
// values. Slightly un-intuitive but it allows for the great MemoryMap and
// Player structure designs. Also allows for verbose yet concise printing
// of each interface because any player that has not been
// played/initialized has a map size of 0.
type PlayerContainer [9]Player

type Player struct {
	Values      map[StateID]interface{}
	valuesMutex sync.Mutex
}

func NewPlayer() Player {
	p := Player{
		Values:      make(map[StateID]interface{}),
		valuesMutex: sync.Mutex{},
	}
	return p
}

type PlayerState struct {
	StateID     StateID
	PlayerIndex int
}

// Setters. Used for R/W locking
func (p *Player) SetFloat(state StateID, val float32) {
	p.valuesMutex.Lock()
	p.Values[state] = val
	p.valuesMutex.Unlock()
}
func (p *Player) SetUint(state StateID, val uint32) {
	p.valuesMutex.Lock()
	p.Values[state] = val
	p.valuesMutex.Unlock()
}
func (p *Player) SetStage(val uint32) {
	p.valuesMutex.Lock()
	p.Values[STAGE] = Stage(val)
	p.valuesMutex.Unlock()
}

func (p *Player) SetMenuState(val uint32) {
	p.valuesMutex.Lock()
	p.Values[MENU_STATE] = MenuState(val)
	p.valuesMutex.Unlock()
}
func (p *Player) SetCharacter(val uint32) {
	p.valuesMutex.Lock()
	p.Values[CHARACTER] = Character(val)
	p.valuesMutex.Unlock()
}
func (p *Player) SetAction(val uint32) {
	p.valuesMutex.Lock()
	p.Values[ACTION] = Action(val)
	p.valuesMutex.Unlock()
}
func (p *Player) SetController(val uint32) {
	p.valuesMutex.Lock()
	p.Values[CONTROLLER_DATA] = ControllerBits(uint32(val))
	p.valuesMutex.Unlock()
}
func (p *Player) SetControllerPrevious(val uint32) {
	p.valuesMutex.Lock()
	p.Values[CONTROLLER_DATA_PREVIOUS] = ControllerBits(uint32(val))
	p.valuesMutex.Unlock()
}

// Getters
func (p *Player) GetFloat(state StateID) (ret float32, err error) {
	p.valuesMutex.Lock()
	if val, ok := p.Values[state].(float32); ok {
		ret = val
		err = nil
	} else {
		ret = 0.0
		err = errors.New("Cannot assert the provided StateID to float32")
	}
	p.valuesMutex.Unlock()

	return
}

func (p *Player) GetUint(state StateID) (ret uint32, err error) {
	p.valuesMutex.Lock()
	if val, ok := p.Values[state].(uint32); ok {
		ret = val
		err = nil
	} else {
		ret = 0x0
		err = errors.New("Cannot assert the provided StateID to uint32")
	}
	p.valuesMutex.Unlock()

	return
}

func (p *Player) GetAction() (ret Action, err error) {
	p.valuesMutex.Lock()
	if val, ok := p.Values[ACTION].(Action); ok {
		ret = val
		err = nil
	} else {
		ret = UNKNOWN_ANIMATION
		err = errors.New("Cannot assert the interface at the ACTION index to an Action. Invalid assignment?")
	}
	p.valuesMutex.Unlock()

	return
}

func (p *Player) GetCharacter() (ret Character, err error) {
	p.valuesMutex.Lock()
	if val, ok := p.Values[CHARACTER].(Character); ok {
		ret = val
		err = nil
	} else {
		ret = UNKNOWN_CHARACTER
		err = errors.New("Cannot assert the interface at the CHARACTER index to a Character. Invalid assignment?")
	}
	p.valuesMutex.Unlock()

	return
}

func (p *Player) GetController() (ret ControllerBits) {
	p.valuesMutex.Lock()
	if val, ok := p.Values[CONTROLLER_DATA].(ControllerBits); ok {
		ret = ControllerBits(val)
	} else {
		ret = ControllerBits(0)
	}
	p.valuesMutex.Unlock()

	return
}

func (p *Player) GetControllerPrevious() (ret ControllerBits) {
	p.valuesMutex.Lock()
	if val, ok := p.Values[CONTROLLER_DATA_PREVIOUS].(ControllerBits); ok {
		ret = ControllerBits(val)
	} else {
		ret = ControllerBits(0)
	}
	p.valuesMutex.Unlock()

	return
}

func (p *Player) GetCharacterString() string {
	c, _ := p.GetCharacter()
	return CharacterNames[c]
}
