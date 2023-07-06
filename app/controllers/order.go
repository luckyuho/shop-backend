package controllers

import (
	"basic/db"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

func HandleOrder() {
	// 設定處理器數量
	workerCount := 5

	// 建立 WaitGroup 來追蹤處理器的執行
	var wg sync.WaitGroup
	wg.Add(workerCount)

	// 建立指定數量的處理器
	for i := 1; i <= workerCount; i++ {
		go OrderProcessor(db.GetRedis(), i, &wg)
	}

	// 等待所有處理器完成執行
	wg.Wait()
}

func PushOrder(orderID string) error {
	// 將訂單推送到 Redis 佇列
	err := PushOrderToQueue(db.GetRedis(), orderID)
	if err != nil {
		return errors.New("無法將訂單推送到佇列")
	}
	return nil
}

// 將訂單推送到 Redis 佇列
func PushOrderToQueue(rdb *redis.Client, orderID string) error {
	// 使用 Redis 的 LPUSH 命令將訂單 ID 推送到佇列的最前面
	err := rdb.LPush(context.Background(), "order_queue", orderID).Err()
	if err != nil {
		return errors.New("推送訂單到佇列時發生錯誤")
	}
	return nil
}

// 訂單處理器
func OrderProcessor(rdb *redis.Client, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 持續從佇列中讀取訂單並進行處理
	for {
		orderID, err := PopOrderFromQueue(rdb)
		if err != nil {
			log.Printf("處理器 %d: 無法從佇列中獲取訂單: %v\n", workerID, err)
			continue
		}

		// 執行訂單處理的相關邏輯
		fmt.Printf("處理器 %d 正在處理訂單: %s\n", workerID, orderID)
		// 在這裡執行相應的訂單處理操作
		fmt.Printf("處理器 %d 完成訂單處理: %s\n", workerID, orderID)
	}
}

// 從 Redis 佇列中彈出訂單
func PopOrderFromQueue(rdb *redis.Client) (string, error) {
	// 使用 Redis 的 BRPOP 命令從佇列的尾部彈出訂單 ID，並設定超時時間
	result, err := rdb.BRPop(context.Background(), 0, "order_queue").Result()
	if err != nil {
		return "", fmt.Errorf("彈出訂單時發生錯誤: %v", err)
	}

	// 獲取彈出的訂單 ID
	orderID := result[1]

	return orderID, nil
}
