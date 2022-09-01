package models

import "github.com/kamva/mgm/v3"

type Recipe struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func NewRecipe(name string) *Recipe {
	return &Recipe{
		Name: name,
	}
}
