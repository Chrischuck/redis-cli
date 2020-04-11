package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "My subcommand",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("before subCmd Run with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside subCmd Run with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
	},
}
