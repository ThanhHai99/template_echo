package services

import (
	"Template_Echo/pkg/configs"
	"Template_Echo/pkg/constants"
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"sort"
	"time"
)

var RedisClient *redis.Client = nil

func Redis() *redis.Client {
	var isInitialized bool
	if RedisClient == nil {
		isInitialized = false
	} else {
		if e0 := RedisClient.Ping().Err(); e0 != nil {
			isInitialized = false
		} else {
			isInitialized = true
		}
	}

	if isInitialized == true {
		return RedisClient
	} else {
		RedisClient = redis.NewClient(&redis.Options{
			Network:            "",
			Addr:               configs.RedisHost(),
			Dialer:             nil,
			OnConnect:          nil,
			Password:           "",
			DB:                 0,
			MaxRetries:         0,
			MinRetryBackoff:    0,
			MaxRetryBackoff:    0,
			DialTimeout:        0,
			ReadTimeout:        0,
			WriteTimeout:       0,
			PoolSize:           0,
			MinIdleConns:       0,
			MaxConnAge:         0,
			PoolTimeout:        0,
			IdleTimeout:        0,
			IdleCheckFrequency: 0,
			TLSConfig:          nil,
		})

		return RedisClient
	}
}

func IsConnected() bool {
	//if Redis().Ping()

	Redis()
	return true
}

func BindKey(funcName string, query map[string]string) string {
	finalKey := fmt.Sprintf("%s%s%s", constants.REDIS_PREFIX_MAIN, constants.REDIS_SEPARATOR, funcName)

	// delete item if is zero
	for key, value := range query {
		isZero := reflect.ValueOf(value).IsZero()
		isEmpty := value == "0"
		if isZero || isEmpty {
			delete(query, key)
		}
	}

	// sort key
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		finalKey += fmt.Sprintf("%s%s%s", constants.REDIS_SEPARATOR, k, constants.REDIS_SEPARATOR+query[k])
	}

	return finalKey
}

func Set(key string, value interface{}, expireTime string) interface{} {
	cachingPower := configs.RedisPower()
	if cachingPower == constants.OFF.String() {
		return nil
	}
	expireTimes, e0 := time.ParseDuration(expireTime)
	if expireTime == "" || e0 != nil {
		expireTimeDefault := fmt.Sprintf("%ds", constants.REDIS_EXPIRE_TIME_DEFAULT)
		expireTimes, _ = time.ParseDuration(expireTimeDefault)
	}
	result := Redis().Set(key, value, expireTimes)
	if result != nil {
		fmt.Printf("%s:cache:set\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_SUCCESS_STATUS)
	} else {
		fmt.Printf("%s:cache:set\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_FAIL_STATUS)
	}

	return result
}

func Get(key string) interface{} {
	cachingPower := configs.RedisPower()
	if cachingPower == constants.OFF.String() {
		return nil
	}
	result := Redis().MGet(key)
	if result != nil {
		fmt.Printf("%s:cache:get\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_SUCCESS_STATUS)
	} else {
		fmt.Printf("%s:cache:get\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_FAIL_STATUS)
	}

	return result
}

func GetRedis() {
}

func SetRedis() {
}
