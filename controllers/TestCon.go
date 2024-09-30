package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"PixApp/models"
	"PixApp/repositories"
	"gorm.io/gorm"
)

type TodoController struct {
	db *gorm.DB
}

// コンストラクタ
func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{db: db}
}

// 追加
func (con *TodoController) Add(c echo.Context) error {
	todo := models.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}

	repo := repositories.NewTodoRepository(con.db)
	if err := repo.Add(&todo); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusNoContent) // 204エラー
}

// リスト取得
func (con *TodoController) List(c echo.Context) error {
	Qtitle := c.QueryParam("Title")
	repo := repositories.NewTodoRepository(con.db)
	TodoList, err := repo.List(Qtitle)
	if err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"list": TodoList,
	})
}

// 更新
func (con *TodoController) Update(c echo.Context) error {
	todo := models.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}

	repo := repositories.NewTodoRepository(con.db)

	var existingTodo models.Todo
	if err := con.db.First(&existingTodo, todo.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound) // 404エラー
		}
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}

	if err := repo.Update(&todo); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusNoContent) // 204エラー
}

// 削除
func (con *TodoController) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest) // 400エラー
	}

	var existingTodo models.Todo
	if err := con.db.First(&existingTodo, intId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound) // 404エラー
		}
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}

	repo := repositories.NewTodoRepository(con.db)
	if err := repo.Delete(intId); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusNoContent) // 204エラー
}
