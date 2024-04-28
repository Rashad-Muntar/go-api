package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadInitializers() {
	env := os.Getenv("ENVIRONMENT")

	if env == "" {
		env = "develop"
	  }

	  godotenv.Load(".env." + env + ".local")
	  if env != "test" {
		godotenv.Load(".env.local")
	  }
	  godotenv.Load(".env." + env)
	  godotenv.Load()
}