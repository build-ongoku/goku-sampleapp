package svcauth_typ

import (
	"github.com/teejays/gokutil/scalars"

	app_typ "sampleapp/backend/.goku/generated/typ"
)

// AuthenticateResponse: <comments>
type AuthenticateResponse struct {
	Token string `json:"token" validate:"required"`
}

// AuthenticateTokenRequest: <comments>
type AuthenticateTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

// LoginRequest: <comments>
type LoginRequest struct {
	Email    scalars.Email `json:"email" validate:"required"`
	Password string        `json:"password" validate:"required"`
}

// RegisterUserRequest: <comments>
type RegisterUserRequest struct {
	Email    scalars.Email           `json:"email" validate:"required"`
	Name     app_typ.PersonNameInput `json:"name" validate:"required"`
	Password string                  `json:"password" validate:"required"`
	TeamID   scalars.ID              `json:"teamID" validate:"required,uuid"`
}

// AuthenticateResponse:
// AuthenticateTokenRequest:
// LoginRequest:
// RegisterUserRequest:
