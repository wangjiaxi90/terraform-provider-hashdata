package provider

import (
	"context"
	"fmt"
	"github.com/wangjiaxi90/terraform-provider-scaffolding/internal/provider/cloudmgr"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWarehouse() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Warehouse.",

		CreateContext: resourceWarehouseCreate,
		ReadContext:   resourceWarehouseRead,
		UpdateContext: resourceWarehouseUpdate,
		DeleteContext: resourceWarehouseDelete,

		Schema: map[string]*schema.Schema{
			"warehouse_param": {
				// This description is used by the documentation generator and the language server.
				Description: "warehouse param.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceWarehouseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	body := *cloudmgr.NewCoreCreateWarehouseRequest() // CoreCreateWarehouseRequest |

	configuration := cloudmgr.NewConfiguration()
	api_client := cloudmgr.NewAPIClient(configuration)

	master_properties := d.Get("master").(map[string]interface{}) //TODO  master.iaas 是不是这么调用

	master := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        master_properties["count"].(*int32), //TODO 可以在这里统计参数
			InstanceType: master_properties["instance_type"].(*string),
			VolumeType:   master_properties["volume_type"].(*string),
			VolumeSize:   master_properties["volume_size"].(*int32),
			Image:        master_properties["image"].(*string),
			Zone:         master_properties["zone"].(*string),
		},
	}

	segment_properties := d.Get("segment").(map[string]interface{})
	segment := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        segment_properties["count"].(*int32),
			InstanceType: segment_properties["instance_type"].(*string),
			VolumeType:   segment_properties["volume_type"].(*string),
			VolumeSize:   segment_properties["volume_size"].(*int32),
			Image:        segment_properties["image"].(*string),
			Zone:         segment_properties["zone"].(*string),
		},
	}
	extra_properties := d.Get("extra").(map[string]interface{})
	extra := cloudmgr.CoreCreateServiceIaasExtraRequest{
		Vpc:     extra_properties["vpc"].(*string),
		Subnet:  extra_properties["subnet"].(*string),
		Keypair: extra_properties["keypair"].(*string),
	}

	metadata_properties := d.Get("metadata").(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadata_properties["defaultDatabase"].(string)
	metadata["default_user"] = metadata_properties["defaultUser"].(string)
	metadata["default_password"] = metadata_properties["defaultPassword"].(string)
	metadata["logic_part"] = metadata_properties["logicPart"].(string)

	feature_properties := d.Get("feature").(map[string]interface{})
	feature := cloudmgr.CoreCreateServiceFeatureRequest{
		LocalStorage:  feature_properties["is_local_storage"].(*bool),
		MirrorStandby: feature_properties["is_local_storage"].(*bool),
	}
	name := d.Get("name").(*string) //TODO 记得加这个参数
	body.Name = name
	body.Master = &master
	body.Segment = &segment
	body.Extras = &extra
	body.Metadata = &metadata
	body.Features = &feature
	resp, r, err := api_client.CoreWarehouseServiceApi.CreateWarehouse(ctx).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreWarehouseServiceApi.CreateWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarehouse`: CommonDescribeJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreWarehouseServiceApi.CreateWarehouse`: %v\n", resp)

	return nil
}

func resourceWarehouseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}

func resourceWarehouseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}

func resourceWarehouseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}
