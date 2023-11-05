package recipe_repository

import (
	"context"
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/entity"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	_mongo "duck-cook-recipe/pkg/mongo"
)

type repositoryImpl struct {
	recipeCollection *mongo.Collection
}

func (repo repositoryImpl) CreateRecipe(recipe entity.Recipe) (entity.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	var recipeModel Recipe
	recipeModel = *recipeModel.FromEntity(recipe)
	timeNow := time.Now()
	recipeModel.CreatedAt = timeNow
	recipeModel.UpdatedAt = timeNow
	res, err := repo.recipeCollection.InsertOne(ctx, &recipeModel)

	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, writeErr := range writeErr.WriteErrors {
				if writeErr.Code == 11000 {
					errorMsg := writeErr.Message
					startIdx := strings.Index(errorMsg, "{")
					endIdx := strings.Index(errorMsg, "}")
					if startIdx != -1 && endIdx != -1 {
						fieldInfo := errorMsg[startIdx+1 : endIdx]

						re := regexp.MustCompile(`(\w+):`)
						match := re.FindStringSubmatch(fieldInfo)
						if len(match) >= 2 {
							fieldName := match[1]
							return recipeModel.ToEntityRecipe(), errors.New("duplicate " + fieldName)
						}
					}

				} else {
					log.Fatal(err)
				}
			}
		} else {
			log.Fatal(err)
		}
	}

	recipeModel.ID = res.InsertedID.(primitive.ObjectID)

	return recipeModel.ToEntityRecipe(), nil
}

func (repo repositoryImpl) DeleteRecipe(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	_, err := repo.recipeCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repo repositoryImpl) GetRecipe(id string) (recipe entity.Recipe, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	var recipeModel Recipe
	err = repo.recipeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(recipeModel)
	return recipeModel.ToEntityRecipe(), err
}

func (repo repositoryImpl) GetRecipesByUser(user string) (recipes []entity.Recipe, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	curso, err := repo.recipeCollection.Find(ctx, bson.M{"idUser": user})

	for curso.Next(ctx) {
		var recipeModel Recipe

		if err := curso.Decode(&recipeModel); err != nil {
			fmt.Println(err)
		}

		recipes = append(recipes, recipeModel.ToEntityRecipe())
	}

	return
}

func (repo repositoryImpl) UpdateRecipe(recipe entity.Recipe) (entity.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	var recipeModel Recipe
	recipeModel = *recipeModel.FromEntity(recipe)
	recipeModel.UpdatedAt = time.Now()

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(recipeModel)
	json.Unmarshal(inrec, &inInterface)

	update := bson.M{
		"$set": inInterface,
	}

	_, err := repo.recipeCollection.UpdateOne(ctx, bson.M{"_id": recipeModel.ID}, update)

	return recipeModel.ToEntityRecipe(), err
}

func New(mongoDb mongo.Database) repository.RecipeRepository {
	customerCollection := mongoDb.Collection(_mongo.COLLETCTION_RECIPE)
	return &repositoryImpl{
		customerCollection,
	}
}
