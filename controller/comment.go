package controller

import (
	"duck-cook-recipe/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Comenta na receita
// @Description	Comenta na receita baseado no usuário
// @Tags		comment-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      int  true  "User ID"
// @Param       idRecipe  path      int  true  "User ID"
// @Param		payload	  body		entity.CommentRecipe	true	"Comentário"
// @Success		200		{object}	entity.Recipe
// @Router		/user/{id}/recipe/{idRecipe}/comment [post]
func (c *Controller) CommentRecipeUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")

	var commentRecipe entity.CommentRecipe
	if err := ctx.ShouldBindJSON(&commentRecipe); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON"})
		return
	}

	commentRecipe.IdUser = userId
	commentRecipe.IdRecipe = recipeId

	comment, err := c.commentRecipeUseCase.CommentRecipeByUser(commentRecipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// @Summary		Comentarios da receita
// @Description	Comentarios da receita
// @Tags		comment-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      int  true  "Recipe ID"
// @Success		200		{object}	[]entity.CommentRecipe
// @Router		/recipe/{id}/comment [get]
func (c *Controller) GetCommentsHandler(ctx *gin.Context) {
	recipeId := ctx.Param("id")
	auth := ctx.GetHeader("authorization")
	comments, err := c.commentRecipeUseCase.GetCommentsByRecipe(recipeId, auth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// @Summary		Deleta o comentário da receita
// @Description	Delete o comentário da receita
// @Tags		comment-recipe
// @Accept		json
// @Produce		json
// @Param       id        path      int  true   "User ID"
// @Param       idRecipe  path      int  true   "Recipe ID"
// @Param       idComment  path      int  true  "Comment ID"
// @Success     204 {string} string "No Content"
// @Router		/user/{id}/recipe/{idRecipe}/comment/{idComment} [delete]
func (c *Controller) DeleteCommentHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")
	idComment := ctx.Param("idComment")

	err := c.commentRecipeUseCase.DeleteCommentRecipeByUser(entity.CommentRecipe{
		IdComment: idComment,
		IdRecipe:  recipeId,
		IdUser:    userId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusNoContent, "")
}
