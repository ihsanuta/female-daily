package repository

import (
	"female-daily/app/repository/user"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	User user.UserRepo
}

func Init(db *gorm.DB) *Repository {
	repo := &Repository{
		User: user.NewUserRepo(
			db,
		),
	}
	return repo
}
