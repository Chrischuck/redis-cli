package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Chrischuck/redis-cli/redis"
	"github.com/spf13/cobra"
)

func getTTL(ttlArg string) (int, error) {
	ttlParam, paramErr := strconv.Atoi(ttlArg)
	if paramErr == nil {
		return ttlParam, nil
	}

	ttlEnv, envErr := strconv.Atoi(os.Getenv("REDIS_TTL"))
	if envErr == nil {
		return ttlEnv, nil
	}

	return 0, errors.New("invalid TTL parameter")
}

var setCmd = &cobra.Command{
	Use: "set",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires a key and value arguments")
		}

		if len(args) > 4 {
			return errors.New("can only have two arguments")
		}

		_, ttlError := getTTL(args[2])

		if ttlError != nil {
			return ttlError
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := redis.Client

		key := args[0]
		value := args[1]
		ttl, _ := getTTL(args[2])

		timeString := fmt.Sprintf("%dms", ttl)
		ttlInMs, _ := time.ParseDuration(timeString)

		client.Set(key, value, ttlInMs)
		fmt.Printf("Set %s: %s\n", key, value)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		redis.Client.Close()
	},
}
