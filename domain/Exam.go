package domain

import "wire_test/domain/service"

type Exam struct {
	userService service.UserService
}

func (e *Exam) Edit() int {
	return e.userService.LoadUserId(1)
}

func NewExam(user service.UserService) *Exam {
	return &Exam{userService: user}
}
