// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
)

func DataSourceIbmSmEnRegistration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmEnRegistrationRead,

		Schema: map[string]*schema.Schema{
			"event_notifications_instance_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A CRN that uniquely identifies an IBM Cloud resource.",
			},
		},
	}
}

func dataSourceIbmSmEnRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	getNotificationsRegistrationOptions := &secretsmanagerv2.GetNotificationsRegistrationOptions{}

	notificationsRegistration, response, err := secretsManagerClient.GetNotificationsRegistrationWithContext(context, getNotificationsRegistrationOptions)
	if err != nil {
		log.Printf("[DEBUG] GetNotificationsRegistrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetNotificationsRegistrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmSmEnRegistrationID(d))

	if err = d.Set("event_notifications_instance_crn", notificationsRegistration.EventNotificationsInstanceCrn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting event_notifications_instance_crn: %s", err))
	}

	return nil
}

// dataSourceIbmSmEnRegistrationID returns a reasonable ID for the list.
func dataSourceIbmSmEnRegistrationID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
