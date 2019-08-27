package configuration

import (
	"wire_test/domain/repository"
	repositoryImpl "wire_test/gateway/repository"
)

var ExamRepositoryBuilder = func() (repository.ExamRepository, error) {
	return repositoryImpl.NewExamRepository(), nil
}

func RepositoryInit() {
	repository.SetUserServiceBuilder(ExamRepositoryBuilder)
}
