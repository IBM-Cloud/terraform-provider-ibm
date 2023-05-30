// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	dnsLinkedZoneInstanceID             = "instance_id"
	dnsLinkedZoneName                   = "name"
	dnsLinkedZoneDescription            = "description"
	dnsLinkedZoneLinkedTo               = "linked_to"
	dnsLinkedZoneState                  = "state"
	dnsLinkedZoneLabel                  = "label"
	dnsLinkedZoneApprovalRequiredBefore = "approval_required_before"
	dnsLinkedZoneCreatedOn              = "created_on"
	dnsLinkedZoneModifiedOn             = "modified_on"
	dnsLZOffset                         = "offset"
	dnLSZLimit                          = "limit"
)

func DataSourceIBMDNSLinkedZone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMDNSLinkedZonesRead,

		Schema: map[string]*schema.Schema{
			dnsLinkedZoneInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID of the DNS Services instance.",
			},
			dnsLinkedZoneName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the linked zone.",
			},
			dnsLinkedZoneDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Descriptive text of the linked zone.",
			},
			dnsLinkedZoneLinkedTo: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the custom resolver.",
			},
			dnsLinkedZoneState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the linked zone.",
			},
			dnsLinkedZoneLabel: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The label of the linked zone.",
			},
			dnsLinkedZoneApprovalRequiredBefore: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time when the approval is required before.",
			},
			dnsLinkedZoneCreatedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time when the linked zone is created.",
			},
			dnsLinkedZoneModifiedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The recent time when the linked zone is modified.",
			},
		},
	}
}

func dataSourceIBMDNSLinkedZonesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	instanceID := d.Get(dnsLinkedZoneInstanceID).(string)

	//listLinkedZonesOptions := sess.NewListLinkedZonesOptions(instanceID)
	opt := sess.NewListLinkedZonesOptions(instanceID)
	//availableDNSZones, resp, err := sess.ListLinkedZonesWithContext(context, listLinkedZonesOptions)
	result, resp, err := sess.ListLinkedZonesWithContext(context, opt)
	if err != nil || result == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error listing the Linked Zones %s:%s", err, resp))
	}
	dnsLinkedZones := make([]map[string]interface{}, 0)
	for _, instance := range result.LinkedDnszones {
		dnsLinkedZone := map[string]interface{}{}
		dnsLinkedZone[dnsLinkedZoneInstanceID] = instance.InstanceID
		dnsLinkedZone[dnsLinkedZoneName] = instance.Name
		dnsLinkedZone[dnsLinkedZoneDescription] = instance.Description
		dnsLinkedZone[dnsLinkedZoneLinkedTo] = instance.LinkedTo
		dnsLinkedZone[dnsLinkedZoneState] = instance.State
		dnsLinkedZone[dnsLinkedZoneLabel] = instance.Label
		dnsLinkedZone[dnsLinkedZoneApprovalRequiredBefore] = instance.ApprovalRequiredBefore
		dnsLinkedZone[dnsLinkedZoneCreatedOn] = instance.CreatedOn
		dnsLinkedZone[dnsLinkedZoneModifiedOn] = instance.ModifiedOn
		dnsLinkedZones = append(dnsLinkedZones, dnsLinkedZone)
	}
	d.SetId(dataSourceIBMDNSLinkedZoneID(d))
	//d.Set(dnsLinkedZoneName)
	return nil
}

func dataSourceIBMDNSLinkedZoneID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
