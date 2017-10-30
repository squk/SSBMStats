package melee

import arrayList "github.com/emirpasic/gods/lists/arraylist"

// hold 1 second of frames in memory. Arbitraty ATM.
const FRAME_BUFFER_SIZE = 240

type FrameBuffer struct {
	Frames *arrayList.List
}

func NewFrameBuffer() FrameBuffer {
	return FrameBuffer{arrayList.New()}
}

func (fb *FrameBuffer) Insert(f Frame) {
	fb.Frames.Insert(0, f)
	//fb.Frames.Add(f)

	if fb.Frames.Size() >= FRAME_BUFFER_SIZE {
		//fb.Frames.Remove(0)
		fb.Frames.Remove(FRAME_BUFFER_SIZE - 1)
	}
}

// second return parameter is true if index is withing bounds of the list and
// it is not empty
func (fb *FrameBuffer) Get(i int) (Frame, bool) {
	v, exists := fb.Frames.Get(i)

	if exists {
		f, ok := v.(Frame)
		return f, ok
	} else {
		return EmptyFrame(), false
	}
}

// this method should only be called AFTER the most recent frame is inserted
func (fb *FrameBuffer) GetPrevious() (Frame, bool) {
	return fb.Get(1)
}

func (fb *FrameBuffer) GetCurrent() (Frame, bool) {
	return fb.Get(0)
}

func (fb *FrameBuffer) GetRange(start, end int) []Frame {
	return fb.GetAscendingRange(start, end)
}

// NEWER FRAMES PRECEDE OLDER ONES. Inclusive.
func (fb *FrameBuffer) GetAscendingRange(start, end int) []Frame {
	frames := make([]Frame, end-start)

	for i := start; i < end; i++ {
		f, exists := fb.Frames.Get(i)
		if exists {
			frames[i] = f.(Frame)
		} else {
			frames[i] = Frame{empty: true}
		}
	}

	return frames
}

// OLDER FRAMES PRECEDE NEWER ONES. Inclusive
func (fb *FrameBuffer) GetDescendingRange(start, end int) []Frame {
	frames := make([]Frame, end-start)
	//var frames []Frame

	for i := end - 1; i >= start; i-- {
		f, exists := fb.Frames.Get(i)
		if exists {
			frames[i] = f.(Frame)
			//frames = append(frames, f.(Frame))
		} else {
			frames[i] = Frame{empty: true}
			//frames = append(frames, Frame{empty: true})
		}
	}

	return frames
}

// NEWER FRAMES PRECEDE OLDER ONES.
func (fb *FrameBuffer) GetXFramesAscending(num int) []Frame {
	return fb.GetAscendingRange(0, num)
}

// OLDER FRAMES PRECEDE NEWER ONES.
func (fb *FrameBuffer) GetXFramesDescending(num int) []Frame {
	return fb.GetDescendingRange(0, num)
}
