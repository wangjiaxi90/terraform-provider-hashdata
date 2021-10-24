package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wangjiaxi90/terraform-provider-hashdata/internal/provider/cloudmgr"
	_nethttp "net/http"
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
			"name": {
				Description: "name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"master": {
				Description: "master.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Description: "master count.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"instance_type": {
							Description: "master instance_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_type": {
							Description: "master volume_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_size": {
							Description: "master volume_size.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"image": {
							Description: "master image.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"zone": {
							Description: "master zone.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"segment": {
				Description: "segment.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Description: "segment count.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"instance_type": {
							Description: "segment instance_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_type": {
							Description: "segment volume_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_size": {
							Description: "segment volume_size.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"image": {
							Description: "segment image.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"zone": {
							Description: "segment zone.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"extra": {
				Description: "extra.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc": {
							Description: "vpc.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"subnet": {
							Description: "subnet.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"keypair": {
							Description: "keypair.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"metadata": {
				Description: "metadata.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_database": {
							Description: "default_database.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"default_user": {
							Description: "default_user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"default_password": {
							Description: "default_password.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"number_segments": {
							Description: "number_segments.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"logic_part": {
							Description: "logic_part.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
			"feature": {
				Description: "feature.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_storage": {
							Description: "is_local_storage.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"mirror_standby": {
							Description: "is_mirror_standby.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWarehouseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ctx = context.Background()
	body := *cloudmgr.NewCoreCreateWarehouseRequest() // CoreCreateWarehouseRequest |
	apiClient := meta.(*cloudmgr.APIClient)
	masterPropertiesRaw := d.Get("master").(*schema.Set).List()
	var masterProperties = masterPropertiesRaw[0].(map[string]interface{})
	var masterCount int32 = 1
	master := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        &masterCount,
			InstanceType: String(masterProperties["instance_type"].(string)),
			VolumeType:   String(masterProperties["volume_type"].(string)),
			VolumeSize:   Int32(masterProperties["volume_size"].(int)),
			Image:        String(masterProperties["image"].(string)),
			Zone:         String(masterProperties["zone"].(string)),
		},
	}

	segmentPropertiesRaw := d.Get("segment").(*schema.Set).List()
	var segmentProperties = segmentPropertiesRaw[0].(map[string]interface{})
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
	var extraProperties = extraPropertiesRaw[0].(map[string]interface{})
	extra := cloudmgr.CoreCreateServiceIaasExtraRequest{
		Vpc:     String(extraProperties["vpc"].(string)),
		Subnet:  String(extraProperties["subnet"].(string)),
		Keypair: String(extraProperties["keypair"].(string)),
	}

	metadataPropertiesRaw := d.Get("metadata").(*schema.Set).List()
	var metadataProperties = metadataPropertiesRaw[0].(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["default_database"].(string)
	metadata["default_user"] = metadataProperties["default_user"].(string)
	metadata["default_password"] = metadataProperties["default_password"].(string)
	metadata["logic_part"] = metadataProperties["logic_part"].(int)

	featurePropertiesRaw := d.Get("feature").(*schema.Set).List()
	var featureProperties = featurePropertiesRaw[0].(map[string]interface{})
	feature := cloudmgr.CoreCreateServiceFeatureRequest{
		LocalStorage:  Bool(featureProperties["local_storage"].(bool)),
		MirrorStandby: Bool(featureProperties["mirror_standby"].(bool)),
	}
	name := d.Get("name").(string)
	body.Name = &name
	body.Master = &master
	body.Segment = &segment
	body.Extras = &extra
	body.Metadata = &metadata
	body.Features = &feature

	if featureProperties["local_storage"].(bool) {
		ossProperties := d.Get("oss").(map[string]interface{})
		oss := cloudmgr.CoreCreateServiceOssZoneRequest{
			Name: String(ossProperties["name"].(string)),
		}
		body.Oss = &oss
	}

	if featureProperties["mirror_standby"].(bool) {
		standbyProperties := d.Get("standby").(map[string]interface{})
		standby := cloudmgr.CoreCreateServiceComponentRequest{
			Iaas: &cloudmgr.CloudmgrcoreIaasResource{
				Count:        Int32(standbyProperties["count"].(int)),
				InstanceType: String(standbyProperties["instance_type"].(string)),
				VolumeType:   String(standbyProperties["volume_type"].(string)),
				VolumeSize:   Int32(standbyProperties["volume_size"].(int)),
				Image:        String(standbyProperties["image"].(string)),
				Zone:         String(standbyProperties["zone"].(string)),
			},
		}
		body.Standby = &standby
	}
	var resp cloudmgr.CommonDescribeJobResponse
	var r *_nethttp.Response
	var err error
	//Do Not retry here
	resp, r, err = apiClient.CoreWarehouseServiceApi.CreateWarehouse(ctx).Body(body).Execute()
	if err != nil {
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse``: %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse``: %s\n", r.Status)
	}

	d.SetId(resp.GetId()) //TODO这里应该判断一下
	d.Set(WAREHOUSE_ID, resp.GetResourceIds()[0])
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}
	return nil
	//return resourceWarehouseUpdate(ctx, d, meta)
}

func resourceWarehouseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, ok := d.GetOk(WAREHOUSE_ID)
	if !ok {
		return diag.Errorf("Can not read warehouse,because id is empty")
	}

	apiClient := meta.(*cloudmgr.APIClient)
	var resp cloudmgr.CoreDescribeInstanceResponse
	var r *_nethttp.Response
	var err error

	resp, r, err = apiClient.CoreInstanceServiceApi.DescribeInstance(ctx, id.(string)).Execute()

	if err != nil {
		return nil
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Error status code when calling `CoreWarehouseServiceApi.CreateWarehouse``: %d \n", r.StatusCode)
	}
	if param, ok := resp.GetArchOk(); !ok {
		d.Set("arch", param)
	}
	if param, ok := resp.GetComponentsOk(); !ok {
		d.Set("components", param)
	}
	if param, ok := resp.GetCreatedAtOk(); !ok {
		d.Set("created_at", param)
	}
	if param, ok := resp.GetDestroyedAtOk(); !ok {
		d.Set("destroyed_at", param)
	}
	if nic, ok := resp.GetElasticNicOk(); !ok {
		var nic_map = make(map[string]interface{})
		if param, ok2 := nic.GetElasticNicIdOk(); !ok2 {
			nic_map["elastic_nic_id"] = param
		}
		if param, ok2 := nic.GetIpaddressesOk(); !ok2 {
			nic_map["ipaddresses"] = param
		}
		d.Set("nic", nic_map)
	}
	if param, ok := resp.GetHealthStatusOk(); !ok {
		d.Set("health_status", param)
	}
	if param, ok := resp.GetHostnameOk(); !ok {
		d.Set("hostname", param)
	}
	if param, ok := resp.GetIdOk(); !ok {
		d.Set("id", param)
	}
	if param, ok := resp.GetImageOk(); !ok {
		d.Set("image", param)
	}
	if internet, ok := resp.GetInternetOk(); !ok {
		var internet_map = make(map[string]interface{})
		if param, ok2 := internet.GetBandwidthOk(); !ok2 {
			internet_map["band_width"] = param
		}
		if param, ok2 := internet.GetElasticIpOk(); !ok2 {
			internet_map["elastic_ip"] = param
		}
		if param, ok2 := internet.GetElasticIpIdOk(); !ok2 {
			internet_map["elastic_ip_id"] = param
		}
		if param, ok2 := internet.GetPublicIpOk(); !ok2 {
			internet_map["public_ip"] = param
		}
		d.Set("internet", internet_map)
	}
	if param, ok := resp.GetIpaddressesOk(); !ok {
		d.Set("ipaddresses", param)
	}
	if param, ok := resp.GetMessageOk(); !ok {
		d.Set("message", param)
	}
	if param, ok := resp.GetNameOk(); !ok {
		d.Set("name", param)
	}
	if param, ok := resp.GetResourcePoolOk(); !ok {
		d.Set("resource_pool", param)
	}
	if param, ok := resp.GetScaleTypeOk(); !ok {
		d.Set("scale_type", param)
	}
	if param, ok := resp.GetServiceOk(); !ok {
		d.Set("service", param)
	}
	if param, ok := resp.GetStatusOk(); !ok {
		d.Set("status", param)
	}
	if param, ok := resp.GetTenantOk(); !ok {
		d.Set("tenant", param)
	}
	if param, ok := resp.GetTypeOk(); !ok {
		d.Set("type", param)
	}
	if param, ok := resp.GetVendorOk(); !ok {
		d.Set("vendor", param)
	}
	if param, ok := resp.GetZoneOk(); !ok {
		d.Set("zone", param)
	}

	//TODO 判断是否被删除
	return nil
}

func resourceWarehouseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//TODO not support update
	return resourceWarehouseRead(ctx, d, meta)
}

func resourceWarehouseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*cloudmgr.APIClient)
	resourceId, ok := d.GetOk(WAREHOUSE_ID)
	if !ok {
		return diag.Errorf(WAREHOUSE_ID + " not found! ")
	}
	resp, r, err := apiClient.CoreServiceApi.DeleteService(ctx, resourceId.(string)).Execute()
	if err != nil {
		return diag.Errorf("Error when calling `CoreServiceApi.DeleteService``: %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Delete resource fail with %d . ", r.StatusCode)
	}
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}
	d.SetId("")
	return nil
}
