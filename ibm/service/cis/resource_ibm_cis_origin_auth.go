// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package cis

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisOriginAuthID          = "cis_origin_auth_id"
	cisOriginAuthLevel       = "auth_level"
	cisOriginAuthHost        = "host_name"
	cisOriginAuthEnable      = "auth_setting"
	cisOriginAuthCertContent = "cert"
	cisOriginAuthCertKey     = "key"
	cisOriginAuthCertId      = "cert_id"
	cisOriginAuthZoneID      = "zone_id"
	cisOriginAuthInstance    = "cis_instance"
)

func ResourceIBMCISOriginAuthPull() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISOriginAuthPullCreate,
		Read:     ResourceIBMCISOriginAuthPullRead,
		Update:   ResourceIBMCISOriginAuthPullUpdate,
		Delete:   ResourceIBMCISOriginAuthPullDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Optional:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisOriginAuthLevel: {
				Type:        schema.TypeString,
				Description: "Authentication level either zone or host",
				Optional:    true,
				Default:     "zone",
			},
			cisOriginAuthHost: {
				Type:        schema.TypeString,
				Description: "Host Name",
				Optional:    true,
			},
			cisOriginAuthEnable: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			cisOriginAuthZoneID: {
				Type:        schema.TypeString,
				Description: "Enable or disable origin auth",
				Optional:    true,
			},
			cisOriginAuthCertContent: {
				Type:        schema.TypeString,
				Description: "Enable or disable origin auth",
				Optional:    true,
			},
			cisOriginAuthCertKey: {
				Type:        schema.TypeString,
				Description: "Enable or disable origin auth",
				Optional:    true,
			},
			cisOriginAuthID: {
				Type:        schema.TypeString,
				Description: "Associated CIS auth pull job id",
				Computed:    true,
			},
		},
	}

}

func ResourceIBMCISOriginAuthPullCreate(d *schema.ResourceData, meta interface{}) error {
	var certValue string
	var keyValue string
	origin_type := "zone"
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
	}

	crn := d.Get(cisID).(string)
	zoneID, _, _ := flex.ConvertTftoCisTwoVar(d.Get(cisDomainID).(string))

	sess.ZoneIdentifier = core.StringPtr(zoneID)

	if authContent, ok := d.GetOk(cisOriginAuthLevel); ok {
		origin_type = authContent.(string)
	}

	if certContent, ok := d.GetOk(cisOriginAuthCertContent); ok {
		certValue = certContent.(string)

	}

	if certKey, ok := d.GetOk(cisOriginAuthCertContent); ok {
		keyValue = certKey.(string)

	}

	if origin_type == "zone" {
		log.Printf(" --> Uploading Zone level certificate %v<-- ", origin_type)
		options := sess.NewUploadZoneOriginPullCertificateOptions()
		options.SetCertificate(certValue)
		options.SetPrivateKey(keyValue)

		result, resp, opErr := sess.UploadZoneOriginPullCertificate(options)
		if opErr != nil {
			log.Printf("Uploading Zone level certificate pull failed %s", resp)
			return opErr
		}
		d.SetId(flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn))

	} else {
		log.Printf(" --> Uploading Host level certificate <-- ")
		options := sess.NewUploadHostnameOriginPullCertificateOptions()
		options.SetCertificate(certValue)
		options.SetPrivateKey(keyValue)
		result, resp, opErr := sess.UploadHostnameOriginPullCertificate(options)
		if opErr != nil {
			log.Printf("Creating Host level certificate pull failed %s", resp)
			return opErr
		}
		d.SetId(flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn))

	}

	return ResourceIBMCISOriginAuthPullRead(d, meta)
}

func ResourceIBMCISOriginAuthPullRead(d *schema.ResourceData, meta interface{}) error {
	origin_type := "zone"
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
	}

	authPullID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	if authContent, ok := d.GetOk(cisOriginAuthLevel); ok {
		origin_type = authContent.(string)
	}

	if origin_type == "zone" {

		getOptions := sess.NewGetZoneOriginPullCertificateOptions(authPullID)
		getOptions.SetCertIdentifier(authPullID)

		result, response, err := sess.GetZoneOriginPullCertificate(getOptions)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("[ERROR] Error while reading the zone origin auth pull %s:%s:%v", err, response, origin_type)
		}
		d.Set(cisOriginAuthID, *result.Result.ID)
		d.Set(cisOriginAuthCertContent, *result.Result.Certificate)

	} else {
		getOptions := sess.NewGetHostnameOriginPullCertificateOptions(authPullID)
		getOptions.SetCertIdentifier(authPullID)

		result, response, err := sess.GetHostnameOriginPullCertificate(getOptions)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("[ERROR] Error while reading the host origin auth pull %s:%s", err, response)
		}
		d.Set(cisOriginAuthID, *result.Result.ID)
		d.Set(cisOriginAuthCertContent, *result.Result.Certificate)

	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)

	return nil
}

func ResourceIBMCISOriginAuthPullUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
	}
	_, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	if d.HasChange(cisOriginAuthEnable) {
		updateOption := sess.NewSetZoneOriginPullSettingsOptions()
		updateOption.SetEnabled(d.Get(cisOriginAuthEnable).(bool))
		result, response, err := sess.SetZoneOriginPullSettings(updateOption)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("[ERROR] Error while updating the zone origin auth pull setting %s:%s", err, response)
		}
		log.Printf("Origin auth pull update result : %v", result)

	}
	return ResourceIBMCISOriginAuthPullRead(d, meta)

}

func ResourceIBMCISOriginAuthPullDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
	}

	authPullID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	delOpt := sess.NewDeleteZoneOriginPullCertificateOptions(authPullID)
	result, resp, err := sess.DeleteZoneOriginPullCertificate(delOpt)

	if err != nil {
		log.Printf("Error deleting origin auth pull: %s", resp)
		return err
	}

	log.Printf("Origin auth pull ID: %s", *result.Result.ID)
	d.SetId("")
	return nil

}
