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
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// データベースのマイグレーション
	err = db.AutoMigrate(&models.User{}, &models.PostedImage{},&models.Comment{},&models.Hashtag{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Echoのインスタンスを作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// コントローラーの作成
	con := controllers.NewController(db)

	// ルーティングの設定
	method := e.Group("/segon_pix")
	{
		method.POST("/add/user", con.AddUser)
		method.POST("/add/image", con.AddPostedImage)
		method.POST("/add/like", con.AddLike)
		method.POST("/add/comment", con.AddComment)
		method.GET("/list/user", con.UserInfo)
		method.GET("/list/image", con.SearchImage)
		method.PUT("/update/comment", con.UpdateComment)
		method.DELETE("/delete/user", con.DeleteUser)
		method.DELETE("/delete/image", con.DeletePostedImage)
		method.DELETE("/delete/like", con.RemoveLike)
		method.DELETE("/delete/comment", con.DeleteComment)
	}

	// サーバーの開始
	e.Logger.Fatal(e.Start(":8080"))
}
