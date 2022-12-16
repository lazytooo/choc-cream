package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lazytooo/choc-cream/project/config"
	"github.com/lazytooo/choc-cream/project/conn"
	"github.com/lazytooo/choc-cream/project/repository"
	"github.com/lazytooo/choc-cream/project/router"
	"github.com/lazytooo/choc-cream/project/usecase"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	initLog()
	db := conn.GetDBConn()
	redisClient := conn.GetRedisConn()
	repo := repository.NewRepository(db, redisClient)
	ucase := usecase.NewUsecase(repo)

	runHttpServer(ucase)
}

func runHttpServer(ucase usecase.Usecase) {
	echoServer := echo.New()
	echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.CORS())

	router.RegisterHTTPHandlers(echoServer, ucase)

	err := echoServer.Start(config.HTTPAddr)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func initLog() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	logrus.Info("logrus init success")
}
