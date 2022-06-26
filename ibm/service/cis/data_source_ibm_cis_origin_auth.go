// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisOriginAuthHostStatus  = "host_status"
	CisOriginAuthHostName    = "hostname"
	CisOriginAuthRequestType = "request_type"
)

func DataSourceIBMCISOriginAuthPull() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISOriginAuthRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
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
			"origin_pull_certs_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Certficate list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate id",
						},

						"certificate": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate detail",
						},
						"issuer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate issue",
						},
						"signature": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate signature",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate active or not",
						},
						"expires_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certificate expiry time",
						},
						"uploaded_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certifaate upldate time",
						},
						"serial_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth certifaate upldate time",
						},
					},
				},
			},
		},
	}

}

func dataIBMCISOriginAuthRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisOrigAuthSession()
	if err != nil {
		return err
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
			return fmt.Errorf("[ERROR] Error Getting Zone Level Origin Pull Settings: %s", zoneSettingsErr)
		}

		zoneSettings := zoneSettingsResult.Result.Enabled
		d.Set("origin_pull_settings_enabled", zoneSettings)

		// Get Zone Origin Pull Certificate List
		zoneCertListOpt := sess.NewListZoneOriginPullCertificatesOptions()
		zoneCertListResult, zoneCertListResponse, zoneCertListErr := sess.ListZoneOriginPullCertificates(zoneCertListOpt)

		if zoneCertListErr != nil || zoneCertListResponse == nil {
			return fmt.Errorf("[ERROR] Error Getting Zone Level Origin Pull Settings: %s", zoneCertListErr)
		}

		zoneCertLists := make([]map[string]interface{}, 0)
		zoneCertList := map[string]interface{}{}

		for _, certObj := range zoneCertListResult.Result {

			zoneCertList["id"] = *certObj.ID
			zoneCertList["certificate"] = *certObj.Certificate
			zoneCertList["issuer"] = *certObj.Issuer
			zoneCertList["signature"] = *certObj.Signature
			zoneCertList["status"] = *certObj.Status
			zoneCertList["expires_on"] = *certObj.ExpiresOn
			zoneCertList["uploaded_on"] = *certObj.UploadedOn

		}
		zoneCertLists = append(zoneCertLists, zoneCertList)
		d.Set("origin_pull_certs_list", zoneCertLists)

	} else if request_type == "per_hostname" {

		// Get Hostname Origin Pull Settings
		hostname := d.Get(CisOriginAuthHostName).(string)
		hostnameSettingsOpt := sess.NewGetHostnameOriginPullSettingsOptions(hostname)
		hostnameSettingsOpt.SetHostname(hostname)

		hostnameSettingsResult, hostnameSettingsResponse, hostnameSettingsErr := sess.GetHostnameOriginPullSettings(hostnameSettingsOpt)

		if hostnameSettingsErr != nil || hostnameSettingsResponse == nil {
			return fmt.Errorf("[ERROR] Error Getting Zone Level Origin Pull Settings: %s", hostnameSettingsErr)
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
