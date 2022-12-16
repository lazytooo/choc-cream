package usecase

import (
	"github.com/lazytooo/choc-cream/project/models"
	"github.com/lazytooo/choc-cream/project/repository"
)

func NewUsecase(repo repository.Repository) Usecase {
	return &implUsecase{
		repo: repo,
	}
}

type implUsecase struct {
	repo repository.Repository
}

func (u *implUsecase) GetPhotoList(params models.PageInfo) (list []models.PhotoInfo, err error) {
	start, end := params.Limit()
	list, err = u.repo.QueryPhotoList(start, end)
	if err != nil {
		err = models.ErrDB
		return
	}
	return
}
