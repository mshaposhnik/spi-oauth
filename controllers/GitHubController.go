package controllers

import (
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"strings"
	"context"
	"fmt"
	"io/ioutil"
)

var gitHubConf = &oauth2.Config{
	ClientID:     "<edited>",
	ClientSecret: "<edited>",
	RedirectURL:  "http://localhost:8000/github/callback",
	Endpoint: github.Endpoint,
}


const gitHubUserAPI = "https://api.github.com/user?access_token="



var GitHubAuthenticate = func(w http.ResponseWriter, r *http.Request) {

	scopes := r.FormValue("scopes")
	gitHubConf.Scopes = strings.Split(scopes, ",")

	state := r.FormValue("state")
	url := gitHubConf.AuthCodeURL(state)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}




var GitHubCallback = func(w http.ResponseWriter, r *http.Request) {

	//state := r.FormValue("state");
	//TODO: validate state
	code := r.FormValue("code")

	token, err := gitHubConf.Exchange(context.Background(), code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error in GitHub token exchange: %s",err.Error())
		return
	}


	req, err := http.NewRequest("GET", gitHubUserAPI, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed making GitHub request: %s", err.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer " + token.AccessToken)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Failed getting GitHub user: %s", err.Error())
		return
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Failed pasring GitHub user data: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Oauth Token: %s <br/>", token.AccessToken)
	fmt.Fprintf(w, "User data: %s", string(content))

}
