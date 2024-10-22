package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"PixApp/repositories"
)


func (con *Controller) AddFollow(c echo.Context) error {
	followingID := c.QueryParam("followingID")
	uintFollowingID64, err := strconv.ParseUint(followingID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintFollowingID := uint(uintFollowingID64)

	followedID := c.QueryParam("followedID")
	uintFollowedID64, err := strconv.ParseUint(followedID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintFollowedID := uint(uintFollowedID64)
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }

	if err := repo.AddFollow(uintFollowingID,uintFollowedID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusOK)
}


func (con *Controller) RemoveFollow(c echo.Context) error {
	followingID := c.QueryParam("followingID")
	uintFollowingID64, err := strconv.ParseUint(followingID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintFollowingID := uint(uintFollowingID64)

	followedID := c.QueryParam("followedID")
	uintFollowedID64, err := strconv.ParseUint(followedID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintFollowedID := uint(uintFollowedID64)
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }

	if err := repo.RemoveFollow(uintFollowingID,uintFollowedID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusOK)
}