package commands

import (
	"fmt"
	"os/exec"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var UserCommandCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Manage commands",
	Long:  `Add, list, and remove commands.`,
}

var (
	addCommandCmd = &cobra.Command{
		Use:   "add [name] [command] [description] [alias]",
		Short: "Add a new command",
		Args:  cobra.RangeArgs(2, 4),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			command := args[1]

			description := ""
			if len(args) > 2 {
				description = args[2]
			}

			alias := ""
			if len(args) > 3 {
				alias = args[3]
			}

			commandList := getCommandsMap()
			for key := range commandList {
				if key == name {
					fmt.Printf("Command already exists: %s\n", name)
					return
				}
			}

			commandList[name] = BashCommand{
				Name:           name,
				Description:    description,
				Command:        command,
				Pinned:         false,
				Alias:          alias,
				DefaultMachine: "",
			}

			setCommandsMap(commandList)
			fmt.Printf("Added command: %s: %s, %s\n", name, description, command)
		},
	}

	listCommandCmd = &cobra.Command{
		Use:   "list",
		Short: "List all machines",
		Run: func(cmd *cobra.Command, args []string) {
			commandList := getCommandsMap()
			for _, c := range commandList {
				pinIcon := ""
				if c.Pinned {
					pinIcon = "\ueba0"
				}
				fmt.Printf("%s %s | `%s`- %s \n", pinIcon, ansi.Color(c.Name, "blue+b"), ansi.Color(c.Command, "green+b"), c.Description)
			}
		},
	}

	removeCommandCmd = &cobra.Command{
		Use:   "remove [name]",
		Short: "Remove a machine",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			machineList := getCommandsMap()
			if _, exists := machineList[name]; exists {
				delete(machineList, name)
				setCommandsMap(machineList)
				fmt.Printf("Removed machine: %s\n", name)
			} else {
				fmt.Printf("BashCommand not found: %s\n", name)
			}
		},
	}

	pinCommandCmd = &cobra.Command{
		Use:   "pin [name] [defaultMachineName] [alias]",
		Short: "Pin a command and optionally set an alias",
		Args:  cobra.RangeArgs(0, 3),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			var alias, defaultMachine string

			defaultMachine = ""
			if len(args) > 1 {
				defaultMachine = args[1]
			}

			alias = ""
			if len(args) > 2 {
				alias = args[2]
			}
			commandList := getCommandsMap()
			if c, exists := commandList[name]; exists {
				nc := BashCommand{
					Name:           c.Name,
					Command:        c.Command,
					Description:    c.Description,
					Pinned:         true,
					Alias:          alias,
					DefaultMachine: defaultMachine,
				}
				delete(commandList, name)
				commandList[name] = nc
				setCommandsMap(commandList)
				fmt.Printf("Pinned command: %s with alias: %s on machine: %s\n", name, alias, defaultMachine)
			} else {
				fmt.Printf("Command not found: %s\n", name)
			}
		},
	}

	unpinCommandCmd = &cobra.Command{
		Use:   "unpin [name]",
		Short: "Unpin a command and optionally set an alias",
		Args:  cobra.RangeArgs(0, 3),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			var alias, defaultMachine string

			commandList := getCommandsMap()
			if c, exists := commandList[name]; exists {
				nc := BashCommand{
					Name:           c.Name,
					Command:        c.Command,
					Description:    c.Description,
					Pinned:         false,
					Alias:          c.Alias,
					DefaultMachine: c.DefaultMachine,
				}
				delete(commandList, name)
				commandList[name] = nc
				setCommandsMap(commandList)
				fmt.Printf("Unpinned command: %s with alias: %s on machine: %s\n", name, alias, defaultMachine)
			} else {
				fmt.Printf("Command not found: %s\n", name)
			}
		},
	}
)

var addMacroCmd = &cobra.Command{
	Use:   "am [macro name] [command names...]",
	Short: "Add a new macro",
	Long:  `Create a new macro with a sequence of commands.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		commandNames := args[1:]

		macros := getMacrosMap()
		if _, exists := macros[name]; exists {
			fmt.Printf("Macro already exists: %s\n", name)
			return
		}

		commands := getCommandsMap()
		for _, name := range commandNames {
			if _, exists := commands[name]; !exists {
				fmt.Printf("Command not found: %s make sure to have it in commands first\n", name)
				return
			}
		}

		newMacro := Macro{
			Name:        name,
			Description: "Custom macro",
			Commands:    commandNames,
		}

		macros[name] = newMacro
		setMacroMap(macros)
		fmt.Printf("Macro '%s' added with commands: %v\n", name, commandNames)
	},
}

var listMacroCmd = &cobra.Command{
	Use:   "lm",
	Short: "List a macro list",
	Long:  `Execute all commands associated with a given macro.`,
	Run: func(cmd *cobra.Command, args []string) {
		macros := getMacrosMap()
		for _, macro := range macros {
			fmt.Printf("Macro: %s\n", macro.Name)
			fmt.Printf("Description: %s\n", macro.Description)
			fmt.Printf("Commands: %v\n\n", macro.Commands)
		}
	},
}

var executeMacroCmd = &cobra.Command{
	Use:   "em [macro name]",
	Short: "Execute a macro",
	Long:  `Execute all commands associated with a given macro.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		macroName := args[0]
		executeMacro(macroName)
	},
}

func executeMacro(macroName string) {
	macros := getMacrosMap()
	if macro, exists := macros[macroName]; exists {
		for _, cmdName := range macro.Commands {
			if command, exists := getCommandsMap()[cmdName]; exists {
				fmt.Printf("Executing command '%s': %s\n", cmdName, command.Command)
				executeCommand(command.Command)
			}
		}
	} else {
		fmt.Println("Macro not found:", macroName)
	}
}

func executeCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func getMacrosMap() map[string]Macro {
	macros := make(map[string]Macro)
	rawMacros := viper.GetStringMap(MACROS_KEY)

	for key, raw := range rawMacros {
		macroMap := raw.(map[string]interface{})
		commands := []string{}
		for _, cmd := range macroMap["commands"].([]interface{}) {
			commands = append(commands, cmd.(string))
		}
		macros[key] = Macro{
			Name:        macroMap["name"].(string),
			Description: macroMap["description"].(string),
			Commands:    commands,
		}
	}

	return macros
}

func setMacroMap(macros map[string]Macro) {
	viper.Set(MACROS_KEY, macros)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error writing macros to config: %s\n", err)
	}
}

func getCommandsMap() map[string]BashCommand {
	commandList := make(map[string]BashCommand)
	for key, raw := range viper.GetStringMap(COMMAND_KEY) {
		commandMap := raw.(map[string]interface{})

		description := ""
		if commandMap["description"] != nil {
			description = commandMap["description"].(string)
		}

		alias := ""
		if commandMap["alias"] != nil {
			alias = commandMap["alias"].(string)
		}

		defaultMachine := ""
		if commandMap["defaultMachine"] != nil {
			defaultMachine = commandMap["defaultMachine"].(string)
		}

		pinned := false
		if commandMap["pinned"] != nil {
			pinned = commandMap["pinned"].(bool)
		}

		command := BashCommand{
			Name:           commandMap["name"].(string),
			Command:        commandMap["command"].(string),
			Description:    description,
			Pinned:         pinned,
			Alias:          alias,
			DefaultMachine: defaultMachine,
		}
		commandList[key] = command
	}
	return commandList
}

func setCommandsMap(commands map[string]BashCommand) {
	viper.Set(COMMAND_KEY, commands)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error writing commands to config: %s\n", err)
	}
}

func init() {
	UserCommandCmd.AddCommand(pinCommandCmd)
	UserCommandCmd.AddCommand(unpinCommandCmd)
	UserCommandCmd.AddCommand(addCommandCmd)
	UserCommandCmd.AddCommand(listCommandCmd)
	UserCommandCmd.AddCommand(removeCommandCmd)
	UserCommandCmd.AddCommand(addMacroCmd)
	UserCommandCmd.AddCommand(executeMacroCmd)
	UserCommandCmd.AddCommand(listMacroCmd)
}
