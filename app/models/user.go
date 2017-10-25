package models

import (
	"github.com/train-cat/api-train/app/rest"
	"golang.org/x/crypto/bcrypt"
	"aahframework.org/aah.v0"
)

const costBcrypt = 13

type (
	UserData struct {
		Username *string `gorm:"column:username" json:"username" validate:"required,unique=username,min=2,max=255"`
		Email    *string `gorm:"column:email"    json:"email"    validate:"required,unique=email,email,max=255"`
		Password *string `gorm:"-"               json:"password,omitempty" validate:"required,min=6,max=255"`
	}

	User struct {
		Entity
		UserData
		EncodedPassword string `gorm:"column:password" json:"-"`
		rest.Hateoas
	}

	UserToken struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserInput UserData
)

// TableName return the table for gorm, used for 'unique' validation
func (i *UserInput) TableName() string {
	return "user"
}

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

func (u *User) GenerateHateoas(ctx *aah.Context) error {
	u.Links = rest.Links{
		"self": rest.Link{
		//	Href: rest.GenerateURI(ctx, "get_user", u.ID),
		},
	}

	return nil
}
