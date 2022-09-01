package main

import (
	"github.com/cdelacruzpinto/controllers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/recipes", controllers.GetRecipes)
	router.POST("/recipes", controllers.AddRecipe)
	router.DELETE("/recipes/:recipeID", controllers.DeleteRecipe)
	router.GET("/users", controllers.GetUsers)
	return router

}
func main() {
	router := setupRouter()
	router.Run(":8080")
}
