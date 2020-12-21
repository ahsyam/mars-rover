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
	if step == FORWARD {
		switch p.Direction {
			case NORTH:
				p.Y++
			case EAST:
				p.X++
			case SOUTH:
				p.Y--
			case WEST:
				p.X--
		}
	}
	if step == RIGHT {
		switch p.Direction {
			case NORTH:
				p.Direction = EAST
			case EAST:
				p.Direction = SOUTH
			case SOUTH:
				p.Direction = WEST
			case WEST:
				p.Direction = NORTH
		}
	}
	if step == LEFT {
		switch p.Direction {
		case NORTH:
			p.Direction = WEST
		case EAST:
			p.Direction = NORTH
		case SOUTH:
			p.Direction = EAST
		case WEST:
			p.Direction = SOUTH
		}
	}
	if step == BACKWARD {
		switch p.Direction {
		case NORTH:
			p.Y--
		case EAST:
			p.X--
		case SOUTH:
			p.Y++
		case WEST:
			p.X++
		}
	}

	return p
}