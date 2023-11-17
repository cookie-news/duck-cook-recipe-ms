package usecase

import (
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/entity"
)

type CommentRecipeUseCase interface {
	GetCommentsByRecipe(idRecipe, token string) (comments []entity.CommentRecipe, err error)
	CommentRecipeByUser(commentRecipe entity.CommentRecipe) (entity.CommentRecipe, error)
	DeleteCommentRecipeByUser(commentRecipe entity.CommentRecipe) error
}

type commentRecipeUseCaseImpl struct {
	commentRecipeRepository repository.CommentRecipeRepository
	userUseCase             UserUseCase
}

func (usecase commentRecipeUseCaseImpl) CommentRecipeByUser(commentRecipe entity.CommentRecipe) (entity.CommentRecipe, error) {
	return usecase.commentRecipeRepository.CommentRecipeByUser(commentRecipe)
}

func (usecase commentRecipeUseCaseImpl) DeleteCommentRecipeByUser(commentRecipe entity.CommentRecipe) error {
	return usecase.commentRecipeRepository.DeleteCommentRecipeByUser(commentRecipe)
}

func (usecase commentRecipeUseCaseImpl) GetCommentsByRecipe(idRecipe, token string) (comments []entity.CommentRecipe, err error) {
	comments, err = usecase.commentRecipeRepository.GetCommentsByRecipe(idRecipe)
	if err != nil {
		return
	}

	for index, comment := range comments {
		user, err := usecase.userUseCase.GetUserById(comment.IdUser, token)
		if err != nil {
			break
		}
		comments[index].User = user
	}
	return
}

func NewCommentRecipeUseCase(
	commentRecipeRepository repository.CommentRecipeRepository,
	userUseCase UserUseCase,
) CommentRecipeUseCase {
	return &commentRecipeUseCaseImpl{
		commentRecipeRepository,
		userUseCase,
	}
}
