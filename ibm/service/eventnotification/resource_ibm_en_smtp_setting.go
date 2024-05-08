// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMEnSMTPSetting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMEnSMTPSettingCreate,
		ReadContext:   resourceIBMEnSMTPSettingRead,
		UpdateContext: resourceIBMEnSMTPSettingUpdate,
		DeleteContext: resourceIBMEnSMTPSettingsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"smtp_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnets": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The SMTP allowed Ips.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceIBMEnSMTPSettingCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	options := &en.UpdateSMTPAllowedIpsOptions{}

	options.SetInstanceID(d.Get("instance_guid").(string))
	options.SetID(d.Get("smtp_config_id").(string))

	subnets := AllowedIPSMap(d.Get("settings").(map[string]interface{}))
	options.SetSubnets(subnets)

	_, response, err := enClient.UpdateSMTPAllowedIpsWithContext(context, options)
	if err != nil {
		return diag.FromErr(fmt.Errorf("UpdateSMTPAllowedIpsWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *options.InstanceID, *options.ID))

	return resourceIBMEnSMTPSettingRead(context, d, meta)
}

func resourceIBMEnSMTPSettingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	options := &en.GetSMTPAllowedIpsOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	_, response, err := enClient.GetSMTPAllowedIpsWithContext(context, options)
	if err != nil {
		return diag.FromErr(fmt.Errorf("GetSMTPAllowedIpsWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("instance_guid", options.InstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting instance_guid: %s", err))
	}

	if err = d.Set("smtp_config_id", options.ID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error smtp_config_id: %s", err))
	}

	return nil
}

func resourceIBMEnSMTPSettingUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	options := &en.UpdateSMTPAllowedIpsOptions{}

	options.SetInstanceID(d.Get("instance_guid").(string))
	options.SetID(d.Get("smtp_config_id").(string))

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	subnets := AllowedIPSMap(d.Get("settings.0").(map[string]interface{}))
	options.SetSubnets(subnets)

	_, response, err := enClient.UpdateSMTPAllowedIpsWithContext(context, options)
	if err != nil {
		return diag.FromErr(fmt.Errorf("UpdateSMTPAllowedIpsWithContext failed %s\n%s", err, response))
	}

	return resourceIBMEnSMTPSettingRead(context, d, meta)
}

func AllowedIPSMap(allowedip map[string]interface{}) []string {

	ips := []string{}
	if allowedip["subnets"] != nil {

		for _, ip := range allowedip["subnets"].([]interface{}) {
			ips = append(ips, ip.(string))
		}

	}

	return ips
}

func resourceIBMEnSMTPSettingsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}
