package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Configuration *configImpl

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) error {
	err := godotenv.Load(filenames...)
	if err != nil {
		return err
	}
	Configuration = &configImpl{}
	return nil
}
