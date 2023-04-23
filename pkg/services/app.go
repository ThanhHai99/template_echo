package services

import (
	"context"
	"time"
)

func GetData(page string, limit string) string {
	ctx := context.Background()
	redisService := RedisService{}
	key := redisService.BindKey("GatData", map[string]string{"page": page, "limit": limit})
	redisIns := redisService.GetInstant(ctx)
	dataRedis := redisIns.Get(ctx, key)
	if dataRedis != "" {
		return dataRedis
	} else {
		data := time.Now().String()
		redisIns.Set(ctx, key, data, "60s")
		return data
	}
}
