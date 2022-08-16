package pkg

import (
	"github.com/spf13/viper"
	remote "github.com/yoyofxteam/nacos-viper-remote"
)

var RemoteViper *viper.Viper

func Nacos() {
	RemoteViper = viper.New()

	remote.SetOptions(&remote.Option{
		Url:         "139.159.182.159",               // nacos server 多地址需要地址用;号隔开，如 Url: "loc1;loc2;loc3"
		Port:        8850,                            // nacos server端口号
		NamespaceId: "public",                        // nacos namespace
		GroupName:   "DEFAULT_GROUP",                 // nacos group
		Config:      remote.Config{DataId: "dining"}, // nacos DataID
		Auth:        nil,                             // 如果需要验证登录,需要此参数
	})

	err := RemoteViper.AddRemoteProvider("nacos", "localhost", "")
	if err != nil {
		panic(err)
	}
	RemoteViper.SetConfigType("yaml")

	_ = RemoteViper.ReadRemoteConfig()           //sync get remote configs to RemoteViper instance memory . for example , RemoteViper.GetString(key)
	_ = RemoteViper.WatchRemoteConfigOnChannel() //异步监听Nacos中的配置变化，如发生配置更改，会直接同步到 viper实例中。

}
