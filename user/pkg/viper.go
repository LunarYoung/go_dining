package pkg

import (
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)

var (
	//Client redis.Conn
	Config *viper.Viper
)

func ViperInit() {

	path, vipErr := os.Getwd()
	if vipErr != nil {
		panic(vipErr)
	}

	Config = viper.New()
	Config.AddConfigPath(path)     //设置读取的文件路径
	Config.SetConfigName("config") //设置读取的文件名
	Config.SetConfigType("yaml")   //设置文件的类型

	if err := Config.ReadInConfig(); err != nil {
		panic("failed to vp:" + err.Error())
	}

	var rErr error

	Client, rErr = redis.Dial("tcp", Config.GetString("redis.host"), redis.DialDatabase(Config.GetInt("redis.db")))

	if rErr != nil {
		panic("failed to redis:" + rErr.Error())
	}

}
