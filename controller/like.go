package controller

import (
	"duck-cook-recipe/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Like na receita
// @Description	Like na receita baseado no usuário
// @Tags		like-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      string  true  "User ID"
// @Param       idRecipe  path      string  true  "Recipe ID"
// @Param       authorization        header      string  true  "Token Bearer"
// @Success		200		{object}	entity.Recipe
// @Router		/user/{id}/recipe/{idRecipe}/like [post]
func (c *Controller) LikeRecipeUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")

	likeRecipe := entity.LikeRecipe{
		Recipe: entity.Recipe{
			Id:     recipeId,
			IdUser: userId,
		},
	}

	_, err := c.likeRecipeRepository.LikeRecipeByUser(likeRecipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusNoContent, "")
}

// @Summary		Verifica o Like do usuário na receita
// @Description	Verifica o Like do usuário na receita
// @Tags		like-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      string  true  "User ID"
// @Param       idRecipe  path      string  true  "Recipe ID"
// @Param       authorization        header      string  true  "Token Bearer"
// @Router		/user/{id}/recipe/{idRecipe}/like [get]
func (c *Controller) CheckUserLikedRecipeHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")

	likeRecipe := entity.LikeRecipe{
		Recipe: entity.Recipe{
			Id:     recipeId,
			IdUser: userId,
		},
	}

	liked, err := c.likeRecipeRepository.CheckRecipeIsLikedByUser(likeRecipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"liked": liked,
	})
}

// @Summary		Likes da receita
// @Description	likes da receita
// @Tags		like-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      int  true  "Recipe ID"
// @Success		200		{int}  100
// @Router		/recipe/{id}/like [get]
func (c *Controller) GetLikesHandler(ctx *gin.Context) {
	recipeId := ctx.Param("id")

	countLikes, err := c.likeRecipeRepository.GetLikesByRecipe(recipeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": countLikes,
	})
}

// @Summary		Deleta o like da receita
// @Description	Delete o like da receita
// @Tags		like-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      string  true   "User ID"
// @Param       idRecipe  path      string  true   "Recipe ID"
// @Param       idLike  path      string  true    "Like ID"
// @Param       authorization        header      string  true  "Token Bearer"
// @Success     204 {string} string "No Content"
// @Router		/user/{id}/recipe/{idRecipe}/like [delete]
func (c *Controller) DeleteLikeHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")

	err := c.likeRecipeRepository.DeleteLikeRecipeByUser(entity.LikeRecipe{
		Recipe: entity.Recipe{
			Id:     recipeId,
			IdUser: userId,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusNoContent, "")
}
