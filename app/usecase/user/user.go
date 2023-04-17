package user

import (
	"female-daily/app/model"
	"female-daily/app/repository/user"
	"fmt"
	"log"
	"sync"
	"time"
)

type UserUsecase interface {
	FetchUser(param model.FetchUserParam) ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	GetList(param model.UserParam) ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
}

type userUsecase struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) FetchUser(param model.FetchUserParam) ([]*model.User, error) {
	data, err := u.userRepo.FetchUsers(param)
	if err != nil {
		return data, err
	}

	wg := sync.WaitGroup{}
	ch := make(chan error, len(data))
	for _, v := range data {
		wg.Add(1)
		go func(user *model.User) {
			defer wg.Done()
			// check id user exists
			check, err := u.userRepo.GetByID(user.ID)
			if err != nil && err.Error() != "record not found" {
				ch <- err
				return
			}

			if check != nil {
				ch <- fmt.Errorf("%s", "user id exists")
				return
			}

			err = u.userRepo.Create(user)
			if err != nil {
				ch <- err
				return
			}
		}(v)
	}
	wg.Wait()
	close(ch)

	if len(ch) > 0 {
		for err := range ch {
			log.Println(err.Error())
		}
	}

	return data, nil
}

func (u *userUsecase) GetByID(id int) (*model.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUsecase) GetList(param model.UserParam) ([]*model.User, error) {
	return u.userRepo.GetList(param)
}

func (u *userUsecase) Create(user *model.User) error {
	id, err := u.userRepo.GetLastID()
	if err != nil {
		return err
	}

	user.ID = id + 1
	err = u.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Update(user *model.User) error {
	_, err := u.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}

	check, err := u.userRepo.GetByEmail(user.Email)
	if err != nil {
		return err
	}

	if check.ID != user.ID {
		return fmt.Errorf("error email exists other user : %s", check.Email)
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Delete(id int) error {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	deletedAt := time.Now()
	user.DeletedAt = &deletedAt

	err = u.userRepo.Update(user)
	if err != nil {
		return err
	}

	return nil
}
