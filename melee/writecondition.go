package melee

import "log"

type WriteFlag byte

const (
	FALSE WriteFlag = iota
	TRUE
	INDETERMINATE
)

// checks for a condition requiring multiple frames
type BufferCondition func(fs *FrameBuffer) WriteFlag

// checks for a condition only requiring a single frame
type FrameCondition func(f Frame) WriteFlag

// TODO: Make the player index that we check dynamic. Right now we assume the
// player is P1
type FrameValidator struct {
	FrameBuffer *FrameBuffer

	FrameConditions  map[string]FrameCondition
	BufferConditions map[string]BufferCondition
}

func NewFrameValidator(fb *FrameBuffer) FrameValidator {
	f_conditions := make(map[string]FrameCondition)
	b_conditions := make(map[string]BufferCondition)

	f_conditions["L_CANCEL_PASS"] = L_CANCEL_PASS
	f_conditions["L_CANCEL_MISS"] = L_CANCEL_MISS

	fv := FrameValidator{fb, f_conditions, b_conditions}
	return fv
}

func (fv *FrameValidator) AnyPassed() bool {
	current_frame, exists := fv.FrameBuffer.GetCurrent()

	if exists && !current_frame.Empty() {
		for _, fc := range fv.FrameConditions {
			frame, exists := fv.FrameBuffer.GetCurrent()
			if exists {
				state := fc(frame)
				if state == TRUE {
					return true
				}
			}
		}

		for _, bc := range fv.BufferConditions {
			state := bc(fv.FrameBuffer)
			if state == TRUE {
				return true
			}
		}

	}

	return false
}

// A FALSE value from here does not necessarily indicate that the player has
// missed an L-cancel. It only means that their current action can be
// L-cancelled and they have yet to do so.
func L_CANCEL_PASS(f Frame) WriteFlag {
	if !f.Empty() {
		action, _ := f.Players[1].GetAction()
		animation_speed, _ := f.Players[1].GetFloat(SPEED_ANIMATION)
		char, _ := f.Players[1].GetCharacter()

		if CAN_L_CANCEL(action) {
			//log.Printf("%s: %v \n", f.Players[1].GetCharacterString(), animation_speed)

			animation := 0

			if action == BAIR_LANDING {
				animation = BAIR_SPEED
			} else if action == DAIR_LANDING {
				animation = DAIR_SPEED
			} else if action == FAIR_LANDING {
				animation = FAIR_SPEED
			} else if action == NAIR_LANDING {
				animation = NAIR_SPEED
			} else if action == UAIR_LANDING {
				animation = UAIR_SPEED
			}

			l_speed := (LANDING_SPEEDS[char][animation] * 2.0) - CANCEL_TOLERANCE

			if animation_speed >= l_speed {
				log.Println("L_CANCEL_PASSSSSSSS")
				return TRUE
			} else {
				log.Println("L_CANCEL_FAIL")
				return FALSE
			}
		}
	}
	return INDETERMINATE
}

func L_CANCEL_MISS(f Frame) WriteFlag {
	pass := L_CANCEL_PASS(f)

	if pass == TRUE {
		return FALSE
	} else if pass == FALSE {
		return TRUE
	}

	return INDETERMINATE
}

func CAN_L_CANCEL(a Action) bool {
	return (a == UAIR_LANDING || a == BAIR_LANDING || a == NAIR_LANDING || a ==
		DAIR_LANDING || a == FAIR_LANDING)
}
