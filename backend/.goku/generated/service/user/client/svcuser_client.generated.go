package svcuser_client

import (
	"github.com/teejays/gokutil/log"

	svcuser_entteam_client "sampleapp/backend/.goku/generated/service/user/entity/team/client"
	svcuser_entuser_client "sampleapp/backend/.goku/generated/service/user/entity/user/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcuser")

// Service Client

// Client is the interface for any User service client.
// A service is made up of it's constitient entity clients, and any service level custom methods.
type Client interface {
	EntitiesClient
	CustomClient
}

// EntitiesClient provides access to the clients of entities in this service.
type EntitiesClient interface {
	Team() svcuser_entteam_client.Client
	User() svcuser_entuser_client.Client
}

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
}
