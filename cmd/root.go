package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "startctl",
	Short: "Manage programs that run after user logon",
	Long: `startctl is a CLI tool for managing programs that run after user logon on Windows.

Features include:
- Add programs to run after logon
- Remove programs from running after logon
- Enable/Disable programs without removing them
- List all programs set to run after logon

Usage examples:
- startctl add "MyApp" "C:\Path\To\MyApp.exe"
- startctl remove "MyApp"
- startctl enable "MyApp"
- startctl disable "MyApp"
- startctl list`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
