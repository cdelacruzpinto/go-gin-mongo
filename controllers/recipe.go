package controllers

import (
	"net/http"

	"github.com/cdelacruzpinto/models"
	"github.com/cdelacruzpinto/repository"
	"github.com/gin-gonic/gin"
)

func GetRecipes(ctx *gin.Context) {
	recipes := repository.GetRecipes()
	if len(recipes) == 0 {
		ctx.JSON(201, gin.H{"result": ""})
		return
	}
	ctx.JSON(http.StatusOK, recipes)

}
func AddRecipe(ctx *gin.Context) {
	var recipe *models.Recipe
	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.AddRecipe(recipe)
	ctx.JSON(http.StatusAccepted, nil)

}
func DeleteRecipe(ctx *gin.Context) {
	id := ctx.Param("recipeID")
	if id == "" {
		ctx.JSON(404, nil)
	}
	repository.DeleteRecipe(id)
	ctx.JSON(http.StatusAccepted, nil)

}
