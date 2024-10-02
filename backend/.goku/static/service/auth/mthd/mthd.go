package custom_svcauth_mthd

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/teejays/gokutil/errutil"
	"github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/panics"

	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
	svcauth_typ "sampleapp/backend/.goku/generated/service/auth/typ"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
	"sampleapp/backend/.goku/static/service/auth/auth"
)

func LoginUser(ctx context.Context, c app_client.Client, req svcauth_typ.LoginRequest) (svcauth_typ.AuthenticateResponse, error) {
	var resp svcauth_typ.AuthenticateResponse

	// Get the DAL wrapper
	listUsersResp, err := c.User().User().List(ctx, svcuser_entuser_typ.ListRequest{
		Filter: svcuser_entuser_typ.UserFilter{
			Email: filter.NewEmailCondition(filter.EQUAL, req.Email),
		},
	})
	if err != nil {
		return resp, fmt.Errorf("Fetching User with email [%s]: %w", req.Email, err)
	}
	if len(listUsersResp.Items) < 1 {
		err := fmt.Errorf("User: %w", errutil.ErrNotFound)
		return resp, errutil.WrapGerror(err, http.StatusUnauthorized, "Invalid email and/or password")
	}
	panics.If(len(listUsersResp.Items) > 1, "Multiple (%d) users found with the same email: %s", len(listUsersResp.Items), req.Email)

	user := listUsersResp.Items[0]

	// Fetch user auth credentials
	authCreds, err := c.Auth().Secret().List(ctx, svcauth_entsecret_typ.ListRequest{
		Filter: svcauth_entsecret_typ.SecretFilter{
			UserID: filter.NewIDCondition(filter.EQUAL, user.ID),
			Type:   &svcauth_entsecret_typ.TypeCondition{Op: filter.EQUAL, Values: []svcauth_entsecret_typ.Type{svcauth_entsecret_typ.Type_Password}},
		},
	})
	if err != nil {
		return resp, fmt.Errorf("Fetching user auth credentials: %w", err)
	}
	hashedPassword := authCreds.Items[0].Secret

	// Compare the hashed password with the submitted password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return resp, errutil.NewGerror(http.StatusUnauthorized, "Invalid email and/or password", "submitted & hashed password did not match")
	}
	if err != nil {
		return resp, fmt.Errorf("comparing hashed password: %w", err)
	}

	// Create and generate a token
	token, err := auth.CreateTokenForUser(ctx, user)
	if err != nil {
		return resp, err
	}

	resp.Token = token

	return resp, err
}

func RegisterUser(ctx context.Context, c app_client.Client, req svcauth_typ.RegisterUserRequest) (svcauth_typ.AuthenticateResponse, error) {
	var resp svcauth_typ.AuthenticateResponse
	var err error

	// Make sure we have no user with the same email
	existingUsers, err := c.User().User().List(ctx, svcuser_entuser_typ.ListRequest{
		Filter: svcuser_entuser_typ.UserFilter{
			Email: filter.NewEmailCondition(filter.EQUAL, req.Email),
		},
	})
	if err != nil {
		return resp, fmt.Errorf("Checking if email is already in use: %w", err)
	}
	if len(existingUsers.Items) > 0 {
		return resp, fmt.Errorf("Email is already in use")
	}

	// The hashed password from bcrypt already has the salt embedded in it, so we don't need to store it separately
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return resp, fmt.Errorf("Could not handle the password: %w", err)
	}

	var _u = svcuser_entuser_typ.UserInput{
		Email:  req.Email,
		Name:   req.Name,
		TeamID: req.TeamID,
	}

	u, err := c.User().User().Add(ctx, svcuser_entuser_typ.AddRequest{Object: _u})
	if err != nil {
		return resp, fmt.Errorf("Cannot create user: %w", err)
	}

	// Add user credentials
	_scrt := svcauth_entsecret_typ.SecretInput{
		UserID: u.ID,
		Secret: string(hashed),
		Type:   svcauth_entsecret_typ.Type_Password,
	}
	_, err = c.Auth().Secret().Add(ctx, svcauth_entsecret_typ.AddRequest{Object: _scrt})
	if err != nil {
		return resp, fmt.Errorf("Cannot create user secret: %w", err)
	}

	resp, err = LoginUser(ctx, c, svcauth_typ.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return resp, err
	}

	return resp, err
}

func AuthenticateToken(ctx context.Context, c app_client.Client, req svcauth_typ.AuthenticateTokenRequest) (svcauth_typ.AuthenticateResponse, error) {
	var resp svcauth_typ.AuthenticateResponse
	var err error

	// Create and generate a token
	_, err = auth.AuthenticateTokenWithDefaultClient(ctx, req.Token)
	if err != nil {
		return resp, err
	}

	resp.Token = req.Token

	return resp, nil
}
