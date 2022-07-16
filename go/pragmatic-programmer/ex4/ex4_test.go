package ex4_test

import (
	"errors"
	"testing"

	"github.com/eduardoths/exercises/go/pragmatic-programmer/ex4"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

func TestTurtleParse(t *testing.T) {
	type args struct {
		s string
	}

	type want struct {
		commands []ex4.TurtleCommand
		err      error
	}

	type testCase struct {
		args args
		want want
	}

	testCases := map[string]testCase{
		"it should return a command to set the pen down": {
			args: args{s: "D"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.DOWN_COMMAND, nil},
				},
				err: nil,
			},
		},
		"it should case insensitive": {
			args: args{s: "d"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.DOWN_COMMAND, nil},
				},
				err: nil,
			},
		},
		"it should return a command to set the pen up": {
			args: args{s: "U"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.UP_COMMAND, nil},
				},
				err: nil,
			},
		},
		"it should return multiple up and down commands": {
			args: args{s: "U\nD\nU\nD"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.UP_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
					{ex4.UP_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
				},
				err: nil,
			},
		},
		"it should ignore blank lines": {
			args: args{s: "U\nD\nU\nD\n\n"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.UP_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
					{ex4.UP_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
				},
				err: nil,
			},
		},
		"it should ignore comments": {
			args: args{s: "U\nD # comentario\nU\nD\n\n"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.UP_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
					{ex4.UP_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
				},
				err: nil,
			},
		},
		"it should ignore spaces": {
			args: args{s: " D\n\tD\n D\t\n U "},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.DOWN_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
					{ex4.DOWN_COMMAND, nil},
					{ex4.UP_COMMAND, nil},
				},
				err: nil,
			},
		},
		"it should move to the west by 1": {
			args: args{s: "W 1"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.WEST_COMMAND, pointy.Float64(1)},
				},
				err: nil,
			},
		},
		"it should move to the west by 2": {
			args: args{s: "W 2"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.WEST_COMMAND, pointy.Float64(2)},
				},
				err: nil,
			},
		},
		"it should move to the west by 1.3": {
			args: args{s: "W 1.3"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.WEST_COMMAND, pointy.Float64(1.3)},
				},
				err: nil,
			},
		},
		"it should move to west by 1.5 then by 0.3": {
			args: args{s: " W    1.5\t\nW -0.3"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.WEST_COMMAND, pointy.Float64(1.5)},
					{ex4.WEST_COMMAND, pointy.Float64(-0.3)},
				},
				err: nil,
			},
		},
		"it should move to east by 0.1": {
			args: args{s: "E 0.1"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.EAST_COMMAND, pointy.Float64(0.1)},
				},
				err: nil,
			},
		},
		"it should move to north by 10": {
			args: args{s: "N 10"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.NORTH_COMMAND, pointy.Float64(10)},
				},
				err: nil,
			},
		},
		"it should move to south by 5.3": {
			args: args{s: "S 5.3"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.SOUTH_COMMAND, pointy.Float64(5.3)},
				},
				err: nil,
			},
		},
		"it should return error if command is invalid": {
			args: args{s: "Test"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid command 'T'"),
			},
		},
		"it should return error if there are multiple commands and one is invalid": {
			args: args{s: "N\nT\nE\nW"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid command 'T'"),
			},
		},
		"it should return error if with invalid command": {
			args: args{s: "A"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid command 'A'"),
			},
		},
		"it should return error if argument is not a float64 number": {
			args: args{s: "N 1.64batata"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:could not parse arguments"),
			},
		},
		"it should return error if pen down has argument": {
			args: args{s: "D 1"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid arguments"),
			},
		},
		"it should return error if pen up has argument": {
			args: args{s: "U 0"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid arguments"),
			},
		},
		"it should return error if move north has no argument": {
			args: args{s: "N"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid arguments"),
			},
		},
		"it should return error if move south has no argument": {
			args: args{s: "S"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid arguments"),
			},
		},
		"it should return error if move east has no argument": {
			args: args{s: "E"},
			want: want{
				commands: nil,
				err:      errors.New("syntax_error:invalid arguments"),
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotCommands, gotErr := ex4.TurtleParse(tc.args.s)
			assert.Equal(t, tc.want.commands, gotCommands)
			assert.Equal(t, tc.want.err, gotErr)
		})
	}
}
