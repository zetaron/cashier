package main // import "github.com/zetaron/cashier"

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/goji/param"
	"github.com/google/go-querystring/query"
	"github.com/spf13/viper"
)

type AuthorizationRequest struct {
	ClientID    string `url:"client_id" param:"client_id"`
	RedirectUri string `url:"redirect_uri,omitempty" param:"redirect_uri"`
	Scope       string `url:"scope,omitempty" param:"scope"`
	State       string `url:"state,omitempty" param:"state"`
	AllowSignup bool   `url:"allow_signup,omitempty" param:"allow_signup"`
}

func (r *AuthorizationRequest) EncodeUri(baseUri string) string {
	v, _ := query.Values(r)

	buffer := bytes.NewBufferString(baseUri)
	buffer.WriteString("?")
	buffer.WriteString(v.Encode())

	return buffer.String()
}

type AccessTokenRequest struct {
	ClientID     string `url:"client_id" param:"client_id"`
	ClientSecret string `url:"client_secret" param:"client_secret"`
	Code         string `url:"code" param:"code"`
	RedirectUri  string `url:"redirect_uri,omitempty" param:"redirect_uri"`
	State        string `url:"state,omitempty" param:"state"`
}

func (r *AccessTokenRequest) EncodeUri(baseUri string) string {
	v, _ := query.Values(r)

	buffer := bytes.NewBufferString(baseUri)
	buffer.WriteString("?")
	buffer.WriteString(v.Encode())

	return buffer.String()
}

func AccessControlAllowOrigin(origin string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", origin)
		fn(w, r)
	}
}

func mustBeGET(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fn(w, r)
		} else {
			http.NotFound(w, r)
		}
	}
}

func mustBePOST(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fn(w, r)
		} else {
			http.NotFound(w, r)
		}
	}
}

func decodeAccessTokenRequest(r *http.Request) (AccessTokenRequest, error) {
	access_token_request := AccessTokenRequest{}

	err := param.Parse(r.URL.Query(), &access_token_request)
	if err != nil {
		fmt.Println("decodeAccessTokenRequest:", err)
		return access_token_request, err
	}

	return access_token_request, nil
}

func decodeAuthorizationRequest(r *http.Request) (AuthorizationRequest, error) {
	authorization_request := AuthorizationRequest{}

	err := param.Parse(r.URL.Query(), &authorization_request)
	if err != nil {
		fmt.Println("decodeAuthorizationRequest:", err)
		return authorization_request, err
	}

	return authorization_request, nil
}

func exchangeCodeForAccessToken(w http.ResponseWriter, r *http.Request) {
	access_token_request, err := decodeAccessTokenRequest(r)
	if err != nil || len(access_token_request.Code) < 1 {
		fmt.Println("error while decoding the access token request: ", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	access_token_request.ClientID = viper.GetString("client_id")
	access_token_request.ClientSecret = viper.GetString("client_secret")

	token_uri := viper.GetString("access_token_uri")
	get_token_url := access_token_request.EncodeUri(token_uri)

	client := http.Client{}
	req, err := http.NewRequest("POST", get_token_url, nil)
	req.Header.Add("Accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("error while being the proxy", err)
		w.WriteHeader(http.StatusFailedDependency)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseBody, err := ioutil.ReadAll(response.Body)
	w.Write(responseBody)

}

func redirectToAuthorization(w http.ResponseWriter, r *http.Request) {
	authorization_request, err := decodeAuthorizationRequest(r)
	if err != nil {
		fmt.Println("error while decoding the authorization request", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	authorization_request.ClientID = viper.GetString("client_id")

	auth_uri := viper.GetString("authorization_uri")
	http.Redirect(w, r, authorization_request.EncodeUri(auth_uri), 301)
}

func main() {
	viper.SetDefault("access_token_uri", "https://github.com/login/oauth/access_token")
	viper.SetDefault("authorization_uri", "https://github.com/login/oauth/authorize")
	viper.SetDefault("port", 80)
	viper.SetDefault("allowed_origin", "*")

	viper.AutomaticEnv()

	if !viper.IsSet("client_id") {
		log.Fatal("No ClientID defined.")
		return
	}

	if !viper.IsSet("client_secret") {
		log.Fatal("No ClientSecret defined.")
		return
	}

	http.HandleFunc(
		"/login/oauth/access_token",
		mustBePOST(
			AccessControlAllowOrigin(
				viper.GetString("allowed_origin"),
				exchangeCodeForAccessToken,
			),
		),
	)

	http.HandleFunc(
		"/login/oauth/authorize",
		mustBeGET(redirectToAuthorization),
	)

	log.Infof("Starting the cashier on Port %s", viper.GetString("port"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("port")), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
