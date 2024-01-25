// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
)

func TestAccIbmMqcloudKeystoreCertificateBasic(t *testing.T) {
	t.Parallel()
	var conf mqcloudv1.KeyStoreCertificateDetails
	serviceInstanceGuid := acc.MqcloudInstanceID
	queueManagerID := acc.MqcloudQueueManagerID
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	certificateFile := acc.MqcloudKSCertFilePath

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckMqcloud(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudKeystoreCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudKeystoreCertificateConfigBasic(serviceInstanceGuid, queueManagerID, label, certificateFile),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudKeystoreCertificateExists("ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "queue_manager_id", queueManagerID),
					resource.TestCheckResourceAttr("ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "label", label),
				),
			},
			{
				ResourceName:            "ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"certificate_file"},
			},
		},
	})
}

func testAccCheckIbmMqcloudKeystoreCertificateConfigBasic(serviceInstanceGuid string, queueManagerID string, label string, certificateFile string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = "%s"
			queue_manager_id = "%s"
			label = "%s"
			certificate_file = filebase64("%s")
		}
	`, serviceInstanceGuid, queueManagerID, label, certificateFile)
}

func testAccCheckIbmMqcloudKeystoreCertificateExists(n string, obj mqcloudv1.KeyStoreCertificateDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
		if err != nil {
			return err
		}

		getKeyStoreCertificateOptions := &mqcloudv1.GetKeyStoreCertificateOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getKeyStoreCertificateOptions.SetServiceInstanceGuid(parts[0])
		getKeyStoreCertificateOptions.SetQueueManagerID(parts[1])
		getKeyStoreCertificateOptions.SetCertificateID(parts[2])

		keyStoreCertificateDetails, _, err := mqcloudClient.GetKeyStoreCertificate(getKeyStoreCertificateOptions)
		if err != nil {
			return err
		}

		obj = *keyStoreCertificateDetails
		return nil
	}
}

func testAccCheckIbmMqcloudKeystoreCertificateDestroy(s *terraform.State) error {
	mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_mqcloud_keystore_certificate" {
			continue
		}

		getKeyStoreCertificateOptions := &mqcloudv1.GetKeyStoreCertificateOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getKeyStoreCertificateOptions.SetServiceInstanceGuid(parts[0])
		getKeyStoreCertificateOptions.SetQueueManagerID(parts[1])
		getKeyStoreCertificateOptions.SetCertificateID(parts[2])

		// Try to find the key
		_, response, err := mqcloudClient.GetKeyStoreCertificate(getKeyStoreCertificateOptions)

		if err == nil {
			return fmt.Errorf("mqcloud_keystore_certificate still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for mqcloud_keystore_certificate (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
