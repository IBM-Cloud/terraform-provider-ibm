// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"log"

	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMDLGatewayMacsecCaks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDLGatewayMacsecCaksRead,
		Schema: map[string]*schema.Schema{
			dlGatewayId: {
				Type:        schema.TypeString,
				Description: "Gateway ID",
				Required:    true,
			},
			dlGatewayMAcsecVersion: {
				Type:        schema.TypeString,
				Description: "Requests the version of the API as a date in the format YYYY-MM-DD.",
				Required:    true,
			},
			dlGatewayMacsecCaksList: {
				Type:        schema.TypeList,
				Description: "Determines how SAK rekeying occurs.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewayMacsecCakID: {
							Type:        schema.TypeString,
							Description: "CAK ID",
							Computed:    true,
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
				},
			},
		},
	}
}

func dataSourceIBMDLGatewayMacsecCaksRead(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	gatewayID := d.Get(dlGatewayId).(string)

	if err != nil {
		return err
	}

	// Get Gateway MAcsec CAK
	// Construct an instance of the GetGatewayMacsecCakOptions model
	listGatewayMacsecCaksOptions := new(directlinkv1.ListGatewayMacsecCaksOptions)
	listGatewayMacsecCaksOptions.ID = &gatewayID
	// listGatewayMacsecCaksOptions.Version = &dlGatewayMAcsecVersion

	listGatewayMacsecCaksOptions.Headers = map[string]string{"x-custom-header": "x-custom-value"}
	// Expect response parsing to fail since we are receiving a text/plain response
	result, response, err := directLink.ListGatewayMacsecCaks(listGatewayMacsecCaksOptions)

	if err != nil {
		log.Println("[WARN] Error Get DL Gateway Macsec", response, err)
		return err
	}

	caksList := make([]map[string]interface{}, 0)
	if len(result.Caks) > 0 {
		for _, cak := range result.Caks {
			cakItem := map[string]interface{}{}

			cakItem[dlGatewayMacsecCakName] = *cak.Name
			cakItem[dlGatewayMacsecCakSession] = *cak.Session
			cakItem[dlGatewayMacsecCakStatus] = *cak.Status
			cakItem[dlCreatedAt] = cak.CreatedAt.String()
			cakItem[dlUpdatedAt] = cak.UpdatedAt.String()
			cakItem[dlGatewayMacsecCakID] = *cak.ID

			hpcsKey := map[string]interface{}{}
			if cak.Key != nil {
				hpcsKey[dlGatewayMacsecHPCSCrn] = *cak.Key.Crn
				cakItem[dlGatewayMacsecHPCSKey] = []map[string]interface{}{hpcsKey}
			}

			activeDelta := map[string]interface{}{}
			if cak.ActiveDelta != nil {
				hpcsKey := map[string]interface{}{}
				if cak.ActiveDelta.Key != nil {
					hpcsKey[dlGatewayMacsecHPCSCrn] = cak.ActiveDelta.Key.Crn
					activeDelta[dlGatewayMacsecHPCSKey] = []map[string]interface{}{hpcsKey}
				}
				activeDelta[dlGatewayMacsecCakName] = cak.ActiveDelta.Name
				// activeDelta[dlGatewayMacsecCakStatus] = *result.ActiveDelta.Status
				cakItem[dlGatewayMacsecCakActiveDelta] = []map[string]interface{}{activeDelta}
			}

			caksList = append(caksList, cakItem)
		}
	}

	d.Set(dlGatewayMacsecCaksList, caksList)
	d.SetId(gatewayID)

	return nil
}
