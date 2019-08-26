package repository

import "wire_test/domain"

type ExamRepository interface {
	Load(id int) *domain.Exam
}
