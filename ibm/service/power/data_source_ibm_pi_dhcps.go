// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
Datasource to get the list of dhcp servers in a power instance
*/
// Attributes and Arguments defined in data_source_ibm_pi_dhcp.go
func DataSourceIBMPIDhcps() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMDhcpsServersRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed Attributes
			DhcpServers: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DhcpID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the DHCP Server",
						},
						DhcpStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the DHCP Server",
						},
						DhcpNetwork: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The DHCP Server private network",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMDhcpsServersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	client := st.NewIBMPIDhcpClient(ctx, sess, cloudInstanceID)
	dhcpServers, err := client.GetAll()
	if err != nil {
		log.Printf("[DEBUG] get all DHCP failed %v", err)
		return diag.FromErr(err)
	}

	result := make([]map[string]interface{}, 0, len(dhcpServers))
	for _, dhcpServer := range dhcpServers {
		server := map[string]interface{}{
			DhcpID:     *dhcpServer.ID,
			DhcpStatus: *dhcpServer.Status,
		}

		dhcpNetwork := dhcpServer.Network
		if dhcpNetwork != nil {
			server[DhcpNetwork] = *dhcpNetwork.ID
		}
		result = append(result, server)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(DhcpServers, result)

	return nil
}
