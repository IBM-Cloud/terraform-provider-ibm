package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

const (
	isInstanceName                    = "name"
	isInstanceKeys                    = "keys"
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
	isInstanceCPU                     = "cpu"
	isInstanceCPUArch                 = "architecture"
	isInstanceCPUCores                = "cores"
	isInstanceCPUFrequency            = "frequency"
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

	isInstanceBootName       = "name"
	isInstanceBootSize       = "size"
	isInstanceBootIOPS       = "iops"
	isInstanceBootEncryption = "encryption"
	isInstanceBootProfile    = "profile"
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

		Schema: map[string]*schema.Schema{
			isInstanceName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
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
				ForceNew: false,
				Required: true,
			},

			isInstanceKeys: {
				Type:             schema.TypeSet,
				Required:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
			},

			isInstancePrimaryNetworkInterface: {
				Type:             schema.TypeList,
				MinItems:         1,
				MaxItems:         1,
				Required:         true,
				DiffSuppressFunc: applyOnce,
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
						},
					},
				},
			},

			isInstanceNetworkInterfaces: {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: applyOnce,
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
						},
					},
				},
			},

			isInstanceGeneration: {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "gc",
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
				Optional: true,
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

			isInstanceCPU: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceCPUArch: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceCPUCores: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceCPUFrequency: {
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
		},
	}
}

func resourceIBMisInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	profile := d.Get(isInstanceProfile).(string)
	var body = &models.PostInstancesParamsBody{
		Name: d.Get(isInstanceName).(string),
		Vpc: &models.PostInstancesParamsBodyVpc{
			ID: strfmt.UUID(d.Get(isInstanceVPC).(string)),
		},
		Zone: &models.NameReference{
			Name: d.Get(isInstanceZone).(string),
		},
		Profile: &models.PostInstancesParamsBodyProfile{
			Name: profile,
		},
	}

	if img, ok := d.GetOk(isInstanceImage); ok {
		body.Image = &models.PostInstancesParamsBodyImage{
			ID: strfmt.UUID(img.(string)),
		}
	}

	if boot, ok := d.GetOk(isInstanceBootVolume); ok {
		bootvol := boot.([]interface{})[0].(map[string]interface{})
		template := &models.PostInstancesParamsBodyBootVolumeAttachmentVolume{}
		name, ok := bootvol[isInstanceBootName]
		if ok {
			template.Name = name.(string)
		}
		enc, ok := bootvol[isInstanceBootEncryption]
		if ok && enc.(string) != "" {
			template.EncryptionKey = &models.PostInstancesParamsBodyBootVolumeAttachmentVolumeEncryptionKey{
				Crn: enc.(string),
			}
		}
		body.BootVolumeAttachment = &models.PostInstancesParamsBodyBootVolumeAttachment{

			Volume: template,
		}
	}

	// implement boovol, nics, vols

	if primnicintf, ok := d.GetOk(isInstancePrimaryNetworkInterface); ok {
		primnic := primnicintf.([]interface{})[0].(map[string]interface{})
		subnetintf, _ := primnic[isInstanceNicSubnet]
		var primnicobj = models.PostInstancesParamsBodyPrimaryNetworkInterface{}
		primnicobj.Subnet = &models.PostInstancesParamsBodyPrimaryNetworkInterfaceSubnet{
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
				var secgrpobjs = make([]*models.PostInstancesParamsBodyPrimaryNetworkInterfaceSecurityGroupsItems, secgrpSet.Len())
				for i, secgrpIntf := range secgrpSet.List() {
					secgrpobjs[i] = &models.PostInstancesParamsBodyPrimaryNetworkInterfaceSecurityGroupsItems{
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
		var intfs []*models.PostInstancesParamsBodyNetworkInterfacesItems
		for _, resource := range nics {
			nic := resource.(map[string]interface{})
			nwInterface := &models.PostInstancesParamsBodyNetworkInterfacesItems{}
			subnetintf, _ := nic[isInstanceNicSubnet]
			nwInterface.Subnet = &models.PostInstancesParamsBodyNetworkInterfacesItemsSubnet{
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
					var secgrpobjs = make([]*models.PostInstancesParamsBodyNetworkInterfacesItemsSecurityGroupsItems, secgrpSet.Len())
					for i, secgrpIntf := range secgrpSet.List() {
						secgrpobjs[i] = &models.PostInstancesParamsBodyNetworkInterfacesItemsSecurityGroupsItems{
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
		keyobjs := make([]*models.PostInstancesParamsBodyKeysItems, keySet.Len())
		for i, key := range keySet.List() {
			keyobjs[i] = &models.PostInstancesParamsBodyKeysItems{
				ID: strfmt.UUID(key.(string)),
			}
		}
		body.Keys = keyobjs
	}

	if userdata, ok := d.GetOk(isInstanceUserData); ok {
		body.UserData = userdata.(string)
	}

	instanceC := compute.NewInstanceClient(sess)
	instance, err := instanceC.Create(body)
	if err != nil {
		log.Printf("[DEBUG] Instance err %s", isErrorToString(err))
		return err
	}

	d.SetId(instance.ID.String())
	log.Printf("[INFO] Instance : %s", instance.ID.String())

	_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
	if err != nil {
		return err
	}

	return resourceIBMisInstanceUpdate(d, meta)
}

func isWaitForInstanceAvailable(instanceC *compute.InstanceClient, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for instance (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isInstanceProvisioning},
		Target:     []string{isInstanceProvisioningDone},
		Refresh:    isInstanceRefreshFunc(instanceC, id),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceRefreshFunc(instanceC *compute.InstanceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		instance, err := instanceC.Get(id)
		if err != nil {
			return nil, "", err
		}

		if instance.Status == "available" || instance.Status == "failed" || instance.Status == "running" {
			return instance, isInstanceProvisioningDone, nil
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
	if instance.CPU != nil {
		currentCPU := map[string]interface{}{}
		currentCPU[isInstanceCPUArch] = instance.CPU.Architecture
		currentCPU[isInstanceCPUCores] = instance.CPU.Cores
		currentCPU[isInstanceCPUFrequency] = instance.CPU.Frequency
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

	if instance.BootVolumeAttachment != nil {
		bootVol := map[string]interface{}{}
		bootVol[isInstanceBootName] = instance.BootVolumeAttachment.Name
		stg := storage.NewStorageClient(sess)
		vol, err := stg.Get(instance.BootVolumeAttachment.Volume.ID.String())
		if err != nil {
			return err
		}
		if vol.EncryptionKey != nil {
			bootVol[isInstanceBootEncryption] = vol.EncryptionKey.Crn
		}
		bootVol[isInstanceBootSize] = vol.Capacity
		bootVol[isInstanceBootIOPS] = vol.Iops
		bootVol[isInstanceBootProfile] = vol.Profile
	}

	return nil
}

func resourceIBMisInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

	if d.HasChange(isInstanceVolumes) {
		ovs, nvs := d.GetChange(isInstanceVolumes)
		ov := ovs.(*schema.Set)
		nv := nvs.(*schema.Set)

		remove := expandStringList(ov.Difference(nv).List())
		add := expandStringList(nv.Difference(ov).List())

		if len(add) > 0 {
			for i := range add {
				_, err := instanceC.AttachVolume(d.Id(), add[i], "", "")
				if err != nil {
					return fmt.Errorf("Error while attaching volume %q for instance %s: %q", add[i], d.Id(), err)
				}
				_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
				if err != nil {
					return err
				}
			}

		}
		if len(remove) > 0 {
			for i := range remove {
				if remove[i] != "" {
					err := instanceC.DeleteVolAttachment(d.Id(), remove[i])
					if err != nil {
						return fmt.Errorf("Error while removing volume %q for instance %s: %q", remove[i], d.Id(), err)
					}
					_, err = isWaitForInstanceAvailable(instanceC, d.Id(), d)
					if err != nil {
						return err
					}
				}

			}

		}
	}

	name := ""
	profile := ""
	if d.HasChange(isInstanceName) {
		name = d.Get(isInstanceName).(string)
	}
	if d.HasChange(isInstanceProfile) {
		profile = d.Get(isInstanceProfile).(string)
	}

	_, err = instanceC.Update(d.Id(), name, profile)
	if err != nil {
		return err
	}

	return resourceIBMisInstanceRead(d, meta)
}

func resourceIBMisInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)
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
