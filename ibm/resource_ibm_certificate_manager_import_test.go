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
MIICZTCCAc4CCQD0OJQAR6NJvjANBgkqhkiG9w0BAQsFADB3MQswCQYDVQQGEwJ1
czELMAkGA1UECAwCa2ExDDAKBgNVBAcMA2libTEMMAoGA1UECgwDaWJtMQwwCgYD
VQQLDANpZXoxEDAOBgNVBAMMB2libS5jb20xHzAdBgkqhkiG9w0BCQEWEGFzZGZn
aGprQGlibS5jb20wHhcNMjAwNTE5MTIwMDA5WhcNMjAwNjE4MTIwMDA5WjB3MQsw
CQYDVQQGEwJ1czELMAkGA1UECAwCa2ExDDAKBgNVBAcMA2libTEMMAoGA1UECgwD
aWJtMQwwCgYDVQQLDANpZXoxEDAOBgNVBAMMB2libS5jb20xHzAdBgkqhkiG9w0B
CQEWEGFzZGZnaGprQGlibS5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGB
ALEahudA3Wubs8rHEBe1cyml5QoAkGAbCnpm2yhGyL6ld2OkOhVjCjmEWnMHHmFc
BndeB0NeCyOxdw+hFEK0QvU/dOjDnSXKkVM4c2lcy4H4rWIXrGxPWi+BCwY+hmEX
hf1ERuodADS+hi8jLQVOAjzykuMxTcCEqXrrYFKaO7EXAgMBAAEwDQYJKoZIhvcN
AQELBQADgYEAYGBJ2C5LRJ1Ul2zhHMPXN1AX7VgaOrs320UYinFqh5JB2W9NngZA
ZUQXSnjstIzqZCbEj6dTau2yaxjyNlFGDQLuiyxzyRfW6COuWt9cGULdL+Md/JX1
phHiFNtuMr8s2Pu4kz/SSN39DjuKG7XYcZQOYf6q8IW4JIHfQ5U9jOU=
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
