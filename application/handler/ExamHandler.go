package handler

import (
	"wire_test/domain/repository"
)

func EditExam() int {
	examRepository, _ := repository.NewExamRepository()
	exam := examRepository.Load(1)
	return exam.Edit()
}
