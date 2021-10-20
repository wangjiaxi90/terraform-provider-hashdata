package provider

import (
	"context"
	"github.com/wangjiaxi90/terraform-provider-hashdata/internal/provider/cloudmgr"
	"golang.org/x/oauth2"
)

type Config struct {
	Username     string
	Password     string
	ClientId     string
	ClientSecret string
	EndPoint     string
}

func (c *Config) Client() (interface{}, error) {
	cfg := cloudmgr.NewConfiguration()
	configOauth2 := oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: c.EndPoint + TOKEN_PATH,
		},
	}
	token, err := configOauth2.PasswordCredentialsToken(context.Background(), c.Username, c.Password)
	if err != nil {
		return nil, err
	}

	cfg.Host = c.EndPoint
	var header = make(map[string]string)
	header["Authorization"] = AUTH_PREFIX + token.AccessToken
	cfg.DefaultHeader = header
	cfg.Servers = cloudmgr.ServerConfigurations{
		{
			URL:         c.EndPoint,
			Description: "Default server",
		},
	}
	return cloudmgr.NewAPIClient(cfg), nil
}
