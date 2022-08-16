package middleware

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"user/model"
)

var Db *gorm.DB

func Mysql() {
	dsn := RemoteViper.GetString("mysql")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	//defer Db.Close()

	// 自动迁移
	err1 := db.AutoMigrate(
		&model.AppUser{},
		&model.OrgUser{},
		&model.Org{},
	)
	if err1 != nil {
		panic(err1)
		return
	}
	Db = db
}
