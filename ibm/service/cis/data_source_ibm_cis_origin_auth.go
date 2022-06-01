// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis


import (
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisOriginAuthCerts      = "certs"
	cisOriginAuthHostStatus = "host_status"
	CisOriginAuthHostName   = "host_name"
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
				Default:     "no-host",
			},
			cisOriginAuthHostStatus: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Host Level Setting",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"cert_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"cert_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"issuer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"signature": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"serial_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"certificate": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"cert_uploaded_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"cert_updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
						"expires_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CIS origin auth enabled or disabled",
						},
					},
				},
			},
			cisOriginAuthCerts: {
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

	// Get Hostname-Level Origin Pull Settings
	statusLevels := make([]map[string]interface{}, 0)
	statusLevel := map[string]interface{}{}

	if d.Get(CisOriginAuthHostName).(string) != "no-host" {
		getHostnameOption := sess.NewGetHostnameOriginPullSettingsOptions(d.Get(CisOriginAuthHostName).(string))
		getHostnameOption.SetHostname(d.Get(CisOriginAuthHostName).(string))
		hResult, _, hErr := sess.GetHostnameOriginPullSettings(getHostnameOption)
		if hErr != nil || hResult == nil {
			return fmt.Errorf("[ERROR] Error getting Hostname-Level Origin Pull Settings : %s", hErr)
		}

		statusLevel["hostname"] = *hResult.Result.Hostname
		statusLevel["cert_id"] = *hResult.Result.CertID
		statusLevel["enabled"] = *hResult.Result.Enabled
		statusLevel["status"] = *hResult.Result.Status
		statusLevel["created_at"] = *hResult.Result.CreatedAt
		statusLevel["updated_at"] = *hResult.Result.UpdatedAt
		statusLevel["cert_status"] = *hResult.Result.CertStatus
		statusLevel["issuer"] = *hResult.Result.Issuer
		statusLevel["signature"] = *hResult.Result.Signature
		statusLevel["serial_number"] = *hResult.Result.SerialNumber
		statusLevel["certificate"] = *hResult.Result.Certificate
		statusLevel["cert_uploaded_on"] = *hResult.Result.CertUploadedOn
		statusLevel["cert_updated_at"] = *hResult.Result.CertUpdatedAt
		statusLevel["expires_on"] = *hResult.Result.ExpiresOn

	}

	statusLevels = append(statusLevels, statusLevel)

	// Get Zone Origin Pull Settings
	listOpt := sess.NewListZoneOriginPullCertificatesOptions()
	result, response, err := sess.ListZoneOriginPullCertificates(listOpt)

	log.Printf("Response code : %v : ", response)

	if err != nil || result == nil {
		return fmt.Errorf("[ERROR] Error getting origin auth certificate: %s", err)
	}

	certLists := make([]map[string]interface{}, 0)
	certList := map[string]interface{}{}
	for _, instance := range result.Result {

		certList["id"] = *instance.ID
		certList["certificate"] = *instance.Certificate
		certList["issuer"] = *instance.Issuer
		certList["signature"] = *instance.Signature
		certList["status"] = *instance.Status
		certList["expires_on"] = *instance.ExpiresOn
		certList["uploaded_on"] = *instance.UploadedOn

	}
	certLists = append(certLists, certList)
	d.SetId(DataSourceIBMCISOriginAuthPullID(d))
	d.Set(cisOriginAuthHostStatus, statusLevels)
	d.Set(cisOriginAuthCerts, certLists)

	return nil
}

func DataSourceIBMCISOriginAuthPullID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
