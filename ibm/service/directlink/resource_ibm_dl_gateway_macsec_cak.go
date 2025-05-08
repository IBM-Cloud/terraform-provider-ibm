// Copyright IBM Corp. 2017, 2021, 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"fmt"
	"log"

	"github.com/IBM/networking-go-sdk/directlinkv1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

func ResourceIBMDLGatewayMacsecCak() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMdlGatewayMacsecCakCreate,
		Read:     resourceIBMdlGatewayMacsecCakRead,
		Delete:   resourceIBMdlGatewayMacsecCakDelete,
		Exists:   resourceIBMdlGatewayMacsecCakExists,
		Update:   resourceIBMdlGatewayMacsecCakUpdate,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			dlGatewayId: {
				Type:        schema.TypeString,
				Description: "Gateway ID",
				Computed:    true,
			},
			dlGatewayMacsecCakID: {
				Type:        schema.TypeString,
				Description: "CAK ID",
				Computed:    true,
			},
			dlGatewayMAcsecVersion: {
				Type:        schema.TypeString,
				Description: "Requests the version of the API as a date in the format YYYY-MM-DD.",
				Required:    true,
			},
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
							Description: "Current status of the instance.",
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
									// 	Description: "Current status of the instance.",
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

func ResourceIBMdlGatewayMacsecCakValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	dlSessionValues := "primary, fallback"
	dlStatusValues := "operational, rotating, active, inactive, failed"

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 dlGatewayMacsecCakSession,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              dlSessionValues})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 dlGatewayMacsecCakStatus,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              dlStatusValues})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 dlGatewayMacsecCakName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([0-9a-fA-F]{2}){1,32}$`,
			MinValueLength:             2,
			MaxValueLength:             64})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 dlGatewayMacsecHPCSCrn,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@/]|%[0-9A-Z]{2})*){2}:hs-crypto(:([A-Za-z0-9-._~!$&'()*+,;=@/]|%[0-9A-Z]{2})*){5}$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISDLGatewayResourceValidator := validate.ResourceValidator{ResourceName: "ibm_dl_gateway_macsec_cak", Schema: validateSchema}
	return &ibmISDLGatewayResourceValidator
}

func resourceIBMdlGatewayMacsecCakCreate(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	gatewayID := d.Get(dlGatewayId).(string)
	name := d.Get(dlGatewayMacsecCakName).(string)
	// version := d.Get(dlGatewayMAcsecVersion).(string)
	session := d.Get(dlGatewayMacsecCakSession).(string)
	keyMapIntf := d.Get(dlGatewayMacsecHPCSKey).(map[string]interface{})
	crn := keyMapIntf[dlCrn].(string)
	key, _ := directLink.NewHpcsKeyIdentity(crn)

	createGatewayMacsecCakOptions := directLink.NewCreateGatewayMacsecCakOptions(gatewayID, key, name, session)

	result, response, err := directLink.CreateGatewayMacsecCak(createGatewayMacsecCakOptions)
	if err != nil {
		return fmt.Errorf("[DEBUG] Create Direct Link Gateway Macsec CAK - err %s\n%s", err, response)
	}

	d.SetId(gatewayID)
	d.Set(dlGatewayMacsecCakID, *result.ID)

	return resourceIBMdlGatewayMacsecCakRead(d, meta)
}

func resourceIBMdlGatewayMacsecCakRead(d *schema.ResourceData, meta interface{}) error {
	dtype := d.Get(dlType).(string)
	log.Printf("[INFO] Inside resourceIBMdlGatewayMacsecCakRead: %s", dtype)

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	gatewayID := d.Id()
	getMacsecCakID := d.Get(dlGatewayMacsecCakID).(string)

	// Get Gateway MAcsec CAK
	// Construct an instance of the GetGatewayMacsecCakOptions model
	getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
	getGatewayMacsecCakOptionsModel.ID = &gatewayID
	getGatewayMacsecCakOptionsModel.CakID = &getMacsecCakID
	getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
	// Expect response parsing to fail since we are receiving a text/plain response
	instance, response, err := directLink.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)
	if err != nil {
		log.Println("[WARN] Error Get DL Gateway Macsec CAK ", response, err)
		return err
	}

	cakItem := map[string]interface{}{}
	if instance.Status != nil {
		cakItem[dlGatewayMacsecCakStatus] = *instance.Status
	}
	if instance.Name != nil {
		cakItem[dlGatewayMacsecCakName] = *instance.Name
	}
	if instance.Session != nil {
		cakItem[dlGatewayMacsecCakSession] = *instance.Session
	}
	if instance.CreatedAt != nil {
		cakItem[dlCreatedAt] = *instance.CreatedAt
	}
	if instance.UpdatedAt != nil {
		cakItem[dlUpdatedAt] = *instance.UpdatedAt
	}

	cakItem[dlGatewayMacsecCakID] = *instance.ID

	hpcsKey := map[string]interface{}{}
	if instance.Key != nil {
		hpcsKey[dlGatewayMacsecHPCSCrn] = *instance.Key.Crn
		cakItem[dlGatewayMacsecHPCSKey] = hpcsKey
	}

	activeDelta := map[string]interface{}{}
	if instance.ActiveDelta != nil {
		hpcsKey := map[string]interface{}{}
		if instance.ActiveDelta.Key != nil {
			hpcsKey[dlGatewayMacsecHPCSCrn] = *instance.ActiveDelta.Key.Crn
			activeDelta[dlGatewayMacsecHPCSKey] = hpcsKey
		}

		activeDelta[dlGatewayMacsecCakName] = *instance.ActiveDelta.Name
		// activeDelta[dlGatewayMacsecCakStatus] = *instance.ActiveDelta.Status
		cakItem[dlGatewayMacsecCakActiveDelta] = activeDelta
	}

	d.Set(dlGatewayMacsecCak, cakItem)
	d.Set(dlGatewayMacsecCakID, instance.ID)
	d.SetId(gatewayID)
	return nil
}

func resourceIBMdlGatewayMacsecCakUpdate(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	gatewayID := d.Id()
	getMacsecCakID := d.Get(dlGatewayMacsecCakID).(string)

	gatewayMacsecCakPatch := map[string]interface{}{}

	if d.HasChange(dlGatewayMacsecCakName) {
		name := d.Get(dlGatewayMacsecCakName).(string)
		gatewayMacsecCakPatch[dlGatewayMacsecCakName] = &name
	}

	if d.HasChange(dlGatewayMacsecHPCSKey) {
		keyMapIntf := d.Get(dlGatewayMacsecHPCSKey).(map[string]interface{})
		crn := keyMapIntf[dlCrn].(string)
		key, _ := directLink.NewHpcsKeyIdentity(crn)
		gatewayMacsecCakPatch[dlGatewayMacsecHPCSKey] = &key
	}

	patchGatewayOptions := directLink.NewUpdateGatewayMacsecCakOptions(gatewayID, getMacsecCakID, gatewayMacsecCakPatch)
	_, response, err := directLink.UpdateGatewayMacsecCak(patchGatewayOptions)
	if err != nil {
		log.Printf("[DEBUG] Update Direct Link Gateway Macsec CAK err %s\n%s", err, response)
		return err
	}

	return resourceIBMdlGatewayMacsecCakRead(d, meta)
}

func resourceIBMdlGatewayMacsecCakDelete(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	gatewayID := d.Id()
	getMacsecCakID := d.Get(dlGatewayMacsecCakID).(string)

	delOptions := &directlinkv1.DeleteGatewayMacsecCakOptions{
		ID:    &gatewayID,
		CakID: &getMacsecCakID,
	}

	response, err := directLink.DeleteGatewayMacsecCak(delOptions)

	if err != nil && response.StatusCode != 404 {
		log.Printf("Error deleting Direct Link Gateway Macsec CAK : %s", response)
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMdlGatewayMacsecCakExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return false, err
	}

	gatewayID := d.Id()
	getMacsecCakID := d.Get(dlGatewayMacsecCakID).(string)

	// Get Gateway MAcsec CAK
	// Construct an instance of the GetGatewayMacsecCakOptions model
	getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
	getGatewayMacsecCakOptionsModel.ID = &gatewayID
	getGatewayMacsecCakOptionsModel.CakID = &getMacsecCakID
	getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
	// Expect response parsing to fail since we are receiving a text/plain response
	instance, response, err := directLink.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)

	if (err != nil) || (instance == nil) {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error Getting Direct Link Gateway Macsec CAK : %s\n%s", err, response)
	}

	return true, nil
}
