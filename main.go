package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/zipkero/use-gin-in-golang/config"
	"github.com/zipkero/use-gin-in-golang/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection
var recipesHandler *handlers.RecipeHandler

func init() {
	ctx = context.Background()
	mongodbConfigGetter := config.MongodbConfig{Name: "mongodb"}
	redisConfigGetter := config.RedisConfig{Name: "redis"}
	mongodbConfig, _ := mongodbConfigGetter.Get()
	redisConfig, _ := redisConfigGetter.Get()

	log.Println(mongodbConfig)
	log.Println(redisConfig)
	client, err = mongo.Connect(ctx,
		options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@%s:%d", mongodbConfig.Username, mongodbConfig.Password, mongodbConfig.Host, mongodbConfig.Port),
		),
	)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s: %d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0,
	})
	status := redisClient.Ping(ctx)
	fmt.Println(status)

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("connected mongodbConfig")
	collection = client.Database("SAMPLE").Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)

}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.CreateRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipeHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.Run()
}
