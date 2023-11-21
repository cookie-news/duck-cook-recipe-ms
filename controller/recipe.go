package controller

import (
	"duck-cook-recipe/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
// @Param       authorization        header      string  true  "Token Bearer"
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
		c.recipeRepository.DeleteRecipe(recipeResponse.Id)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
// @Param       authorization        header      string  true  "Token Bearer"
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
// @Param        id   path      string  true  "Recipe ID"
// @Success		200		{object}	entity.RecipeCountLikeManyComments
// @Router		/recipe/{id} [get]
func (c *Controller) GetRecipeHandler(ctx *gin.Context) {
	recipeId := ctx.Param("id")

	recipe, err := c.recipeUseCase.GetRecipe(recipeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.Images, _ = c.storageUseCase.ListFiles(recipeId)

	ctx.JSON(http.StatusOK, recipe)
}

// @Summary		Retorna as receita paginadas
// @Description	Retorna as receita paginadas
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param        page   path      int  true  "Número da page"
// @Param        nameRecipe   query      string  false  "Recipe name"
// @Param        nameIngredient   query      string  false  "Número da page"
// @Success		200		{object}	entity.Pagination
// @Router		/recipe/page/{page} [get]
func (c *Controller) GetPageRecipesHandler(ctx *gin.Context) {
	pageStr := ctx.Param("page")
	nameRecipe := ctx.Query("nameRecipe")
	nameIngredient := ctx.Query("nameIngredient")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page not valid"})
		return
	}

	pageRecipes, err := c.recipeUseCase.GetRecipeByPage(page, nameRecipe, nameIngredient)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipes := pageRecipes.Items.([]entity.RecipeCountLikeManyComments)

	for index, item := range recipes {
		files, _ := c.storageUseCase.ListFiles(item.Id)
		recipes[index].Images = files
	}

	pageRecipes.Items = recipes

	ctx.JSON(http.StatusOK, pageRecipes)
}

// @Summary		Retonar as receitas do usuário
// @Description	Retonar as receitas do usuário
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param        id   path      string  true  "User ID"
// @Param       authorization        header      string  true  "Token Bearer"
// @Success		200		{object}	entity.Recipe
// @Router		/user/{id}/recipe [get]
func (c *Controller) GetRecipeUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	recipes, err := c.recipeRepository.GetRecipesByUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, recipe := range recipes {
		files, _ := c.storageUseCase.ListFiles(recipe.Id)
		recipes[index].Images = files
	}
	ctx.JSON(http.StatusOK,
		recipes,
	)
}

// @Summary		Adicionar nova receita
// @Description	Adicionar uma nova receita
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Param        id   path      int  true  "Recipe ID"
// @Param       authorization        header      string  true  "Token Bearer"
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

// @Summary		Adicionar nova receita
// @Description	Adicionar uma nova receita
// @Tags		recipe
// @Accept		json
// @Produce		json
// @Success     200   {object}	entity.RecipeResponse[]
// @Router		/recipe/more-like [GET]
func (c Controller) GetRecipesMoreLikeHandler(ctx *gin.Context) {
	recipes, err := c.recipeRepository.GetRecipesMoreLike()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, recipes)
}
