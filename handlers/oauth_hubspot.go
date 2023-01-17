package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var hubspotOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/hubspot/callback",
	ClientID:     "",
	ClientSecret: "",
	Scopes:       []string{"crm.objects.companies.read", "crm.objects.contacts.read"},
	Endpoint:     oauth2.Endpoint {
		AuthURL: "https://app.hubspot.com/oauth/authorize",
		TokenURL: "https://api.hubspot.com/oauth/v1/token",
		AuthStyle: 1,
	},
}

// const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

//- DEPRECATED
func oauthHubspotLogin(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
	AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	validate that it matches the the state query parameter on your redirect callback.
	*/
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

//- https://app.hubspot.com/oauth/authorize?client_id=11ce62ba-d0f0-400a-a90c-b7d8f1bf5d6f&redirect_uri=http://localhost:8000/auth/hubspot/callback&scope=crm.objects.contacts.read%20crm.objects.companies.read
func oauthHubspotCallback(w http.ResponseWriter, r *http.Request) {
	viper.ReadInConfig()
	// Read oauthState from Cookie
	// oauthState, _ := r.Cookie("oauthstate")

	// if r.FormValue("state") != oauthState.Value {
	// 	log.Println("invalid oauth google state")
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }

	code := r.FormValue("code")
	fmt.Printf("code in request parameter %s", code)
	
	hubspotOauthConfig.ClientID = viper.GetString("HUBSPOT_OAUTH_CLIENT_ID")
	hubspotOauthConfig.ClientSecret = viper.GetString("HUBSPOT_OAUTH_CLIENT_SECRET")
	

	//- get code
	tokens, err := hubspotOauthConfig.Exchange(context.Background(), code)
	fmt.Println("=========")
	if err != nil {
		fmt.Printf("code exchange wrong: %s", err.Error())
		return 
	}

	//- save token to somewhere to reuse
	//- TODO: Check token expiration
	//- TODO: Save tokens to database and check token expiration
	fmt.Println(tokens)

	// data, err := getUserDataFromHubspot(r.FormValue("code"))
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(w, "UserInfo: %s\n", "")
}

// func generateStateOauthCookie(w http.ResponseWriter) string {
// 	var expiration = time.Now().Add(20 * time.Minute)

// 	b := make([]byte, 16)
// 	rand.Read(b)
// 	state := base64.URLEncoding.EncodeToString(b)
// 	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
// 	http.SetCookie(w, &cookie)

// 	return state
// }

func getUserDataFromHubspot(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := hubspotOauthConfig.Exchange(context.Background(), code)
	fmt.Println(token)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	// response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	// }

	// defer response.Body.Close()
	// contents, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed read response: %s", err.Error())
	// }
	// return contents, nil
	return nil, nil
}