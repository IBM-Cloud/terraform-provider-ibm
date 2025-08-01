// Copyright IBM Corp. 2025 All Rights Reserved.
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

func TestAccIbmMqcloudTruststoreCertificateBasic(t *testing.T) {
	var conf mqcloudv1.TrustStoreCertificateDetails
	serviceInstanceGuid := acc.MqcloudDeploymentID
	queueManagerID := acc.MqcloudQueueManagerID
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	certificateFile := acc.MqcloudTSCertFilePath

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckMqcloud(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudTruststoreCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudTruststoreCertificateConfigBasic(serviceInstanceGuid, queueManagerID, label, certificateFile),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudTruststoreCertificateExists("ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance", "queue_manager_id", queueManagerID),
					resource.TestCheckResourceAttr("ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance", "label", label),
				),
			},
			{
				ResourceName:            "ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"certificate_file"},
			},
		},
	})
}

func testAccCheckIbmMqcloudTruststoreCertificateConfigBasic(serviceInstanceGuid string, queueManagerID string, label string, certificateFile string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
			service_instance_guid = "%s"
			queue_manager_id = "%s"
			label = "%s"
			certificate_file = filebase64("%s")
		}
	`, serviceInstanceGuid, queueManagerID, label, certificateFile)
}

func testAccCheckIbmMqcloudTruststoreCertificateExists(n string, obj mqcloudv1.TrustStoreCertificateDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
		if err != nil {
			return err
		}

		getTrustStoreCertificateOptions := &mqcloudv1.GetTrustStoreCertificateOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTrustStoreCertificateOptions.SetServiceInstanceGuid(parts[0])
		getTrustStoreCertificateOptions.SetQueueManagerID(parts[1])
		getTrustStoreCertificateOptions.SetCertificateID(parts[2])

		trustStoreCertificateDetails, _, err := mqcloudClient.GetTrustStoreCertificate(getTrustStoreCertificateOptions)
		if err != nil {
			return err
		}

		obj = *trustStoreCertificateDetails
		return nil
	}
}

func testAccCheckIbmMqcloudTruststoreCertificateDestroy(s *terraform.State) error {
	mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_mqcloud_truststore_certificate" {
			continue
		}

		getTrustStoreCertificateOptions := &mqcloudv1.GetTrustStoreCertificateOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTrustStoreCertificateOptions.SetServiceInstanceGuid(parts[0])
		getTrustStoreCertificateOptions.SetQueueManagerID(parts[1])
		getTrustStoreCertificateOptions.SetCertificateID(parts[2])

		// Try to find the key
		_, response, err := mqcloudClient.GetTrustStoreCertificate(getTrustStoreCertificateOptions)

		if err == nil {
			return fmt.Errorf("mqcloud_truststore_certificate still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for mqcloud_truststore_certificate (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
