package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDIDPSamlDataSource_basic(t *testing.T) {
	dispName := fmt.Sprintf("testacc_saml_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDIDPSAMLDataSourceConfig(acc.AppIDTenantID, dispName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "is_active", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "config.0.entity_id", "https://test-saml-idp"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "config.0.sign_in_url", "https://test-saml-idp/login"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "config.0.display_name", dispName),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "config.0.encrypt_response", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "config.0.sign_request", "false"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "config.0.certificates.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml.test_saml", "config.0.certificates.0", `MIIDNTCCAh2gAwIBAgIRAPbl3OBL5oXq47d98l2s/3IwDQYJKoZIhvcNAQELBQAw
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
`),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDIDPSAMLDataSourceConfig(tenantID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_idp_saml" "saml" {
			tenant_id = "%s"
  			is_active = true
			config {
				entity_id = "https://test-saml-idp"
				sign_in_url = "https://test-saml-idp/login"
				display_name = "%s"
				encrypt_response = true
				sign_request = false
				certificates = [					
					<<EOF
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
EOF
				]
			}
		}	
		data "ibm_appid_idp_saml" "test_saml" {
			tenant_id = ibm_appid_idp_saml.saml.tenant_id

			depends_on = [
				ibm_appid_idp_saml.saml
			]
		}
	`, tenantID, name)
}
