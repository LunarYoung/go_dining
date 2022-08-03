package pkg

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"user/model"
)

var Db *gorm.DB

func Mysql() {
	dsn := "root:123456@tcp(106.13.192.119:3306)/ding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	// 自动迁移
	err1 := db.AutoMigrate(
		&model.AppUser{},
		&model.OrgUser{},
		&model.Org{},
	)
	if err != nil {
		panic(err1)
		return
	}
	Db = db
}
