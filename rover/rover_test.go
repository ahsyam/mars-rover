package rover

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRover_Init(t *testing.T) {
	// add location
	p := position{0,0,NORTH}
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
			position{0,0,NORTH},
			"RRRR",
			position{X: 0,Y: 0,Direction: NORTH},
		},
		{
			"test Moving",
			position{0,0,NORTH},
			"FFRFLB",
			position{X: 1,Y: 1,Direction: NORTH},
		},
		{
			"test Moving from non-origin",
			position{2,1,EAST},
			"BRBLBR",
			position{X: 0,Y: 2,Direction: SOUTH},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcasename, func(t *testing.T) {
			rover.Init(tc.initposition.X,tc.initposition.Y,tc.initposition.Direction)
			rover.Command(tc.command)
			cp := rover.CurrentPosition()
			if cp.X != tc.expectedposition.X {
				t.Errorf("expected x to be = %d, given %d", tc.expectedposition.X, cp.X)
			}
			if cp.Y != tc.expectedposition.Y {
				t.Errorf("expected Y to be = %d, given %d", tc.expectedposition.Y, cp.Y)
			}
			if cp.Direction != tc.expectedposition.Direction {
				t.Errorf("expected Direction to be = %s, given %s", tc.expectedposition.Direction, cp.Direction)
			}
		})
	}
}




