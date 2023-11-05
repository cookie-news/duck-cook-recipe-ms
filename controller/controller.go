package controller

import "duck-cook-recipe/api/repository"

type Controller struct {
	recipeRepository     repository.RecipeRepository
	commentRecipe        repository.CommentRecipeRepository
	likeRecipeRepository repository.LikeRecipeRepository
}

func NewController(
	recipeRepository repository.RecipeRepository,
	commentRecipeRepository repository.CommentRecipeRepository,
	likeRecipeRepository repository.LikeRecipeRepository,
) Controller {
	return Controller{
		recipeRepository,
		commentRecipeRepository,
		likeRecipeRepository,
	}
}
