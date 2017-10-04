package melee

import (
	"errors"
	"sync"
)

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
	ValuesMutex sync.Mutex
}

func NewPlayer() Player {
	p := Player{
		Values:      make(map[StateID]interface{}),
		ValuesMutex: sync.Mutex{},
	}
	return p
}

type PlayerState struct {
	StateID     StateID
	PlayerIndex int
}

// Setters. Used for R/W locking
func (p *Player) SetFloat(state StateID, val float32) {
	p.ValuesMutex.Lock()
	p.Values[state] = val
	p.ValuesMutex.Unlock()
}
func (p *Player) SetUint(state StateID, val uint32) {
	p.ValuesMutex.Lock()
	p.Values[state] = val
	p.ValuesMutex.Unlock()
}
func (p *Player) SetStage(state StateID, val Stage) {
	p.ValuesMutex.Lock()
	p.Values[state] = Stage(val)
	p.ValuesMutex.Unlock()
}

func (p *Player) SetMenuState(state StateID, val MenuState) {
	p.ValuesMutex.Lock()
	p.Values[state] = MenuState(val)
	p.ValuesMutex.Unlock()
}
func (p *Player) SetCharacter(state StateID, val Character) {
	p.ValuesMutex.Lock()
	p.Values[state] = Character(val)
	p.ValuesMutex.Unlock()
}
func (p *Player) SetAction(state StateID, val Action) {
	p.ValuesMutex.Lock()
	p.Values[state] = Action(val)
	p.ValuesMutex.Unlock()
}

// Getters
func (p *Player) GetFloat(state StateID) (float32, error) {
	var ret float32
	var err error

	p.ValuesMutex.Lock()
	if val, ok := p.Values[state].(float32); ok {
		ret = val
		err = nil
	} else {
		ret = 0.0
		err = errors.New("Cannot assert the provided StateID to float32")
	}
	p.ValuesMutex.Unlock()

	return ret, err
}

func (p *Player) GetUint(state StateID) (uint32, error) {
	var ret uint32
	var err error

	p.ValuesMutex.Lock()
	if val, ok := p.Values[state].(uint32); ok {
		ret = val
		err = nil
	} else {
		ret = 0x0
		err = errors.New("Cannot assert the provided StateID to uint32")
	}
	p.ValuesMutex.Unlock()

	return ret, err
}

func (p *Player) GetAction() (Action, error) {
	var ret Action
	var err error

	p.ValuesMutex.Lock()
	if val, ok := p.Values[ACTION].(Action); ok {
		ret = val
		err = nil
	} else {
		ret = UNKNOWN_ANIMATION
		err = errors.New("Cannot assert the interface at the ACTION index to an Action. Invalid assignment?")
	}
	p.ValuesMutex.Unlock()

	return ret, err
}

func (p *Player) GetCharacter() (Character, error) {
	var ret Character
	var err error

	p.ValuesMutex.Lock()
	if val, ok := p.Values[CHARACTER].(Character); ok {
		ret = val
		err = nil
	} else {
		ret = UNKNOWN_CHARACTER
		err = errors.New("Cannot assert the interface at the CHARACTER index to a Character. Invalid assignment?")
	}
	p.ValuesMutex.Unlock()

	return ret, err
}

func (p *Player) GetCharacterString() string {
	c, _ := p.GetCharacter()
	return CharacterNames[c]
}
