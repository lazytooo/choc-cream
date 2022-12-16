package conn

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/lazytooo/choc-cream/project/config"
	"github.com/sirupsen/logrus"
	"time"
)

func GetDBConn() *sqlx.DB {
	if config.DBMaxIdleConns <= 0 || config.DBMaxOpenConns <= 0 {
		logrus.Fatal("DB max connection config err")
	}
	dsnConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_general_ci&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	var err error
	db, err := sqlx.Open("mysql", dsnConfig)

	if err == nil {
		err = db.Ping()
	}

	if err != nil {
		logrus.Fatalf("db connect error: %#v", err)
	}
	// 最大连接数限制
	db.SetMaxOpenConns(config.DBMaxOpenConns)
	db.SetMaxIdleConns(config.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Minute * 10)

	logrus.Infoln("connect to db")

	return db
}

func GetRedisConn() *redis.Client {
	return nil
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
		DB:       0, // use default DB
		PoolSize: 300,
		// 最低维持连接数
		MinIdleConns: 3,
		PoolTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		logrus.WithError(err).WithField("pong", pong).Fatalln("redis error")
	}
	logrus.WithField("pong", pong).Infoln("redis ping")

	return client
}
