package comment_repository

import (
	"context"
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/entity"
	"fmt"
	"time"

	_mongo "duck-cook-recipe/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentRepositortImpl struct {
	commentCollection *mongo.Collection
}

func (repo commentRepositortImpl) GetCommentsByRecipe(idRecipe string) (comments []entity.CommentRecipe, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updatedAt": -1})
	curso, err := repo.commentCollection.Find(ctx, bson.M{"idRecipe": idRecipe}, findOptions)

	for curso.Next(ctx) {
		var commentModel Comment

		if err := curso.Decode(&commentModel); err != nil {
			fmt.Println(err)
		}

		comments = append(comments, commentModel.ToEntityComment())
	}

	return
}

func (repo commentRepositortImpl) CommentRecipeByUser(commentRecipe entity.CommentRecipe) (entity.CommentRecipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	var commentModel Comment
	commentModel = commentModel.FromEntity(commentRecipe)
	timeNow := time.Now()
	commentModel.CreatedAt = timeNow
	commentModel.UpdatedAt = timeNow

	res, err := repo.commentCollection.InsertOne(ctx, &commentModel)

	if err != nil {
		return commentRecipe, err
	}
	commentModel.ID = res.InsertedID.(primitive.ObjectID)

	return commentModel.ToEntityComment(), nil
}

func (repo commentRepositortImpl) DeleteCommentRecipeByUser(commentRecipe entity.CommentRecipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	_, err := repo.commentCollection.DeleteOne(ctx, bson.M{"_id": commentRecipe.IdComment})
	return err
}

func New(mongoDb mongo.Database) repository.CommentRecipeRepository {
	commentCollection := mongoDb.Collection(_mongo.COLLETCTION_COMMENT_RECIPE)
	return &commentRepositortImpl{
		commentCollection,
	}
}
