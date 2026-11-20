// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/networking-go-sdk/dnssvcsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMPrivateDNSCustomResolverCacheClearActions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmDnsCrCacheClearActionsRead,

		Schema: map[string]*schema.Schema{
			pdnsInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of a service instance.",
			},
			pdnsCRFRResolverID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of a custom resolver.",
			},
			pdnsCRCacheClearActions: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of cache clear actions for the custom resolver.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						pdnsCRCacheClearActionID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the cache clear action.",
						},
						pdnsCRCacheClearStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cache clear action.",
						},
						pdnsCRCacheClearCreatedOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the cache clear action was created.",
						},
						pdnsCRCacheClearModOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the cache clear action was modified.",
						},
						pdnsCRCacheClearExpOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the cache clear action expires.",
						},
					},
				},
			},
			pdnsCRCacheClearTotal: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of cache clear actions.",
			},
		},
	}
}

func dataSourceIbmDnsCrCacheClearActionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIbmDnsCrCacheClearActionsRead Client initialization failed: %s", err.Error()), "ibm_dns_custom_resolver_cache_clear_actions", "read")
		return tfErr.GetDiag()
	}

	instanceID := d.Get(pdnsInstanceID).(string)
	resolverID := d.Get(pdnsCRFRResolverID).(string)

	opt := sess.NewListClearCacheActionsOptions(instanceID, resolverID)
	result, resp, err := sess.ListClearCacheActionsWithContext(ctx, opt)
	if err != nil || result == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIbmDnsCrCacheClearActionsRead ListClearCacheActionsWithContext failed with error: %s and response:\n%s", err, resp), "ibm_dns_custom_resolver_cache_clear_actions", "read")
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", instanceID, resolverID))
	d.Set(pdnsInstanceID, instanceID)
	d.Set(pdnsCRFRResolverID, resolverID)

	if result.ClearCacheActions != nil {
		d.Set(pdnsCRCacheClearActions, flattenPDNSCacheClearActions(result.ClearCacheActions))
	} else {
		d.Set(pdnsCRCacheClearActions, []interface{}{})
	}

	if result.Count != nil {
		d.Set(pdnsCRCacheClearTotal, int(*result.Count))
	} else {
		d.Set(pdnsCRCacheClearTotal, len(result.ClearCacheActions))
	}

	return nil
}

func flattenPDNSCacheClearActions(actions []dnssvcsv1.ClearCacheAction) []interface{} {
	actionList := make([]interface{}, 0, len(actions))
	for _, a := range actions {
		item := make(map[string]interface{})
		if a.ID != nil {
			item[pdnsCRCacheClearActionID] = *a.ID
		}
		if a.Status != nil {
			item[pdnsCRCacheClearStatus] = *a.Status
		}
		if a.CreatedOn != nil {
			item[pdnsCRCacheClearCreatedOn] = a.CreatedOn.String()
		}
		if a.ModifiedOn != nil {
			item[pdnsCRCacheClearModOn] = a.ModifiedOn.String()
		}
		if a.ExpiresOn != nil {
			item[pdnsCRCacheClearExpOn] = a.ExpiresOn.String()
		}
		actionList = append(actionList, item)
	}
	return actionList
}
