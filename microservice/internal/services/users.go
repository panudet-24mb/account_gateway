package services

import "account_gateway/internal/repository"

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{userRepository: userRepository}
}

func (u userService) NewUserAccount(n *NewUserAccountRequest) (*NewUserAccountResponse, error) {
	user := repository.User{
		UserName: n.UserName,
		//เข้ารหัสด้วย bcrypt
		HashPassword: n.HashPassword,
		Email:        n.Email,
		AcceptTerms:  n.AcceptTerms,
	}
	_, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &NewUserAccountResponse{
		UserName: n.UserName,
	}, nil
}
