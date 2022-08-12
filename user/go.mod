module user

go 1.16

require (
	github.com/Shopify/sarama v1.35.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.8.1
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.9
	github.com/jinzhu/gorm v1.9.16
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/nacos/v2 v2.0.0-20201025091542-fa097e59f8ac
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.12.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	golang.org/x/net v0.0.0-20220708220712-1185a9018129
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
