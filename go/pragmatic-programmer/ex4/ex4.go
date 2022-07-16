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
	Args    *[]interface{}
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

		if line == "D" {
			commands = append(commands, TurtleCommand{DOWN_COMMAND, nil})
		} else {
			commands = append(commands, TurtleCommand{UP_COMMAND, nil})
		}
	}

	return commands, nil
}
