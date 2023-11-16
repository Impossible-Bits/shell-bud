package commands

var COMMAND_KEY = "commands"
var MACROS_KEY = "macros"

type BashCommand struct {
	Name           string
	Description    string
	Command        string
	Pinned         bool
	Alias          string
	DefaultMachine string
}

type Macro struct {
	Name        string
	Description string
	Commands    []string
}
