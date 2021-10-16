package provider

import (
	"context"
	"fmt"
	huge "github.com/dablelv/go-huge-util"
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

func resourceWarehouseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	fmt.Println("================================")
	s, _ := huge.ToIndentJSON(&d)
	fmt.Printf("schema=%v\n", s)
	fmt.Println("================================")
	body := *cloudmgr.NewCoreCreateWarehouseRequest() // CoreCreateWarehouseRequest |

	configuration := cloudmgr.NewConfiguration() //TODO
	apiClient := cloudmgr.NewAPIClient(configuration)

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

	metadataProperties := d.Get("metadata").(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["default_database"].(string)
	metadata["default_user"] = metadataProperties["default_user"].(string)
	metadata["default_password"] = metadataProperties["default_password"].(string)
	metadata["logic_part"] = metadataProperties["logic_part"].(string)

	featureProperties := d.Get("feature").(map[string]interface{})
	feature := cloudmgr.CoreCreateServiceFeatureRequest{
		LocalStorage:  featureProperties["is_local_storage"].(*bool),
		MirrorStandby: featureProperties["is_mirror_standby"].(*bool),
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
			Name: ossProperties["name"].(*string),
		}
		body.Oss = &oss
	}

	if featureProperties["is_mirror_standby"].(bool) {
		standbyProperties := d.Get("standby").(map[string]interface{})
		standby := cloudmgr.CoreCreateServiceComponentRequest{
			Iaas: &cloudmgr.CloudmgrcoreIaasResource{
				Count:        standbyProperties["count"].(*int32),
				InstanceType: standbyProperties["instance_type"].(*string),
				VolumeType:   standbyProperties["volume_type"].(*string),
				VolumeSize:   standbyProperties["volume_size"].(*int32),
				Image:        standbyProperties["image"].(*string),
				Zone:         standbyProperties["zone"].(*string),
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
