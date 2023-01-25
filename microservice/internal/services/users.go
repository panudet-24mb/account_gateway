package services

import (
	"account_services/internal/repository"
	"log"
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

// func (s userService) NewUserAccount(n *NewUserAccountRequest) (*NewUserAccountResponse, error) {
// 	user := repository.User{
// 		UserName: n.UserName,
// 		//เข้ารหัสด้วย bcrypt
// 		HashPassword: n.HashPassword,
// 		Email:        n.Email,
// 		AcceptTerms:  n.AcceptTerms,
// 	}
// 	_, err := s.userRepository.Create(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &NewUserAccountResponse{
// 		UserName: n.UserName,
// 	}, nil
// }

func (s userService) NewUserAccount(n *NewUserAccountRequest) (*NewUserAccountResponse, error) {

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
	}

	data, err := s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	userResponse := &NewUserAccountResponse{
		UserName:     data.UserName,
		HashPassword: data.HashPassword,
		Email:        data.Email,
		AcceptTerms:  data.AcceptTerms,
		CreateAt:     data.CreateAt,
	}

	return userResponse, nil
}

func (s userService) GetAccounts() ([]FindAccountReponse, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	userResponses := []FindAccountReponse{}
	for _, user := range users {
		userResponse := FindAccountReponse{
			UserName: user.UserName,
			Email:    user.Email,
			CreateAt: user.CreateAt,
			UpdateAt: user.UpdateAt,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (s userService) GetAccountByID(id string) (*FindAccountReponse, error) {
	user, err := s.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	userResponse := FindAccountReponse{
		UserName: user.UserName,
		Email:    user.Email,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
	return &userResponse, nil

}

func (s userService) UpdateAccounts(updateUser *UpdateUserRequest) (*UpdateUserResponse, error) {
	userData := repository.User{
		UserName: updateUser.UserName,
		Email:    updateUser.Email,
		UpdateAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.userRepository.Update(userData)
	if err != nil {
		return nil, err
	}

	userUpdate := &UpdateUserResponse{
		UserName: dataUpdate.UserName,
		Email:    dataUpdate.Email,
		UpdateAt: dataUpdate.UpdateAt,
	}
	return userUpdate, nil
}

func (s userService) UpdateAccount(id string, updateUser *UpdateUserRequest) (*UpdateUserResponse, error) {
	userData := repository.User{
		UserName: updateUser.UserName,
		Email:    updateUser.Email,
		UpdateAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.userRepository.UpdateOne(id, userData)
	if err != nil {
		log.Fatal(err)
	}
	userUpdate := &UpdateUserResponse{
		UserName: dataUpdate.UserName,
		Email:    dataUpdate.Email,
		UpdateAt: dataUpdate.UpdateAt,
	}
	return userUpdate, nil
}

func (s userService) DeleteAccount(username string, deleteUser *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	userData := repository.User{
		UserName: deleteUser.UserName,
		DeleteAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataDelete, err := s.userRepository.DeleteOne(username, userData)
	if err != nil {
		log.Fatal(err)
	}
	userDelete := &DeleteAccountResponse{
		UserName: deleteUser.UserName,
		DeleteAt: dataDelete.DeleteAt,
	}

	return userDelete, nil
}

func (s userService) FindAccount(findAccount *FindAccountRequest) (*FindAccountReponse, error) {
	userData := repository.User{
		UserName: findAccount.UserName,
		Email:    findAccount.Email,
	}
	user, err := s.userRepository.FindUser(userData)
	if err != nil {
		return nil, err
	}
	userResponse := &FindAccountReponse{
		UserName: user.UserName,
		Email:    user.Email,
	}

	return userResponse, nil
}
