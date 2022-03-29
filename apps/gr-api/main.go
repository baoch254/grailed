package main

import (
	"log"
	"net/http"
	"os"

	"grailed/libs/gr-api/feature-auth/authtransport/ginauth"
	common "grailed/libs/gr-api/shared-common"
	middleware "grailed/libs/gr-api/shared-middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DBConnectionStr")
	secretKey := os.Getenv("SECRET_KEY")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := runService(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, secretKey string) error {
	appCtx := common.NewAppContext(db, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	v1.POST("/register", ginauth.Register(appCtx))
	v1.POST("/login", ginauth.Login(appCtx))

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
