package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sb",
	Short: "sb is second brain management tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello! Welcome to sb.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
