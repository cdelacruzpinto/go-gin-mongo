package main

import (
	"github.com/cdelacruzpinto/controllers"
	"github.com/cdelacruzpinto/middleware"

	"github.com/gin-gonic/gin"
)




func setupRouter() *gin.Engine {
	authMiddleware := middleware.NewAuthMiddleware()	
	router := gin.Default()
	router.GET("/recipes", controllers.GetRecipes)
	router.POST("/recipes", controllers.AddRecipe)
	router.DELETE("/recipes/:recipeID", controllers.DeleteRecipe)
	router.GET("/users", controllers.GetUsers)
	router.POST("/login", authMiddleware.LoginHandler)
	router.Use(authMiddleware.MiddlewareFunc()) 
	{
		router.GET("/hello", controllers.HelloHandler)
	}
	return router

}
func main() {
	
	router := setupRouter()
	router.Run(":8080")
}
