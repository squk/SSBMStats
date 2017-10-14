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

func FAILED_APPROACH(fb *FrameBuffer) (auth WriteAuth, data DiskFrame, descriptor FrameDescriptor) {
	// we want older frames first so we analyze any hitstun in context with an
	// approach
	frames := fb.GetXFramesDescending(30)
	//f0, _ := fb.Get(0)
	//a0 := ACTION_NAMES[f0.SelfAction()]
	//f1, _ := fb.Get(30)
	//a1 := ACTION_NAMES[f1.SelfAction()]
	//a0 := ACTION_NAMES[frames[0].SelfAction()]
	//a1 := ACTION_NAMES[frames[30].SelfAction()]

	//log.Println(a0, a1)

	// stores the frame index where we approached
	approach_frame := -1

	for i, f := range frames {
		if f.SelfIs(ATTACKING) && f.OpponentIs(IN_IMMEDIATE_NEUTRAL) {
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

		if f.SelfIs(HIT) || f.SelfIs(DEAD) {
			auth = TRUE

			frame_number, _ := f.Players[0].GetUint(FRAME)

			// failed approach. Write to disk
			data = NewBasicDiskFrame(
				frames[i].SelfAction(),
				frames[i].SelfAction(), // TODO: self last attack
				frames[i].OpponentAction(),
				frames[i].OpponentAction(), // TODO: opponent last attack
			)

			descriptor = FrameDescriptor{
				FrameNumber:  frame_number,
				WriteInvoker: "FAILED_APPROACH",
			}
			return
		}
	}

	return
}

func SUCCESSFUL_APPROACH(fb *FrameBuffer) (auth WriteAuth, data DiskFrame, descriptor FrameDescriptor) {
	// we want older frames first so we analyze any hitstun in context with an
	// approach
	frames := fb.GetXFramesDescending(30)

	// stores the frame index where we approached
	approach_frame := -1
	//a0 := ACTION_NAMES[frames[0].SelfAction()]
	//a1 := ACTION_NAMES[frames[30].SelfAction()]

	//log.Println(a0, a1)

	for i, f := range frames {
		if f.SelfIs(ATTACKING) {
		}

		if f.SelfIs(ATTACKING) && f.OpponentIs(IN_IMMEDIATE_NEUTRAL) {
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

		if !f.SelfIs(HIT) && (f.OpponentIs(DEAD) || f.OpponentIs(HIT)) {
			auth = TRUE

			frame_number, _ := f.Players[0].GetUint(FRAME)

			// successful approach. Write to disk
			data = NewBasicDiskFrame(
				frames[i].SelfAction(),
				frames[i].SelfAction(), // TODO: self last attack
				frames[i].OpponentAction(),
				frames[i].OpponentAction(), // TODO: opponent last attack
			)

			descriptor = FrameDescriptor{
				FrameNumber:  frame_number,
				WriteInvoker: "SUCCESSFUL_APPROACH",
			}

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
