package user

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
	return c.Database("gonews").Collection("user")
}

func (r *mongoRepository) FindByUsername(ctx context.Context, username string) (*types.User, error) {
	var user *types.User
	collection := r.collection(r.client)
	filter := bson.D{{Key: "username", Value: username}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mongoRepository) CheckByUsername(ctx context.Context, username string) bool {
	var user *types.User
	collection := r.collection(r.client)
	filter := bson.D{{Key: "username", Value: username}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return false
	}
	return true
}

func (r *mongoRepository) FindAll(ctx context.Context) ([]types.User, error) {
	var user []types.User
	collection := r.collection(r.client)
	cur, err := collection.Find(ctx, bson.D{{}})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var elem types.User
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		user = append(user, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mongoRepository) Create(ctx context.Context, user *types.User) (string, error) {
	user.ID = uuid.New()
	user.Priority = 1
	collection := r.collection(r.client)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func (r *mongoRepository) Update(ctx context.Context, user *types.User) error {
	collection := r.collection(r.client)
	filter := bson.D{{Key: "username", Value: user.Username}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "fullname", Value: user.Fullname},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
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
