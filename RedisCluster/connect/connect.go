package connect

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var Cluster *redis.ClusterClient

func init() {
	Cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{
			"192.168.31.251:7001",
			"192.168.31.251:7002",
			"192.168.31.251:7003",
			"192.168.31.251:7004",
			"192.168.31.251:7005",
			"192.168.31.251:7006",
		},
		DialTimeout:        100 * time.Microsecond,
		ReadTimeout:        100 * time.Microsecond,
		WriteTimeout:       100 * time.Microsecond,
	})
}