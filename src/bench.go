package main

import (
	"flag"
	"fmt"
	"gopkg.in/redis.v4"
	"strings"
	// "sync"
	"time"
)

var client *redis.Client
var redis_addrs string
var pool_size int
var max_counts int
var redis_key string

func main() {
	flag.IntVar(&pool_size, "pool", 1000, "pool size")
	flag.IntVar(&max_counts, "max_counts", 50, "concurrence")
	flag.StringVar(&redis_addrs, "addrs", "", "redis cluster addr")
	flag.StringVar(&redis_key, "key", "test", "redis get key")
	flag.Parse()
	if len(redis_addrs) == 0 {
		flag.PrintDefaults()
		return
	}

	cluster_addrs := []string{redis_addrs}
	if strings.Contains(redis_addrs, ",") {
		cluster_addrs = strings.Split(redis_addrs, ",")
	}

	for i := 0; i < max_counts; i++ {
		conStart := time.Now()
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:          cluster_addrs,
			PoolSize:       1000,
			ReadOnly:       false,
			RouteByLatency: false,
		})
		// client.Ping()
		_, err := client.Ping().Result()
		connElapsed := time.Since(conStart)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("redis connection init time: ", connElapsed)
		}
	}

	// var wg sync.WaitGroup
	// connStart := time.Now()
	// for i := 0; i < max_counts; i++ {
	// 	wg.Add(1)
	// 	go client.Get(redis_key)
	// }
	// wg.Wait()
	// connElapsed := time.Since(connStart)
	// fmt.Println("max count:%v time: %v", max_counts, connElapsed)
	// fmt.Println("USING CONN : Written %v records to redis: %v", max_counts, connElapsed)
}
