package provider

import (
	"github.com/wangjiaxi90/terraform-provider-scaffolding/internal/provider/cloudmgr"
)

type Config struct {
	AccessKey string
	SecretKey string
	Zone      string
	EndPoint  string
}

func (c *Config) Client() (interface{}, error) {
	cfg := cloudmgr.NewConfiguration()
	cfg.Host = "host" //TODO
	cfg.Scheme = "scheme"
	return cloudmgr.NewAPIClient(cfg), nil
}

func Int32(v int) *int32 {
	t := int32(v)
	return &t
}
func String(v string) *string {
	return &v
}
func Bool(v bool) *bool {
	return &v
}
