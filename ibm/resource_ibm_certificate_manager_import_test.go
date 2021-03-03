// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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

		if err != nil && !strings.Contains(err.Error(), "404") && !strings.Contains(err.Error(), "412") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil

}
func TestAccIBMCertificateManager_Import(t *testing.T) {
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
MIIDNTCCAh2gAwIBAgIRAPbl3OBL5oXq47d98l2s/3IwDQYJKoZIhvcNAQELBQAw
KDEQMA4GA1UEChMHRXhhbXBsZTEUMBIGA1UEAxMLZXhhbXBsZS5jb20wHhcNMjAw
OTIxMTEyMjI1WhcNMzAwOTI5MTEyMjI1WjAoMRAwDgYDVQQKEwdFeGFtcGxlMRQw
EgYDVQQDEwtleGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBAJqMcqnms1XpCuKz+CIVrqppMog9aerAQEV5wY6XuvakZ/w89zrA7YX3vwgi
+0ZO9ldDBh5Wvl8Li8vDFALJc42MxxyENk4qB6zee1O+zYu1Bwynkp7nIxqyKKRd
+0tvc+WHUbPFHvXc94rajT/csHOvBRiLmABMBx/IqF1nEAG/+KAEh7+KZYbvQ6wk
OoiPZlW+B0HR/DL/uO/v1Q7eq2Z8pAVTGikHefckvolkOqiCIRZx8HDe8DxTojEm
ygiR1aeT29XV8frI3Y2C8e7vgDpuZ8nV+0JUzqi5tAfl8bUfuq/W0eng6BYk2hBD
uuS66fHb1hnW96WIaExlK6T096sCAwEAAaNaMFgwDgYDVR0PAQH/BAQDAgKkMA8G
A1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFP3hKtk3MVXSF1H79ukO7oBVwwAkMBYG
A1UdEQQPMA2CC2V4YW1wbGUuY29tMA0GCSqGSIb3DQEBCwUAA4IBAQA9TbumFQHA
SHS6DBzzJz8GeX451AelW8UtpIuc5mRDFvTFEuNn/wMikxi+m8SQgkcuO5wfQi+0
FzLQO8DYH6fnAo1BYqooT1bt4lXflt74FnyUYbZ75yUhsddYF00FYOX6eOxrAU/U
qaPXw2N/e6S859hsUMq79/g3ES9sdNedtiwgiQv7roh4WNSvgTLh+sD32Ehl+x/I
eE80MljFLf5bfu2bQqV7C17lszGxTQWI2Xj56gLr2jcITjltcHCuBwnRDyXJkNhq
/2KRyIGAaRkkCOJAJxiz82wxkuQ8aL4sD3dctfGNu2Qe1JXHB65M1P2m0j/IcrLT
iUCoFQ0xO5VC
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
