package melee

/*
* TYPES SUMMARY: Different Frame types are used for different reasons.
*
* Frame:      Completely verbose. Holds all the data we collect each frame.
* BasicFrame: Only stores a PlayerPlayerAction and a message for it
*               e.g. "L_CANCEL_MISS"
* DiskFrame:  The only data type(other than Match) that is written to disk
 */
type Frame struct {
	Players PlayerContainer `json:"players"`
	empty   bool            `json:"empty"`
}

type BasicFrame struct {
	// no discrete use now for FrameNumber but may be useful later to
	// always store this.  Violates YAGNI for good reason
	FrameNumber    uint32 `json:"frame_number"`
	PlayerAction   string `json:"player_action"`
	OpponentAction string `json:"opponent_action"`
	Message        string `json:"message"`
}

// Data actually written to disk. For the most part this should be used
// similar to a C union
type DiskFrame struct {
	BasicFrame    BasicFrame `json:"basic_frame"`
	DetailedFrame Frame      `json:"detailed_frame"`
}

// Creates our Players slice only containing ports 1-4 (ignoring the global 0
// port and the psuedo 5-8 ports
func NewFrame(players PlayerContainer) Frame {
	f := Frame{
		players, false,
	}
	return f
}

func (f *Frame) Empty() bool {
	return f.empty
}

type OpponentIterator struct {
	current Port
	players PlayerContainer
}

func NewOpponentIterator(f Frame) *OpponentIterator {
	return &OpponentIterator{0, f.Players}
}

func (it *OpponentIterator) Next() bool {
	it.current++

	action, _ := it.players[it.current].GetPlayerAction()
	if it.current == FWriter.GetSelfPort() || action == UNKNOWN_ANIMATION {
		it.current++
	}

	if int(it.current) >= len(it.players) {
		return false
	}

	return true
}

func (it *OpponentIterator) Value() (Port, Player) {
	return it.current, it.players[it.current]
}

func (f Frame) SelfAction() PlayerAction {
	a, _ := f.Players[FWriter.GetSelfPort()].GetPlayerAction()
	return a
}

// TODO: report more than one? I don't care much about dubs...
func (f Frame) OpponentAction() PlayerAction {
	action := PlayerAction(UNKNOWN_ANIMATION)

	it := NewOpponentIterator(f)
	for it.Next() {
		_, player := it.Value()
		action, _ = player.GetPlayerAction()
		return action
	}

	return action
}

func NewBasicDiskFrame(f_num uint32, action PlayerAction, opponent PlayerAction, msg string) DiskFrame {
	return DiskFrame{
		BasicFrame: BasicFrame{
			FrameNumber:    f_num,
			PlayerAction:   ACTION_NAMES[action],
			OpponentAction: ACTION_NAMES[opponent],
			Message:        msg,
		},
	}
}
