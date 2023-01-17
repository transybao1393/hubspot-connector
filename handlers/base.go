package handlers

import (
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func New() http.Handler {

	viper.SetConfigFile(".env")

	//- declare env
	//- This not work, please move to file
	os.Setenv("GOOGLE_OAUTH_CLIENT_ID", "1")
	os.Setenv("GOOGLE_OAUTH_CLIENT_SECRET", "1")
	os.Setenv("HUBSPOT_OAUTH_CLIENT_ID", "11ce62ba-d0f0-400a-a90c-b7d8f1bf5d6f")
	os.Setenv("HUBSPOT_OAUTH_CLIENT_SECRET", "e754f12e-a4da-474a-aabc-108bf6c3af9b")

	mux := http.NewServeMux()
	// Root
	mux.Handle("/",  http.FileServer(http.Dir("templates/")))

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	//- OAuth Hubspot
	mux.HandleFunc("/auth/hubspot/login", oauthHubspotLogin)
	mux.HandleFunc("/auth/hubspot/callback", oauthHubspotCallback)

	return mux
}
