package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bcstest",
	Short: "bcstest for helm3 charts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("t")
	},
}

// Execute execute unittest command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute bcstest error, %v", err)
		os.Exit(1)
	}
}
