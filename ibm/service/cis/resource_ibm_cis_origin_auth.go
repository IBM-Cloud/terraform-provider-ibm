// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package cis

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/authenticatedoriginpullapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisOriginAuthID          = "auth_id"
	cisOriginAuthHost        = "hostname"
	cisOriginAuthEnable      = "enabled"
	cisOriginAuthCertContent = "certificate"
	cisOriginAuthCertKey     = "private_key"
	cisOriginAuthCertId      = "cert_id"
	CisOriginAuthStatus      = "status"
	cisOriginAuthExpiresOn   = "expires_on"
	cisOriginAuthUploadedOn  = "uploaded_on"
)

func ResourceIBMCISOriginAuthPull() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCISOriginAuthPullCreate,
		ReadContext:   resourceIBMCISOriginAuthPullRead,
		UpdateContext: resourceIBMCISOriginAuthPullUpdate,
		DeleteContext: resourceIBMCISOriginAuthPullDelete,
		Importer:      &schema.ResourceImporter{},
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
			cisOriginAuthHost: {
				Type:        schema.TypeString,
				Description: "Host name needed for host level authentication",
				Optional:    true,
			},
			cisOriginAuthEnable: {
				Type:        schema.TypeBool,
				Description: "Enabel-disable origin auth for a zone or host",
				Optional:    true,
				Default:     true,
			},
			cisOriginAuthCertContent: {
				Type:        schema.TypeString,
				Description: "Certificate content which needs to be uploaded",
				Required:    true,
			},
			cisOriginAuthCertKey: {
				Type:        schema.TypeString,
				Description: "Private key content which needs to be uploaded",
				Required:    true,
			},
			CisOriginAuthStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Authentication status whether active or not",
			},
			cisOriginAuthCertId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate ID which is uploaded",
			},
			cisOriginAuthExpiresOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate expires on",
			},
			cisOriginAuthUploadedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate uploaded on",
			},
			cisOriginAuthID: {
				Type:        schema.TypeString,
				Description: "Associated CIS auth pull job id",
				Computed:    true,
			},
		},
	}

}

func resourceIBMCISOriginAuthPullCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var cert_val string
	var key_val string
	var zone_config bool

	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %v", err))
	}

	crn := d.Get(cisID).(string)
	zoneID, _, _ := flex.ConvertTftoCisTwoVar(d.Get(cisDomainID).(string))

	sess.ZoneIdentifier = core.StringPtr(zoneID)
	sess.Crn = core.StringPtr(crn)

	if cert_content, ok := d.GetOk(cisOriginAuthCertContent); ok {
		cert_val = cert_content.(string)

	}

	if cert_key, ok := d.GetOk(cisOriginAuthCertKey); ok {
		key_val = cert_key.(string)

	}
	zone_config = true
	if _, ok := d.GetOk(cisOriginAuthHost); ok {
		zone_config = false
	}

	// Check host level certificate creation or zone level
	if zone_config {
		options := sess.NewUploadZoneOriginPullCertificateOptions()
		options.SetCertificate(cert_val)
		options.SetPrivateKey(key_val)

		result, resp, opErr := sess.UploadZoneOriginPullCertificate(options)
		if opErr != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while uploading certificate zone level %v", resp))
		}
		d.SetId(flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn))

	} else {
		options := sess.NewUploadHostnameOriginPullCertificateOptions()
		options.SetCertificate(cert_val)
		options.SetPrivateKey(key_val)
		result, resp, opErr := sess.UploadHostnameOriginPullCertificate(options)
		if opErr != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while uploading certificate host level %v", resp))
		}
		d.SetId(flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn))

	}

	return resourceIBMCISOriginAuthPullRead(context, d, meta)
}

func resourceIBMCISOriginAuthPullRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var zone_config bool
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %v", err))
	}

	certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	zone_config = true
	if _, ok := d.GetOk(cisOriginAuthHost); ok {
		zone_config = false
	}

	if zone_config {
		getOptions := sess.NewGetZoneOriginPullCertificateOptions(certID)
		getOptions.SetCertIdentifier(certID)

		result, response, err := sess.GetZoneOriginPullCertificate(getOptions)

		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while getting detail of zone origin auth pull %v:%v", err, response))
		}
		d.Set(cisOriginAuthID, *result.Result.ID)
		d.Set(cisOriginAuthCertContent, *result.Result.Certificate)
		d.Set(CisOriginAuthStatus, *result.Result.Status)
		d.Set(cisOriginAuthExpiresOn, *result.Result.ExpiresOn)
		d.Set(cisOriginAuthUploadedOn, *result.Result.UploadedOn)
		d.Set(cisOriginAuthCertId, *result.Result.ID)

	} else {
		getOptions := sess.NewGetHostnameOriginPullCertificateOptions(certID)
		getOptions.SetCertIdentifier(certID)

		result, response, err := sess.GetHostnameOriginPullCertificate(getOptions)

		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while getting detail of host origin auth pull %v:%v", err, response))
		}
		d.Set(cisOriginAuthID, *result.Result.ID)
		d.Set(cisOriginAuthCertContent, *result.Result.Certificate)
		d.Set(CisOriginAuthStatus, *result.Result.Status)
		d.Set(cisOriginAuthExpiresOn, *result.Result.ExpiresOn)
		d.Set(cisOriginAuthUploadedOn, *result.Result.UploadedOn)
		d.Set(cisOriginAuthCertId, *result.Result.ID)
	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)

	return nil
}

func resourceIBMCISOriginAuthPullUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var zone_config bool
	var host_name string
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %v", err))
	}
	certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	zone_config = true
	if host_val, ok := d.GetOk(cisOriginAuthHost); ok {
		zone_config = false
		host_name = host_val.(string)

	}

	if zone_config {

		if d.HasChange(cisOriginAuthEnable) {
			updateOption := sess.NewSetZoneOriginPullSettingsOptions()
			updateOption.SetEnabled(d.Get(cisOriginAuthEnable).(bool))
			_, response, err := sess.SetZoneOriginPullSettings(updateOption)

			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error while updaing the zone origin auth pull setting %v:%v", err, response))
			}

		}

	} else {

		if d.HasChange(cisOriginAuthEnable) {

			model := &authenticatedoriginpullapiv1.HostnameOriginPullSettings{
				Hostname: core.StringPtr(host_name),
				CertID:   core.StringPtr(certID),
				Enabled:  core.BoolPtr(d.Get(cisOriginAuthEnable).(bool)),
			}
			setOption := sess.NewSetHostnameOriginPullSettingsOptions()
			setOption.SetConfig([]authenticatedoriginpullapiv1.HostnameOriginPullSettings{*model})
			_, setResp, setErr := sess.SetHostnameOriginPullSettings(setOption)
			if setErr != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error while updaing the host origin auth pull setting %v:%v", setErr, setResp))
			}

		}

	}
	return resourceIBMCISOriginAuthPullRead(context, d, meta)

}

func resourceIBMCISOriginAuthPullDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var zone_config bool
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %v", err))
	}

	certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	zone_config = true
	if _, ok := d.GetOk(cisOriginAuthHost); ok {
		zone_config = false
	}

	if zone_config {
		delOpt := sess.NewDeleteZoneOriginPullCertificateOptions(certID)
		_, resp, err := sess.DeleteZoneOriginPullCertificate(delOpt)

		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while deleting the certificate zone level %v: %v", certID, resp))
		}

	} else {
		delOpt := sess.NewDeleteHostnameOriginPullCertificateOptions(certID)
		_, resp, err := sess.DeleteHostnameOriginPullCertificate(delOpt)

		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while deleting the certificate host level %v: %v", certID, resp))
		}

	}
	d.SetId("")
	return nil

}
