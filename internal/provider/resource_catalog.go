package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wangjiaxi90/terraform-provider-scaffolding/internal/provider/cloudmgr"
	"os"
)

func resourceCatalog() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Catalog.",

		CreateContext: resourceCatalogCreate,
		ReadContext:   resourceCatalogRead,
		UpdateContext: resourceCatalogUpdate,
		DeleteContext: resourceCatalogDelete,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceCatalogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	body := *cloudmgr.NewCoreCreateCatalogRequest()

	configuration := cloudmgr.NewConfiguration() //TODO 客户端的生成
	apiClient := cloudmgr.NewAPIClient(configuration)

	etcdPropertiesRaw := d.Get("etcd").(*schema.Set).List()
	var etcdProperties map[string]interface{}
	for _, raw := range etcdPropertiesRaw {
		etcdProperties = raw.(map[string]interface{})
		break
	}
	etcd := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(etcdProperties["count"].(int)),
			InstanceType: String(etcdProperties["instance_type"].(string)),
			VolumeType:   String(etcdProperties["instance_type"].(string)),
			VolumeSize:   Int32(etcdProperties["instance_type"].(int)),
			Image:        String(etcdProperties["instance_type"].(string)),
			Zone:         String(etcdProperties["instance_type"].(string)),
		},
	}

	catalogPropertiesRaw := d.Get("catalog").(*schema.Set).List()
	var catalogProperties map[string]interface{}
	for _, raw := range catalogPropertiesRaw {
		catalogProperties = raw.(map[string]interface{})
		break
	}
	catalog := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(catalogProperties["count"].(int)),
			InstanceType: String(catalogProperties["instance_type"].(string)),
			VolumeType:   String(catalogProperties["instance_type"].(string)),
			VolumeSize:   Int32(catalogProperties["instance_type"].(int)),
			Image:        String(catalogProperties["instance_type"].(string)),
			Zone:         String(catalogProperties["instance_type"].(string)),
		},
	}

	fdbPropertiesRaw := d.Get("fdb").(*schema.Set).List()
	var fdbProperties map[string]interface{}
	for _, raw := range fdbPropertiesRaw {
		fdbProperties = raw.(map[string]interface{})
		break
	}
	fdb := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(fdbProperties["count"].(int)),
			InstanceType: String(fdbProperties["instance_type"].(string)),
			VolumeType:   String(fdbProperties["instance_type"].(string)),
			VolumeSize:   Int32(fdbProperties["instance_type"].(int)),
			Image:        String(fdbProperties["instance_type"].(string)),
			Zone:         String(fdbProperties["instance_type"].(string)),
		},
	}

	extraPropertiesRaw := d.Get("extra").(*schema.Set).List()
	var extraProperties map[string]interface{}
	for _, raw := range extraPropertiesRaw {
		extraProperties = raw.(map[string]interface{})
		break
	}
	extra := cloudmgr.CoreCreateServiceIaasExtraRequest{
		Vpc:     String(extraProperties["vpc"].(string)),
		Subnet:  String(extraProperties["subnet"].(string)),
		Keypair: String(extraProperties["keypair"].(string)),
	}

	ossPropertiesRaw := d.Get("oss").(*schema.Set).List()
	var ossProperties map[string]interface{}
	for _, raw := range ossPropertiesRaw {
		ossProperties = raw.(map[string]interface{})
		break
	}
	oss := cloudmgr.CoreCreateServiceOssZoneRequest{
		Name: String(ossProperties["name"].(string)),
	}

	metadataPropertiesRaw := d.Get("metadata").(*schema.Set).List()
	var metadataProperties map[string]interface{}
	for _, raw := range metadataPropertiesRaw {
		metadataProperties = raw.(map[string]interface{})
		break
	}
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["default_database"].(string)
	metadata["default_user"] = metadataProperties["default_user"].(string)
	metadata["default_password"] = metadataProperties["default_password"].(string)
	metadata["logic_part"] = metadataProperties["logic_part"].(int)

	name := d.Get("name").(string) //TODO 校验这个参数
	body.Name = &name
	body.Etcd = &etcd
	body.Catalog = &catalog
	body.Fdb = &fdb
	body.Extras = &extra
	body.Oss = &oss
	body.Metadata = &metadata

	resp, r, err := apiClient.CoreWarehouseServiceApi.CreateCatalog(ctx).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreWarehouseServiceApi.CreateWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarehouse`: CommonDescribeJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreWarehouseServiceApi.CreateWarehouse`: %v\n", resp)

	return nil
}

func resourceCatalogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}

func resourceCatalogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}

func resourceCatalogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}
