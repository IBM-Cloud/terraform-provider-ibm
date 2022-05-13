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

func DataSourceIBMPICloudConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPICloudConnectionRead,
		Schema: map[string]*schema.Schema{
			Arg_CloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_CloudConnectionName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Cloud Connection Name to be used",
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
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
	}
}

func dataSourceIBMPICloudConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	cloudConnectionName := d.Get(Arg_CloudConnectionName).(string)
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

	cloudConnection, err = client.Get(*cloudConnection.CloudConnectionID)
	if err != nil {
		log.Printf("[DEBUG] get cloud connection failed %v", err)
		return diag.FromErr(err)
	}

	d.SetId(*cloudConnection.CloudConnectionID)
	d.Set(Arg_CloudConnectionName, cloudConnection.Name)
	d.Set(Attr_CloudConnectionRouting, cloudConnection.GlobalRouting)
	d.Set(Attr_CloudConnectionMetered, cloudConnection.Metered)
	d.Set(Attr_CloudConnectionIbmIP, cloudConnection.IbmIPAddress)
	d.Set(Attr_CloudConnectionUserIP, cloudConnection.UserIPAddress)
	d.Set(Attr_CloudConnectionStatus, cloudConnection.LinkStatus)
	d.Set(Attr_CloudConnectionPort, cloudConnection.Port)
	d.Set(Attr_CloudConnectionSpeed, cloudConnection.Speed)
	d.Set(Arg_CloudInstanceID, cloudInstanceID)
	if cloudConnection.Networks != nil {
		networks := make([]string, len(cloudConnection.Networks))
		for i, ccNetwork := range cloudConnection.Networks {
			if ccNetwork != nil {
				networks[i] = *ccNetwork.NetworkID
			}
		}
		d.Set(Attr_CloudConnectionNetworks, networks)
	}
	if cloudConnection.Classic != nil {
		d.Set(Attr_CloudConnectionClassic, cloudConnection.Classic.Enabled)
		if cloudConnection.Classic.Gre != nil {
			d.Set(AttrCloudConnectionGreDestAddr, cloudConnection.Classic.Gre.DestIPAddress)
			d.Set(Attr_CloudConnectionSourceGreAddr, cloudConnection.Classic.Gre.SourceIPAddress)
		}
	}
	if cloudConnection.Vpc != nil {
		d.Set(Attr_CloudConnectionVPC, cloudConnection.Vpc.Enabled)
		if cloudConnection.Vpc.Vpcs != nil && len(cloudConnection.Vpc.Vpcs) > 0 {
			vpcCRNs := make([]string, len(cloudConnection.Vpc.Vpcs))
			for i, vpc := range cloudConnection.Vpc.Vpcs {
				vpcCRNs[i] = *vpc.VpcID
			}
			d.Set(Attr_CloudConnectionVPCCrns, vpcCRNs)
		}
	}

	return nil
}
