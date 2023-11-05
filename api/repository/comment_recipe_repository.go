package repository

import "duck-cook-recipe/entity"

type CommentRecipeRepository interface {
	CommentRecipeByUser(commentRecipe entity.CommentRecipe) (entity.CommentRecipe, error)
	GetCommentsByRecipe(idRecipe string) (comments []entity.CommentRecipe, err error)
	DeleteCommentRecipeByUser(commentRecipe entity.CommentRecipe) error
}
