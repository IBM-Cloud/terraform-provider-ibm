// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	//"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	DnsLinkedZoneInstanceID             = "instance_id"
	DnsLinkedZoneName                   = "name"
	DnsLinkedZoneDescription            = "description"
	DnsLinkedZoneLinkedTo               = "linked_to"
	DnsLinkedZoneState                  = "state"
	DnsLinkedZoneLabel                  = "label"
	DnsLinkedZoneApprovalRequiredBefore = "approval_required_before"
	DnsLinkedZoneCreatedOn              = "created_on"
	DnsLinkedZoneModifiedOn             = "modified_on"

// DnsLinkedZoneOwnerInstanceID        = "owner_instance_id"
)

func ResourceIBMDNSLinkedZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMDNSLinkedZoneCreate,
		ReadContext:   resourceIBMDNSLinkedZoneRead,
		UpdateContext: resourceIBMDNSLinkedZoneUpdate,
		DeleteContext: resourceIBMDNSLinkedZoneDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			DnsLinkedZoneInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of a service instance.",
			},
			DnsLinkedZoneName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of a service instance.",
			},
			DnsLinkedZoneDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Descriptive text of the secondary zone",
			},
			DnsLinkedZoneLinkedTo: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique identifier of the primary zone",
			},
			DnsLinkedZoneState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the secondary zone",
			},
			DnsLinkedZoneLabel: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The label of the secondary zone",
			},
			DnsLinkedZoneApprovalRequiredBefore: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The date when the secondary zone will be deleted",
			},
			DnsLinkedZoneCreatedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secondary Zone Creation date",
			},
			DnsLinkedZoneModifiedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secondary Zone Modification date",
			},
		},
	}
}

func resourceIBMDNSLinkedZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceID := d.Get(DnsLinkedZoneInstanceID).(string)
	description := d.Get(DnsLinkedZoneDescription).(string)

	createLinkedZoneOptions := sess.NewCreateLinkedZoneOptions(instanceID)

	createLinkedZoneOptions.SetDescription(description)
	mk := "dns_linked_zone_" + instanceID
	conns.IbmMutexKV.Lock(mk)
	defer conns.IbmMutexKV.Unlock(mk)

	resource, response, err := sess.CreateLinkedZone(createLinkedZoneOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error creating DNS Linked zone:%s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", instanceID, *resource.ID))
	return resourceIBMDNSLinkedZoneRead(ctx, d, meta)
}

func resourceIBMDNSLinkedZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	idSet := strings.Split(d.Id(), "/")
	if len(idSet) < 2 {
		return diag.FromErr(fmt.Errorf("[ERROR] Incorrect ID %s: Id should be a combination of InstanceID/resolverID/secondaryZoneID", d.Id()))
	}
	instanceID := idSet[0]
	linkedDnsZoneID := idSet[1]
	getLinkedZoneOptions := sess.NewGetLinkedZoneOptions(instanceID, linkedDnsZoneID)
	resource, response, err := sess.GetLinkedZone(getLinkedZoneOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error reading DNS Linked zone:%s\n%s", err, response))
	}

	//transferFrom := []string{}
	//for _, value := range resource.TransferFrom {
	//	values := strings.Split(value, ":")
	//	transferFrom = append(transferFrom, values[0])
	//}
	d.Set(DnsLinkedZoneInstanceID, idSet[0])
	d.Set(DnsLinkedZoneDescription, *resource.Description)
	d.Set(DnsLinkedZoneCreatedOn, resource.CreatedOn)
	d.Set(DnsLinkedZoneModifiedOn, resource.ModifiedOn)

	return nil
}

func resourceIBMDNSLinkedZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}

	idSet := strings.Split(d.Id(), "/")
	if len(idSet) < 2 {
		return diag.FromErr(fmt.Errorf("[ERROR] Incorrect ID %s: Id should be a combination of InstanceID/resolverID/secondaryZoneID", d.Id()))
	}
	instanceID := idSet[0]
	linkedDnsZoneID := idSet[1]

	// Check DNS zone is present
	getLinkedZoneOptions := sess.NewGetLinkedZoneOptions(instanceID, linkedDnsZoneID)
	_, response, err := sess.GetLinkedZone(getLinkedZoneOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error fetching secondary zone:%s\n%s", err, response))
	}

	// Update DNS Linked zone if attributes has any change
	if d.HasChange(DnsLinkedZoneDescription) ||
		d.HasChange(DnsLinkedZoneLabel) {
		updateLinkedZoneOptions := sess.NewUpdateLinkedZoneOptions(instanceID, linkedDnsZoneID)
		//transferFrom := flex.ExpandStringList(d.Get(pdnsSecondaryZoneTransferFrom).([]interface{}))
		description := d.Get(DnsLinkedZoneDescription).(string)
		//enabled := d.Get(pdnsSecZoneEnabled).(bool)
		//updateSecondaryZoneOptions.SetTransferFrom(transferFrom)
		updateLinkedZoneOptions.SetDescription(description)
		//updateSecondaryZoneOptions.SetEnabled(enabled)

		mk := "dns_linked_zone_" + instanceID
		conns.IbmMutexKV.Lock(mk)
		defer conns.IbmMutexKV.Unlock(mk)

		_, response, err := sess.UpdateLinkedZone(updateLinkedZoneOptions)

		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error updating DNS Services zone:%s\n%s", err, response))
		}
	}

	return resourceIBMDNSLinkedZoneRead(ctx, d, meta)
}
func resourceIBMDNSLinkedZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	idSet := strings.Split(d.Id(), "/")
	if len(idSet) < 3 {
		return diag.FromErr(fmt.Errorf("[ERROR] Incorrect ID %s: Id should be a combination of InstanceID/resolverID/secondaryZoneID", d.Id()))
	}
	instanceID := idSet[0]
	linkedDnsZoneID := idSet[1]
	deleteLinkedZoneOptions := sess.NewDeleteLinkedZoneOptions(instanceID, linkedDnsZoneID)

	mk := "linked_dns_zone_" + instanceID
	conns.IbmMutexKV.Lock(mk)
	defer conns.IbmMutexKV.Unlock(mk)
	response, err := sess.DeleteLinkedZone(deleteLinkedZoneOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error reading DNS Services secondary zone:%s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
