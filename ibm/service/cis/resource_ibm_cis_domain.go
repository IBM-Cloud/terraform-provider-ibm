// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"log"
	"net/http"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisDomain                    = "domain"
	cisDomainPaused              = "paused"
	cisDomainStatus              = "status"
	cisDomainNameServers         = "name_servers"
	cisDomainOriginalNameServers = "original_name_servers"
	cisDomainType                = "type"
	cisDomainVerificationKey     = "verification_key"
	cisDomainCnameSuffix         = "cname_suffix"
	ibmCISDomain                 = "ibm_cis_domain"
)

func ResourceIBMCISDomain() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cis_domain",
					"cis_id"),
			},
			cisDomain: {
				Type:        schema.TypeString,
				Description: "CISzone - Domain",
				Required:    true,
			},
			cisDomainType: {
				Type:        schema.TypeString,
				Description: "CISzone - Domain Type",
				Default:     "full",
				Optional:    true,
				ValidateFunc: validate.InvokeValidator(ibmCISDomain,
					cisDomainType),
			},
			cisDomainPaused: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			cisDomainStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			cisDomainNameServers: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			cisDomainOriginalNameServers: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			cisDomainID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			cisDomainVerificationKey: {
				Type:     schema.TypeString,
				Computed: true,
			},
			cisDomainCnameSuffix: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create:   resourceCISdomainCreate,
		Read:     resourceCISdomainRead,
		Exists:   resourceCISdomainExists,
		Update:   resourceCISdomainUpdate,
		Delete:   resourceCISdomainDelete,
		Importer: &schema.ResourceImporter{},
	}
}

func resourceCISdomainCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisZonesV1ClientSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	cisClient.Crn = core.StringPtr(crn)
	zoneName := d.Get(cisDomain).(string)
	zoneType := d.Get(cisDomainType).(string)

	opt := cisClient.NewCreateZoneOptions()
	opt.SetName(zoneName)
	opt.SetType(zoneType)

	result, resp, err := cisClient.CreateZone(opt)
	if err != nil {
		log.Printf("CreateZones Failed %s", resp)
		return err
	}
	d.SetId(flex.ConvertCisToTfTwoVar(*result.Result.ID, crn))
	return resourceCISdomainRead(d, meta)
}

func resourceCISdomainRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisZonesV1ClientSession()
	if err != nil {
		return err
	}

	zoneID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	opt := cisClient.NewGetZoneOptions(zoneID)
	result, resp, err := cisClient.GetZone(opt)
	if err != nil {
		if isCISDomainDeleted(resp) {
			log.Printf("[WARN] Zone not found or already deleted, removing from state")
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone %v\n", resp)
		return err
	}
	d.Set(cisID, crn)
	d.Set(cisDomainID, result.Result.ID)
	d.Set(cisDomain, result.Result.Name)
	d.Set(cisDomainStatus, result.Result.Status)
	d.Set(cisDomainPaused, result.Result.Paused)
	d.Set(cisDomainNameServers, result.Result.NameServers)
	d.Set(cisDomainOriginalNameServers, result.Result.OriginalNameServers)
	d.Set(cisDomainType, result.Result.Type)

	if cisDomainType == "partial" {
		d.Set(cisDomainVerificationKey, result.Result.VerificationKey)
		d.Set(cisDomainCnameSuffix, result.Result.CnameSuffix)
	}

	return nil
}
func resourceCISdomainExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cisClient, err := meta.(conns.ClientSession).CisZonesV1ClientSession()
	if err != nil {
		return false, err
	}

	zoneID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	log.Println("resource exist :", d.Id())
	if err != nil {
		return false, err
	}
	log.Println("resource exist :", d.Id())
	cisClient.Crn = core.StringPtr(crn)
	opt := cisClient.NewGetZoneOptions(zoneID)
	_, resp, err := cisClient.GetZone(opt)
	if err != nil {
		if isCISDomainDeleted(resp) {
			log.Printf("[WARN] zone is not found or already deleted")
			return false, nil
		}
		log.Printf("[WARN] Error getting zone %v\n", resp)
		return false, err
	}
	return true, nil
}

func resourceCISdomainUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCISdomainRead(d, meta)
}

func resourceCISdomainDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisZonesV1ClientSession()
	if err != nil {
		return err
	}

	zoneID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	log.Println("resource delete :", d.Id())

	if err != nil {
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	opt := cisClient.NewGetZoneOptions(zoneID)
	_, resp, err := cisClient.GetZone(opt)
	if err != nil {
		if isCISDomainDeleted(resp) {
			log.Printf("[WARN] Zone already deleted, removing from state")
			return nil
		}
		log.Printf("[WARN] Error getting zone %v\n", resp)
		return err
	}
	delOpt := cisClient.NewDeleteZoneOptions(zoneID)
	_, resp, err = cisClient.DeleteZone(delOpt)
	if err != nil {
		// Zone already gone or deletion already in progress - treat as success so
		// Terraform removes it from state without blocking schematics workspace runs.
		if isCISDomainDeleted(resp) {
			log.Printf("[WARN] Zone already deleted or deletion in progress, removing from state")
			return nil
		}
		log.Printf("[ERR] Error deleting zone %v\n", resp)
		return err
	}
	return nil
}

func isCISDomainDeleted(resp *core.DetailedResponse) bool {
	if resp == nil {
		return false
	}
	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusGone {
		return true
	}
	// CIS API returns 400 with error code 1001 ("Invalid zone identifier") when a zone no longer exists.
	if resp.StatusCode == http.StatusBadRequest {
		if m, ok := resp.GetResultAsMap(); ok {
			if errs, ok := m["errors"].([]interface{}); ok && len(errs) > 0 {
				if e, ok := errs[0].(map[string]interface{}); ok {
					if code, ok := e["code"].(float64); ok && code == 1001 {
						return true
					}
				}
			}
		}
	}
	return false
}

func ResourceIBMCISDomainValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 cisDomainType,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "full, partial"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	ibmCISDomainResourceValidator := validate.ResourceValidator{
		ResourceName: ibmCISDomain,
		Schema:       validateSchema}
	return &ibmCISDomainResourceValidator
}
