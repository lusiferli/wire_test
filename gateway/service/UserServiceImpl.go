package service

import "wire_test/domain/service"

type UserServiceImpl struct {
}

func (u *UserServiceImpl) LoadUserId(id int) int {
	return 1001
}

func NewUserServiceImpl() service.UserService {
	return &UserServiceImpl{}
}
