package repository

import "github.com/lazytooo/choc-cream/project/models"

type Repository interface {
	QueryPhotoList(start, end int64) (list []models.PhotoInfo, err error)
}
