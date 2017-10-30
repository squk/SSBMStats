package melee

type ApproachData struct {
	StartingFrame  uint32 `json:"starting_frame"`
	SelfAction     string `json:"self_action"`
	OpponentAction string `json:"opponent_action"`
	ApproachOption string `json:"last_self_attack"`
	DefenseOption  string `json:"last_opponentf_attack"`
}

func APPROACH_CHECK(fb *FrameBuffer) (auth WriteAuth, diskframe DiskFrame) {
	// we want older frames first so we analyze any hitstun in context with an
	// approach
	frames := fb.GetXFramesDescending(30)

	// stores the frame index where we approached
	var approach_frame uint32 = 0
	var approach_index uint32 = 0
	//var defense_option PlayerAction = UNKNOWN_ANIMATION
	var approach_option PlayerAction = UNKNOWN_ANIMATION

	for i, f := range frames {
		if f.Self().Is(ATTACKING) && f.Opponent().Is(IN_IMMEDIATE_NEUTRAL) {
			approach_index = uint32(i)
			approach_frame = f.Number()
			approach_option = f.SelfAction()
			break
		}
	}

	if approach_frame == 0 || approach_index == 0 {
		auth = INDETERMINATE
		return
	}

	// Check if punished for approach. Only check frames newer than the one we
	// approached on
	for i := approach_index + 1; i < uint32(len(frames)); i++ {
		f := frames[i]

		// Successful approach
		if !f.Self().Is(HIT) && (f.Opponent().Is(DEAD) || f.Opponent().Is(HIT)) {
			diskframe = DiskFrame{
				BasicFrame: BasicFrame{
					FrameNumber: approach_frame,
					Data: ApproachData{
						StartingFrame:  approach_frame,
						SelfAction:     frames[approach_index].SelfAction().String(),
						OpponentAction: frames[approach_index].OpponentAction().String(),
						ApproachOption: approach_option.String(),
						DefenseOption:  frames[i].OpponentAction().String(),
					},

					WriteInvoker: "SUCCESSFUL_APPROACH",
				},
				DetailedFrame: nil,
			}
			return
		}

		// Failed approach
		if f.Self().Is(HIT) || f.Self().Is(DEAD) {
			auth = TRUE

			// failed approach. Write to disk
			diskframe = DiskFrame{
				BasicFrame: BasicFrame{
					FrameNumber: approach_frame,
					Data: ApproachData{
						StartingFrame:  approach_frame,
						SelfAction:     frames[i].SelfAction().String(),
						ApproachOption: approach_option.String(),
						OpponentAction: frames[i].OpponentAction().String(),
						DefenseOption:  frames[i].OpponentAction().String(),
						//DefenseOption:  fb.LastOpponentAttack(InRange(approach_index, i)).String(),
					},

					WriteInvoker: "FAILED_APPROACH",
				},
				DetailedFrame: nil,
			}
			return

		}
	}

	return
}

func (fb *FrameBuffer) LastSelfAttack(start int) PlayerAction {
	return UNKNOWN_ANIMATION
}

func (fb *FrameBuffer) LastAction(start, end int) []PlayerAction {
	attacks := []PlayerAction{
		UNKNOWN_ANIMATION, UNKNOWN_ANIMATION,
		UNKNOWN_ANIMATION, UNKNOWN_ANIMATION,
	}
	frames := fb.GetRange(start, end)

	for _, f := range frames {
		for i := 0; i < 4; i++ {
			action := f.Players[i].Action()
			if action.IsAttack() {
				attacks[i] = action
			}
		}
	}

	return attacks
}

// alias for GetAscendingRange
func (fb *FrameBuffer) InRange(start, end int) []Frame {
	return fb.GetAscendingRange(start, end)
}

func (fb *FrameBuffer) At(i int) []Frame {
	f, _ := fb.Get(i)
	return []Frame{f}
}
