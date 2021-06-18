package services

import "go-electricity/repositories"

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	return "获取到的信息是: " + m.repo.GetMovieName()
}