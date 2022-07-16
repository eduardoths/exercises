package ex4_test

import (
	"reflect"
	"testing"

	"github.com/eduardoths/exercises/go/pragmatic-programmer/ex4"
	"github.com/openlyinc/pointy"
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
					{ex4.WEST_COMMAND, pointy.Int(1)},
				},
				err: nil,
			},
		},
		"it should move to the west by 2": {
			args: args{s: "W 2"},
			want: want{
				commands: []ex4.TurtleCommand{
					{ex4.WEST_COMMAND, pointy.Int(2)},
				},
				err: nil,
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotCommands, gotErr := ex4.TurtleParse(tc.args.s)
			if !reflect.DeepEqual(tc.want.commands, gotCommands) {
				t.Errorf("wanted comand %v, got %v", tc.want.commands, gotCommands)
			}

			if tc.want.err != gotErr {
				t.Errorf("wanted err %s, got %s", tc.want.err, gotErr)
			}

		})
	}
}
