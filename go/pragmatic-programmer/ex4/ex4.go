package ex4

import (
	"strconv"
	"strings"
)

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
	Args    *float64
}

func TurtleParse(s string) ([]TurtleCommand, error) {
	commands := []TurtleCommand{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		turtleCommand, _ := turtleParseLine(line)
		if turtleCommand != nil {
			commands = append(commands, *turtleCommand)
		}
	}

	return commands, nil
}

func turtleParseLine(line string) (*TurtleCommand, error) {
	line = strings.TrimSpace(line)
	codeString := strings.TrimSpace(removeComments(line))
	if codeString == "" {
		return nil, nil
	}

	var command Command
	var args *float64

	commandChar := codeString[0]
	commandMap := map[byte]Command{
		'D': DOWN_COMMAND,
		'W': WEST_COMMAND,
		'U': UP_COMMAND,
		'E': EAST_COMMAND,
		'N': NORTH_COMMAND,
		'S': SOUTH_COMMAND,
	}
	command = commandMap[commandChar]
	argStr := strings.TrimSpace(codeString[1:])
	if argStr != "" {
		argInt, _ := strconv.ParseFloat(argStr, 64)
		args = &argInt
	}
	return &TurtleCommand{command, args}, nil
}

func removeComments(s string) string {
	commandAndComment := strings.Split(s, "#")
	return commandAndComment[0]
}
