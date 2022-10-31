// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.ibm.com/ibmcloud/networking-go-sdk/directlinkv1"
)

func ResourceIBMDLASPrepends() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMDLASPrependsCreate,
		Read:     resourceIBMDLASPrependsRead,
		Delete:   resourceIBMDLASPrependsDelete,
		Exists:   resourceIBMDLASPrependsExists,
		Update:   resourceIBMDLASPrependsUpdate,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			dlGatewayId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Direct Link gateway identifier",
			},
			dlAsPrepends: {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    false,
				Description: "List of AS Prepend configuration information",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlCreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time AS Prepend was created",
						},
						dlResourceId: {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    false,
							Computed:    true,
							Description: "The unique identifier for this AS Prepend",
						},
						dlLength: {
							Type:         schema.TypeInt,
							Required:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_dl_as_prepends", dlLength),
							Description:  "Number of times the ASN to appended to the AS Path",
						},
						dlPolicy: {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_dl_as_prepends", dlPolicy),
							Description:  "Route type this AS Prepend applies to",
						},
						dlPrefix: {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    false,
							Description: "Comma separated list of prefixes this AS Prepend applies to. Maximum of 10 prefixes. If not specified, this AS Prepend applies to all prefixes",
						},
						dlSpecificPrefixes: {
							Type:        schema.TypeList,
							Description: "Array of prefixes this AS Prepend applies to",
							Optional:    true,
							ForceNew:    false,
							MinItems:    1,
							MaxItems:    10,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						dlUpdatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time AS Prepend was updated",
						},
					},
				},
			},
		},
	}
}
func ResourceIBMDLASPrependsValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	dlPolicyAllowedValues := "export, import"

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 dlPolicy,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              dlPolicyAllowedValues})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 dlLength,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Required:                   true,
			MinValue:                   "3",
			MaxValue:                   "10"})

	ibmDLASPrependsValidator := validate.ResourceValidator{ResourceName: "ibm_dl_as_prepends", Schema: validateSchema}

	return &ibmDLASPrependsValidator
}
func resourceIBMDLASPrependsCreate(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	createGatewayVCOptions := &directlinkv1.CreateGatewayVirtualConnectionOptions{}

	gatewayId := d.Get(dlGatewayId).(string)
	createGatewayVCOptions.SetGatewayID(gatewayId)
	vcName := d.Get(dlVCName).(string)
	createGatewayVCOptions.SetName(vcName)
	vcType := d.Get(dlVCType).(string)
	createGatewayVCOptions.SetType(vcType)

	if _, ok := d.GetOk(dlVCNetworkId); ok {
		vcNetworkId := d.Get(dlVCNetworkId).(string)
		createGatewayVCOptions.SetNetworkID(vcNetworkId)
	}

	gatewayVC, response, err := directLink.CreateGatewayVirtualConnection(createGatewayVCOptions)
	if err != nil {
		log.Printf("[DEBUG] Create Direct Link Gateway (Dedicated) Virtual connection err %s\n%s", err, response)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", gatewayId, *gatewayVC.ID))
	d.Set(dlVirtualConnectionId, *gatewayVC.ID)
	return resourceIBMdlGatewayVCRead(d, meta)
}

func resourceIBMDLASPrependsRead(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}

	gatewayId := parts[0]
	ID := parts[1]

	getGatewayVirtualConnectionOptions := &directlinkv1.GetGatewayVirtualConnectionOptions{}
	getGatewayVirtualConnectionOptions.SetGatewayID(gatewayId)
	getGatewayVirtualConnectionOptions.SetID(ID)
	instance, response, err := directLink.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error Getting Directlink Gateway Connection (%s): %s\n%s", ID, err, response)
	}

	if instance.Name != nil {
		d.Set(dlVCName, *instance.Name)
	}
	if instance.Type != nil {
		d.Set(dlVCType, *instance.Type)
	}
	if instance.NetworkAccount != nil {
		d.Set(dlVCNetworkAccount, *instance.NetworkAccount)
	}
	if instance.NetworkID != nil {
		d.Set(dlVCNetworkId, *instance.NetworkID)
	}
	if instance.CreatedAt != nil {
		d.Set(dlVCCreatedAt, instance.CreatedAt.String())
	}
	if instance.Status != nil {
		d.Set(dlVCStatus, *instance.Status)
	}
	d.Set(dlVirtualConnectionId, *instance.ID)
	d.Set(dlGatewayId, gatewayId)
	getGatewayOptions := &directlinkv1.GetGatewayOptions{
		ID: &gatewayId,
	}
	dlgw, response, err := directLink.GetGateway(getGatewayOptions)
	if err != nil {
		return fmt.Errorf("[ERROR] Error Getting Direct Link Gateway (Dedicated Template): %s\n%s", err, response)
	}
	d.Set(flex.RelatedCRN, *dlgw.Crn)
	return nil
}

func resourceIBMDLASPrependsUpdate(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}

	gatewayId := parts[0]
	ID := parts[1]

	getVCOptions := &directlinkv1.GetGatewayVirtualConnectionOptions{
		ID: &ID,
	}
	getVCOptions.SetGatewayID(gatewayId)
	_, detail, err := directLink.GetGatewayVirtualConnection(getVCOptions)

	if err != nil {
		log.Printf("Error fetching Direct Link Gateway (Dedicated Template) Virtual Connection:%s", detail)
		return err
	}

	updateGatewayVCOptions := &directlinkv1.UpdateGatewayVirtualConnectionOptions{}
	updateGatewayVCOptions.ID = &ID
	updateGatewayVCOptions.SetGatewayID(gatewayId)
	if d.HasChange(dlName) {
		if d.Get(dlName) != nil {
			name := d.Get(dlName).(string)
			updateGatewayVCOptions.Name = &name
		}
	}

	_, response, err := directLink.UpdateGatewayVirtualConnection(updateGatewayVCOptions)
	if err != nil {
		log.Printf("[DEBUG] Update Direct Link Gateway (Dedicated) Virtual Connection err %s\n%s", err, response)
		return err
	}

	return resourceIBMdlGatewayVCRead(d, meta)
}

func resourceIBMDLASPrependsDelete(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}

	gatewayId := parts[0]
	ID := parts[1]
	delVCOptions := &directlinkv1.DeleteGatewayVirtualConnectionOptions{
		ID: &ID,
	}
	delVCOptions.SetGatewayID(gatewayId)
	response, err := directLink.DeleteGatewayVirtualConnection(delVCOptions)

	if err != nil && response.StatusCode != 404 {
		log.Printf("Error deleting Direct Link Gateway (Dedicated Template) Virtual Connection: %s", response)
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMDLASPrependsExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return false, err
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, err
	}
	if len(parts) < 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: Id should be a combination of gatewayID/gatewayVCID", d.Id())
	}
	gatewayId := parts[0]
	ID := parts[1]

	getVCOptions := &directlinkv1.GetGatewayVirtualConnectionOptions{
		ID: &ID,
	}
	getVCOptions.SetGatewayID(gatewayId)
	_, response, err := directLink.GetGatewayVirtualConnection(getVCOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error Getting Direct Link Gateway (Dedicated Template) Virtual Connection: %s\n%s", err, response)
	}

	if response.StatusCode == 404 {
		d.SetId("")
		return false, nil
	}
	return true, nil
}
