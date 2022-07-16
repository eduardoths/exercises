package ex4_test

import (
	"testing"

	"github.com/eduardoths/exercises/go/pragmatic-programmer/ex4"
)

func TestPen_Down(t *testing.T) {
	type want struct {
		isDown bool
	}

	type testCase struct {
		pen  ex4.Pen
		want want
	}

	testCases := map[string]testCase{
		"it should turn down if the pen is not down": {
			pen:  ex4.Pen{IsDown: false},
			want: want{true},
		},
		"it should keep down if the pen is down": {
			pen:  ex4.Pen{IsDown: true},
			want: want{true},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			tc.pen.Down()
			if tc.want.isDown != tc.pen.IsDown {
				t.Errorf("wanted pen.IsDown to be %v but it was %v", tc.want.isDown, tc.pen.IsDown)
			}
		})
	}
}

func TestPen_Up(t *testing.T) {
	type want struct {
		isDown bool
	}

	type testCase struct {
		pen  ex4.Pen
		want want
	}

	testCases := map[string]testCase{
		"it should turn up if the pen is down": {
			pen:  ex4.Pen{IsDown: true},
			want: want{false},
		},
		"it should keep up if the pen is up": {
			pen:  ex4.Pen{IsDown: false},
			want: want{false},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			tc.pen.Up()
			if tc.want.isDown != tc.pen.IsDown {
				t.Errorf("wanted pen.IsDown to be %v but it was %v", tc.want.isDown, tc.pen.IsDown)
			}
		})
	}
}

func TestPen_Move(t *testing.T) {
	type want struct {
		position ex4.Position
	}

	type args struct {
		direction ex4.Direction
		distance  float64
	}

	type testCase struct {
		Pen  ex4.Pen
		args args
		want want
	}

	testCases := map[string]testCase{
		"it should move up by 1": {
			Pen:  ex4.Pen{Position: ex4.Position{0, 0}},
			args: args{direction: ex4.NORTH, distance: 1},
			want: want{position: ex4.Position{0, 1}},
		},

		"it should move right by 1": {
			Pen:  ex4.Pen{Position: ex4.Position{0, 0}},
			args: args{direction: ex4.EAST, distance: 1},
			want: want{position: ex4.Position{1, 0}},
		},
		"it should move down by 1": {
			Pen:  ex4.Pen{Position: ex4.Position{0, 0}},
			args: args{direction: ex4.SOUTH, distance: 1},
			want: want{position: ex4.Position{0, -1}},
		},
		"it should move left by 1": {
			Pen:  ex4.Pen{Position: ex4.Position{0, 0}},
			args: args{direction: ex4.WEST, distance: 1},
			want: want{position: ex4.Position{-1, 0}},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			tc.Pen.Move(tc.args.direction, tc.args.distance)
			if tc.want.position != tc.Pen.Position {
				t.Errorf("wanted position to be (%.2f, %.2f), got (%.2f, %.2f)",
					tc.want.position.X,
					tc.want.position.Y,
					tc.Pen.Position.X,
					tc.Pen.Position.Y,
				)
			}
		})
	}
}
