package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectionsMovements(t *testing.T) {

	//test cases
	testcases := []struct {
		testcasename string
		direction direction
		initdirection string
		step      string
		expectedX int
		expectedY int
		expectedDirection string
	}{
		{
			"test North move forward",
			North{},
			NORTH,
			FORWARD,
			0,
			1,
			NORTH,
		},
		{
			"test North move backward",
			North{},
			NORTH,
			BACKWARD,
			0,
			-1,
			NORTH,
		},
		{
			"test North turn left",
			North{},
			NORTH,
			LEFT,
			0,
			0,
			WEST,
		},
		{
			"test North turn right",
			North{},
			NORTH,
			RIGHT,
			0,
			0,
			EAST,
		},
		{
			"test East move forward",
			East{},
			EAST,
			FORWARD,
			1,
			0,
			EAST,
		},
		{
			"test East move backward",
			East{},
			EAST,
			BACKWARD,
			-1,
			0,
			EAST,
		},
		{
			"test East turn left",
			East{},
			EAST,
			LEFT,
			0,
			0,
			NORTH,
		},
		{
			"test East turn right",
			East{},
			EAST,
			RIGHT,
			0,
			0,
			SOUTH,
		},
		{
			"test West  move forward",
			West{},
			WEST,
			FORWARD,
			-1,
			0,
			WEST,
		},
		{
			"test West move backward",
			West{},
			WEST,
			BACKWARD,
			1,
			0,
			WEST,
		},
		{
			"test West turn left",
			West{},
			WEST,
			LEFT,
			0,
			0,
			SOUTH,
		},
		{
			"test West turn right",
			West{},
			WEST,
			RIGHT,
			0,
			0,
			NORTH,
		},
		{
			"test South move forward",
			South{},
			SOUTH,
			FORWARD,
			0,
			-1,
			SOUTH,
		},
		{
			"test South move backward",
			South{},
			SOUTH,
			BACKWARD,
			0,
			1,
			SOUTH,
		},
		{
			"test South turn left",
			South{},
			SOUTH,
			LEFT,
			0,
			0,
			EAST,
		},
		{
			"test South turn right",
			South{},
			SOUTH,
			RIGHT,
			0,
			0,
			WEST,
		},
	}

	directions := Directions{North{}, East{}, West{}, South{}}
	for _, tc := range testcases {
		t.Run(tc.testcasename, func(t *testing.T) {
			//init position
			position := position{0,0,tc.initdirection, directions}

			switch tc.step {
				case FORWARD:
					tc.direction.forward(&position)
				case BACKWARD:
					tc.direction.backward(&position)
				case LEFT:
					tc.direction.left(&position)
				case RIGHT:
					tc.direction.right(&position)
			}

			assert.Equal(t, tc.expectedX, position.X)
			assert.Equal(t, tc.expectedY, position.Y)
			assert.Equal(t, tc.expectedDirection, position.Direction)
		})
	}
}