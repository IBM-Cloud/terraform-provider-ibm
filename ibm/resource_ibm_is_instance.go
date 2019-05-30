package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/models"
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
							Required:         true,
							DiffSuppressFunc: applyOnce,
						},
						isInstanceNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
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
			},

			isInstanceUserData: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},

			isInstanceImage: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				ConflictsWith: []string{isInstanceBootVolume},
			},

			isInstanceBootVolume: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isInstanceImage},
			},

			isInstanceVolumes: {
				Type:             schema.TypeSet,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
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
		Vpc: &models.ResourceLocator{
			ID: strfmt.UUID(d.Get(isInstanceVPC).(string)),
		},
		Zone: &models.ZoneReference{
			Name: d.Get(isInstanceZone).(string),
		},
		Profile: &models.NameLocator{
			Name: profile,
		},
		Flavor: &models.NameLocator{
			Name: profile,
		},
		Generation: models.Generation(d.Get(isInstanceGeneration).(string)),
	}

	var imageID, bootvol string

	if img, ok := d.GetOk(isInstanceImage); ok {
		imageID = img.(string)
		body.Image = &models.ResourceLocator{
			ID: strfmt.UUID(imageID),
		}
	}

	if boot, ok := d.GetOk(isInstanceBootVolume); ok {
		bootvol = boot.(string)

		template := &models.VolumeAttachmentTemplateVolume{}
		template.ID = strfmt.UUID(bootvol)

		body.BootVolumeAttachment = &models.VolumeAttachmentTemplate{

			Volume: template,
		}
	}

	if imageID == "" && bootvol == "" {
		return fmt.Errorf("%s or %s need to be provided", isInstanceImage, isInstanceBootVolume)
	}

	// implement boovol, nics, vols

	if primnicintf, ok := d.GetOk(isInstancePrimaryNetworkInterface); ok {
		primnic := primnicintf.([]interface{})[0].(map[string]interface{})
		portspeedintf := (primnic[isInstanceNicPortSpeed].(int))
		subnetintf, _ := primnic[isInstanceNicSubnet]
		var primnicobj = models.PrimaryNetworkInterfaceTemplate{}
		primnicobj.Subnet = &models.ResourceLocator{
			ID: strfmt.UUID(subnetintf.(string)),
		}

		primnicobj.PortSpeed = int64(portspeedintf)
		secgrpintf, ok := primnic[isInstanceNicSecurityGroups]
		if ok {
			secgrpSet := secgrpintf.(*schema.Set)
			if secgrpSet.Len() != 0 {
				var secgrpobjs = make([]*models.ResourceLocator, secgrpSet.Len())
				for i, secgrpIntf := range secgrpSet.List() {
					secgrpobjs[i] = &models.ResourceLocator{
						ID: strfmt.UUID(secgrpIntf.(string)),
					}
				}
				primnicobj.SecurityGroups = secgrpobjs
			}
		}

		body.PrimaryNetworkInterface = &primnicobj
	}

	keySet := d.Get(isInstanceKeys).(*schema.Set)
	if keySet.Len() != 0 {
		keyobjs := make([]*models.KeyLocator, keySet.Len())
		for i, key := range keySet.List() {
			keyobjs[i] = &models.KeyLocator{
				ID: strfmt.UUID(key.(string)),
			}
		}
		body.Keys = keyobjs
	}

	volSet := d.Get(isInstanceVolumes).(*schema.Set)
	if volSet.Len() != 0 {
		volobjs := make([]*models.VolumeAttachmentTemplate, volSet.Len())
		for i, vol := range volSet.List() {

			template := &models.VolumeAttachmentTemplateVolume{}
			template.ID = strfmt.UUID(vol.(string))

			volobjs[i] = &models.VolumeAttachmentTemplate{
				Volume: template,
			}
		}
		body.VolumeAttachments = volobjs
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

	return resourceIBMisInstanceRead(d, meta)
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

	primaryNicList := make([]map[string]interface{}, 0)
	if instance.PrimaryNetworkInterface != nil {
		currentPrimNic := map[string]interface{}{}
		currentPrimNic["id"] = instance.PrimaryNetworkInterface.ID.String()
		currentPrimNic[isInstanceNicName] = instance.PrimaryNetworkInterface.Name
		currentPrimNic[isInstanceNicPrimaryIpv4Address] = instance.PrimaryNetworkInterface.PrimaryIPV4Address
		primaryNicList = append(primaryNicList, currentPrimNic)
	}
	d.Set(isInstancePrimaryNetworkInterface, primaryNicList)

	if instance.Image != nil {
		d.Set(isInstanceImage, instance.Image.ID.String())
	}

	if instance.BootVolumeAttachment != nil {
		d.Set(isInstanceBootVolume, instance.BootVolumeAttachment.ID.String())

	}

	d.Set(isInstanceStatus, instance.Status)
	d.Set(isInstanceVPC, instance.Vpc.ID.String())
	d.Set(isInstanceZone, instance.Zone.Name)

	var volumes []string
	volumes = make([]string, len(instance.VolumeAttachments), len(instance.VolumeAttachments))
	if instance.VolumeAttachments != nil {
		for i := 0; i < len(instance.VolumeAttachments); i++ {
			if instance.VolumeAttachments[i] != nil {
				volumes[i] = instance.VolumeAttachments[i].ID.String()
			}
		}
	}
	d.Set(isInstanceVolumes, volumes)

	return nil
}

func resourceIBMisInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

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
			instance, err := instanceC.Get(d.Id())
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
			return instance, instance.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}
