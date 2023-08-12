package sdk

import (
	"os"

	"github.com/joho/godotenv"
)

type Interface interface {
	CanLoad(envPath string) error
	Get(key string) string
}

type config struct{}

func Init() Interface {
	return &config{}
}

func (c *config) CanLoad(envPath string) error {
	return godotenv.Load(envPath)
}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}
