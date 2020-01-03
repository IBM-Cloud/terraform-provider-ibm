package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/IBM-Cloud/bluemix-go/models"
)

func TestAccIBMCertificateManager_Basic(t *testing.T) {
	var conf models.CertificateGetData
	name1 := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	name2 := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCertificateManagerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCertificateManager_basicImport(name1, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCMExists("ibm_certificate_manager_import.cert", conf),
					resource.TestCheckResourceAttr("ibm_certificate_manager_import.cert", "name", name2),
				),
			},
		},
	})
}

func testAccCheckIBMCertificateManagerDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_certificate_manager_import" {
			continue
		}
		certID := rs.Primary.ID
		cmClient, err := testAccProvider.Meta().(ClientSession).CertificateManagerAPI()
		if err != nil {
			return err
		}
		certAPI := cmClient.Certificate()
		_, err = certAPI.GetCertData(certID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil

}
func TestAccIBMCMimport(t *testing.T) {
	var conf models.CertificateGetData
	name1 := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	name2 := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCertificateManagerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCertificateManager_basicImport(name1, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCMExists("ibm_certificate_manager_import.cert", conf),
					resource.TestCheckResourceAttr("ibm_certificate_manager_import.cert", "name", name2),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_certificate_manager_import.cert",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCertificateManager_basicImport(name1, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "cm" {
	name                = "%s"
	location            = "us-south"
	service             = "cloudcerts"
	plan                = "free"
	}
	
	resource "ibm_certificate_manager_import" "cert"{
	certificate_manager_instance_id=ibm_resource_instance.cm.id
	name = "%s"
	data = {
		content = <<EOF
-----BEGIN CERTIFICATE-----
MIICezCCAeQCCQDBQ+jPkch9PjANBgkqhkiG9w0BAQsFADCBgTELMAkGA1UEBhMC
dXMxDTALBgNVBAgMBGFzZGYxDzANBgNVBAcMBmFzZGZnaDESMBAGA1UECgwJYXNk
ZmdoamtsMQ4wDAYDVQQLDAVzZGZnaDEQMA4GA1UEAwwHaWJtLmNvbTEcMBoGCSqG
SIb3DQEJARYNYXNkZmdoamtsLmNvbTAeFw0xOTEyMjMwNzEwNTdaFw0yMDAxMjIw
NzEwNTdaMIGBMQswCQYDVQQGEwJ1czENMAsGA1UECAwEYXNkZjEPMA0GA1UEBwwG
YXNkZmdoMRIwEAYDVQQKDAlhc2RmZ2hqa2wxDjAMBgNVBAsMBXNkZmdoMRAwDgYD
VQQDDAdpYm0uY29tMRwwGgYJKoZIhvcNAQkBFg1hc2RmZ2hqa2wuY29tMIGfMA0G
CSqGSIb3DQEBAQUAA4GNADCBiQKBgQDF80KUhdb0efLkWFXlVllC2vP020zwfh9C
vQ7T1AamBJPzHZHcJw48xb9TJP+hlA2iV4jyeECAYzwJlv/Fo2my17pTl1haKyw2
D9gpxny0FRv3GngIFFqgluVfOBZdp6quyA9uPyrPzny9JRKypaGKtLG1wYnvQKg8
gg+0HDdX3wIDAQABMA0GCSqGSIb3DQEBCwUAA4GBABefpuSBwTkXi9U9v+agb97y
yufg0Zb+8T0TQbcdtdPb/IR2q3XQdyhfyeR0CfRlR94jeyJYmBEXVvc7J5BO2ram
54TJWEeeZkEaz36ozpYjPpOttZ8S22di3ediYiNsoAUhzeDZ24/X1YhBTK71VuKR
9x90g2Y3xPf6KLEZukR2
-----END CERTIFICATE-----
		EOF
	  }
	}
	  `, name1, name2)
}

func testAccCheckIBMCMExists(n string, obj models.CertificateGetData) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cmClient, err := testAccProvider.Meta().(ClientSession).CertificateManagerAPI()
		if err != nil {
			return err
		}
		certID := rs.Primary.ID

		crt, err := cmClient.Certificate().GetCertData(certID)
		if err != nil {
			return err
		}

		obj = crt
		return nil
	}
}
