package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shell-buddy/interfaces"
)

var COMMAND_KEY = "commands"

var commandCmd = &cobra.Command{
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
		commands[name] = interfaces.BashCommand{
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
			fmt.Printf("%s(%b) | %s(%s)\n", name, command["pinned"], command["command"], command["description"])
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
	Short: "Pin a command",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		commands := viper.GetStringMap(COMMAND_KEY)
		if len(args) == 1 {
			name = args[0]
			if command, exists := commands[name]; exists {
				command := command.(map[string]interface{})
				command["pinned"] = true
				viper.Set(COMMAND_KEY, commands)
				err := viper.WriteConfig()
				if err != nil {
					return
				}
				fmt.Printf("Pinned machine: %s\n", name)
			}
		} else {
			for name, command := range commands {
				command := command.(map[string]interface{})
				if command["pinned"] == true {
					fmt.Printf("`%s` | %s(%s)\n", name, command["command"], command["description"])
				}
			}
		}
	},
}

func init() {
	commandCmd.AddCommand(addCommandCmd)
	commandCmd.AddCommand(listCommandCmd)
	commandCmd.AddCommand(removeCommandCmd)
	commandCmd.AddCommand(pinCommandCmd)
	rootCmd.AddCommand(commandCmd)
}
