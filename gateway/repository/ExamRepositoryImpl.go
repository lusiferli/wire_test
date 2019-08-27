package repository

import (
	"wire_test/domain"
	service2 "wire_test/domain/service"
)

type ExamRepository struct {
	userService service2.UserService
}

func (u *ExamRepository) Load(id int) *domain.Exam {
	return domain.NewExam()
}

func NewExamRepository() *ExamRepository {
	return &ExamRepository{}
}
