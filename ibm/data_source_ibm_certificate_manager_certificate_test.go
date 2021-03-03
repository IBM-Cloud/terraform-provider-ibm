// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCertificateManagerCertificateDataSource_Basic(t *testing.T) {
	cmsName := fmt.Sprintf("tf-acc-test1-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCertificateManagerCertificateDataSourceConfig_basic(cmsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_certificate_manager_certificate.certificate", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMCertificateManagerCertificateDataSourceConfig_basic(cmsName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "cm" {
		name     = "%s"
		location = "us-south"
		service  = "cloudcerts"
		plan     = "free"
	}
	resource "ibm_certificate_manager_import" "cert"{
		certificate_manager_instance_id=ibm_resource_instance.cm.id
		name = "testimport"
		data = {
			content = <<EOF
-----BEGIN CERTIFICATE-----
MIIDZjCCAk4CCQDi5oda747xYzANBgkqhkiG9w0BAQsFADB1MQswCQYDVQQGEwJ1
czELMAkGA1UECAwCc2ExCzAJBgNVBAcMAnVzMQwwCgYDVQQKDANpZG0xDDAKBgNV
BAsMA2VkeDEQMA4GA1UEAwwHaWJtLmNvbTEeMBwGCSqGSIb3DQEJARYPbGtqaGdm
ZEBpYm0uY29tMB4XDTIwMDcxNDEwMjg0NFoXDTMwMDcxMjEwMjg0NFowdTELMAkG
A1UEBhMCdXMxCzAJBgNVBAgMAnNhMQswCQYDVQQHDAJ1czEMMAoGA1UECgwDaWRt
MQwwCgYDVQQLDANlZHgxEDAOBgNVBAMMB2libS5jb20xHjAcBgkqhkiG9w0BCQEW
D2xramhnZmRAaWJtLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJhVD1drJAWO0QY8UFGZzut1mOgD+2+Or8AKOjrp6yAMWTAw57rM0pK1lxxOwcYl
vCNHg12FVsjHpKzXsnxNbj7YQ0NKxz3l1bZXfAKxiZ00CVV1/e6AFIChaH3sgYY+
JZLH2TtKDlzg+14UvuencsUa+4tN+8lWSfJUVxCsIbdcOO78nVAurS/61Sk7/zei
+i1podwtNcjir1Xbuzlai5U+K1H8Y7NKnPlTx8gUXBEcDHdtnvdZGUpRFEDgw+dl
AeLrabQz76qDZcVDgqDueT+Uo8Ri6XaN8pbkG4Ei8M7ceoO6tOgD/yaCRPToq3Ov
8b7V6navCDsVxG/V2n6cYu0CAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAKzq4CD+O
TYMILLdQUGQJ+Ftq7/3J/0TRd6IshOXGVF/O6eZq4nxY+o4ZAeJlmBNSkoRI5+U+
gLJ6c41bBCIUloAZec68d5NG4Vz26uYy2JeGLCeIK9g/A9E8XTfVXUgrMvbLNd6q
3Dvduit044KhItuyEvInfVT1pmMsGOmclyEH/4mAOAb/WIY0kMpL3Oqy3g+lp2nY
XBOSiRxWgWsPCWUO34J5NEc0QU54l4VIuOOc7gcex6PP8eRDqmEgv7UEFZJiaQ0h
lEVEJFVWW8YRGX8n+NJb9LWq7HHNvBTcvZl2T8IUZx7oLAi2U4tyfa0v8iEyvju/
9HNwUobLq0q7/Q==
-----END CERTIFICATE-----
			EOF
		}
	}
	data "ibm_certificate_manager_certificate" "certificate"{
		certificate_manager_instance_id=ibm_resource_instance.cm.id
		name = ibm_certificate_manager_import.cert.name
	}
	  
	`, cmsName)
}
