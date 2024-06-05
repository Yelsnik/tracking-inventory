package main

import (
	"fmt"
	"net/http"
	"os"

	//"gorm.io/gorm"

	"github.com/Yelsnik/tracking-inventory.git/internal/app"
	"github.com/Yelsnik/tracking-inventory.git/internal/initializers"
	"github.com/gin-gonic/gin"
)

func initialize() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World!")
	})

	app.Routes(r)

	initialize()

	initializers.ConnectToDatabase(
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
	)
	fmt.Println("success!")

	r.Run(":5000")
}
