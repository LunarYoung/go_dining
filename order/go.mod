module order

go 1.16

require (
	github.com/Shopify/sarama v1.35.0
	github.com/aliyun/aliyun-oss-go-sdk v2.2.4+incompatible
	github.com/gin-gonic/gin v1.8.1
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.9
	github.com/jinzhu/copier v0.3.5
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/nacos/v2 v2.0.0-20200823081325-d250a107782a
	github.com/olivere/elastic/v7 v7.0.32
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.12.0
	github.com/yoyofxteam/nacos-viper-remote v0.4.0

	go.mongodb.org/mongo-driver v1.10.1
	golang.org/x/net v0.0.0-20220811182439-13a9a731de15
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
