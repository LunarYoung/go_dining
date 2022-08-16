package websocket

import (
	"fmt"
	"github.com/robfig/cron"

	"sync"
	"time"
)

var lock sync.Mutex

func Time() {
	c := cron.New()
	// 给对象增加定时任务
	err := c.AddFunc("@every 500s", func() {
		myCron()
	})
	if err != nil {
		return
	}
	c.Start()
}

func myCron() {

	if len(Manager.Clients) == 0 {
	} else {
		lock.Lock()
		for k, v := range Manager.Clients {

			timeFlag, _ := time.ParseDuration("3000s")
			timeLength := time.Now().Sub(v.LastOnlineTime)

			if timeLength > timeFlag {
				err := v.Socket.Close()
				if err != nil {
					//log.Log.Info("用户:" + v.Id + "关闭socket失败")
					fmt.Println(err.Error())
				} else {
					//log.Log.Info("用户:" + v.Id + "关闭socket成功")
					//_ = cache.RedisClient.DelByKey(v.Id)
					delete(Manager.Clients, k)
				}
			}
		}
		lock.Unlock()
	}
}
