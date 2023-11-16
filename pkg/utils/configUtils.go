package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// InitializeConfig initializes the Viper configuration.
func InitializeConfig() {
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

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found. Creating a new one.")
		createConfigFile(configPath)
	}
	//fmt.Println("Using config file:", viper.ConfigFileUsed())
}

// createConfigFile creates a new YAML configuration file.
func createConfigFile(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Printf("Failed to create config directory: %s\n", err)
		return
	}

	configFilePath := filepath.Join(path, "config.yaml")
	file, err := os.Create(configFilePath)
	if err != nil {
		fmt.Printf("Failed to create config file: %s\n", err)
		return
	}
	defer file.Close()

	defaultConfig := []byte("commands: {}\nmachines: {}\n")
	if _, err := file.Write(defaultConfig); err != nil {
		fmt.Printf("Failed to write to config file: %s\n", err)
		return
	}

	fmt.Println("Created new config file:", configFilePath)
}
