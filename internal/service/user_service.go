package service

type userService struct{}

func NewUserService(UserRepository) *userService {
	return &userService{}
}
