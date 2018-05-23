package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func TestAccIBMLbaas_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMLbaasConfig_update(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "updated desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "protocols.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMLbaasConfig_updateHTTPS(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "updated desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "protocols.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMLbaas_with_more_protocols(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasConfig_MoreThanTwoProtocols(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "protocols.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMLbaas_importBasic(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasConfig_basic(name),
			},

			resource.TestStep{
				ResourceName:      "ibm_lbaas.lbaas",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"subnets.#", "subnets.0", "wait_time_minutes"},
			},
		},
	})
}

func TestAccIBMLbaas_InvalidProtocol(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaas_InvalidProtocol,
				ExpectError: regexp.MustCompile("must contain a value from"),
			},
		},
	})
}

func TestAccIBMLbaas_InvalidPort(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaas_InvalidPort,
				ExpectError: regexp.MustCompile("must be in the range of"),
			},
		},
	})
}

func TestAccIBMLbaas_InvalidMethod(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaas_InvalidMethod,
				ExpectError: regexp.MustCompile("must contain a value from"),
			},
		},
	})
}

func TestAccIBMLbaas_InvalidMaxConn(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaas_InvalidMaxConn,
				ExpectError: regexp.MustCompile("must be between 1 and 64000"),
			},
		},
	})
}

func TestAccIBMLbaas_certificate_with_http_invalid_config(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaas_certificate_with_http_invalid_config(name),
				ExpectError: regexp.MustCompile("tls_certificate_id may be set only when frontend protocol is 'HTTPS'"),
			},
		},
	})
}

func testAccCheckIBMLbaasDestroy(s *terraform.State) error {
	sess := testAccProvider.Meta().(ClientSession).SoftLayerSession()
	service := services.GetNetworkLBaaSLoadBalancerService(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_lbaas" {
			continue
		}

		// Try to find the key
		_, err := service.GetLoadBalancer(sl.String(rs.Primary.ID))

		if err == nil {
			return fmt.Errorf("load balancer (%s) to be destroyed still exists", rs.Primary.ID)
		} else if apiErr, ok := err.(sl.Error); ok && apiErr.Exception != NOT_FOUND {
			return fmt.Errorf("Error waiting for load balancer (%s) to be destroyed: %s", rs.Primary.ID, err)
		}

	}

	return nil
}

func testAccCheckIBMLbaas_certificate_with_http_invalid_config(name string) string {
	return fmt.Sprintf(`
resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "desc-used for terraform uat"
  subnets     = ["%s"]

  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "round_robin"
    "max_conn" = 65
    "tls_certificate_id" = 1234567
  }]

}
`, name, lbaasSubnetId)
}

const testAccCheckIBMLbaas_InvalidMaxConn = `
resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "desc-used for terraform uat"
  subnets     = [1511875]

  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTPs"
    "backend_port" = 80
    "load_balancing_method" = "round_robin"
    "max_conn" = 65000
  }]

}
`

const testAccCheckIBMLbaas_InvalidProtocol = `
resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "desc-used for terraform uat"
  subnets     = [1511875]

  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTPS"
    "backend_port" = 80
    "load_balancing_method" = "round_robin"
  }]

}
`

const testAccCheckIBMLbaas_InvalidPort = `
resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "desc-used for terraform uat"
  subnets     = [1511875]

  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 65536
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "round_robin"
  }]

}
`

const testAccCheckIBMLbaas_InvalidMethod = `
resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "desc-used for terraform uat"
  subnets     = [1511875]

  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "ROUND_ROUND_ROBIN"
  }]

}
`

func testAccCheckIBMLbaasConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
}
`, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasConfig_MoreThanTwoProtocols(name string) string {
	return fmt.Sprintf(`
resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 9090
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "round_robin"
    "session_stickiness" = "SOURCE_IP"
  },
  {

    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80

    "load_balancing_method" = "round_robin"
  },
  {

    "frontend_protocol" = "HTTP"
    "frontend_port" = 8081
    "backend_protocol" = "HTTP"
    "backend_port" = 80

    "load_balancing_method" = "round_robin"
  }]
}
`, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasConfig_update(name string) string {
	return fmt.Sprintf(`
resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "updated desc-used for terraform uat"
  subnets     = ["%s"]

  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "weighted_round_robin"
  }]

}
`, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasConfig_updateHTTPS(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_ssl_certificate" "test-cert" {
    certificate = <<EOF
-----BEGIN CERTIFICATE-----
MIIEujCCA6KgAwIBAgIJAKMRot3rBodEMA0GCSqGSIb3DQEBBQUAMIGZMQswCQYD
VQQGEwJVUzEQMA4GA1UECBMHR2VvcmdpYTEQMA4GA1UEBxMHQXRsYW50YTEMMAoG
A1UEChMDVFdDMQ0wCwYDVQQLEwRHcmlkMRYwFAYDVQQDFA0qLndlYXRoZXIuY29t
MTEwLwYJKoZIhvcNAQkBFiJ0aW0ubXVsaGVybi5jb250cmFjdG9yQHdlYXRoZXIu
Y29tMB4XDTE2MDYwMjE5MjcwOVoXDTE3MDYwMjE5MjcwOVowgZkxCzAJBgNVBAYT
AlVTMRAwDgYDVQQIEwdHZW9yZ2lhMRAwDgYDVQQHEwdBdGxhbnRhMQwwCgYDVQQK
EwNUV0MxDTALBgNVBAsTBEdyaWQxFjAUBgNVBAMUDSoud2VhdGhlci5jb20xMTAv
BgkqhkiG9w0BCQEWInRpbS5tdWxoZXJuLmNvbnRyYWN0b3JAd2VhdGhlci5jb20w
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDgVW1J8vhrOFBCBx7Rqz5I
/3WKChjxYe8MK/TkfVfCyHBe7dAdaiRyP4YLU5O1wyTvk6XNOM2I2W1l6Hmoa2RV
eo20k3NILLAZvhPeNoCQDMvRUdo8jKXxuerz+1oxYb4ip/BUZDN6EBDkBckptciP
yeB/cwCZI+thdnuEgp3H74nZrQQmOxow+HTSY00hd92IF4Jz8Qb/C2relyJB1bMZ
uk5BQc39FyBFJLYp5yiRUSVU22GtbaLFuQsdtVfxEwPCRG5a1piy3MLq9VIQYcbv
/1y02EmnMCM/Zfhw+rjz53XCy6e0lT/02w6fp2TEIGuFVKAvZrUsLkM6XGLoqDn7
AgMBAAGjggEBMIH+MB0GA1UdDgQWBBTI9DVDsxajJ/EQ1SdjnpEmCrHahzCBzgYD
VR0jBIHGMIHDgBTI9DVDsxajJ/EQ1SdjnpEmCrHah6GBn6SBnDCBmTELMAkGA1UE
BhMCVVMxEDAOBgNVBAgTB0dlb3JnaWExEDAOBgNVBAcTB0F0bGFudGExDDAKBgNV
BAoTA1RXQzENMAsGA1UECxMER3JpZDEWMBQGA1UEAxQNKi53ZWF0aGVyLmNvbTEx
MC8GCSqGSIb3DQEJARYidGltLm11bGhlcm4uY29udHJhY3RvckB3ZWF0aGVyLmNv
bYIJAKMRot3rBodEMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADggEBABrz
RWXhnGKSJj3isBFjdVgb6oIymW4bHeCMRVKxm5p+yJqv1LiCZzUah0aNjRRua4k3
nUBIs+c2SO7WVuyDgQ87oq+shEL2H3G07cvl8vVESr4r/K7R5fwYUCobOeAr6qSB
sj9ZiJqQ02NfD4q4E0gS/P8CuL9w76M8350WSahKDx3VNUs/QIm6nZy/8OhCQYqq
Q2xmxuSPiI9MNEAh8IfYVBH4qi51SlSRiDJoGXmmbkwa+YZyfpEiZeisHVNNdVrm
DDtf0yuw5VRx2wnTWhv+ezUkhRGCL80fnqkWB94IS66UHlO5WyHw1cgQEVW1ie2y
baU37Sk90FDVrroBgNY=
-----END CERTIFICATE-----
    EOF
    private_key = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4FVtSfL4azhQQgce0as+SP91igoY8WHvDCv05H1XwshwXu3Q
HWokcj+GC1OTtcMk75OlzTjNiNltZeh5qGtkVXqNtJNzSCywGb4T3jaAkAzL0VHa
PIyl8bnq8/taMWG+IqfwVGQzehAQ5AXJKbXIj8ngf3MAmSPrYXZ7hIKdx++J2a0E
JjsaMPh00mNNIXfdiBeCc/EG/wtq3pciQdWzGbpOQUHN/RcgRSS2KecokVElVNth
rW2ixbkLHbVX8RMDwkRuWtaYstzC6vVSEGHG7/9ctNhJpzAjP2X4cPq48+d1wsun
tJU/9NsOn6dkxCBrhVSgL2a1LC5DOlxi6Kg5+wIDAQABAoIBAHwOgduNI9eXUrrQ
2Tg1rMINk2B86QJDmEBw5oKc1jV/RrUYaih6FCGiA2ysEVlIy1o5mkz9BpyRMLBU
eUKr8NZcaZTcnbniDJiPxsjx9vKyQNxGmZs2ZGZi3A2EiIIafV0I5hylNNphnBWd
JXuNbZYmm6GfZUtK09YYAYJsAPkY8xxk274YfPOQQbWFMl5sR1QqXCzDDJ23hgIS
9pw05oHx++HliC+rsExOJ3K+j3X2HGBlQgQJjEJBDxs1ttSLxoFAHcUSyGJGsXud
fgvJf6GkcJ/JnAi8qhH5IV50/X3YWdosY2fGBzR7Naasfh4IrNq6tZ+1L5c/6agP
RfKU+0ECgYEA+t6fPgcE211inH4H0i8H5HrI/sgmsF7uXiobbcUCFBJR3rT8XUq0
9x7SEj5CokvpDm1pM3ktv/fffB2W74pcpn63n8rWjHOu3/LMvnab8Ad54wA7IMF8
/vvjhbqZaWhbYt93o5bFP6U3QlfLaMRItr+0KLm7kyJ4GBC6QGRSDhECgYEA5Ovh
oBILLZriVcuVwYeLxuzjCCJohpFkUtXmxUpwLKYRVAsC0MSNTjvZfJkVOvR9G8Ki
Cmy7wGt1VIo8M7DKmetHTsXn6H9S0SN62ykKX/ob/D1g0tFETsEFkVt7mha3Q1AB
6VR9LiohCQAevoOLn+Vm8B4aHyOGjah2FgPta0sCgYAN3lbBUBQFqID2E8WM6gqu
p9cKtrfk0iqtS/ieNeDqiSS7ghfddG7SpoKIfaajYDzvDj9dmBpeXW6eZuhcL7L1
hVXTYJxBwXdua/bDpLz0JQWo9e9O3UNyuSwXzXwDpsA+lAoCIiifXxvR8BaPoSI/
8BMemT30YVhwRCR3wNQEcQKBgQCwcULRTrcA6p1DBYyiwuewZotCjMrF1bBezHF3
ZT16nHFEtsvvv18uiqDCEXe0nhcD24trv40i7XBcvcNTEBPIePjYNV/e6qwZeGBM
JaDSgwMo8uH6+8LLdKjm9X0aMiIEptkiT7XAbEZUGpyXuOpYTsd9kaYOlCI0c0C5
DUPkawKBgGlwzHX3dr7jYldmB9/g94jWeNkX6KPtSDNaKZ9WzIuywCB6wua7AVXa
NXMjAHErbX2J+8k85TccHR1ps3MgBbFHdiuJwx2vUPLfVj53GWUXmg4Gw4zUs5mq
ykXbeuyhK6AL6V3NsJyP454bM8dmZnxBrZvRo5FnqQInGgwGSjgc
-----END RSA PRIVATE KEY-----
    EOF
}

resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "updated desc-used for terraform uat"
  subnets     = ["%s"]


  protocols = [{
    "frontend_protocol" = "HTTPS"
    "frontend_port" = 443
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "round_robin"
    "session_stickiness" = "SOURCE_IP"
    "tls_certificate_id" = "${ibm_compute_ssl_certificate.test-cert.id}"
  },
  {

    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80

    "load_balancing_method" = "round_robin"
  }]

}
`, name, lbaasSubnetId)
}
