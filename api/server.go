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

	docs.SwaggerInfo.Title = "Duck Cook Receipe"
	docs.SwaggerInfo.Description = "Duck Cook Receipe"
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
		auth := ctx.GetHeader("authorization")

		client := resty.New()
		client.BaseURL = os.Getenv("URL_AUTH")

		resp, _ := client.R().
			SetHeader("authorization", auth).
			Post("/v1/auth/verify-jwt")

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
		v1.GET("/user/:id/recipe", s.controller.GetRecipeUserHandler)
		recipe := v1.Group("/recipe")
		{
			recipe.GET("/", s.controller.GetRecipeHandler)
			recipe.POST("/", s.controller.CreateRecipeHandler)
			recipe.PUT("/", s.controller.UpdateRecipeHandler)
			recipe.DELETE("/:id", s.controller.DeleteRecipeHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r.Run(":" + addr)
}
