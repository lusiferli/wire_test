package configuration

import (
	"wire_test/domain/service"
	serviceImpl "wire_test/gateway/service"
)

var UserServiceBuilder = func() (service.UserService, error) {
	return serviceImpl.NewUserServiceImpl(), nil
}

func ServiceInit() {
	service.SetUserServiceBuilder(UserServiceBuilder)
}
