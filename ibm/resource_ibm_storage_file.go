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
	storagePackageType = "STORAGE_AS_A_SERVICE"
	storageMask        = "id,billingItem.orderItem.order.id"
	storageDetailMask  = "id,capacityGb,iops,storageType,username,serviceResourceBackendIpAddress,properties[type]" +
		",serviceResourceName,allowedIpAddresses[id,ipAddress,subnetId,allowedHost[name,credential[username,password]]],allowedSubnets[allowedHost[name,credential[username,password]]],allowedHardware[allowedHost[name,credential[username,password]]],allowedVirtualGuests[id,allowedHost[name,credential[username,password]]],snapshotCapacityGb,osType,notes,billingItem[hourlyFlag],serviceResource[datacenter[name]],schedules[dayOfWeek,hour,minute,retentionCount,type[keyname,name]]"
	itemMask        = "id,capacity,description,units,keyName,capacityMinimum,capacityMaximum,prices[id,categories[id,name,categoryCode],capacityRestrictionMinimum,capacityRestrictionMaximum,capacityRestrictionType,locationGroupId],itemCategory[categoryCode]"
	enduranceType   = "Endurance"
	performanceType = "Performance"
	fileStorage     = "file"
	blockStorage    = "block"
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
				Required: true,
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

	storageOrderContainer, err = buildStorageProductOrderContainer(sess, storageType, iops, capacity, snapshotCapacity, fileStorage, datacenter, hourlyBilling)
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

	default:
		return fmt.Errorf("Error during creation of storage: Invalid storageType %s", storageType)
	}

	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}

	// Find the storage device
	fileStorage, err := findStorageByOrderId(sess, *receipt.OrderId)

	if err != nil {
		return fmt.Errorf("Error during creation of storage: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", *fileStorage.Id))

	// Wait for storage availability
	_, err = WaitForStorageAvailable(d, meta)

	if err != nil {
		return fmt.Errorf(
			"Error waiting for storage (%s) to become ready: %s", d.Id(), err)
	}

	// SoftLayer changes the device ID after completion of provisioning. It is necessary to refresh device ID.
	fileStorage, err = findStorageByOrderId(sess, *receipt.OrderId)

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

	// Calculate IOPS
	iops, err := getIops(storage, storageType)
	if err != nil {
		return fmt.Errorf("Error retrieving storage information: %s", err)
	}
	d.Set("iops", iops)

	d.Set("type", storageType)
	d.Set("capacity", *storage.CapacityGb)
	d.Set("volumename", *storage.Username)
	d.Set("hostname", *storage.ServiceResourceBackendIpAddress)

	if storage.SnapshotCapacityGb != nil {
		snapshotCapacity, _ := strconv.Atoi(*storage.SnapshotCapacityGb)
		d.Set("snapshot_capacity", snapshotCapacity)
	}

	// Parse data center short name from ServiceResourceName. For example,
	// if SoftLayer API returns "'serviceResourceName': 'PerfStor Aggr aggr_staasdal0601_p01'",
	// the data center short name is "dal06".
	r, _ := regexp.Compile("[a-zA-Z]{3}[0-9]{2}")
	d.Set("datacenter", strings.ToLower(r.FindString(*storage.ServiceResourceName)))
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
	storageID, _ := strconv.Atoi(d.Id())

	// Get billing item associated with the storage
	billingItem, err := storageService.Id(storageID).GetBillingItem()

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

	_, err = services.GetNetworkStorageService(sess).
		Id(storageID).
		GetObject()

	if err != nil {
		if apiErr, ok := err.(sl.Error); ok && apiErr.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error retrieving storage information: %s", err)
	}
	return true, nil
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

	// Get a package type)
	pkg, err := product.GetPackageByType(sess, storagePackageType)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// Get all prices
	productItems, err := product.GetPackageProducts(sess, *pkg.Id, itemMask)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// Add IOPS price
	targetItemPrices := []datatypes.Product_Item_Price{}

	if storageType == "Performance" {
		price, err := getPriceByCategory(productItems, "storage_as_a_service")
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)
		price, err = getPriceByCategory(productItems, "storage_"+storageProtocol)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)

		price, err = getSaaSPerformSpacePrice(productItems, capacity)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)

		price, err = getSaaSPerformIOPSPrice(productItems, capacity, int(iops))
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)

	} else {

		price, err := getPriceByCategory(productItems, "storage_as_a_service")
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)
		price, err = getPriceByCategory(productItems, "storage_"+storageProtocol)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)

		price, err = getSaaSEnduranceSpacePrice(productItems, capacity, iops)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)

		price, err = getSaaSEnduranceTierPrice(productItems, iops)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)

	}

	if snapshotCapacity > 0 {
		price, err := getSaaSSnapshotSpacePrice(productItems, snapshotCapacity, iops, storageType)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		targetItemPrices = append(targetItemPrices, price)

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

func findStorageByOrderId(sess *session.Session, orderId int) (datatypes.Network_Storage, error) {
	filterPath := "networkStorage.billingItem.orderItem.order.id"

	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"complete"},
		Refresh: func() (interface{}, string, error) {
			storage, err := services.GetAccountService(sess).
				Filter(filter.Build(
					filter.Path(filterPath).
						Eq(strconv.Itoa(orderId)))).
				Mask(storageMask).
				GetNetworkStorage()
			if err != nil {
				return datatypes.Network_Storage{}, "", err
			}

			if len(storage) == 1 {
				return storage[0], "complete", nil
			} else if len(storage) == 0 {
				return nil, "pending", nil
			} else {
				return nil, "", fmt.Errorf("Expected one Storage: %s", err)
			}
		},
		Timeout:        45 * time.Minute,
		Delay:          10 * time.Second,
		MinTimeout:     10 * time.Second,
		NotFoundChecks: 300,
	}

	pendingResult, err := stateConf.WaitForState()

	if err != nil {
		return datatypes.Network_Storage{}, err
	}

	var result, ok = pendingResult.(datatypes.Network_Storage)

	if ok {
		return result, nil
	}

	return datatypes.Network_Storage{},
		fmt.Errorf("Cannot find Storage with order id '%d'", orderId)
}

// Waits for storage provisioning
func WaitForStorageAvailable(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	log.Printf("Waiting for storage (%s) to be available.", d.Id())
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, fmt.Errorf("The storage ID %s must be numeric", d.Id())
	}
	sess := meta.(ClientSession).SoftLayerSession()
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

			return result, "available", nil
		},
		Timeout:    45 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
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

func getPrice(prices []datatypes.Product_Item_Price, category, restrictionType string, restrictionValue int) datatypes.Product_Item_Price {
	for _, price := range prices {

		if price.LocationGroupId != nil || *price.Categories[0].CategoryCode != category {
			continue
		}

		if restrictionType != "" && restrictionValue > 0 {

			capacityRestrictionMinimum, _ := strconv.Atoi(*price.CapacityRestrictionMinimum)
			capacityRestrictionMaximum, _ := strconv.Atoi(*price.CapacityRestrictionMaximum)
			if restrictionType != *price.CapacityRestrictionType || restrictionValue < capacityRestrictionMinimum || restrictionValue > capacityRestrictionMaximum {
				continue
			}

		}

		return price

	}

	return datatypes.Product_Item_Price{}

}

func getPriceByCategory(productItems []datatypes.Product_Item, priceCategory string) (datatypes.Product_Item_Price, error) {
	for _, item := range productItems {
		price := getPrice(item.Prices, priceCategory, "", 0)
		if price.Id != nil {
			return price, nil
		}
	}

	return datatypes.Product_Item_Price{},
		fmt.Errorf("No product items matching with category %s could be found", priceCategory)
}

func getSaaSPerformSpacePrice(productItems []datatypes.Product_Item, size int) (datatypes.Product_Item_Price, error) {

	for _, item := range productItems {

		category, ok := sl.GrabOk(item, "ItemCategory.CategoryCode")
		if ok && category != "performance_storage_space" {
			continue
		}
		if item.CapacityMinimum == nil || item.CapacityMaximum == nil {
			continue
		}

		capacityMinimum, _ := strconv.Atoi(*item.CapacityMinimum)
		capacityMaximum, _ := strconv.Atoi(*item.CapacityMaximum)

		if size < capacityMinimum ||
			size > capacityMaximum {
			continue
		}

		keyname := fmt.Sprintf("%d_%d_GBS", capacityMinimum, capacityMaximum)
		if *item.KeyName != keyname {
			continue
		}

		price := getPrice(item.Prices, "performance_storage_space", "", 0)
		if price.Id != nil {
			return price, nil
		}
	}

	return datatypes.Product_Item_Price{},
		fmt.Errorf("Could not find price for performance storage space")

}

func getSaaSPerformIOPSPrice(productItems []datatypes.Product_Item, size, iops int) (datatypes.Product_Item_Price, error) {

	for _, item := range productItems {

		category, ok := sl.GrabOk(item, "ItemCategory.CategoryCode")
		if ok && category != "performance_storage_iops" {
			continue
		}

		if item.CapacityMinimum == nil || item.CapacityMaximum == nil {
			continue
		}

		capacityMinimum, _ := strconv.Atoi(*item.CapacityMinimum)
		capacityMaximum, _ := strconv.Atoi(*item.CapacityMaximum)

		if iops < capacityMinimum ||
			iops > capacityMaximum {
			continue
		}

		price := getPrice(item.Prices, "performance_storage_iops", "STORAGE_SPACE", size)
		if price.Id != nil {
			return price, nil
		}
	}

	return datatypes.Product_Item_Price{},
		fmt.Errorf("Could not find price for iops for the given volume")

}

func getSaaSEnduranceSpacePrice(productItems []datatypes.Product_Item, size int, iops float64) (datatypes.Product_Item_Price, error) {

	var keyName string
	if iops != 0.25 {
		tiers := int(iops)
		keyName = fmt.Sprintf("STORAGE_SPACE_FOR_%d_IOPS_PER_GB", tiers)
	} else {

		keyName = "STORAGE_SPACE_FOR_0_25_IOPS_PER_GB"

	}

	for _, item := range productItems {

		if *item.KeyName != keyName {
			continue
		}

		if item.CapacityMinimum == nil || item.CapacityMaximum == nil {
			continue
		}

		capacityMinimum, _ := strconv.Atoi(*item.CapacityMinimum)
		capacityMaximum, _ := strconv.Atoi(*item.CapacityMaximum)

		if size < capacityMinimum ||
			size > capacityMaximum {
			continue
		}

		price := getPrice(item.Prices, "performance_storage_space", "", 0)
		if price.Id != nil {
			return price, nil
		}
	}

	return datatypes.Product_Item_Price{},
		fmt.Errorf("Could not find price for endurance storage space")

}

func getSaaSEnduranceTierPrice(productItems []datatypes.Product_Item, iops float64) (datatypes.Product_Item_Price, error) {

	targetCapacity := enduranceCapacityRestrictionMap[iops]

	for _, item := range productItems {

		category, ok := sl.GrabOk(item, "ItemCategory.CategoryCode")
		if ok && category != "storage_tier_level" {
			continue
		}

		if int(*item.Capacity) != targetCapacity {
			continue
		}

		price := getPrice(item.Prices, "storage_tier_level", "", 0)
		if price.Id != nil {
			return price, nil
		}
	}

	return datatypes.Product_Item_Price{},
		fmt.Errorf("Could not find price for endurance tier level")

}

func getSaaSSnapshotSpacePrice(productItems []datatypes.Product_Item, size int, iops float64, volumeType string) (datatypes.Product_Item_Price, error) {

	var targetValue int
	var targetRestrictionType string
	if volumeType == "Performance" {
		targetValue = int(iops)
		targetRestrictionType = "IOPS"
	} else {

		targetValue = enduranceCapacityRestrictionMap[iops]
		targetRestrictionType = "STORAGE_TIER_LEVEL"

	}

	for _, item := range productItems {

		if int(*item.Capacity) != size {
			continue
		}

		price := getPrice(item.Prices, "storage_snapshot_space", targetRestrictionType, targetValue)
		if price.Id != nil {
			return price, nil
		}
	}

	return datatypes.Product_Item_Price{},
		fmt.Errorf("Could not find price for snapshot space")

}
