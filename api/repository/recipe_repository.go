package repository

import "duck-cook-recipe/entity"

type RecipeRepository interface {
	CreateRecipe(recipe entity.Recipe) (entity.Recipe, error)
	GetRecipesByUser(user string) (recipes []entity.Recipe, err error)
	GetRecipe(id string) (recipe entity.Recipe, err error)
	UpdateRecipe(recipe entity.Recipe) (entity.Recipe, error)
	DeleteRecipe(id string) error
}
