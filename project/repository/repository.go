package repository

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/lazytooo/choc-cream/project/models"
	"github.com/sirupsen/logrus"
)

type implRepo struct {
	dbClient    *sqlx.DB
	redisClient *redis.Client
}

func NewRepository(dbClient *sqlx.DB, redisClient *redis.Client) Repository {
	return &implRepo{
		dbClient:    dbClient,
		redisClient: redisClient,
	}
}

func (r *implRepo) QueryPhotoList(start, end int64) (list []models.Photo, err error) {
	query := "SELECT `photo_id`, `photo_title`, `photo_url`, `photo_shot_date`, `describe`, `create_time`, " +
		"`update_time` FROM `photo` LIMIT ?, ?;"

	err = r.dbClient.Select(&list, query, start, end)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		logrus.WithError(err).WithField("start", start).WithField("end", end).
			Warnln("query photo list error")
		return
	}

	return
}
