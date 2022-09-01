package repository

import (
	"github.com/cdelacruzpinto/models"
	"github.com/kamva/mgm/v3"
)

func GetRecipes() []models.Recipe {
	return []models.Recipe{}
}

func AddRecipe(recipe *models.Recipe) {
	mgm.Coll(recipe).Create(recipe)
}
func DeleteRecipe(id interface{}) {
	recipe := &models.Recipe{}
	id, _ = recipe.PrepareID(id.(string))
	mgm.Coll(recipe).FindByID(id, recipe)
	mgm.Coll(recipe).Delete(recipe)
}
