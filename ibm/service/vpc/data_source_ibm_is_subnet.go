// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISSubnet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISSubnetRead,

		Schema: map[string]*schema.Schema{

			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isSubnetName, "identifier"},
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_is_subnet", "identifier"),
			},

			isSubnetIpv4CidrBlock: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetAvailableIpv4AddressCount: {
				Type:     schema.TypeInt,
				Computed: true,
			},

			isSubnetTotalIpv4AddressCount: {
				Type:     schema.TypeInt,
				Computed: true,
			},

			isSubnetName: {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ExactlyOneOf: []string{isSubnetName, "identifier"},
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_is_subnet", isSubnetName),
			},

			isSubnetTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			isSubnetAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},

			isSubnetCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			isSubnetNetworkACL: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetPublicGateway: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetVPC: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				RequiredWith: []string{isSubnetName},
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_is_subnet", "identifier"),
			},

			isSubnetVPCName: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetZone: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetResourceGroup: {
				Type:     schema.TypeString,
				Computed: true,
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

			"routing_table": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The routing table for this subnet",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this routing table.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this routing table.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this routing table.",
						},
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The crn for this routing table.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMISSubnetValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "identifier",
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSubnetName,
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})

	ibmISSubnetDataSourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_subnet", Schema: validateSchema}
	return &ibmISSubnetDataSourceValidator
}

func dataSourceIBMISSubnetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := subnetGetByNameOrID(context, d, meta)
	if err != nil {
		return err
	}
	return nil
}

func subnetGetByNameOrID(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_subnet", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	var subnet *vpcv1.Subnet
	if v, ok := d.GetOk("identifier"); ok {
		id := v.(string)
		getSubnetOptions := &vpcv1.GetSubnetOptions{
			ID: &id,
		}
		subnetinfo, _, err := sess.GetSubnetWithContext(context, getSubnetOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetWithContext failed: %s", err.Error()), "(Data) ibm_is_subnet", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		subnet = subnetinfo
	} else if v, ok := d.GetOk(isSubnetName); ok {
		name := v.(string)
		start := ""
		allrecs := []vpcv1.Subnet{}
		getSubnetsListOptions := &vpcv1.ListSubnetsOptions{}
		for {
			if start != "" {
				getSubnetsListOptions.Start = &start
			}
			if vpcIdOk, ok := d.GetOk(isSubnetVPC); ok {
				vpcIDOk := vpcIdOk.(string)
				getSubnetsListOptions.VPCID = &vpcIDOk
			}
			subnetsCollection, _, err := sess.ListSubnetsWithContext(context, getSubnetsListOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSubnetsWithContext failed: %s", err.Error()), "(Data) ibm_is_subnet", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(subnetsCollection.Next)
			allrecs = append(allrecs, subnetsCollection.Subnets...)
			if start == "" {
				break
			}
		}

		for _, subnetInfo := range allrecs {
			if *subnetInfo.Name == name {
				subnet = &subnetInfo
				break
			}
		}
		if subnet == nil {
			err = fmt.Errorf("[ERROR] No subnet found with name (%s)", name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSubnetsWithContext failed: %s", err.Error()), "(Data) ibm_is_subnet", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId(*subnet.ID)

	if subnet.RoutingTable != nil {
		err = d.Set("routing_table", dataSourceSubnetFlattenroutingTable(*subnet.RoutingTable))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting available_ipv4_address_count: %s", err), "(Data) ibm_is_subnet", "read", "set-routing_table").GetDiag()
		}
	}
	if err = d.Set("name", subnet.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_subnet", "read", "set-name").GetDiag()
	}
	if err = d.Set("ipv4_cidr_block", subnet.Ipv4CIDRBlock); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ipv4_cidr_block: %s", err), "(Data) ibm_is_subnet", "read", "set-ipv4_cidr_block").GetDiag()
	}
	if err = d.Set("available_ipv4_address_count", flex.IntValue(subnet.AvailableIpv4AddressCount)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting available_ipv4_address_count: %s", err), "(Data) ibm_is_subnet", "read", "set-available_ipv4_address_count").GetDiag()
	}
	if err = d.Set("total_ipv4_address_count", flex.IntValue(subnet.TotalIpv4AddressCount)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_ipv4_address_count: %s", err), "(Data) ibm_is_subnet", "read", "set-total_ipv4_address_count").GetDiag()
	}
	if subnet.NetworkACL != nil {
		if err = d.Set(isSubnetNetworkACL, *subnet.NetworkACL.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_acl: %s", err), "(Data) ibm_is_subnet", "read", "set-network_acl").GetDiag()
		}
	}
	if subnet.PublicGateway != nil {
		if err = d.Set(isSubnetPublicGateway, *subnet.PublicGateway.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting public_gateway: %s", err), "(Data) ibm_is_subnet", "read", "set-public_gateway").GetDiag()
		}
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	if err = d.Set("status", subnet.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_subnet", "read", "set-status").GetDiag()
	}
	if subnet.Zone != nil {
		if err = d.Set(isSubnetZone, *subnet.Zone.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_subnet", "read", "set-zone").GetDiag()
		}
	}
	if subnet.VPC != nil {
		if err = d.Set(isSubnetVPC, *subnet.VPC.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_subnet", "read", "set-vpc").GetDiag()
		}
		if err = d.Set(isSubnetVPCName, *subnet.VPC.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc_name: %s", err), "(Data) ibm_is_subnet", "read", "set-vpc_name").GetDiag()
		}
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "(Data) ibm_is_subnet", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	tags, err := flex.GetGlobalTagsUsingCRN(meta, *subnet.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"An error occured during reading of subnet (%s) tags : %s", d.Id(), err)
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *subnet.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource subnet (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isSubnetTags, tags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_subnet", "read", "set-tags").GetDiag()
	}
	if err = d.Set(isSubnetAccessTags, accesstags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_subnet", "read", "set-access_tags").GetDiag()
	}
	if err = d.Set("crn", subnet.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_subnet", "read", "set-crn").GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/subnets"); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_controller_url: %s", err), "(Data) ibm_is_subnet", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *subnet.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_name: %s", err), "(Data) ibm_is_subnet", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *subnet.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_crn: %s", err), "(Data) ibm_is_subnet", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set(flex.ResourceStatus, *subnet.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_status: %s", err), "(Data) ibm_is_subnet", "read", "set-resource_status").GetDiag()
	}
	if subnet.ResourceGroup != nil {
		if err = d.Set(isSubnetResourceGroup, *subnet.ResourceGroup.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_subnet", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, *subnet.ResourceGroup.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group_name: %s", err), "(Data) ibm_is_subnet", "read", "set-resource_group_name").GetDiag()
		}
	}
	return nil
}

func dataSourcesubnetRoutingTableDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceSubnetFlattenroutingTable(result vpcv1.RoutingTableReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceSubnetRoutingTableToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceSubnetRoutingTableToMap(routingTableItem vpcv1.RoutingTableReference) (routingTableMap map[string]interface{}) {
	routingTableMap = map[string]interface{}{}

	if routingTableItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourcesubnetRoutingTableDeletedToMap(*routingTableItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		routingTableMap["deleted"] = deletedList
	}
	if routingTableItem.Href != nil {
		routingTableMap["href"] = routingTableItem.Href
	}
	if routingTableItem.ID != nil {
		routingTableMap["id"] = routingTableItem.ID
	}
	if routingTableItem.Name != nil {
		routingTableMap["name"] = routingTableItem.Name
	}
	if routingTableItem.CRN != nil {
		routingTableMap["crn"] = routingTableItem.CRN
	}
	if routingTableItem.ResourceType != nil {
		routingTableMap["resource_type"] = routingTableItem.ResourceType
	}

	return routingTableMap
}
