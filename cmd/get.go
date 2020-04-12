package cmd

import (
	"errors"
	"fmt"

	"github.com/Chrischuck/redis-cli/redis"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: "get",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a key argument")
		}

		if len(args) > 1 {
			return errors.New("can only have one key argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		client := redis.Client

		value := client.Get(key).Val()

		message := fmt.Sprintf("Get %s: %v\n", key, value)
		color.Green(message)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		redis.Client.Close()
	},
}
