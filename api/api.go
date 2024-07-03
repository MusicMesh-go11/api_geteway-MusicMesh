package api

import "github.com/gin-gonic/gin"

func API() *gin.Engine {

	router := gin.Default()

	user := router.Group("/user")
	composition := router.Group("/composition")

	user.POST("")
	user.GET("")
	user.GET("/:id")
	user.PUT("/:id")
	user.DELETE("/:id")

	user.POST("/:id/profile")
	user.GET("/profile")
	user.GET("/:id/profile")
	user.PUT("/:id/profile")
	user.DELETE("/:id/profile")

	composition.POST("")
	composition.GET("")
	composition.GET("/:id")
	composition.PUT("/:id")
	composition.DELETE("/:id")
	user.GET("/:id/compositions")

	composition.POST("/:id/track")
	composition.GET("/track")
	composition.GET("/:id/track")
	composition.PUT("/:id/track")
	composition.DELETE("/:id/track")

	return router
}
