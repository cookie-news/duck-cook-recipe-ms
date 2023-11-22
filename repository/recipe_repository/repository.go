package recipe_repository

import (
	"context"
	api_helper "duck-cook-recipe/api/helper"
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
	"go.mongodb.org/mongo-driver/mongo/options"

	_mongo "duck-cook-recipe/pkg/mongo"
)

type repositoryImpl struct {
	recipeCollection *mongo.Collection
}

func (repo repositoryImpl) GetRecipesLikedByUser(idUser string) (recipes []entity.RecipeResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectIdUser, err := primitive.ObjectIDFromHex(idUser)
	if err != nil {
		return
	}
	pipeline := []bson.M{
		{
			"$match": bson.M{"idUser": objectIdUser},
		},
		{
			"$lookup": bson.M{
				"from":         "LikeRecipe",
				"localField":   "_id",
				"foreignField": "idRecipe",
				"as":           "likes",
			},
		},
	}

	curso, err := repo.recipeCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	for curso.Next(ctx) {
		var recipe Recipe
		if err := curso.Decode(&recipe); err != nil {
			fmt.Println(err)
		}

		recipes = append(recipes, recipe.ToEntityRecipeResponse())
	}

	return
}

func (repo repositoryImpl) GetRecipesMoreLike() (recipes []entity.RecipeResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "LikeRecipe",
				"localField":   "_id",
				"foreignField": "idRecipe",
				"as":           "likes",
			},
		},
		{
			"$addFields": bson.M{
				"likeCount": bson.M{"$size": "$likes"},
			},
		},
		{
			"$sort": bson.M{
				"likeCount": -1,
			},
		},
		{
			"$limit": 6,
		},
	}

	curso, err := repo.recipeCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	for curso.Next(ctx) {
		var recipe Recipe
		if err := curso.Decode(&recipe); err != nil {
			fmt.Println(err)
		}

		recipes = append(recipes, recipe.ToEntityRecipeResponse())
	}

	return
}

func (repo repositoryImpl) GetAllRecipe(page int, name, ingredient string) (pagination entity.Pagination, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	limit := int64(10)
	p := int64(page)
	skip := int64(p*limit - limit)
	fOpt := options.FindOptions{Limit: &limit, Skip: &skip}

	filterFields := []bson.M{}

	if name != "" {
		filterName := bson.M{
			"title": bson.M{
				"$regex":   name,
				"$options": "i",
			},
		}
		filterFields = append(filterFields, filterName)
	}

	if ingredient != "" {
		filterIngredient := bson.M{
			"ingredients.name": bson.M{
				"$regex":   ingredient,
				"$options": "i",
			},
		}
		filterFields = append(filterFields, filterIngredient)
	}

	var filter bson.M

	if len(filterFields) > 0 {
		filter = bson.M{
			"$or": filterFields,
		}
	}

	curso, err := repo.recipeCollection.Find(ctx, filter, &fOpt)
	if err != nil {
		return pagination, err
	}
	var list []entity.RecipeResponse
	for curso.Next(ctx) {
		var recipe Recipe
		if err := curso.Decode(&recipe); err != nil {
			fmt.Println(err)
		}

		list = append(list, recipe.ToEntityRecipeResponse())
	}
	pagination = api_helper.CreatePage(func() int {
		count, _ := repo.recipeCollection.CountDocuments(ctx, filter)
		return int(count)
	}, int(limit), page)
	pagination.Items = list
	return pagination, nil
}

func (repo repositoryImpl) CreateRecipe(recipe entity.Recipe) (entity.RecipeResponse, error) {
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
							return recipeModel.ToEntityRecipeResponse(), errors.New("duplicate " + fieldName)
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

	return recipeModel.ToEntityRecipeResponse(), nil
}

func (repo repositoryImpl) DeleteRecipe(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)

	_, err := repo.recipeCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}

func (repo repositoryImpl) GetRecipe(id string) (recipe entity.RecipeResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	var recipeModel Recipe
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	err = repo.recipeCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&recipeModel)
	return recipeModel.ToEntityRecipeResponse(), err
}

func (repo repositoryImpl) GetRecipesByUser(user string) (recipes []entity.RecipeResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	objectUserId, _ := primitive.ObjectIDFromHex(user)
	curso, err := repo.recipeCollection.Find(ctx, bson.M{"idUser": objectUserId})

	for curso.Next(ctx) {
		var recipeModel Recipe

		if err := curso.Decode(&recipeModel); err != nil {
			fmt.Println(err)
		}

		recipes = append(recipes, recipeModel.ToEntityRecipeResponse())
	}

	return
}

func (repo repositoryImpl) UpdateRecipe(recipe entity.Recipe) (entity.RecipeResponse, error) {
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

	return recipeModel.ToEntityRecipeResponse(), err
}

func New(mongoDb mongo.Database) repository.RecipeRepository {
	recipeCollection := mongoDb.Collection(_mongo.COLLETCTION_RECIPE)
	return &repositoryImpl{
		recipeCollection,
	}
}
