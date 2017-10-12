package melee

type WriteFlag byte

const (
	FALSE WriteFlag = iota
	TRUE
	INDETERMINATE
)

// checks for a condition requiring multiple frames
type BufferCondition func(fs *FrameBuffer) (WriteFlag, DiskFrame)

// checks for a condition only requiring a single frame
type FrameCondition func(f Frame) (WriteFlag, DiskFrame)

type FrameValidator struct {
	FrameBuffer *FrameBuffer

	FrameConditions  map[string]FrameCondition
	BufferConditions map[string]BufferCondition
}

func NewFrameValidator(fb *FrameBuffer) FrameValidator {
	f_conditions := make(map[string]FrameCondition)
	b_conditions := make(map[string]BufferCondition)

	f_conditions["L_CANCEL_PASS"] = L_CANCEL_PASS

	fv := FrameValidator{fb, f_conditions, b_conditions}
	return fv
}

func (fv *FrameValidator) CheckAll() {
	frame, exists := fv.FrameBuffer.GetCurrent()

	if exists && !frame.Empty() {
		for _, fc := range fv.FrameConditions {
			// a Frame will only be logged by a FrameCondition of the
			// WriteFlag is set to TRUE
			if ok, data := fc(frame); ok == TRUE {
				FWriter.Write(data)
			}
		}

		//for _, bc := range fv.BufferConditions {
		//if ok, _ := bc(fv.FrameBuffer); ok == TRUE {

		//}
		//}

	}
}

var L_CANCEL_OPPORTUNE bool // prevents multiple frames being marked as a successful L-Cancel

// returns a DiskFrame and whether or not that DiskFrame should be logged
func L_CANCEL_PASS(f Frame) (flag WriteFlag, data DiskFrame) {
	flag = INDETERMINATE

	if !f.Empty() {
		action, _ := f.Players[GameState.SelfPort].GetAction()

		if !L_CANCEL_OPPORTUNE {
			animation_speed, _ := f.Players[GameState.SelfPort].GetFloat(SPEED_ANIMATION)
			char, _ := f.Players[GameState.SelfPort].GetCharacter()

			if CAN_L_CANCEL(action) {
				L_CANCEL_OPPORTUNE = true
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

				var msg string
				if animation_speed >= l_speed {
					flag = TRUE
					msg = "L_CANCEL_PASS"
				} else {
					flag = TRUE
					msg = "L_CANCEL_MISS"
				}

				data = DiskFrame{
					BasicFrame: BasicFrame{
						GameState.FrameNumber,
						ACTION_NAMES[action],
						msg,
					},
					DetailedFrame: Frame{},
				}
			}
		} else {
			if !CAN_L_CANCEL(action) {
				L_CANCEL_OPPORTUNE = false
			}
		}
	}

	return flag, data
}

func CAN_L_CANCEL(a Action) bool {
	return (a == UAIR_LANDING || a == BAIR_LANDING || a == NAIR_LANDING || a ==
		DAIR_LANDING || a == FAIR_LANDING)
}
