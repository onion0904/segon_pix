package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"PixApp/repositories"
)


func (con *Controller) AddLike(c echo.Context) error {
	userID := c.QueryParam("userID")
	uintUserID64, err := strconv.ParseUint(userID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintUserID := uint(uintUserID64)
	imageID := c.QueryParam("imageID")
	uintImageID64, err := strconv.ParseUint(imageID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintImageID := uint(uintImageID64)
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	if err := repo.AddLike(uintUserID,uintImageID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusNoContent) // 204エラー
}


func (con *Controller) RemoveLike(c echo.Context) error {
	userID := c.QueryParam("userID")
	uintUserID64, err := strconv.ParseUint(userID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintUserID := uint(uintUserID64)
	imageID := c.QueryParam("imageID")
	uintImageID64, err := strconv.ParseUint(imageID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintImageID := uint(uintImageID64)
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	if err := repo.RemoveLike(uintUserID,uintImageID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusNoContent) // 204エラー
}