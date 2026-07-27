// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsLbListenerDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))
	protocol1 := "http"
	port1 := "8080"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsLbListenerDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "listener_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "accept_proxy_protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "port"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "port_max"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "port_min"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener", "provisioning_status"),
				),
			},
		},
	})
}

func TestAccIBMIsLbListenerDataSource_ClientAuth(t *testing.T) {
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))
	protocol := "https"
	port := "443"

	// Example CRNs - replace with actual values from your test environment
	certCRN := "crn:v1:bluemix:public:secrets-manager:us-south:a/7f75c7b025e54bc5635f754b2f888665:152af435-37ac-4b3e-83c3-828805bfc8e0:secret:0ed079f8-5b93-66e9-86c6-ba79157036d6"
	caCRN := "crn:v1:bluemix:public:secrets-manager:us-south:a/7f75c7b025e54bc5635f754b2f888665:152af435-37ac-4b3e-83c3-828805bfc8e0:secret:78b29d17-a610-9349-a220-32a50f3cd9e6"

	// CRL content - must match what's in testAccCheckIBMISLBListenerClientAuthConfigUpdate
	crlContent := `-----BEGIN X509 CRL-----
MIICvTCBpgIBATANBgkqhkiG9w0BAQsFADBMMQswCQYDVQQGEwJVUzEOMAwGA1UE
CAwFZGVsYXMxDDAKBgNVBAoMA0lCTTENMAsGA1UECwwEcm9vdDEQMA4GA1UEAwwH
cm9vdC1jYRcNMjUwOTA4MDUwMjQwWhcNMjUxMDA4MDUwMjQwWjAVMBMCAhAAFw0y
NTA5MDgwNTAxNTlaoA8wDTALBgNVHRQEBAICEAAwDQYJKoZIhvcNAQELBQADggIB
ACeEcj7ompUepc5qTvTrNA5PoK5bN71gNI7Rbhq/Bxf1YPMp2iU3qMSj7YpVP7aw
GNrxFoIZcQ4X7PYyHMfDk6Z83PSTVMnSOVk09fZW49tyVTWmzBVLz3R1bPasnWTZ
0hRIv9j9n7Lemin+0ubIR/2zmsfBs1JFAFEbbRcgwg+qotsfZNLkX6bjHDpsRQzE
mXUEu4/AqAsWPbFzG2uMKZ9pKOK+Nn3bt/NEK+AFlnSmgjEqzQ+0zhsrCExIReJV
c2oiLBkLG6rBwxlGDog+PqwjP+1wGNIL1J3c2lMW1IGMNcts/aDBO5LtPVIY1LsQ
FoeaTfm3U3GKC/pTczoDk/pKN756f8O05nTWUHgktcNsPvgqDKnpvEkI3VPf9Y4a
fMOzKgVTgY1dSgjzHO8+4ZfcVGpBePsjOe0/RCUwkgtgOyGtcmBPTMJa0elJzjaM
jD9myqIXkB359sqbuEmcrjgo5uUUvubFYpmT/W0YxOi/py/bDK+7uUs38nUElNkZ
+YFRpNWjLF9JtAghX5MhA5BwhTTuATvWYuDdK769ifi9qcYvE4u+VNxYfOpPY6sv
x4FnkZ9+A7s2hk11d+DEq29Efa0xak8rO1LzT5hCSFT0P3KfZEZMpbuXpzVGiZoM
g5cWHgYcNnzhUatKodvzZizAOVGRR7UFg42O4ylhxDVe
-----END X509 CRL-----
`

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsLbListenerDataSourceConfigClientAuth(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, certCRN, caCRN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener_mtls", "listener_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener_mtls", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener_mtls", "id"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_listener.is_lb_listener_mtls", "protocol", protocol),
					resource.TestCheckResourceAttr("data.ibm_is_lb_listener.is_lb_listener_mtls", "port", port),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener_mtls", "client_authentication.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener.is_lb_listener_mtls", "client_authentication.0.certificate_authority.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_listener.is_lb_listener_mtls", "client_authentication.0.certificate_authority.0.crn", caCRN),
					resource.TestCheckResourceAttr("data.ibm_is_lb_listener.is_lb_listener_mtls", "client_authentication.0.certificate_revocation_list", crlContent),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbListenerDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
	return testAccCheckIBMISLBListenerConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol) + fmt.Sprintf(`

	data "ibm_is_lb_listener" "is_lb_listener" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		listener_id = ibm_is_lb_listener.testacc_lb_listener.listener_id
	}
	`)
}

func testAccCheckIBMIsLbListenerDataSourceConfigClientAuth(vpcname, subnetname, zone, cidr, lbname, port, protocol, certCRN, caCRN string) string {
	return testAccCheckIBMISLBListenerClientAuthConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, certCRN, caCRN) + fmt.Sprintf(`

	data "ibm_is_lb_listener" "is_lb_listener_mtls" {
		lb = ibm_is_lb.testacc_LB.id
		listener_id = ibm_is_lb_listener.testacc_lb_listener_mtls.listener_id
	}
	`)
}
