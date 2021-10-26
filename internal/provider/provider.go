package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"username": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: descriptions["username"],
				},
				"password": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: descriptions["password"],
				},
				"client_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: descriptions["client_id"],
				},
				"client_secret": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: descriptions["client_secret"],
				},
				"endpoint": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: descriptions["endpoint"],
				},
			},
			DataSourcesMap: map[string]*schema.Resource{ //这个在hcl文件中用data字段获取
				"hashdata_data_source": dataSourceEmpty(),
			},
			ResourcesMap: map[string]*schema.Resource{ //这个在hcl文件中用resource字段获取
				"hashdata_warehouse": resourceWarehouse(),
				"hashdata_catalog":   resourceCatalog(),
				"hashdata_computing": resourceComputing(),
			},
		}
		p.ConfigureContextFunc = configure
		return p
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	username, ok := d.GetOk("username")
	if !ok {
		username = os.Getenv("HASHDATA_USERNAME")
		if username == "" {
			return nil, diag.Errorf("Don't have username variable in provider scope.")
		}
	}
	password, ok := d.GetOk("password")
	if !ok {
		password = os.Getenv("HASHDATA_PASSWORD")
		if password == "" {
			return nil, diag.Errorf("Don't have password variable in provider scope.")
		}
	}
	client_id, ok := d.GetOk("client_id")
	if !ok {
		client_id = os.Getenv("HASHDATA_CLIENT_ID")
		if client_id == "" {
			return nil, diag.Errorf("Don't have client_id variable in provider scope.")
		}
	}
	client_secret, ok := d.GetOk("client_secret")
	if !ok {
		client_secret = os.Getenv("HASHDATA_CLIENT_SECRET")
		if client_secret == "" {
			return nil, diag.Errorf("Don't have client_secret variable in provider scope.")
		}
	}
	endpoint, ok := d.GetOk("endpoint")
	if !ok {
		endpoint = os.Getenv("HASHDATA_END_POINT")
		if endpoint == "" {
			endpoint = DEFAULT_ENDPOINT
		}
	}
	config := Config{
		Username:     username.(string),
		Password:     password.(string),
		ClientId:     client_id.(string),
		ClientSecret: client_secret.(string),
		EndPoint:     endpoint.(string),
	}
	client, err := config.Client() // TODO 校验参数啥的
	if err != nil {
		return nil, diag.Errorf(err.Error())
	}
	return client, nil
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"username":      "hashdata username ",
		"password":      "hashdata password",
		"client_id":     "hashdata client_id",
		"client_secret": "hashdata client_secret",
		"endpoint":      "hashdata endpoint",
	}
}
