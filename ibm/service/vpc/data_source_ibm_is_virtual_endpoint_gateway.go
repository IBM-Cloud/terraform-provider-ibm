// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISEndpointGateway() *schema.Resource {
	return &schema.Resource{
		Read:     dataSourceIBMISEndpointGatewayRead,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isVirtualEndpointGatewayName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Endpoint gateway name",
			},
			isVirtualEndpointGatewayResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway resource type",
			},
			isVirtualEndpointGatewayCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this Endpoint gateway",
			},
			isVirtualEndpointGatewayResourceGroupID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group id",
			},
			isVirtualEndpointGatewayCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway created date and time",
			},
			isVirtualEndpointGatewayHealthState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway health state",
			},
			isVirtualEndpointGatewayLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway lifecycle state",
			},
			isVirtualEndpointGatewaySecurityGroups: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Endpoint gateway securitygroups list",
			},
			isVirtualEndpointGatewayIPs: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Endpoint gateway IPs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVirtualEndpointGatewayIPsID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPs id",
						},
						isVirtualEndpointGatewayIPsName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPs name",
						},
						isVirtualEndpointGatewayIPsResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint gateway IP resource type",
						},
						isVirtualEndpointGatewayIPsAddress: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint gateway IP Address",
						},
					},
				},
			},
			isVirtualEndpointGatewayTarget: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Endpoint gateway target",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVirtualEndpointGatewayTargetName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The target name",
						},
						isVirtualEndpointGatewayTargetResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The target resource type",
						},
					},
				},
			},
			isVirtualEndpointGatewayVpcID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The VPC id",
			},
			isVirtualEndpointGatewayTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "List of tags for VPE",
			},
			isVirtualEndpointGatewayAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "List of access management tags",
			},
		},
	}
}

func dataSourceIBMISEndpointGatewayRead(
	d *schema.ResourceData, meta interface{}) error {
	var found bool
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	name := d.Get(isVirtualEndpointGatewayName).(string)

	start := ""
	allrecs := []vpcv1.EndpointGateway{}
	for {
		options := sess.NewListEndpointGatewaysOptions()
		if start != "" {
			options.Start = &start
		}
		result, response, err := sess.ListEndpointGateways(options)
		if err != nil {
			return fmt.Errorf("[ERROR] Error fetching endpoint gateways %s\n%s", err, response)
		}
		start = flex.GetNext(result.Next)
		allrecs = append(allrecs, result.EndpointGateways...)
		if start == "" {
			break
		}
	}
	for _, endpointGateway := range allrecs {
		if *endpointGateway.Name == name {
			d.SetId(*endpointGateway.ID)
			d.Set(isVirtualEndpointGatewayName, endpointGateway.Name)
			d.Set(isVirtualEndpointGatewayCRN, endpointGateway.CRN)
			d.Set(isVirtualEndpointGatewayHealthState, endpointGateway.HealthState)
			d.Set(isVirtualEndpointGatewayCreatedAt, endpointGateway.CreatedAt.String())
			d.Set(isVirtualEndpointGatewayLifecycleState, endpointGateway.LifecycleState)
			d.Set(isVirtualEndpointGatewayResourceType, endpointGateway.ResourceType)
			d.Set(isVirtualEndpointGatewayIPs, flattenIPs(endpointGateway.Ips))
			d.Set(isVirtualEndpointGatewayResourceGroupID, endpointGateway.ResourceGroup.ID)
			d.Set(isVirtualEndpointGatewayTarget, flattenEndpointGatewayTarget(
				endpointGateway.Target.(*vpcv1.EndpointGatewayTarget)))
			if endpointGateway.SecurityGroups != nil {
				d.Set(isVirtualEndpointGatewaySecurityGroups, flattenDataSourceSecurityGroups(endpointGateway.SecurityGroups))
			}
			d.Set(isVirtualEndpointGatewayVpcID, endpointGateway.VPC.ID)
			tags, err := flex.GetGlobalTagsUsingCRN(meta, *endpointGateway.CRN, "", isUserTagType)
			if err != nil {
				log.Printf(
					"Error on get of VPE (%s) tags: %s", d.Id(), err)
			}
			d.Set(isVirtualEndpointGatewayTags, tags)

			accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *endpointGateway.CRN, "", isAccessTagType)
			if err != nil {
				log.Printf(
					"Error on get of VPE (%s) access tags: %s", d.Id(), err)
			}
			d.Set(isVirtualEndpointGatewayAccessTags, accesstags)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("[ERROR] No Virtual Endpoints Gateway found with given name %s", name)
	}
	return nil
}
