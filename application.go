package cli

import (
	"fmt"
	"os"
	"strings"
)

type Application struct {
	Commands       map[string]Command
	DefaultCommand Command
}

func NewApplication() *Application {
	a := &Application{}
	a.Commands = make(map[string]Command)
	return a
}

func (a *Application) Register(alias string, c Command) {
	a.Commands[alias] = c
}

func (a *Application) RegisterDefault(c Command) {
	a.DefaultCommand = c
}

func (a *Application) Handle(cli *CLI) {
	command := cli.GetCommand()
	if command == "" {
		a.DefaultCommand.Invoke(cli, a)
		return
	}
	if a.CanHandle(command) {
		c := a.Commands[command]
		if c != nil {
			c.Invoke(cli, a)
		}
	} else {
		fmt.Printf("Cannot handle %v\n", command)
		os.Exit(1)
	}
}

func (a *Application) CanHandle(command string) bool {
	return a.Commands[command] != nil
}

type Command interface {
	Invoke(cli *CLI, a *Application)
	Help() string
	Description() string
}

type CommandUsage struct {
}

func (c CommandUsage) Invoke(cli *CLI, a *Application) {
	for key, cmd := range a.Commands {
		fmt.Println(key + " - " + cmd.Help())
	}
}
func (c CommandUsage) Help() string {
	return "Used to describe the workings of the application."
}
func (c CommandUsage) Description() string {
	return "Displays usage of all commands."
}

type CommandHelp struct {
}

func (c CommandHelp) Invoke(cli *CLI, a *Application) {
	command := cli.GetCommand()
	splits := strings.Split(command, ",")
	cmdname := splits[1]
	cmd := a.Commands[cmdname]
	fmt.Println(cmd.Help())
}
func (c CommandHelp) Help() string {
	return "Used to how the help of other commands."
}
func (c CommandHelp) Description() string {
	return "Displays help for a specific command."
}
