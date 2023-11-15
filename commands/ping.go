package commands

import (
	"fmt"
	"github.com/spf13/viper"
	"self-hosted-manager/interfaces"
	"self-hosted-manager/style"
	"self-hosted-manager/utils"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping machines to check their status",
	Long:  `This command pings all machines listed in the configuration to check if they are online or offline.`,
	Run: func(cmd *cobra.Command, args []string) {
		pingMachines()
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

func pingMachines() {
	machinesData := viper.GetStringMapString("machines")
	var machines []interfaces.Machine

	for name, ip := range machinesData {
		machines = append(machines, interfaces.Machine{Name: name, IP: ip, Status: utils.IsMachineOnline(ip)})
	}

	// Display machine table using term-ui
	if err := style.DisplayMachineTable(machines); err != nil {
		fmt.Println("Error displaying machine table:", err)
	}
}
