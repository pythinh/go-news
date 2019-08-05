package types

import "time"

// Server config
type Server struct {
	DB struct {
		Type     string `env:"DB_TYPE" default:"mongodb"`
		ConfigDB Database
	}
	HTTP struct {
		Address           string        `env:"HTTP_ADDRESS" default:""`
		Port              int           `env:"PORT" default:"8080"`
		ReadTimeout       time.Duration `env:"HTTP_READ_TIMEOUT" default:"5m"`
		WriteTimeout      time.Duration `env:"HTTP_WRITE_TIMEOUT" default:"5m"`
		ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT" default:"30s"`
		ShutdownTimeout   time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT" default:"10s"`
	}
}
