package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "Shell Buddy",
	Short: "Manage your services, commands and machines with ease",
	Long:  `Shell Buddy is a tool to manage your commands, services and machines. It is designed to be simple and easy to use.`,
}

// Execute initializes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory. Using current directory instead.")
		home = "."
	}

	configPath := filepath.Join(home, ".config/sb")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found. A new one will be created upon adding machines.")
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
