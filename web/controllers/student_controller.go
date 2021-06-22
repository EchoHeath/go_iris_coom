package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"go-electricity/repositories"
	"go-electricity/services"
)

type StudentController struct {

}

func (s *StudentController) Get() mvc.View {
	studentRepository := repositories.NewStudentManage()
	studentManager := services.NewStudentServiceManage(studentRepository)
	studentResult := studentManager.ShowStudentName()
	return mvc.View{
		Name: "student/index.html",
		Data: studentResult,
	}
}
