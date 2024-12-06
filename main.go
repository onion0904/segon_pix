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

	e.POST("/signup", con.Signup)
    e.POST("/verifyAddUser", con.VerifyAddUser)
	e.POST("/login", con.Login)
	
	method := e.Group("/segon_pix")
	{	
		method.GET("/get/user", con.UserInfo)
		method.GET("/get/list/search", con.SearchImage)
		method.GET("/get/list/like", con.GetLikeImages)
		method.GET("/get/list/recent", con.GetRecentImages)
		method.GET("/get/image_detail", con.ImageInfo)
	}


	secret := os.Getenv("JWT_SECRET_KEY")
	jwtSecret := []byte(secret)
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.MyCustomClaims)
		},
		SigningKey: jwtSecret,
	}

	// 認証が必要なルートグループを作成
	r := e.Group("/segon_pix_auth")
	r.Use(echojwt.WithConfig(config))

	// 認証が必要なルート
	{
		r.POST("/add/image", con.AddPostedImage)
		r.POST("/add/like", con.AddLike)
		r.POST("/add/comment", con.AddComment)
		r.POST("/add/follow", con.AddFollow)
		r.GET("", con.Restricted)
		r.GET("/get/user",con.UserInfoAuth)
		r.PUT("/update/user", con.UpdateUserInfo)
		r.PUT("/update/user/icon", con.UpdateUserIcon)
		r.PUT("/update/user/header", con.UpdateUserHeader)
		r.PUT("/update/comment", con.UpdateComment)
		r.DELETE("/delete/user", con.DeleteUser)
		r.DELETE("/delete/image", con.DeletePostedImage)
		r.DELETE("/delete/like", con.RemoveLike)
		r.DELETE("/delete/comment", con.DeleteComment)
		r.DELETE("/delete/follow", con.RemoveFollow)
	}
	// サーバーの開始
	e.Logger.Fatal(e.Start(":8080"))
}
