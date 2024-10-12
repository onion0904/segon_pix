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
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 環境変数の取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")

	// データベース接続の設定

	dsn := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?parseTime=true", dbUser, dbPassword, os.Getenv("DB_PORT"),dbDatabase)
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

	e.POST("/login", con.Login)
    e.POST("/verify", con.Verify)

	method := e.Group("/segon_pix")
	{	
		method.POST("/add/user", con.AddUser)
		method.POST("/add/image", con.AddPostedImage)
		method.POST("/add/like", con.AddLike)
		method.POST("/add/comment", con.AddComment)
		method.GET("/get/user", con.UserInfo)
		method.GET("/get/list/image", con.SearchImage)
		method.GET("/get/image", con.ImageInfo)
		method.PUT("/update/user", con.UpdateUserInfo)
		method.PUT("/update/user/icon", con.UpdateUserIcon)
		method.PUT("/update/user/header", con.UpdateUserHeader)
		method.PUT("/update/comment", con.UpdateComment)
		method.DELETE("/delete/user", con.DeleteUser)
		method.DELETE("/delete/image", con.DeletePostedImage)
		method.DELETE("/delete/like", con.RemoveLike)
		method.DELETE("/delete/comment", con.DeleteComment)
	}

    // 認証が必要なルートグループを作成
    r := e.Group("/restricted")
	secret := os.Getenv("JWT_SECRET_KEY")
	jwtSecret := []byte(secret)
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(controllers.MyCustomClaims)
		},
		SigningKey: jwtSecret,
	}
	r.Use(echojwt.WithConfig(config))

	// 認証が必要なルート
	{
		r.GET("", con.Restricted)
	}
	// サーバーの開始
	e.Logger.Fatal(e.Start(":8080"))
}
