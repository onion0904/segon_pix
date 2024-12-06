package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"PixApp/repositories"
	"log"
)


func (con *Controller) AddLike(c echo.Context) error {
    userID := c.QueryParam("userID")
    imageID := c.QueryParam("imageID")

    if userID == "" || imageID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ユーザーIDと画像IDが必要です"})
    }

    uintuserID := uintID(userID)
    if uintuserID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "fail to convert userID to uintID"})
    }
    err := con.VerifyUserID(c, uintuserID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
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

    if err := repo.AddLike(uintuserID, uintImageID); err != nil {
        log.Printf("Error adding like: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "「いいね」の追加に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "「いいね」が追加されました"})
}

func (con *Controller) RemoveLike(c echo.Context) error {
    userID := c.QueryParam("userID")
    imageID := c.QueryParam("imageID")

    if userID == "" || imageID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ユーザーIDと画像IDが必要です"})
    }

    uintuserID := uintID(userID)
    if uintuserID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "fail to convert userID to uintID"})
    }
    err := con.VerifyUserID(c, uintuserID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
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

    if err := repo.RemoveLike(uintuserID, uintImageID); err != nil {
        log.Printf("Error removing like: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "「いいね」の削除に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "「いいね」が削除されました"})
}
