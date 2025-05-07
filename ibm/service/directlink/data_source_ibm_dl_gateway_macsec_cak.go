// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"log"

	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMDLGatewayMacsecCak() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDLGatewayMacsecCakRead,
		Schema: map[string]*schema.Schema{
			ID: {
				Type:        schema.TypeString,
				Description: "Gateway ID",
				Required:    true,
			},
			dlGatewayMacsecCakID: {
				Type:        schema.TypeString,
				Description: "CAK ID",
				Required:    true,
			},
			dlGatewayMAcsecVersion: {
				Type:        schema.TypeString,
				Description: "Requests the version of the API as a date in the format YYYY-MM-DD.",
				Required:    true,
			},
			dlCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the resource was created",
			},
			dlGatewayMacsecCakName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name identifies the connectivity association key (CAK) within the MACsec key chain.",
			},
			dlGatewayMacsecCakSession: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The intended session the key will be used to secure.",
			},
			dlGatewayMacsecCakStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the CAK.",
			},
			dlUpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the resource was last updated",
			},
			dlGatewayMacsecHPCSKey: {
				Type:        schema.TypeSet,
				Description: "HPCS Key",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewayMacsecHPCSCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the referenced key.",
						},
					},
				},
			},
			dlGatewayMacsecCakActiveDelta: {
				Type:        schema.TypeSet,
				Description: "CAK Active Delta",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewayMacsecHPCSKey: {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "HPCS Key",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									dlGatewayMacsecHPCSCrn: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN of the referenced key.",
									},
								},
							},
						},
						dlGatewayMacsecCakName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name identifies the connectivity association key (CAK) within the MACsec key chain.",
						},
						// dlGatewayMacsecCakStatus: {
						// 	Type:        schema.TypeString,
						// 	Computed:    true,
						// 	Description: "Current status of the CAK.",
						// },
					},
				},
			},
		},
	}
}

func dataSourceIBMDLGatewayMacsecCakRead(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	gatewayID := d.Get(ID).(string)
	getMacsecCakID := d.Get(dlGatewayMacsecCakID).(string)

	if err != nil {
		return err
	}

	// Get Gateway MAcsec CAK
	// Construct an instance of the GetGatewayMacsecCakOptions model
	getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
	getGatewayMacsecCakOptionsModel.ID = &gatewayID
	getGatewayMacsecCakOptionsModel.CakID = &getMacsecCakID
	// getGatewayMacsecCakOptionsModel.Version = &dlGatewayMAcsecVersion

	getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
	// Expect response parsing to fail since we are receiving a text/plain response
	result, response, err := directLink.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)

	if err != nil {
		log.Println("[WARN] Error Get DL Gateway Macsec", response, err)
		return err
	}

	if result.Status != nil {
		d.Set(dlMacSecConfigStatus, *result.Status)
	}
	if result.Name != nil {
		d.Set(dlGatewayMacsecCakName, *result.Name)
	}
	if result.Session != nil {
		d.Set(dlGatewayMacsecCakSession, *result.Session)
	}
	if result.Status != nil {
		d.Set(dlGatewayMacsecCakStatus, *result.Status)
	}
	if result.CreatedAt != nil {
		d.Set(dlCreatedAt, result.CreatedAt.String())
	}
	if result.UpdatedAt != nil {
		d.Set(dlUpdatedAt, result.UpdatedAt.String())
	}

	hpcsKey := map[string]interface{}{}
	if result.Key != nil {
		hpcsKey[dlGatewayMacsecHPCSCrn] = *result.Key.Crn
		d.Set(dlGatewayMacsecHPCSKey, hpcsKey)
	}

	activeDelta := map[string]interface{}{}
	if result.ActiveDelta != nil {
		hpcsKey := map[string]interface{}{}
		if result.ActiveDelta.Key != nil {
			hpcsKey[dlGatewayMacsecHPCSCrn] = *result.ActiveDelta.Key.Crn
			activeDelta[dlGatewayMacsecHPCSKey] = hpcsKey
		}

		activeDelta[dlGatewayMacsecCakName] = *result.ActiveDelta.Name
		// activeDelta[dlGatewayMacsecCakStatus] = *result.ActiveDelta.Status
		d.Set(dlGatewayMacsecCakActiveDelta, activeDelta)
	}

	d.Set(dlGatewayMacsecCakID, result.ID)
	d.SetId(gatewayID)

	return nil
}
