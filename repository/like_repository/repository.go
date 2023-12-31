package like_repository

import (
	"context"
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/entity"
	"time"

	_mongo "duck-cook-recipe/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type likeRepositortImpl struct {
	likeCollection *mongo.Collection
}

func (repo likeRepositortImpl) CheckRecipeIsLikedByUser(like entity.LikeRecipe) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	var likeModel Like
	likeModel = likeModel.FromEntity(like)

	count, err := repo.likeCollection.CountDocuments(ctx, bson.M{"idRecipe": likeModel.IdRecipe, "idUser": likeModel.IdUser})

	return count > 0, err
}

func (repo likeRepositortImpl) DeleteLikeRecipeByUser(like entity.LikeRecipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	objectIdRecipe, err := primitive.ObjectIDFromHex(like.Id)
	if err != nil {
		return err
	}
	objectIdUser, err := primitive.ObjectIDFromHex(like.IdUser)
	if err != nil {
		return err
	}
	_, err = repo.likeCollection.DeleteOne(ctx, bson.M{"idRecipe": objectIdRecipe, "idUser": objectIdUser})
	return err
}

func (repo likeRepositortImpl) GetLikesByRecipe(idRecipe string) (count int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	objectIdRecipe, err := primitive.ObjectIDFromHex(idRecipe)
	if err != nil {
		return
	}
	count, err = repo.likeCollection.CountDocuments(ctx, bson.M{"idRecipe": objectIdRecipe})
	return
}

func (repo likeRepositortImpl) LikeRecipeByUser(like entity.LikeRecipe) (entity.LikeRecipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	var likeModel Like
	likeModel = likeModel.FromEntity(like)
	timeNow := time.Now()
	likeModel.CreatedAt = timeNow
	likeModel.UpdatedAt = timeNow

	res, err := repo.likeCollection.InsertOne(ctx, &likeModel)

	if err != nil {
		return like, err
	}
	likeModel.ID = res.InsertedID.(primitive.ObjectID)

	return likeModel.ToEntityLike(), nil
}

func New(mongoDb mongo.Database) repository.LikeRecipeRepository {
	commentCollection := mongoDb.Collection(_mongo.COLLETCTION_LIKE_RECIPE)
	return &likeRepositortImpl{
		commentCollection,
	}
}
