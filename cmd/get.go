package cmd

import (
	"errors"

	"github.com/Chrischuck/redis-cli/redis"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: "get",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a key argument")
		}

		if len(args) > 1 {
			return errors.New("Can only have one key argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := redis.Client
		value := client.Get(args[0]).String()
		println(value)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		redis.Client.Close()
	},
}
