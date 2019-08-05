package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Connections all types of database connections
type Connections struct {
	Type    string
	MongoDB *mongo.Client
}

const (
	// TypeMongoDB type of mongodb
	TypeMongoDB = "mongodb"
	// TypeMySQL type of mysql
	TypeMySQL = "mysql"
)

// Close all underlying connections
func (c *Connections) Close() error {
	switch c.Type {
	case TypeMongoDB:
		if c.MongoDB != nil {
			err := c.MongoDB.Disconnect(context.Background())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
