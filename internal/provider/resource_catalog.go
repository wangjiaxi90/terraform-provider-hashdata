package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wangjiaxi90/terraform-provider-hashdata/internal/provider/cloudmgr"
	_nethttp "net/http"
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
			"name": {
				Description: "name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"etcd": {
				Description: "etcd.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": { //TODO count 默认为1吗？
							Description: "etcd count.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"instance_type": {
							Description: "etcd instance_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_type": {
							Description: "etcd volume_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_size": {
							Description: "etcd volume_size.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"image": {
							Description: "etcd image.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"zone": {
							Description: "etcd zone.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"catalog": {
				Description: "etcd.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": { //TODO count 默认为1吗？
							Description: "catalog count.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"instance_type": {
							Description: "catalog instance_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_type": {
							Description: "catalog volume_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_size": {
							Description: "catalog volume_size.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"image": {
							Description: "catalog image.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"zone": {
							Description: "catalog zone.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"fdb": {
				Description: "etcd.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": { //TODO count 默认为1吗？
							Description: "fdb count.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"instance_type": {
							Description: "fdb instance_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_type": {
							Description: "fdb volume_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_size": {
							Description: "fdb volume_size.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"image": {
							Description: "fdb image.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"zone": {
							Description: "fdb zone.",
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
			"oss": {
				Description: "oss.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "oss name.",
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
		},
	}
}

func resourceCatalogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	body := *cloudmgr.NewCoreCreateCatalogRequest()
	apiClient := meta.(*cloudmgr.APIClient)
	etcdPropertiesRaw := d.Get("etcd").(*schema.Set).List()
	var etcdProperties = etcdPropertiesRaw[0].(map[string]interface{})
	etcd := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(etcdProperties["count"].(int)),
			InstanceType: String(etcdProperties["instance_type"].(string)),
			VolumeType:   String(etcdProperties["volume_type"].(string)),
			VolumeSize:   Int32(etcdProperties["volume_size"].(int)),
			Image:        String(etcdProperties["image"].(string)),
			Zone:         String(etcdProperties["zone"].(string)),
		},
	}

	catalogPropertiesRaw := d.Get("catalog").(*schema.Set).List()
	var catalogProperties = catalogPropertiesRaw[0].(map[string]interface{})
	catalog := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(catalogProperties["count"].(int)),
			InstanceType: String(catalogProperties["instance_type"].(string)),
			VolumeType:   String(catalogProperties["volume_type"].(string)),
			VolumeSize:   Int32(catalogProperties["volume_size"].(int)),
			Image:        String(catalogProperties["image"].(string)),
			Zone:         String(catalogProperties["zone"].(string)),
		},
	}

	fdbPropertiesRaw := d.Get("fdb").(*schema.Set).List()
	var fdbProperties = fdbPropertiesRaw[0].(map[string]interface{})
	fdb := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(fdbProperties["count"].(int)),
			InstanceType: String(fdbProperties["instance_type"].(string)),
			VolumeType:   String(fdbProperties["volume_type"].(string)),
			VolumeSize:   Int32(fdbProperties["volume_size"].(int)),
			Image:        String(fdbProperties["image"].(string)),
			Zone:         String(fdbProperties["zone"].(string)),
		},
	}

	extraPropertiesRaw := d.Get("extra").(*schema.Set).List()
	var extraProperties = extraPropertiesRaw[0].(map[string]interface{})
	extra := cloudmgr.CoreCreateServiceIaasExtraRequest{
		Vpc:     String(extraProperties["vpc"].(string)),
		Subnet:  String(extraProperties["subnet"].(string)),
		Keypair: String(extraProperties["keypair"].(string)),
	}

	ossPropertiesRaw := d.Get("oss").(*schema.Set).List()
	var ossProperties = ossPropertiesRaw[0].(map[string]interface{})
	oss := cloudmgr.CoreCreateServiceOssZoneRequest{
		Name: String(ossProperties["name"].(string)),
	}

	metadataPropertiesRaw := d.Get("metadata").(*schema.Set).List()
	var metadataProperties = metadataPropertiesRaw[0].(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["default_database"].(string)
	metadata["default_user"] = metadataProperties["default_user"].(string)
	metadata["default_password"] = metadataProperties["default_password"].(string)
	metadata["logic_part"] = metadataProperties["logic_part"].(int)

	name := d.Get("name").(string)
	body.Name = &name
	body.Etcd = &etcd
	body.Catalog = &catalog
	body.Fdb = &fdb
	body.Extras = &extra
	body.Oss = &oss
	body.Metadata = &metadata

	resp, r, err := apiClient.CoreWarehouseServiceApi.CreateCatalog(ctx).Body(body).Execute()
	if err != nil {
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateCatalog``: %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateCatalog``: %s\n", r.Status)
	}
	d.SetId(resp.GetResourceIds()[0])
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}

	//return nil
	return resourceCatalogUpdate(ctx, d, meta)
}

func resourceCatalogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	if id == "" {
		return diag.Errorf("Can not read catalog,because id is empty")
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

func resourceCatalogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceCatalogRead(ctx, d, meta)
}

func resourceCatalogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*cloudmgr.APIClient)
	resourceId := d.Id()
	if resourceId == "" {
		return diag.Errorf(CATALOG_ID + " not found! ")
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
