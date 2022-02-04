// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisWebhooksID     = "webhooks_id"
	cisWebhooksName   = "name"
	cisWebhooksURL    = "url"
	cisWebhooksType   = "type"
	cisWebhooksSecret = "secret"
)

func ResourceIBMCISWebhooks() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISWebhooksCreate,
		Read:     ResourceIBMCISWebhooksRead,
		Update:   ResourceIBMCISWebhooksUpdate,
		Delete:   ResourceIBMCISWebhooksDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisWebhooksID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Webhook ID",
			},
			cisWebhooksName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Webhook Name",
			},
			cisWebhooksURL: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Webhook URL",
			},
			cisWebhooksType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Webhook Type",
			},
			cisWebhooksSecret: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "API key needed to use the webhook",
			},
		},
	}
}
func ResourceIBMCISWebhooksCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisWebhooksSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisWebhooksSession %s", err)
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewCreateAlertWebhookOptions()

	if name, ok := d.GetOk(cisWebhooksName); ok {
		opt.SetName(name.(string))
	}
	if url, ok := d.GetOk(cisWebhooksURL); ok {
		opt.SetURL(url.(string))
	}
	if secret, ok := d.GetOk(cisWebhooksSecret); ok {
		opt.SetSecret((secret.(string)))
	}

	result, resp, err := sess.CreateAlertWebhook(opt)
	if err != nil || result == nil {
		return fmt.Errorf("[ERROR] Error creating Webhooks  %s %s", err, resp)
	}
	log.Printf("Alert Webhooks created successfully : %s", *result.Result.ID)
	d.SetId(flex.ConvertCisToTfTwoVar(*result.Result.ID, crn))
	return ResourceIBMCISWebhooksRead(d, meta)

}
func ResourceIBMCISWebhooksRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisWebhooksSession()
	if err != nil {
		return err
	}
	webhooksID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	sess.Crn = core.StringPtr(crn)
	opt := sess.NewGetWebhookOptions(webhooksID)

	result, resp, err := sess.GetWebhook(opt)
	if err != nil {
		log.Printf("Error reading Alert Webhook detail: %s", resp)
		return err
	}
	if result.Result != nil {
		d.Set(cisID, crn)
		d.Set(cisWebhooksID, result.Result.ID)
		d.Set(cisWebhooksName, result.Result.Name)
		d.Set(cisWebhooksURL, result.Result.URL)
		d.Set(cisWebhooksType, result.Result.Type)
	}
	return nil
}
func ResourceIBMCISWebhooksUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisWebhooksSession()
	if err != nil {
		return err
	}
	webhooksID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	sess.Crn = core.StringPtr(crn)
	opt := sess.NewUpdateAlertWebhookOptions(webhooksID)
	if d.HasChange(cisWebhooksName) ||
		d.HasChange(cisWebhooksURL) ||
		d.HasChange(cisWebhooksType) ||
		d.HasChange(cisWebhooksSecret) {

		if name, ok := d.GetOk(cisWebhooksName); ok {
			opt.SetName(name.(string))
		}
		if url, ok := d.GetOk(cisWebhooksURL); ok {
			opt.SetURL(url.(string))
		}
		if secret, ok := d.GetOk(cisWebhooksSecret); ok {
			opt.SetSecret((secret.(string)))
		}

		result, resp, err := sess.UpdateAlertWebhook(opt)
		if err != nil {
			log.Printf("Error updating Alert webhook detail: %s", resp)
			return err
		}
		log.Printf("Alert Webhook update succesful : %s", *result.Result.ID)
	}
	return ResourceIBMCISWebhooksRead(d, meta)
}
func ResourceIBMCISWebhooksDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisWebhooksSession()
	if err != nil {
		return err
	}
	webhooksID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	sess.Crn = core.StringPtr(crn)
	opt := sess.NewDeleteWebhookOptions(webhooksID)
	result, resp, err := sess.DeleteWebhook(opt)
	if err != nil {
		log.Printf("Error deleting Alert Webhooks detail: %s", resp)
		return err
	}
	log.Printf("Webhook ID : %s", *result.Result.ID)
	return nil

}
