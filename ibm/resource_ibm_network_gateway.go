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
	"github.com/softlayer/softlayer-go/helpers/location"
	"github.com/softlayer/softlayer-go/helpers/product"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

const packageKeyName = "NETWORK_GATEWAY_APPLIANCE"

func resourceIBMNetworkGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMNetworkGatewayCreate,
		Read:     resourceIBMNetworkGatewayRead,
		Update:   resourceIBMNetworkGatewayUpdate,
		Delete:   resourceIBMNetworkGatewayDelete,
		Exists:   resourceIBMNetworkGatewayExists,
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

			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"datacenter": {
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

			"tcp_monitoring": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},

			"process_key_name": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Default:          "INTEL_SINGLE_XEON_1270_3_40_2",
				DiffSuppressFunc: applyOnce,
			},

			"os_key_name": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Default:          "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT",
				DiffSuppressFunc: applyOnce,
			},

			"redundant_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"unbonded_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"public_bandwidth": {
				Type:             schema.TypeInt,
				Optional:         true,
				ForceNew:         true,
				Default:          20000,
				DiffSuppressFunc: applyOnce,
			},

			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			// // Base for Gateway vlans array
			// "storage_groups": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	ForceNew: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"array_type_id": {
			// 				Type:     schema.TypeInt,
			// 				Required: true,
			// 			},
			// 			"hard_drives": {
			// 				Type:     schema.TypeList,
			// 				Elem:     &schema.Schema{Type: schema.TypeInt},
			// 				Required: true,
			// 			},
			// 			"array_size": {
			// 				Type:     schema.TypeInt,
			// 				Optional: true,
			// 			},
			// 			"partition_template_id": {
			// 				Type:     schema.TypeInt,
			// 				Optional: true,
			// 			},
			// 		},
			// 	},
			// 	DiffSuppressFunc: applyOnce,
			// },

			// Quote based provisioning only
			"quote_id": {
				Type:             schema.TypeInt,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},

			// Quote based provisioning, Monthly
			"public_vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			// Quote based provisioning, Monthly
			"private_vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"public_ipv4_address": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"private_ipv4_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},

			"ipv6_address": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"ipv6_address_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			// SoftLayer does not support public_ipv6_subnet configuration in vm creation. So, public_ipv6_subnet
			// is defined as a computed parameter.
			"public_ipv6_subnet": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"private_network_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"associated_vlans": {
				Type:        schema.TypeSet,
				Description: "The VLAN instances associated with this Network Gateway",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"networkVlanID": {
							Type:        schema.TypeInt,
							Description: "The Identifier of the VLAN to be associated",
							Optional:    true,
						},
						"bypass": {
							Type:        schema.TypeBool,
							Description: "Indicates if the VLAN should be in bypass or routed modes",
							Default:     true,
							Optional:    true,
						},
						"networkGatewayID": {
							Type:        schema.TypeInt,
							Description: "The identifier of the Network Gateway where the VLAN should be configured",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMNetworkGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	var order datatypes.Container_Product_Order
	var err error
	quote_id := d.Get("quote_id").(int)
	hardware := datatypes.Hardware{
		Hostname: sl.String(d.Get("hostname").(string)),
		Domain:   sl.String(d.Get("domain").(string)),
	}

	if quote_id > 0 {
		// Build a Network Gateway from a quote.
		order, err = services.GetBillingOrderQuoteService(sess).
			Id(quote_id).GetRecalculatedOrderContainer(nil, sl.Bool(false))
		if err != nil {
			return fmt.Errorf(
				"Encountered problem trying to get the Network Gateway order template from quote: %s", err)
		}
		order.Quantity = sl.Int(1)
		order.Hardware = make([]datatypes.Hardware, 0, 1)
		order.Hardware = append(
			order.Hardware,
			hardware,
		)
	} else {
		// Build a montly Network gateway
		order, err = getMonthlyGatewayOrder(d, meta)
		if err != nil {
			return fmt.Errorf(
				"Encountered problem trying to get the Gateway order template: %s", err)
		}
	}

	order, err = setCommonGatewayOrderOptions(d, meta, order)
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to configure Gateway options: %s", err)
	}

	var ProductOrder datatypes.Container_Product_Order
	ProductOrder.OrderContainers = make([]datatypes.Container_Product_Order, 1)
	ProductOrder.OrderContainers[0] = order

	_, err = services.GetProductOrderService(sess).VerifyOrder(&ProductOrder)
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to verify the order: %s", err)
	}
	_, err = services.GetProductOrderService(sess).PlaceOrder(&ProductOrder, sl.Bool(false))
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to place the order: %s", err)
	}

	log.Printf("[INFO] Gateway ID: %s", d.Id())

	bm, err := waitForNetworkGatewayProvision(&hardware, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for Gateway (%s) to become ready: %s", d.Id(), err)
	}

	id := *bm.(datatypes.Hardware).Id
	d.SetId(fmt.Sprintf("%d", id))

	if v, ok := d.GetOk("associated_vlans"); ok && v.(*schema.Set).Len() > 0 {
		debugvar := expandVlans(v.(*schema.Set).List())
		resourceIBMNetworkGatewayVlanAssociate(d, meta, debugvar, id)
	}

	// Set tags
	err = setHardwareTags(id, d, meta)
	if err != nil {
		return err
	}

	// Set notes
	if d.Get("notes").(string) != "" {
		err = setHardwareNotes(id, d, meta)
		if err != nil {
			return err
		}
	}

	return resourceIBMNetworkGatewayRead(d, meta)
}

func resourceIBMNetworkGatewayRead(d *schema.ResourceData, meta interface{}) error {
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
			"datacenter[id,name,longName]," +
			"primaryNetworkComponent[networkVlan[id,primaryRouter,vlanNumber],maxSpeed]," +
			"primaryBackendNetworkComponent[networkVlan[id,primaryRouter,vlanNumber],maxSpeed,redundancyEnabledFlag]," +
			"memoryCapacity,powerSupplyCount," +
			"operatingSystem[softwareLicense[softwareDescription[referenceCode]]]",
	).GetObject()

	if err != nil {
		return fmt.Errorf("Error retrieving Network Gateway: %s", err)
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

	if result.PrimaryNetworkComponent.NetworkVlan != nil {
		d.Set("public_vlan_id", *result.PrimaryNetworkComponent.NetworkVlan.Id)
	}

	if result.PrimaryBackendNetworkComponent.NetworkVlan != nil {
		d.Set("private_vlan_id", *result.PrimaryBackendNetworkComponent.NetworkVlan.Id)
	}

	d.Set("notes", sl.Get(result.Notes, nil))
	d.Set("memory", *result.MemoryCapacity)

	d.Set("redundant_network", false)
	d.Set("unbonded_network", false)

	backendNetworkComponent, err := service.Filter(
		filter.Build(
			filter.Path("backendNetworkComponents.status").Eq("ACTIVE"),
		),
	).Id(id).GetBackendNetworkComponents()

	if err != nil {
		return fmt.Errorf("Error retrieving Network Gateway network: %s", err)
	}

	if len(backendNetworkComponent) > 2 && result.PrimaryBackendNetworkComponent != nil {
		if *result.PrimaryBackendNetworkComponent.RedundancyEnabledFlag {
			d.Set("redundant_network", true)
		} else {
			d.Set("unbonded_network", true)
		}
	}

	tagReferences := result.TagReferences
	tagReferencesLen := len(tagReferences)
	if tagReferencesLen > 0 {
		tags := make([]string, 0, tagReferencesLen)
		for _, tagRef := range tagReferences {
			tags = append(tags, *tagRef.Tag.Name)
		}
		d.Set("tags", tags)
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

func resourceIBMNetworkGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
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

func resourceIBMNetworkGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetHardwareService(sess)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	_, err = waitForNoBareMetalActiveTransactions(id, meta)
	if err != nil {
		return fmt.Errorf("Error deleting Network Gateway while waiting for zero active transactions: %s", err)
	}

	billingItem, err := service.Id(id).GetBillingItem()

	if err != nil {
		return fmt.Errorf("Error getting billing item for Network Gateway: %s", err)
	}

	if billingItem.Id == nil {
		return fmt.Errorf("Error identifying the resource to delete, billing item is empty, please check the resource has not been deleted directly in Softlayer: %s", err)
	}

	// Monthly  Softlayer items only support an anniversary date cancellation option.
	billingItemService := services.GetBillingItemService(sess)
	_, err = billingItemService.Id(*billingItem.Id).CancelItem(
		sl.Bool(false), sl.Bool(true), sl.String("No longer required"), sl.String("Please cancel this Network Gateway"),
	)
	if err != nil {
		return fmt.Errorf("Error canceling the Network Gateway (%d): %s", id, err)
	}

	return nil
}

func resourceIBMNetworkGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	service := services.GetHardwareService(meta.(ClientSession).SoftLayerSession())

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	result, err := service.Id(id).GetObject()
	if err != nil {
		if apiErr, ok := err.(sl.Error); !ok || apiErr.StatusCode != 404 {
			return false, fmt.Errorf("Error trying to retrieve Network Gateway: %s", err)
		}
	}

	return result.Id != nil && *result.Id == id, nil
}

func getMonthlyGatewayOrder(d *schema.ResourceData, meta interface{}) (datatypes.Container_Product_Order, error) {
	sess := meta.(ClientSession).SoftLayerSession()

	// Validate attributes for network gateway ordering.
	model := packageKeyName

	datacenter, ok := d.GetOk("datacenter")
	if !ok {
		return datatypes.Container_Product_Order{}, fmt.Errorf("The attribute 'datacenter' is not defined.")
	}

	osKeyName := d.Get("os_key_name")

	process_key_name := d.Get("process_key_name")

	dc, err := location.GetDatacenterByName(sess, datacenter.(string), "id")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// 1. Find a package id using Gateway package key name.
	pkg, err := getPackageByModelGateway(sess, model)

	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	if pkg.Id == nil {
		return datatypes.Container_Product_Order{}, err
	}

	// 2. Get all prices for the package
	items, err := product.GetPackageProducts(sess, *pkg.Id, productItemMaskWithPriceLocationGroupID)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// 3. Build price items
	server, err := getItemPriceId(items, "server", process_key_name.(string))
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	os, err := getItemPriceId(items, "os", osKeyName.(string))
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	ram, err := findMemoryItemPriceId(items, d)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	portSpeed, err := findNetworkItemPriceId(items, d)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	monitoring, err := getItemPriceId(items, "monitoring", "MONITORING_HOST_PING")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	if d.Get("tcp_monitoring").(bool) {
		monitoring, err = getItemPriceId(items, "monitoring", "MONITORING_HOST_PING_AND_TCP_SERVICE")
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
	}
	// Other common default options
	priIpAddress, err := getItemPriceId(items, "pri_ip_addresses", "1_IP_ADDRESS")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	pri_ipv6_addresses, err := getItemPriceId(items, "pri_ipv6_addresses", "1_IPV6_ADDRESS")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	remoteManagement, err := getItemPriceId(items, "remote_management", "REBOOT_KVM_OVER_IP")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	vpnManagement, err := getItemPriceId(items, "vpn_management", "UNLIMITED_SSL_VPN_USERS_1_PPTP_VPN_USER_PER_ACCOUNT")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	notification, err := getItemPriceId(items, "notification", "NOTIFICATION_EMAIL_AND_TICKET")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	response, err := getItemPriceId(items, "response", "AUTOMATED_NOTIFICATION")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	vulnerabilityScanner, err := getItemPriceId(items, "vulnerability_scanner", "NESSUS_VULNERABILITY_ASSESSMENT_REPORTING")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// Define an order object using basic paramters.

	order := datatypes.Container_Product_Order{
		ContainerIdentifier: sl.String(d.Get("hostname").(string)),
		Quantity:            sl.Int(1),
		Hardware: []datatypes.Hardware{{
			Hostname: sl.String(d.Get("hostname").(string)),
			Domain:   sl.String(d.Get("domain").(string)),
		},
		},
		Location:  sl.String(strconv.Itoa(*dc.Id)),
		PackageId: pkg.Id,
		Prices: []datatypes.Product_Item_Price{
			server,
			os,
			ram,
			portSpeed,
			priIpAddress,
			pri_ipv6_addresses,
			remoteManagement,
			vpnManagement,
			monitoring,
			notification,
			response,
			vulnerabilityScanner,
		},
	}

	// Add optional price ids.
	// Add public bandwidth

	publicBandwidth := d.Get("public_bandwidth")
	publicBandwidthStr := "BANDWIDTH_" + strconv.Itoa(publicBandwidth.(int)) + "_GB"
	bandwidth, err := getItemPriceId(items, "bandwidth", publicBandwidthStr)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	order.Prices = append(order.Prices, bandwidth)

	// Add prices of disks.
	var arrayDrives []interface{}
	var arrayDrivestemp []string
	arrayDrivestemp = append(arrayDrivestemp, "HARD_DRIVE_1_00_TB_SATA_III")
	arrayDrivestemp = append(arrayDrivestemp, "HARD_DRIVE_2_00TB_SATA_II")

	for _, val := range arrayDrivestemp {
		arrayDrives = append(arrayDrives, val)
	}

	diskLen := len(arrayDrives)
	if diskLen > 0 {
		for i, disk := range arrayDrives {
			diskPrice, err := getItemPriceId(items, "disk"+strconv.Itoa(i), disk.(string))
			if err != nil {
				return datatypes.Container_Product_Order{}, err
			}
			order.Prices = append(order.Prices, diskPrice)
		}
	}

	// Add storage_groups for RAID configuration
	//hard coded support for disk controller, just RAID 1 is supported in this release of the resource
	diskController, err := getItemPriceId(items, "disk_controller", "DISK_CONTROLLER_RAID_1")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	order.Prices = append(order.Prices, diskController)

	return order, nil
}

func getPackageByModelGateway(sess *session.Session, model string) (datatypes.Product_Package, error) {
	objectMask := "id,keyName,name,description,isActive,type[keyName],categories[id,name,categoryCode]"
	service := services.GetProductPackageService(sess)
	availableModels := ""
	filterStr := "{\"items\": {\"categories\": {\"categoryCode\": {\"operation\":\"server\"}}},\"type\": {\"keyName\": {\"operation\":\"BARE_METAL_GATEWAY\"}}}"

	// Get package id
	packages, err := service.Mask(objectMask).
		Filter(filterStr).GetAllObjects()
	if err != nil {
		return datatypes.Product_Package{}, err
	}
	for _, pkg := range packages {
		availableModels = availableModels + *pkg.KeyName
		if pkg.Description != nil {
			availableModels = availableModels + " ( " + *pkg.Description + " ), "
		} else {
			availableModels = availableModels + ", "
		}
		if *pkg.KeyName == model {
			return pkg, nil
		}
	}
	return datatypes.Product_Package{}, fmt.Errorf("No Gateway package key name for %s. Available package key name(s) is(are) %s", model, availableModels)
}

func setCommonGatewayOrderOptions(d *schema.ResourceData, meta interface{}, order datatypes.Container_Product_Order) (datatypes.Container_Product_Order, error) {
	public_vlan_id := d.Get("public_vlan_id").(int)

	if public_vlan_id > 0 {
		order.Hardware[0].PrimaryNetworkComponent = &datatypes.Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{Id: sl.Int(public_vlan_id)},
		}
	}

	private_vlan_id := d.Get("private_vlan_id").(int)
	if private_vlan_id > 0 {
		order.Hardware[0].PrimaryBackendNetworkComponent = &datatypes.Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{Id: sl.Int(private_vlan_id)},
		}
	}

	return order, nil
}

func waitForNoGatewayActiveTransactions(id int, meta interface{}) (interface{}, error) {
	log.Printf("Waiting for Gateway (%d) to have zero active transactions", id)
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
			}
			return bm, "active", nil

		},
		Timeout:        24 * time.Hour,
		Delay:          10 * time.Second,
		MinTimeout:     1 * time.Minute,
		NotFoundChecks: 24 * 60,
	}

	return stateConf.WaitForState()
}

// Network gateways or Bare metal creation does not return a  object with an Id.
// Have to wait on provision date to become available on server that matches
// hostname and domain.
// http://sldn.softlayer.com/blog/bpotter/ordering-bare-metal-servers-using-softlayer-api
func waitForNetworkGatewayProvision(d *datatypes.Hardware, meta interface{}) (interface{}, error) {
	hostname := *d.Hostname
	domain := *d.Domain
	log.Printf("Waiting for Gateway (%s.%s) to be provisioned", hostname, domain)

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
		Timeout:        24 * time.Hour,
		Delay:          10 * time.Second,
		MinTimeout:     1 * time.Minute,
		NotFoundChecks: 24 * 60,
	}

	return stateConf.WaitForState()
}

func flattenVLANInstances(list []datatypes.Network_Gateway_Vlan) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"bypass":           *i.BypassFlag,
			"networkGatewayID": *i.NetworkGatewayId,
			"networkVlanID":    *i.NetworkVlanId,
		}
		result = append(result, l)
	}
	return result
}

func expandVlans(configured []interface{}) []datatypes.Network_Gateway_Vlan {
	vlans := make([]datatypes.Network_Gateway_Vlan, 0, len(configured))
	for _, lRaw := range configured {
		data := lRaw.(map[string]interface{})
		p := &datatypes.Network_Gateway_Vlan{}
		if v, ok := data["networkVlanID"]; ok && v.(int) != 0 {
			p.NetworkVlanId = sl.Int(v.(int))
		}
		if v, ok := data["bypass"]; ok {
			p.BypassFlag = sl.Bool(v.(bool))
		}
		if v, ok := data["networkGatewayID"]; ok && v.(int) != 0 {
			p.NetworkGatewayId = sl.Int(v.(int))
		}

		vlans = append(vlans, *p)
	}
	return vlans
}

func resourceIBMNetworkGatewayVlanAssociate(d *schema.ResourceData, meta interface{}, vlanObject []datatypes.Network_Gateway_Vlan, id int) error {
	sess := meta.(ClientSession).SoftLayerSession()
	processingarray := vlanObject

	for i := 0; i < len(processingarray); i++ {
		bypass := *processingarray[i].BypassFlag
		networkGatewayID := id
		networkVlanID := *processingarray[i].NetworkVlanId
		vlanObjectTemplate := datatypes.Network_Gateway_Vlan{
			BypassFlag:       &bypass,
			NetworkGatewayId: &networkGatewayID,
			NetworkVlanId:    &networkVlanID,
		}
		_, err := services.GetNetworkGatewayVlanService(sess).CreateObject(&vlanObjectTemplate)
		if err != nil {
			return fmt.Errorf(
				"Encountered problem trying to associate the VLAN %d : %s", networkVlanID, err)
		}
	}
	return nil

}
