package controller

import "duck-cook-recipe/api/repository"

type Controller struct {
	recipeRepository     repository.RecipeRepository
	commentRecipe        repository.CommentRecipeRepository
	likeRecipeRepository repository.LikeRecipeRepository
	recipeStorage        repository.RecipeStorage
}

func NewController(
	recipeRepository repository.RecipeRepository,
	commentRecipeRepository repository.CommentRecipeRepository,
	likeRecipeRepository repository.LikeRecipeRepository,
	recipeStorage repository.RecipeStorage,
) Controller {
	return Controller{
		recipeRepository,
		commentRecipeRepository,
		likeRecipeRepository,
		recipeStorage,
	}
}
