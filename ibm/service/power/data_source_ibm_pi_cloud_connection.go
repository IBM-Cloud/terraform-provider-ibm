// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	// Arguments
	PICloudConnectionName           = "pi_cloud_connection_name"
	PICloudConnectionClassicEnabled = "pi_cloud_connection_classic_enabled"
	PICloudConnectionGlobalRouting  = "pi_cloud_connection_global_routing"
	PICloudConnectionGreCIDR        = "pi_cloud_connection_gre_cidr"
	PICloudConnectionGREDestination = "pi_cloud_connection_gre_destination_address"
	PICloudConnectionMetered        = "pi_cloud_connection_metered"
	PICloudConnectionNetworks       = "pi_cloud_connection_networks"
	PICloudConnectionSpeed          = "pi_cloud_connection_speed"
	PICloudConnectionVPCEnabled     = "pi_cloud_connection_vpc_enabled"
	PICloudConnectionVPCCrns        = "pi_cloud_connection_vpc_crns"

	// Attributes
	CloudConnections                     = "connections"
	CloudConnectionID                    = "cloud_connection_id"
	CloudConnectionClassicEnabled        = "classic_enabled"
	CloudConnectionGlobalRouting         = "global_routing"
	CloudConnectionGREDestinationAddress = "gre_destination_address"
	CloudConnectionGreSource             = "gre_source_address"
	CloudConnectionIbmIP                 = "ibm_ip_address"
	CloudConnectionMetered               = "metered"
	CloudConnectionNetworks              = "networks"
	CloudConnectionPort                  = "port"
	CloudConnectionSpeed                 = "speed"
	CloudConnectionStatus                = "status"
	CloudConnectionUserIP                = "user_ip_address"
	CloudConnectionVPCCRNs               = "vpc_crns"
	CloudConnectionVPCEnabled            = "vpc_enabled"
	CloudConnectionName                  = "name"
)

func DataSourceIBMPICloudConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPICloudConnectionRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			PICloudConnectionName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Cloud Connection Name to be used",
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
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
	}
}

func dataSourceIBMPICloudConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	cloudConnectionName := d.Get(PICloudConnectionName).(string)
	client := instance.NewIBMPICloudConnectionClient(ctx, sess, cloudInstanceID)

	// Get API does not work with name for Cloud Connection hence using GetAll (max 2)
	// TODO: Uncomment Get call below when avaiable and remove GetAll
	// cloudConnectionD, err := client.GetWithContext(ctx, cloudConnectionName, cloudInstanceID)
	// if err != nil {
	// 	log.Printf("[DEBUG] get cloud connection failed %v", err)
	// 	return diag.Errorf(errors.GetCloudConnectionOperationFailed, cloudConnectionName, err)
	// }
	cloudConnections, err := client.GetAll()
	if err != nil {
		log.Printf("[DEBUG] get cloud connections failed %v", err)
		return diag.FromErr(err)
	}
	var cloudConnection *models.CloudConnection
	if cloudConnections != nil {
		for _, cc := range cloudConnections.CloudConnections {
			if cloudConnectionName == *cc.Name {
				cloudConnection = cc
				break
			}
		}
	}
	if cloudConnection == nil {
		log.Printf("[DEBUG] cloud connection not found")
		return diag.Errorf("failed to perform get cloud connection operation for name %s", cloudConnectionName)
	}

	d.SetId(*cloudConnection.CloudConnectionID)
	d.Set(CloudConnectionGlobalRouting, cloudConnection.GlobalRouting)
	d.Set(CloudConnectionMetered, cloudConnection.Metered)
	d.Set(CloudConnectionIbmIP, cloudConnection.IbmIPAddress)
	d.Set(CloudConnectionUserIP, cloudConnection.UserIPAddress)
	d.Set(CloudConnectionStatus, cloudConnection.LinkStatus)
	d.Set(CloudConnectionPort, cloudConnection.Port)
	d.Set(CloudConnectionSpeed, cloudConnection.Speed)
	d.Set(PICloudInstanceID, cloudInstanceID)
	if cloudConnection.Networks != nil {
		networks := make([]string, len(cloudConnection.Networks))
		for i, ccNetwork := range cloudConnection.Networks {
			if ccNetwork != nil {
				networks[i] = *ccNetwork.NetworkID
			}
		}
		d.Set(CloudConnectionNetworks, networks)
	}
	if cloudConnection.Classic != nil {
		d.Set(CloudConnectionClassicEnabled, cloudConnection.Classic.Enabled)
		if cloudConnection.Classic.Gre != nil {
			d.Set(CloudConnectionGREDestinationAddress, cloudConnection.Classic.Gre.DestIPAddress)
			d.Set(CloudConnectionGreSource, cloudConnection.Classic.Gre.SourceIPAddress)
		}
	}
	if cloudConnection.Vpc != nil {
		d.Set(CloudConnectionVPCEnabled, cloudConnection.Vpc.Enabled)
		if cloudConnection.Vpc.Vpcs != nil && len(cloudConnection.Vpc.Vpcs) > 0 {
			vpcCRNs := make([]string, len(cloudConnection.Vpc.Vpcs))
			for i, vpc := range cloudConnection.Vpc.Vpcs {
				vpcCRNs[i] = *vpc.VpcID
			}
			d.Set(CloudConnectionVPCCRNs, vpcCRNs)
		}
	}

	return nil
}
