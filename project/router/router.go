package router

import (
	"github.com/labstack/echo"
	"github.com/lazytooo/choc-cream/project/handler"
	"github.com/lazytooo/choc-cream/project/usecase"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func RegisterHTTPHandlers(server *echo.Echo, usecase usecase.Usecase) {
	handler := &handler.Handler{
		Usecase: usecase,
	}

	bg := server.Group("/photo-wall")
	// setLogLevel 设置日志记录等级
	bg.GET("/debug/loglevel/:level", func(c echo.Context) error {
		level, err := logrus.ParseLevel(c.Param("level"))
		logrus.WithError(err).WithField("setlevel", c.Param("level")).Infoln("set log level")

		if err == nil {
			logrus.SetLevel(level)
			return c.String(http.StatusOK, logrus.GetLevel().String())
		}

		return c.String(http.StatusOK, err.Error())
	})

	bg.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339Nano))
	})

	bg.GET("/photos", handler.GetPhotoList)
}
