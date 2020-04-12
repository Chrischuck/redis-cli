package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "redis-cli",
}

// Execute starts the cli
func Execute() {

	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(setCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
