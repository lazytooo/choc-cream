package handler

import (
	"github.com/labstack/echo"
	"github.com/lazytooo/choc-cream/project/models"
	"github.com/lazytooo/choc-cream/project/usecase"
	"net/http"
	"strconv"
)

type Handler struct {
	Usecase usecase.Usecase
}

func (h Handler) GetPhotoList(c echo.Context) (err error) {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.QueryParam("page_size"), 10, 64)

	list, err := h.Usecase.GetPhotoList(models.PageInfo{
		Page:     page,
		PageSize: pageSize,
		UsePage:  true,
	})

	result := models.GetErrorMap(err)

	if err == nil {
		if len(list) == 0 {
			list = []models.Photo{}
		}
		result["data"] = list
	}
	return c.JSON(http.StatusOK, result)
}
