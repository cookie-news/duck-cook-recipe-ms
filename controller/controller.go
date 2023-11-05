package controller

import "duck-cook-recipe/api/repository"

type Controller struct {
	recipeRepository repository.RecipeRepository
}

func NewController(
	recipeRepository repository.RecipeRepository,
) Controller {
	return Controller{
		recipeRepository,
	}
}
