// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmMqcloudKeystoreCertificateDataSourceBasic(t *testing.T) {
	t.Parallel()
	keyStoreCertificateDetailsServiceInstanceGuid := acc.MqcloudInstanceID
	keyStoreCertificateDetailsQueueManagerID := acc.MqcloudQueueManagerID
	keyStoreCertificateDetailsLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	keyStoreCertificateDetailsCertificateFile := acc.MqcloudKSCertFilePath

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudKeystoreCertificateDataSourceConfigBasic(keyStoreCertificateDetailsServiceInstanceGuid, keyStoreCertificateDetailsQueueManagerID, keyStoreCertificateDetailsLabel, keyStoreCertificateDetailsCertificateFile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "queue_manager_id"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudKeystoreCertificateDataSourceConfigBasic(keyStoreCertificateDetailsServiceInstanceGuid string, keyStoreCertificateDetailsQueueManagerID string, keyStoreCertificateDetailsLabel string, keyStoreCertificateDetailsCertificateFile string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = "%s"
			queue_manager_id = "%s"
			label = "%s"
			certificate_file = filebase64("%s")
		}

		data "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.service_instance_guid
			queue_manager_id = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.queue_manager_id
			label = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.label
		}
	`, keyStoreCertificateDetailsServiceInstanceGuid, keyStoreCertificateDetailsQueueManagerID, keyStoreCertificateDetailsLabel, keyStoreCertificateDetailsCertificateFile)
}
