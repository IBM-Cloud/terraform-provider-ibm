// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
)

func DataSourceIBMDatabaseRemotes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMDatabaseRemotesRead,

		Schema: map[string]*schema.Schema{
			"deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Deployment ID.",
			},
			"leader": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Leader ID, if applicable.",
			},
			"replicas": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Replica IDs, if applicable.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIBMDatabaseRemotesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return diag.FromErr(err)
	}

	listRemotesOptions := &clouddatabasesv5.ListRemotesOptions{}

	listRemotesOptions.SetID(d.Get("deployment_id").(string))

	remotes, response, err := cloudDatabasesClient.ListRemotesWithContext(context, listRemotesOptions)
	if err != nil {
		log.Printf("[DEBUG] ListRemotesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListRemotesWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIBMDatabaseRemotesID(d))

	remotesLeader := remotes.Remotes.Leader
	remotesReplicas := remotes.Remotes.Replicas

	if err = d.Set("leader", remotesLeader); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting leader: %s", err))
	}

	if err = d.Set("replicas", remotesReplicas); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting replicas: %s", err))
	}

	return nil
}

// dataSourceIBMDatabaseRemotesID returns a reasonable ID for the list.
func dataSourceIBMDatabaseRemotesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
