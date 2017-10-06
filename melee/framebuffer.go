package melee

import (
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

type Frame struct {
	MenuState MenuState
	Stage     Stage
	Players   []Player
	empty     bool
}

func NewFrame(players PlayerContainer, m MenuState, s Stage) Frame {
	f := Frame{
		m, s, players[0 : PLAYER4+1], false,
	}
	return f
}

func (f *Frame) Empty() bool {
	return f.empty
}

const FRAME_BUFFER_SIZE = 30 // hold 0.5 seconds of frames in memory
type FrameBuffer struct {
	Frames *sll.List
}

func NewFrameBuffer() FrameBuffer {
	return FrameBuffer{sll.New()}
}

func (fb *FrameBuffer) Insert(f Frame) {
	fb.Frames.Prepend(f)

	if fb.Frames.Size() > FRAME_BUFFER_SIZE {
		fb.Frames.Remove(FRAME_BUFFER_SIZE)
	}
}

// second return parameter is true if index is withing bounds of the list and
// it is not empty
func (fb *FrameBuffer) Get(i int) (Frame, bool) {
	f, success := fb.Frames.Get(i)
	return f.(Frame), success
}

// this method should only be called AFTER the most recent frame is inserted
func (fb *FrameBuffer) GetPrevious() (Frame, bool) {
	return fb.Get(1)
}

func (fb *FrameBuffer) GetCurrent() (Frame, bool) {
	return fb.Get(0)
}

/* Since this can be a bit confusing to understand, here's an explanation:
* GetRange(6, 10) would return a slice of size 4 equivalent to the following:
*
* frames[0] = fb.Get(6)
* frames[1] = fb.Get(7)
* frames[2] = fb.Get(8)
* frames[3] = fb.Get(9)
*
* Meaning that the indexing of the returned slice is the same as the list(e.g.
* lower index = newer frame)
 */
func (fb *FrameBuffer) GetRange(start, end int) []Frame {
	frames := make([]Frame, end-start)

	for i := start; i < end; i++ {
		f, exists := fb.Frames.Get(i)
		if exists {
			frames[i-start] = f.(Frame)
		} else {
			frames[i-start] = Frame{empty: true}
		}
	}

	return frames
}
