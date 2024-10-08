package cache

import (
	"log"
	"os"

	"github.com/NoCapCbas/webadmin/queue"
	"github.com/go-redis/redis"
)

var rc *redis.Client

func init() {
	host := os.Getenv("REDIS_ADDR")
	if len(host) == 0 {
		host = "redis:6379"
	}

	c := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: os.Getenv("REDIS_KEY"),
		DB:       0, // use default DB
	})

	if _, err := c.Ping().Result(); err != nil {
		log.Fatal("unable to connect to redis", err)
	}

	rc = c
}

func New(queueProcessor, isDev bool) {
	queue.New(rc, isDev)

	if queueProcessor {
		queue.SetAsSubscriber()
	}
}
