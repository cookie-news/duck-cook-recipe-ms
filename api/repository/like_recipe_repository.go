package repository

import "duck-cook-recipe/entity"

type LikeRecipeRepository interface {
	LikeRecipeByUser(like entity.LikeRecipe) (entity.LikeRecipe, error)
	CheckRecipeIsLikedByUser(like entity.LikeRecipe) (bool, error)
	GetLikesByRecipe(idRecipe string) (count int64, err error)
	DeleteLikeRecipeByUser(like entity.LikeRecipe) error
}
