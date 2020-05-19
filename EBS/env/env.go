package env

import (
	"os"
)

// extension settings
var (
	ClientId           = os.Getenv("CLIENT_ID")
	ClientSecret       = os.Getenv("CLIENT_SECRET")
	ExtSecret          = os.Getenv("EXT_SECRET")
	OwnerID            = os.Getenv("OWNER_ID")
	ExtConfigVersion   = os.Getenv("EXT_CONFIG_VERSION")
	ExtVersion         = os.Getenv("EXT_VERSION")
	DebugLevel         = os.Getenv("DEBUG_LEVEL")
	AppAccessToken     = os.Getenv("APP_ACCESS_TOKEN")
	TwitchAuthRedirect = os.Getenv("AUTH_URL")
)

// aws resources
var (
	UsersTable = os.Getenv("USERS_TBL")
)
