package save

import (
	"context"
	"maker/log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	RedisServer         string
	RedisServerPassword string
	RedisClient         *redis.Client
)

func init() {
	RedisServer = os.Getenv("REDIS_SERVER_ADDRESS")
	RedisServerPassword = os.Getenv("REDIS_SERVER_PASSWORD")
	if RedisServer == "" {
		log.Lx.WithFields(logrus.Fields{
			"REDIS_SERVER_ADDRESS": RedisServer,
		}).Fatal("Failed to initail environments setting, pls configure & Reboot.")

	} else {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     RedisServer,
			Password: RedisServerPassword,
		})
		pong, err := RedisClient.Ping(context.TODO()).Result()
		if err != nil {
			log.Lx.WithFields(logrus.Fields{
				"to":     RedisServer,
				"result": pong,
			}).Fatal("Failed to connect Redis.")

		}
	}
}

func SaveAq2Redis(city string, buf []byte) error {

	pipeline := RedisClient.Pipeline()
	ctx := context.TODO()
	pipeline.HSet(ctx, "air_quality_cache", city, buf)
	cmds, err := pipeline.Exec(ctx)

	for _, cmd := range cmds {
		log.Lx.WithFields(logrus.Fields{
			"args": cmd.Args(),
			"name": cmd.FullName(),
		}).Debug("Execute Redis' cmd")
	}

	return err
}

func QueryAq4Redis(city string) ([]byte, error) {
	ctx := context.TODO()
	return RedisClient.HGet(ctx, "air_quality_cache", city).Bytes()

}
