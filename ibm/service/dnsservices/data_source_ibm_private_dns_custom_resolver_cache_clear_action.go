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

const (
	pdnsCRCacheClearActionID  = "action_id"
	pdnsCRCacheClearStatus    = "status"
	pdnsCRCacheClearErrors    = "errors"
	pdnsCRCacheClearErrorCode = "code"
	pdnsCRCacheClearErrorMsg  = "message"
	pdnsCRCacheClearLocID     = "location_id"
	pdnsCRCacheClearCreatedOn = "created_on"
	pdnsCRCacheClearModOn     = "modified_on"
	pdnsCRCacheClearExpOn     = "expires_on"
	pdnsCRCacheClearActions   = "actions"
	pdnsCRCacheClearTotal     = "total_count"
)

func DataSourceIBMPrivateDNSCustomResolverCacheClearAction() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmDnsCrCacheClearActionRead,

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
			pdnsCRCacheClearActionID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of a cache clear action.",
			},
			pdnsCRCacheClearStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the cache clear action (pending, success, failed).",
			},
			pdnsCRCacheClearErrors: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of error details if the cache clear action failed.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						pdnsCRCacheClearErrorCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error code identifier.",
						},
						pdnsCRCacheClearErrorMsg: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error message text.",
						},
						pdnsCRCacheClearLocID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the location that failed to clear its cache.",
						},
					},
				},
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
	}
}

func dataSourceIbmDnsCrCacheClearActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIbmDnsCrCacheClearActionRead Client initialization failed: %s", err.Error()), "ibm_dns_custom_resolver_cache_clear_action", "read")
		return tfErr.GetDiag()
	}

	instanceID := d.Get(pdnsInstanceID).(string)
	resolverID := d.Get(pdnsCRFRResolverID).(string)
	actionID := d.Get(pdnsCRCacheClearActionID).(string)

	opt := sess.NewGetClearCacheActionOptions(instanceID, resolverID, actionID)
	result, resp, err := sess.GetClearCacheActionWithContext(ctx, opt)
	if err != nil || result == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIbmDnsCrCacheClearActionRead GetClearCacheActionWithContext failed with error: %s and response:\n%s", err, resp), "ibm_dns_custom_resolver_cache_clear_action", "read")
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, resolverID, actionID))
	d.Set(pdnsInstanceID, instanceID)
	d.Set(pdnsCRFRResolverID, resolverID)
	d.Set(pdnsCRCacheClearActionID, actionID)

	if result.Status != nil {
		d.Set(pdnsCRCacheClearStatus, *result.Status)
	}
	if result.CreatedOn != nil {
		d.Set(pdnsCRCacheClearCreatedOn, result.CreatedOn.String())
	}
	if result.ModifiedOn != nil {
		d.Set(pdnsCRCacheClearModOn, result.ModifiedOn.String())
	}
	if result.ExpiresOn != nil {
		d.Set(pdnsCRCacheClearExpOn, result.ExpiresOn.String())
	}
	if len(result.Errors) > 0 {
		d.Set(pdnsCRCacheClearErrors, flattenPDNSCacheClearErrors(result.Errors))
	} else {
		d.Set(pdnsCRCacheClearErrors, []interface{}{})
	}

	return nil
}

func flattenPDNSCacheClearErrors(errors []dnssvcsv1.ClearCacheError) []interface{} {
	errorList := make([]interface{}, 0, len(errors))
	for _, e := range errors {
		item := make(map[string]interface{})
		if e.Code != nil {
			item[pdnsCRCacheClearErrorCode] = *e.Code
		}
		if e.Message != nil {
			item[pdnsCRCacheClearErrorMsg] = *e.Message
		}
		if e.LocationID != nil {
			item[pdnsCRCacheClearLocID] = *e.LocationID
		}
		errorList = append(errorList, item)
	}
	return errorList
}
