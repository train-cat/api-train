package controllers

import (
	"time"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
	"github.com/train-cat/api-train/app/security"
)

// UserController regroup all endpoints concern the user
type UserController struct {
	Controller
}

// Post create new user
func (c *UserController) Post(i *models.UserInput) {
	if !c.validatePost(i) {
		return
	}

	u, err := repositories.User.Persist(i)

	if c.serverError(err) || c.serverError(c.hateoas(u)) {
		return
	}

	c.Reply().Created().JSON(u)
}

// Token authenticate one user
func (c *UserController) Token(t *models.UserToken) {
	// NOTE: Validation feature is upcoming :)
	if ess.IsStrEmpty(t.Username) || ess.IsStrEmpty(t.Password) {
		c.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	// get the user details by username
	user, err := repositories.User.FindByUsername(t.Username)
	if user == nil || err != nil {
		c.Reply().Unauthorized().JSON(aah.Data{
			"message": "invalid credentials",
		})
		return
	}

	// validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.EncodedPassword), []byte(t.Password)); err != nil {
		c.Reply().Unauthorized().JSON(aah.Data{
			"message": "invalid credentials",
		})
		return
	}

	// Generate JWT token
	token := security.CreateJWTToken()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = *user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString(security.JWTSigningKey)
	if c.serverError(err) {
		return
	}

	// everything went good, respond token
	c.Reply().Ok().JSON(aah.Data{
		"token": signedToken,
	})
}
