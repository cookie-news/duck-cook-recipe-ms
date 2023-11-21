package usecase

import (
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/entity"
)

type RecipeUseCase interface {
	GetRecipeByPage(page int, name, ingredient string) (pagination entity.Pagination, err error)
	GetRecipe(id string) (recipe entity.RecipeCountLikeManyComments, err error)
}

type recipeUseCaseImpl struct {
	recipeRepository        repository.RecipeRepository
	likeRecipeRepository    repository.LikeRecipeRepository
	commentRecipeRepository repository.CommentRecipeRepository
}

func (usecase recipeUseCaseImpl) GetRecipe(id string) (recipe entity.RecipeCountLikeManyComments, err error) {
	recipeResult, err := usecase.recipeRepository.GetRecipe(id)
	if err != nil {
		return
	}
	countLikes, err := usecase.likeRecipeRepository.GetLikesByRecipe(id)
	if err != nil {
		return
	}

	comments, err := usecase.commentRecipeRepository.GetCommentsByRecipe(id)
	if err != nil {
		return
	}
	recipe = entity.RecipeCountLikeManyComments{
		RecipeResponse: recipeResult,
		CountLikes:     int(countLikes),
		CountComments:  len(comments),
	}
	return
}

func (usecase recipeUseCaseImpl) GetRecipeByPage(page int, name string, ingredient string) (pagination entity.Pagination, err error) {
	pagination, err = usecase.recipeRepository.GetAllRecipe(page, name, ingredient)

	recipes := pagination.Items.([]entity.RecipeResponse)
	var recipesResult []entity.RecipeCountLikeManyComments

	for _, recipe := range recipes {
		countLikes, err := usecase.likeRecipeRepository.GetLikesByRecipe(recipe.Id)
		if err != nil {
			break
		}

		comments, err := usecase.commentRecipeRepository.GetCommentsByRecipe(recipe.Id)
		if err != nil {
			break
		}

		recipesResult = append(recipesResult, entity.RecipeCountLikeManyComments{
			RecipeResponse: recipe,
			CountLikes:     int(countLikes),
			CountComments:  len(comments),
		})
	}

	pagination.Items = recipesResult

	return
}

func NewRecipeUseCase(
	recipeRepository repository.RecipeRepository,
	likeRecipeRepository repository.LikeRecipeRepository,
	commentRecipeRepository repository.CommentRecipeRepository,

) RecipeUseCase {
	return &recipeUseCaseImpl{
		recipeRepository,
		likeRecipeRepository,
		commentRecipeRepository,
	}
}
