package provider

import (
	"github.com/wangjiaxi90/terraform-provider-scaffolding/internal/provider/cloudmgr"
)

type Config struct {
	//AccessKey string
	//SecretKey string
	//Zone      string
	//EndPoint  string
}

func (c *Config) Client() (interface{}, error) {
	return cloudmgr.NewAPIClient(cloudmgr.NewConfiguration()), nil // TODO 九折九折？
}
