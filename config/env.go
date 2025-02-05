package config

import "childgo/utils/env"

var (
	Port           = env.Fetch("PORT", "8080")
	SecretKey      = env.Fetch("SECRET_KEY", "key")
	ContextKeyUser = env.Fetch("KEY_USER", "user")
	ContextJwtUser = env.Fetch("KEY_JWT", "jwt_user")
)
