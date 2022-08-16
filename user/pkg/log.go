package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var (
	Log *logrus.Logger
)

// 日志记录到文件
func Logger() gin.HandlerFunc {
	logFilePath := RemoteViper.GetString("log.path")
	logFileName := RemoteViper.GetString("log.name")
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	fmt.Println(fileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	Log = logrus.New()
	// 设置输出
	Log.Out = src

	// 设置日志级别
	Log.SetLevel(logrus.DebugLevel)
	// 设置 rotatelogs
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logwriter, err := rotatelogs.New(
		// 分割后的文件名称
		logFileName+".%y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logFileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writemap := lfshook.WriterMap{
		logrus.InfoLevel:  logwriter,
		logrus.FatalLevel: logwriter,
		logrus.DebugLevel: logwriter,
		logrus.WarnLevel:  logwriter,
		logrus.ErrorLevel: logwriter,
		logrus.PanicLevel: logwriter,
	}
	lfhook := lfshook.NewHook(writemap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 hook
	Log.AddHook(lfhook)
	return func(c *gin.Context) {
		// 开始时间
		starttime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endtime := time.Now()
		// 执行时间
		latencytime := endtime.Sub(starttime)
		// 请求方式
		reqmethod := c.Request.Method
		// 请求路由
		requri := c.Request.RequestURI
		// 状态码
		statuscode := c.Writer.Status()
		// 请求ip
		clientip := c.ClientIP()
		// 日志格式
		Log.WithFields(logrus.Fields{
			"status_code ": statuscode,
			"latency_time": latencytime,
			"client_ip":    clientip,
			"req_method":   reqmethod,
			"req_uri":      requri,
		}).Info()
	}

}
