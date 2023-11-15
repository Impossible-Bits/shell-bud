package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shell-buddy/interfaces"
	"shell-buddy/utils"
)

var machines = make([]interfaces.Machine, 0)

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "Manage machines",
	Long:  `Add, list, and remove machines.`,
}

var addMachineCmd = &cobra.Command{
	Use:   "add [name] [ip]",
	Short: "Add a new machine",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ip := args[1]
		machines := viper.GetStringMap("machines")
		machines[name] = ip
		viper.Set("machines", machines)
		err := viper.WriteConfig()
		if err != nil {
			return
		}
		fmt.Printf("Added machine: %s (%s)\n", name, ip)
	},
}

var listMachineCmd = &cobra.Command{
	Use:   "list",
	Short: "List all machines",
	Run: func(cmd *cobra.Command, args []string) {
		machines := viper.GetStringMap("machines")
		ms := make([]interfaces.Machine, 0, len(machines))
		fmt.Printf("Machines:\n")
		for name, ip := range machines {
			m := interfaces.Machine{Name: name, IP: ip.(string), Status: utils.IsMachineOnline(ip.(string))}

			var status string
			if m.Status {
				status = "Online"
			} else {
				status = "Offline"
			}

			fmt.Printf("\t%s: %s - %s\n", m.Name, m.IP, status)
			ms = append(ms, m)
		}
	},
}

var removeMachineCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a machine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		machines := viper.GetStringMap("machines")
		if _, exists := machines[name]; exists {
			delete(machines, name)
			viper.Set("machines", machines)
			err := viper.WriteConfig()
			if err != nil {
				return
			}
			fmt.Printf("Removed machine: %s\n", name)
		} else {
			fmt.Printf("Machine not found: %s\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(machineCmd)
	machineCmd.AddCommand(addMachineCmd)
	machineCmd.AddCommand(listMachineCmd)
	machineCmd.AddCommand(removeMachineCmd)
}
