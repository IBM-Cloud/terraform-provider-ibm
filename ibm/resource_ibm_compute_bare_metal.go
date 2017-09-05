package ibm

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func resourceIBMComputeBareMetal() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMComputeBareMetalCreate,
		Read:     resourceIBMComputeBareMetalRead,
		Update:   resourceIBMComputeBareMetalUpdate,
		Delete:   resourceIBMComputeBareMetalDelete,
		Exists:   resourceIBMComputeBareMetalExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				DefaultFunc: genID,
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					// FIXME: Work around another bug in terraform.
					// When a default function is used with an optional property,
					// terraform will always execute it on apply, even when the property
					// already has a value in the state for it. This causes a false diff.
					// Making the property Computed:true does not make a difference.
					if strings.HasPrefix(o, "terraformed-") && strings.HasPrefix(n, "terraformed-") {
						return true
					}

					return o == n
				},
			},

			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"os_reference_code": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"image_template_id"},
			},

			"hourly_billing": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},

			"private_network_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"datacenter": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"public_vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"public_subnet": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"private_vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"private_subnet": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"network_speed": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
				ForceNew: true,
			},

			"public_ipv4_address": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"private_ipv4_address": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"ssh_key_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				ForceNew: true,
			},

			"user_metadata": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"file_storage_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Set: func(v interface{}) int {
					return v.(int)
				},
			},

			"block_storage_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Set: func(v interface{}) int {
					return v.(int)
				},
			},

			"post_install_script_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
				ForceNew: true,
			},

			"fixed_config_preset": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"image_template_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"os_reference_code"},
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func getBareMetalOrderFromResourceData(d *schema.ResourceData, meta interface{}) (datatypes.Hardware, error) {
	dc := datatypes.Location{
		Name: sl.String(d.Get("datacenter").(string)),
	}

	networkComponent := datatypes.Network_Component{
		MaxSpeed: sl.Int(d.Get("network_speed").(int)),
	}

	hardware := datatypes.Hardware{
		Hostname:               sl.String(d.Get("hostname").(string)),
		Domain:                 sl.String(d.Get("domain").(string)),
		HourlyBillingFlag:      sl.Bool(d.Get("hourly_billing").(bool)),
		PrivateNetworkOnlyFlag: sl.Bool(d.Get("private_network_only").(bool)),
		Datacenter:             &dc,
		NetworkComponents:      []datatypes.Network_Component{networkComponent},
		PostInstallScriptUri:   sl.String(d.Get("post_install_script_uri").(string)),
		BareMetalInstanceFlag:  sl.Int(1),

		FixedConfigurationPreset: &datatypes.Product_Package_Preset{
			KeyName: sl.String(d.Get("fixed_config_preset").(string)),
		},
	}

	if operatingSystemReferenceCode, ok := d.GetOk("os_reference_code"); ok {
		hardware.OperatingSystemReferenceCode = sl.String(operatingSystemReferenceCode.(string))
	}

	public_vlan_id := d.Get("public_vlan_id").(int)
	if public_vlan_id > 0 {
		hardware.PrimaryNetworkComponent = &datatypes.Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{Id: sl.Int(public_vlan_id)},
		}
	}

	private_vlan_id := d.Get("private_vlan_id").(int)
	if private_vlan_id > 0 {
		hardware.PrimaryBackendNetworkComponent = &datatypes.Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{Id: sl.Int(private_vlan_id)},
		}
	}

	if public_subnet, ok := d.GetOk("public_subnet"); ok {
		subnet := public_subnet.(string)
		subnetID, err := getSubnetID(subnet, meta)
		if err != nil {
			return hardware, fmt.Errorf("Error determining id for subnet %s: %s", subnet, err)
		}

		hardware.PrimaryNetworkComponent.NetworkVlan.PrimarySubnetId = sl.Int(subnetID)
	}

	if private_subnet, ok := d.GetOk("private_subnet"); ok {
		subnet := private_subnet.(string)
		subnetID, err := getSubnetID(subnet, meta)
		if err != nil {
			return hardware, fmt.Errorf("Error determining id for subnet %s: %s", subnet, err)
		}

		hardware.PrimaryBackendNetworkComponent.NetworkVlan.PrimarySubnetId = sl.Int(subnetID)
	}

	if userMetadata, ok := d.GetOk("user_metadata"); ok {
		hardware.UserData = []datatypes.Hardware_Attribute{
			{Value: sl.String(userMetadata.(string))},
		}
	}

	// Get configured ssh_keys
	ssh_key_ids := d.Get("ssh_key_ids").([]interface{})
	if len(ssh_key_ids) > 0 {
		hardware.SshKeys = make([]datatypes.Security_Ssh_Key, 0, len(ssh_key_ids))
		for _, ssh_key_id := range ssh_key_ids {
			hardware.SshKeys = append(hardware.SshKeys, datatypes.Security_Ssh_Key{
				Id: sl.Int(ssh_key_id.(int)),
			})
		}
	}

	return hardware, nil
}

func resourceIBMComputeBareMetalCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	hwService := services.GetHardwareService(sess)
	orderService := services.GetProductOrderService(sess)

	hardware, err := getBareMetalOrderFromResourceData(d, meta)
	if err != nil {
		return err
	}

	order, err := hwService.GenerateOrderTemplate(&hardware)
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to get the bare metal order template: %s", err)
	}

	// Set image template id if it exists
	if rawImageTemplateId, ok := d.GetOk("image_template_id"); ok {
		imageTemplateId := rawImageTemplateId.(int)
		order.ImageTemplateId = sl.Int(imageTemplateId)
	}

	log.Println("[INFO] Ordering bare metal server")

	_, err = orderService.PlaceOrder(&order, sl.Bool(false))
	if err != nil {
		return fmt.Errorf("Error ordering bare metal server: %s", err)
	}

	log.Printf("[INFO] Bare Metal Server ID: %s", d.Id())

	// wait for machine availability
	bm, err := waitForBareMetalProvision(&hardware, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for bare metal server (%s) to become ready: %s", d.Id(), err)
	}

	id := *bm.(datatypes.Hardware).Id
	d.SetId(fmt.Sprintf("%d", id))

	// Set tags
	err = setHardwareTags(id, d, meta)
	if err != nil {
		return err
	}

	var storageIds []int
	if storageIdsSet := d.Get("file_storage_ids").(*schema.Set); len(storageIdsSet.List()) > 0 {
		storageIds = expandIntList(storageIdsSet.List())

	}
	if storageIdsSet := d.Get("block_storage_ids").(*schema.Set); len(storageIdsSet.List()) > 0 {
		storageIds = append(storageIds, expandIntList(storageIdsSet.List())...)
	}
	if len(storageIds) > 0 {
		err := addAccessToStorageList(hwService.Id(id), id, storageIds, meta)
		if err != nil {
			return err
		}
	}

	// Set notes
	if d.Get("notes").(string) != "" {
		err = setHardwareNotes(id, d, meta)
		if err != nil {
			return err
		}
	}

	return resourceIBMComputeBareMetalRead(d, meta)
}

func resourceIBMComputeBareMetalRead(d *schema.ResourceData, meta interface{}) error {
	service := services.GetHardwareService(meta.(ClientSession).SoftLayerSession())

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	result, err := service.Id(id).Mask(
		"hostname,domain," +
			"primaryIpAddress,primaryBackendIpAddress,privateNetworkOnlyFlag," +
			"notes,userData[value],tagReferences[id,tag[name]]," +
			"allowedNetworkStorage[id,nasType]," +
			"hourlyBillingFlag," +
			"datacenter[id,name,longName]," +
			"primaryNetworkComponent[networkVlan[id,primaryRouter,vlanNumber],maxSpeed]," +
			"primaryBackendNetworkComponent[networkVlan[id,primaryRouter,vlanNumber],maxSpeed]",
	).GetObject()

	if err != nil {
		return fmt.Errorf("Error retrieving bare metal server: %s", err)
	}

	d.Set("hostname", *result.Hostname)
	d.Set("domain", *result.Domain)

	if result.Datacenter != nil {
		d.Set("datacenter", *result.Datacenter.Name)
	}

	d.Set("network_speed", *result.PrimaryNetworkComponent.MaxSpeed)
	if result.PrimaryIpAddress != nil {
		d.Set("public_ipv4_address", *result.PrimaryIpAddress)
	}
	d.Set("private_ipv4_address", *result.PrimaryBackendIpAddress)

	d.Set("private_network_only", *result.PrivateNetworkOnlyFlag)
	d.Set("hourly_billing", *result.HourlyBillingFlag)

	if result.PrimaryNetworkComponent.NetworkVlan != nil {
		d.Set("public_vlan_id", *result.PrimaryNetworkComponent.NetworkVlan.Id)
	}

	if result.PrimaryBackendNetworkComponent.NetworkVlan != nil {
		d.Set("private_vlan_id", *result.PrimaryBackendNetworkComponent.NetworkVlan.Id)
	}

	userData := result.UserData
	if len(userData) > 0 && userData[0].Value != nil {
		d.Set("user_metadata", *userData[0].Value)
	}

	d.Set("notes", sl.Get(result.Notes, nil))

	tagReferences := result.TagReferences
	tagReferencesLen := len(tagReferences)
	if tagReferencesLen > 0 {
		tags := make([]string, 0, tagReferencesLen)
		for _, tagRef := range tagReferences {
			tags = append(tags, *tagRef.Tag.Name)
		}
		d.Set("tags", tags)
	}

	storages := result.AllowedNetworkStorage
	if len(storages) > 0 {
		d.Set("block_storage_ids", flattenBlockStorageID(storages))
		d.Set("file_storage_ids", flattenFileStorageID(storages))
	}

	connInfo := map[string]string{"type": "ssh"}
	if !*result.PrivateNetworkOnlyFlag && result.PrimaryIpAddress != nil {
		connInfo["host"] = *result.PrimaryIpAddress
	} else {
		connInfo["host"] = *result.PrimaryBackendIpAddress
	}
	d.SetConnInfo(connInfo)

	return nil
}

func resourceIBMComputeBareMetalUpdate(d *schema.ResourceData, meta interface{}) error {
	id, _ := strconv.Atoi(d.Id())
	service := services.GetHardwareService(meta.(ClientSession).SoftLayerSession())

	if d.HasChange("tags") {
		err := setHardwareTags(id, d, meta)
		if err != nil {
			return err
		}
	}

	if d.HasChange("notes") {
		err := setHardwareNotes(id, d, meta)
		if err != nil {
			return err
		}
	}
	err := modifyStorageAccess(service.Id(id), id, meta, d)
	if err != nil {
		return err
	}

	return nil
}

func resourceIBMComputeBareMetalDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetHardwareService(sess)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	_, err = waitForNoBareMetalActiveTransactions(id, meta)
	if err != nil {
		return fmt.Errorf("Error deleting bare metal server while waiting for zero active transactions: %s", err)
	}

	billingItem, err := service.Id(id).GetBillingItem()
	if err != nil {
		return fmt.Errorf("Error getting billing item for bare metal server: %s", err)
	}

	billingItemService := services.GetBillingItemService(sess)
	_, err = billingItemService.Id(*billingItem.Id).CancelItem(
		sl.Bool(true), sl.Bool(true), sl.String("No longer required"), sl.String("Please cancel this server"),
	)
	if err != nil {
		return fmt.Errorf("Error canceling the bare metal server (%d): %s", id, err)
	}

	return nil
}

func resourceIBMComputeBareMetalExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	service := services.GetHardwareService(meta.(ClientSession).SoftLayerSession())

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	result, err := service.Id(id).GetObject()
	if err != nil {
		if apiErr, ok := err.(sl.Error); !ok || apiErr.StatusCode != 404 {
			return false, fmt.Errorf("Error trying to retrieve the Bare Metal server: %s", err)
		}
	}

	return result.Id != nil && *result.Id == id, nil
}

// Bare metal creation does not return a bare metal object with an Id.
// Have to wait on provision date to become available on server that matches
// hostname and domain.
// http://sldn.softlayer.com/blog/bpotter/ordering-bare-metal-servers-using-softlayer-api
func waitForBareMetalProvision(d *datatypes.Hardware, meta interface{}) (interface{}, error) {
	hostname := *d.Hostname
	domain := *d.Domain
	log.Printf("Waiting for server (%s.%s) to have to be provisioned", hostname, domain)

	stateConf := &resource.StateChangeConf{
		Pending: []string{"retry", "pending"},
		Target:  []string{"provisioned"},
		Refresh: func() (interface{}, string, error) {
			service := services.GetAccountService(meta.(ClientSession).SoftLayerSession())
			bms, err := service.Filter(
				filter.Build(
					filter.Path("hardware.hostname").Eq(hostname),
					filter.Path("hardware.domain").Eq(domain),
				),
			).Mask("id,provisionDate").GetHardware()
			if err != nil {
				return false, "retry", nil
			}

			if len(bms) == 0 || bms[0].ProvisionDate == nil {
				return datatypes.Hardware{}, "pending", nil
			} else {
				return bms[0], "provisioned", nil
			}
		},
		Timeout:    4 * time.Hour,
		Delay:      30 * time.Second,
		MinTimeout: 2 * time.Minute,
	}

	return stateConf.WaitForState()
}

func waitForNoBareMetalActiveTransactions(id int, meta interface{}) (interface{}, error) {
	log.Printf("Waiting for server (%d) to have zero active transactions", id)
	service := services.GetHardwareServerService(meta.(ClientSession).SoftLayerSession())

	stateConf := &resource.StateChangeConf{
		Pending: []string{"retry", "active"},
		Target:  []string{"idle"},
		Refresh: func() (interface{}, string, error) {
			bm, err := service.Id(id).Mask("id,activeTransactionCount").GetObject()
			if err != nil {
				return false, "retry", nil
			}

			if bm.ActiveTransactionCount != nil && *bm.ActiveTransactionCount == 0 {
				return bm, "idle", nil
			} else {
				return bm, "active", nil
			}
		},
		Timeout:    4 * time.Hour,
		Delay:      5 * time.Second,
		MinTimeout: 1 * time.Minute,
	}

	return stateConf.WaitForState()
}

func setHardwareTags(id int, d *schema.ResourceData, meta interface{}) error {
	service := services.GetHardwareService(meta.(ClientSession).SoftLayerSession())

	tags := getTags(d)
	if tags != "" {
		_, err := service.Id(id).SetTags(sl.String(tags))
		if err != nil {
			return fmt.Errorf("Could not set tags on bare metal server %d", id)
		}
	}

	return nil
}

func setHardwareNotes(id int, d *schema.ResourceData, meta interface{}) error {
	service := services.GetHardwareServerService(meta.(ClientSession).SoftLayerSession())

	result, err := service.Id(id).GetObject()
	if err != nil {
		return err
	}

	result.Notes = sl.String(d.Get("notes").(string))

	_, err = service.Id(id).EditObject(&result)
	if err != nil {
		return err
	}

	return nil
}
