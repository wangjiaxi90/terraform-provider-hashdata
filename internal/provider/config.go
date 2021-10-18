package provider

import (
	"context"
	"fmt"
	"github.com/wangjiaxi90/terraform-provider-scaffolding/internal/provider/cloudmgr"
	"golang.org/x/oauth2"
	"net/url"
	"strings"
)

type Config struct {
	Username     string
	Password     string
	ClientId     string
	ClientSecret string
	EndPoint     string
}

func (c *Config) Client() (interface{}, error) {
	hashdataURL, err := url.Parse(c.EndPoint)
	if err != nil {
		return nil, err
	}
	cfg := cloudmgr.NewConfiguration()
	config_outh2 := oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: c.EndPoint,
		},
	}
	token, err := config_outh2.PasswordCredentialsToken(context.Background(), c.Username, c.Password)
	if err != nil {
		return nil, err
	}

	//TODO 在这里继续

	if !strings.Contains(c.EndPoint, ":") {
		if strings.Contains(c.EndPoint, "https") {
			c.EndPoint = c.EndPoint + ":443"
		} else if strings.Contains(c.EndPoint, "http") {
			c.EndPoint = c.EndPoint + ":80"
		} else {
			return nil, fmt.Errorf("can not find default port ")
		}
	}
	cfg.Host = c.EndPoint
	var header = make(map[string]string)
	header["Authorization"] = AUTH_PREFIX + token.AccessToken
	cfg.DefaultHeader = header
	return cloudmgr.NewAPIClient(cfg), nil
}
