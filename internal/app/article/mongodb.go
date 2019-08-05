package article

import (
	"context"

	"github.com/pythinh/go-news/internal/pkg/types"
	"github.com/pythinh/go-news/internal/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	client *mongo.Client
}

func newMongoRepository(c *mongo.Client) *mongoRepository {
	return &mongoRepository{c}
}

func (r *mongoRepository) collection(c *mongo.Client) *mongo.Collection {
	return c.Database("gonews").Collection("article")
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*types.Article, error) {
	var article *types.Article
	collection := r.collection(r.client)
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(ctx, filter).Decode(&article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (r *mongoRepository) FindAll(ctx context.Context) ([]types.Article, error) {
	var article []types.Article
	collection := r.collection(r.client)
	cur, err := collection.Find(ctx, bson.D{{}})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var elem types.Article
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		article = append(article, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return article, nil
}

func (r *mongoRepository) Create(ctx context.Context, article types.Article) (string, error) {
	article.ID = uuid.New()
	collection := r.collection(r.client)
	_, err := collection.InsertOne(ctx, article)
	if err != nil {
		return "", err
	}
	return article.ID, nil
}

func (r *mongoRepository) Update(ctx context.Context, article types.Article) error {
	collection := r.collection(r.client)
	filter := bson.D{{Key: "_id", Value: article.ID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: article.Title},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) Delete(ctx context.Context, id string) error {
	collection := r.collection(r.client)
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := collection.DeleteOne(ctx, filter, nil)
	return err
}
