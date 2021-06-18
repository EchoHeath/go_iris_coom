package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"go-electricity/repositories"
	"go-electricity/services"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	MovieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: MovieResult,
	}
}
