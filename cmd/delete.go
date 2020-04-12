package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Chrischuck/redis-cli/redis"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a key(s) argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := redis.Client

		client.Del(args...)

		message := fmt.Sprintf("Delete %s\n", strings.Join(args, ", "))
		color.Green(message)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		redis.Client.Close()
	},
}
