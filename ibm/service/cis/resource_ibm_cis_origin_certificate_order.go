// Copyright IBM Corp. 2017, 2021, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISOriginCertificateOrder     = "ibm_cis_origin_certificate_order"
	cisOriginCertificate             = "certificate"
	cisOriginCertificateID           = "certificate_id"
	cisOriginCertificateHosts        = "hostnames"
	cisOriginCertificateType         = "request_type"
	cisOriginCertificateValidityDays = "requested_validity"
	cisOriginCertificateCSR          = "csr"
	cisOriginCertificateExpiresOn    = "expires_on"
	cisOriginCertificatePrivateKey   = "private_key"
)

func ResourceIBMCISOriginCertificateOrder() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISOriginCertificateCreate,
		Update:   ResourceIBMCISOriginCertificateRead,
		Read:     ResourceIBMCISOriginCertificateRead,
		Delete:   ResourceIBMCISOriginCertificateDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS object id or CRN",
				Required:    true,
				ValidateFunc: validate.InvokeValidator(ibmCISOriginCertificateOrder,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisOriginCertificateID: {
				Type:        schema.TypeString,
				Description: "certificate id",
				Computed:    true,
			},
			cisOriginCertificateType: {
				Type:        schema.TypeString,
				Description: "certificate type",
				Required:    true,
			},
			cisOriginCertificateHosts: {
				Type:        schema.TypeList,
				Description: "Hosts which certificate need to be ordered",
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			cisOriginCertificateValidityDays: {
				Type:        schema.TypeInt,
				Description: "validity days",
				Required:    true,
			},
			cisOriginCertificateCSR: {
				Type:        schema.TypeString,
				Description: "validity days",
				Required:    true,
			},
			cisOriginCertificatePrivateKey: {
				Type:        schema.TypeString,
				Description: "certificate id",
				Computed:    true,
			},
			cisOriginCertificate: {
				Type:        schema.TypeString,
				Description: "certificate id",
				Computed:    true,
			},
			cisOriginCertificateExpiresOn: {
				Type:        schema.TypeString,
				Description: "certificate id",
				Computed:    true,
			},
		},
	}
}

func ResourceIBMCISOriginCertificateOrderValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	cisCertificateOrderValidator := validate.ResourceValidator{
		ResourceName: ibmCISOriginCertificateOrder,
		Schema:       validateSchema}
	return &cisCertificateOrderValidator
}

func ResourceIBMCISOriginCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisSSLClientSession()
	if err != nil {
		return err
	}
	crn := d.Get(cisID).(string)
	zoneID := d.Get(cisDomainID).(string)
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	certType := d.Get(cisOriginCertificateType).(string)
	hosts := d.Get(cisOriginCertificateHosts)
	hostsList := flex.ExpandStringList(hosts.([]interface{}))
	validityDays := int64(d.Get(cisOriginCertificateValidityDays).(int))
	csr := d.Get(cisOriginCertificateCSR).(string)

	opt := cisClient.NewCreateOriginCertificateOptions(crn, zoneID)
	opt.SetHostnames(hostsList)
	opt.SetCsr(csr)
	opt.SetRequestType(certType)
	opt.SetRequestedValidity(validityDays)

	result, resp, err := cisClient.CreateOriginCertificate(opt)
	if err != nil {
		log.Printf("Origin Certificate  order failed: %v", resp)
		return err
	}

	d.SetId(flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn))
	return ResourceIBMCISOriginCertificateRead(d, meta)
}

func ResourceIBMCISOriginCertificateRead(d *schema.ResourceData, meta interface{}) error {

	cisClient, err := meta.(conns.ClientSession).CisSSLClientSession()
	if err != nil {
		return err
	}
	certificateID, zoneID, crn, err := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Println("Error in reading certificate id")
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)
	opt := cisClient.NewGetOriginCertificateOptions(crn, zoneID, certificateID)
	result, resp, err := cisClient.GetOriginCertificate(opt)
	if err != nil {
		log.Printf("Certificate read failed: %v", resp)
		return err
	}
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisOriginCertificateID, result.Result.ID)
	d.Set(cisOriginCertificate, result.Result.Certificate)
	d.Set(cisOriginCertificateHosts, flex.FlattenStringList(result.Result.Hostnames))
	d.Set(cisOriginCertificateExpiresOn, result.Result.ExpiresOn)
	d.Set(cisOriginCertificateType, result.Result.ExpiresOn)
	d.Set(cisOriginCertificateValidityDays, result.Result.RequestedValidity)
	d.Set(cisOriginCertificateCSR, result.Result.Csr)
	d.Set(cisOriginCertificatePrivateKey, result.Result.PrivateKey)
	return nil
}

func ResourceIBMCISOriginCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisSSLClientSession()
	if err != nil {
		return err
	}
	certificateID, zoneID, crn, err := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Println("Error in reading certificate id")
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)
	opt := cisClient.NewRevokeOriginCertificateOptions(crn, zoneID, certificateID)
	resp, _, err := cisClient.RevokeOriginCertificate(opt)
	if err != nil {
		log.Printf("Origin Certificate delete failed: %v", resp)
		return err
	}

	// _, err = waitForCISCertificateOrderDelete(d, meta)
	// if err != nil {
	// 	return err
	// }

	return nil
}

/*
func ResourceIBMCISCertificateOrderExist(d *schema.ResourceData, meta interface{}) (bool, error) {
	cisClient, err := meta.(conns.ClientSession).CisSSLClientSession()
	if err != nil {
		return false, err
	}
	certificateID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Println("Error in reading certificate id")
		return false, err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)
	opt := cisClient.NewGetCustomCertificateOptions(certificateID)
	_, response, err := cisClient.GetCustomCertificate(opt)
	if err != nil {
		if response != nil && response.StatusCode == 400 {
			log.Printf("Certificate is not found")
			return false, nil
		}
		log.Printf("Get Certificate failed: %v", response)
		return false, err
	}
	return true, nil
}
*/
/*
func waitForCISCertificateOrderDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	cisClient, err := meta.(conns.ClientSession).CisSSLClientSession()
	if err != nil {
		return nil, err
	}
	certificateID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Println("Error in reading certificate id")
		return nil, err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)
	opt := cisClient.NewGetCustomCertificateOptions(certificateID)
	stateConf := &resource.StateChangeConf{
		Pending: []string{cisCertificateOrderDeletePending},
		Target:  []string{cisCertificateOrderDeleted},
		Refresh: func() (interface{}, string, error) {
			_, detail, err := cisClient.GetCustomCertificate(opt)
			if err != nil {
				if detail != nil && detail.StatusCode == 400 {
					return detail, cisCertificateOrderDeleted, nil
				}
				return nil, "", err
			}
			return detail, cisCertificateOrderDeletePending, nil
		},
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        10 * time.Second,
		MinTimeout:   10 * time.Second,
		PollInterval: 10 * time.Second,
	}

	return stateConf.WaitForState()
}
*/
