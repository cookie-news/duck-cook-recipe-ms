package controller

import (
	"duck-cook-recipe/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Adicionar nova receita
// @Description	Adicionar uma nova receita
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param		payload	body		entity.Recipe	true	"Dados da receita"
// @Success		200		{object}	entity.Recipe
// @Router		/recipe [post]
func (c *Controller) CreateRecipeHandler(ctx *gin.Context) {
	var recipe entity.Recipe
	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON"})
		return
	}

	recipe, err := c.recipeRepository.CreateRecipe(recipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated,
		recipe,
	)
}

// @Summary		Atualizar a receita
// @Description	Atualizar a receita
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param		payload	body		entity.Recipe	true	"Dados da receita"
// @Success		200		{object}	entity.Recipe
// @Router		/recipe [put]
func (c *Controller) UpdateRecipeHandler(ctx *gin.Context) {
	var recipe entity.Recipe
	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON"})
		return
	}

	recipe, err := c.recipeRepository.UpdateRecipe(recipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

// @Summary		Retorna a receita
// @Description	Retorna a receita
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param        id   path      int  true  "Recipe ID"
// @Success		200		{object}	entity.Recipe
// @Router		/recipe/{id} [get]
func (c *Controller) GetRecipeHandler(ctx *gin.Context) {
	recipeId := ctx.Param("id")

	recipe, err := c.recipeRepository.GetRecipe(recipeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

// @Summary		Retonar as receitas do usuário
// @Description	Retonar as receitas do usuário
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param        id   path      int  true  "User ID"
// @Success		200		{object}	entity.Recipe
// @Router		/user/{id}/recipe [get]
func (c *Controller) GetRecipeUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	recipe, err := c.recipeRepository.GetRecipesByUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,
		recipe,
	)
}

// @Summary		Adicionar nova receita
// @Description	Adicionar uma nova receita
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param        id   path      int  true  "Recipe ID"
// @Success     204   "No Content"
// @Router		/recipe/{id} [delete]
func (c *Controller) DeleteRecipeHandler(ctx *gin.Context) {
	recipeId := ctx.Param("id")

	err := c.recipeRepository.DeleteRecipe(recipeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusNoContent, "")
}

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON"})
		return
	}

	commentRecipe.IdUser = userId
	commentRecipe.Id = recipeId

	comment, err := c.commentRecipe.CommentRecipeByUser(commentRecipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)
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
func (c *Controller) DeleteRecipeUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	recipeId := ctx.Param("idRecipe")
	idComment := ctx.Param("idComment")

	err := c.commentRecipe.DeleteCommentRecipeByUser(entity.CommentRecipe{
		Id: idComment,
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
