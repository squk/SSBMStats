package melee

var L_CANCEL_OPPORTUNE bool // prevents multiple frames being marked as a successful L-Cancel

// returns a DiskFrame and whether or not that DiskFrame should be logged
func L_CANCEL_PASS(f Frame) (auth WriteAuth, data DiskFrame) {
	auth = INDETERMINATE

	if !f.Empty() {
		action := f.Players[GameState.SelfPort].Action()
		//log.Println(ACTION_NAMES[action])

		if !L_CANCEL_OPPORTUNE {
			char, _ := f.Players[GameState.SelfPort].GetCharacter()
			animation_speed, _ := f.Players[GameState.SelfPort].GetFloat(SPEED_ANIMATION)
			if action.CanLCancel() {
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
					auth = TRUE
					msg = "L_CANCEL_PASS"
				} else {
					auth = TRUE
					msg = "L_CANCEL_MISS"
				}

				data = DiskFrame{
					BasicFrame: BasicFrame{
						FrameNumber:  f.Number(),
						Data:         struct{ Action PlayerAction }{f.SelfAction()},
						WriteInvoker: msg,
					},
					DetailedFrame: nil,
				}

			}
		} else {
			if !action.CanLCancel() {
				L_CANCEL_OPPORTUNE = false
			}
		}
	}

	return auth, data
}
