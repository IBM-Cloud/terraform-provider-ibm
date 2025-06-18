// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"context"
	"log"

	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMDLGatewayMacsecConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMDLGatewayMacsecConfigRead,
		Schema: map[string]*schema.Schema{
			dlGatewayId: {
				Type:        schema.TypeString,
				Description: "Gateway ID",
				Required:    true,
			},
			dlActive: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicate whether MACsec protection should be active (true) or inactive (false) for this MACsec enabled gateway",
			},
			dlCipherSuite: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cipher suite used in generating the security association key (SAK).",
			},
			dlConfidentialityOffset: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The confidentiality offset determines the number of octets in an Ethernet frame that are not encrypted.",
			},
			dlCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the resource was created",
			},
			dlKeyServerPriority: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Key Server Priority",
			},
			dlSecurityPolicy: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Determines how packets without MACsec headers are handled.",
			},
			dlMacSecConfigStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of MACsec on the device for this gateway",
			},
			dlUpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the resource was last updated",
			},
			dlWindowSize: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The window size determines the number of frames in a window for replay protection.",
			},
			dlGatewaySakRekey: {
				Type:        schema.TypeSet,
				Description: "Determines how SAK rekeying occurs.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewaySakRekeyInterval: {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "SAK ReKey Interval",
						},
						dlGatewaySakRekeyMode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SAK ReKey Mode",
						},
					},
				},
			},
			dlGatewayMacsecSatusReasons: {
				Type:        schema.TypeList,
				Description: "A reason for the current status.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewaySakRekeyTimerMode: {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "SAK rekey mode based on length of time since last rekey.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									dlGatewayMacsecSatusReasonCode: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Code",
									},
									dlGatewayMacsecSatusReasonMessage: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Message",
									},
									dlGatewayMacsecSatusReasonMoreInfo: {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "More Info",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMDLGatewayMacsecConfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	dlGatewayID := d.Get(dlGatewayId).(string)

	// Get MacSec gateway
	// Construct an instance of the GetGatewayMacsecOptions model
	getGatewayMacsecOptionsModel := new(directlinkv1.GetGatewayMacsecOptions)
	getGatewayMacsecOptionsModel.ID = &dlGatewayID
	getGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	result, response, err := directLink.GetGatewayMacsec(getGatewayMacsecOptionsModel)
	if err != nil {
		log.Println("[WARN] Error Get DL Gateway Macsec", response, err)
		return diag.FromErr(err)
	}

	if result.Active != nil {
		d.Set(dlActive, *result.Active)
	}
	if result.CipherSuite != nil {
		d.Set(dlCipherSuite, *result.CipherSuite)
	}
	if result.ConfidentialityOffset != nil {
		d.Set(dlConfidentialityOffset, *result.ConfidentialityOffset)
	}
	if result.KeyServerPriority != nil {
		d.Set(dlKeyServerPriority, *result.KeyServerPriority)
	}
	if result.SecurityPolicy != nil {
		d.Set(dlSecurityPolicy, *result.SecurityPolicy)
	}
	if result.Status != nil {
		d.Set(dlMacSecConfigStatus, *result.Status)
	}
	if result.WindowSize != nil {
		d.Set(dlWindowSize, *result.WindowSize)
	}
	if result.CreatedAt != nil {
		d.Set(dlCreatedAt, result.CreatedAt.String())
	}
	if result.UpdatedAt != nil {
		d.Set(dlUpdatedAt, result.UpdatedAt.String())
	}

	sakReKey := map[string]interface{}{}
	if result.SakRekey != nil {
		gatewaySakRekeyIntf := result.SakRekey
		gatewaySakRekey := gatewaySakRekeyIntf.(*directlinkv1.SakRekey)
		sakReKey[dlGatewaySakRekeyMode] = *gatewaySakRekey.Mode
		if gatewaySakRekey.Interval != nil {
			sakReKey[dlGatewaySakRekeyInterval] = *gatewaySakRekey.Interval
		}
	}
	d.Set(dlGatewaySakRekey, []map[string]interface{}{sakReKey})

	statusReasonsList := make([]map[string]interface{}, 0)
	if len(result.StatusReasons) > 0 {
		for _, statusReason := range result.StatusReasons {
			statusReasonItem := map[string]interface{}{}
			statusReasonItem[dlGatewayMacsecSatusReasonCode] = statusReason.Code
			statusReasonItem[dlGatewayMacsecSatusReasonMessage] = statusReason.Message
			statusReasonItem[dlGatewayMacsecSatusReasonMoreInfo] = statusReason.MoreInfo
			statusReasonsList = append(statusReasonsList, statusReasonItem)
		}
		d.Set(dlGatewayMacsecSatusReasons, statusReasonsList)
	}

	d.SetId(dlGatewayID)

	return nil
}
