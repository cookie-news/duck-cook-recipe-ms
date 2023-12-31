package api

import (
	"duck-cook-recipe/controller"
	"duck-cook-recipe/docs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	controller controller.Controller
}

func NewServer(controller controller.Controller) *Server {
	return &Server{controller}
}

func (s *Server) Start(addr string) error {

	docs.SwaggerInfo.Title = "Duck Cook Recipe"
	docs.SwaggerInfo.Description = "Duck Cook Recipe"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("DOMAIN_ALLOW"))
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.Use(func(ctx *gin.Context) {
		switch ctx.FullPath() {
		case
			"/swagger/*any",
			"/v1/recipe/page/:page",
			"/v1/recipe/:id/comment",
			"/v1/recipe/:id/like",
			"/v1/recipe/:id",
			"/v1/recipe/more-like":
			ctx.Next()
			return
		}

		auth := ctx.GetHeader("authorization")

		client := resty.New()
		client.BaseURL = os.Getenv("URL_AUTH")

		resp, err := client.R().
			SetHeader("authorization", auth).
			Post("/v1/auth/verify-jwt")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		if resp.StatusCode() == http.StatusNoContent {
			ctx.Next()
			return
		} else {
			ctx.String(resp.StatusCode(), resp.String())
			ctx.Abort()
			return
		}
	})

	v1 := r.Group("/v1")
	{
		user := v1.Group("/user/:id")
		{
			user.GET("/recipe", s.controller.GetRecipeUserHandler)
			user.GET("/recipe/like", s.controller.GetRecipeLikeUserHandler)
			recipe := user.Group("/recipe/:idRecipe")
			{
				comment := recipe.Group("comment")
				{
					comment.POST("", s.controller.CommentRecipeUserHandler)
					comment.DELETE("/:idComment", s.controller.DeleteCommentHandler)
				}
				like := recipe.Group("/like")
				{
					like.POST("", s.controller.LikeRecipeUserHandler)
					like.GET("", s.controller.CheckUserLikedRecipeHandler)
					like.DELETE("", s.controller.DeleteLikeHandler)
				}
			}
		}
		recipe := v1.Group("/recipe")
		{
			recipe.GET("/:id", s.controller.GetRecipeHandler)
			recipe.GET("/page/:page", s.controller.GetPageRecipesHandler)
			recipe.GET("/more-like", s.controller.GetRecipesMoreLikeHandler)
			recipe.POST("", s.controller.CreateRecipeHandler)
			recipe.PUT("", s.controller.UpdateRecipeHandler)
			recipe.DELETE("/:id", s.controller.DeleteRecipeHandler)
			recipe.GET("/:id/comment", s.controller.GetCommentsHandler)
			recipe.GET("/:id/like", s.controller.GetLikesHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r.Run(":" + addr)
}
