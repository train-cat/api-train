package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
	"golang.org/x/crypto/bcrypt"
)

const costBcrypt = 13

type (
	// UserData regroup all values can be received by client
	UserData struct {
		Username *string `gorm:"column:username" json:"username" validate:"required,unique=username,min=2,max=255"`
		Email    *string `gorm:"column:email"    json:"email"    validate:"required,unique=email,email,max=255"`
		Password *string `gorm:"-"               json:"password,omitempty" validate:"required,min=6,max=255"`
	}

	// User of the API
	User struct {
		Entity
		UserData
		EncodedPassword string      `gorm:"column:password" json:"-"`
		Roles           SliceString `gorm:"column:roles"    json:"roles"`
		rest.Hateoas
	}

	// UserToken used for retrieve new JsonWebToken
	UserToken struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// UserInput send by client
	UserInput UserData
)

// TableName for UserInput, used for 'unique' validation
func (i *UserInput) TableName() string {
	return "user"
}

// ToEntity transform UserInput to User
func (i *UserInput) ToEntity() (*User, error) {
	u := &User{UserData: UserData(*i)}

	p, err := bcrypt.GenerateFromPassword([]byte(*u.Password), costBcrypt)
	u.Password = nil

	if err != nil {
		return nil, err
	}

	u.EncodedPassword = string(p)

	return u, nil
}

// GenerateHateoas content
func (u *User) GenerateHateoas(ctx *aah.Context) error {
	return nil
}
