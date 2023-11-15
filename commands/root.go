package commands

import (
	"github.com/spf13/cobra"
	"shell-buddy/pkg/commands"
	"shell-buddy/pkg/machines"
	"shell-buddy/pkg/utils"
)

var RootCmd = &cobra.Command{
	Use:   "Shell Buddy",
	Short: "Manage your services, commands and machines with ease",
	Long:  `Shell Buddy is a tool to manage your commands, services and machines. It is designed to be simple and easy to use.`,
}

func Execute() error {
	return RootCmd.Execute()
}
func init() {
	cobra.OnInitialize(utils.InitializeConfig)
	RootCmd.AddCommand(commands.UserCommandCmd)
	RootCmd.AddCommand(machines.MachineCmd)
}
