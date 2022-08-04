package pkg

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Bucket *oss.Bucket

func NewOss() {

	client, err := oss.New(Config.GetString("oss.endpoint"), Config.GetString("oss.accessKeyId"), Config.GetString("oss.accessKeySecret"))
	if err != nil {

	}
	Bucket, err = client.Bucket(Config.GetString("oss.bucketName"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Upload(fileName string, fileByte []byte) (url string, err error) {

	err = Bucket.PutObject(fileName, bytes.NewReader([]byte(fileByte)))
	if err != nil {
		return url, err
	}
	return Config.GetString("oss.pre") + fileName, nil
}
