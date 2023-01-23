package services

import (
	"account_gateway/internal/repository"
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
	}

	return userResponse, nil
}

func (u userService) GetAccounts() ([]FindAccountReponse, error) {
	users, err := u.userRepository.GetAll()
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

func (u userService) GetAccountByID(id string) (*FindAccountReponse, error) {
	user, err := u.userRepository.GetByID(id)
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

func (u userService) UpdateAccounts(updateUser *UpdateUserRequest) (*UpdateUserResponse, error) {
	userData := repository.User{
		UserName: updateUser.UserName,
		Email:    updateUser.Email,
		UpdateAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := u.userRepository.Update(userData)
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

func (u userService) UpdateAccount(updateUser *UpdateUserRequest) (*UpdateUserResponse, error) {
	userData := repository.User{
		UserName: updateUser.UserName,
		Email:    updateUser.Email,
		UpdateAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := u.userRepository.UpdateOne(userData)
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

func (u userService) DeleteAccount(username string, deleteUser *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	userData := repository.User{
		UserName: deleteUser.UserName,
		DeleteAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataDelete, err := u.userRepository.DeleteOne(username, userData)
	if err != nil {
		log.Fatal(err)
	}
	userDelete := &DeleteAccountResponse{
		UserName: deleteUser.UserName,
		DeleteAt: dataDelete.DeleteAt,
	}
	return userDelete, nil
}

func (u userService) FindAccount(findAccount FindAccountRequest) (*FindAccountReponse, error) {
	userData := repository.User{
		UserName: findAccount.UserName,
		Email:    findAccount.Email,
	}
	user, err := u.userRepository.FindUser(userData)
	if err != nil {
		return nil, err
	}
	userResponse := &FindAccountReponse{
		UserName: user.UserName,
		Email:    user.Email,
	}

	return userResponse, nil
}
