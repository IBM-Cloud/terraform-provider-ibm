package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isInstanceTemplateName                    = "name"
	isInstanceTemplateKeys                    = "keys"
	isInstanceTemplateNetworkInterfaces       = "network_interfaces"
	isInstanceTemplatePrimaryNetworkInterface = "primary_network_interface"
	isInstanceTemplateNicName                 = "name"
	isInstanceTemplateProfile                 = "profile"
	isInstanceTemplateNicPortSpeed            = "port_speed"
	//isInstanceTemplateNicAllowIpSpoofing      = "allow_ip_spoofing"
	isInstanceTemplateNicPrimaryIpv4Address  = "primary_ipv4_address"
	isInstanceTemplateNicPrimaryIpv6Address  = "primary_ipv6_address"
	isInstanceTemplateNicSecondaryAddress    = "secondary_addresses"
	isInstanceTemplateNicSecurityGroups      = "security_groups"
	isInstanceTemplateNicSubnet              = "subnet"
	isInstanceTemplateNicFloatingIPs         = "floating_ips"
	isInstanceTemplateUserData               = "user_data"
	isInstanceTemplateVPC                    = "vpc"
	isInstanceTemplateZone                   = "zone"
	isInstanceTemplateBootVolume             = "boot_volume"
	isInstanceTemplateVolAttName             = "name"
	isInstanceTemplateVolAttVolume           = "volume"
	isInstanceTemplateVolAttVolAutoDelete    = "auto_delete"
	isInstanceTemplateVolAttVolCapacity      = "capacity"
	isInstanceTemplateVolAttVolIops          = "iops"
	isInstanceTemplateVolAttVolName          = "name"
	isInstanceTemplateVolAttVolBillingTerm   = "billing_term"
	isInstanceTemplateVolAttVolEncryptionKey = "encryption_key"
	isInstanceTemplateVolAttVolType          = "type"
	isInstanceTemplateVolAttVolProfile       = "profile"
	isInstanceTemplateImage                  = "image"
	isInstanceTemplateGeneration             = "generation"

	isInstanceTemplateProvisioning     = "provisioning"
	isInstanceTemplateProvisioningDone = "done"
	isInstanceTemplateAvailable        = "available"
	isInstanceTemplateDeleting         = "deleting"
	isInstanceTemplateDeleteDone       = "done"
	isInstanceTemplateFailed           = "failed"

	isInstanceTemplateBootName       = "name"
	isInstanceTemplateBootSize       = "size"
	isInstanceTemplateBootIOPS       = "iops"
	isInstanceTemplateBootEncryption = "encryption"
	isInstanceTemplateBootProfile    = "profile"

	isInstanceTemplateVolumeAttachments = "volume_attachments"
	isInstanceTemplateVolumeAttaching   = "attaching"
	isInstanceTemplateVolumeAttached    = "attached"
	isInstanceTemplateVolumeDetaching   = "detaching"
	isInstanceTemplateResourceGroup     = "resource_group"
	isInstanceTemplateDeleteVolume      = "delete_volume_on_instance_delete"
)

func resourceIBMISInstanceTemplate() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMisInstanceTemplateCreate,
		Read:     resourceIBMisInstanceTemplateRead,
		Update:   resourceIBMisInstanceTemplateUpdate,
		Delete:   resourceIBMisInstanceTemplateDelete,
		Exists:   resourceIBMisInstanceTemplateExists,
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
			isInstanceTemplateName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
			},

			isInstanceTemplateVPC: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceTemplateZone: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceTemplateProfile: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceTemplateKeys: {
				Type:             schema.TypeSet,
				Required:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
			},

			isInstanceTemplateVolumeAttachments: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_volume_on_instance_delete": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"volume": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			isInstanceTemplatePrimaryNetworkInterface: {
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
						/*
							isInstanceTemplateNicAllowIpSpoofing: {
								Type:     schema.TypeBool,
								Optional: true,
								Default:  false,
							},
						*/
						isInstanceTemplateNicName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceTemplateNicPortSpeed: {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: applyOnce,
							Deprecated:       "This field is deprected",
						},
						isInstanceTemplateNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceTemplateNicSubnet: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			isInstanceTemplateNetworkInterfaces: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						/*
							isInstanceTemplateNicAllowIpSpoofing: {
								Type:     schema.TypeBool,
								Optional: true,
								Default:  false,
							},
						*/
						isInstanceTemplateNicName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceTemplateNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceTemplateNicSubnet: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			isInstanceTemplateGeneration: {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: applyOnce,
				ValidateFunc:     validateGeneration,
				Removed:          "This field is removed",
			},

			isInstanceTemplateUserData: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},

			isInstanceTemplateImage: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isInstanceTemplateBootVolume: {
				Type:             schema.TypeList,
				DiffSuppressFunc: applyOnce,
				Optional:         true,
				Computed:         true,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateBootName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceTemplateBootEncryption: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isInstanceTemplateBootSize: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceTemplateBootIOPS: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceTemplateBootProfile: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isInstanceTemplateResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func instanceTemplateCreate(d *schema.ResourceData, meta interface{}, profile, name, vpcID, zone, image string) error {
	sess, err := vpcClient(meta)
	//boolAllowIpSpoofing := false
	if err != nil {
		return err
	}
	instanceproto := &vpcv1.InstanceTemplatePrototype{
		Image: &vpcv1.ImageIdentity{
			ID: &image,
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
		Profile: &vpcv1.InstanceProfileIdentity{
			Name: &profile,
		},
		Name: &name,
		Vpc: &vpcv1.VPCIdentity{
			ID: &vpcID,
		},
	}

	if boot, ok := d.GetOk(isInstanceTemplateBootVolume); ok {
		bootvol := boot.([]interface{})[0].(map[string]interface{})
		var volTemplate = &vpcv1.VolumePrototypeInstanceByImageContext{}
		name, ok := bootvol[isInstanceTemplateBootName]
		namestr := name.(string)
		if ok {
			volTemplate.Name = &namestr
		}
		// enc, ok := bootvol[isInstanceBootEncryption]
		// encstr := enc.(string)
		// if ok && encstr != "" {
		// 	volTemplate.EncryptionKey = &vpcv1.EncryptionKeyIdentity{
		// 		Crn: &encstr,
		// 	}
		// }
		volcap := 100
		volcapint64 := int64(volcap)
		volprof := "general-purpose"
		volTemplate.Capacity = &volcapint64
		volTemplate.Profile = &vpcv1.VolumeProfileIdentity{
			Name: &volprof,
		}
		deletebool := true
		instanceproto.BootVolumeAttachment = &vpcv1.VolumeAttachmentPrototypeInstanceByImageContext{
			DeleteVolumeOnInstanceDelete: &deletebool,
			Volume:                       volTemplate,
		}
	}

	if volsintf, ok := d.GetOk("volume_attachments"); ok {
		vols := volsintf.([]interface{})
		var intfs []vpcv1.VolumeAttachmentPrototypeInstanceContext
		for _, resource := range vols {
			vol := resource.(map[string]interface{})
			volInterface := &vpcv1.VolumeAttachmentPrototypeInstanceContext{}
			deleteVol, _ := vol["delete_volume_on_instance_delete"]
			deleteVolBool := deleteVol.(bool)
			volInterface.DeleteVolumeOnInstanceDelete = &deleteVolBool
			name, _ := vol["name"]
			namestr := name.(string)
			volInterface.Name = &namestr
			volintf, _ := vol["volume"]
			volintfstr := volintf.(string)
			volInterface.Volume = &vpcv1.VolumeAttachmentPrototypeInstanceContextVolume{
				ID: &volintfstr,
			}
			intfs = append(intfs, *volInterface)
		}
		instanceproto.VolumeAttachments = intfs
	}

	if primnicintf, ok := d.GetOk(isInstanceTemplatePrimaryNetworkInterface); ok {
		primnic := primnicintf.([]interface{})[0].(map[string]interface{})
		subnetintf, _ := primnic[isInstanceTemplateNicSubnet]
		subnetintfstr := subnetintf.(string)
		var primnicobj = &vpcv1.NetworkInterfacePrototype{}
		primnicobj.Subnet = &vpcv1.SubnetIdentity{
			ID: &subnetintfstr,
		}
		name, _ := primnic[isInstanceTemplateNicName]
		namestr := name.(string)
		if namestr != "" {
			primnicobj.Name = &namestr
		}
		/*
			allowIpSpoofing, ok := primnic[isInstanceTemplateNicAllowIpSpoofing]
			allowIpSpoofingbool := allowIpSpoofing.(bool)
			if ok {
				// primnicobj.AllowIpSpoofing = &allowIpSpoofingbool
				primnicobj.AllowIpSpoofing = &boolAllowIpSpoofing
			}
			log.Printf("[INFO] ****** Create: Allow Primary IP Spoofing ****** : %t", allowIpSpoofingbool)
		*/
		secgrpintf, ok := primnic[isInstanceTemplateNicSecurityGroups]
		if ok {
			secgrpSet := secgrpintf.(*schema.Set)
			if secgrpSet.Len() != 0 {
				var secgrpobjs = make([]vpcv1.SecurityGroupIdentityIntf, secgrpSet.Len())
				for i, secgrpIntf := range secgrpSet.List() {
					secgrpIntfstr := secgrpIntf.(string)
					secgrpobjs[i] = &vpcv1.SecurityGroupIdentity{
						ID: &secgrpIntfstr,
					}
				}
				primnicobj.SecurityGroups = secgrpobjs
			}
		}
		instanceproto.PrimaryNetworkInterface = primnicobj
	}

	if nicsintf, ok := d.GetOk(isInstanceTemplateNetworkInterfaces); ok {
		nics := nicsintf.([]interface{})
		var intfs []vpcv1.NetworkInterfacePrototype
		for _, resource := range nics {
			nic := resource.(map[string]interface{})
			nwInterface := &vpcv1.NetworkInterfacePrototype{}
			subnetintf, _ := nic[isInstanceTemplateNicSubnet]
			subnetintfstr := subnetintf.(string)
			nwInterface.Subnet = &vpcv1.SubnetIdentity{
				ID: &subnetintfstr,
			}
			name, ok := nic[isInstanceTemplateNicName]
			namestr := name.(string)
			if ok && namestr != "" {
				nwInterface.Name = &namestr
			}
			/*
				allowIpSpoofing, ok := nic[isInstanceTemplateNicAllowIpSpoofing]
				allowIpSpoofingbool := allowIpSpoofing.(bool)
				if ok {
					// nwInterface.AllowIpSpoofing = &allowIpSpoofingbool
					nwInterface.AllowIpSpoofing = &boolAllowIpSpoofing
				}
				log.Printf("[INFO] ****** Create: Allow Network Interfaces IP Spoofing ****** : %t", allowIpSpoofingbool)
			*/
			secgrpintf, ok := nic[isInstanceTemplateNicSecurityGroups]
			if ok {
				secgrpSet := secgrpintf.(*schema.Set)
				if secgrpSet.Len() != 0 {
					var secgrpobjs = make([]vpcv1.SecurityGroupIdentityIntf, secgrpSet.Len())
					for i, secgrpIntf := range secgrpSet.List() {
						secgrpIntfstr := secgrpIntf.(string)
						secgrpobjs[i] = &vpcv1.SecurityGroupIdentity{
							ID: &secgrpIntfstr,
						}
					}
					nwInterface.SecurityGroups = secgrpobjs
				}
			}
			intfs = append(intfs, *nwInterface)
		}
		instanceproto.NetworkInterfaces = intfs
	}

	keySet := d.Get(isInstanceTemplateKeys).(*schema.Set)
	if keySet.Len() != 0 {
		keyobjs := make([]vpcv1.KeyIdentityIntf, keySet.Len())
		for i, key := range keySet.List() {
			keystr := key.(string)
			keyobjs[i] = &vpcv1.KeyIdentity{
				ID: &keystr,
			}
		}
		instanceproto.Keys = keyobjs
	}

	if userdata, ok := d.GetOk(isInstanceTemplateUserData); ok {
		userdatastr := userdata.(string)
		instanceproto.UserData = &userdatastr
	}

	if grp, ok := d.GetOk(isInstanceTemplateResourceGroup); ok {
		grpstr := grp.(string)
		instanceproto.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &grpstr,
		}

	}

	options := &vpcv1.CreateInstanceTemplateOptions{
		InstanceTemplatePrototype: instanceproto,
	}

	instanceIntf, response, err := sess.CreateInstanceTemplate(options)
	if err != nil {
		log.Printf("[DEBUG] Instance err %s\n%s", err, response)
		return err
	}
	instance := instanceIntf.(*vpcv1.InstanceTemplate)
	d.SetId(*instance.ID)

	log.Printf("[INFO] Instance Template : %s", *instance.ID)
	/*
		_, err = isWaitForInstanceTemplateAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate), d)
		if err != nil {
			return err
		}
	*/
	return nil
}

func resourceIBMisInstanceTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	profile := d.Get(isInstanceTemplateProfile).(string)
	name := d.Get(isInstanceTemplateName).(string)
	vpcID := d.Get(isInstanceTemplateVPC).(string)
	zone := d.Get(isInstanceTemplateZone).(string)
	image := d.Get(isInstanceTemplateImage).(string)

	err := instanceTemplateCreate(d, meta, profile, name, vpcID, zone, image)
	if err != nil {
		return err
	}

	return resourceIBMisInstanceTemplateUpdate(d, meta)
}

func resourceIBMisInstanceTemplateRead(d *schema.ResourceData, meta interface{}) error {
	ID := d.Id()
	err := instanceTemplateGet(d, meta, ID)
	if err != nil {
		return err
	}
	return nil
}

func instanceTemplateGet(d *schema.ResourceData, meta interface{}, ID string) error {
	instanceC, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getinsOptions := &vpcv1.GetInstanceTemplateOptions{
		ID: &ID,
	}
	instanceIntf, response, err := instanceC.GetInstanceTemplate(getinsOptions)
	instance := instanceIntf.(*vpcv1.InstanceTemplate)
	if err != nil {
		return fmt.Errorf("Error Getting Instance: %s\n%s", err, response)
	}
	d.Set(isInstanceTemplateName, instance.Name)
	if instance.Profile != nil {
		instanceProfileIntf := instance.Profile
		identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
		d.Set(isInstanceTemplateProfile, identity.Name)
	}

	if instance.PrimaryNetworkInterface != nil {
		primaryNicList := make([]map[string]interface{}, 0)
		currentPrimNic := map[string]interface{}{}
		currentPrimNic[isInstanceNicName] = instance.PrimaryNetworkInterface.Name
		currentPrimNic[isInstanceNicPrimaryIpv4Address] = instance.PrimaryNetworkInterface.PrimaryIpv4Address
		subInf := instance.PrimaryNetworkInterface.Subnet
		subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
		currentPrimNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
		//currentPrimNic[isInstanceTemplateNicAllowIpSpoofing] = instance.PrimaryNetworkInterface.AllowIpSpoofing
		if len(instance.PrimaryNetworkInterface.SecurityGroups) != 0 {
			secgrpList := []string{}
			for i := 0; i < len(instance.PrimaryNetworkInterface.SecurityGroups); i++ {
				secGrpInf := instance.PrimaryNetworkInterface.SecurityGroups[i]
				subnetIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
				secgrpList = append(secgrpList, string(*subnetIdentity.ID))
			}
			currentPrimNic[isInstanceTemplateNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
		}
		primaryNicList = append(primaryNicList, currentPrimNic)
		d.Set(isInstanceTemplatePrimaryNetworkInterface, primaryNicList)
	}

	if instance.NetworkInterfaces != nil {
		interfacesList := make([]map[string]interface{}, 0)
		for _, intfc := range instance.NetworkInterfaces {
			currentNic := map[string]interface{}{}
			currentNic[isInstanceTemplateNicName] = intfc.Name
			currentNic[isInstanceTemplateNicPrimaryIpv4Address] = intfc.PrimaryIpv4Address
			//currentNic[isInstanceTemplateNicAllowIpSpoofing] = intfc.AllowIpSpoofing
			subInf := intfc.Subnet
			subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
			currentNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
			if len(intfc.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(intfc.SecurityGroups); i++ {
					secGrpInf := intfc.SecurityGroups[i]
					subnetIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
					secgrpList = append(secgrpList, string(*subnetIdentity.ID))
				}
				currentNic[isInstanceTemplateNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
			}
			interfacesList = append(interfacesList, currentNic)
		}
		d.Set(isInstanceTemplateNetworkInterfaces, interfacesList)
	}

	if instance.Image != nil {
		imageInf := instance.Image
		imageIdentity := imageInf.(*vpcv1.ImageIdentity)
		d.Set(isInstanceTemplateImage, imageIdentity.ID)
	}
	vpcInf := instance.Vpc
	vpcRef := vpcInf.(*vpcv1.VPCIdentity)
	d.Set(isInstanceTemplateVPC, vpcRef.ID)
	zoneInf := instance.Zone
	zone := zoneInf.(*vpcv1.ZoneIdentity)
	d.Set(isInstanceTemplateZone, zone.Name)

	interfacesList := make([]map[string]interface{}, 0)
	if instance.VolumeAttachments != nil {
		for _, volume := range instance.VolumeAttachments {
			volumeAttach := map[string]interface{}{}
			volumeAttach[isInstanceTemplateVolAttName] = volume.Name
			volumeAttach[isInstanceTemplateDeleteVolume] = volume.DeleteVolumeOnInstanceDelete
			volumeId := map[string]interface{}{}
			volumeIntf := volume.Volume
			volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeInstanceContextVolume)
			volumeId["name"] = volumeInst.Name
			volumeId["iops"] = volumeInst.Iops
			volumeAttach[isInstanceTemplateVolAttVolume] = volumeId
			interfacesList = append(interfacesList, volumeAttach)
		}
		d.Set(isInstanceTemplateVolumeAttachments, interfacesList)
	}
	if instance.BootVolumeAttachment != nil {
		bootVolList := make([]map[string]interface{}, 0)
		bootVol := map[string]interface{}{}
		bootVol[isInstanceTemplateBootName] = instance.BootVolumeAttachment.Name
		bootVol[isInstanceTemplateVolAttVolume] = instance.BootVolumeAttachment.Volume
		bootVol[isInstanceTemplateDeleteVolume] = instance.BootVolumeAttachment.DeleteVolumeOnInstanceDelete

		// getvolattoptions := &vpcclassicv1.GetVolumeAttachmentOptions{
		// 	InstanceID: &ID,
		// 	ID:         instance.BootVolumeAttachment.Volume.ID,
		// }
		// vol, _, err := instanceC.GetVolumeAttachment(getvolattoptions)
		// if err != nil {
		// 	return fmt.Errorf("Error while retrieving boot volume %s for instance %s: %v", getvolattoptions.ID, d.Id(), err)
		// }

		// bootVol[isInstanceBootSize] = instance.BootVolumeAttachment.Capacity
		// bootVol[isInstanceBootIOPS] = instance.BootVolumeAttachment.Iops
		// bootVol[isInstanceBootProfile] = instance.BootVolumeAttachment.Name
		bootVolList = append(bootVolList, bootVol)

		d.Set(isInstanceBootVolume, bootVolList)
	}

	if instance.ResourceGroup != nil {
		d.Set(isInstanceTemplateResourceGroup, instance.ResourceGroup.ID)
	}
	return nil
}

func instanceTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	instanceC, err := vpcClient(meta)
	if err != nil {
		return err
	}
	ID := d.Id()

	if d.HasChange(isInstanceName) {
		name := d.Get(isInstanceTemplateName).(string)
		updnetoptions := &vpcv1.UpdateInstanceTemplateOptions{
			ID:   &ID,
			Name: &name,
		}

		_, _, err = instanceC.UpdateInstanceTemplate(updnetoptions)
		if err != nil {
			return err
		}
	}

	getinsOptions := &vpcv1.GetInstanceTemplateOptions{
		ID: &ID,
	}
	instanceIntf, response, err := instanceC.GetInstanceTemplate(getinsOptions)
	instance := instanceIntf.(*vpcv1.InstanceTemplate)

	log.Printf("[INFO] Instance Template : %s", *instance.ID)
	if err != nil {
		return fmt.Errorf("Error Getting Instance Template: %s\n%s", err, response)
	}

	return nil
}

func resourceIBMisInstanceTemplateUpdate(d *schema.ResourceData, meta interface{}) error {

	err := instanceTemplateUpdate(d, meta)
	if err != nil {
		return err
	}

	return resourceIBMisInstanceTemplateRead(d, meta)
}

func instanceTemplateDelete(d *schema.ResourceData, meta interface{}, ID string) error {
	instanceC, err := vpcClient(meta)
	if err != nil {
		return err
	}

	deleteinstanceTemplateOptions := &vpcv1.DeleteInstanceTemplateOptions{
		ID: &ID,
	}
	_, err = instanceC.DeleteInstanceTemplate(deleteinstanceTemplateOptions)
	if err != nil {
		return err
	}
	return nil
}

func resourceIBMisInstanceTemplateDelete(d *schema.ResourceData, meta interface{}) error {

	ID := d.Id()

	err := instanceTemplateDelete(d, meta, ID)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func instanceTemplateExists(d *schema.ResourceData, meta interface{}, ID string) error {
	instanceC, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getinsOptions := &vpcv1.GetInstanceTemplateOptions{
		ID: &ID,
	}
	_, response, err := instanceC.GetInstanceTemplate(getinsOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting Instance: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}

func resourceIBMisInstanceTemplateExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	ID := d.Id()
	err := instanceTemplateExists(d, meta, ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
