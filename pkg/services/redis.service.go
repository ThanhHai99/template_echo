package services

import (
	"Template_Echo/pkg/configs"
	"Template_Echo/pkg/constants"
	"fmt"
	"github.com/go-redis/redis"
)

func Init() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: configs.RedisHost(),
		Password: "",
		DB: 0,
	})
}

func BindKey(prefix string, data [][]interface{}) string {
	cachingPower := configs.RedisPower()
	if cachingPower == constants.OFF.String() {
		return ""
	}

	key := fmt.Sprintf("%s%s%s", constants.REDIS_PREFIX_MAIN, constants.REDIS_SEPARATOR, prefix)
	for i := 0; i < len(data); i++ {
		for item := 0; item < len(data[i]); item++ {
			if data[i][item] == nil {
				continue
			}
			key += fmt.Sprintf("%s%d%s", constants.REDIS_SEPARATOR, item, data[i][item])
		}
	}

	return key
}

func Set(key string, value interface{}, expireTime int16) interface{} {
	//const cachingPower = this.configService.get('redis.power')
	//if (cachingPower == CachingEnum.OFF) return null
	//
	//let result = null
	//if (expireTime) result = (await this.redis.set(key, value, 'EX', expireTime)) ? REDIS_SUCCESS_STATUS : REDIS_FAIL_STATUS
	//else result = (await this.redis.set(key, value, 'EX', REDIS_EXPIRE_TIME_DEFAULT)) ? REDIS_SUCCESS_STATUS : REDIS_FAIL_STATUS
	//this.logger.log(REDIS_PREFIX_MAIN + ':cache:set\t' + key + ' ==> ' + result)
	//return result

	cachingPower := configs.RedisPower()
	if cachingPower == constants.OFF.String() {
		return nil
	}

	result :=
}

func GetRedis() {
}

func SetRedis() {
}
