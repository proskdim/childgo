package config

import (
	"fmt"
	"os"
)

var (
	PORT             = getEnv("PORT", "8080")
	SECRET_KEY       = getEnv("SECRET_KEY", "key")
	CONTEXT_KEY_USER = getEnv("KEY_USER", "user")
	CONTEXT_JWT_USER = getEnv("KEY_JWT", "jwt_user")
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
