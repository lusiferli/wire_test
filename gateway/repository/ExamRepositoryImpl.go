package repository

import (
	"wire_test/domain"
	service2 "wire_test/domain/service"
)

type ExamRepository struct {
	userService service2.UserService
}

func (u *ExamRepository) Load(id int) *domain.Exam {
	return domain.NewExam(u.userService)
}

func NewExamRepository(userService service2.UserService) *ExamRepository {
	return &ExamRepository{userService: userService}
}
