// Copyright IBM Corp. 2017, 2021, 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/networking-go-sdk/directlinkv1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func ResourceIBMDLGatewayMacsecConfig() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMdlGatewayMacsecConfigUpdate,
		Read:     resourceIBMdlGatewayMacsecConfigRead,
		Delete:   resourceIBMdlGatewayMacsecConfigDelete,
		Exists:   resourceIBMdlGatewayMacsecConfigExists,
		Update:   resourceIBMdlGatewayMacsecConfigUpdate,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourceTagsCustomizeDiff(diff)
			},
		),
		Schema: map[string]*schema.Schema{
			ID: {
				Type:        schema.TypeString,
				Description: "Gateway ID",
				Required:    true,
			},
			dlActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Indicate whether MACsec protection should be active (true) or inactive (false) for this MACsec enabled gateway",
			},
			dlSecurityPolicy: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Determines how packets without MACsec headers are handled.",
			},
			dlWindowSize: {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The window size determines the number of frames in a window for replay protection.",
			},
			dlGatewaySakRekey: {
				Type:        schema.TypeList,
				Description: "Determines how SAK rekeying occurs.",
				Computed:    true,
				Optional:    true,
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
			dlGatewayMacsecCaksList: {
				Type:        schema.TypeList,
				Description: "Determines how SAK rekeying occurs.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewayMacsecCakName: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name identifies the connectivity association key (CAK) within the MACsec key chain.",
						},
						dlGatewayMacsecCakSession: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The intended session the key will be used to secure.",
						},
						dlGatewayMacsecHPCSKey: {
							Type:        schema.TypeSet,
							Description: "HPCS Key",
							Required:    true,
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
					},
				},
			},
			dlGatewayMacsecCak: {
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

// func ResourceIBMdlGatewayMacsecConfigValidator() *validate.ResourceValidator {

// 	validateSchema := make([]validate.ValidateSchema, 0)
// 	dlSessionValues := "primary, fallback"
// 	dlStatusValues := "operational, rotating, active, inactive, failed"

// 	validateSchema = append(validateSchema,
// 		validate.ValidateSchema{
// 			Identifier:                 dlGatewayMacsecConfigSession,
// 			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
// 			Type:                       validate.TypeString,
// 			Required:                   true,
// 			AllowedValues:              dlSessionValues})
// 	validateSchema = append(validateSchema,
// 		validate.ValidateSchema{
// 			Identifier:                 dlGatewayMacsecConfigStatus,
// 			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
// 			Type:                       validate.TypeString,
// 			Required:                   true,
// 			AllowedValues:              dlStatusValues})
// 	validateSchema = append(validateSchema,
// 		validate.ValidateSchema{
// 			Identifier:                 dlGatewayMacsecConfigName,
// 			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
// 			Type:                       validate.TypeString,
// 			Required:                   true,
// 			Regexp:                     `^([0-9a-fA-F]{2}){1,32}$`,
// 			MinValueLength:             2,
// 			MaxValueLength:             64})
// 	validateSchema = append(validateSchema,
// 		validate.ValidateSchema{
// 			Identifier:                 dlGatewayMacsecHPCSCrn,
// 			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
// 			Type:                       validate.TypeString,
// 			Optional:                   true,
// 			Regexp:                     `^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@/]|%[0-9A-Z]{2})*){2}:hs-crypto(:([A-Za-z0-9-._~!$&'()*+,;=@/]|%[0-9A-Z]{2})*){5}$`,
// 			MinValueLength:             1,
// 			MaxValueLength:             128})

// 	ibmISDLGatewayResourceValidator := validate.ResourceValidator{ResourceName: "ibm_dl_gateway_macsec_cak", Schema: validateSchema}
// 	return &ibmISDLGatewayResourceValidator
// }

// func resourceIBMdlGatewayMacsecConfigCreate(d *schema.ResourceData, meta interface{}) error {
// 	directLink, err := directlinkClient(meta)
// 	if err != nil {
// 		return err
// 	}

// 	gatewayID := d.Get(ID).(string)

// 	active := d.Get(dlActive).(bool)
// 	securityPolicy := d.Get(dlSecurityPolicy).(string)
// 	// window := d.Get(dlWindowSize).(string)

// 	caksPrototype := d.Get(dlGatewayMacsecCaksList).([]map[string]interface{})
// 	caksList := new([]directlinkv1.GatewayMacsecCakPrototype)
// 	for _, cakPrototype := range caksPrototype {
// 		cakItem := new(directlinkv1.GatewayMacsecCakPrototype)
// 		cakItem.Name = cakPrototype[dlGatewayMacsecCakName].(string)
// 		cakItem.Session = cakPrototype[dlGatewayMacsecCakSession]

// 		keyMap := cakPrototype[dlGatewayMacsecHPCSKey]
// 		keyMapIntf := keyMap.(map[string]interface{})
// 		crn := keyMapIntf[dlGatewayMacsecHPCSCrn].(string)
// 		key, _ := directLink.NewHpcsKeyIdentity(crn)
// 		cakItem.Key = key

// 		caksList = append(caksList, cakItem)
// 	}

// 	sakReKey := d.Get(dlGatewaySakRekey).(map[string]interface{})
// 	sakReKeyIntf := new(directlinkv1.SakRekey)
// 	if sakReKey[dlGatewaySakRekeyInterval] != nil {
// 		sakReKeyIntf.Interval = sakReKey[dlGatewaySakRekeyInterval].(int)
// 	}
// 	if sakReKey[dlGatewaySakRekeyMode] != nil {
// 		sakReKeyIntf.Mode = sakReKey[dlGatewaySakRekeyMode].(string)
// 	}

// 	setGatewayMacsecConfigOptions := directLink.NewSetGatewayMacsecOptions(gatewayID, active, caksList, sakReKeyIntf, securityPolicy)

// 	_, response, err := directLink.SetGatewayMacsec(setGatewayMacsecConfigOptions)
// 	if err != nil {
// 		return fmt.Errorf("[DEBUG] Set Direct Link Gateway Macsec - err %s\n%s", err, response)
// 	}

// 	d.SetId(gatewayID)

// 	return resourceIBMdlGatewayMacsecConfigRead(d, meta)
// }

func resourceIBMdlGatewayMacsecConfigRead(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	dlGatewayID := d.Get(ID).(string)

	// Get MacSec gateway
	// Construct an instance of the GetGatewayMacsecOptions model
	getGatewayMacsecOptionsModel := new(directlinkv1.GetGatewayMacsecOptions)
	getGatewayMacsecOptionsModel.ID = &dlGatewayID
	getGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	result, response, err := directLink.GetGatewayMacsec(getGatewayMacsecOptionsModel)
	if err != nil {
		log.Println("[WARN] Error Get DL Gateway Macsec", response, err)
		return err
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
		d.Set(dlGatewaySakRekey, sakReKey)
	}

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

func resourceIBMdlGatewayMacsecConfigUpdate(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
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

		if result.SakRekey != nil {
			gatewaySakRekeyIntf := result.SakRekey
			gatewaySakRekey := gatewaySakRekeyIntf.(*directlinkv1.SakRekey)
			sakReKey[dlGatewaySakRekeyMode] = *gatewaySakRekey.Mode
			if gatewaySakRekey.Interval != nil {
				sakReKey[dlGatewaySakRekeyInterval] = *gatewaySakRekey.Interval
			}
			d.Set(dlGatewaySakRekey, sakReKey)
		}

		GatewayMacsecConfigPatch[dlWindowSize] = &windoSize
	}




	patchGatewayOptions := directLink.NewUpdateGatewayMacsecConfigOptions(gatewayID, getMacsecCakID, GatewayMacsecConfigPatch)
	_, response, err := directLink.UpdateGatewayMacsecConfig(patchGatewayOptions)
	if err != nil {
		log.Printf("[DEBUG] Update Direct Link Gateway Macsec CAK err %s\n%s", err, response)
		return err
	}

	return resourceIBMdlGatewayMacsecConfigRead(d, meta)
}

func resourceIBMdlGatewayMacsecConfigDelete(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	gatewayID := d.Id()
	getMacsecCakID := d.Get(dlGatewayMacsecConfigID).(string)

	delOptions := &directlinkv1.DeleteGatewayMacsecConfigOptions{
		ID:    &gatewayID,
		CakID: &getMacsecCakID,
	}

	response, err := directLink.DeleteGatewayMacsecConfig(delOptions)

	if err != nil && response.StatusCode != 404 {
		log.Printf("Error deleting Direct Link Gateway Macsec CAK : %s", response)
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMdlGatewayMacsecConfigExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return false, err
	}

	gatewayID := d.Id()
	getMacsecCakID := d.Get(dlGatewayMacsecConfigID).(string)

	// Get Gateway MAcsec CAK
	// Construct an instance of the GetGatewayMacsecConfigOptions model
	getGatewayMacsecConfigOptionsModel := new(directlinkv1.GetGatewayMacsecConfigOptions)
	getGatewayMacsecConfigOptionsModel.ID = &gatewayID
	getGatewayMacsecConfigOptionsModel.CakID = &getMacsecCakID
	getGatewayMacsecConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
	// Expect response parsing to fail since we are receiving a text/plain response
	instance, response, err := directLink.GetGatewayMacsecConfig(getGatewayMacsecConfigOptionsModel)

	if (err != nil) || (instance == nil) {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error Getting Direct Link Gateway Macsec CAK : %s\n%s", err, response)
	}

	return true, nil
}
