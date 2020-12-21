package rover

import "strings"

//directions
const FORWARD = "F"
const BACKWARD = "B"
const LEFT = "L"
const RIGHT = "R"

type Rover struct {
	currentposition position
}

func (r *Rover) Init(x, y int, direction string) *Rover {
	r.currentposition = position{x, y, direction}
	return r
}

func (r *Rover) CurrentPosition() position {
	return r.currentposition
}

func (r *Rover) Command(command string) *Rover {
	commandSlice := strings.Split(command, "")
	for _, step := range commandSlice {
		r.Move(step)
	}
	return r
}

func (r *Rover) Move(step string) *Rover {
	cp := r.CurrentPosition()
	cp.MovePosition(step)
	r.currentposition = cp
	return r
}
