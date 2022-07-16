package ex4

import "strings"

type Command string

const (
	DOWN_COMMAND  Command = "DOWN"
	UP_COMMAND    Command = "UP"
	NORTH_COMMAND Command = "NORTH"
	SOUTH_COMMAND Command = "SOUTH"
	EAST_COMMAND  Command = "EAST"
	WEST_COMMAND  Command = "WEST"
)

type TurtleCommand struct {
	Command Command
	Args    *int
}

func TurtleParse(s string) ([]TurtleCommand, error) {
	commands := []TurtleCommand{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		commandAndComment := strings.Split(line, "#")
		line = strings.TrimSpace(commandAndComment[0])
		if line == "" {
			continue
		}

		var command Command
		var args *int

		if line == "D" {
			command = DOWN_COMMAND
		} else if line[0] == 'W' {
			one := 1
			args = &one
			command = WEST_COMMAND
		} else {
			command = UP_COMMAND
		}
		commands = append(commands, TurtleCommand{command, args})
	}

	return commands, nil
}
