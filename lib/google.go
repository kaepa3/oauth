package google

import (
	"github.com/BurntSushi/toml"
	"golang.org/x/oauth2"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

var config Config

type Config struct {
	GoogleClientID     string
	GoogleClientSecret string
}

func init() {
	toml.DecodeFile("./config.toml", &config)
}

// GetConnect 接続を取得する
func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     config.GoogleClientID,
		ClientSecret: config.GoogleClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: "http://localhost:8000/callback",
	}

	return config
}
