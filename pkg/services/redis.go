package services

import (
	"Template_Echo/pkg/configs"
	"Template_Echo/pkg/constants"
	"Template_Echo/pkg/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"reflect"
	"sort"
	"time"
)

type RedisService struct {
	client *redis.Client
}

var (
	redisInstants  *RedisService
	lostConnection bool
)

func (s *RedisService) GetInstant(ctx context.Context) *RedisService {
	var isInitialized bool
	if redisInstants == nil {
		isInitialized = false
	} else {
		if e0 := redisInstants.client.Ping(ctx).Err(); e0 != nil {
			isInitialized = false
		} else {
			isInitialized = true
		}
	}

	if isInitialized == true {
		return redisInstants
	} else {
		redisOptions := &redis.Options{
			Addr:        configs.RedisHost(),
			Password:    "",
			DB:          0, // use default DB
			DialTimeout: 150 * time.Millisecond,
			ReadTimeout: 100 * time.Millisecond,
			MaxRetries:  0,
		}

		rdb := redis.NewClient(redisOptions)
		redisInstants = &RedisService{client: rdb}
		return redisInstants
	}
}

func (s *RedisService) IsConnected(ctx context.Context) bool {
	if s.client == nil {
		utils.Log().Error().Msg("REDIS IS NOT READY")
		lostConnection = true
		return false
	}

	e0 := s.client.Ping(ctx).Err()
	if e0 != nil {
		utils.Log().Error().Msg("REDIS IS NOT READY")
		lostConnection = true
		return false
	}

	utils.Log().Info().Msg("REDIS IS READY")
	lostConnection = false
	return true
}

func (s *RedisService) BindKey(funcName string, query map[string]string) string {
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

func (s *RedisService) Set(ctx context.Context, key string, value string, expireTime string) bool {
	cachingPower := configs.RedisPower()
	if cachingPower == constants.OFF.String() {
		return false
	}
	expireTimes, e0 := time.ParseDuration(expireTime)
	if expireTime == "" || e0 != nil {
		expireTimeDefault := fmt.Sprintf("%ds", constants.REDIS_EXPIRE_TIME_DEFAULT)
		expireTimes, _ = time.ParseDuration(expireTimeDefault)
	}

	if e1 := redisInstants.client.Set(ctx, key, value, expireTimes).Err(); e1 == nil {
		fmt.Printf("%s:cache:set\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_SUCCESS_STATUS)
	} else {
		fmt.Printf("%s:cache:set\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_FAIL_STATUS)
	}

	return true
}

func (s *RedisService) Get(ctx context.Context, key string) string {
	cachingPower := configs.RedisPower()
	if cachingPower == constants.OFF.String() {
		return ""
	}
	result, e0 := redisInstants.client.Get(ctx, key).Result()
	if e0 != nil {
		fmt.Printf("%s:cache:get\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_FAIL_STATUS)
		fmt.Println(e0)
		return ""
	}
	fmt.Printf("%s:cache:get\t %s ==> %s \n", constants.REDIS_PREFIX_MAIN, key, constants.REDIS_SUCCESS_STATUS)
	return result
}

func (s *RedisService) Keys(ctx context.Context) ([]string, error) {
	return s.client.Keys(ctx, constants.REDIS_PREFIX_MAIN+"*").Result()
}

func (s *RedisService) Clear(ctx context.Context) {
	keys, _ := s.Keys(ctx)
	pipe := s.client.Pipeline()
	for _, key := range keys {
		_ = pipe.Del(ctx, key).Err()
	}
	_, _ = pipe.Exec(ctx)

	utils.Log().Warn().Msg("RedisClient.Clear")
}

/*
package model

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
	_log "github.com/sirupsen/logrus"
	conf "ms-setting/config"
	"reflect"
	"sort"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}

var (
	instanceRedis *RedisClient

	redisConfig     = conf.RedisConfig{}
	mainConfig      = conf.Config{}
	_               = env.Parse(&redisConfig)
	_               = env.Parse(&mainConfig)
	lostConnections bool // used to lost connection
)

func GetInstance(ctx context.Context) *RedisClient {
	if !redisConfig.RedisInUse || mainConfig.AppEnv == "stg" {
		return nil
	}
	if instanceRedis == nil {
		rdOption := &redis.Options{
			Addr:        redisConfig.RedisHost,
			Password:    redisConfig.RedisPass,
			DB:          0, // use default DB
			DialTimeout: 150 * time.Millisecond,
			ReadTimeout: 100 * time.Millisecond,
			MaxRetries:  0,
		}
		if mainConfig.AppEnv == "prd" {
			rdOption.TLSConfig = &tls.Config{}
		}
		rdb := redis.NewClient(rdOption)
		instanceRedis = &RedisClient{rdb}
		return instanceRedis
	} else {
		if instanceRedis.connected(ctx) {
			_log.Info("REDIS IS READY")
			if lostConnections {
				instanceRedis.Clear(ctx)
			}
			return instanceRedis
		}
		_log.Error("REDIS IS NOT READY")
		return nil
	}
}

func (r *RedisClient) connected(ctx context.Context) bool {
	if instanceRedis == nil {
		lostConnections = true
		return false
	}

	err := instanceRedis.Client.Ping(ctx).Err()
	if err != nil {
		lostConnections = true
		return false
	}

	lostConnections = false
	return true
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, duration time.Duration) {
	err := r.Client.Set(ctx, key, value, duration).Err()
	_log.SetFormatter(&_log.JSONFormatter{})
	_log.WithFields(_log.Fields{"key": key}).Warning("RedisClient.Set")
	if err != nil {
		_log.Error(err)
	}
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	_log.SetFormatter(&_log.JSONFormatter{})
	_log.WithField("key", key).Warning("RedisClient.Get")
	s, err := r.Client.Get(ctx, key).Result()

	return s, err
}

func (r *RedisClient) Gets(ctx context.Context, keys []string) ([]interface{}, error) {
	s, err := r.Client.MGet(ctx, keys...).Result()
	_log.SetFormatter(&_log.JSONFormatter{})
	_log.Warning("RedisClient.Get")

	return s, err
}

func (r *RedisClient) Remove(ctx context.Context, key string) error {
	err := r.Client.Del(ctx, key).Err()
	_log.SetFormatter(&_log.JSONFormatter{})
	_log.WithField("key", key).Warning("RedisClient.Remove")

	return err
}

//func (r *RedisClient) Expire(ctx context.Context, key string, timeToLive time.Duration) error {
//	return r.Client.Expire(ctx, key, timeToLive).Err()
//}

func (r *RedisClient) Clear(ctx context.Context) {
	keys, _ := r.Keys(ctx)
	pipe := r.Client.Pipeline()
	for _, key := range keys {
		pipe.Del(ctx, key).Err()
	}
	pipe.Exec(ctx)
	_log.SetFormatter(&_log.JSONFormatter{})
	_log.Warning("RedisClient.Clear")
}

func BindKey(funcName string, query map[string]string) string {
	finalKey := CACHING_PREFIX_MAIN + CACHING_SEPARATOR + funcName

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
		finalKey += (CACHING_SEPARATOR + k + CACHING_SEPARATOR + query[k])
	}

	return fmt.Sprintf(finalKey)
}

func (r *RedisClient) Keys(ctx context.Context) ([]string, error) {
	s, err := r.Client.Keys(ctx, CACHING_PREFIX_MAIN+"*").Result()
	_log.SetFormatter(&_log.JSONFormatter{})
	_log.Warning("RedisClient.Keys")

	return s, err
}

*/
