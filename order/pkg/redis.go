package pkg

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisUtil struct {
	client redis.Conn
}

var RedisClient RedisUtil

func NewRedisClient() RedisUtil {
	RedisClient = RedisUtil{
		client: Client,
	}
	return RedisClient
}

func (rs *RedisUtil) SetStr(key string, value string) error {
	fmt.Println("redis")
	_, err := rs.client.Do("setex", key, 86400, value)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
func (rs *RedisUtil) DelByKey(key string) error {
	_, err := rs.client.Do("DEL", key)
	return err
}

func (rs *RedisUtil) GetStr(key string) string {
	val, err := rs.client.Do("Get", key)
	if err != nil {
		return ""
	} else {
		if val == nil {
			return ""
		} else {
			return string(val.([]byte))
		}
	}
}
