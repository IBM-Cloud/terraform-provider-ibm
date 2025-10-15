// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSubnetIpv4CidrBlock             = "ipv4_cidr_block"
	isSubnetTotalIpv4AddressCount     = "total_ipv4_address_count"
	isSubnetIPVersion                 = "ip_version"
	isSubnetName                      = "name"
	isSubnetTags                      = "tags"
	isSubnetCRN                       = "crn"
	isSubnetNetworkACL                = "network_acl"
	isSubnetPublicGateway             = "public_gateway"
	isSubnetStatus                    = "status"
	isSubnetVPC                       = "vpc"
	isSubnetVPCName                   = "vpc_name"
	isSubnetZone                      = "zone"
	isSubnetAvailableIpv4AddressCount = "available_ipv4_address_count"
	isSubnetResourceGroup             = "resource_group"

	isSubnetProvisioning     = "provisioning"
	isSubnetProvisioningDone = "done"
	isSubnetDeleting         = "deleting"
	isSubnetDeleted          = "done"
	isSubnetRoutingTableID   = "routing_table"
	isSubnetRoutingTableCrn  = "routing_table_crn"
	isSubnetInUse            = "resources_attached"
	isSubnetAccessTags       = "access_tags"
	isUserTagType            = "user"
	isAccessTagType          = "access"
)

func ResourceIBMISSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISSubnetCreate,
		ReadContext:   resourceIBMISSubnetRead,
		UpdateContext: resourceIBMISSubnetUpdate,
		DeleteContext: resourceIBMISSubnetDelete,
		Exists:        resourceIBMISSubnetExists,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				}),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{
			isSubnetIpv4CidrBlock: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isSubnetTotalIpv4AddressCount},
				ValidateFunc:  validate.InvokeValidator("ibm_is_subnet", isSubnetIpv4CidrBlock),
				Description:   "IPV4 subnet - CIDR block",
			},

			isSubnetAvailableIpv4AddressCount: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of IPv4 addresses in this subnet that are not in-use, and have not been reserved by the user or the provider.",
			},

			isSubnetTotalIpv4AddressCount: {
				Type:          schema.TypeInt,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isSubnetIpv4CidrBlock},
				Description:   "The total number of IPv4 addresses in this subnet.",
			},
			isSubnetIPVersion: {
				Type:         schema.TypeString,
				ForceNew:     true,
				Default:      "ipv4",
				Optional:     true,
				ValidateFunc: validate.ValidateIPVersion,
				Description:  "The IP version(s) to support for this subnet.",
			},

			isSubnetName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_subnet", isSubnetName),
				Description:  "Subnet name",
			},

			isSubnetTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_subnet", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			isSubnetAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_subnet", isSubnetAccessTags)},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},

			isSubnetCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			isSubnetNetworkACL: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    false,
				Description: "The network ACL for this subnet",
			},

			isSubnetPublicGateway: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    false,
				Description: "Public Gateway of the subnet",
			},

			isSubnetStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the subnet",
			},

			isSubnetVPC: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "VPC instance ID",
			},

			isSubnetZone: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Subnet zone info",
			},

			isSubnetResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "The resource group for this subnet",
			},
			isSubnetRoutingTableID: {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{isSubnetRoutingTableCrn},
				Computed:      true,
				Description:   "routing table id that is associated with the subnet",
			},
			isSubnetRoutingTableCrn: {
				Type:          schema.TypeString,
				Computed:      true,
				ForceNew:      true,
				Optional:      true,
				ConflictsWith: []string{isSubnetRoutingTableID},
				Description:   "routing table crn that is associated with the subnet.",
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func ResourceIBMISSubnetValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSubnetName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSubnetIpv4CidrBlock,
			ValidateFunctionIdentifier: validate.ValidateCIDRAddress,
			Type:                       validate.TypeString,
			ForceNew:                   true,
			Optional:                   true})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSubnetAccessTags,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISSubnetResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_subnet", Schema: validateSchema}
	return &ibmISSubnetResourceValidator
}

func resourceIBMISSubnetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get(isSubnetName).(string)
	vpc := d.Get(isSubnetVPC).(string)
	zone := d.Get(isSubnetZone).(string)

	ipv4cidr := ""
	if cidr, ok := d.GetOk(isSubnetIpv4CidrBlock); ok {
		ipv4cidr = cidr.(string)
	}
	ipv4addrcount64 := int64(0)
	ipv4addrcount := 0
	if ipv4addrct, ok := d.GetOk(isSubnetTotalIpv4AddressCount); ok {
		ipv4addrcount = ipv4addrct.(int)
		ipv4addrcount64 = int64(ipv4addrcount)
	}
	if ipv4cidr == "" && ipv4addrcount == 0 {
		err := fmt.Errorf("%s or %s need to be provided", isSubnetIpv4CidrBlock, isSubnetTotalIpv4AddressCount)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "create", "parse-ipv4").GetDiag()
	}

	if ipv4cidr != "" && ipv4addrcount != 0 {
		err := fmt.Errorf("only one of %s or %s needs to be provided", isSubnetIpv4CidrBlock, isSubnetTotalIpv4AddressCount)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "create", "parse-ipv4").GetDiag()
	}
	isSubnetKey := "subnet_key_" + vpc + "_" + zone
	conns.IbmMutexKV.Lock(isSubnetKey)
	defer conns.IbmMutexKV.Unlock(isSubnetKey)

	acl := ""
	if nwacl, ok := d.GetOk(isSubnetNetworkACL); ok {
		acl = nwacl.(string)
	}

	gw := ""
	if pgw, ok := d.GetOk(isSubnetPublicGateway); ok {
		gw = pgw.(string)
	}

	// route table association related
	rtID := ""
	if rt, ok := d.GetOk(isSubnetRoutingTableID); ok {
		rtID = rt.(string)
	}

	rtCrn := ""
	if rtcrn, ok := d.GetOk(isSubnetRoutingTableCrn); ok {
		rtCrn = rtcrn.(string)
	}

	err := subnetCreate(context, d, meta, name, vpc, zone, ipv4cidr, acl, gw, rtID, rtCrn, ipv4addrcount64)
	if err != nil {
		return err
	}

	return resourceIBMISSubnetRead(context, d, meta)
}

func subnetCreate(context context.Context, d *schema.ResourceData, meta interface{}, name, vpc, zone, ipv4cidr, acl, gw, rtID, rtCrn string, ipv4addrcount64 int64) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	subnetTemplate := &vpcv1.SubnetPrototype{
		Name: &name,
		VPC: &vpcv1.VPCIdentity{
			ID: &vpc,
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
	}
	if ipv4cidr != "" {
		subnetTemplate.Ipv4CIDRBlock = &ipv4cidr
	}
	if ipv4addrcount64 != int64(0) {
		subnetTemplate.TotalIpv4AddressCount = &ipv4addrcount64
	}
	if gw != "" {
		subnetTemplate.PublicGateway = &vpcv1.PublicGatewayIdentity{
			ID: &gw,
		}
	}

	if acl != "" {
		subnetTemplate.NetworkACL = &vpcv1.NetworkACLIdentity{
			ID: &acl,
		}
	}
	if rtID != "" {
		rt := rtID
		subnetTemplate.RoutingTable = &vpcv1.RoutingTableIdentity{
			ID: &rt,
		}
	}
	if rtCrn != "" {
		subnetTemplate.RoutingTable = &vpcv1.RoutingTableIdentity{
			CRN: &rtCrn,
		}
	}

	rg := ""
	if grp, ok := d.GetOk(isSubnetResourceGroup); ok {
		rg = grp.(string)
		subnetTemplate.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	//create a subnet
	createSubnetOptions := &vpcv1.CreateSubnetOptions{
		SubnetPrototype: subnetTemplate,
	}
	subnet, _, err := sess.CreateSubnetWithContext(context, createSubnetOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*subnet.ID)
	log.Printf("[INFO] Subnet : %s", *subnet.ID)
	_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSubnetAvailable failed: %s", err.Error()), "ibm_is_subnet", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isSubnetTags); ok || v != "" {
		oldList, newList := d.GetChange(isSubnetTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *subnet.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource subnet (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(isSubnetAccessTags); ok {
		oldList, newList := d.GetChange(isSubnetAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *subnet.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource subnet (%s) access tags: %s", d.Id(), err)
		}
	}

	return nil
}

func isWaitForSubnetAvailable(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetProvisioning},
		Target:     []string{isSubnetProvisioningDone, ""},
		Refresh:    isSubnetRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetRefreshFunc(subnetC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSubnetOptions := &vpcv1.GetSubnetOptions{
			ID: &id,
		}
		subnet, response, err := subnetC.GetSubnet(getSubnetOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting Subnet : %s\n%s", err, response)
		}

		if *subnet.Status == "available" || *subnet.Status == "failed" {
			return subnet, isSubnetProvisioningDone, nil
		}

		return subnet, isSubnetProvisioning, nil
	}
}

func resourceIBMISSubnetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()

	err := subnetGet(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func subnetGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnetWithContext(context, getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(subnet.Name) {
		if err = d.Set("name", subnet.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(subnet.IPVersion) {
		if err = d.Set("ip_version", subnet.IPVersion); err != nil {
			err = fmt.Errorf("Error setting ip_version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-ip_version").GetDiag()
		}
	}
	if !core.IsNil(subnet.Ipv4CIDRBlock) {
		if err = d.Set("ipv4_cidr_block", subnet.Ipv4CIDRBlock); err != nil {
			err = fmt.Errorf("Error setting ipv4_cidr_block: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-ipv4_cidr_block").GetDiag()
		}
	}
	if err = d.Set("available_ipv4_address_count", flex.IntValue(subnet.AvailableIpv4AddressCount)); err != nil {
		err = fmt.Errorf("Error setting available_ipv4_address_count: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-available_ipv4_address_count").GetDiag()
	}
	if !core.IsNil(subnet.TotalIpv4AddressCount) {
		if err = d.Set("total_ipv4_address_count", flex.IntValue(subnet.TotalIpv4AddressCount)); err != nil {
			err = fmt.Errorf("Error setting total_ipv4_address_count: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-total_ipv4_address_count").GetDiag()
		}
	}
	if subnet.NetworkACL != nil {
		if err = d.Set(isSubnetNetworkACL, *subnet.NetworkACL.ID); err != nil {
			err = fmt.Errorf("Error setting network_acl: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-network_acl").GetDiag()
		}
	}
	if subnet.PublicGateway != nil {
		if err = d.Set(isSubnetPublicGateway, *subnet.PublicGateway.ID); err != nil {
			err = fmt.Errorf("Error setting public_gateway: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-public_gateway").GetDiag()
		}
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	if subnet.RoutingTable != nil {
		if err = d.Set(isSubnetRoutingTableID, *subnet.RoutingTable.ID); err != nil {
			err = fmt.Errorf("Error setting routing_table: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-routing_table").GetDiag()
		}
		if err = d.Set(isSubnetRoutingTableCrn, *subnet.RoutingTable.CRN); err != nil {
			err = fmt.Errorf("Error setting routing_table_crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-routing_table_crn").GetDiag()
		}
	} else {
		d.Set(isSubnetRoutingTableID, nil)
		d.Set(isSubnetRoutingTableCrn, nil)
	}
	if err = d.Set("status", subnet.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-status").GetDiag()
	}
	if !core.IsNil(subnet.Zone) {
		if err = d.Set(isSubnetZone, *subnet.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-zone").GetDiag()
		}
	}
	if err = d.Set(isSubnetVPC, *subnet.VPC.ID); err != nil {
		err = fmt.Errorf("Error setting vpc: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-vpc").GetDiag()
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_subnet", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	tags, err := flex.GetGlobalTagsUsingCRN(meta, *subnet.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource subnet (%s) tags: %s", d.Id(), err)
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *subnet.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource subnet (%s) access tags: %s", d.Id(), err)
	}

	if err = d.Set(isSubnetTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-tags").GetDiag()
	}
	if err = d.Set(isSubnetAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-access_tags").GetDiag()
	}
	if err = d.Set("crn", subnet.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-crn").GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/subnets"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *subnet.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set("resource_crn", subnet.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set(flex.ResourceStatus, *subnet.Status); err != nil {
		err = fmt.Errorf("Error setting resource_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-resource_status").GetDiag()
	}
	if subnet.ResourceGroup != nil {
		if err = d.Set(isSubnetResourceGroup, *subnet.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, *subnet.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "read", "set-resource_group_name").GetDiag()
		}
	}
	return nil
}

func resourceIBMISSubnetUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	if d.HasChange(isSubnetTags) {
		oldList, newList := d.GetChange(isSubnetTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isSubnetCRN).(string), "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource subnet (%s) tags: %s", d.Id(), err)
		}
	}

	if d.HasChange(isSubnetAccessTags) {
		oldList, newList := d.GetChange(isSubnetAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isSubnetCRN).(string), "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource subnet (%s) access tags: %s", d.Id(), err)
		}
	}

	err := subnetUpdate(context, d, meta, id)
	if err != nil {
		return err
	}

	return resourceIBMISSubnetRead(context, d, meta)
}

func subnetUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	hasChanged := false
	name := ""
	acl := ""
	updateSubnetOptions := &vpcv1.UpdateSubnetOptions{}
	subnetPatchModel := &vpcv1.SubnetPatch{}
	if d.HasChange(isSubnetName) {
		name = d.Get(isSubnetName).(string)
		subnetPatchModel.Name = &name
		hasChanged = true
	}
	if d.HasChange(isSubnetNetworkACL) {
		acl = d.Get(isSubnetNetworkACL).(string)
		subnetPatchModel.NetworkACL = &vpcv1.NetworkACLIdentity{
			ID: &acl,
		}
		hasChanged = true
	}
	if d.HasChange(isSubnetPublicGateway) {
		gw := d.Get(isSubnetPublicGateway).(string)
		if gw == "" {
			unsetSubnetPublicGatewayOptions := &vpcv1.UnsetSubnetPublicGatewayOptions{
				ID: &id,
			}
			_, err = sess.UnsetSubnetPublicGatewayWithContext(context, unsetSubnetPublicGatewayOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UnsetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_subnet", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSubnetAvailable failed: %s", err.Error()), "ibm_is_subnet", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		} else {
			setSubnetPublicGatewayOptions := &vpcv1.SetSubnetPublicGatewayOptions{
				ID: &id,
				PublicGatewayIdentity: &vpcv1.PublicGatewayIdentity{
					ID: &gw,
				},
			}
			_, _, err = sess.SetSubnetPublicGatewayWithContext(context, setSubnetPublicGatewayOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_subnet", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSubnetAvailable failed: %s", err.Error()), "ibm_is_subnet", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	}
	if d.HasChange(isSubnetRoutingTableID) {
		hasChanged = true
		rtID := d.Get(isSubnetRoutingTableID).(string)
		// Construct an instance of the RoutingTableIdentityByID model
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = &rtID
		subnetPatchModel.RoutingTable = routingTableIdentityModel
		/*rt := &vpcv1.RoutingTableIdentity{
			ID: corev3.StringPtr(rtID),
		}
		setSubnetRoutingTableBindingOptions := sess.NewReplaceSubnetRoutingTableOptions(id, rt)
		setSubnetRoutingTableBindingOptions.SetRoutingTableIdentity(rt)
		setSubnetRoutingTableBindingOptions.SetID(id)
		_, _, err = sess.ReplaceSubnetRoutingTable(setSubnetRoutingTableBindingOptions)
		if err != nil {
			log.Printf("SetSubnetRoutingTableBinding eroor: %s", err)
			return err
		}*/
	}
	if d.HasChange(isSubnetRoutingTableCrn) {
		hasChanged = true
		rtCrn := d.Get(isSubnetRoutingTableCrn).(string)
		// Construct an instance of the RoutingTableIdentityByCRN model
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByCRN)
		routingTableIdentityModel.CRN = &rtCrn
		subnetPatchModel.RoutingTable = routingTableIdentityModel
	}
	if hasChanged {
		subnetPatch, err := subnetPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("subnetPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_subnet", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateSubnetOptions.SubnetPatch = subnetPatch
		updateSubnetOptions.ID = &id
		_, _, err = sess.UpdateSubnetWithContext(context, updateSubnetOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISSubnetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	err := subnetDelete(context, d, meta, id)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func subnetDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()

	}
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnetWithContext(context, getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if subnet.PublicGateway != nil {
		unsetSubnetPublicGatewayOptions := &vpcv1.UnsetSubnetPublicGatewayOptions{
			ID: &id,
		}
		_, err = sess.UnsetSubnetPublicGatewayWithContext(context, unsetSubnetPublicGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UnsetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_subnet", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSubnetAvailable failed: %s", err.Error()), "ibm_is_subnet", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	deleteSubnetOptions := &vpcv1.DeleteSubnetOptions{
		ID: &id,
	}
	response, err = sess.DeleteSubnetWithContext(context, deleteSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 409 {
			log.Printf("[DEBUG] Delete subnet response status code: 409 conflict, provider will try again. %s", err)
			_, err = isWaitForSubnetDeleteRetry(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSubnetWithContext on retry failed: %s", err.Error()), "ibm_is_subnet", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		} else {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	_, err = isWaitForSubnetDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSubnetDeleted failed: %s", err.Error()), "ibm_is_subnet", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func isWaitForSubnetDeleteRetry(vpcClient *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("[DEBUG] Retrying subnet (%s) delete", id)
	stateConf := &resource.StateChangeConf{
		Pending: []string{isSubnetInUse},
		Target:  []string{isSubnetDeleting, isSubnetDeleted, ""},
		Refresh: func() (interface{}, string, error) {
			deleteSubnetOptions := &vpcv1.DeleteSubnetOptions{
				ID: &id,
			}
			log.Printf("[DEBUG] Retrying subnet (%s) delete", id)
			response, err := vpcClient.DeleteSubnet(deleteSubnetOptions)
			if err != nil {
				if response != nil && response.StatusCode == 409 {
					return response, isSubnetInUse, nil
				} else if response != nil && response.StatusCode == 404 {
					return response, isSubnetDeleted, nil
				}
				return response, "", fmt.Errorf("[ERROR] Error deleting subnet: %s\n%s", err, response)
			}
			return response, isSubnetDeleting, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isWaitForSubnetDeleted(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetDeleting},
		Target:     []string{isSubnetDeleted, ""},
		Refresh:    isSubnetDeleteRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetDeleteRefreshFunc(subnetC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] is subnet delete function here")
		getSubnetOptions := &vpcv1.GetSubnetOptions{
			ID: &id,
		}
		subnet, response, err := subnetC.GetSubnet(getSubnetOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return subnet, isSubnetDeleted, nil
			}
			if response != nil && strings.Contains(err.Error(), "please detach all network interfaces from subnet before deleting it") {
				return subnet, isSubnetDeleting, nil
			}
			return subnet, "", fmt.Errorf("[ERROR] The Subnet %s failed to delete: %s\n%s", id, err, response)
		}
		return subnet, isSubnetDeleting, err
	}
}

func resourceIBMISSubnetExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()
	exists, err := subnetExists(d, meta, id)
	return exists, err
}

func subnetExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	getsubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	_, response, err := sess.GetSubnet(getsubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnet failed: %s", err.Error()), "ibm_is_subnet", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
