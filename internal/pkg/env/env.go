package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Load the environment variables into struct
func Load(t interface{}) {
	err := envconfig.Process("", t)
	if err != nil {
		log.Panicf("unable to load config for %T: %s", t, err)
	}
}
