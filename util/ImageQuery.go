package util

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ListQueryParams struct {
	Hashtag   string
	Search    bool
	Like      bool
	LikeNum   int
	CurrentID int
}

func ParseListQueryParams(c echo.Context) (*ListQueryParams, error) {
	q := &ListQueryParams{}

	// Hashtag: 文字列のまま
	q.Hashtag = c.QueryParam("Hashtag")

	// Search: "true" → true、それ以外 → false
	q.Search = strings.ToLower(c.QueryParam("Search")) == "true"

	// Like: "true" → true、それ以外 → false
	q.Like = strings.ToLower(c.QueryParam("Like")) == "true"

	// like_num: 数値（未指定や変換失敗時は -1）
	if numStr := c.QueryParam("like_num"); numStr != "" {
		if num, err := strconv.Atoi(numStr); err == nil {
			q.LikeNum = num
		} else {
			q.LikeNum = -1
		}
	} else {
		q.LikeNum = -1
	}

	// current: 数値（未指定や変換失敗時は -1）
	if idStr := c.QueryParam("current"); idStr != "" {
		if id, err := strconv.Atoi(idStr); err == nil {
			q.CurrentID = id
		} else {
			q.CurrentID = -1
		}
	} else {
		q.CurrentID = -1
	}

	return q, nil
}