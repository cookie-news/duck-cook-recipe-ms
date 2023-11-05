package repository

import "duck-cook-recipe/entity"

type CommentRecipeRepository interface {
	CommentRecipeByUser(commentRecipe entity.CommentRecipe) (entity.CommentRecipe, error)
	DeleteCommentRecipeByUser(commentRecipe entity.CommentRecipe) error
}
