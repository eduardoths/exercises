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

	command, _ := parseCommand(codeString)
	args, _ := parseArgument(codeString)

	return &TurtleCommand{command, args}, nil
}

func removeComments(s string) string {
	commandAndComment := strings.Split(s, "#")
	return commandAndComment[0]
}

func parseCommand(s string) (Command, error) {
	s = strings.TrimSpace(s)
	commandChar := s[0]
	commandMap := map[byte]Command{
		'D': DOWN_COMMAND,
		'W': WEST_COMMAND,
		'U': UP_COMMAND,
		'E': EAST_COMMAND,
		'N': NORTH_COMMAND,
		'S': SOUTH_COMMAND,
	}
	return commandMap[commandChar], nil
}

func parseArgument(s string) (*float64, error) {
	s = strings.TrimSpace(s)
	argStr := strings.TrimSpace(s[1:])
	if argStr != "" {
		argFloat, _ := strconv.ParseFloat(argStr, 64)
		return &argFloat, nil
	}

	return nil, nil
}
