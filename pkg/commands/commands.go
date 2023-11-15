package commands

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shell-buddy/commands"
)

var COMMAND_KEY = "commands"

var UserCommandCmd = &cobra.Command{
	Use:   "commands",
	Short: "Manage commands",
	Long:  `Add, list, and remove commands.`,
}

var addCommandCmd = &cobra.Command{
	Use:   "add [name] [description] [command]",
	Short: "Add a new command",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		description := args[1]
		command := args[2]

		commands := viper.GetStringMap(COMMAND_KEY)
		commands[name] = BashCommand{
			Description: description,
			Command:     command,
			Pinned:      false,
		}

		viper.Set(COMMAND_KEY, commands)
		err := viper.WriteConfig()
		if err != nil {
			return
		}
		fmt.Printf("Added command: %s: %s, %s\n", name, description, command)
	},
}

var listCommandCmd = &cobra.Command{
	Use:   "list",
	Short: "List all machines",
	Run: func(cmd *cobra.Command, args []string) {
		commands := viper.GetStringMap(COMMAND_KEY)
		for name, command := range commands {
			command := command.(map[string]interface{})
			pinIcon := ""
			if command["pinned"].(bool) {
				pinIcon = "\uf44d " // Nerd Fonts icon or emoji
			}
			fmt.Printf("%s%s | %s\n", pinIcon, ansi.Color(name, "blue+b"), command["command"])
		}
	},
}

var removeCommandCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a machine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		machines := viper.GetStringMap(COMMAND_KEY)
		if _, exists := machines[name]; exists {
			delete(machines, name)
			viper.Set(COMMAND_KEY, machines)
			err := viper.WriteConfig()
			if err != nil {
				return
			}
			fmt.Printf("Removed machine: %s\n", name)
		} else {
			fmt.Printf("BashCommand not found: %s\n", name)
		}
	},
}

var pinCommandCmd = &cobra.Command{
	Use:   "pin [name]",
	Short: "Pin a command and optionally set an alias",
	Args:  cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		var alias, defaultMachine string
		if len(args) > 1 {
			alias = args[1]
		}
		if len(args) > 2 {
			defaultMachine = args[2]
		}
		commands := viper.GetStringMap(COMMAND_KEY)
		if command, exists := commands[name]; exists {
			command := command.(map[string]interface{})
			command["pinned"] = true
			command["alias"] = alias
			command["defaultMachine"] = defaultMachine
			viper.Set(COMMAND_KEY, commands)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println("Error updating configuration:", err)
				return
			}
			fmt.Printf("Pinned command: %s with alias: %s on machine: %s\n", name, alias, defaultMachine)
		} else {
			fmt.Printf("Command not found: %s\n", name)
		}
	},
}

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

func init() {
	UserCommandCmd.AddCommand(addCommandCmd)
	UserCommandCmd.AddCommand(listCommandCmd)
	UserCommandCmd.AddCommand(removeCommandCmd)
	UserCommandCmd.AddCommand(pinCommandCmd)
	commands.RootCmd.AddCommand(UserCommandCmd)
}
