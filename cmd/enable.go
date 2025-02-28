package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a program to run after user logon",
	Long:  "Enable a program in the Windows registry by renaming its key back to the original name to run after user logon.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: startctl enable <name>")
			return
		}
		name := args[0]
		disabledName := name + "_disabled"
		key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
		if err != nil {
			fmt.Println("Error opening registry key:", err)
			return
		}
		defer key.Close()
		value, _, err := key.GetStringValue(disabledName)
		if err != nil {
			fmt.Println("Error retrieving disabled program:", err)
			return
		}
		err = key.SetStringValue(name, value)
		if err != nil {
			fmt.Println("Error enabling program:", err)
			return
		}
		err = key.DeleteValue(disabledName)
		if err != nil {
			fmt.Println("Error cleaning up disabled key:", err)
			return
		}
		fmt.Printf("Successfully enabled %s.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
