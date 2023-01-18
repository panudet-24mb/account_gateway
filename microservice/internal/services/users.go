package services

import (
	"account_gateway/internal/repository"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{userRepository: userRepository}
}

// func (u userService) NewUserAccount(n *NewUserAccountRequest) (*NewUserAccountResponse, error) {
// 	user := repository.User{
// 		UserName: n.UserName,
// 		//เข้ารหัสด้วย bcrypt
// 		HashPassword: n.HashPassword,
// 		Email:        n.Email,
// 		AcceptTerms:  n.AcceptTerms,
// 	}
// 	_, err := u.userRepository.Create(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &NewUserAccountResponse{
// 		UserName: n.UserName,
// 	}, nil
// }

func (u userService) NewUserAccount(n *NewUserAccountRequest) (*NewUserAccountResponse, error) {

	userInputPassword := n.HashPassword
	password, err := bcrypt.GenerateFromPassword([]byte(userInputPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := repository.User{
		UserName:     n.UserName,
		HashPassword: string(password),
		Email:        n.Email,
		AcceptTerms:  n.AcceptTerms,
		CreateAt:     strconv.Itoa(int(time.Now().Unix())),
		UpdateAt:     strconv.Itoa(int(time.Now().Unix())),
	}

	data, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	userResponse := &NewUserAccountResponse{
		UserName:     data.UserName,
		HashPassword: data.HashPassword,
		Email:        data.Email,
		AcceptTerms:  data.AcceptTerms,
		CreateAt:     data.CreateAt,
		UpdateAt:     data.UpdateAt,
	}

	return userResponse, nil
}
