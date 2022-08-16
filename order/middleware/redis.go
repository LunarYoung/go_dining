package middleware

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var Client redis.Conn

func InitRedis() {
	var rErr error
	Client, rErr = redis.Dial("tcp", RemoteViper.GetString("redis.host"), redis.DialDatabase(RemoteViper.GetInt("redis.db")))
	if rErr != nil {
		panic("failed to redis:" + rErr.Error())
	}
}

func SetStr(key string, value string) error {
	_, err := Client.Do("setex", key, 86400, value)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func SetMap(mapName string, k string, v string) error {
	_, err := Client.Do("HSet", mapName, k, v)
	return err
}

func GetMap(mapName string, k string) string {
	val, err := Client.Do("HGet", mapName, k)
	if err != nil {
		return ""
	}
	return string(val.([]byte))
}
func DelMap(mapName string, k string) error {
	//_, err := Client.Do("HDel", mapName, k)
	_, err := Client.Do("HDel", "books", "abc")
	return err
}

func DelByKey(key string) error {
	_, err := Client.Do("DEL", key)
	return err
}

func GetStr(key string) string {
	val, err := Client.Do("Get", key)
	if err != nil {
		return " "
	} else {
		if val == nil {
			return ""
		} else {
			return string(val.([]byte))
		}
	}
}
