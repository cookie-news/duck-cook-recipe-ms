package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLETCTION_RECIPE         = "Recipe"
	COLLETCTION_COMMENT_RECIPE = "CommentRecipe"
)

func Connect() mongo.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_SRV")).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	db := mongoClient.Database(os.Getenv("MONGO_DB"))

	return *db
}
