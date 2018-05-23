package ibm

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/helpers/location"
	"github.com/softlayer/softlayer-go/helpers/product"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

const (
	storagePerformancePackageType = "STORAGE_AS_A_SERVICE"
	storageEndurancePackageType   = "STORAGE_AS_A_SERVICE"
	storagePortablePackageType    = "ADDITIONAL_SERVICES_PORTABLE_STORAGE"
	storageNasPackageType         = "ADDITIONAL_SERVICES_NETWORK_ATTACHED_STORAGE"
	storageMask                   = "id,billingItem.orderItem.order.id"
	storageDetailMask             = "id,capacityGb,iops,storageType,username,serviceResourceBackendIpAddress,properties[type]" +
		",serviceResourceName,allowedIpAddresses[ipAddress,subnetId,allowedHost[name,credential[username,password]]],allowedSubnets[allowedHost[name,credential[username,password]]],allowedHardware[allowedHost[name,credential[username,password]]],allowedVirtualGuests[id,allowedHost[name,credential[username,password]]],snapshotCapacityGb,osType,notes,billingItem[hourlyFlag],serviceResource[datacenter[name]],schedules[dayOfWeek,hour,minute,retentionCount,type[keyname,name]]"
	itemMask        = "id,capacity,description,units,keyName,prices[id,categories[id,name,categoryCode],capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]"
	enduranceType   = "Endurance"
	performanceType = "Performance"
	portableType    = "Portable"
	nasType         = "NAS/FTP"
	fileStorage     = "FILE_STORAGE"
	blockStorage    = "BLOCK_STORAGE"
	retryTime       = 5
)

var (
	// Map IOPS value to endurance storage tier keyName in SoftLayer_Product_Item
	enduranceIopsMap = map[float64]string{
		0.25: "LOW_INTENSITY_TIER",
		2:    "READHEAVY_TIER",
		4:    "WRITEHEAVY_TIER",
		10:   "10_IOPS_PER_GB",
	}

	// Map IOPS value to endurance storage tier capacityRestrictionMaximum/capacityRestrictionMinimum in SoftLayer_Product_Item
	enduranceCapacityRestrictionMap = map[float64]int{
		0.25: 100,
		2:    200,
		4:    300,
		10:   1000,
	}

	// Map IOPS value to endurance storage space keyName in SoftLayer_Product_Item
	enduranceStorageMap = map[float64]string{
		0.25: "STORAGE_SPACE_FOR_0_25_IOPS_PER_GB",
		2:    "STORAGE_SPACE_FOR_2_IOPS_PER_GB",
		4:    "STORAGE_SPACE_FOR_4_IOPS_PER_GB",
		10:   "STORAGE_SPACE_FOR_10_IOPS_PER_GB",
	}

	performanceStorageMap = map[int]string{
		20:    "20_39_GBS",
		40:    "40_79_GBS",
		80:    "80_99_GBS",
		100:   "100_499_GBS",
		250:   "100_499_GBS",
		500:   "500_999_GBS",
		1000:  "1000_1999_GBS",
		2000:  "2000_2999_GBS",
		3000:  "3000_3999_GBS",
		4000:  "4000_7999_GBS",
		5000:  "4000_7999_GBS",
		6000:  "4000_7999_GBS",
		7000:  "4000_7999_GBS",
		8000:  "8000_9999_GBS",
		9000:  "8000_9999_GBS",
		10000: "10000_12000_GBS",
		11000: "10000_12000_GBS",
		12000: "10000_12000_GBS",
	}

	portablestorageMap = map[int]string{
		10:   "GUEST_DISK_10_GB_SAN",
		20:   "GUEST_DISK_20_GB_SAN",
		25:   "GUEST_DISK_25_GB_SAN_4",
		30:   "GUEST_DISK_30_GB_SAN",
		40:   "GUEST_DISK_40_GB_SAN",
		50:   "GUEST_DISK_50_GB_SAN",
		75:   "GUEST_DISK_75_GB_SAN",
		100:  "GUEST_DISK_100_GB_SAN_3",
		125:  "GUEST_DISK_125_GB_SAN",
		150:  "GUEST_DISK_150_GB_SAN",
		175:  "GUEST_DISK_175_GB_SAN",
		200:  "GUEST_DISK_200_GB_SAN",
		250:  "GUEST_DISK_250_GB_SAN",
		300:  "GUEST_DISK_300_GB_SAN",
		350:  "GUEST_DISK_350_GB_SAN",
		400:  "GUEST_DISK_400_GB_SAN",
		500:  "GUEST_DISK_500_GB_SAN",
		750:  "GUEST_DISK_750_GB_SAN_2",
		1000: "GUEST_DISK_1000_GB_SAN_2",
		1500: "GUEST_DISK_1500_GB_SAN",
		2000: "GUEST_DISK_2000_GB_SAN",
	}

	// Map monthly storage value to performance IOPS keyName in SoftLayer_Product_Item
	performanceMonthlyIopsMap = map[int]string{
		20:    "100_1000_IOPS",
		40:    "100_2000_IOPS",
		80:    "100_4000_IOPS",
		100:   "100_6000_IOPS",
		250:   "100_6000_IOPS",
		500:   "100_6000_IOPS",
		1000:  "100_6000_IOPS",
		2000:  "200_6000_IOPS",
		3000:  "200_6000_IOPS",
		4000:  "300_6000_IOPS",
		5000:  "300_6000_IOPS",
		6000:  "300_6000_IOPS",
		7000:  "300_6000_IOPS",
		8000:  "500_6000_IOPS",
		9000:  "500_6000_IOPS",
		10000: "1000_6000_IOPS",
		11000: "1000_6000_IOPS",
		12000: "1000_6000_IOPS",
	}

	// Map hourly storage value to performance IOPS keyName in SoftLayer_Product_Item
	performanceHourlyIopsMap = map[int]string{
		20:    "100_1000_IOPS",
		40:    "100_2000_IOPS",
		80:    "100_4000_IOPS",
		100:   "100_6000_IOPS",
		250:   "100_6000_IOPS",
		500:   "100_10000_IOPS",
		1000:  "100_20000_IOPS",
		2000:  "200_40000_IOPS",
		3000:  "200_48000_IOPS",
		4000:  "300_48000_IOPS",
		5000:  "300_48000_IOPS",
		6000:  "300_48000_IOPS",
		7000:  "300_48000_IOPS",
		8000:  "500_48000_IOPS",
		9000:  "500_48000_IOPS",
		10000: "1000_48000_IOPS",
		11000: "1000_48000_IOPS",
		12000: "1000_48000_IOPS",
	}

	// storagePackageType is a storage package keyName for SoftLayer_Product_Package. It is used to filter storage package.
	// iopsCategoryCode is a storage IOPS categoryCode for SoftLayer_Product_Item. It is used to filter storage IOPS price.
	// storageProtocolCategoryCode is a storage protocol categoryCode for SoftLayer_Product_Item. It is used to filter storage protocol price.
	storagePackageMap = map[string](map[string](map[string]string)){
		fileStorage: {
			performanceType: {
				"storagePackageType":          storagePerformancePackageType,
				"iopsCategoryCode":            "performance_storage_iops",
				"storageProtocolCategoryCode": "storage_file",
			},
			enduranceType: {
				"storagePackageType":          storageEndurancePackageType,
				"iopsCategoryCode":            "storage_tier_level",
				"storageProtocolCategoryCode": "storage_file",
			},
		},
		blockStorage: {
			performanceType: {
				"storagePackageType":          storagePerformancePackageType,
				"iopsCategoryCode":            "performance_storage_iops",
				"storageProtocolCategoryCode": "storage_block",
			},
			enduranceType: {
				"storagePackageType":          storageEndurancePackageType,
				"iopsCategoryCode":            "storage_tier_level",
				"storageProtocolCategoryCode": "storage_block",
			},
			portableType: {
				"storagePackageType": storagePortablePackageType,
			},
		},
	}

	snapshotDay = map[string]string{
		"0": "SUNDAY",
		"1": "MONDAY",
		"2": "TUESDAY",
		"3": "WEDNESDAY",
		"4": "THURSDAY",
		"5": "FRIDAY",
		"6": "SATURDAY",
	}
)

func resourceIBMStorageFile() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMStorageFileCreate,
		Read:     resourceIBMStorageFileRead,
		Update:   resourceIBMStorageFileUpdate,
		Delete:   resourceIBMStorageFileDelete,
		Exists:   resourceIBMStorageFileExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateStorageType,
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

			"snapshot_capacity": {
				Type:     schema.TypeInt,
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

			"allowed_hardware_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Set: func(v interface{}) int {
					return v.(int)
				},
			},

			"allowed_subnets": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"allowed_ip_addresses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"snapshot_schedule": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 3,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateScheduleType,
						},

						"retention_count": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"minute": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateMinute(0, 59),
						},

						"hour": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateHour(0, 23),
						},

						"day_of_week": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateDayOfWeek,
						},

						"enable": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
				Set: resourceIBMFilSnapshotHash,
			},
			"mountpoint": {
				Type:     schema.TypeString,
				Computed: true,
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
		},
	}
}

func resourceIBMStorageFileCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()

	storageType := d.Get("type").(string)
	iops := d.Get("iops").(float64)
	datacenter := d.Get("datacenter").(string)
	capacity := d.Get("capacity").(int)
	snapshotCapacity := d.Get("snapshot_capacity").(int)
	hourlyBilling := d.Get("hourly_billing").(bool)

	var (
		storageOrderContainer datatypes.Container_Product_Order
		err                   error
	)

	if storageType == nasType {
		if _, ok := d.GetOk("iops"); ok {
			return fmt.Errorf("Error while creating storage: iops value can't be specified if type is %s", nasType)
		}
		storageOrderContainer, err = buildNasProductOrderContainer(sess, capacity, datacenter, hourlyBilling)
	} else {
		storageOrderContainer, err = buildStorageProductOrderContainer(sess, storageType, iops, capacity, snapshotCapacity, fileStorage, datacenter, hourlyBilling)
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
				VolumeSize:              &capacity,
			}, sl.Bool(false))
	case performanceType:
		receipt, err = services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(
			&datatypes.Container_Product_Order_Network_Storage_AsAService{
				Container_Product_Order: storageOrderContainer,
				VolumeSize:              &capacity,
				Iops:                    sl.Int(int(iops)),
			}, sl.Bool(false))
	case nasType:
		receipt, err = services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(
			&datatypes.Container_Product_Order_Network_Storage_Nas{
				Container_Product_Order: storageOrderContainer,
			}, sl.Bool(false))

	default:
		return fmt.Errorf("Error during creation of storage: Invalid storageType %s", storageType)
	}

	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}

	// Find the storage device
	fileStorage, _, err := findStorageByOrderId(sess, *receipt.OrderId, "")

	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", *fileStorage.Id))

	// Wait for storage availability
	_, err = WaitForStorageAvailable(d, meta, "")

	if err != nil {
		return fmt.Errorf(
			"Error waiting for storage (%s) to become ready: %s", d.Id(), err)
	}

	// SoftLayer changes the device ID after completion of provisioning. It is necessary to refresh device ID.
	fileStorage, _, err = findStorageByOrderId(sess, *receipt.OrderId, "")

	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", *fileStorage.Id))

	log.Printf("[INFO] Storage ID: %s", d.Id())

	return resourceIBMStorageFileUpdate(d, meta)
}

func resourceIBMStorageFileRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()

	storageId, _ := strconv.Atoi(d.Id())

	storage, err := services.GetNetworkStorageService(sess).
		Id(storageId).
		Mask(storageDetailMask).
		GetObject()

	if err != nil {
		return fmt.Errorf("Error retrieving storage information: %s", err)
	}

	storageType, err := getStorageTypeFromKeyName(*storage.StorageType.KeyName)
	if err != nil {
		return fmt.Errorf("Error retrieving storage information: %s", err)
	}

	if storageType != nasType {
		// Calculate IOPS
		iops, err := getIops(storage, storageType)
		if err != nil {
			return fmt.Errorf("Error retrieving storage information: %s", err)
		}
		d.Set("iops", iops)
	}

	d.Set("type", storageType)
	d.Set("capacity", *storage.CapacityGb)
	d.Set("volumename", *storage.Username)
	d.Set("hostname", *storage.ServiceResourceBackendIpAddress)

	if storage.SnapshotCapacityGb != nil {
		snapshotCapacity, _ := strconv.Atoi(*storage.SnapshotCapacityGb)
		d.Set("snapshot_capacity", snapshotCapacity)
	}

	if storageType == nasType {
		if storage.ServiceResource != nil {
			d.Set("datacenter", *storage.ServiceResource.Datacenter.Name)
		}
	} else {
		// Parse data center short name from ServiceResourceName. For example,
		// if SoftLayer API returns "'serviceResourceName': 'PerfStor Aggr aggr_staasdal0601_p01'",
		// the data center short name is "dal06".
		r, _ := regexp.Compile("[a-zA-Z]{3}[0-9]{2}")
		d.Set("datacenter", strings.ToLower(r.FindString(*storage.ServiceResourceName)))
	}
	// Read allowed_ip_addresses
	allowedIpaddressesList := make([]string, 0, len(storage.AllowedIpAddresses))
	for _, allowedIpaddress := range storage.AllowedIpAddresses {
		allowedIpaddressesList = append(allowedIpaddressesList, *allowedIpaddress.IpAddress)
	}
	d.Set("allowed_ip_addresses", allowedIpaddressesList)

	// Read allowed_subnets
	allowedSubnetsList := make([]string, 0, len(storage.AllowedSubnets))
	for _, allowedSubnets := range storage.AllowedSubnets {
		allowedSubnetsList = append(allowedSubnetsList, *allowedSubnets.NetworkIdentifier+"/"+strconv.Itoa(*allowedSubnets.Cidr))
	}
	d.Set("allowed_subnets", allowedSubnetsList)

	// Read allowed_virtual_guest_ids
	allowedVirtualGuestIdsList := make([]int, 0, len(storage.AllowedVirtualGuests))
	for _, allowedVirtualGuest := range storage.AllowedVirtualGuests {
		allowedVirtualGuestIdsList = append(allowedVirtualGuestIdsList, *allowedVirtualGuest.Id)
	}
	d.Set("allowed_virtual_guest_ids", allowedVirtualGuestIdsList)

	// Read allowed_hardware_ids
	allowedHardwareIdsList := make([]int, 0, len(storage.AllowedHardware))
	for _, allowedHW := range storage.AllowedHardware {
		allowedHardwareIdsList = append(allowedHardwareIdsList, *allowedHW.Id)
	}
	d.Set("allowed_hardware_ids", allowedHardwareIdsList)

	if storage.OsType != nil {
		d.Set("os_type", *storage.OsType.Name)
	}

	if storage.Notes != nil {
		d.Set("notes", *storage.Notes)
	}

	mountpoint, err := services.GetNetworkStorageService(sess).Id(storageId).GetFileNetworkMountAddress()
	if err != nil {
		return fmt.Errorf("Error retrieving storage information: %s", err)
	}
	d.Set("mountpoint", mountpoint)

	if storage.BillingItem != nil {
		d.Set("hourly_billing", storage.BillingItem.HourlyFlag)
	}

	schds := make([]interface{}, len(storage.Schedules))
	for i, schd := range storage.Schedules {
		s := make(map[string]interface{})
		s["retention_count"], _ = strconv.Atoi(*schd.RetentionCount)
		if *schd.Minute != "-1" {

			s["minute"], _ = strconv.Atoi(*schd.Minute)
		}
		if *schd.Hour != "-1" {
			s["hour"], _ = strconv.Atoi(*schd.Hour)
		}
		if *schd.Active > 0 {
			s["enable"], _ = strconv.ParseBool("true")
		} else {
			s["enable"], _ = strconv.ParseBool("false")
		}

		if *schd.DayOfWeek != "-1" {
			s["day_of_week"] = snapshotDay[*schd.DayOfWeek]
		}

		stype := *schd.Type.Keyname
		stype = stype[strings.LastIndex(stype, "_")+1:]
		s["schedule_type"] = stype
		schds[i] = s
	}
	d.Set("snapshot_schedule", schds)

	return nil
}

func resourceIBMStorageFileUpdate(d *schema.ResourceData, meta interface{}) error {
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

	// Enable Storage Snapshot Schedule
	if d.HasChange("snapshot_schedule") {
		err := enableStorageSnapshot(d, sess, storage)
		if err != nil {
			return fmt.Errorf("Error creating storage snapshot schedule: %s", err)
		}
	}

	return resourceIBMStorageFileRead(d, meta)
}

func resourceIBMStorageFileDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	storageService := services.GetNetworkStorageService(sess)
	var billingItem datatypes.Billing_Item
	var err error
	storageID, _ := strconv.Atoi(d.Id())
	if d.Get("type") == portableType {
		billingItems, err := services.GetVirtualDiskImageService(sess).Id(storageID).GetBillingItem()
		billingItem = billingItems.Billing_Item
		if err != nil {
			return fmt.Errorf("Error while looking up billing item associated with the storage: No billing item for ID:%d", storageID)
		}

	} else {
		// Get billing item associated with the storage
		billingItem, err = storageService.Id(storageID).GetBillingItem()
	}

	if err != nil {
		return fmt.Errorf("Error while looking up billing item associated with the storage: %s", err)
	}

	if billingItem.Id == nil {
		return fmt.Errorf("Error while looking up billing item associated with the storage: No billing item for ID:%d", storageID)
	}

	success, err := services.GetBillingItemService(sess).Id(*billingItem.Id).CancelService()
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("SoftLayer reported an unsuccessful cancellation")
	}
	return nil
}

func resourceIBMStorageFileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess := meta.(ClientSession).SoftLayerSession()

	storageID, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}
	if d.Get("type") == portableType {
		_, err = services.GetVirtualDiskImageService(sess).Id(storageID).GetObject()
	} else {
		_, err = services.GetNetworkStorageService(sess).
			Id(storageID).
			GetObject()
	}

	if err != nil {
		if apiErr, ok := err.(sl.Error); ok && apiErr.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error retrieving storage information: %s", err)
	}
	return true, nil
}

func buildNasProductOrderContainer(
	sess *session.Session,
	capacity int,
	datacenter string,
	hourlyBilling bool) (datatypes.Container_Product_Order, error) {

	pkg, err := product.GetPackageByType(sess, storageNasPackageType)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// Get all prices
	productItems, err := product.GetPackageProducts(sess, *pkg.Id, itemMask)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	targetItemPrices := []datatypes.Product_Item_Price{}

	capacityPrice, err := getPrice(productItems, fmt.Sprintf("NAS_%d_GB", capacity), "nas", "", 0)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// Lookup the data center ID
	dc, err := location.GetDatacenterByName(sess, datacenter)
	if err != nil {
		return datatypes.Container_Product_Order{},
			fmt.Errorf("No data centers matching %s could be found", datacenter)
	}

	targetItemPrices = append(targetItemPrices, capacityPrice)
	productOrderContainer := datatypes.Container_Product_Order{
		PackageId:        pkg.Id,
		Location:         sl.String(strconv.Itoa(*dc.Id)),
		Prices:           targetItemPrices,
		Quantity:         sl.Int(1),
		UseHourlyPricing: sl.Bool(hourlyBilling),
	}
	return productOrderContainer, nil
}
func buildStorageProductOrderContainer(
	sess *session.Session,
	storageType string,
	iops float64,
	capacity int,
	snapshotCapacity int,
	storageProtocol string,
	datacenter string,
	hourlyBilling bool) (datatypes.Container_Product_Order, error) {

	// Build product item filters for performance storage
	capacityKeyName, err := getCapacityKeyName(iops, capacity, storageType)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	// Get a package type)
	storagePackageType := storagePackageMap[storageProtocol][storageType]["storagePackageType"]
	pkg, err := product.GetPackageByType(sess, storagePackageType)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	// Get all prices
	productItems, err := product.GetPackageProducts(sess, *pkg.Id, itemMask)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	var capacityPrice datatypes.Product_Item_Price
	targetItemPrices := []datatypes.Product_Item_Price{}
	// Add capacity price
	if storageType == enduranceType {
		capacityPrice, err = getPrice(productItems, capacityKeyName, "performance_storage_space", "STORAGE_TIER_LEVEL", enduranceCapacityRestrictionMap[iops])
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
	} else if storageType == performanceType {
		capacityPrice, err = getPrice(productItems, capacityKeyName, "performance_storage_space", "", 0)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
	} else {
		capacityPrice, err = getPrice(productItems, capacityKeyName, "", "portablestorage", 0)
	}
	targetItemPrices = append(targetItemPrices, capacityPrice)

	if storageType != portableType {
		iopsKeyName, err := getIopsKeyName(iops, capacity, storageType, hourlyBilling)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		snapshotCapacityKeyName := fmt.Sprintf("%d_GB_", snapshotCapacity)

		iopsCategoryCode := storagePackageMap[storageProtocol][storageType]["iopsCategoryCode"]
		storageProtocolCategoryCode := storagePackageMap[storageProtocol][storageType]["storageProtocolCategoryCode"]

		// Add IOPS price
		var iopsPrice datatypes.Product_Item_Price

		if storageType == enduranceType {
			iopsPrice, err = getPrice(productItems, iopsKeyName, iopsCategoryCode, "", 0)
			if err != nil {
				return datatypes.Container_Product_Order{}, err
			}
		} else {
			iopsPrice, err = getPrice(productItems, iopsKeyName, iopsCategoryCode, "STORAGE_SPACE", capacity)
			if err != nil {
				return datatypes.Container_Product_Order{}, err
			}
		}

		targetItemPrices = append(targetItemPrices, iopsPrice)

		// Add storageProtocol price
		storageProtocolPrice, err := getPrice(productItems, storageProtocol, storageProtocolCategoryCode, "", 0)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, storageProtocolPrice)

		endurancePrice, err := getPrice(productItems, "STORAGE_AS_A_SERVICE", "storage_as_a_service", "", 0)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, endurancePrice)

		// Add snapshot capacity price
		if storageType == enduranceType && snapshotCapacity > 0 {
			snapshotCapacityPrice, err := getPrice(productItems, snapshotCapacityKeyName, "storage_snapshot_space", "STORAGE_TIER_LEVEL", enduranceCapacityRestrictionMap[iops])
			if err != nil {
				return datatypes.Container_Product_Order{}, err
			}
			targetItemPrices = append(targetItemPrices, snapshotCapacityPrice)
		}
	}
	// Lookup the data center ID
	dc, err := location.GetDatacenterByName(sess, datacenter)
	if err != nil {
		return datatypes.Container_Product_Order{},
			fmt.Errorf("No data centers matching %s could be found", datacenter)
	}
	productOrderContainer := datatypes.Container_Product_Order{
		PackageId:        pkg.Id,
		Location:         sl.String(strconv.Itoa(*dc.Id)),
		Prices:           targetItemPrices,
		Quantity:         sl.Int(1),
		UseHourlyPricing: sl.Bool(hourlyBilling),
	}

	return productOrderContainer, nil
}

func findStorageByOrderId(sess *session.Session, orderId int, storagetype string) (datatypes.Network_Storage, datatypes.Virtual_Disk_Image, error) {
	filterPath := "networkStorage.billingItem.orderItem.order.id"
	portablestoragefilter := "portableStorageVolumes.billingItem.orderItem.order.id"
	var storage []datatypes.Network_Storage
	var portablestorage []datatypes.Virtual_Disk_Image
	var err error
	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"complete"},
		Refresh: func() (interface{}, string, error) {
			if storagetype != portableType {
				storage, err = services.GetAccountService(sess).
					Filter(filter.Build(
						filter.Path(filterPath).
							Eq(strconv.Itoa(orderId)))).
					Mask(storageMask).
					GetNetworkStorage()
				if err != nil {
					return datatypes.Network_Storage{}, "", err
				}
			} else {
				portablestorage, err = services.GetAccountService(sess).
					Filter(filter.Build(
						filter.Path(portablestoragefilter).
							Eq(strconv.Itoa(orderId)))).
					Mask(storageMask).
					GetPortableStorageVolumes()
				if err != nil {
					return datatypes.Network_Storage{}, "", err
				}
			}
			if len(storage) == 1 {
				return storage[0], "complete", nil
			} else if len(portablestorage) == 1 {
				fmt.Println("----------------------------Finally Found it-------------------------")
				return portablestorage[0], "complete", nil
			} else if len(storage) == 0 || len(portablestorage) == 0 {
				return nil, "pending", nil
			}
			return nil, "", fmt.Errorf("Expected one Storage: %s", err)
		},
		Timeout:        45 * time.Minute,
		Delay:          10 * time.Second,
		MinTimeout:     10 * time.Second,
		NotFoundChecks: 300,
	}

	pendingResult, err := stateConf.WaitForState()

	if err != nil {
		return datatypes.Network_Storage{}, datatypes.Virtual_Disk_Image{}, err
	}

	var result, ok = pendingResult.(datatypes.Network_Storage)
	if storagetype == portableType {
		if result, ok := pendingResult.(datatypes.Virtual_Disk_Image); ok {
			return datatypes.Network_Storage{}, result, nil
		}
		return datatypes.Network_Storage{}, datatypes.Virtual_Disk_Image{},
			fmt.Errorf("Cannot find Storage with order id from line 856 '%d'", orderId)
	}

	if ok {
		return result, datatypes.Virtual_Disk_Image{}, nil
	}

	return datatypes.Network_Storage{}, datatypes.Virtual_Disk_Image{},
		fmt.Errorf("Cannot find Storage with order id '%d'", orderId)
}

// Waits for storage provisioning
func WaitForStorageAvailable(d *schema.ResourceData, meta interface{}, storagetype string) (interface{}, error) {
	log.Printf("Waiting for storage (%s) to be available.", d.Id())
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, fmt.Errorf("The storage ID %s must be numeric", d.Id())
	}
	sess := meta.(ClientSession).SoftLayerSession()
	storageType := d.Get("type").(string)
	stateConf := &resource.StateChangeConf{
		Pending: []string{"retry", "provisioning"},
		Target:  []string{"available"},
		Refresh: func() (interface{}, string, error) {
			// Check active transactions
			service := services.GetNetworkStorageService(sess)
			result, err := service.Id(id).Mask("activeTransactionCount").GetObject()
			if err != nil {
				if apiErr, ok := err.(sl.Error); ok && apiErr.StatusCode == 404 {
					return nil, "", fmt.Errorf("Error retrieving storage: %s", err)
				}
				return false, "retry", nil
			}

			log.Println("Checking active transactions.")
			if *result.ActiveTransactionCount > 0 {
				return result, "provisioning", nil
			}

			// Check volume status.
			if storageType != nasType || storagetype != portableType {
				log.Println("Checking volume status.")
				resultStr := ""
				err = sess.DoRequest(
					"SoftLayer_Network_Storage",
					"getObject",
					nil,
					&sl.Options{Id: &id, Mask: "volumeStatus"},
					&resultStr,
				)
				if err != nil {
					return false, "retry", nil
				}

				if !strings.Contains(resultStr, "PROVISION_COMPLETED") &&
					!strings.Contains(resultStr, "Volume Provisioning has completed") {
					return result, "provisioning", nil
				}
			}

			return result, "available", nil
		},
		Timeout:    45 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func getIopsKeyName(iops float64, capacity int, storageType string, hourlyBilling bool) (string, error) {
	switch storageType {
	case enduranceType:
		return enduranceIopsMap[iops], nil
	case performanceType:
		if hourlyBilling {
			return performanceHourlyIopsMap[capacity], nil
		}
		return performanceMonthlyIopsMap[capacity], nil
	}
	return "", fmt.Errorf("Invalid storageType %s.", storageType)
}

func getCapacityKeyName(iops float64, capacity int, storageType string) (string, error) {
	switch storageType {
	case enduranceType:
		return enduranceStorageMap[iops], nil
	case performanceType:
		return performanceStorageMap[capacity], nil
	case portableType:
		return portablestorageMap[capacity], nil
	}
	return "", fmt.Errorf("Invalid storageType %s.", storageType)
}

func getPrice(productItems []datatypes.Product_Item, keyName string, categoryCode string, capacityRestrictionType string, capacityRestriction int) (datatypes.Product_Item_Price, error) {
	for _, item := range productItems {
		if strings.HasPrefix(*item.KeyName, keyName) {
			for _, price := range item.Prices {
				if price.LocationGroupId == nil && capacityRestrictionType == "portablestorage" {
					return price, nil
				}
				if *price.Categories[0].CategoryCode == categoryCode && price.LocationGroupId == nil {
					if capacityRestrictionType == "STORAGE_SPACE" {
						if price.CapacityRestrictionMinimum == nil ||
							price.CapacityRestrictionMaximum == nil {
							continue
						}
						capacityRestrictionMinimum, _ := strconv.Atoi(*price.CapacityRestrictionMinimum)
						capacityRestrictionMaximum, _ := strconv.Atoi(*price.CapacityRestrictionMaximum)
						if capacityRestrictionMinimum > 0 &&
							capacityRestriction >= capacityRestrictionMinimum &&
							capacityRestriction <= capacityRestrictionMaximum {
							return price, nil
						}
					}

					if capacityRestrictionType == "STORAGE_TIER_LEVEL" {
						if price.CapacityRestrictionMinimum == nil ||
							price.CapacityRestrictionMaximum == nil {
							continue
						}
						capacityRestrictionMinimum, _ := strconv.Atoi(*price.CapacityRestrictionMinimum)
						capacityRestrictionMaximum, _ := strconv.Atoi(*price.CapacityRestrictionMaximum)
						if capacityRestrictionMinimum > 0 &&
							capacityRestriction == capacityRestrictionMinimum &&
							capacityRestriction == capacityRestrictionMaximum {
							return price, nil
						}
					}

					if capacityRestrictionType == "" && capacityRestriction == 0 {
						return price, nil
					}
				}
			}
		}
	}
	return datatypes.Product_Item_Price{},
		fmt.Errorf("No product items matching with keyName %s and categoryCode %s could be found", keyName, categoryCode)
}

func getIops(storage datatypes.Network_Storage, storageType string) (float64, error) {
	switch storageType {
	case enduranceType:
		for _, property := range storage.Properties {
			if *property.Type.Keyname == "PROVISIONED_IOPS" {
				provisionedIops, err := strconv.Atoi(*property.Value)
				if err != nil {
					return 0, err
				}
				enduranceIops := float64(provisionedIops / *storage.CapacityGb)
				if enduranceIops < 1 {
					enduranceIops = 0.25
				}
				return enduranceIops, nil
			}
		}
	case performanceType:
		if storage.Iops == nil {
			return 0, fmt.Errorf("Failed to retrieve iops information.")
		}
		iops, err := strconv.Atoi(*storage.Iops)
		if err != nil {
			return 0, err
		}
		return float64(iops), nil
	}
	return 0, fmt.Errorf("Invalid storage type %s", storageType)
}

func updateAllowedIpAddresses(d *schema.ResourceData, sess *session.Session, storage datatypes.Network_Storage) error {
	id := *storage.Id
	newIps := d.Get("allowed_ip_addresses").(*schema.Set).List()

	// Add new allowed_ip_addresses
	for _, newIp := range newIps {
		isNewIp := true
		for _, oldAllowedIpAddresses := range storage.AllowedIpAddresses {
			if newIp.(string) == *oldAllowedIpAddresses.IpAddress {
				isNewIp = false
				break
			}
		}
		if isNewIp {
			ipObject, err := services.GetAccountService(sess).
				Filter(filter.Build(
					filter.Path("ipAddresses.ipAddress").
						Eq(newIp.(string)))).GetIpAddresses()
			if err != nil {
				return err
			}
			if len(ipObject) != 1 {
				return fmt.Errorf("Number of IP address is %d", len(ipObject))
			}
			for {
				_, err = services.GetNetworkStorageService(sess).
					Id(id).
					AllowAccessFromHostList([]datatypes.Container_Network_Storage_Host{
						{
							Id:         ipObject[0].Id,
							ObjectType: sl.String("SoftLayer_Network_Subnet_IpAddress"),
						},
					})
				if err != nil {
					if strings.Contains(err.Error(), "SoftLayer_Exception_Network_Storage_Group_MassAccessControlModification") {
						time.Sleep(retryTime * time.Second)
						continue
					}
					return err
				}
				break
			}
		}
	}

	// Remove deleted allowed_hardware_ids
	for _, oldAllowedIpAddresses := range storage.AllowedIpAddresses {
		isDeletedId := true
		for _, newIp := range newIps {
			if newIp.(string) == *oldAllowedIpAddresses.IpAddress {
				isDeletedId = false
				break
			}
		}
		if isDeletedId {
			for {
				_, err := services.GetNetworkStorageService(sess).
					Id(id).
					RemoveAccessFromHostList([]datatypes.Container_Network_Storage_Host{
						{
							Id:         oldAllowedIpAddresses.Id,
							ObjectType: sl.String("SoftLayer_Network_Subnet_IpAddress"),
						},
					})
				if err != nil {
					if strings.Contains(err.Error(), "SoftLayer_Exception_Network_Storage_Group_MassAccessControlModification") {
						time.Sleep(retryTime * time.Second)
						continue
					}
					return err
				}
				break
			}
		}
	}
	return nil
}

func updateAllowedSubnets(d *schema.ResourceData, sess *session.Session, storage datatypes.Network_Storage) error {
	id := *storage.Id
	newSubnets := d.Get("allowed_subnets").(*schema.Set).List()

	// Add new allowed_subnets
	for _, newSubnet := range newSubnets {
		isNewSubnet := true
		newSubnetArr := strings.Split(newSubnet.(string), "/")
		newNetworkIdentifier := newSubnetArr[0]
		newCidr, err := strconv.Atoi(newSubnetArr[1])
		if err != nil {
			return err
		}
		for _, oldAllowedSubnets := range storage.AllowedSubnets {
			if newNetworkIdentifier == *oldAllowedSubnets.NetworkIdentifier && newCidr == *oldAllowedSubnets.Cidr {
				isNewSubnet = false
				break
			}
		}
		if isNewSubnet {
			filterStr := fmt.Sprintf("{\"subnets\":{\"networkIdentifier\":{\"operation\":\"%s\"},\"cidr\":{\"operation\":\"%d\"}}}", newNetworkIdentifier, newCidr)
			subnetObject, err := services.GetAccountService(sess).
				Filter(filterStr).GetSubnets()
			if err != nil {
				return err
			}
			if len(subnetObject) != 1 {
				return fmt.Errorf("Number of subnet is %d", len(subnetObject))
			}
			_, err = services.GetNetworkStorageService(sess).
				Id(id).
				AllowAccessFromHostList([]datatypes.Container_Network_Storage_Host{
					{
						Id:         subnetObject[0].Id,
						ObjectType: sl.String("SoftLayer_Network_Subnet"),
					},
				})
			if err != nil {
				return err
			}
		}
	}

	// Remove deleted allowed_subnets
	for _, oldAllowedSubnets := range storage.AllowedSubnets {
		isDeletedSubnet := true
		for _, newSubnet := range newSubnets {
			newSubnetArr := strings.Split(newSubnet.(string), "/")
			newNetworkIdentifier := newSubnetArr[0]
			newCidr, err := strconv.Atoi(newSubnetArr[1])
			if err != nil {
				return err
			}

			if newNetworkIdentifier == *oldAllowedSubnets.NetworkIdentifier && newCidr == *oldAllowedSubnets.Cidr {
				isDeletedSubnet = false
				break
			}
		}
		if isDeletedSubnet {
			_, err := services.GetNetworkStorageService(sess).
				Id(id).
				RemoveAccessFromHostList([]datatypes.Container_Network_Storage_Host{
					{
						Id:         sl.Int(*oldAllowedSubnets.Id),
						ObjectType: sl.String("SoftLayer_Network_Subnet"),
					},
				})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func updateAllowedVirtualGuestIds(d *schema.ResourceData, sess *session.Session, storage datatypes.Network_Storage) error {
	id := *storage.Id
	newIds := d.Get("allowed_virtual_guest_ids").(*schema.Set).List()

	// Add new allowed_virtual_guest_ids
	for _, newId := range newIds {
		isNewId := true
		for _, oldAllowedVirtualGuest := range storage.AllowedVirtualGuests {
			if newId.(int) == *oldAllowedVirtualGuest.Id {
				isNewId = false
				break
			}
		}
		if isNewId {
			for {
				_, err := services.GetNetworkStorageService(sess).
					Id(id).
					AllowAccessFromHostList([]datatypes.Container_Network_Storage_Host{
						{
							Id:         sl.Int(newId.(int)),
							ObjectType: sl.String("SoftLayer_Virtual_Guest"),
						},
					})
				if err != nil {
					if strings.Contains(err.Error(), "SoftLayer_Exception_Network_Storage_Group_MassAccessControlModification") {
						time.Sleep(retryTime * time.Second)
						continue
					}
					return err
				}
				break
			}
		}
	}

	// Remove deleted allowed_virtual_guest_ids
	for _, oldAllowedVirtualGuest := range storage.AllowedVirtualGuests {
		isDeletedId := true
		for _, newId := range newIds {
			if newId.(int) == *oldAllowedVirtualGuest.Id {
				isDeletedId = false
				break
			}
		}
		if isDeletedId {
			for {
				_, err := services.GetNetworkStorageService(sess).
					Id(id).
					RemoveAccessFromHostList([]datatypes.Container_Network_Storage_Host{
						{
							Id:         sl.Int(*oldAllowedVirtualGuest.Id),
							ObjectType: sl.String("SoftLayer_Virtual_Guest"),
						},
					})
				if err != nil {
					if strings.Contains(err.Error(), "SoftLayer_Exception_Network_Storage_Group_MassAccessControlModification") {
						time.Sleep(retryTime * time.Second)
						continue
					}
					return err
				}
				break
			}
		}
	}
	return nil
}

func updateAllowedHardwareIds(d *schema.ResourceData, sess *session.Session, storage datatypes.Network_Storage) error {
	id := *storage.Id
	newIds := d.Get("allowed_hardware_ids").(*schema.Set).List()

	// Add new allowed_hardware_ids
	for _, newId := range newIds {
		isNewId := true
		for _, oldAllowedHardware := range storage.AllowedHardware {
			if newId.(int) == *oldAllowedHardware.Id {
				isNewId = false
				break
			}
		}
		if isNewId {
			_, err := services.GetNetworkStorageService(sess).
				Id(id).
				AllowAccessFromHostList([]datatypes.Container_Network_Storage_Host{
					{
						Id:         sl.Int(newId.(int)),
						ObjectType: sl.String("SoftLayer_Hardware"),
					},
				})
			if err != nil {
				return err
			}
		}
	}

	// Remove deleted allowed_hardware_ids
	for _, oldAllowedHardware := range storage.AllowedHardware {
		isDeletedId := true
		for _, newId := range newIds {
			if newId.(int) == *oldAllowedHardware.Id {
				isDeletedId = false
				break
			}
		}
		if isDeletedId {
			_, err := services.GetNetworkStorageService(sess).
				Id(id).
				RemoveAccessFromHostList([]datatypes.Container_Network_Storage_Host{
					{
						Id:         sl.Int(*oldAllowedHardware.Id),
						ObjectType: sl.String("SoftLayer_Hardware"),
					},
				})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func enableStorageSnapshot(d *schema.ResourceData, sess *session.Session, storage datatypes.Network_Storage) error {
	id := *storage.Id
	for _, e := range d.Get("snapshot_schedule").(*schema.Set).List() {
		value := e.(map[string]interface{})
		enable := value["enable"].(bool)
		_, err := services.GetNetworkStorageService(sess).
			Id(id).
			EnableSnapshots(sl.String(value["schedule_type"].(string)), sl.Int(value["retention_count"].(int)), sl.Int(value["minute"].(int)), sl.Int(value["hour"].(int)), sl.String(value["day_of_week"].(string)))
		if err != nil {
			return err
		}
		if !enable {
			_, err := services.GetNetworkStorageService(sess).
				Id(id).
				DisableSnapshots(sl.String(value["schedule_type"].(string)))
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func updateNotes(d *schema.ResourceData, sess *session.Session, storage datatypes.Network_Storage) error {
	id := *storage.Id
	notes := d.Get("notes").(string)

	if (storage.Notes != nil && *storage.Notes != notes) || (storage.Notes == nil && notes != "") {
		_, err := services.GetNetworkStorageService(sess).
			Id(id).
			EditObject(&datatypes.Network_Storage{Notes: sl.String(notes)})
		if err != nil {
			return fmt.Errorf("Error adding note to storage (%d): %s", id, err)
		}
	}

	return nil
}

func getStorageTypeFromKeyName(key string) (string, error) {
	switch key {
	case "FILE_STORAGE":
		return nasType, nil
	case "ENDURANCE_FILE_STORAGE":
		return enduranceType, nil
	case "PERFORMANCE_FILE_STORAGE":
		return performanceType, nil
	}
	return "", fmt.Errorf("Couldn't find storage type for key %s", key)
}

func resourceIBMFilSnapshotHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-",
		m["schedule_type"].(string)))
	buf.WriteString(fmt.Sprintf("%s-",
		m["day_of_week"].(string)))
	buf.WriteString(fmt.Sprintf("%d-",
		m["hour"].(int)))

	buf.WriteString(fmt.Sprintf("%d-",
		m["minute"].(int)))

	buf.WriteString(fmt.Sprintf("%d-",
		m["retention_count"].(int)))

	return hashcode.String(buf.String())
}
