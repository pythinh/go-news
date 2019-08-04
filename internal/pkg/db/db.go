package db

import (
	"context"

	"github.com/pythinh/go-news/internal/app/router"
	"github.com/pythinh/go-news/internal/app/types"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	typeMongoDB = "mongodb"
	typeMySQL   = "mysql"
)

// Init connections to database
func Init(conf *types.Server) *router.DBConns {
	conns := &router.DBConns{}
	conns.Database.Type = conf.DB.Type

	switch conf.DB.Type {
	case db.typeMongoDB:

	}
}

// Close all underlying connections
func (c *Connections) Close() error {
	switch c.Type {
	case typeMongoDB:
		if c.MongoDB != nil {
			err := c.MongoDB.Disconnect(context.Background())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
