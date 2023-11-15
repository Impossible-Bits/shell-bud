package commands

var COMMAND_KEY = "commands"

type BashCommand struct {
	Name           string
	Description    string
	Command        string
	Pinned         bool
	Alias          string
	DefaultMachine string
}
