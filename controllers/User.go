package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"PixApp/models"
	"PixApp/repositories"
	"gorm.io/gorm"
)


func (con *Controller) AddUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.NoContent(500) // 503エラー
	}
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(501) // 503エラー
    }
	if err := repo.AddUser(&user); err != nil {
		return c.NoContent(502) // 503エラー
	}
	return c.NoContent(http.StatusNoContent) // 204エラー
}

func (con *Controller) SearchImage(c echo.Context) error {
	Qhashtag := c.QueryParam("Hashtag")
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	PostedImage, err := repo.SearchImage(Qhashtag)
	if err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.JSON(http.StatusOK, PostedImage)
}


func (con *Controller) UserInfo(c echo.Context) error {
	userID := c.QueryParam("ID")
	uintID64, err := strconv.ParseUint(userID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintID := uint(uintID64)
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	User, err := repo.UserInfo(uintID)
	if err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.JSON(http.StatusOK, User)
}

func (con *Controller) DeleteUser(c echo.Context) error {
	userID := c.QueryParam("ID")
	uintID64, err := strconv.ParseUint(userID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintID := uint(uintID64)

	var existinguser models.User
	if err := con.db.First(&existinguser, uintID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound) // 404エラー
		}
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}

	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	if err := repo.DeleteUser(uintID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusNoContent) // 204エラー
}