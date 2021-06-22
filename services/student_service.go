package services

import "go-electricity/repositories"

type StudentService interface {
	ShowStudentName() string
}

type StudentServiceManage struct {
	repo repositories.StudentRepository
}

func NewStudentServiceManage(repo repositories.StudentRepository) StudentService {
	return &StudentServiceManage{repo: repo}
}

func (s StudentServiceManage) ShowStudentName() string {
	return "获取学生姓名: " + s.repo.GetStudentName()
}
