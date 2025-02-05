package config

import (
	"fmt"
	"os"
)

var (
	Port           = getEnv("PORT", "8080")
	SecretKey      = getEnv("SECRET_KEY", "key")
	ContextKeyUser = getEnv("KEY_USER", "user")
	ContextJwtUser = getEnv("KEY_JWT", "jwt_user")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
