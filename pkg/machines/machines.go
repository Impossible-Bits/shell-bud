package machines

import (
	"fmt"
	"net"
	"time"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var MachineCmd = &cobra.Command{
	Use:   "machine",
	Short: "Manage machines",
	Long:  `Add, list, and remove machines.`,
}

var (
	addMachineCmd = &cobra.Command{
		Use:   "add [name] [ip]",
		Short: "Add a new machine",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			ip := args[1]
			machines := viper.GetStringMap(MACHINE_KEY)
			machines[name] = ip
			viper.Set(MACHINE_KEY, machines)
			err := viper.WriteConfig()
			if err != nil {
				return
			}
			fmt.Printf("Added machine: %s (%s)\n", name, ip)
		},
	}

	listMachineCmd = &cobra.Command{
		Use:   "list",
		Short: "List all machines",
		Run: func(cmd *cobra.Command, args []string) {
			machines := viper.GetStringMap(MACHINE_KEY)
			fmt.Printf("Machines:\n")
			for name, ip := range machines {
				m := Machine{Name: name, IP: ip.(string), Status: isMachineOnline(ip.(string))}
				statusColor := "red+b"
				status := "Offline"
				if m.Status {
					status = "Online"
					statusColor = "green+b"
				}
				fmt.Printf("\t%s: %s - %s\n", ansi.Color(m.Name, "yellow+b"), m.IP, ansi.Color(status, statusColor))
			}
		},
	}

	removeMachineCmd = &cobra.Command{
		Use:   "remove [name]",
		Short: "Remove a machine",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			machines := viper.GetStringMap(MACHINE_KEY)
			if _, exists := machines[name]; exists {
				delete(machines, name)
				viper.Set(MACHINE_KEY, machines)
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
)

func isMachineOnline(ip string) bool {
	timeout := time.Second * 1
	_, err := net.DialTimeout("tcp", ip+":22", timeout)
	return err == nil
}

func getMachinesMap() map[string]Machine {
	machineList := make(map[string]Machine)
	for key, raw := range viper.GetStringMap("machines") {
		machineMap := raw.(map[string]interface{})
		machine := Machine{
			Name:   machineMap["name"].(string),
			IP:     machineMap["ip"].(string),
			Status: machineMap["status"].(bool),
		}
		machineList[key] = machine
	}
	return machineList
}

func setMachinesMap(machines map[string]Machine) {
	viper.Set(MACHINE_KEY, machines)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error writing machines to config: %s\n", err)
	}
}

func init() {
	MachineCmd.AddCommand(addMachineCmd)
	MachineCmd.AddCommand(listMachineCmd)
	MachineCmd.AddCommand(removeMachineCmd)
}
