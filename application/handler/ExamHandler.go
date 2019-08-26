package handler

import (
	"github.com/google/wire"
	"wire_test/domain/repository"
	repository2 "wire_test/gateway/repository"
	"wire_test/gateway/service"
)

func EditExam() int {
	examRepository, _ := InitExamRepository()
	exam := examRepository.Load(1)
	return exam.Edit()
}

func InitExamRepository() (repository.ExamRepository, error) {
	// wire的构造
	wire.Build(repository2.NewExamRepository, service.NewUserServiceImpl)

	//正常的构造
	return repository2.NewExamRepository(service.NewUserServiceImpl()), nil
}
