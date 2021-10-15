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

	body := *cloudmgr.NewCoreCreateCatalogRequest() // CoreCreateWarehouseRequest |

	configuration := cloudmgr.NewConfiguration() //TODO
	apiClient := cloudmgr.NewAPIClient(configuration)

	etcdProperties := d.Get("etcd").(map[string]interface{}) //TODO  master.iaas 是不是这么调用
	etcd := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        etcdProperties["count"].(*int32),
			InstanceType: etcdProperties["instance_type"].(*string),
			VolumeType:   etcdProperties["instance_type"].(*string),
			VolumeSize:   etcdProperties["instance_type"].(*int32),
			Image:        etcdProperties["instance_type"].(*string),
			Zone:         etcdProperties["instance_type"].(*string),
		},
	}
	catalogProperties := d.Get("catalog").(map[string]interface{})
	catalog := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        catalogProperties["count"].(*int32),
			InstanceType: catalogProperties["instance_type"].(*string),
			VolumeType:   catalogProperties["instance_type"].(*string),
			VolumeSize:   catalogProperties["instance_type"].(*int32),
			Image:        catalogProperties["instance_type"].(*string),
			Zone:         catalogProperties["instance_type"].(*string),
		},
	}
	fdbProperties := d.Get("fdb").(map[string]interface{})
	fdb := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        fdbProperties["count"].(*int32),
			InstanceType: fdbProperties["instance_type"].(*string),
			VolumeType:   fdbProperties["instance_type"].(*string),
			VolumeSize:   fdbProperties["instance_type"].(*int32),
			Image:        fdbProperties["instance_type"].(*string),
			Zone:         fdbProperties["instance_type"].(*string),
		},
	}

	extraProperties := d.Get("extra").(map[string]interface{})
	extra := cloudmgr.CoreCreateServiceIaasExtraRequest{
		Vpc:     extraProperties["vpc"].(*string),
		Subnet:  extraProperties["subnet"].(*string),
		Keypair: extraProperties["keypair"].(*string),
	}
	ossProperties := d.Get("extra").(map[string]interface{})
	oss := cloudmgr.CoreCreateServiceOssZoneRequest{
		Name: ossProperties["name"].(*string),
	}

	metadataProperties := d.Get("metadata").(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["defaultDatabase"].(string)
	metadata["default_user"] = metadataProperties["defaultUser"].(string)
	metadata["default_password"] = metadataProperties["defaultPassword"].(string)
	metadata["logic_part"] = metadataProperties["logicPart"].(string)

	name := d.Get("name").(string) //TODO 记得加这个参数
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
