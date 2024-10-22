package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"PixApp/repositories"
	"log"
)


func (con *Controller) AddFollow(c echo.Context) error {
    followingID := c.QueryParam("followingID")
    followedID := c.QueryParam("followedID")

    if followingID == "" || followedID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "フォローするユーザーIDとフォローされるユーザーIDが必要です"})
    }

    uintFollowingID := uintID(followingID)
    if uintFollowingID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintFollowingID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    uintFollowedID64, err := strconv.ParseUint(followedID, 10, 64)
    if err != nil {
        log.Printf("Invalid followedID: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "無効なフォローされるユーザーIDです"})
    }
    uintFollowedID := uint(uintFollowedID64)

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error initializing repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    if err := repo.AddFollow(uintFollowingID, uintFollowedID); err != nil {
        log.Printf("Error adding follow: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "フォローに失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "フォローしました"})
}

func (con *Controller) RemoveFollow(c echo.Context) error {
    followingID := c.QueryParam("followingID")
    followedID := c.QueryParam("followedID")

    if followingID == "" || followedID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "フォロー解除するユーザーIDとフォロー解除されるユーザーIDが必要です"})
    }

    uintFollowingID := uintID(followingID)
    if uintFollowingID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintFollowingID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    uintFollowedID64, err := strconv.ParseUint(followedID, 10, 64)
    if err != nil {
        log.Printf("Invalid followedID: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "無効なフォロー解除されるユーザーIDです"})
    }
    uintFollowedID := uint(uintFollowedID64)

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error initializing repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    if err := repo.RemoveFollow(uintFollowingID, uintFollowedID); err != nil {
        log.Printf("Error removing follow: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "フォロー解除に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "フォロー解除しました"})
}
