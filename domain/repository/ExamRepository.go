package repository

import "wire_test/domain"

type ExamRepository interface {
	Load(id int) *domain.Exam
}

type ExamRepositoryBuilder func() (ExamRepository, error)

var examRepositoryBuilder ExamRepositoryBuilder

func NewExamRepository() (ExamRepository, error) {
	return examRepositoryBuilder()
}

func SetUserServiceBuilder(b ExamRepositoryBuilder) {
	examRepositoryBuilder = b
}
