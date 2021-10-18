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

	configuration := cloudmgr.NewConfiguration() //TODO 客户端的生成
	apiClient := cloudmgr.NewAPIClient(configuration)

	catalog := d.Get("catalog").(string) //TODO 这里判断一下catalog是否为nil 或者为空 如果true的话

	masterPropertiesRaw := d.Get("master").(*schema.Set).List()
	var masterProperties map[string]interface{}
	for _, raw := range masterPropertiesRaw {
		masterProperties = raw.(map[string]interface{})
		break
	}

	master := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(masterProperties["count"].(int)),
			InstanceType: String(masterProperties["instance_type"].(string)),
			VolumeType:   String(masterProperties["volume_type"].(string)),
			VolumeSize:   Int32(masterProperties["volume_size"].(int)),
			Image:        String(masterProperties["image"].(string)),
			Zone:         String(masterProperties["zone"].(string)),
		},
	}

	segmentPropertiesRaw := d.Get("segment").(*schema.Set).List()
	var segmentProperties map[string]interface{}
	for _, raw := range segmentPropertiesRaw {
		segmentProperties = raw.(map[string]interface{})
		break
	}
	segment := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(segmentProperties["count"].(int)),
			InstanceType: String(segmentProperties["instance_type"].(string)),
			VolumeType:   String(segmentProperties["volume_type"].(string)),
			VolumeSize:   Int32(segmentProperties["volume_size"].(int)),
			Image:        String(segmentProperties["image"].(string)),
			Zone:         String(segmentProperties["zone"].(string)),
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

	name := d.Get("name").(string)
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
