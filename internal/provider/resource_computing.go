package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wangjiaxi90/terraform-provider-hashdata/internal/provider/cloudmgr"
	_nethttp "net/http"
)

func resourceComputing() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Computing.",

		CreateContext: resourceComputingCreate,
		ReadContext:   resourceComputingRead,
		UpdateContext: resourceComputingUpdate,
		DeleteContext: resourceComputingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"catalog": {
				Description: "catalog UUID.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"image": {
				Description: "image.",
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
		},
	}
}

func resourceComputingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	body := *cloudmgr.NewCoreCreateWarehouseRequest()
	apiClient := meta.(*cloudmgr.APIClient)
	errMsg, err2 := checkComputingCreateSchema(d)
	if err2 != nil {
		return diag.Errorf(errMsg)
	}
	image := d.Get("image")
	catalog := d.Get("catalog").(string)

	masterPropertiesRaw := d.Get("master").(*schema.Set).List()
	var masterProperties = masterPropertiesRaw[0].(map[string]interface{})
	var masterCount int32 = 1

	master := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        &masterCount,
			InstanceType: String(masterProperties["instance_type"].(string)),
			VolumeType:   String(masterProperties["volume_type"].(string)),
			VolumeSize:   Int32(masterProperties["volume_size"].(int)),
			Image:        String(image.(string)),
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
			Image:        String(image.(string)),
			Zone:         String(segmentProperties["zone"].(string)),
		},
	}

	if extraRaw, ok := d.GetOk("extra"); ok {
		extraMap := extraRaw.(map[string]interface{})
		extra := cloudmgr.CoreCreateServiceIaasExtraRequest{}
		if vpc, ok := extraMap["vpc"]; ok {
			extra.Vpc = String(vpc.(string))
		}
		if subnet, ok := extraMap["subnet"]; ok {
			extra.Vpc = String(subnet.(string))
		}
		if keypair, ok := extraMap["keypair"]; ok {
			extra.Vpc = String(keypair.(string))
		}
		body.Extras = &extra
	}

	name := d.Get("name").(string)
	body.Name = &name
	body.Master = &master
	body.Segment = &segment
	body.Catalog = &catalog
	var resp cloudmgr.CommonDescribeJobResponse
	var r *_nethttp.Response
	var err error
	resp, r, err = apiClient.CoreWarehouseServiceApi.CreateWarehouse(ctx).Body(body).Execute()
	if err != nil {
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse_Computing``: %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse_Computing``: %s\n", r.Status)
	}
	// response from `CreateWarehouse`: CommonDescribeJobResponse
	d.SetId(resp.GetResourceIds()[0])
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}
	return nil
}

func resourceComputingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	if id == "" {
		return diag.Errorf("Can not read warehouse,because id is empty")
	}

	apiClient := meta.(*cloudmgr.APIClient)
	var resp cloudmgr.CoreDescribeInstanceResponse
	var r *_nethttp.Response
	var err error

	resp, r, err = apiClient.CoreInstanceServiceApi.DescribeInstance(ctx, id).Execute()

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

func resourceComputingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceComputingRead(ctx, d, meta)
}

func resourceComputingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*cloudmgr.APIClient)
	resourceId := d.Id()
	if resourceId == "" {
		return diag.Errorf(COMPUTING_ID + " not found! ")
	}

	False := false
	resp1, r1, err1 := apiClient.CoreServiceApi.StopService(ctx, resourceId).Body(cloudmgr.CoreStopServiceRequest{
		Force: &False,
	}).Execute()
	if err1 != nil {
		return diag.Errorf("Error when calling computing `CoreServiceApi.StopService``: %v\n", err1)
	}
	if r1.StatusCode != 200 {
		return diag.Errorf("Stop resource computing fail with %d . ", r1.StatusCode)
	}
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp1.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}

	resp, r, err := apiClient.CoreServiceApi.DeleteService(ctx, resourceId).Execute()
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

func checkComputingCreateSchema(d *schema.ResourceData) (string, error) {
	res := ""
	_, ok := d.GetOk("name")
	if !ok {
		res += "schema name field is missing\n"
	}
	_, ok = d.GetOk("image")
	if !ok {
		res += "schema image field is missing\n"
	}
	_, ok = d.GetOk("catalog")
	if !ok {
		res += "schema name field is missing\n"
	}

	masterRaw, ok := d.GetOk("master")
	if !ok {
		res += "schema master field is missing\n"
	}
	masterMap := masterRaw.(map[string]interface{})
	if _, ok := masterMap["instance_type"]; !ok {
		res += "schema master.instance_type field is missing\n"
	}
	if _, ok := masterMap["volume_type"]; !ok {
		res += "schema master.volume_type field is missing\n"
	}
	if _, ok := masterMap["volume_size"]; !ok {
		res += "schema master.volume_size field is missing\n"
	}
	if _, ok := masterMap["image"]; !ok {
		res += "schema master.image field is missing\n"
	}
	if _, ok := masterMap["zone"]; !ok {
		res += "schema master.zone field is missing\n"
	}

	segmentRaw, ok := d.GetOk("segment")
	if !ok {
		res += "schema segment field is missing\n"
	}
	segmentMap := segmentRaw.(map[string]interface{})
	if _, ok := segmentMap["count"]; !ok {
		res += "schema segment.count field is missing\n"
	}
	if _, ok := segmentMap["instance_type"]; !ok {
		res += "schema segment.instance_type field is missing\n"
	}
	if _, ok := segmentMap["volume_type"]; !ok {
		res += "schema segment.volume_type field is missing\n"
	}
	if _, ok := segmentMap["volume_size"]; !ok {
		res += "schema segment.volume_size field is missing\n"
	}
	if _, ok := segmentMap["image"]; !ok {
		res += "schema segment.image field is missing\n"
	}
	if _, ok := segmentMap["zone"]; !ok {
		res += "schema segment.zone field is missing\n"
	}
	//extraRaw, ok := d.GetOk("extra")
	//if !ok {
	//	res += "schema extra field is missing\n"
	//}
	//extraMap := extraRaw.(map[string]interface{})
	//if _, ok := extraMap["vpc"]; !ok {
	//	res += "schema extra.vpc field is missing\n"
	//}
	//if _, ok := extraMap["subnet"]; !ok {
	//	res += "schema extra.subnet field is missing\n"
	//}
	//if _, ok := extraMap["keypair"]; !ok {
	//	res += "schema extra.keypair field is missing\n"
	//}
	if res != "" {
		return res, fmt.Errorf("Input is illegal. ")
	}
	return "", nil
}
