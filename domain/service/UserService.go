package service

type UserService interface {
	LoadUserId(id int) int
}

type UserServiceBuilder func() (UserService, error)

var userServiceBuilder UserServiceBuilder

func NewUserService() (UserService, error) {
	return userServiceBuilder()
}

func SetUserServiceBuilder(b UserServiceBuilder) {
	userServiceBuilder = b
}
