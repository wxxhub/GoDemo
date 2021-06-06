package demo

import (
	"RedisCluster/connect"
	"context"
	"fmt"
	"time"
)

const SampleDemoKey = "SampleDemoKey"

func SampleDemo()  {
	// 写入数据，并设置10分钟缓存
	connect.Cluster.Set(context.TODO(), SampleDemoKey, "666", 10 * time.Minute)

	cmd := connect.Cluster.Get(context.TODO(), SampleDemoKey)

	result, err := cmd.Result()
	fmt.Println("err:", err)
	fmt.Println("result:", result)
}