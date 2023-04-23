package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBindKey(t *testing.T) {
	expect := "ms_example:f1:a:1:b:2"
	redisService := RedisService{}
	actual := redisService.BindKey("f1", map[string]string{"a": "1", "b": "2"})
	assert.Equal(t, expect, actual)
}

func TestIsConnected(t *testing.T) {
	ctx := context.Background()
	expect := true
	redisService := RedisService{}
	actual := redisService.GetInstant(ctx).IsConnected(ctx)
	assert.Equal(t, expect, actual)
}

func TestSet(t *testing.T) {
	ctx := context.Background()
	expect := true
	redisService := RedisService{}
	key := redisService.BindKey("f1", map[string]string{"page": "1", "limit": "10"})
	redisIns := redisService.GetInstant(ctx)
	actual := redisIns.Set(ctx, key, "heheboi", "60s")
	assert.Equal(t, expect, actual)
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	expect := "heheboi"
	redisService := RedisService{}
	key := redisService.BindKey("f1", map[string]string{"page": "1", "limit": "10"})
	redisIns := redisService.GetInstant(ctx)
	actual := redisIns.Get(ctx, key)
	assert.Equal(t, expect, actual)
}

/*
package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBindKey(t *testing.T) {
	expect := "ms_setting:f1:a:1:b:2"
	actual := BindKey("f1", map[string]string{"a": "1", "b": "2"})
	assert.Equal(t, expect, actual)
}

func TestCachingGet(t *testing.T) {
	redisUtil := GetInstance(context.Background())
	redisUtil.Set(context.Background(), "ms_setting:test", "testnha", 5*time.Second)

	data1, _ := redisUtil.Get(context.Background(), "ms_setting:test")
	assert.Equal(t, "testnha", data1)
	assert.NotEqual(t, "testnhaaaaa", data1)

	data2, _ := redisUtil.Get(context.Background(), "ms_setting:testtt")
	assert.Equal(t, "", data2)
}

func TestCachingSet(t *testing.T) {
	redisUtil := GetInstance(context.Background())
	redisUtil.Set(context.Background(), "ms_setting:test2", "test2nha", 5*time.Second)
	data, _ := redisUtil.Get(context.Background(), "ms_setting:test2")
	assert.NotEqual(t, "", data)
	assert.Equal(t, "test2nha", data)
}

func TestRemove(t *testing.T) {
	redisUtil := GetInstance(context.Background())
	redisUtil.Set(context.Background(), "t1", "test1", 5*time.Second)
	redisUtil.Remove(context.Background(), "t1")
	data, _ := redisUtil.Get(context.Background(), "t1")
	assert.NotEqual(t, "test1", data)
	assert.Equal(t, "", data)
}

func TestKeys(t *testing.T) {
	redisUtil := GetInstance(context.Background())
	redisUtil.Set(context.Background(), CACHING_PREFIX_MAIN+":test1", "setting1", 5*time.Second)
	redisUtil.Set(context.Background(), CACHING_PREFIX_MAIN+":test2", "setting2", 5*time.Second)
	redisUtil.Set(context.Background(), "test", "test", 5*time.Second)
	keys, _ := redisUtil.Keys(context.Background())
	assert.Equal(t, []string{"ms_setting:test1", "ms_setting:test2"}, keys)
}

func TestCachingClear(t *testing.T) {
	redisUtil := GetInstance(context.Background())
	redisUtil.Set(context.Background(), "ms_setting:test", "ms_setting:test", 5*time.Second)
	redisUtil.Set(context.Background(), "test", "test", 5*time.Second)
	redisUtil.Clear(context.Background())
	data1, _ := redisUtil.Get(context.Background(), "ms_setting:test")
	data2, _ := redisUtil.Get(context.Background(), "test")
	assert.Equal(t, "", data1)
	assert.Equal(t, "test", data2)
}

func TestConnected(t *testing.T) {
	redisUtil := GetInstance(context.Background())
	assert.NotNil(t, redisUtil)
}

func TestDisconnected(t *testing.T) {
	redisUtil := GetInstance(context.Background())
	assert.Nil(t, redisUtil)
}

*/
