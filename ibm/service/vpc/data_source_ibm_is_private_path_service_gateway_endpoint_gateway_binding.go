// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMIsPrivatePathServiceGatewayEndpointGatewayBinding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingRead,

		Schema: map[string]*schema.Schema{
			"private_path_service_gateway": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The private path service gateway identifier.",
			},
			"endpoint_gateway_binding": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The endpoint gateway binding identifier.",
			},
			"account": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The account that created the endpoint gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the endpoint gateway binding was created.",
			},
			"expiration_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expiration date and time for the endpoint gateway binding.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this endpoint gateway binding.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the endpoint gateway binding.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the endpoint gateway binding- `denied`: endpoint gateway binding was denied- `expired`: endpoint gateway binding has expired- `pending`: endpoint gateway binding is awaiting review- `permitted`: endpoint gateway binding was permittedThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the endpoint gateway binding was updated.",
			},
		},
	}
}

func dataSourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPrivatePathServiceGatewayEndpointGatewayBindingOptions := &vpcv1.GetPrivatePathServiceGatewayEndpointGatewayBindingOptions{}

	getPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetPrivatePathServiceGatewayID(d.Get("private_path_service_gateway").(string))
	getPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetID(d.Get("endpoint_gateway_binding").(string))

	privatePathServiceGatewayEndpointGatewayBinding, response, err := vpcClient.GetPrivatePathServiceGatewayEndpointGatewayBindingWithContext(context, getPrivatePathServiceGatewayEndpointGatewayBindingOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPrivatePathServiceGatewayEndpointGatewayBindingWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s//%s", *getPrivatePathServiceGatewayEndpointGatewayBindingOptions.PrivatePathServiceGatewayID, *privatePathServiceGatewayEndpointGatewayBinding.ID))

	account := []map[string]interface{}{}
	if privatePathServiceGatewayEndpointGatewayBinding.Account != nil {
		modelMap, err := dataSourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingAccountReferenceToMap(privatePathServiceGatewayEndpointGatewayBinding.Account)
		if err != nil {
			return diag.FromErr(err)
		}
		account = append(account, modelMap)
	}
	if err = d.Set("account", account); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account: %s", err), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "set-account").GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(privatePathServiceGatewayEndpointGatewayBinding.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("expiration_at", flex.DateTimeToString(privatePathServiceGatewayEndpointGatewayBinding.ExpirationAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting expiration_at: %s", err), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "set-expiration_at").GetDiag()
	}

	if err = d.Set("href", privatePathServiceGatewayEndpointGatewayBinding.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "set-href").GetDiag()
	}

	if err = d.Set("lifecycle_state", privatePathServiceGatewayEndpointGatewayBinding.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "set-lifecycle_state").GetDiag()
	}

	if err = d.Set("resource_type", privatePathServiceGatewayEndpointGatewayBinding.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "set-resource_type").GetDiag()
	}

	if err = d.Set("status", privatePathServiceGatewayEndpointGatewayBinding.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_private_path_service_gateway_endpoint_gateway_binding", "read", "set-status").GetDiag()
	}

	// if privatePathServiceGatewayEndpointGatewayBinding.UpdatedAt != nil {
	// 	if err = d.Set("updated_at", flex.DateTimeToString(privatePathServiceGatewayEndpointGatewayBinding.UpdatedAt)); err != nil {
	// 		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	// 	}
	// }

	return nil
}

func dataSourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingAccountReferenceToMap(model *vpcv1.AccountReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}
