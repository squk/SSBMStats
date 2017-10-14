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
	SelfAction         string `json:"self_action"`
	SelfLastAttack     string `json:"self_last_attack"`
	OpponentAction     string `json:"opponent_action"`
	OpponentLastAttack string `json:"opponent_last_attack"`
}

type FrameDescriptor struct {
	FrameNumber  uint32
	WriteInvoker string
}

// Data actually written to disk. For the most part this should be used
// similar to a C union
type DiskFrame struct {
	BasicFrame    BasicFrame `json:"basic_frame"`
	DetailedFrame *Frame     `json:"detailed_frame"`
}

func NewFrame(players PlayerContainer) Frame {
	return Frame{players.DeepCopy(), false}
}

func EmptyFrame() Frame {
	return Frame{empty: true}
}

func (f Frame) Empty() bool {
	return f.empty
}

func (f Frame) DeepCopy() Frame {
	return Frame{f.Players.DeepCopy(), f.Empty()}
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

func (f *Frame) SelfAction() PlayerAction {
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

func NewBasicDiskFrame(action PlayerAction, self_attack PlayerAction, opponent PlayerAction, opponent_attack PlayerAction) DiskFrame {
	return DiskFrame{
		BasicFrame: BasicFrame{
			SelfAction:         ACTION_NAMES[action],
			SelfLastAttack:     ACTION_NAMES[self_attack],
			OpponentAction:     ACTION_NAMES[opponent],
			OpponentLastAttack: ACTION_NAMES[opponent_attack],
		},
		DetailedFrame: nil,
	}
}
