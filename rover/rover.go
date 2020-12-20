package rover

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
	//do nothing now

	return r
}

func (r *Rover) Move(step string) *Rover {

	//do nothing now
	return r
}
