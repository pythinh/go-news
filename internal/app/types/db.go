package types

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Connections all types of database connections
	Connections struct {
		Type    string
		MongoDB *mongo.Client
	}
	// Database config
	Database struct {
		Host    string        `env:"HOST_DB" default:"127.0.0.1:27017"`
		Name    string        `env:"DATABASE_DB" default:"gonews"`
		User    string        `env:"USERNAME_DB"`
		Pass    string        `env:"PASSWORD_DB"`
		Timeout time.Duration `env:"TIMEOUT_DB" default:"10s"`
	}
)
