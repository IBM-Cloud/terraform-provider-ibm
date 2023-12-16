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
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "queue_manager_id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.#"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.label", keyStoreCertificateDetailsLabel),
				),
			},
		},
	})
}

func TestAccIbmMqcloudKeystoreCertificateDataSourceAllArgs(t *testing.T) {
	keyStoreCertificateDetailsServiceInstanceGuid := acc.MqcloudInstanceID
	keyStoreCertificateDetailsQueueManagerID := acc.MqcloudQueueManagerID
	keyStoreCertificateDetailsLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	keyStoreCertificateDetailsCertificateFile := acc.MqcloudKSCertFilePath

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudKeystoreCertificateDataSourceConfig(keyStoreCertificateDetailsServiceInstanceGuid, keyStoreCertificateDetailsQueueManagerID, keyStoreCertificateDetailsLabel, keyStoreCertificateDetailsCertificateFile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "queue_manager_id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.#"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.label", keyStoreCertificateDetailsLabel),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.certificate_type"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.fingerprint_sha256"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.subject_dn"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.subject_cn"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.issuer_dn"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.issuer_cn"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.issued"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.expiry"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.is_default"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.dns_names_total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "key_store.0.href"),
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
			certificate_file = "%s"
		}

		data "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.service_instance_guid
			queue_manager_id = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.queue_manager_id
			label = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.label
		}
	`, keyStoreCertificateDetailsServiceInstanceGuid, keyStoreCertificateDetailsQueueManagerID, keyStoreCertificateDetailsLabel, keyStoreCertificateDetailsCertificateFile)
}

func testAccCheckIbmMqcloudKeystoreCertificateDataSourceConfig(keyStoreCertificateDetailsServiceInstanceGuid string, keyStoreCertificateDetailsQueueManagerID string, keyStoreCertificateDetailsLabel string, keyStoreCertificateDetailsCertificateFile string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = "%s"
			queue_manager_id = "%s"
			label = "%s"
			certificate_file = "%s"
		}

		data "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.service_instance_guid
			queue_manager_id = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.queue_manager_id
			label = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.label
		}
	`, keyStoreCertificateDetailsServiceInstanceGuid, keyStoreCertificateDetailsQueueManagerID, keyStoreCertificateDetailsLabel, keyStoreCertificateDetailsCertificateFile)
}
