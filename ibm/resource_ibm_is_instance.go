package ibm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	computec "github.ibm.com/Bluemix/riaas-go-client/riaas/client/compute"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

const (
	isInstanceName                    = "name"
	isInstanceKeys                    = "keys"
	isInstanceTags                    = "tags"
	isInstanceNetworkInterfaces       = "network_interfaces"
	isInstancePrimaryNetworkInterface = "primary_network_interface"
	isInstanceNicName                 = "name"
	isInstanceProfile                 = "profile"
	isInstanceNicPortSpeed            = "port_speed"
	isInstanceNicPrimaryIpv4Address   = "primary_ipv4_address"
	isInstanceNicPrimaryIpv6Address   = "primary_ipv6_address"
	isInstanceNicSecondaryAddress     = "secondary_addresses"
	isInstanceNicSecurityGroups       = "security_groups"
	isInstanceNicSubnet               = "subnet"
	isInstanceNicFloatingIPs          = "floating_ips"
	isInstanceUserData                = "user_data"
	isInstanceVolumes                 = "volumes"
	isInstanceVPC                     = "vpc"
	isInstanceZone                    = "zone"
	isInstanceBootVolume              = "boot_volume"
	isInstanceVolAttName              = "name"
	isInstanceVolAttVolume            = "volume"
	isInstanceVolAttVolAutoDelete     = "auto_delete"
	isInstanceVolAttVolCapacity       = "capacity"
	isInstanceVolAttVolIops           = "iops"
	isInstanceVolAttVolName           = "name"
	isInstanceVolAttVolBillingTerm    = "billing_term"
	isInstanceVolAttVolEncryptionKey  = "encryption_key"
	isInstanceVolAttVolType           = "type"
	isInstanceVolAttVolProfile        = "profile"
	isInstanceImage                   = "image"
	isInstanceCPU                     = "vcpu"
	isInstanceCPUArch                 = "architecture"
	isInstanceCPUCores                = "cores"
	isInstanceCPUCount                = "count"
	isInstanceGpu                     = "gpu"
	isInstanceGpuCores                = "cores"
	isInstanceGpuCount                = "count"
	isInstanceGpuManufacturer         = "manufacturer"
	isInstanceGpuMemory               = "memory"
	isInstanceGpuModel                = "model"
	isInstanceMemory                  = "memory"
	isInstanceStatus                  = "status"
	isInstanceGeneration              = "generation"

	isInstanceProvisioning     = "provisioning"
	isInstanceProvisioningDone = "done"
	isInstanceAvailable        = "available"
	isInstanceDeleting         = "deleting"
	isInstanceDeleteDone       = "done"
	isInstanceFailed           = "failed"

	isInstanceActionStatusStopping = "stopping"
	isInstanceActionStatusStopped  = "stopped"
	isInstanceStatusPending        = "pending"
	isInstanceStatusRunning        = "running"
	isInstanceStatusFailed         = "failed"

	isInstanceBootName       = "name"
	isInstanceBootSize       = "size"
	isInstanceBootIOPS       = "iops"
	isInstanceBootEncryption = "encryption"
	isInstanceBootProfile    = "profile"

	isInstanceVolumeAttachments = "volume_attachments"
	isInstanceVolumeAttaching   = "attaching"
	isInstanceVolumeAttached    = "attached"
	isInstanceVolumeDetaching   = "detaching"
	isInstanceResourceGroup     = "resource_group"
)

func resourceIBMISInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMisInstanceCreate,
		Read:     resourceIBMisInstanceRead,
		Update:   resourceIBMisInstanceUpdate,
		Delete:   resourceIBMisInstanceDelete,
		Exists:   resourceIBMisInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isInstanceName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
			},

			isInstanceVPC: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceZone: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceProfile: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceKeys: {
				Type:             schema.TypeSet,
				Required:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
			},

			isInstanceTags: {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      resourceIBMVPCHash,
			},

			isInstanceVolumeAttachments: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_crn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isInstancePrimaryNetworkInterface: {
				Type:     schema.TypeList,
				MinItems: 1,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceNicName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceNicPortSpeed: {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: applyOnce,
							Deprecated:       "This field is deprected",
						},
						isInstanceNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceNicSubnet: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			isInstanceNetworkInterfaces: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceNicName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceNicSubnet: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			isInstanceGeneration: {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: applyOnce,
				ValidateFunc:     validateGeneration,
				Removed:          "This field is removed",
			},

			isInstanceUserData: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},

			isInstanceImage: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceBootVolume: {
				Type:             schema.TypeList,
				DiffSuppressFunc: applyOnce,
				Optional:         true,
				Computed:         true,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceBootName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceBootEncryption: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceBootSize: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceBootIOPS: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceBootProfile: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isInstanceVolumes: {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			isInstanceResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			isInstanceCPU: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceCPUArch: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceCPUCount: {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			isInstanceGpu: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceGpuCores: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceGpuCount: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceGpuMemory: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceGpuManufacturer: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceGpuModel: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isInstanceMemory: {
				Type:     schema.TypeInt,
				Computed: true,
			},

			isInstanceStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMisInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	profile := d.Get(isInstanceProfile).(string)
	var body = computec.PostInstancesBody{
		Name: d.Get(isInstanceName).(string),
		Vpc: &computec.PostInstancesParamsBodyVpc{
			ID: strfmt.UUID(d.Get(isInstanceVPC).(string)),
		},
		Zone: &models.NameReference{
			Name: d.Get(isInstanceZone).(string),
		},
		Profile: &computec.PostInstancesParamsBodyProfile{
			Name: profile,
		},
		Image: &computec.PostInstancesParamsBodyImage{
			ID: strfmt.UUID(d.Get(isInstanceImage).(string)),
		},
	}

	if boot, ok := d.GetOk(isInstanceBootVolume); ok {
		bootvol := boot.([]interface{})[0].(map[string]interface{})
		template := &computec.PostInstancesParamsBodyBootVolumeAttachmentVolume{}
		name, ok := bootvol[isInstanceBootName]
		if ok {
			template.Name = name.(string)
		}
		enc, ok := bootvol[isInstanceBootEncryption]
		if ok && enc.(string) != "" {
			template.EncryptionKey = &computec.PostInstancesParamsBodyBootVolumeAttachmentVolumeEncryptionKey{
				Crn: enc.(string),
			}
		}
		template.Capacity = 100
		template.Profile = &computec.PostInstancesParamsBodyBootVolumeAttachmentVolumeProfile{
			Name: "general-purpose",
		}

		body.BootVolumeAttachment = &computec.PostInstancesParamsBodyBootVolumeAttachment{
			DeleteVolumeOnInstanceDelete: true,
			Volume:                       template,
		}
	}

	// implement boovol, nics, vols

	if primnicintf, ok := d.GetOk(isInstancePrimaryNetworkInterface); ok {
		primnic := primnicintf.([]interface{})[0].(map[string]interface{})
		subnetintf, _ := primnic[isInstanceNicSubnet]
		var primnicobj = computec.PostInstancesParamsBodyPrimaryNetworkInterface{}
		primnicobj.Subnet = &computec.PostInstancesParamsBodyPrimaryNetworkInterfaceSubnet{
			ID: strfmt.UUID(subnetintf.(string)),
		}
		name, ok := primnic[isInstanceNicName]
		if ok {
			primnicobj.Name = name.(string)
		}
		secgrpintf, ok := primnic[isInstanceNicSecurityGroups]
		if ok {
			secgrpSet := secgrpintf.(*schema.Set)
			if secgrpSet.Len() != 0 {
				var secgrpobjs = make([]*computec.PostInstancesParamsBodyPrimaryNetworkInterfaceSecurityGroupsItems0, secgrpSet.Len())
				for i, secgrpIntf := range secgrpSet.List() {
					secgrpobjs[i] = &computec.PostInstancesParamsBodyPrimaryNetworkInterfaceSecurityGroupsItems0{
						ID: strfmt.UUID(secgrpIntf.(string)),
					}
				}
				primnicobj.SecurityGroups = secgrpobjs
			}
		}

		body.PrimaryNetworkInterface = &primnicobj
	}

	if nicsintf, ok := d.GetOk(isInstanceNetworkInterfaces); ok {
		nics := nicsintf.([]interface{})
		var intfs []*computec.NetworkInterfacesItems0
		for _, resource := range nics {
			nic := resource.(map[string]interface{})
			nwInterface := &computec.NetworkInterfacesItems0{}
			subnetintf, _ := nic[isInstanceNicSubnet]
			nwInterface.Subnet = &computec.NetworkInterfacesItems0Subnet{
				ID: strfmt.UUID(subnetintf.(string)),
			}
			name, ok := nic[isInstanceNicName]
			if ok && name.(string) != "" {
				nwInterface.Name = name.(string)
			}
			secgrpintf, ok := nic[isInstanceNicSecurityGroups]
			if ok {
				secgrpSet := secgrpintf.(*schema.Set)
				if secgrpSet.Len() != 0 {
					var secgrpobjs = make([]*computec.NetworkInterfacesItems0SecurityGroupsItems0, secgrpSet.Len())
					for i, secgrpIntf := range secgrpSet.List() {
						secgrpobjs[i] = &computec.NetworkInterfacesItems0SecurityGroupsItems0{
							ID: strfmt.UUID(secgrpIntf.(string)),
						}
					}
					nwInterface.SecurityGroups = secgrpobjs
				}
			}

			intfs = append(intfs, nwInterface)
		}

		body.NetworkInterfaces = intfs
	}

	keySet := d.Get(isInstanceKeys).(*schema.Set)
	if keySet.Len() != 0 {
		keyobjs := make([]*computec.KeysItems0, keySet.Len())
		for i, key := range keySet.List() {
			keyobjs[i] = &computec.KeysItems0{
				ID: strfmt.UUID(key.(string)),
			}
		}
		body.Keys = keyobjs
	}

	if userdata, ok := d.GetOk(isInstanceUserData); ok {
		body.UserData = userdata.(string)
	}

	if grp, ok := d.GetOk(isInstanceResourceGroup); ok {
		body.ResourceGroup = &computec.PostInstancesParamsBodyResourceGroup{
			ID: strfmt.UUID(grp.(string)),
		}

	}

	instanceC := compute.NewInstanceClient(sess)
	instance, err := instanceC.Create(body)
	if err != nil {
		log.Printf("[DEBUG] Instance err %s", isErrorToString(err))
		return err
	}

	d.SetId(instance.ID.String())
	log.Printf("[INFO] Instance : %s", instance.ID.String())
	d.Set(isInstanceStatus, instance.Status)

	_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
	if err != nil {
		return err
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isInstanceTags); ok || v != "" {
		oldList, newList := d.GetChange(isInstanceTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, instance.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc instance (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMisInstanceUpdate(d, meta)
}

func isWaitForInstanceAvailable(instanceC *compute.InstanceClient, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for instance (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isInstanceProvisioning},
		Target:     []string{isInstanceStatusRunning},
		Refresh:    isInstanceRefreshFunc(instanceC, id, d),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceRefreshFunc(instanceC *compute.InstanceClient, id string, d *schema.ResourceData) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		instance, err := instanceC.Get(id)
		if err != nil {
			return nil, "", err
		}
		d.Set(isInstanceStatus, instance.Status)

		if instance.Status == "available" || instance.Status == "failed" || instance.Status == "running" {
			return instance, isInstanceStatusRunning, nil
		}

		return instance, isInstanceProvisioning, nil
	}
}

func resourceIBMisInstanceRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

	instance, err := instanceC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set(isInstanceName, instance.Name)
	if instance.Profile != nil {
		d.Set(isInstanceProfile, instance.Profile.Name)
	}
	cpuList := make([]map[string]interface{}, 0)
	if instance.Vcpu != nil {
		currentCPU := map[string]interface{}{}
		currentCPU[isInstanceCPUArch] = *instance.Vcpu.Architecture
		currentCPU[isInstanceCPUCount] = *instance.Vcpu.Count
		cpuList = append(cpuList, currentCPU)
	}
	d.Set(isInstanceCPU, cpuList)

	d.Set(isInstanceMemory, instance.Memory)
	gpuList := make([]map[string]interface{}, 0)
	if instance.Gpu != nil {
		currentGpu := map[string]interface{}{}
		currentGpu[isInstanceGpuManufacturer] = instance.Gpu.Manufacturer
		currentGpu[isInstanceGpuModel] = instance.Gpu.Model
		currentGpu[isInstanceGpuCores] = instance.Gpu.Cores
		currentGpu[isInstanceGpuCount] = instance.Gpu.Count
		currentGpu[isInstanceGpuMemory] = instance.Gpu.Memory
		gpuList = append(gpuList, currentGpu)

	}
	d.Set(isInstanceGpu, gpuList)

	if instance.PrimaryNetworkInterface != nil {
		primaryNicList := make([]map[string]interface{}, 0)
		currentPrimNic := map[string]interface{}{}
		currentPrimNic["id"] = instance.PrimaryNetworkInterface.ID.String()
		currentPrimNic[isInstanceNicName] = instance.PrimaryNetworkInterface.Name
		currentPrimNic[isInstanceNicPrimaryIpv4Address] = instance.PrimaryNetworkInterface.PrimaryIPV4Address
		insnic, err := instanceC.GetInterface(d.Id(), instance.PrimaryNetworkInterface.ID.String())
		if err != nil {
			return err
		}
		currentPrimNic[isInstanceNicSubnet] = insnic.Subnet.ID
		if len(insnic.SecurityGroups) != 0 {
			secgrpList := []string{}
			for i := 0; i < len(insnic.SecurityGroups); i++ {
				secgrpList = append(secgrpList, string((insnic.SecurityGroups[i].ID)))
			}
			currentPrimNic[isInstanceNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
		}

		primaryNicList = append(primaryNicList, currentPrimNic)
		d.Set(isInstancePrimaryNetworkInterface, primaryNicList)
	}

	if instance.NetworkInterfaces != nil {
		interfacesList := make([]map[string]interface{}, 0)
		for _, intfc := range instance.NetworkInterfaces {
			if intfc.ID != instance.PrimaryNetworkInterface.ID {
				currentNic := map[string]interface{}{}
				currentNic["id"] = intfc.ID
				currentNic[isInstanceNicName] = intfc.Name
				currentNic[isInstanceNicPrimaryIpv4Address] = intfc.PrimaryIPV4Address
				insnic, err := instanceC.GetInterface(d.Id(), intfc.ID.String())
				if err != nil {
					return err
				}
				currentNic[isInstanceNicSubnet] = insnic.Subnet.ID
				if len(insnic.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(insnic.SecurityGroups); i++ {
						secgrpList = append(secgrpList, string((insnic.SecurityGroups[i].ID)))
					}
					currentNic[isInstanceNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
				}
				interfacesList = append(interfacesList, currentNic)

			}

		}

		d.Set(isInstanceNetworkInterfaces, interfacesList)
	}

	if instance.Image != nil {
		d.Set(isInstanceImage, instance.Image.ID.String())
	}

	d.Set(isInstanceStatus, instance.Status)
	d.Set(isInstanceVPC, instance.Vpc.ID.String())
	d.Set(isInstanceZone, instance.Zone.Name)

	var volumes []string
	volumes = make([]string, 0)
	if instance.VolumeAttachments != nil {
		for _, volume := range instance.VolumeAttachments {
			if volume != nil && volume.Volume.ID != instance.BootVolumeAttachment.Volume.ID {
				volumes = append(volumes, string(volume.Volume.ID))
			}
		}
	}
	d.Set(isInstanceVolumes, newStringSet(schema.HashString, volumes))
	if instance.VolumeAttachments != nil {
		volList := make([]map[string]interface{}, 0)
		vol := map[string]interface{}{}
		for _, volume := range instance.VolumeAttachments {
			if volume != nil {
				vol["id"] = volume.ID
				vol["volume_id"] = volume.Volume.ID
				vol["name"] = volume.Name
				vol["volume_name"] = volume.Volume.Name
				vol["volume_crn"] = volume.Volume.Crn
				volList = append(volList, vol)
			}
		}
		d.Set(isInstanceVolumeAttachments, volList)
	}
	if instance.BootVolumeAttachment != nil {
		bootVolList := make([]map[string]interface{}, 0)
		bootVol := map[string]interface{}{}
		bootVol[isInstanceBootName] = instance.BootVolumeAttachment.Name
		stg := storage.NewStorageClient(sess)
		volId := instance.BootVolumeAttachment.Volume.ID.String()
		vol, err := stg.Get(volId)
		if err != nil {
			return fmt.Errorf("Error while retrieving boot volume %s for instance %s: %v", volId, d.Id(), err)
		}
		if vol.EncryptionKey != nil {
			bootVol[isInstanceBootEncryption] = vol.EncryptionKey.Crn
		}
		bootVol[isInstanceBootSize] = vol.Capacity
		bootVol[isInstanceBootIOPS] = vol.Iops
		bootVol[isInstanceBootProfile] = vol.Profile.Name
		bootVolList = append(bootVolList, bootVol)

		d.Set(isInstanceBootVolume, bootVolList)
	}
	tags, err := GetTagsUsingCRN(meta, instance.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc Instance (%s) tags: %s", d.Id(), err)
	}
	d.Set(isInstanceTags, tags)
	d.Set(isInstanceResourceGroup, instance.ResourceGroup.ID)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/compute/vs")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/compute/vs")
	}
	d.Set(ResourceName, instance.Name)
	d.Set(ResourceCRN, instance.Crn)
	d.Set(ResourceStatus, instance.Status)
	if instance.ResourceGroup != nil {
		d.Set(ResourceGroupName, instance.ResourceGroup.Name)
	}
	return nil
}

func resourceIBMisInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

	instance, err := instanceC.Get(d.Id())
	if err != nil {
		return err
	}

	if d.HasChange(isInstanceVolumes) {
		ovs, nvs := d.GetChange(isInstanceVolumes)
		ov := ovs.(*schema.Set)
		nv := nvs.(*schema.Set)

		remove := expandStringList(ov.Difference(nv).List())
		add := expandStringList(nv.Difference(ov).List())

		if len(add) > 0 {
			for i := range add {
				vol, err := instanceC.AttachVolume(d.Id(), add[i], "", "")
				if err != nil {
					return fmt.Errorf("Error while attaching volume %q for instance %s: %q", add[i], d.Id(), err)
				}
				_, err = isWaitForInstanceVolumeAttached(instanceC, d.Id(), vol.ID.String(), d)
				if err != nil {
					return err
				}
			}

		}
		if len(remove) > 0 {
			for i := range remove {
				vols, err := instanceC.ListVolAttachments(d.Id())
				if err != nil {
					return err
				}
				for _, vol := range vols {
					if vol.Volume.ID.String() == remove[i] {
						err := instanceC.DeleteVolAttachment(d.Id(), vol.ID.String())
						if err != nil {
							return fmt.Errorf("Error while removing volume %q for instance %s: %q", remove[i], d.Id(), err)
						}
						_, err = isWaitForInstanceVolumeDetached(vol.ID.String(), d, meta)
						if err != nil {
							return err
						}
						break
					}

				}

			}

		}
	}

	if d.HasChange("primary_network_interface.0.security_groups") && !d.IsNewResource() {
		ovs, nvs := d.GetChange("primary_network_interface.0.security_groups")
		ov := ovs.(*schema.Set)
		nv := nvs.(*schema.Set)
		remove := expandStringList(ov.Difference(nv).List())
		add := expandStringList(nv.Difference(ov).List())
		if len(add) > 0 {
			sgC := network.NewSecurityGroupClient(sess)
			networkID := d.Get("primary_network_interface.0.id").(string)
			for i := range add {
				_, err := sgC.AddNetworkInterface(add[i], networkID)
				if err != nil {
					return fmt.Errorf("Error while attaching security group %q for primary network interface of instance %s: %q", add[i], d.Id(), err)
				}
				_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
				if err != nil {
					return err
				}
			}

		}
		if len(remove) > 0 {
			sgC := network.NewSecurityGroupClient(sess)
			networkID := d.Get("primary_network_interface.0.id").(string)
			for i := range remove {
				err := sgC.DeleteNetworkInterface(remove[i], networkID)
				if err != nil {
					return fmt.Errorf("Error while removing security group %q for primary network interface of instance %s: %q", remove[i], d.Id(), err)
				}
				_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
				if err != nil {
					return err
				}
			}

		}

	}

	if d.HasChange("primary_network_interface.0.name") && !d.IsNewResource() {
		newName := d.Get("primary_network_interface.0.name").(string)
		networkID := d.Get("primary_network_interface.0.id").(string)
		_, err := instanceC.UpdateInterface(d.Id(), networkID, newName, 0)
		if err != nil {
			return fmt.Errorf("Error while updating name %s for primary network interface of instance %s: %q", newName, d.Id(), err)
		}
		_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
		if err != nil {
			return err
		}
	}

	if d.HasChange(isInstanceNetworkInterfaces) && !d.IsNewResource() {
		nics := d.Get(isInstanceNetworkInterfaces).([]interface{})
		for i, _ := range nics {
			securitygrpKey := fmt.Sprintf("network_interfaces.%d.security_groups", i)
			networkNameKey := fmt.Sprintf("network_interfaces.%d.name", i)
			if d.HasChange(securitygrpKey) {
				ovs, nvs := d.GetChange(securitygrpKey)
				ov := ovs.(*schema.Set)
				nv := nvs.(*schema.Set)
				remove := expandStringList(ov.Difference(nv).List())
				add := expandStringList(nv.Difference(ov).List())
				if len(add) > 0 {
					sgC := network.NewSecurityGroupClient(sess)
					networkIDKey := fmt.Sprintf("network_interfaces.%d.id", i)
					networkID := d.Get(networkIDKey).(string)
					for i := range add {
						_, err := sgC.AddNetworkInterface(add[i], networkID)
						if err != nil {
							return fmt.Errorf("Error while attaching security group %q for network interface %s of instance %s: %q", add[i], networkID, d.Id(), err)
						}
						_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
						if err != nil {
							return err
						}
					}

				}
				if len(remove) > 0 {
					sgC := network.NewSecurityGroupClient(sess)
					networkIDKey := fmt.Sprintf("network_interfaces.%d.id", i)
					networkID := d.Get(networkIDKey).(string)
					for i := range remove {
						err := sgC.DeleteNetworkInterface(remove[i], networkID)
						if err != nil {
							return fmt.Errorf("Error while removing security group %q for network interface %s of instance %s: %q", remove[i], networkID, d.Id(), err)
						}
						_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
						if err != nil {
							return err
						}
					}
				}

			}

			if d.HasChange(networkNameKey) {
				newName := d.Get(networkNameKey).(string)
				networkIDKey := fmt.Sprintf("network_interfaces.%d.id", i)
				networkID := d.Get(networkIDKey).(string)
				_, err := instanceC.UpdateInterface(d.Id(), networkID, newName, 0)
				if err != nil {
					return fmt.Errorf("Error while updating name %s for network interface %s of instance %s: %q", newName, networkID, d.Id(), err)
				}
				_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
				if err != nil {
					return err
				}
			}
		}

	}

	if d.HasChange(isInstanceName) {
		name := d.Get(isInstanceName).(string)
		_, err = instanceC.Update(d.Id(), name, "")
		if err != nil {
			return err
		}
	}
	if d.HasChange(isInstanceTags) {
		oldList, newList := d.GetChange(isInstanceTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, instance.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc Instance (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMisInstanceRead(d, meta)
}

func resourceIBMisInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)
	_, err = instanceC.CreateAction(d.Id(), "stop")
	if err != nil {
		return err
	}
	_, err = isWaitForInstanceActionStop(d, meta)
	if err != nil {
		return err
	}
	vols, err := instanceC.ListVolAttachments(d.Id())
	if err != nil {
		return err
	}
	for _, vol := range vols {
		if vol.Type == "data" {
			err := instanceC.DeleteVolAttachment(d.Id(), vol.ID.String())
			if err != nil {
				return fmt.Errorf("Error while removing volume %q for instance %s: %q", vol.ID.String(), d.Id(), err)
			}
			_, err = isWaitForInstanceVolumeDetached(vol.ID.String(), d, meta)
			if err != nil {
				return err
			}
			break
		}
	}
	err = instanceC.Delete(d.Id())
	if err != nil {
		return err
	}
	_, err = isWaitForInstanceDelete(d, meta)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMisInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	instanceC := compute.NewInstanceClient(sess)

	_, err = instanceC.Get(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func isWaitForInstanceDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	instanceC := compute.NewInstanceClient(sess)

	stateConf := &resource.StateChangeConf{
		Pending: []string{isInstanceDeleting, isInstanceAvailable},
		Target:  []string{isInstanceDeleteDone},
		Refresh: func() (interface{}, string, error) {
			instance, err := instanceC.Get(d.Id()) //Only in case there's a rias error with code "not found", resource is deleted, all other cases we keep attempting to delete
			if err != nil {
				iserror, ok := err.(iserrors.RiaasError)
				if ok {
					if len(iserror.Payload.Errors) == 1 &&
						iserror.Payload.Errors[0].Code == "not_found" {
						return instance, isInstanceDeleteDone, nil
					}
				}
				return instance, "", err
			}
			if instance.Status == isInstanceFailed {
				return instance, instance.Status, fmt.Errorf("The  instance %s failed to delete: %v", d.Id(), err)
			}
			return instance, isInstanceDeleting, nil
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isWaitForInstanceActionStop(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	instanceC := compute.NewInstanceClient(sess)

	stateConf := &resource.StateChangeConf{
		Pending: []string{isInstanceStatusRunning, isInstanceStatusPending, isInstanceActionStatusStopping},
		Target:  []string{isInstanceActionStatusStopped, isInstanceStatusFailed},
		Refresh: func() (interface{}, string, error) {
			instance, err := instanceC.Get(d.Id())
			if err != nil {
				return nil, "", err
			}
			return instance, instance.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isWaitForInstanceVolumeAttached(instanceC *compute.InstanceClient, id, volID string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for instance volume (%s) to be attched.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isInstanceVolumeAttaching},
		Target:     []string{isInstanceVolumeAttached},
		Refresh:    isInstanceVolumeRefreshFunc(instanceC, id, volID),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceVolumeRefreshFunc(instanceC *compute.InstanceClient, id, volID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := instanceC.GetVolAttachment(id, volID)
		if err != nil {
			return nil, "", err
		}

		if vol.Status == isInstanceVolumeAttached {
			return vol, isInstanceVolumeAttached, nil
		}

		return vol, isInstanceVolumeAttaching, nil
	}
}

func isWaitForInstanceVolumeDetached(volID string, d *schema.ResourceData, meta interface{}) (interface{}, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	instanceC := compute.NewInstanceClient(sess)

	stateConf := &resource.StateChangeConf{
		Pending: []string{isInstanceVolumeAttached, isInstanceVolumeDetaching},
		Target:  []string{isInstanceDeleteDone},
		Refresh: func() (interface{}, string, error) {
			vol, err := instanceC.GetVolAttachment(d.Id(), volID)
			if err != nil {
				iserror, ok := err.(iserrors.RiaasError)
				if ok {
					if len(iserror.Payload.Errors) == 1 &&
						iserror.Payload.Errors[0].Code == "not_found" {
						return vol, isInstanceDeleteDone, nil
					}
				}
				return vol, "", err
			}
			if vol.Status == isInstanceFailed {
				return vol, vol.Status, fmt.Errorf("The instance %s failed to detach volume %s: %v", d.Id(), volID, err)
			}
			return vol, isInstanceVolumeDetaching, nil
		},
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}
