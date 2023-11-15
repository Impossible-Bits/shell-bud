package commands

import (
	"fmt"

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
					pinIcon = "\uf44d " // Nerd Fonts icon or emoji
				}
				fmt.Printf("%v %s | `%s`- %s \n", pinIcon, ansi.Color(c.Name, "blue+b"), ansi.Color(c.Command, "green+b"), c.Description)
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
				delete(commandList, name)
				commandList[name] = BashCommand{
					Name:           c.Name,
					Description:    c.Description,
					Command:        c.Command,
					Pinned:         true,
					Alias:          alias,
					DefaultMachine: defaultMachine,
				}
				setCommandsMap(commandList)
				fmt.Printf("Pinned command: %s with alias: %s on machine: %s\n", name, alias, defaultMachine)
			} else {
				fmt.Printf("Command not found: %s\n", name)
			}
		},
	}
)

func executeCommand(name string) {
	commands := viper.GetStringMap(COMMAND_KEY)
	if command, exists := commands[name]; exists {
		cmd := command.(map[string]interface{})
		defaultMachine := cmd["defaultMachine"].(string)

		if defaultMachine != "" {
			// Logic to execute the command on the default machine
			fmt.Printf("Executing '%s' on machine: %s\n", cmd["command"].(string), defaultMachine)
		} else {
			// Local execution logic
			fmt.Printf("Executing '%s' locally\n", cmd["command"].(string))
		}
	} else {
		fmt.Printf("Command not found: %s\n", name)
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
		if commandMap["alias"] == nil {
			commandMap["alias"] = ""
		}

		defaultMachine := ""
		if commandMap["defaultMachine"] == nil {
			commandMap["defaultMachine"] = ""
		}

		pinned := false
		if commandMap["pinned"] == nil {
			commandMap["pinned"] = false
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
	UserCommandCmd.AddCommand(addCommandCmd)
	UserCommandCmd.AddCommand(listCommandCmd)
	UserCommandCmd.AddCommand(removeCommandCmd)
}
