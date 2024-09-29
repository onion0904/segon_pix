package main

import (
	"fmt"
	"log"
	"os"

	"PixApp/controllers"
	"PixApp/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 環境変数の取得
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")

	// データベース接続の設定
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// データベースのマイグレーション
	err = db.AutoMigrate(&models.User{}, &models.PostedImage{}, &models.LikedImage{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Echoのインスタンスを作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// コントローラーの作成
	con := controllers.NewTodoController(db)

	// ルーティングの設定
	method := e.Group("/Todo")
	{
		method.POST("/add", con.Add)
		method.GET("/list", con.List)
		method.PUT("/update", con.Update)
		method.DELETE("/delete", con.Delete)
	}

	// サーバーの開始
	e.Logger.Fatal(e.Start(":8080"))
}
