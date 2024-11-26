// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.95.2-120e65bc-20240924-152329
 */

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/mqcloud"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmMqcloudTruststoreCertificateDataSourceBasic(t *testing.T) {
	trustStoreCertificateDetailsServiceInstanceGuid := acc.MqcloudDeploymentID
	trustStoreCertificateDetailsQueueManagerID := acc.MqcloudQueueManagerID
	trustStoreCertificateDetailsLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	trustStoreCertificateDetailsCertificateFile := acc.MqcloudTSCertFilePath

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheckMqcloud(t)
		},
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudTruststoreCertificateDataSourceConfigBasic(trustStoreCertificateDetailsServiceInstanceGuid, trustStoreCertificateDetailsQueueManagerID, trustStoreCertificateDetailsLabel, trustStoreCertificateDetailsCertificateFile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance", "queue_manager_id"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudTruststoreCertificateDataSourceConfigBasic(trustStoreCertificateDetailsServiceInstanceGuid string, trustStoreCertificateDetailsQueueManagerID string, trustStoreCertificateDetailsLabel string, trustStoreCertificateDetailsCertificateFile string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
			service_instance_guid = "%s"
			queue_manager_id = "%s"
			label = "%s"
			certificate_file = filebase64("%s")
		}

		data "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
			service_instance_guid = ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance.service_instance_guid
			queue_manager_id = ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance.queue_manager_id
			label = ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance.label
		}
	`, trustStoreCertificateDetailsServiceInstanceGuid, trustStoreCertificateDetailsQueueManagerID, trustStoreCertificateDetailsLabel, trustStoreCertificateDetailsCertificateFile)
}

func TestDataSourceIbmMqcloudTruststoreCertificateTrustStoreCertificateDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["label"] = "testString"
		model["certificate_type"] = "trust_store"
		model["fingerprint_sha256"] = "testString"
		model["subject_dn"] = "testString"
		model["subject_cn"] = "testString"
		model["issuer_dn"] = "testString"
		model["issuer_cn"] = "testString"
		model["issued"] = "2019-01-01T12:00:00.000Z"
		model["expiry"] = "2019-01-01T12:00:00.000Z"
		model["trusted"] = true
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(mqcloudv1.TrustStoreCertificateDetails)
	model.ID = core.StringPtr("testString")
	model.Label = core.StringPtr("testString")
	model.CertificateType = core.StringPtr("trust_store")
	model.FingerprintSha256 = core.StringPtr("testString")
	model.SubjectDn = core.StringPtr("testString")
	model.SubjectCn = core.StringPtr("testString")
	model.IssuerDn = core.StringPtr("testString")
	model.IssuerCn = core.StringPtr("testString")
	model.Issued = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Expiry = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Trusted = core.BoolPtr(true)
	model.Href = core.StringPtr("testString")

	result, err := mqcloud.DataSourceIbmMqcloudTruststoreCertificateTrustStoreCertificateDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
