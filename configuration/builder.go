package configuration

import (
	"wire_test/domain/repository"
	"wire_test/domain/service"
	repositoryImpl "wire_test/gateway/repository"
	serviceImpl "wire_test/gateway/service"
)

var UserServiceBuilder = func() (service.UserService, error) {
	return serviceImpl.NewUserServiceImpl(), nil
}

var ExamRepositoryBuilder = func() (repository.ExamRepository, error) {
	return repositoryImpl.NewExamRepository(), nil
}

func Init() {
	service.SetUserServiceBuilder(UserServiceBuilder)
	repository.SetUserServiceBuilder(ExamRepositoryBuilder)
}
