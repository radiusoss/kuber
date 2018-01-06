package twitter

import (
	"fmt"
	"log"

	"github.com/mrjones/oauth"
)

func NewTwitterSession(consumerKey, consumerSecret string) *TwitterSession {

	newServer := new(TwitterSession)

	newServer.OAuthConsumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   OAUTH_REQUEST_TOKEN,
			AuthorizeTokenUrl: OAUTH_AUTH_TOKEN,
			AccessTokenUrl:    OAUTH_ACCESS_TOKEN,
		},
	)

	//Enable debug info
	newServer.OAuthConsumer.Debug(false)

	// newServer.Client = *newClient
	newServer.OAuthTokens = make(map[string]*oauth.RequestToken)
	return newServer
}

type TwitterSession struct {
	Client
	OAuthConsumer *oauth.Consumer
	OAuthTokens   map[string]*oauth.RequestToken
}

func (s *TwitterSession) GetAuthURL(tokenUrl string) string {
	token, requestUrl, err := s.OAuthConsumer.GetRequestTokenAndUrl(tokenUrl)
	if err != nil {
		log.Println(err)
	}
	// Make sure to save the token, we'll need it for AuthorizeToken()
	s.OAuthTokens[token.Token] = token
	return requestUrl
}

func (s *TwitterSession) CompleteAuth(tokenKey, verificationCode string) error {
	accessToken, err := s.OAuthConsumer.AuthorizeToken(s.OAuthTokens[tokenKey], verificationCode)
	if err != nil {
		log.Println(err)
	}
	s.HttpConn, err = s.OAuthConsumer.MakeHttpClient(accessToken)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Access Token: ", accessToken)
	return nil
}
