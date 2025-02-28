package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all programs set to run after user logon",
	Long:  "List all programs currently set to run after user logon in Windows.",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE)
		if err != nil {
			fmt.Println("Error opening registry key:", err)
			return
		}
		defer key.Close()

		valueNames, err := key.ReadValueNames(0)
		if err != nil {
			fmt.Println("Error reading registry values:", err)
			return
		}

		fmt.Println("Startup Programs:")
		for _, name := range valueNames {
			value, _, err := key.GetStringValue(name)
			if err != nil {
				fmt.Printf("Error reading value for %s: %v\n", name, err)
				continue
			}
			fmt.Printf("- %s: %s\n", name, value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
