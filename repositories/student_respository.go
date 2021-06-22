package repositories

import (
	"go-electricity/datamodels"
)

type StudentRepository interface {
	GetStudentName() string
}

type StudentManage struct {
}

func NewStudentManage() StudentRepository {
	return &StudentManage{}
}

func (s StudentManage) GetStudentName() string {
	student := &datamodels.Student{
		Id:   1,
		Name: "aa",
		Age:  15,
	}
	return student.Name
}
