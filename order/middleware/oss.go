package middleware

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Bucket *oss.Bucket

func NewOss() {

	client, err := oss.New(RemoteViper.GetString("oss.endpoint"), RemoteViper.GetString("oss.accessKeyId"), RemoteViper.GetString("oss.accessKeySecret"))
	if err != nil {

	}
	Bucket, err = client.Bucket(RemoteViper.GetString("oss.bucketName"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Upload(fileName string, fileByte []byte) (url string, err error) {

	err = Bucket.PutObject(fileName, bytes.NewReader([]byte(fileByte)))
	if err != nil {
		return url, err
	}
	return RemoteViper.GetString("oss.pre") + fileName, nil
}
