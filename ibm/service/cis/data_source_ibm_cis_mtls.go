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
	cisMtlsOutput                  = "mtls_certificates"
	cisMtlsCertID                  = "cert_id"
	cisMtlsCertName                = "cert_name"
	cisMtlsCertFingerprint         = "cert_fingerprint"
	cisMtlsCertAssociatedHostnames = "cert_associated_hostnames"
	cisMtlsCertCreatedAt           = "cert_created_at"
	cisMtlsCertUpdatedAt           = "cert_updated_at"
	cisMtlsCertExpiresOn           = "cert_expires_on"
)

func DataSourceIBMCISMtls() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISMtlsRead,
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
			cisMtlsOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Container for response information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisMtlsCertID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate ID",
						},
						cisMtlsCertName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate name",
						},
						cisMtlsCertFingerprint: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate Fingerprint",
						},
						cisMtlsCertAssociatedHostnames: {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "Certificate Associated Hostnames",
						},
						cisMtlsCertCreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate Created At",
						},
						cisMtlsCertUpdatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate Updated At",
						},
						cisMtlsCertExpiresOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate Expires On",
						},
					},
				},
			},
		},
	}
}

func dataIBMCISMtlsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()

	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneID, _, _ := flex.ConvertTftoCisTwoVar(d.Get(cisDomainID).(string))

	opt := sess.NewListAccessCertificatesOptions(zoneID)

	result, resp, err := sess.ListAccessCertificates(opt)
	if err != nil {
		log.Printf("[WARN] List all certificates failed: %v\n", resp)
		return err
	}
	mtlsCertLists := make([]map[string]interface{}, 0)
	for _, certObj := range result.Result {
		mtlsCertList := map[string]interface{}{}
		mtlsCertList[cisMtlsCertID] = *certObj.ID
		mtlsCertList[cisMtlsCertName] = *certObj.Name
		mtlsCertList[cisMtlsCertFingerprint] = *certObj.Fingerprint
		mtlsCertList[cisMtlsCertAssociatedHostnames] = certObj.AssociatedHostnames
		mtlsCertList[cisMtlsCertCreatedAt] = *certObj.CreatedAt
		mtlsCertList[cisMtlsCertUpdatedAt] = *certObj.UpdatedAt
		mtlsCertList[cisMtlsCertExpiresOn] = *certObj.ExpiresOn

		mtlsCertLists = append(mtlsCertLists, mtlsCertList)

	}
	d.SetId(dataSourceCISMtlsCheckID(d))

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisMtlsOutput, mtlsCertLists)

	return nil
}
func dataSourceCISMtlsCheckID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
