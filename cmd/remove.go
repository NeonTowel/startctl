package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a program from running after user logon",
	Long:  "Remove a program from the Windows registry so it no longer launches after user logon.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: startctl remove <name>")
			return
		}
		name := args[0]
		key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
		if err != nil {
			fmt.Println("Error opening registry key:", err)
			return
		}
		defer key.Close()
		err = key.DeleteValue(name)
		if err != nil {
			fmt.Println("Error removing registry value:", err)
			return
		}
		fmt.Printf("Successfully removed %s from startup.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
