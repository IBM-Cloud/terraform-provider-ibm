// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package distributionlistapi

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
	"github.com/go-openapi/strfmt"
)

func DataSourceIbmDistributionListDestination() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmDistributionListDestinationRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IBM Cloud account ID.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID of the destination.",
			},
			"destination_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the destination.",
			},
		},
	}
}

func dataSourceIbmDistributionListDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListApiClient, err := meta.(conns.ClientSession).DistributionListApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_distribution_list_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDistributionListDestinationOptions := &distributionlistapiv1.GetDistributionListDestinationOptions{}

	getDistributionListDestinationOptions.SetAccountID(d.Get("account_id").(string))
	getDistributionListDestinationOptions.SetDestinationID(core.UUIDPtr(strfmt.UUID(d.Get("destination_id").(string))))

	eventNotificationDestinationIntf, _, err := distributionListApiClient.GetDistributionListDestinationWithContext(context, getDistributionListDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDistributionListDestinationWithContext failed: %s", err.Error()), "(Data) ibm_distribution_list_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	eventNotificationDestination := eventNotificationDestinationIntf.(*distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination)

	d.SetId(fmt.Sprintf("%s/%s", *getDistributionListDestinationOptions.AccountID, *getDistributionListDestinationOptions.DestinationID))

	if err = d.Set("destination_type", eventNotificationDestination.DestinationType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination_type: %s", err), "(Data) ibm_distribution_list_destination", "read", "set-destination_type").GetDiag()
	}

	return nil
}
