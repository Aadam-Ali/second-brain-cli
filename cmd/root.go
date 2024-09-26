package cmd

import (
	"fmt"
	"os"

	"github.com/aadam-ali/second-brain-cli/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sb",
	Short: "sb is second brain management tool",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.GetConfig()

		fmt.Println("Hello! Welcome to sb.")
		fmt.Printf("Root: %s\n", config.Root)
		fmt.Printf("Inbox: %s\n", config.Inbox)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
