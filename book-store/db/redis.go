package db

import "github.com/redis/go-redis/v9"

var RDB *redis.Client

func InitRedis(addr string, password string, dbcount int) {
	// 创建一个 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // 如果没有密码，则留空
		DB:       dbcount,  // 使用默认数据库
	})
	RDB = rdb
}
