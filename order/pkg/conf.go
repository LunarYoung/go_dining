package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	remote "github.com/yoyofxteam/nacos-viper-remote"
)

func Nacos() {
	remoteViper := viper.New()

	remote.SetOptions(&remote.Option{
		Url:         "139.159.182.159",               // nacos server 多地址需要地址用;号隔开，如 Url: "loc1;loc2;loc3"
		Port:        8850,                            // nacos server端口号
		NamespaceId: "public",                        // nacos namespace
		GroupName:   "DEFAULT_GROUP",                 // nacos group
		Config:      remote.Config{DataId: "dining"}, // nacos DataID
		Auth:        nil,                             // 如果需要验证登录,需要此参数
	})

	err := remoteViper.AddRemoteProvider("nacos", "localhost", "")
	if err != nil {
		panic(err)
	}
	remoteViper.SetConfigType("yaml")

	_ = remoteViper.ReadRemoteConfig()           //sync get remote configs to remoteViper instance memory . for example , remoteViper.GetString(key)
	_ = remoteViper.WatchRemoteConfigOnChannel() //异步监听Nacos中的配置变化，如发生配置更改，会直接同步到 viper实例中。

	appName := remoteViper.GetString("rpcHost") // sync get config by key
	a := remoteViper.GetString("redis.host")    // sync get config by key

	fmt.Println(appName)
	fmt.Println(a)

}
