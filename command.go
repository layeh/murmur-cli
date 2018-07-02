package main

import (
	"errors"
	"flag"
)

type Command struct {
	Commands map[string]*Command
	Run      func(Args)
}

func NewCommand() *Command {
	return &Command{
		Commands: make(map[string]*Command),
	}
}

func (c *Command) Add(name string, run ...func(arguments Args)) *Command {
	cmd := NewCommand()
	c.Commands[name] = cmd
	if len(run) >= 1 {
		cmd.Run = run[0]
	}
	return cmd
}

func (c *Command) Do() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	args := flag.Args()
	current := c
	for i, arg := range flag.Args() {
		if sub := current.Commands[arg]; sub != nil {
			current = sub
		} else {
			if current.Run == nil {
				return errors.New("unknown command")
			}
			current.Run(args[i:])
			return nil
		}
	}
	if current.Run == nil {
		return errors.New("unknown command")
	}
	current.Run(nil)
	return nil
}
