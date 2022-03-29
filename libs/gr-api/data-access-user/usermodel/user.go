package usermodel

import (
	"errors"

	common "grailed/libs/gr-api/shared-common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	FirstName       string `json:"firstName" gorm:"column:first_name;"`
	LastName        string `json:"lastName" gorm:"column:last_name;"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	FirstName       string `json:"firstName" gorm:"column:first_name;"`
	LastName        string `json:"lastName" gorm:"column:last_name;"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (User) TableName() string {
	return "users"
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)
