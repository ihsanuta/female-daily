package user

import (
	"encoding/json"
	"female-daily/app/model"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	FetchUsers(params model.FetchUserParam) ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	GetList(param model.UserParam) ([]*model.User, error)
	GetLastID() (int, error)
	Create(user *model.User) error
	Update(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) FetchUsers(params model.FetchUserParam) ([]*model.User, error) {
	var users []*model.User
	client := http.Client{
		Timeout: time.Minute * 15,
	}

	baseURL, _ := url.Parse("https://reqres.in/api/users")
	p := url.Values{}

	if params.Page > 0 {
		p.Add("page", fmt.Sprintf("%d", params.Page))
	}

	baseURL.RawQuery = p.Encode()
	req, err := http.NewRequest(http.MethodGet, baseURL.String(), nil)
	if err != nil {
		return users, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return users, err
	}

	var response model.ResponseFetch
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return users, err
	}

	users = response.Data
	return users, nil
}

func (u *userRepo) GetByID(id int) (*model.User, error) {
	user := &model.User{
		ID:        id,
		DeletedAt: nil,
	}
	db := u.db.Find(&user)
	if db.Error != nil {
		return nil, db.Error
	}

	return user, nil
}

func (u *userRepo) GetList(param model.UserParam) ([]*model.User, error) {
	var list []*model.User

	q := ""
	p := []interface{}{}
	if param.Email != "" {
		if q == "" {
			q = fmt.Sprint("WHERE email = ?")
		} else {
			q = fmt.Sprintf("%s AND email = ?", q)
		}
		p = append(p, param.Email)
	}

	if param.FirstName != "" {
		if q == "" {
			q = fmt.Sprint("WHERE first_name = ?")
		} else {
			q = fmt.Sprintf("%s AND first_name = ?", q)
		}
		p = append(p, param.FirstName)
	}

	if param.LastName != "" {
		if q == "" {
			q = fmt.Sprint("WHERE last_name = ?")
		} else {
			q = fmt.Sprintf("%s AND last_name = ?", q)
		}
		p = append(p, param.LastName)
	}

	// deltedat
	if q == "" {
		q = fmt.Sprint("WHERE deleted_at = ?")
	} else {
		q = fmt.Sprintf("%s AND deleted_at = ?", q)
	}
	p = append(p, nil)

	limit := 10
	if param.Limit > 0 {
		limit = param.Limit
	}

	offset := 0
	if param.Page > 0 {
		offset = (param.Page - 1) * limit
	}

	query := fmt.Sprintf("SELECT * FROM users %s LIMIT %d OFFSET %d", q, limit, offset)
	db := u.db.Raw(query, p...).Scan(&list)
	if db.Error != nil {
		return list, db.Error
	}

	return list, nil
}

func (u *userRepo) Create(user *model.User) error {
	db := u.db.Create(&user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (u *userRepo) GetLastID() (int, error) {
	user := &model.User{}
	db := u.db.Find(&user).Order("id DESC")
	if db.Error != nil {
		return 0, db.Error
	}

	return user.ID, nil
}

func (u *userRepo) Update(user *model.User) error {
	db := u.db.Save(&user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (u *userRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{
		Email:     email,
		DeletedAt: nil,
	}
	db := u.db.Find(&user)
	if db.Error != nil {
		return nil, db.Error
	}

	return user, nil
}
