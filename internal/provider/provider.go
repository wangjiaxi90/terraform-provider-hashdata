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
			DataSourcesMap: map[string]*schema.Resource{ //这个在hcl文件中用data字段获取
				"scaffolding_data_source": dataSourceEmpty(),
			},
			ResourcesMap: map[string]*schema.Resource{ //这个在hcl文件中用resource字段获取
				"scaffolding_warehouse": resourceWarehouse(),
				"scaffolding_catalog":   resourceCatalog(),
				"scaffolding_computing": resourceComputing(),
			},
		}
		p.ConfigureContextFunc = configure
		return p
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accesskey, ok := d.GetOk("access_key")
	if !ok {
		accesskey = os.Getenv("HASHDATA_ACCESS_KEY")
	}
	secretkey, ok := d.GetOk("secret_key")
	if !ok {
		secretkey = os.Getenv("HASHDATA_SECRET_KEY")
	}
	zone, ok := d.GetOk("zone")
	if !ok {
		zone = os.Getenv("HASHDATA_ZONE")
		if zone == "" {
			zone = "DEFAULT_ZONE" //TODO 默认区
		}
	}
	endpoint, ok := d.GetOk("endpoint")
	if !ok {
		endpoint = os.Getenv("HASHDATA_ENDPOINT")
		if endpoint == "" {
			endpoint = "DEFAULT_ENDPOINT" //TODO 默认endpoint
		}
	}
	config := Config{
		AccessKey: accesskey.(string),
		SecretKey: secretkey.(string),
		Zone:      accesskey.(string),
		EndPoint:  endpoint.(string),
	}
	client, err := config.Client() // TODO 校验参数啥的
	if err != nil {
		return nil, diag.Errorf(err.Error())
	}
	return client, nil
}
