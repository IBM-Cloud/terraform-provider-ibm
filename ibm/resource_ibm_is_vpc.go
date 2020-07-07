package ibm

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isVPCDefaultNetworkACL       = "default_network_acl"
	isVPCIsDefault               = "is_default"
	isVPCIDefaultSecurityGroup   = "default_security_group"
	isVPCName                    = "name"
	isVPCResourceGroup           = "resource_group"
	isVPCStatus                  = "status"
	isVPCDeleting                = "deleting"
	isVPCDeleted                 = "done"
	isVPCTags                    = "tags"
	isVPCClassicAccess           = "classic_access"
	isVPCAvailable               = "available"
	isVPCFailed                  = "failed"
	isVPCPending                 = "pending"
	isVPCAddressPrefixManagement = "address_prefix_management"
	cseSourceAddresses           = "cse_source_addresses"
	subnetsList                  = "subnets"
	totalIPV4AddressCount        = "total_ipv4_address_count"
	availableIPV4AddressCount    = "available_ipv4_address_count"
	isVPCCRN                     = "crn"
)

func resourceIBMISVPC() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVPCCreate,
		Read:     resourceIBMISVPCRead,
		Update:   resourceIBMISVPCUpdate,
		Delete:   resourceIBMISVPCDelete,
		Exists:   resourceIBMISVPCExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isVPCAddressPrefixManagement: {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "auto",
				DiffSuppressFunc: applyOnce,
				ValidateFunc:     InvokeValidator("ibm_is_vpc", isVPCAddressPrefixManagement),
				Description:      "Address Prefix management value",
			},

			isVPCDefaultNetworkACL: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     nil,
				Computed:    true,
				Deprecated:  "This field is deprecated",
				Description: "Default network ACL",
			},

			isVPCIsDefault: {
				Type:     schema.TypeBool,
				ForceNew: true,
				Optional: true,
				Removed:  "This field is removed use classic_access",
			},

			isVPCClassicAccess: {
				Type:        schema.TypeBool,
				ForceNew:    true,
				Default:     false,
				Optional:    true,
				Description: "Set to true if classic access needs to enabled to VPC",
			},

			isVPCName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
				Description:  "VPC name",
			},

			isVPCResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group info",
			},

			isVPCStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPC status",
			},

			isVPCIDefaultSecurityGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security group associated with VPC",
			},
			isVPCTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "List of tags",
			},

			isVPCCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
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

			cseSourceAddresses: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud service endpoint IP Address",
						},

						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location info of CSE Address",
						},
					},
				},
			},

			subnetsList: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subent name",
						},

						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subnet ID",
						},

						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subnet status",
						},

						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subnet location",
						},

						totalIPV4AddressCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total IPv4 address count in the subnet",
						},

						availableIPV4AddressCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Available IPv4 address count in the subnet",
						},
					},
				},
			},
		},
	}
}

func resourceIBMISVPCValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)
	address_prefix_management := "auto, manual"

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isVPCAddressPrefixManagement,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Optional:                   true,
			Default:                    "auto",
			AllowedValues:              address_prefix_management})

	ibmISVPCResourceValidator := ResourceValidator{ResourceName: "ibm_is_vpc", Schema: validateSchema}
	return &ibmISVPCResourceValidator
}

func resourceIBMISVPCCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] VPC create")
	name := d.Get(isVPCName).(string)
	apm := ""
	rg := ""
	isClassic := false

	if addprefixmgmt, ok := d.GetOk(isVPCAddressPrefixManagement); ok {
		apm = addprefixmgmt.(string)
	}
	if classic, ok := d.GetOk(isVPCClassicAccess); ok {
		isClassic = classic.(bool)
	}

	if grp, ok := d.GetOk(isVPCResourceGroup); ok {
		rg = grp.(string)
	}
	if userDetails.generation == 1 {
		err := classicVpcCreate(d, meta, name, apm, rg, isClassic)
		if err != nil {
			return err
		}
	} else {
		err := vpcCreate(d, meta, name, apm, rg, isClassic)
		if err != nil {
			return err
		}
	}
	return resourceIBMISVPCRead(d, meta)
}

func classicVpcCreate(d *schema.ResourceData, meta interface{}, name, apm, rg string, isClassic bool) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.CreateVpcOptions{
		Name: &name,
	}
	if rg != "" {
		options.ResourceGroup = &vpcclassicv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	if apm != "" {
		options.AddressPrefixManagement = &apm
	}
	options.ClassicAccess = &isClassic

	vpc, response, err := sess.CreateVpc(options)
	if err != nil {
		return fmt.Errorf("Error while creating VPC err %s\n%s", err, response)
	}
	d.SetId(*vpc.ID)
	log.Printf("[INFO] VPC : %s", *vpc.ID)
	_, err = isWaitForClassicVPCAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVPCTags); ok || v != "" {
		oldList, newList := d.GetChange(isVPCTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *vpc.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func isWaitForClassicVPCAvailable(vpc *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPC (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isVPCPending},
		Target:     []string{isVPCAvailable, isVPCFailed},
		Refresh:    isClassicVPCRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicVPCRefreshFunc(vpc *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getvpcOptions := &vpcclassicv1.GetVpcOptions{
			ID: &id,
		}
		vpc, response, err := vpc.GetVpc(getvpcOptions)
		if err != nil {
			return nil, isVPCFailed, fmt.Errorf("Error getting VPC : %s\n%s", err, response)
		}

		if *vpc.Status == isVPCAvailable || *vpc.Status == isVPCFailed {
			return vpc, *vpc.Status, nil
		}

		return vpc, isVPCPending, nil
	}
}

func vpcCreate(d *schema.ResourceData, meta interface{}, name, apm, rg string, isClassic bool) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateVpcOptions{
		Name: &name,
	}
	if rg != "" {
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	if apm != "" {
		options.AddressPrefixManagement = &apm
	}
	options.ClassicAccess = &isClassic

	vpc, response, err := sess.CreateVpc(options)
	if err != nil {
		return fmt.Errorf("Error while creating VPC err %s\n%s", err, response)
	}
	d.SetId(*vpc.ID)
	log.Printf("[INFO] VPC : %s", *vpc.ID)
	_, err = isWaitForVPCAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVPCTags); ok || v != "" {
		oldList, newList := d.GetChange(isVPCTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *vpc.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func isWaitForVPCAvailable(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPC (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isVPCPending},
		Target:     []string{isVPCAvailable, isVPCFailed},
		Refresh:    isVPCRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPCRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getvpcOptions := &vpcv1.GetVpcOptions{
			ID: &id,
		}
		vpc, response, err := vpc.GetVpc(getvpcOptions)
		if err != nil {
			return nil, isVPCFailed, fmt.Errorf("Error getting VPC : %s\n%s", err, response)
		}

		if *vpc.Status == isVPCAvailable || *vpc.Status == isVPCFailed {
			return vpc, *vpc.Status, nil
		}

		return vpc, isVPCPending, nil
	}
}

func resourceIBMISVPCRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicVpcGet(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := vpcGet(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicVpcGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getvpcOptions := &vpcclassicv1.GetVpcOptions{
		ID: &id,
	}
	vpc, response, err := sess.GetVpc(getvpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting VPC : %s\n%s", err, response)
	}

	d.Set("id", *vpc.ID)
	d.Set(isVPCName, *vpc.Name)
	d.Set(isVPCClassicAccess, *vpc.ClassicAccess)
	d.Set(isVPCStatus, *vpc.Status)
	if vpc.DefaultNetworkAcl != nil {
		log.Printf("[DEBUG] vpc default network acl is not null :%s", *vpc.DefaultNetworkAcl.ID)
		d.Set(isVPCDefaultNetworkACL, *vpc.DefaultNetworkAcl.ID)
	} else {
		log.Printf("[DEBUG] vpc default network acl is  null")
		d.Set(isVPCDefaultNetworkACL, nil)
	}
	if vpc.DefaultSecurityGroup != nil {
		d.Set(isVPCIDefaultSecurityGroup, *vpc.DefaultSecurityGroup.ID)
	} else {
		d.Set(isVPCIDefaultSecurityGroup, nil)
	}
	tags, err := GetTagsUsingCRN(meta, *vpc.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc (%s) tags: %s", d.Id(), err)
	}
	d.Set(isVPCTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(isVPCCRN, *vpc.Crn)
	d.Set(ResourceControllerURL, controller+"/vpc/network/vpcs")
	d.Set(ResourceName, *vpc.Name)
	d.Set(ResourceCRN, *vpc.Crn)
	d.Set(ResourceStatus, *vpc.Status)
	if vpc.ResourceGroup != nil {
		d.Set(isVPCResourceGroup, *vpc.ResourceGroup.ID)
		d.Set(ResourceGroupName, *vpc.ResourceGroup.ID)
	}
	//set the cse ip addresses info
	if vpc.CseSourceIps != nil {
		cseSourceIpsList := make([]map[string]interface{}, 0)
		for _, sourceIP := range vpc.CseSourceIps {
			currentCseSourceIp := map[string]interface{}{}
			if sourceIP.Ip != nil {
				currentCseSourceIp["address"] = *sourceIP.Ip.Address
				currentCseSourceIp["zone_name"] = *sourceIP.Zone.Name
				cseSourceIpsList = append(cseSourceIpsList, currentCseSourceIp)
			}
		}
		d.Set(cseSourceAddresses, cseSourceIpsList)
	}
	// set the subnets list
	options := &vpcclassicv1.ListSubnetsOptions{}
	s, response, err := sess.ListSubnets(options)
	if err != nil {
		return fmt.Errorf("Error Fetching subnets %s\n%s", err, response)
	} else {
		subnetsInfo := make([]map[string]interface{}, 0)
		for _, subnet := range s.Subnets {
			if *subnet.Vpc.ID == d.Id() {
				l := map[string]interface{}{
					"name":                    *subnet.Name,
					"id":                      *subnet.ID,
					"status":                  *subnet.Status,
					"zone":                    *subnet.Zone.Name,
					totalIPV4AddressCount:     *subnet.TotalIpv4AddressCount,
					availableIPV4AddressCount: *subnet.AvailableIpv4AddressCount,
				}
				subnetsInfo = append(subnetsInfo, l)
			}
		}
		d.Set(subnetsList, subnetsInfo)
	}
	return nil
}

func vpcGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getvpcOptions := &vpcv1.GetVpcOptions{
		ID: &id,
	}
	vpc, response, err := sess.GetVpc(getvpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting VPC : %s\n%s", err, response)
	}

	d.Set("id", *vpc.ID)
	d.Set(isVPCName, *vpc.Name)
	d.Set(isVPCClassicAccess, *vpc.ClassicAccess)
	d.Set(isVPCStatus, *vpc.Status)
	if vpc.DefaultNetworkAcl != nil {
		log.Printf("[DEBUG] vpc default network acl is not null :%s", *vpc.DefaultNetworkAcl.ID)
		d.Set(isVPCDefaultNetworkACL, *vpc.DefaultNetworkAcl.ID)
	} else {
		log.Printf("[DEBUG] vpc default network acl is  null")
		d.Set(isVPCDefaultNetworkACL, nil)
	}
	if vpc.DefaultSecurityGroup != nil {
		d.Set(isVPCIDefaultSecurityGroup, *vpc.DefaultSecurityGroup.ID)
	} else {
		d.Set(isVPCIDefaultSecurityGroup, nil)
	}
	tags, err := GetTagsUsingCRN(meta, *vpc.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc (%s) tags: %s", d.Id(), err)
	}
	d.Set(isVPCTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(isVPCCRN, *vpc.Crn)
	d.Set(ResourceControllerURL, controller+"/vpc-ext/network/vpcs")
	d.Set(ResourceName, *vpc.Name)
	d.Set(ResourceCRN, *vpc.Crn)
	d.Set(ResourceStatus, *vpc.Status)
	if vpc.ResourceGroup != nil {
		d.Set(isVPCResourceGroup, *vpc.ResourceGroup.ID)
		d.Set(ResourceGroupName, *vpc.ResourceGroup.Name)
	}
	//set the cse ip addresses info
	if vpc.CseSourceIps != nil {
		cseSourceIpsList := make([]map[string]interface{}, 0)
		for _, sourceIP := range vpc.CseSourceIps {
			currentCseSourceIp := map[string]interface{}{}
			if sourceIP.Ip != nil {
				currentCseSourceIp["address"] = *sourceIP.Ip.Address
				currentCseSourceIp["zone_name"] = *sourceIP.Zone.Name
				cseSourceIpsList = append(cseSourceIpsList, currentCseSourceIp)
			}
		}
		d.Set(cseSourceAddresses, cseSourceIpsList)
	}
	// set the subnets list
	options := &vpcv1.ListSubnetsOptions{}
	s, response, err := sess.ListSubnets(options)
	if err != nil {
		return fmt.Errorf("Error Fetching subnets %s\n%s", err, response)
	} else {
		subnetsInfo := make([]map[string]interface{}, 0)
		for _, subnet := range s.Subnets {
			if *subnet.Vpc.ID == d.Id() {
				l := map[string]interface{}{
					"name":                    *subnet.Name,
					"id":                      *subnet.ID,
					"status":                  *subnet.Status,
					"zone":                    *subnet.Zone.Name,
					totalIPV4AddressCount:     *subnet.TotalIpv4AddressCount,
					availableIPV4AddressCount: *subnet.AvailableIpv4AddressCount,
				}
				subnetsInfo = append(subnetsInfo, l)
			}
		}
		d.Set(subnetsList, subnetsInfo)
	}
	return nil
}

func resourceIBMISVPCUpdate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()

	name := ""
	hasChanged := false

	if d.HasChange(isVPCName) {
		name = d.Get(isVPCName).(string)
		hasChanged = true
	}
	if userDetails.generation == 1 {
		err := classicVpcUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	} else {
		err := vpcUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	}
	return resourceIBMISVPCRead(d, meta)
}

func classicVpcUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isVPCTags) {
		getvpcOptions := &vpcclassicv1.GetVpcOptions{
			ID: &id,
		}
		vpc, response, err := sess.GetVpc(getvpcOptions)
		if err != nil {
			return fmt.Errorf("Error getting VPC : %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isVPCTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *vpc.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc (%s) tags: %s", id, err)
		}
	}
	if hasChanged {
		updateVpcOptions := &vpcclassicv1.UpdateVpcOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateVpc(updateVpcOptions)
		if err != nil {
			return fmt.Errorf("Error Updating VPC : %s\n%s", err, response)
		}
	}
	return nil
}

func vpcUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isVPCTags) {
		getvpcOptions := &vpcv1.GetVpcOptions{
			ID: &id,
		}
		vpc, response, err := sess.GetVpc(getvpcOptions)
		if err != nil {
			return fmt.Errorf("Error getting VPC : %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isVPCTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *vpc.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc (%s) tags: %s", d.Id(), err)
		}
	}
	if hasChanged {
		updateVpcOptions := &vpcv1.UpdateVpcOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateVpc(updateVpcOptions)
		if err != nil {
			return fmt.Errorf("Error Updating VPC : %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISVPCDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicVpcDelete(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := vpcDelete(d, meta, id)
		if err != nil {
			return err
		}
	}
	d.SetId("")
	return nil
}

func classicVpcDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	getVpcOptions := &vpcclassicv1.GetVpcOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpc(getVpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error Getting VPC (%s): %s\n%s", id, err, response)

	}

	deletevpcOptions := &vpcclassicv1.DeleteVpcOptions{
		ID: &id,
	}
	response, err = sess.DeleteVpc(deletevpcOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting VPC : %s\n%s", err, response)
	}
	_, err = isWaitForClassicVPCDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func vpcDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getVpcOptions := &vpcv1.GetVpcOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpc(getVpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting VPC (%s): %s\n%s", id, err, response)
	}

	deletevpcOptions := &vpcv1.DeleteVpcOptions{
		ID: &id,
	}
	response, err = sess.DeleteVpc(deletevpcOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting VPC : %s\n%s", err, response)
	}
	_, err = isWaitForVPCDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForClassicVPCDeleted(vpc *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPC (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPCDeleting},
		Target:     []string{isVPCDeleted, isVPCFailed},
		Refresh:    isClassicVPCDeleteRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicVPCDeleteRefreshFunc(vpc *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		getvpcOptions := &vpcclassicv1.GetVpcOptions{
			ID: &id,
		}
		vpc, response, err := vpc.GetVpc(getvpcOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return vpc, isVPCDeleted, nil
			}
			return nil, isVPCFailed, fmt.Errorf("The VPC %s failed to delete: %s\n%s", id, err, response)
		}

		return vpc, isVPCDeleting, nil
	}
}

func isWaitForVPCDeleted(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPC (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPCDeleting},
		Target:     []string{isVPCDeleted, isVPCFailed},
		Refresh:    isVPCDeleteRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPCDeleteRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		getvpcOptions := &vpcv1.GetVpcOptions{
			ID: &id,
		}
		vpc, response, err := vpc.GetVpc(getvpcOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return vpc, isVPCDeleted, nil
			}
			return nil, isVPCFailed, fmt.Errorf("The VPC %s failed to delete: %s\n%s", id, err, response)
		}

		return vpc, isVPCDeleting, nil
	}
}

func resourceIBMISVPCExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		exists, err := classicVpcExists(d, meta, id)
		return exists, err
	} else {
		exists, err := vpcExists(d, meta, id)
		return exists, err
	}
}

func classicVpcExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return false, err
	}
	getvpcOptions := &vpcclassicv1.GetVpcOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpc(getvpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting VPC: %s\n%s", err, response)
	}

	return true, nil
}

func vpcExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	getvpcOptions := &vpcv1.GetVpcOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpc(getvpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting VPC: %s\n%s", err, response)
	}
	return true, nil
}

func resourceIBMVPCHash(v interface{}) int {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s",
		strings.ToLower(v.(string))))
	return hashcode.String(buf.String())
}
