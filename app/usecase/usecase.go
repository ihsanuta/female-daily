package usecase

import (
	"female-daily/app/repository"
	"female-daily/app/usecase/user"
)

type Usecase struct {
	User user.UserUsecase
}

func Init(repository *repository.Repository) *Usecase {
	uc := &Usecase{
		User: user.NewUserUsecase(
			repository.User,
		),
	}
	return uc
}
