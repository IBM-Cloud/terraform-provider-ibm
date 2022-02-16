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
Datasource to get the list of Cloud Connections in a power instance
*/
// Attributes and Arguments defined in data_source_ibm_pi_cloud_connection.go
func DataSourceIBMPICloudConnections() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPICloudConnectionsRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed Attributes
			CloudConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						CloudConnectionID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudConnectionName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudConnectionSpeed: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						CloudConnectionGlobalRouting: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						CloudConnectionMetered: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						CloudConnectionStatus: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudConnectionIbmIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudConnectionUserIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudConnectionPort: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudConnectionNetworks: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Set of Networks attached to this cloud connection",
						},
						CloudConnectionClassicEnabled: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable classic endpoint destination",
						},
						CloudConnectionGREDestinationAddress: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GRE destination IP address",
						},
						CloudConnectionGreSource: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GRE auto-assigned source IP address",
						},
						CloudConnectionVPCEnabled: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable VPC for this cloud connection",
						},
						CloudConnectionVPCCRNs: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Set of VPCs attached to this cloud connection",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPICloudConnectionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	client := st.NewIBMPICloudConnectionClient(ctx, sess, cloudInstanceID)

	cloudConnections, err := client.GetAll()
	if err != nil {
		log.Printf("[DEBUG] get cloud connections failed %v", err)
		return diag.FromErr(err)
	}

	result := make([]map[string]interface{}, 0, len(cloudConnections.CloudConnections))
	for _, cloudConnection := range cloudConnections.CloudConnections {
		cc := map[string]interface{}{
			CloudConnectionID:            *cloudConnection.CloudConnectionID,
			CloudConnectionName:          *cloudConnection.Name,
			CloudConnectionGlobalRouting: *cloudConnection.GlobalRouting,
			CloudConnectionMetered:       *cloudConnection.Metered,
			CloudConnectionIbmIP:         *cloudConnection.IbmIPAddress,
			CloudConnectionUserIP:        *cloudConnection.UserIPAddress,
			CloudConnectionStatus:        *cloudConnection.LinkStatus,
			CloudConnectionPort:          *cloudConnection.Port,
			CloudConnectionSpeed:         *cloudConnection.Speed,
		}

		if cloudConnection.Networks != nil {
			networks := make([]string, len(cloudConnection.Networks))
			for i, ccNetwork := range cloudConnection.Networks {
				if ccNetwork != nil {
					networks[i] = *ccNetwork.NetworkID
				}
			}
			cc[CloudConnectionNetworks] = networks
		}
		if cloudConnection.Classic != nil {
			cc[CloudConnectionClassicEnabled] = cloudConnection.Classic.Enabled
			if cloudConnection.Classic.Gre != nil {
				cc[CloudConnectionGREDestinationAddress] = cloudConnection.Classic.Gre.DestIPAddress
				cc[CloudConnectionGreSource] = cloudConnection.Classic.Gre.SourceIPAddress
			}
		}
		if cloudConnection.Vpc != nil {
			cc[CloudConnectionVPCEnabled] = cloudConnection.Vpc.Enabled
			if cloudConnection.Vpc.Vpcs != nil && len(cloudConnection.Vpc.Vpcs) > 0 {
				vpcCRNs := make([]string, len(cloudConnection.Vpc.Vpcs))
				for i, vpc := range cloudConnection.Vpc.Vpcs {
					vpcCRNs[i] = *vpc.VpcID
				}
				cc[CloudConnectionVPCCRNs] = vpcCRNs
			}
		}

		result = append(result, cc)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(CloudConnections, result)

	return nil
}
