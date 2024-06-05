package app

import (
	"github.com/Yelsnik/tracking-inventory.git/internal/controllers"
	"github.com/gin-gonic/gin"
)

// routes

func Routes(r *gin.Engine) {

	r.POST("/sign-in", controllers.SignIn)
	r.POST("/login", controllers.Login)
}
