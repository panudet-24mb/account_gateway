package services

import (
	"account_gateway/internal/errs"
	"account_gateway/internal/repository"
	"account_gateway/internal/utils"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return authService{userRepo: userRepo}
}

func (a authService) DefaultLogin(loginForm *LoginDefaultForm) (*ResponseLogin, error) {

	var (
		EmailExist    bool
		UserNameExist bool
	)

	if len(loginForm.UserName) > 0 && len(loginForm.Email) == 0 {
		if len(loginForm.UserName) < 5 && len(loginForm.UserName) >= 20 {
			return nil, errs.CustomError(
				"UserName must be between 5 and 20 characters",
				400,
			)

		}
		UserNameExist = true
	}
	if len(loginForm.Email) > 0 && len(loginForm.UserName) == 0 {
		email := utils.VerifyEmail(loginForm.Email)
		if !email {
			return nil, errs.CustomError(
				"Email is not valid",
				400,
			)
		}
		EmailExist = true

	}
	sevenOrMore, number, upper, special := utils.VerifyPassword(loginForm.Password)
	if !sevenOrMore || !number || !upper || !special {
		return nil, errs.CustomError(
			"Password must be at least 7 characters long, contain at least one number, one upper case letter, and one special character",
			400,
		)
	}
	//create empty struct

	userInfomation := []repository.User{}
	switch {
	case UserNameExist:
		user, err := a.userRepo.FindUser(repository.User{
			UserName: loginForm.UserName,
		})
		if err == errors.New("record not found") {
			return nil, errs.CustomError(
				"Not Found This UserName",
				400,
			)
		}
		userInfomation = append(userInfomation, *user)

	case EmailExist:
		user, err := a.userRepo.FindUser(repository.User{
			Email: loginForm.Email,
		})
		if err == errors.New("record not found") {
			return nil, errs.CustomError(
				"Not Found This Email",
				400,
			)
		}
		userInfomation = append(userInfomation, *user)

	default:
		return nil, errs.CustomError(
			"Only UserName Or Email must be filled Once",
			400,
		)
	}

	err := bcrypt.CompareHashAndPassword([]byte(userInfomation[0].HashPassword), []byte(loginForm.Password))
	if err != nil {
		return nil, errs.CustomError(
			"can't update user",
			500,
		)

	}
	cliams := jwt.StandardClaims{
		Issuer:    fmt.Sprint(userInfomation[0].UserName),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(viper.GetString("app.secret")))
	if err != nil {
		return nil, err
	}
	if len(token) > 0 {
		a.userRepo.UpdateOne(
			repository.User{
				UserName:     userInfomation[0].UserName,
				LoginAttempt: 1,
				LastLogin:    strconv.Itoa(int(time.Now().Unix())),
			},
		)

	}

	return &ResponseLogin{
		Message: token,
		Status:  true,
	}, nil
}
