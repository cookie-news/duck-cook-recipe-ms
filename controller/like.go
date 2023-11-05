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
// @Param       id        path      int  true  "User ID"
// @Param       idRecipe  path      int  true  "User ID"
// @Success		200		{object}	entity.Recipe
// @Router		/user/{id}/recipe/{idRecipe}/like [post]
func (c *Controller) LikeRecipeUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")

	var likeRecipe entity.LikeRecipe

	likeRecipe.IdUser = userId
	likeRecipe.Id = recipeId

	_, err := c.likeRecipeRepository.LikeRecipeByUser(likeRecipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusNoContent, "")
}

// @Summary		Likes da receita
// @Description	likes da receita
// @Tags		like-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      int  true  "Recipe ID"
// @Success		200		{int}
// @Router		/recipe/{id}/comment [get]
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
// @Param       id        path      int  true   "User ID"
// @Param       idRecipe  path      int  true   "Recipe ID"
// @Param       idLike  path      int  true    "Like ID"
// @Success     204 {string} string "No Content"
// @Router		/user/{id}/recipe/{idRecipe}/like/{idLike} [delete]
func (c *Controller) DeleteLikeHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")
	idLike := ctx.Param("idLike")

	err := c.likeRecipeRepository.DeleteLikeRecipeByUser(entity.LikeRecipe{
		IdLike: idLike,
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
