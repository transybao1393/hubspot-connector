package handlers

import (
	"net/http"

	"github.com/spf13/viper"
)

func New() http.Handler {

	viper.SetConfigFile(".env")

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
