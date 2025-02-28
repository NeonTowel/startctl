package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a program to run after user logon",
	Long:  "Add a program to the Windows registry so it launches after user logon.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: startctl add <name> <path>")
			return
		}
		name := args[0]
		path := args[1]
		key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
		if err != nil {
			fmt.Println("Error opening registry key:", err)
			return
		}
		defer key.Close()
		err = key.SetStringValue(name, path)
		if err != nil {
			fmt.Println("Error setting registry value:", err)
			return
		}
		fmt.Printf("Successfully added %s to startup.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
