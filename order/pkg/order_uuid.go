package pkg

import (
	"math/rand"
	"time"
)

func Uuid() string {
	str := "123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 4; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
