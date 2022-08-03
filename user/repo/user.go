package repo

import (
	_ "gorm.io/gorm"
	"user/model"
	"user/pkg"
)

type ICategoryRepository interface {
	Create(req model.Org)
}

type CategoryRepository struct {
}

func (c CategoryRepository) Create(req model.Org) {

	pkg.Db.Create(req)
}
