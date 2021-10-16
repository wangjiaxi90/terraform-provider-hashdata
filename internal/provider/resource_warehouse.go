package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wangjiaxi90/terraform-provider-scaffolding/internal/provider/cloudmgr"
	"os"
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
			"master": {
				Description: "master.",
				Type:        schema.TypeMap,
				Optional:    true,
			},
			"segment": {
				Description: "segment.",
				Type:        schema.TypeMap,
				Optional:    true,
			},
			"extra": {
				Description: "extra.",
				Type:        schema.TypeMap,
				Optional:    true,
			},
			"metadata": {
				Description: "extra.",
				Type:        schema.TypeMap,
				Optional:    true,
			},
			"feature": {
				Description: "extra.",
				Type:        schema.TypeMap,
				Optional:    true,
			},
			"name": {
				Description: "extra.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func Int(v int32) *int32 {
	return &v
}
func String(v string) *string {
	return &v
}
func Bool(v bool) *bool {
	return &v
}
func resourceWarehouseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	body := *cloudmgr.NewCoreCreateWarehouseRequest() // CoreCreateWarehouseRequest |

	configuration := cloudmgr.NewConfiguration() //TODO
	apiClient := cloudmgr.NewAPIClient(configuration)
	masterProperties := d.Get("master").(map[string]interface{}) //TODO  master.iaas 是不是这么调用
	master := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int(masterProperties["count"].(int32)), //TODO 可以在这里统计参数 这块加 【*】  对不对？
			InstanceType: String(masterProperties["instance_type"].(string)),
			VolumeType:   String(masterProperties["volume_type"].(string)),
			VolumeSize:   Int(masterProperties["volume_size"].(int32)),
			Image:        String(masterProperties["image"].(string)),
			Zone:         String(masterProperties["zone"].(string)),
		},
	}

	segmentProperties := d.Get("segment").(map[string]interface{})
	segment := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int(segmentProperties["count"].(int32)), //TODO 可以在这里统计参数 这块加 【*】  对不对？
			InstanceType: String(segmentProperties["instance_type"].(string)),
			VolumeType:   String(segmentProperties["volume_type"].(string)),
			VolumeSize:   Int(segmentProperties["volume_size"].(int32)),
			Image:        String(segmentProperties["image"].(string)),
			Zone:         String(segmentProperties["zone"].(string)),
		},
	}
	extraProperties := d.Get("extra").(map[string]interface{})
	extra := cloudmgr.CoreCreateServiceIaasExtraRequest{
		Vpc:     String(extraProperties["vpc"].(string)),
		Subnet:  String(extraProperties["subnet"].(string)),
		Keypair: String(extraProperties["keypair"].(string)),
	}

	metadataProperties := d.Get("metadata").(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["default_database"].(string)
	metadata["default_user"] = metadataProperties["default_user"].(string)
	metadata["default_password"] = metadataProperties["default_password"].(string)
	metadata["logic_part"] = metadataProperties["logic_part"].(string)

	featureProperties := d.Get("feature").(map[string]interface{})
	feature := cloudmgr.CoreCreateServiceFeatureRequest{
		LocalStorage:  Bool(featureProperties["is_local_storage"].(bool)),
		MirrorStandby: Bool(featureProperties["is_mirror_standby"].(bool)),
	}
	name := d.Get("name").(string) //TODO 记得加这个参数
	body.Name = &name
	body.Master = &master
	body.Segment = &segment
	body.Extras = &extra
	body.Metadata = &metadata
	body.Features = &feature

	if featureProperties["is_local_storage"].(bool) {
		ossProperties := d.Get("oss").(map[string]interface{})
		oss := cloudmgr.CoreCreateServiceOssZoneRequest{
			Name: String(ossProperties["name"].(string)),
		}
		body.Oss = &oss
	}

	if featureProperties["is_mirror_standby"].(bool) {
		standbyProperties := d.Get("standby").(map[string]interface{})
		standby := cloudmgr.CoreCreateServiceComponentRequest{
			Iaas: &cloudmgr.CloudmgrcoreIaasResource{
				Count:        Int(standbyProperties["count"].(int32)),
				InstanceType: String(standbyProperties["instance_type"].(string)),
				VolumeType:   String(standbyProperties["volume_type"].(string)),
				VolumeSize:   Int(standbyProperties["volume_size"].(int32)),
				Image:        String(standbyProperties["image"].(string)),
				Zone:         String(standbyProperties["zone"].(string)),
			},
		}
		body.Standby = &standby
	}

	resp, r, err := apiClient.CoreWarehouseServiceApi.CreateWarehouse(ctx).Body(body).Execute()
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

	//configuration := cloudmgr.NewConfiguration()//TODO
	//apiClient := cloudmgr.NewAPIClient(configuration)
	//apiClient.CloudmgrServiceManager
	//
	//cloudmgr.Service
	return nil
}
