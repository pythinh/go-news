package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/pythinh/go-news/internal/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Dial to target server MongoDB
func Dial(conf *types.Database) (*mongo.Client, error) {
	log.Printf("dialing to target MongoDB at: %s, database: %s", conf.Host, conf.Name)
	mongoURL := fmt.Sprintf("mongodb://%s/%s", conf.Host, conf.Name)
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("successfully dialing to MongoDB at:", conf.Host)
	return client, nil
}
