package controller

import (
	"duck-cook-recipe/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// @Summary		Adicionar nova receita
// @Description	Adicionar uma nova receita
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param id formData string true "ID da Receita"
// @Param idUser formData string true "ID do Usuário"
// @Param title formData string true "Título da Receita"
// @Param description formData string true "Descrição da Receita"
// @Param preparationTime formData int true "Tempo de Preparação da Receita"
// @Param preparationMethod formData string true "Metodo de Preparação da Receita"
// @Param images formData file true "Imagem da Receita"
// @Param ingredients formData string true "ingredientes da Receita"
// @Success		200		{object}	entity.Recipe
// @Router		/recipe [post]
func (c *Controller) CreateRecipeHandler(ctx *gin.Context) {
	var recipe entity.Recipe
	if err := ctx.ShouldBind(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipeResponse, err := c.recipeRepository.CreateRecipe(recipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := c.recipeStorage.UploadImage(recipe.Images, recipeResponse.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred while saving your profile photo, but the username was created successfully"})
		return
	}

	recipeResponse.Images = url

	ctx.JSON(http.StatusCreated,
		recipeResponse,
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

	recipeResponse, err := c.recipeRepository.UpdateRecipe(recipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, recipeResponse)
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

	recipe.Images, _ = c.recipeStorage.ListFiles(recipeId)

	ctx.JSON(http.StatusOK, recipe)
}

// @Summary		Retorna as receita paginadas
// @Description	Retorna as receita paginadas
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param        page   path      int  true  "Número da page"
// @Param        nameRecipe   query      int  true  "Recipe name"
// @Param        nameIngredin   query      int  true  "Número da page"
// @Success		200		{object}	entity.Pagination
// @Router		/page/{page} [get]
func (c *Controller) GetPageRecipesHandler(ctx *gin.Context) {
	pageStr := ctx.Param("page")
	nameRecipe := ctx.Query("nameRecipe")
	nameIngredin := ctx.Query("nameIngredin")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page not valid"})
		return
	}

	pageRecipes, err := c.recipeRepository.GetAllRecipe(page, nameRecipe, nameIngredin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var itens []map[string]interface{}

	mapstructure.Decode(pageRecipes.Items, &itens)

	for _, item := range itens {
		files, _ := c.recipeStorage.ListFiles(item["Recipe"].(map[string]interface{})["Id"].(string))
		item["Images"] = files
	}

	pageRecipes.Items = itens

	ctx.JSON(http.StatusOK, pageRecipes)
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

	/*
		@TODO FAZER URL DA IMAGEM
	*/

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
