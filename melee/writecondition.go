package melee

/*
*  Any check/validation of Frame data that may result in a write to disk should
*  exist as a FrameCondition or BufferCondition.
 */

type WriteAuth byte

const (
	FALSE WriteAuth = iota
	TRUE
	INDETERMINATE
)

// checks for a condition requiring multiple frames
type BufferCondition func(fb *FrameBuffer) (WriteAuth, DiskFrame)

// checks for a condition only requiring a single frame
type FrameCondition func(f Frame) (WriteAuth, DiskFrame)

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
			// WriteAuth is set to TRUE
			if ok, data := fc(frame); ok == TRUE {
				FWriter.Write(data)
			}
		}

		for _, bc := range fv.BufferConditions {
			if ok, _ := bc(fv.FrameBuffer); ok == TRUE {

			}
		}

	}
}
