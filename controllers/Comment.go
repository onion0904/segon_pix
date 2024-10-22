package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"PixApp/models"
	"PixApp/repositories"
)



func (con *Controller) AddComment(c echo.Context) error {
	newcomment := c.QueryParam("comment")
	userID := c.QueryParam("userID")
	comment := models.Comment{}
	comment.Message = newcomment
	uintUserID64, err := strconv.ParseUint(userID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintUserID := uint(uintUserID64)
	comment.UserID = uintUserID
	if err := c.Bind(&comment); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
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
	if err := repo.AddComment(&comment,uintImageID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusOK)
}


func (con *Controller) UpdateComment(c echo.Context) error {
	commentID := c.QueryParam("commentID")
	uintCommentID64, err := strconv.ParseUint(commentID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintCommentID := uint(uintCommentID64)
	imageID := c.QueryParam("imageID")
	uintImageID64, err := strconv.ParseUint(imageID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintImageID := uint(uintImageID64)
	newContent := c.QueryParam("newContent")
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	if err := repo.UpdateComment(uintCommentID,newContent,uintImageID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusOK)
}


func (con *Controller) DeleteComment(c echo.Context) error {
	commentID := c.QueryParam("commentID")
	uintCommentID64, err := strconv.ParseUint(commentID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintCommentID := uint(uintCommentID64)	
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	if err := repo.DeleteComment(uintCommentID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusOK)
}