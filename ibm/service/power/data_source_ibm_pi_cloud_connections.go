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

const PICloudConnections = "connections"

func DataSourceIBMPICloudConnections() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPICloudConnectionsRead,
		Schema: map[string]*schema.Schema{
			Arg_CloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed Attributes
			PICloudConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_CloudConnectionID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_CloudConnectionName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_CloudConnectionSpeed: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						Attr_CloudConnectionRouting: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						Attr_CloudConnectionMetered: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						Attr_CloudConnectionStatus: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_CloudConnectionIbmIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_CloudConnectionUserIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_CloudConnectionPort: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_CloudConnectionNetworks: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Set of Networks attached to this cloud connection",
						},
						Attr_CloudConnectionClassic: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable classic endpoint destination",
						},
						AttrCloudConnectionGreDestAddr: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GRE destination IP address",
						},
						Attr_CloudConnectionSourceGreAddr: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GRE auto-assigned source IP address",
						},
						Attr_CloudConnectionVPC: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable VPC for this cloud connection",
						},
						Attr_CloudConnectionVPCCrns: {
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

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	client := st.NewIBMPICloudConnectionClient(ctx, sess, cloudInstanceID)

	cloudConnections, err := client.GetAll()
	if err != nil {
		log.Printf("[DEBUG] get cloud connections failed %v", err)
		return diag.FromErr(err)
	}

	result := make([]map[string]interface{}, 0, len(cloudConnections.CloudConnections))
	for _, cloudConnection := range cloudConnections.CloudConnections {
		cc := map[string]interface{}{
			Attr_CloudConnectionID:      *cloudConnection.CloudConnectionID,
			Attr_CloudConnectionName:    *cloudConnection.Name,
			Attr_CloudConnectionRouting: *cloudConnection.GlobalRouting,
			Attr_CloudConnectionMetered: *cloudConnection.Metered,
			Attr_CloudConnectionIbmIP:   *cloudConnection.IbmIPAddress,
			Attr_CloudConnectionUserIP:  *cloudConnection.UserIPAddress,
			Attr_CloudConnectionStatus:  *cloudConnection.LinkStatus,
			Attr_CloudConnectionPort:    *cloudConnection.Port,
			Attr_CloudConnectionSpeed:   *cloudConnection.Speed,
		}

		if cloudConnection.Networks != nil {
			networks := make([]string, len(cloudConnection.Networks))
			for i, ccNetwork := range cloudConnection.Networks {
				if ccNetwork != nil {
					networks[i] = *ccNetwork.NetworkID
				}
			}
			cc[Attr_CloudConnectionNetworks] = networks
		}
		if cloudConnection.Classic != nil {
			cc[Attr_CloudConnectionClassic] = cloudConnection.Classic.Enabled
			if cloudConnection.Classic.Gre != nil {
				cc[AttrCloudConnectionGreDestAddr] = cloudConnection.Classic.Gre.DestIPAddress
				cc[Attr_CloudConnectionSourceGreAddr] = cloudConnection.Classic.Gre.SourceIPAddress
			}
		}
		if cloudConnection.Vpc != nil {
			cc[Attr_CloudConnectionVPC] = cloudConnection.Vpc.Enabled
			if cloudConnection.Vpc.Vpcs != nil && len(cloudConnection.Vpc.Vpcs) > 0 {
				vpcCRNs := make([]string, len(cloudConnection.Vpc.Vpcs))
				for i, vpc := range cloudConnection.Vpc.Vpcs {
					vpcCRNs[i] = *vpc.VpcID
				}
				cc[Attr_CloudConnectionVPCCrns] = vpcCRNs
			}
		}

		result = append(result, cc)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(PICloudConnections, result)

	return nil
}
