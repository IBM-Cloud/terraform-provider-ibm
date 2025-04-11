// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsIpsecPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsIpsecPolicyRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "ipsec_policy"},
				Description:  "The IPsec policy name.",
			},
			"ipsec_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "ipsec_policy"},
				Description:  "The IPsec policy identifier.",
			},
			"authentication_algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication algorithm.",
			},
			"connections": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The VPN gateway connections that use this IPsec policy.",
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
							Description: "The VPN connection's canonical URL.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPN gateway connection.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this VPN connection.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this IPsec policy was created.",
			},
			"encapsulation_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encapsulation mode used. Only `tunnel` is supported.",
			},
			"encryption_algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption algorithm.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IPsec policy's canonical URL.",
			},
			"key_lifetime": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The key lifetime in seconds.",
			},

			"pfs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Perfect Forward Secrecy.",
			},
			"resource_group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this IPsec policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this resource group.",
						},
					},
				},
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"transform_protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The transform protocol used. Only `esp` is supported.",
			},
		},
	}
}

func dataSourceIBMIsIpsecPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_ipsec_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	name := d.Get("name").(string)
	identifier := d.Get("ipsec_policy").(string)
	var iPsecPolicy *vpcv1.IPsecPolicy
	if name != "" {
		start := ""
		allrecs := []vpcv1.IPsecPolicy{}
		for {
			listIPSecPoliciesyOptions := &vpcv1.ListIpsecPoliciesOptions{}
			if start != "" {
				listIPSecPoliciesyOptions.Start = &start
			}
			ipSecPolicy, _, err := vpcClient.ListIpsecPoliciesWithContext(context, listIPSecPoliciesyOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListIpsecPoliciesWithContext failed: %s", err.Error()), "(Data) ibm_is_ipsec_policy", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(ipSecPolicy.Next)
			allrecs = append(allrecs, ipSecPolicy.IpsecPolicies...)
			if start == "" {
				break
			}
		}
		ipsec_policy_found := false
		for _, ipSecPolicyItem := range allrecs {
			if *ipSecPolicyItem.Name == name {
				iPsecPolicy = &ipSecPolicyItem
				ipsec_policy_found = true
				break
			}
		}

		if !ipsec_policy_found {
			err = fmt.Errorf("No ipsec policy found with given name %s", name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Not found failed: %s", err.Error()), "(Data) ibm_is_ipsec_policy", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	} else {
		getIPSecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{}

		getIPSecPolicyOptions.SetID(identifier)

		ipsecPolicy1, _, err := vpcClient.GetIpsecPolicyWithContext(context, getIPSecPolicyOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIpsecPolicyWithContext failed: %s", err.Error()), "(Data) ibm_is_ipsec_policy", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		iPsecPolicy = ipsecPolicy1
	}

	d.SetId(*iPsecPolicy.ID)
	if err = d.Set("authentication_algorithm", iPsecPolicy.AuthenticationAlgorithm); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting authentication_algorithm: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-authentication_algorithm").GetDiag()
	}

	if iPsecPolicy.Connections != nil {
		err = d.Set("connections", dataSourceIPsecPolicyFlattenConnections(iPsecPolicy.Connections))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_ipsec_policy", "read", "connections-to-map").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(iPsecPolicy.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("encapsulation_mode", iPsecPolicy.EncapsulationMode); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encapsulation_mode: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-encapsulation_mode").GetDiag()
	}
	if err = d.Set("encryption_algorithm", iPsecPolicy.EncryptionAlgorithm); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption_algorithm: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-encryption_algorithm").GetDiag()
	}
	if err = d.Set("href", iPsecPolicy.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-href").GetDiag()
	}
	if err = d.Set("key_lifetime", flex.IntValue(iPsecPolicy.KeyLifetime)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting key_lifetime: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-key_lifetime").GetDiag()
	}
	if err = d.Set("name", iPsecPolicy.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-name").GetDiag()
	}
	if err = d.Set("pfs", iPsecPolicy.Pfs); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting pfs: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-pfs").GetDiag()
	}

	if iPsecPolicy.ResourceGroup != nil {
		err = d.Set("resource_group", dataSourceIPsecPolicyFlattenResourceGroup(*iPsecPolicy.ResourceGroup))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_ipsec_policy", "read", "resource_group-to-map").GetDiag()
		}
	}
	if err = d.Set("resource_type", iPsecPolicy.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("transform_protocol", iPsecPolicy.TransformProtocol); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting transform_protocol: %s", err), "(Data) ibm_is_ipsec_policy", "read", "set-transform_protocol").GetDiag()
	}
	return nil
}

func dataSourceIPsecPolicyFlattenConnections(result []vpcv1.VPNGatewayConnectionReference) (connections []map[string]interface{}) {
	for _, connectionsItem := range result {
		connections = append(connections, dataSourceIPsecPolicyConnectionsToMap(connectionsItem))
	}

	return connections
}

func dataSourceIPsecPolicyConnectionsToMap(connectionsItem vpcv1.VPNGatewayConnectionReference) (connectionsMap map[string]interface{}) {
	connectionsMap = map[string]interface{}{}

	if connectionsItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceIPsecPolicyConnectionsDeletedToMap(*connectionsItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		connectionsMap["deleted"] = deletedList
	}
	if connectionsItem.Href != nil {
		connectionsMap["href"] = connectionsItem.Href
	}
	if connectionsItem.ID != nil {
		connectionsMap["id"] = connectionsItem.ID
	}
	if connectionsItem.Name != nil {
		connectionsMap["name"] = connectionsItem.Name
	}
	if connectionsItem.ResourceType != nil {
		connectionsMap["resource_type"] = connectionsItem.ResourceType
	}

	return connectionsMap
}

func dataSourceIPsecPolicyConnectionsDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceIPsecPolicyFlattenResourceGroup(result vpcv1.ResourceGroupReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceIPsecPolicyResourceGroupToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceIPsecPolicyResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap["href"] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap["id"] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap["name"] = resourceGroupItem.Name
	}

	return resourceGroupMap
}
