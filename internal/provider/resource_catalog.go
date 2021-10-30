package provider

import (
	"context"
	"fmt"
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"image": {
				Description: "image.",
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
	image := d.Get("image")
	errMsg, err2 := checkCatalogCreateSchema(d)
	if err2 != nil {
		return diag.Errorf(errMsg)
	}
	etcdPropertiesRaw := d.Get("etcd").(*schema.Set).List()
	var etcdProperties = etcdPropertiesRaw[0].(map[string]interface{})
	etcd := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(etcdProperties["count"].(int)),
			InstanceType: String(etcdProperties["instance_type"].(string)),
			VolumeType:   String(etcdProperties["volume_type"].(string)),
			VolumeSize:   Int32(etcdProperties["volume_size"].(int)),
			Image:        String(image.(string)),
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
			Image:        String(image.(string)),
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
			Image:        String(image.(string)),
			Zone:         String(fdbProperties["zone"].(string)),
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

	if ossOk, ok := d.GetOk("oss"); ok {
		ossProperties := ossOk.(map[string]interface{})
		if ossName, ok := ossProperties["name"]; ok {
			oss := cloudmgr.CoreCreateServiceOssZoneRequest{
				Name: String(ossName.(string)),
			}
			body.Oss = &oss
		}
	}

	metadataPropertiesRaw := d.Get("metadata").(*schema.Set).List()
	var metadataProperties = metadataPropertiesRaw[0].(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["default_database"].(string)
	metadata["default_user"] = metadataProperties["default_user"].(string)
	metadata["default_password"] = metadataProperties["default_password"].(string)
	if numberSegments, ok := metadataProperties["number_segments"]; ok {
		metadata["number_segments"] = numberSegments
	}
	if logicPart, ok := metadataProperties["logic_part"]; ok {
		metadata["logic_part"] = logicPart
	}

	name := d.Get("name").(string)
	body.Name = &name
	body.Etcd = &etcd
	body.Catalog = &catalog
	body.Fdb = &fdb
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
	id := d.Id()
	if id == "" {
		return diag.Errorf("Can not read warehouse,because id is empty")
	}
	apiClient := meta.(*cloudmgr.APIClient)
	isServiceStop := false
	//step 1: describe instance
	//step 2: getCount
	//step 3: judgment expend or shrink
	//step 4: stop service
	//step 5: async job
	//step 6: start service

	if d.HasChange("fdb") && !d.IsNewResource() {
		component := []string{"fdb"}
		respListInstance, rListInstance, errListInstance := apiClient.CoreServiceApi.ListServiceInstance(ctx, id).Component(component).Execute()
		if errListInstance != nil {
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance``: %v\n", errListInstance)
		}
		if rListInstance.StatusCode != 200 {
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance``: %s\n", rListInstance.Status)
		}
		countOld := respListInstance.GetCount()

		fdbPropertiesRaw := d.Get("fdb").(*schema.Set).List()
		var fdbProperties = fdbPropertiesRaw[0].(map[string]interface{})
		countNew := fdbProperties["count"].(int)

		if int32(countNew) != countOld {
			respStopService, rStopService, errStopService := apiClient.CoreServiceApi.StopService(ctx, id).Execute()
			if errStopService != nil {
				return diag.Errorf("Error when calling `CoreServiceApi.StopService`: %v\v", errStopService)
			}
			if rStopService.StatusCode != 200 {
				return diag.Errorf("Error when calling `CoreServiceApi.StopService`: %s\n", rStopService.Status)
			}
			if eWait := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStopService.GetId()); eWait != nil {
				return diag.Errorf("Error when waiting stop service %s\n", eWait.Error())
			}
			isServiceStop = true
			componentRequestMap := make(map[string]interface{})
			var respScaleOut cloudmgr.CommonDescribeJobResponse
			var rScaleOut *_nethttp.Response
			var errScaleOut error
			if int32(countNew) > countOld {
				componentRequestMap["fdb"] = cloudmgr.CoreScaleOutServiceComponentRequest{
					Iaas: &cloudmgr.CoreScaleOutIaasResource{
						Count: Int32(countNew),
					},
				}
				respScaleOut, rScaleOut, errScaleOut = apiClient.CoreServiceApi.ScaleOutService(ctx, id).Body(cloudmgr.CoreScaleOutServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			} else {
				var remainInstances = make([]string, countNew)
				for i := 0; i < countNew; i++ {
					remainInstances[i] = (*respListInstance.Content)[i].GetId()
				}
				componentRequestMap["fdb"] = cloudmgr.CoreScaleInServiceComponentRequest{
					Instances: &remainInstances,
				}
				respScaleOut, rScaleOut, errScaleOut = apiClient.CoreServiceApi.ScaleInService(ctx, id).Body(cloudmgr.CoreScaleInServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			}

			if errScaleOut != nil {
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %v\v", errListInstance)
			}
			if rScaleOut.StatusCode != 200 {
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\n", rListInstance.Status)
			}
			err2 := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respScaleOut.GetId())
			if err2 != nil {
				return diag.Errorf("Error when wait calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\v", err2)
			}
		}

		if isServiceStop {
			respStartService, rStartService, errStartService := apiClient.CoreServiceApi.StartService(ctx, id).Execute()
			if errStartService != nil {
				return diag.Errorf("Error when calling `CoreServiceApi.StartService`: %v\v", errStartService)
			}
			if rStartService.StatusCode != 200 {
				return diag.Errorf("Error when calling `CoreServiceApi.StartService`: %s\n", rStartService.Status)
			}
			if eWait := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStartService.GetId()); eWait != nil {
				return diag.Errorf("Error when waiting start service %s\n", eWait.Error())
			}
		}
	}

	if d.HasChange("etcd") && !d.IsNewResource() {
		component := []string{"etcd"}
		respListInstance, rListInstance, errListInstance := apiClient.CoreServiceApi.ListServiceInstance(ctx, id).Component(component).Execute()
		if errListInstance != nil {
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance``: %v\n", errListInstance)
		}
		if rListInstance.StatusCode != 200 {
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance``: %s\n", rListInstance.Status)
		}
		countOld := respListInstance.GetCount()

		etcdPropertiesRaw := d.Get("etcd").(*schema.Set).List()
		var etcdProperties = etcdPropertiesRaw[0].(map[string]interface{})
		countNew := etcdProperties["count"].(int)

		if int32(countNew) != countOld {
			if !isServiceStop {
				if errStopService := stopService(ctx, id, apiClient); errStopService != nil {
					return diag.Errorf(errStopService.Error())
				}
				isServiceStop = true
			}

			componentRequestMap := make(map[string]interface{})
			var respScaleOut cloudmgr.CommonDescribeJobResponse
			var rScaleOut *_nethttp.Response
			var errScaleOut error
			if int32(countNew) > countOld {
				componentRequestMap["etcd"] = cloudmgr.CoreScaleOutServiceComponentRequest{
					Iaas: &cloudmgr.CoreScaleOutIaasResource{
						Count: Int32(countNew),
					},
				}
				respScaleOut, rScaleOut, errScaleOut = apiClient.CoreServiceApi.ScaleOutService(ctx, id).Body(cloudmgr.CoreScaleOutServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			} else {
				var remainInstances = make([]string, countNew)
				for i := 0; i < countNew; i++ {
					remainInstances[i] = (*respListInstance.Content)[i].GetId()
				}
				componentRequestMap["etcd"] = cloudmgr.CoreScaleInServiceComponentRequest{
					Instances: &remainInstances,
				}
				respScaleOut, rScaleOut, errScaleOut = apiClient.CoreServiceApi.ScaleInService(ctx, id).Body(cloudmgr.CoreScaleInServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			}

			if errScaleOut != nil {
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %v\v", errListInstance)
			}
			if rScaleOut.StatusCode != 200 {
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\n", rListInstance.Status)
			}
			err2 := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respScaleOut.GetId())
			if err2 != nil {
				return diag.Errorf("Error when wait calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\v", err2)
			}
		}
	}

	if d.HasChange("catalog") && !d.IsNewResource() {
		component := []string{"catalog"}
		respListInstance, rListInstance, errListInstance := apiClient.CoreServiceApi.ListServiceInstance(ctx, id).Component(component).Execute()
		if errListInstance != nil {
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance``: %v\n", errListInstance)
		}
		if rListInstance.StatusCode != 200 {
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance``: %s\n", rListInstance.Status)
		}
		countOld := respListInstance.GetCount()

		catalogPropertiesRaw := d.Get("catalog").(*schema.Set).List()
		var catalogProperties = catalogPropertiesRaw[0].(map[string]interface{})
		countNew := catalogProperties["count"].(int)

		if int32(countNew) != countOld {
			if !isServiceStop {
				if errStopService := stopService(ctx, id, apiClient); errStopService != nil {
					return diag.Errorf(errStopService.Error())
				}
				isServiceStop = true
			}
			componentRequestMap := make(map[string]interface{})
			var respScaleOut cloudmgr.CommonDescribeJobResponse
			var rScaleOut *_nethttp.Response
			var errScaleOut error
			if int32(countNew) > countOld {
				componentRequestMap["catalog"] = cloudmgr.CoreScaleOutServiceComponentRequest{
					Iaas: &cloudmgr.CoreScaleOutIaasResource{
						Count: Int32(countNew),
					},
				}
				respScaleOut, rScaleOut, errScaleOut = apiClient.CoreServiceApi.ScaleOutService(ctx, id).Body(cloudmgr.CoreScaleOutServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			} else {
				var remainInstances = make([]string, countNew)
				for i := 0; i < countNew; i++ {
					remainInstances[i] = (*respListInstance.Content)[i].GetId()
				}
				componentRequestMap["catalog"] = cloudmgr.CoreScaleInServiceComponentRequest{
					Instances: &remainInstances,
				}
				respScaleOut, rScaleOut, errScaleOut = apiClient.CoreServiceApi.ScaleInService(ctx, id).Body(cloudmgr.CoreScaleInServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			}

			if errScaleOut != nil {
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %v\v", errListInstance)
			}
			if rScaleOut.StatusCode != 200 {
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\n", rListInstance.Status)
			}
			err2 := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respScaleOut.GetId())
			if err2 != nil {
				return diag.Errorf("Error when wait calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\v", err2)
			}
		}
	}

	if isServiceStop {
		respStartService, rStartService, errStartService := apiClient.CoreServiceApi.StartService(ctx, id).Execute()
		if errStartService != nil {
			return diag.Errorf("Error when calling `CoreServiceApi.StartService`: %v\v", errStartService)
		}
		if rStartService.StatusCode != 200 {
			return diag.Errorf("Error when calling `CoreServiceApi.StartService`: %s\n", rStartService.Status)
		}
		if eWait := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStartService.GetId()); eWait != nil {
			return diag.Errorf("Error when waiting start service %s\n", eWait.Error())
		}
	}
	return resourceCatalogRead(ctx, d, meta)
}

func resourceCatalogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//先stop 在 delete

	apiClient := meta.(*cloudmgr.APIClient)
	resourceId := d.Id()
	if resourceId == "" {
		return diag.Errorf(CATALOG_ID + " not found! ")
	}
	False := false
	resp1, r1, err1 := apiClient.CoreServiceApi.StopService(ctx, resourceId).Body(cloudmgr.CoreStopServiceRequest{
		Force: &False,
	}).Execute()
	if err1 != nil {
		return diag.Errorf("Error when calling catalog `CoreServiceApi.StopService``: %v\n", err1)
	}
	if r1.StatusCode != 200 {
		return diag.Errorf("Stop resource fail with %d . ", r1.StatusCode)
	}
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp1.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}

	resp, r, err := apiClient.CoreServiceApi.DeleteService(ctx, resourceId).Execute()
	if err != nil {
		return diag.Errorf("Error when calling catalog `CoreServiceApi.DeleteService``: %v\n", err)
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

func checkCatalogCreateSchema(d *schema.ResourceData) (string, error) {
	res := ""
	_, ok := d.GetOk("name")
	if !ok {
		res += "schema name field is missing\n"
	}
	_, ok = d.GetOk("image")
	if !ok {
		res += "schema image field is missing\n"
	}
	etcdRaw, ok := d.GetOk("etcd")
	if !ok {
		res += "schema etcd field is missing\n"
	}
	etcdMap := etcdRaw.(map[string]interface{})
	if _, ok := etcdMap["count"]; !ok {
		res += "schema etcd.count field is missing\n"
	}
	if _, ok := etcdMap["instance_type"]; !ok {
		res += "schema etcd.instance_type field is missing\n"
	}
	if _, ok := etcdMap["volume_type"]; !ok {
		res += "schema etcd.volume_type field is missing\n"
	}
	if _, ok := etcdMap["volume_size"]; !ok {
		res += "schema etcd.volume_size field is missing\n"
	}
	if _, ok := etcdMap["zone"]; !ok {
		res += "schema etcd.zone field is missing\n"
	}

	catalogRaw, ok := d.GetOk("catalog")
	if !ok {
		res += "schema catalog field is missing\n"
	}
	catalogMap := catalogRaw.(map[string]interface{})
	if _, ok := catalogMap["count"]; !ok {
		res += "schema catalog.count field is missing\n"
	}
	if _, ok := catalogMap["instance_type"]; !ok {
		res += "schema catalog.instance_type field is missing\n"
	}
	if _, ok := catalogMap["volume_type"]; !ok {
		res += "schema catalog.volume_type field is missing\n"
	}
	if _, ok := catalogMap["volume_size"]; !ok {
		res += "schema catalog.volume_size field is missing\n"
	}
	if _, ok := catalogMap["zone"]; !ok {
		res += "schema catalog.zone field is missing\n"
	}

	fdbRaw, ok := d.GetOk("fdb")
	if !ok {
		res += "schema segment field is missing\n"
	}
	fdbMap := fdbRaw.(map[string]interface{})
	if _, ok := fdbMap["count"]; !ok {
		res += "schema fdb.count field is missing\n"
	}
	if _, ok := fdbMap["instance_type"]; !ok {
		res += "schema fdb.instance_type field is missing\n"
	}
	if _, ok := fdbMap["volume_type"]; !ok {
		res += "schema fdb.volume_type field is missing\n"
	}
	if _, ok := fdbMap["volume_size"]; !ok {
		res += "schema fdb.volume_size field is missing\n"
	}
	if _, ok := fdbMap["zone"]; !ok {
		res += "schema fdb.zone field is missing\n"
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

	//ossRaw, ok := d.GetOk("oss")
	//if !ok {
	//	res += "schema oss field is missing\n"
	//}
	//ossMap := ossRaw.(map[string]interface{})
	//if _, ok := ossMap["name"]; !ok {
	//	res += "schema oss.name field is missing\n"
	//}

	metadataRaw, ok := d.GetOk("metadata")
	if !ok {
		res += "schema metadata field is missing\n"
	}
	metadataMap := metadataRaw.(map[string]interface{})
	if _, ok := metadataMap["default_database"]; !ok {
		res += "schema metadata.default_database field is missing\n"
	}
	if _, ok := metadataMap["default_user"]; !ok {
		res += "schema metadata.default_user field is missing\n"
	}
	if _, ok := metadataMap["default_password"]; !ok {
		res += "schema metadata.default_password field is missing\n"
	}
	//if _, ok := metadataMap["logic_part"]; !ok {
	//	res += "schema metadata.logic_part field is missing\n"
	//}

	if res != "" {
		return res, fmt.Errorf("Input is illegal. ")
	}
	return "", nil
}
