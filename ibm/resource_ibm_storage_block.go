package ibm

import (
	"fmt"
	"log"
	"strconv"

	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/helpers/network"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func resourceIBMStorageBlock() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMStorageBlockCreate,
		Read:     resourceIBMStorageBlockRead,
		Update:   resourceIBMStorageBlockUpdate,
		Delete:   resourceIBMStorageBlockDelete,
		Exists:   resourceIBMStorageBlockExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"datacenter": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"capacity": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"iops": {
				Type:     schema.TypeFloat,
				Optional: true,
				ForceNew: true,
			},

			"volumename": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"disk_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"snapshot_capacity": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			"os_format_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"allowed_virtual_guest_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Set: func(v interface{}) int {
					return v.(int)
				},
			},

			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//TODO in v0.9.0
			"allowed_virtual_guest_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_iqn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Set: func(v interface{}) int {
					virtualGuest := v.(map[string]interface{})
					return virtualGuest["id"].(int)
				},
				Deprecated: "Please use 'allowed_host_info' instead",
			},

			"allowed_hardware_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Set: func(v interface{}) int {
					return v.(int)
				},
			},

			//TODO in v0.9.0
			"allowed_hardware_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_iqn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Set: func(v interface{}) int {
					baremetal := v.(map[string]interface{})
					return baremetal["id"].(int)
				},
				Deprecated: "Please use 'allowed_host_info' instead",
			},

			"allowed_ip_addresses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"hourly_billing": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"allowed_host_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_iqn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMStorageBlockCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()

	storageType := d.Get("type").(string)
	var iops float64
	var err error
	var diskdescription string
	var osFormatType string
	if _, ok := d.GetOk("disk_description"); ok {
		diskdescription = d.Get("disk_description").(string)
	}
	if _, ok := d.GetOk("iops"); ok {
		iops = d.Get("iops").(float64)
	}
	datacenter := d.Get("datacenter").(string)
	capacity := d.Get("capacity").(int)
	snapshotCapacity := d.Get("snapshot_capacity").(int)
	var osType datatypes.Network_Storage_Iscsi_OS_Type
	if _, ok := d.GetOk("os_format_type"); ok {
		osFormatType = d.Get("os_format_type").(string)
		osType, err = network.GetOsTypeByName(sess, osFormatType)
	}
	hourlyBilling := d.Get("hourly_billing").(bool)
	if err != nil {
		return err
	}
	var storageOrderContainer datatypes.Container_Product_Order
	if storageType == "Portable" {
		storageOrderContainer, err = buildStorageProductOrderContainer(sess, storageType, 0.00, capacity, snapshotCapacity, blockStorage, datacenter, hourlyBilling)
	} else {
		storageOrderContainer, err = buildStorageProductOrderContainer(sess, storageType, iops, capacity, snapshotCapacity, blockStorage, datacenter, hourlyBilling)
	}
	if err != nil {
		return fmt.Errorf("Error while creating storage:%s", err)
	}

	log.Println("[INFO] Creating storage")

	var receipt datatypes.Container_Product_Order_Receipt

	switch storageType {
	case enduranceType:
		receipt, err = services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(
			&datatypes.Container_Product_Order_Network_Storage_AsAService{
				Container_Product_Order: storageOrderContainer,
				OsFormatType: &datatypes.Network_Storage_Iscsi_OS_Type{
					Id:      osType.Id,
					KeyName: osType.KeyName,
				},
				VolumeSize: &capacity,
			}, sl.Bool(false))
	case performanceType:
		receipt, err = services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(
			&datatypes.Container_Product_Order_Network_Storage_AsAService{
				Container_Product_Order: storageOrderContainer,
				OsFormatType: &datatypes.Network_Storage_Iscsi_OS_Type{
					Id:      osType.Id,
					KeyName: osType.KeyName,
				},
				Iops:       sl.Int(int(iops)),
				VolumeSize: &capacity,
			}, sl.Bool(false))
	case portableType:
		receipt, err = services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(
			&datatypes.Container_Product_Order_Virtual_Disk_Image{
				Container_Product_Order: storageOrderContainer,
				DiskDescription:         &diskdescription,
			}, sl.Bool(false))
	default:
		return fmt.Errorf("Error during creation of storage: Invalid storageType %s", storageType)
	}

	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}

	// Find the storage device
	var blockStorage datatypes.Network_Storage
	var portablestorage datatypes.Virtual_Disk_Image
	if storageType != portableType {
		blockStorage, _, err = findStorageByOrderId(sess, *receipt.OrderId, "")
	} else {
		_, portablestorage, err = findStorageByOrderId(sess, *receipt.OrderId, portableType)
	}
	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}
	if storageType != portableType {
		d.SetId(fmt.Sprintf("%d", *blockStorage.Id))
	} else {
		d.SetId(fmt.Sprintf("%d", *portablestorage.Id))
	}

	// Wait for storage availability
	if storageType != portableType {
		_, err = WaitForStorageAvailable(d, meta, "")
	}

	if err != nil {
		return fmt.Errorf(
			"Error waiting for storage (%s) to become ready: %s", d.Id(), err)
	}

	// SoftLayer changes the device ID after completion of provisioning. It is necessary to refresh device ID.
	if storageType != portableType {
		blockStorage, _, err = findStorageByOrderId(sess, *receipt.OrderId, "")
	} else {
		_, portablestorage, err = findStorageByOrderId(sess, *receipt.OrderId, portableType)
	}

	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}
	if storageType != portableType {
		d.SetId(fmt.Sprintf("%d", *blockStorage.Id))
	} else {
		d.SetId(fmt.Sprintf("%d", *portablestorage.Id))
	}

	log.Printf("[INFO] Storage ID: %s", d.Id())

	if storageType != portableType {
		return resourceIBMStorageBlockUpdate(d, meta)
	}
	return resourceIBMStorageBlockRead(d, meta)
}

func resourceIBMStorageBlockRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	storageId, _ := strconv.Atoi(d.Id())
	var err error
	portablestoragefilter := "portableStorageVolumes.id"
	var portablestorage []datatypes.Virtual_Disk_Image
	if d.Get("type") == portableType {
		portablestorage, err = services.GetAccountService(sess).
			Filter(filter.Build(
				filter.Path(portablestoragefilter).
					Eq(strconv.Itoa(storageId)))).
			GetPortableStorageVolumes()
		if err != nil {
			return fmt.Errorf("Error during fetching details of storage: %s", err)
		}
		d.Set("type", "Portable")
		d.Set("capacity", *portablestorage[0].Capacity)
		d.Set("volumename", *portablestorage[0].Description)
	} else {
		storage, err := services.GetNetworkStorageService(sess).
			Id(storageId).
			Mask(storageDetailMask).
			GetObject()

		storageType := strings.Fields(*storage.StorageType.Description)[0]

		// Calculate IOPS
		iops, err := getIops(storage, storageType)
		if err != nil {
			return fmt.Errorf("Error retrieving storage information: %s", err)
		}

		d.Set("type", storageType)
		d.Set("capacity", *storage.CapacityGb)
		d.Set("volumename", *storage.Username)
		d.Set("hostname", *storage.ServiceResourceBackendIpAddress)
		d.Set("iops", iops)
		if storage.SnapshotCapacityGb != nil {
			snapshotCapacity, _ := strconv.Atoi(*storage.SnapshotCapacityGb)
			d.Set("snapshot_capacity", snapshotCapacity)
		}

		// Parse data center short name from ServiceResourceName. For example,
		// if SoftLayer API returns "'serviceResourceName': 'PerfStor Aggr aggr_staasdal0601_p01'",
		// the data center short name is "dal06".
		r, _ := regexp.Compile("[a-zA-Z]{3}[0-9]{2}")
		d.Set("datacenter", r.FindString(*storage.ServiceResourceName))

		allowedHostInfoList := make([]map[string]interface{}, 0)

		// Read allowed_ip_addresses
		allowedIpaddressesList := make([]string, 0, len(storage.AllowedIpAddresses))
		for _, allowedIpaddress := range storage.AllowedIpAddresses {
			singleHost := make(map[string]interface{})
			singleHost["id"] = *allowedIpaddress.SubnetId
			singleHost["username"] = *allowedIpaddress.AllowedHost.Credential.Username
			singleHost["password"] = *allowedIpaddress.AllowedHost.Credential.Password
			singleHost["host_iqn"] = *allowedIpaddress.AllowedHost.Name
			allowedHostInfoList = append(allowedHostInfoList, singleHost)
			allowedIpaddressesList = append(allowedIpaddressesList, *allowedIpaddress.IpAddress)
		}
		d.Set("allowed_ip_addresses", allowedIpaddressesList)

		// Read allowed_virtual_guest_ids and allowed_host_info
		allowedVirtualGuestInfoList := make([]map[string]interface{}, 0)
		allowedVirtualGuestIdsList := make([]int, 0, len(storage.AllowedVirtualGuests))

		for _, allowedVirtualGuest := range storage.AllowedVirtualGuests {
			singleVirtualGuest := make(map[string]interface{})
			singleVirtualGuest["id"] = *allowedVirtualGuest.Id
			singleVirtualGuest["username"] = *allowedVirtualGuest.AllowedHost.Credential.Username
			singleVirtualGuest["password"] = *allowedVirtualGuest.AllowedHost.Credential.Password
			singleVirtualGuest["host_iqn"] = *allowedVirtualGuest.AllowedHost.Name
			allowedHostInfoList = append(allowedHostInfoList, singleVirtualGuest)
			allowedVirtualGuestInfoList = append(allowedVirtualGuestInfoList, singleVirtualGuest)
			allowedVirtualGuestIdsList = append(allowedVirtualGuestIdsList, *allowedVirtualGuest.Id)
		}
		d.Set("allowed_virtual_guest_ids", allowedVirtualGuestIdsList)
		d.Set("allowed_virtual_guest_info", allowedVirtualGuestInfoList)

		// Read allowed_hardware_ids and allowed_host_info
		allowedHardwareInfoList := make([]map[string]interface{}, 0)
		allowedHardwareIdsList := make([]int, 0, len(storage.AllowedHardware))
		for _, allowedHW := range storage.AllowedHardware {
			singleHardware := make(map[string]interface{})
			singleHardware["id"] = *allowedHW.Id
			singleHardware["username"] = *allowedHW.AllowedHost.Credential.Username
			singleHardware["password"] = *allowedHW.AllowedHost.Credential.Password
			singleHardware["host_iqn"] = *allowedHW.AllowedHost.Name
			allowedHostInfoList = append(allowedHostInfoList, singleHardware)
			allowedHardwareInfoList = append(allowedHardwareInfoList, singleHardware)
			allowedHardwareIdsList = append(allowedHardwareIdsList, *allowedHW.Id)
		}
		d.Set("allowed_hardware_ids", allowedHardwareIdsList)
		d.Set("allowed_hardware_info", allowedHardwareInfoList)
		d.Set("allowed_host_info", allowedHostInfoList)

		if storage.Notes != nil {
			d.Set("notes", *storage.Notes)
		}

		if storage.BillingItem != nil {
			d.Set("hourly_billing", storage.BillingItem.HourlyFlag)
		}

		return nil
	}
	return nil
}

func resourceIBMStorageBlockUpdate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	storage, err := services.GetNetworkStorageService(sess).
		Id(id).
		Mask(storageDetailMask).
		GetObject()

	if err != nil {
		return fmt.Errorf("Error updating storage information: %s", err)
	}

	// Update allowed_ip_addresses
	if d.HasChange("allowed_ip_addresses") {
		err := updateAllowedIpAddresses(d, sess, storage)
		if err != nil {
			return fmt.Errorf("Error updating storage information: %s", err)
		}
	}

	// Update allowed_subnets
	if d.HasChange("allowed_subnets") {
		err := updateAllowedSubnets(d, sess, storage)
		if err != nil {
			return fmt.Errorf("Error updating storage information: %s", err)
		}
	}

	// Update allowed_virtual_guest_ids
	if d.HasChange("allowed_virtual_guest_ids") {
		err := updateAllowedVirtualGuestIds(d, sess, storage)
		if err != nil {
			return fmt.Errorf("Error updating storage information: %s", err)
		}
	}

	// Update allowed_hardware_ids
	if d.HasChange("allowed_hardware_ids") {
		err := updateAllowedHardwareIds(d, sess, storage)
		if err != nil {
			return fmt.Errorf("Error updating storage information: %s", err)
		}
	}

	// Update notes
	if d.HasChange("notes") {
		err := updateNotes(d, sess, storage)
		if err != nil {
			return fmt.Errorf("Error updating storage information: %s", err)
		}
	}

	return resourceIBMStorageBlockRead(d, meta)
}

func resourceIBMStorageBlockDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMStorageFileDelete(d, meta)
}

func resourceIBMStorageBlockExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return resourceIBMStorageFileExists(d, meta)
}
