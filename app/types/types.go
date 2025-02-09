package types

import (
	"time"

	"github.com/google/uuid"
)

type MsgResp struct {
	Message string `json:"message"`
}

type HealthResp struct {
	Status string `json:"status"`
}

type ChildRequest struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`

	Address AddressRequest `json:"address"`
}

type ChildCreateResponse struct {
	Child *ChildResponse `json:"child"`
}

type ChildResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`

	Address AddressResponse `json:"address"`
}

type Child struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

type AddressRequest struct {
	City      string `json:"city"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Apartment string `json:"apartment"`
}

type AddressResponse struct {
	City      string `json:"city"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Apartment string `json:"apartment"`
}

type SigninResponse struct {
	JWTToken string `json:"jwt_token"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Email string `json:"email"`
}

type ProfileResponse struct {
	Email string `json:"email"`
}
