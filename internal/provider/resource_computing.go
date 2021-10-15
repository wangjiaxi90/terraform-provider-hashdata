package provider

import (
	"context"
	"fmt"
	"github.com/wangjiaxi90/terraform-provider-scaffolding/internal/provider/cloudmgr"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputing() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Computing.",

		CreateContext: resourceComputingCreate,
		ReadContext:   resourceComputingRead,
		UpdateContext: resourceComputingUpdate,
		DeleteContext: resourceComputingDelete,

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

func resourceComputingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	body := *cloudmgr.NewCoreCreateWarehouseRequest() // CoreCreateWarehouseRequest |

	configuration := cloudmgr.NewConfiguration() //TODO
	apiClient := cloudmgr.NewAPIClient(configuration)

	catalog := d.Get("catalog").(string) //TODO 这里判断一下catalog是否为nil 或者为空 如果true的话

	masterProperties := d.Get("master").(map[string]interface{}) //TODO  master.iaas 是不是这么调用

	master := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        masterProperties["count"].(*int32), //TODO 可以在这里统计参数 这块加 【*】  对不对？
			InstanceType: masterProperties["instance_type"].(*string),
			VolumeType:   masterProperties["volume_type"].(*string),
			VolumeSize:   masterProperties["volume_size"].(*int32),
			Image:        masterProperties["image"].(*string),
			Zone:         masterProperties["zone"].(*string),
		},
	}

	segmentProperties := d.Get("segment").(map[string]interface{})
	segment := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        segmentProperties["count"].(*int32),
			InstanceType: segmentProperties["instance_type"].(*string),
			VolumeType:   segmentProperties["volume_type"].(*string),
			VolumeSize:   segmentProperties["volume_size"].(*int32),
			Image:        segmentProperties["image"].(*string),
			Zone:         segmentProperties["zone"].(*string),
		},
	}
	extraProperties := d.Get("extra").(map[string]interface{})
	extra := cloudmgr.CoreCreateServiceIaasExtraRequest{
		Vpc:     extraProperties["vpc"].(*string),
		Subnet:  extraProperties["subnet"].(*string),
		Keypair: extraProperties["keypair"].(*string),
	}

	name := d.Get("name").(string) //TODO 记得加这个参数
	body.Name = &name
	body.Master = &master
	body.Segment = &segment
	body.Catalog = &catalog
	body.Extras = &extra

	resp, r, err := apiClient.CoreWarehouseServiceApi.CreateWarehouse(ctx).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreWarehouseServiceApi.CreateWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarehouse`: CommonDescribeJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreWarehouseServiceApi.CreateWarehouse`: %v\n", resp)

	//return diag.Errorf("not implemented")
	return nil
}

func resourceComputingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}

func resourceComputingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}

func resourceComputingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//return diag.Errorf("not implemented")
	return nil
}
