// Copyright IBM Corp. 2017, 2021, 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"context"
	"log"

	"github.com/IBM/networking-go-sdk/directlinkv1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMDLGatewayMacsecConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMdlGatewayMacsecConfigCreate,
		ReadContext:   resourceIBMdlGatewayMacsecConfigRead,
		DeleteContext: resourceIBMdlGatewayMacsecConfigDelete,
		UpdateContext: resourceIBMdlGatewayMacsecConfigUpdate,
		Importer:      &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			dlGatewayId: {
				Type:        schema.TypeString,
				Description: "Gateway ID",
				Required:    true,
			},
			dlActive: {
				Type:        schema.TypeBool,
				Required:    true,
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
				Required:    true,
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
				Optional:    true,
				Description: "The window size determines the number of frames in a window for replay protection.",
			},
			dlGatewaySakRekey: {
				Type:        schema.TypeSet,
				Description: "Determines how SAK rekeying occurs.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewaySakRekeyInterval: {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "SAK ReKey Interval",
						},
						dlGatewaySakRekeyMode: {
							Type:        schema.TypeString,
							Required:    true,
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
			dlGatewayMacsecCaksList: {
				Type:        schema.TypeList,
				Description: "CAKs",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewayMacsecHPCSKey: {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "HPCS Key",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									dlGatewayMacsecHPCSCrn: {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The CRN of the referenced key.",
									},
								},
							},
						},
						dlGatewayMacsecCakName: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name identifies the connectivity association key (CAK) within the MACsec key chain.",
						},
						dlGatewayMacsecCakSession: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Current status of the instance.",
						},
					},
				},
			},
		},
	}
}

func resourceIBMdlGatewayMacsecConfigCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	gatewayID := d.Get(dlGatewayId).(string)
	active := d.Get(dlActive).(bool)

	caks := d.Get(dlGatewayMacsecCaksList).([]interface{})
	caksList := []directlinkv1.GatewayMacsecCakPrototype{}
	for _, cak := range caks {
		cakMap := cak.(map[string]interface{})

		name := cakMap[dlGatewayMacsecCakName].(string)
		session := cakMap[dlGatewayMacsecCakSession].(string)

		keyMap := cakMap[dlGatewayMacsecHPCSKey].(*schema.Set).List()[0].(map[string]interface{})
		crn := keyMap[dlGatewayMacsecHPCSCrn].(string)
		keyItem := directlinkv1.HpcsKeyIdentity{Crn: &crn}

		cakItem := &directlinkv1.GatewayMacsecCakPrototype{
			Key:     &keyItem,
			Name:    &name,
			Session: &session,
		}

		caksList = append(caksList, *cakItem)
	}

	securityPolicy := d.Get(dlSecurityPolicy).(string)
	windowSize := int64(d.Get(dlWindowSize).(int))

	opts := &directlinkv1.SetGatewayMacsecOptions{
		ID:             &gatewayID,
		Active:         &active,
		Caks:           caksList,
		SecurityPolicy: &securityPolicy,
		WindowSize:     &windowSize,
	}
	sakReKeyIntf := d.Get(dlGatewaySakRekey).(*schema.Set).List()[0].(map[string]interface{})
	mode := sakReKeyIntf["mode"].(string)

	if val, ok := sakReKeyIntf["interval"]; ok {
		intervalInt64 := int64(val.(int))
		sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
		sakRekeyPrototypeModel.Interval = &intervalInt64
		sakRekeyPrototypeModel.Mode = &mode
		opts.SetSakRekey(sakRekeyPrototypeModel)

	} else {
		sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyPacketNumberRolloverModePrototype)
		sakRekeyPrototypeModel.Mode = &mode
		opts.SetSakRekey(sakRekeyPrototypeModel)
	}
	opts.Version = IBMCLOUD_DL_VERSION_DEFAULT
	result, response, err := directLink.SetGatewayMacsec(opts)

	if err != nil {
		log.Printf("Error setting Direct Link Gateway Macsec Config : %s", response)
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

	d.SetId(gatewayID)

	return nil
}

func resourceIBMdlGatewayMacsecConfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	getGatewayMacsecOptionsModel.Version = IBMCLOUD_DL_VERSION_DEFAULT

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

func resourceIBMdlGatewayMacsecConfigUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	gatewayID := d.Id()
	GatewayMacsecConfigPatch := map[string]interface{}{}

	if d.HasChange(dlActive) {
		name := d.Get(dlActive).(bool)
		GatewayMacsecConfigPatch[dlActive] = &name
	}

	if d.HasChange(dlSecurityPolicy) {
		policy := d.Get(dlSecurityPolicy).(string)
		GatewayMacsecConfigPatch[dlSecurityPolicy] = &policy
	}

	if d.HasChange(dlWindowSize) {
		windoSize := d.Get(dlWindowSize).(string)
		GatewayMacsecConfigPatch[dlWindowSize] = &windoSize
	}

	if d.HasChange(dlGatewaySakRekey) {
		sakReKey := d.Get(dlGatewaySakRekey).(map[string]interface{})
		GatewayMacsecConfigPatch[dlGatewaySakRekey] = &sakReKey
	}

	updateGatewayMacsecOptions := directLink.NewUpdateGatewayMacsecOptions(gatewayID, GatewayMacsecConfigPatch)
	updateGatewayMacsecOptions.Version = IBMCLOUD_DL_VERSION_DEFAULT

	_, response, err := directLink.UpdateGatewayMacsec(updateGatewayMacsecOptions)
	if err != nil {
		log.Printf("[DEBUG] Update Direct Link Gateway Macsec Config err %s\n%s", err, response)
		return diag.FromErr(err)
	}

	return resourceIBMdlGatewayMacsecConfigRead(context, d, meta)
}

func resourceIBMdlGatewayMacsecConfigDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	gatewayID := d.Id()

	delOptions := &directlinkv1.UnsetGatewayMacsecOptions{
		ID: &gatewayID,
		Version: IBMCLOUD_DL_VERSION_DEFAULT
	}

	response, err := directLink.UnsetGatewayMacsec(delOptions)

	if err != nil && response.StatusCode != 404 {
		log.Printf("Error unsetting Direct Link Gateway Macsec Config : %s", response)
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
