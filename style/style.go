package style

import (
	"github.com/fatih/color"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/viper"
	"shell-buddy/pkg/machines"
)

func GetStyle(styleName string) *color.Color {
	styleConfig := viper.GetStringMapString("styles." + styleName)
	// Initialize a default color
	c := color.New(color.FgWhite)

	// Apply configuration
	if val, ok := styleConfig["color"]; ok {
		c = color.New(getColor(val))
	}
	return c
}

func getColor(colorName string) color.Attribute {
	switch colorName {
	case "red":
		return color.FgRed
	// Add cases for other colors
	default:
		return color.FgWhite
	}
}

func DisplayMachineTable(machines []machines.Machine) error {
	if err := termui.Init(); err != nil {
		return err
	}
	defer termui.Close()

	table := widgets.NewTable()
	table.Rows = [][]string{
		{"Name", "IP Address", "Status"},
	}

	for _, machine := range machines {
		status := "Offline" // Placeholder, add actual check
		table.Rows = append(table.Rows, []string{machine.Name, machine.IP, status})
	}

	table.TextStyle = termui.NewStyle(termui.ColorWhite)
	table.SetRect(0, 0, 50, len(machines)+2) // Adjust size as needed

	termui.Render(table)

	// Wait for user input to exit
	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		if e.Type == termui.KeyboardEvent {
			break
		}
	}

	return nil
}
