package melee

// check if both players have been in neutral for the last 15 frames. Arbitrary
// number
// probably wont use this now that I think about it... keeping it for now
func IN_NEUTRAL(fb *FrameBuffer) bool {
	// we want newer frames first because more recent info has precendence
	// since this gets ran every frame anyway.
	//frames := fb.GetXFramesAscending(30)

	//for _, f := range frames {
	//// If only 1 player is in neutral, it's most likely not considered an
	//// entirely neutral situtation. e.g the other player might be
	//// approaching. This will probably be tweaked later on.
	//if len(PortsInState(f, IN_IMMEDIATE_NEUTRAL)) <= 1 {
	//FWriter.RemoveFlag(NEUTRAL_FLAG)
	//}
	//}

	//FWriter.AddFlag(NEUTRAL_FLAG)
	return false
}

func FAILED_APPROACH(fb *FrameBuffer) (auth WriteAuth, data DiskFrame) {
	// we want older frames first so we analyze any hitstun in context with an
	// approach
	frames := fb.GetXFramesDescending(30)

	// stores the frame index where we approached
	approach_frame := -1

	for i, f := range frames {
		if SELF_IS(ATTACKING, f) && OPPONENT_IS(IN_IMMEDIATE_NEUTRAL, f) {
			approach_frame = i
			break
		}
	}

	if approach_frame == -1 {
		auth = INDETERMINATE
		return
	}

	// Check if punished for approach. Only check frames newer than the one we
	// approached on
	for i := approach_frame + 1; i < len(frames); i++ {
		f := frames[i]

		if SELF_IS(HIT, f) || SELF_IS(DEAD, f) {
			auth = TRUE

			// failed approach. Write to disk
			data = NewBasicDiskFrame(
				//frames[i].Players[0].FrameNumber,
				0,
				frames[i].SelfAction(),
				frames[i].OpponentAction(),
				"FAILED_APPROACH")
			return
		}
	}

	return
}

func SUCCESSFUL_APPROACH(fb *FrameBuffer) (auth WriteAuth, data DiskFrame) {
	// we want older frames first so we analyze any hitstun in context with an
	// approach
	frames := fb.GetXFramesDescending(30)

	// stores the frame index where we approached
	approach_frame := -1

	for i, f := range frames {
		if SELF_IS(ATTACKING, f) && OPPONENT_IS(IN_IMMEDIATE_NEUTRAL, f) {
			approach_frame = i
			break
		}
	}

	if approach_frame == -1 {
		auth = INDETERMINATE
		return
	}

	// Check if punished for approach. Only check frames newer than the one we
	// approached on
	for i := approach_frame + 1; i < len(frames); i++ {
		f := frames[i]

		if !SELF_IS(HIT, f) && (OPPONENT_IS(DEAD, f) || OPPONENT_IS(HIT, f)) {
			auth = TRUE

			// successful approach. Write to disk
			data = NewBasicDiskFrame(
				//frames[i].FrameNumber,
				0,
				frames[i].SelfAction(),
				frames[i].OpponentAction(),
				"SUCCESSFUL_APPROACH")
			return
		}
	}

	return
}

func LAST_ATTACK(fb *FrameBuffer) []PlayerAction {
	attacks := []PlayerAction{
		UNKNOWN_ANIMATION, UNKNOWN_ANIMATION,
		UNKNOWN_ANIMATION, UNKNOWN_ANIMATION,
	}
	frames := fb.GetXFramesAscending(60)

	for _, f := range frames {
		if f.SelfAction().IsAttack() {
			attacks[FWriter.GetSelfPort()] = f.SelfAction()
		}

		it := NewOpponentIterator(f)
		for it.Next() {
			port, player := it.Value()
			attacks[port], _ = player.GetPlayerAction()
		}
	}

	return attacks
}
