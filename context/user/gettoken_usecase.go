package user

import (
	"errors"
	"time"

	"github.com/heroku/go-getting-started/app"
	"github.com/heroku/go-getting-started/domain/model"
	"github.com/heroku/go-getting-started/domain/repository"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mmuflih/datelib"
	"golang.org/x/crypto/bcrypt"
)

type GetTokenUsecase interface {
	Issue(GetTokenRequest) (error, GetTokenResponse)
	claimToken(*model.User) model.AccessToken
}

type GetTokenRequest interface {
	GetUsername() string
	GetPassword() string
}

type GetTokenResponse interface{}

type getTokenUsecase struct {
	userRepo     repository.UserRepository
	signatureKey []byte
}

func NewGetTokenUsecase(userRepo repository.UserRepository, signatureKey []byte) GetTokenUsecase {
	return &getTokenUsecase{userRepo, signatureKey}
}

func (uc *getTokenUsecase) Issue(req GetTokenRequest) (error, GetTokenResponse) {
	err, user := uc.CheckUser(req.GetUsername())
	if err != nil {
		return err, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))
	if err != nil {
		return app.NewError("Invalid username and password"), nil
	}
	at := uc.claimToken(user)
	at.Data = nil
	at.Role = user.Role
	user.LastLogin.Time = datelib.NewTZLocal()
	user.LastLogin.Valid = true
	return nil, at
}

func (uc *getTokenUsecase) CheckUser(username string) (error, *model.User) {
	err, user := uc.userRepo.FindByUsername(username)
	if user.ID == "" {
		return errors.New("Username and password not match" + err.Error()), nil
	}
	return nil, user
}

func (uc *getTokenUsecase) claimToken(u *model.User) model.AccessToken {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS512)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	expiredAt := time.Now().Add(time.Hour * (24 * 1500)).Unix()
	claims["user_id"] = u.ID
	claims["role"] = u.Role
	claims["exp"] = expiredAt

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(uc.signatureKey)

	return model.AccessToken{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expiredAt,
	}
}
