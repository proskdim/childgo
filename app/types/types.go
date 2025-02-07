package types

import (
	"time"

	"github.com/google/uuid"
)

type MsgResp struct {
	Message string `json:"message"`
}

type ChildRequest struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

type ChildResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

type Child struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

type SigninResponse struct {
	JWTToken string `json:"jwt_token"`
}

type SignupResponse struct {
	Email string `json: "email"`
}