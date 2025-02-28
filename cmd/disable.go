package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a program from running after user logon",
	Long:  "Disable a program in the Windows registry by renaming its key to prevent it from launching after user logon.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: startctl disable <name>")
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
		value, _, err := key.GetStringValue(name)
		if err != nil {
			fmt.Println("Error retrieving program:", err)
			return
		}
		err = key.SetStringValue(disabledName, value)
		if err != nil {
			fmt.Println("Error disabling program:", err)
			return
		}
		err = key.DeleteValue(name)
		if err != nil {
			fmt.Println("Error cleaning up original key:", err)
			return
		}
		fmt.Printf("Successfully disabled %s.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
