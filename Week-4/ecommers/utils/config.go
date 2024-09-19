package utils

import (
	"os"
)

var (
	ClientID     = os.Getenv("GOOGLE_CLIENT_ID")
	ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	RedirectURL  = os.Getenv("GOOGLE_REDIRECT_URL")
)
