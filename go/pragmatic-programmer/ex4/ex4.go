package ex4

import (
	"errors"
	"fmt"
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
		turtleCommand, err := turtleParseLine(line)
		if err != nil {
			return nil, err
		}

		if turtleCommand != nil {
			commands = append(commands, *turtleCommand)
		}

	}

	for _, turtleCommand := range commands {
		if turtleCommand.Command == DOWN_COMMAND || turtleCommand.Command == UP_COMMAND {
			if turtleCommand.Args != nil {
				return nil, errors.New("syntax_error:invalid arguments")
			}
		}
		if turtleCommand.Command == NORTH_COMMAND || turtleCommand.Command == SOUTH_COMMAND {
			if turtleCommand.Args == nil {
				return nil, errors.New("syntax_error:invalid arguments")
			}
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

	command, err := parseCommand(codeString)
	if err != nil {
		return nil, err
	}

	args, err := parseArgument(codeString)
	if err != nil {
		return nil, err
	}

	return &TurtleCommand{command, args}, nil
}

func removeComments(s string) string {
	commandAndComment := strings.Split(s, "#")
	return commandAndComment[0]
}

func parseCommand(s string) (Command, error) {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	commandChar := s[0]
	commandMap := map[byte]Command{
		'D': DOWN_COMMAND,
		'W': WEST_COMMAND,
		'U': UP_COMMAND,
		'E': EAST_COMMAND,
		'N': NORTH_COMMAND,
		'S': SOUTH_COMMAND,
	}
	command, ok := commandMap[commandChar]
	if !ok {
		return command, fmt.Errorf("syntax_error:invalid command '%s'", string(commandChar))
	}

	return commandMap[commandChar], nil
}

func parseArgument(s string) (*float64, error) {
	s = strings.TrimSpace(s)
	argStr := strings.TrimSpace(s[1:])
	if argStr != "" {
		argFloat, err := strconv.ParseFloat(argStr, 64)
		if err != nil {
			return nil, errors.New("syntax_error:could not parse arguments")
		}
		return &argFloat, nil
	}

	return nil, nil
}
