package commands

import (
	"github.com/spf13/cobra"
	"shell-buddy/pkg/utils"
)

var RootCmd = &cobra.Command{
	Use:   "Shell Buddy",
	Short: "Manage your services, commands and machines with ease",
	Long:  `Shell Buddy is a tool to manage your commands, services and machines. It is designed to be simple and easy to use.`,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	cobra.OnInitialize(utils.InitializeConfig)
}
