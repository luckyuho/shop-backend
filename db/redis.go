package db

import (
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	// 建立 Redis 客戶端
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 設定 Redis 伺服器的位址
		Password: "",               // 若有密碼，請填入對應密碼
		DB:       0,                // 選擇使用的資料庫編號
	})
}

func GetRedis() *redis.Client {
	return rdb
}
