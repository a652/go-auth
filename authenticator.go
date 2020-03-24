package auth

import (
	"fmt"

	"github.com/a652/go-auth/request"
	"github.com/a652/go-auth/store"
	"github.com/a652/go-auth/token"
)

type APIAuthenticator struct {
	credentialStore store.CredentialStore
}

func NewAPIAuthenticator(credentialStore store.CredentialStore) *APIAuthenticator {
	return &APIAuthenticator{
		credentialStore,
	}
}

func (aauth *APIAuthenticator) Auth(rawURL string) bool {
	req, err := request.NewAPIRequestFromRawUrl(rawURL)
	if err != nil {
		fmt.Println("create api request from raw url err: " + err.Error())
		return false
	}
	return aauth.auth(req)
}

func (aauth *APIAuthenticator) auth(req *request.APIRequest) bool {

	clientToken := token.NewAuthToken(req.GetToken(), req.GetTS(), token.WithDefaultExpiredTimeInterval())
	if clientToken.IsExpired() {
		fmt.Println("token is expired")
		return false
	}
	secret, err := aauth.credentialStore.GetSecretByAppID(req.GetAppID())
	if err != nil {
		fmt.Println("get secret err")
		return false
	}
	fmt.Printf("path: %s, appid: %s, secret: %s, ts: %d\n", req.GetBaseUrl(), req.GetAppID(), secret, req.GetTS())
	serverToken := token.GenerateAuthToken(req.GetBaseUrl(), req.GetAppID(), secret, req.GetTS())
	return serverToken.Match(clientToken)
}
