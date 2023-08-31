// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func DataSourceIbmSccProviderTypeInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSccProviderTypeInstanceRead,

		Schema: map[string]*schema.Schema{
			"provider_type_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The provider type ID.",
			},
			"provider_type_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The provider type instance ID.",
			},
			"x_correlation_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The supplied or generated value of this header is logged for a request and repeated in a response header for the corresponding response. The same value is used for downstream requests and retries of those requests. If a value of this headers is not supplied in a request, the service generates a random (version 4) UUID.",
			},
			"x_request_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The supplied or generated value of this header is logged for a request and repeated in a response header  for the corresponding response.  The same value is not used for downstream requests and retries of those requests.  If a value of this headers is not supplied in a request, the service generates a random (version 4) UUID.",
			},
			"provider_type_instance_item_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the provider type instance.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the provider type.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the provider type instance.",
			},
			"attributes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The attributes for connecting to the provider type instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time at which resource was created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time at which resource was updated.",
			},
		},
	}
}

func dataSourceIbmSccProviderTypeInstanceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	securityAndComplianceCenterApIsClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	getProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions{}

	getProviderTypeInstanceOptions.SetProviderTypeID(d.Get("provider_type_id").(string))
	getProviderTypeInstanceOptions.SetProviderTypeInstanceID(d.Get("provider_type_instance_id").(string))
	if _, ok := d.GetOk("x_correlation_id"); ok {
		getProviderTypeInstanceOptions.SetXCorrelationID(d.Get("x_correlation_id").(string))
	}
	if _, ok := d.GetOk("x_request_id"); ok {
		getProviderTypeInstanceOptions.SetXRequestID(d.Get("x_request_id").(string))
	}

	providerTypeInstanceItem, response, err := securityAndComplianceCenterApIsClient.GetProviderTypeInstanceWithContext(context, getProviderTypeInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProviderTypeInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProviderTypeInstanceWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *getProviderTypeInstanceOptions.ProviderTypeID, *getProviderTypeInstanceOptions.ProviderTypeInstanceID))

	if err = d.Set("provider_type_instance_item_id", providerTypeInstanceItem.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting provider_type_instance_item_id: %s", err))
	}

	if err = d.Set("type", providerTypeInstanceItem.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}

	if err = d.Set("name", providerTypeInstanceItem.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	attributes := []map[string]interface{}{}
	if providerTypeInstanceItem.Attributes != nil {
		attributes = append(attributes, providerTypeInstanceItem.Attributes)
	}
	if err = d.Set("attributes", attributes); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting attributes %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(providerTypeInstanceItem.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(providerTypeInstanceItem.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	return nil
}
