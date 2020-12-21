package rover

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var directions = Directions{North{}, East{}, West{}, South{}}
func TestRover_Init(t *testing.T) {
	// add location
	p := position{0,0,NORTH,directions}
	rover := Rover{}
	rover.Init(0,0,NORTH)
	cp := rover.CurrentPosition()
	assert.Equal(t, p.X, cp.X)
	assert.Equal(t, p.Y, cp.Y)
	assert.Equal(t, p.Direction, cp.Direction)
}


func TestRover_Move(t *testing.T) {
	rover := Rover{}

	//test cases
	testcases := []struct {
		testcasename string
		initposition position
		command string
		expectedposition position
	}{
		{
			"test rotating only",
			position{0,0,NORTH,directions},
			"RRRR",
			position{X: 0,Y: 0,Direction: NORTH},
		},
		{
			"test Moving",
			position{0,0,NORTH,directions},
			"FFRFLB",
			position{X: 1,Y: 1,Direction: NORTH},
		},
		{
			"test Moving from non-origin",
			position{2,1,EAST,directions},
			"BRBLBR",
			position{X: 0,Y: 2,Direction: SOUTH},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcasename, func(t *testing.T) {
			rover.Init(tc.initposition.X,tc.initposition.Y,tc.initposition.Direction)
			rover.Command(tc.command)
			cp := rover.CurrentPosition()

			assert.Equalf(t,cp.X,tc.expectedposition.X,"expected x to be = %d, given %d", tc.expectedposition.X, cp.X)
			assert.Equalf(t,cp.Y,tc.expectedposition.Y,"expected x to be = %d, given %d", tc.expectedposition.Y, cp.Y)
			assert.Equalf(t,cp.Direction,tc.expectedposition.Direction,"expected Direction to be = %s, given %s", tc.expectedposition.Direction, cp.Direction)
		})
	}
}




