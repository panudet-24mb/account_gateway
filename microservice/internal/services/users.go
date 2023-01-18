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

func (u userService) FindAllAccount() ([]FindAccountReponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	userResponses := []FindAccountReponse{}
	for _, user := range users {
		userResponse := FindAccountReponse{
			UserName: user.UserName,
			Email:    user.Email,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (u userService) FindAccount(id string) (*FindAccountReponse, error) {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	userResponse := FindAccountReponse{
		UserName: user.UserName,
		Email:    user.Email,
	}
	return &userResponse, nil

}
