package repository

import "duck-cook-recipe/entity"

type RecipeRepository interface {
	CreateRecipe(recipe entity.Recipe) (entity.RecipeResponse, error)
	GetRecipe(id string) (recipe entity.RecipeResponse, err error)
	UpdateRecipe(recipe entity.Recipe) (entity.RecipeResponse, error)
	GetAllRecipe(page int, name, ingredient string) (pagination entity.Pagination, err error)
	DeleteRecipe(id string) error
	GetRecipesByUser(user string) (recipes []entity.RecipeResponse, err error)
	GetRecipesMoreLike() (recipes []entity.RecipeResponse, err error)
}
