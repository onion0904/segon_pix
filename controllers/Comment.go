package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"PixApp/models"
	"PixApp/repositories"
	"log"
)


func (con *Controller) AddComment(c echo.Context) error {
    newComment := c.QueryParam("comment")
    userID := c.QueryParam("userID")
    if newComment == "" || userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "コメントとユーザーIDが必要です"})
    }

    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    comment := models.Comment{
        Message: newComment,
        UserID:  uintID,
    }

    imageID := c.QueryParam("imageID")
    if imageID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "画像IDが必要です"})
    }

    uintImageID64, err := strconv.ParseUint(imageID, 10, 64)
    if err != nil {
        log.Printf("Invalid imageID: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "無効な画像IDです"})
    }
    uintImageID := uint(uintImageID64)

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error initializing repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    if err := repo.AddComment(&comment, uintImageID); err != nil {
        log.Printf("Error adding comment: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "コメントの追加に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "コメントが追加されました"})
}

func (con *Controller) UpdateComment(c echo.Context) error {
	userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    commentID := c.QueryParam("commentID")
    if commentID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "コメントIDが必要です"})
    }

    uintCommentID64, err := strconv.ParseUint(commentID, 10, 64)
    if err != nil {
        log.Printf("Invalid commentID: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "無効なコメントIDです"})
    }
    uintCommentID := uint(uintCommentID64)

    imageID := c.QueryParam("imageID")
    if imageID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "画像IDが必要です"})
    }

    uintImageID64, err := strconv.ParseUint(imageID, 10, 64)
    if err != nil {
        log.Printf("Invalid imageID: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "無効な画像IDです"})
    }
    uintImageID := uint(uintImageID64)

    newContent := c.QueryParam("newContent")
    if newContent == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "新しいコメント内容が必要です"})
    }

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error initializing repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    if err := repo.UpdateComment(uintCommentID, newContent, uintImageID); err != nil {
        log.Printf("Error updating comment: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "コメントの更新に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "コメントが更新されました"})
}

func (con *Controller) DeleteComment(c echo.Context) error {
	userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    commentID := c.QueryParam("commentID")
    if commentID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "コメントIDが必要です"})
    }

    uintCommentID64, err := strconv.ParseUint(commentID, 10, 64)
    if err != nil {
        log.Printf("Invalid commentID: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "無効なコメントIDです"})
    }
    uintCommentID := uint(uintCommentID64)

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error initializing repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    if err := repo.DeleteComment(uintCommentID); err != nil {
        log.Printf("Error deleting comment: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "コメントの削除に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "コメントが削除されました"})
}
