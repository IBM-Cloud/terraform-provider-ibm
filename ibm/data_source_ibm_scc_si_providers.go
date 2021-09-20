// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/scc-go-sdk/findingsv1"
)

func dataSourceIBMSccSiProviders() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSccSiProvidersRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"providers": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The providers requested.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the provider in the form '{account_id}/providers/{provider_id}'.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the provider.",
						},
					},
				},
			},
			"limit": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of elements returned in the current instance. The default is 200.",
			},
			"skip": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The offset is the index of the item from which you want to start returning data from. The default is 0.",
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of providers available.",
			},
		},
	}
}

func dataSourceIBMSccSiProvidersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	findingsClient, err := meta.(ClientSession).FindingsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	accountID := d.Get("account_id").(string)
	log.Println(fmt.Sprintf("[DEBUG] using specified AccountID %s", accountID))
	if accountID == "" {
		accountID = userDetails.userAccount
		log.Println(fmt.Sprintf("[DEBUG] AccountID not spedified, using %s", accountID))
	}
	findingsClient.AccountID = &accountID

	listProvidersOptions := &findingsv1.ListProvidersOptions{}

	if skip, ok := d.GetOk("skip"); ok {
		listProvidersOptions.SetSkip(int64(skip.(int)))
	}
	if limit, ok := d.GetOk("limit"); ok {
		listProvidersOptions.SetLimit(int64(limit.(int)))
	}

	apiListProvidersResponse, response, err := findingsClient.ListProvidersWithContext(context, listProvidersOptions)
	if err != nil {
		log.Printf("[DEBUG] ListProvidersWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListProvidersWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIBMSccSiProvidersID(d))

	if apiListProvidersResponse.Providers != nil {
		err = d.Set("providers", dataSourceAPIListProvidersResponseFlattenProviders(apiListProvidersResponse.Providers))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting providers %s", err))
		}
	}
	if err = d.Set("total_count", intValue(apiListProvidersResponse.TotalCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting total_count: %s", err))
	}

	return nil
}

// dataSourceIBMSccSiProviderID returns a reasonable ID for the list.
func dataSourceIBMSccSiProvidersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceAPIListProvidersResponseFlattenProviders(result []findingsv1.APIProvider) (providers []map[string]interface{}) {
	for _, providersItem := range result {
		providers = append(providers, dataSourceAPIListProvidersResponseProvidersToMap(providersItem))
	}

	return providers
}

func dataSourceAPIListProvidersResponseProvidersToMap(providersItem findingsv1.APIProvider) (providersMap map[string]interface{}) {
	providersMap = map[string]interface{}{}

	if providersItem.Name != nil {
		providersMap["name"] = providersItem.Name
	}
	if providersItem.ID != nil {
		providersMap["id"] = providersItem.ID
	}

	return providersMap
}
