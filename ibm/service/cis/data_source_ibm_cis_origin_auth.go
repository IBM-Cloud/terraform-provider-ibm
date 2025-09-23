// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	CisOriginAuthHostName    = "hostname"
	CisOriginAuthRequestType = "request_type"
)

func DataSourceIBMCISOriginAuthPull() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataIBMCISOriginAuthRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_cis_origin_auths",
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			CisOriginAuthHostName: {
				Type:        schema.TypeString,
				Description: "Associated CIS host name",
				Optional:    true,
				Default:     "no_host",
			},
			CisOriginAuthRequestType: {
				Type:        schema.TypeString,
				Description: "Associated CIS Request Type",
				Optional:    true,
				Default:     "zone_level",
			},
			"origin_pull_settings_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "CIS origin auth settings enabled or disabled",
			},
			"origin_pull_certs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Certficate list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate id",
						},

						"certificate": {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "CIS origin auth certificate detail",
						},
						"cert_issuer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate issue",
						},
						"cert_signature": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate signature",
						},
						"cert_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate active or not",
						},
						"cert_expires_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate expiry time",
						},
						"cert_uploaded_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate upldate time",
						},
						"cert_serial_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate Serial Number",
						},
					},
				},
			},
		},
	}

}
func DataSourceIBMCISOriginAuthPullValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	iBMCISOriginAuthValidator := validate.ResourceValidator{
		ResourceName: "ibm_cis_origin_auths",
		Schema:       validateSchema}
	return &iBMCISOriginAuthValidator
}

func dataIBMCISOriginAuthRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("dataIBMCISOriginAuthRead CisOrigAuthSession initialization failed: %s", err.Error()),
			"ibm_cis_origin_auth", "read")
		return tfErr.GetDiag()
	}

	crn := d.Get(cisID).(string)
	zoneID, _, _, _ := flex.ConvertTfToCisThreeVar(d.Get(cisDomainID).(string))
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	request_type := d.Get(CisOriginAuthRequestType).(string)

	if request_type == "zone_level" {

		// Get Zone Origin Pull Settings
		zoneSettingsOpt := sess.NewGetZoneOriginPullSettingsOptions()
		zoneSettingsResult, zoneSettingsResponse, zoneSettingsErr := sess.GetZoneOriginPullSettings(zoneSettingsOpt)

		if zoneSettingsErr != nil || zoneSettingsResponse == nil {
			tfErr := flex.TerraformErrorf(zoneSettingsErr,
				fmt.Sprintf("dataIBMCISOriginAuthRead GetZoneOriginPullSettings failed: %s \nResponse: %s", zoneSettingsErr.Error(), zoneSettingsResponse),
				"ibm_cis_origin_auth", "read")
			return tfErr.GetDiag()
		}

		zoneSettings := zoneSettingsResult.Result.Enabled
		d.Set("origin_pull_settings_enabled", zoneSettings)

		// Get Zone Origin Pull Certificate List
		zoneCertListOpt := sess.NewListZoneOriginPullCertificatesOptions()
		zoneCertListResult, zoneCertListResponse, zoneCertListErr := sess.ListZoneOriginPullCertificates(zoneCertListOpt)

		if zoneCertListErr != nil || zoneCertListResponse == nil {
			tfErr := flex.TerraformErrorf(zoneCertListErr,
				fmt.Sprintf("dataIBMCISOriginAuthRead ListZoneOriginPullCertificates failed: %s \nResponse: %s", zoneCertListErr.Error(), zoneCertListResponse),
				"ibm_cis_origin_auth", "read")
			return tfErr.GetDiag()
		}

		zoneCertLists := make([]map[string]interface{}, 0)
		zoneCertList := map[string]interface{}{}

		for _, certObj := range zoneCertListResult.Result {

			zoneCertList["cert_id"] = *certObj.ID
			zoneCertList["certificate"] = *certObj.Certificate
			zoneCertList["cert_issuer"] = *certObj.Issuer
			zoneCertList["cert_signature"] = *certObj.Signature
			zoneCertList["cert_status"] = *certObj.Status
			zoneCertList["cert_expires_on"] = *certObj.ExpiresOn
			zoneCertList["cert_uploaded_on"] = *certObj.UploadedOn

		}
		zoneCertLists = append(zoneCertLists, zoneCertList)
		d.Set("origin_pull_certs", zoneCertLists)

	} else if request_type == "per_hostname" {

		// Get Hostname Origin Pull Settings
		hostname := d.Get(CisOriginAuthHostName).(string)
		hostnameSettingsOpt := sess.NewGetHostnameOriginPullSettingsOptions(hostname)
		hostnameSettingsOpt.SetHostname(hostname)

		hostnameSettingsResult, hostnameSettingsResponse, hostnameSettingsErr := sess.GetHostnameOriginPullSettings(hostnameSettingsOpt)

		if hostnameSettingsErr != nil || hostnameSettingsResponse == nil {
			tfErr := flex.TerraformErrorf(hostnameSettingsErr,
				fmt.Sprintf("dataIBMCISOriginAuthRead GetHostnameOriginPullSettings failed: %s \nResponse: %s", hostnameSettingsErr.Error(), hostnameSettingsResponse),
				"ibm_cis_origin_auth", "read")
			return tfErr.GetDiag()
		}

		hostnameSettings := hostnameSettingsResult.Result.Enabled
		d.Set("origin_pull_settings_enabled", hostnameSettings)
	}

	d.SetId(DataSourceIBMCISOriginAuthPullID(d))
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(CisOriginAuthRequestType, request_type)
	return nil
}

func DataSourceIBMCISOriginAuthPullID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
