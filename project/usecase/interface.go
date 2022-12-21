package usecase

import "github.com/lazytooo/choc-cream/project/models"

type Usecase interface {
	GetPhotoList(params models.PageInfo) (list []models.Photo, err error)
}
