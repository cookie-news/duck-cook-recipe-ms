package controller

import "duck-cook-recipe/api/repository"

type Controller struct {
	recipeRepository repository.RecipeRepository
	commentRecipe    repository.CommentRecipeRepository
}

func NewController(
	recipeRepository repository.RecipeRepository,
	commentRecipe repository.CommentRecipeRepository,
) Controller {
	return Controller{
		recipeRepository,
		commentRecipe,
	}
}
