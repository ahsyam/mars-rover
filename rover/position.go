package rover

const NORTH = "N"
const EAST = "E"
const WEST = "W"
const SOUTH = "S"

type position struct {
	X, Y int
	Direction string
}


func (p *position) MovePosition(step string) *position {
	//check the movement step and change the X,Y and direction based on it

	//Do nothing now
	return p
}