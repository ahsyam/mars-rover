package rover

const NORTH = "N"
const EAST = "E"
const WEST = "W"
const SOUTH = "S"

type position struct {
	X, Y int
	Direction string
	Directions Directions
}

func (p *position) MovePosition(step string) *position {
	//check the movement step and change the X,Y and direction based on it
	var direction direction

	switch p.Direction {
	case NORTH:
		direction = North{}
	case EAST:
		direction = East{}
	case WEST:
		direction = West{}
	case SOUTH:
		direction = South{}
	}

	switch step {
	case FORWARD:
		return direction.forward(p)
	case BACKWARD:
		return direction.backward(p)
	case RIGHT:
		return direction.right(p)
	case LEFT:
		return direction.left(p)
	}
	return p
}