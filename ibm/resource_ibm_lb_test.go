package ibm

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMLbShared_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbSharedConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "250"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLbSharedConfig_UpgradeConnectionLimit,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "500"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMLbDedicated_Basic(t *testing.T) {
	t.SkipNow()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbDedicatedConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "15000"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "true"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "true"),
				),
			},
		},
	})
}

func TestAccIBMLbSharedWithTag(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbSharedConfigWithTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "250"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "tags.#", "2"),
				),
			},
			{

				Config: testAccCheckIBMLbSharedConfigWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "250"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMLbSSL_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbSSLConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "250"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_offload", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLbSSLConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "250"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_offload", "true"),
				),
			},
		},
	})
}

const testAccCheckIBMLbSharedConfig_basic = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "dal09"
    ha_enabled  = false
}`

const testAccCheckIBMLbSharedConfig_UpgradeConnectionLimit = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 500
    datacenter    = "dal09"
	ha_enabled  = false	
}`

const testAccCheckIBMLbDedicatedConfig_basic = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 15000
    datacenter    = "dal09"
    ha_enabled  = false
    dedicated = true	
}`

const testAccCheckIBMLbSharedConfigWithTag = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "dal09"
	ha_enabled  = false
	tags = ["one", "two"]
}`

const testAccCheckIBMLbSharedConfigWithUpdatedTag = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "dal09"
	ha_enabled  = false
	tags = ["one", "two", "three"]
}`

const testAccCheckIBMLbSSLConfig_basic = `

resource "ibm_compute_ssl_certificate" "test-cert" {
    certificate = <<EOF
-----BEGIN CERTIFICATE-----
MIICsTCCAhoCCQDl71v1qDksRTANBgkqhkiG9w0BAQUFADCBnDELMAkGA1UEBhMC
SU4xEjAQBgNVBAgMCUJhbmdhbG9yZTESMBAGA1UEBwwJQmFuZ2Fsb3JlMQwwCgYD
VQQKDANJQk0xEjAQBgNVBAsMCVRlcnJhZm9ybTEfMB0GA1UEAwwWd3d3LnRlcnJh
Zm9ybS10ZXN0LmNvbTEiMCAGCSqGSIb3DQEJARYTaGthbnRhcmVAaW4uaWJtLmNv
bTAeFw0xODEwMTkwNjU5MDBaFw0xOTEwMTkwNjU5MDBaMIGcMQswCQYDVQQGEwJJ
TjESMBAGA1UECAwJQmFuZ2Fsb3JlMRIwEAYDVQQHDAlCYW5nYWxvcmUxDDAKBgNV
BAoMA0lCTTESMBAGA1UECwwJVGVycmFmb3JtMR8wHQYDVQQDDBZ3d3cudGVycmFm
b3JtLXRlc3QuY29tMSIwIAYJKoZIhvcNAQkBFhNoa2FudGFyZUBpbi5pYm0uY29t
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDkgLoMrxpPpoYHs08VAVp56xA7
/DNeVawkbYC693tovvO5HkzIJmXAq/QwDRSLyhanIwzGckTnIFsTS/lD0yF9Wq7w
dhWKum87YAqu/2/FtqIxmDQl4Ma55LDtg0GHhLIqLUrCAIvmoZMUuSrh31PrJDGF
c/bDalB6ioCXzlgMwQIDAQABMA0GCSqGSIb3DQEBBQUAA4GBANO8m68Sa3gVrId6
ziJUxLYMzY0i+e6rq/WyEgbne/nBauOENlzpCUp2lPnTncH6LqnOYu2NUA37GYqC
HVFvrAqGy8JzuW/M4vtipCxFIVGkepy/RlnyWJZas4PYgES2bZES+f8or4UwtrcG
g7b8cLkvBhfZkB2X/NgqBqhxBFVA
-----END CERTIFICATE-----
    EOF
    private_key = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDkgLoMrxpPpoYHs08VAVp56xA7/DNeVawkbYC693tovvO5HkzI
JmXAq/QwDRSLyhanIwzGckTnIFsTS/lD0yF9Wq7wdhWKum87YAqu/2/FtqIxmDQl
4Ma55LDtg0GHhLIqLUrCAIvmoZMUuSrh31PrJDGFc/bDalB6ioCXzlgMwQIDAQAB
AoGBAJG0u+5mocJ0jzbN0fm0+TqQ97MoaKEYxEIeSV3vfZQXX1aFybQ/N9caTwVs
8dMJtFQzd2v7ZZB0A19UrMfhE5KlD/pNaNmsZBz3qTqAs9zpU0ZNExGsWyVQgOFw
MyfgOj18q7ZU6Hujc6XknRj3pOYFtLnv0bmaiKObzuxqtsShAkEA8/Fm7ItRmbsu
k3qrKzhYNFx1beT61DTwr6g/BSaczsmIOnFFnAvnqya8L5/kMNRLzx88QrhXo32K
xea8i17T/QJBAO/L9iDOF1RxlpiA9OHc2mH86boyjkeINQZK12Hce0XHtsVgQtem
KCD6VdDXitdx1DgnGpI36BfqrfkxxtDRHRUCQAeAEZkOQ4kFf04bhG3EwrmBaj7h
vnCN3CSaeK2Q3VtiSOT7HJfKqenSPBD+yoZR0K7il/i5MECfmIezK3LhjIUCQEdT
3gRoCRx/JRJ72VuNvA/FkShnfVbdtxgGDwb29FwPSdhwB7HppKoajIgwdQYcv8ls
KEUyCAGFNvaWzdKzQPkCQQDgg2NQlICyPAxtgrQc02UcqQVh7EEGP8+qSAWM+jE8
vQwcdUsdzh8pdICOLc25yRvlXf3uQ/9ycUncnrS0aJJG
-----END RSA PRIVATE KEY-----

    EOF
}

resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
	datacenter    = "dal09"
	security_certificate_id = "${ibm_compute_ssl_certificate.test-cert.id}"
}

resource "ibm_lb_service_group" "test_service_group5" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTPS"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
	allocation = 50
}
`

const testAccCheckIBMLbSSLConfig_update = `

resource "ibm_compute_ssl_certificate" "test-cert" {
    certificate = <<EOF
-----BEGIN CERTIFICATE-----
MIICsTCCAhoCCQDl71v1qDksRTANBgkqhkiG9w0BAQUFADCBnDELMAkGA1UEBhMC
SU4xEjAQBgNVBAgMCUJhbmdhbG9yZTESMBAGA1UEBwwJQmFuZ2Fsb3JlMQwwCgYD
VQQKDANJQk0xEjAQBgNVBAsMCVRlcnJhZm9ybTEfMB0GA1UEAwwWd3d3LnRlcnJh
Zm9ybS10ZXN0LmNvbTEiMCAGCSqGSIb3DQEJARYTaGthbnRhcmVAaW4uaWJtLmNv
bTAeFw0xODEwMTkwNjU5MDBaFw0xOTEwMTkwNjU5MDBaMIGcMQswCQYDVQQGEwJJ
TjESMBAGA1UECAwJQmFuZ2Fsb3JlMRIwEAYDVQQHDAlCYW5nYWxvcmUxDDAKBgNV
BAoMA0lCTTESMBAGA1UECwwJVGVycmFmb3JtMR8wHQYDVQQDDBZ3d3cudGVycmFm
b3JtLXRlc3QuY29tMSIwIAYJKoZIhvcNAQkBFhNoa2FudGFyZUBpbi5pYm0uY29t
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDkgLoMrxpPpoYHs08VAVp56xA7
/DNeVawkbYC693tovvO5HkzIJmXAq/QwDRSLyhanIwzGckTnIFsTS/lD0yF9Wq7w
dhWKum87YAqu/2/FtqIxmDQl4Ma55LDtg0GHhLIqLUrCAIvmoZMUuSrh31PrJDGF
c/bDalB6ioCXzlgMwQIDAQABMA0GCSqGSIb3DQEBBQUAA4GBANO8m68Sa3gVrId6
ziJUxLYMzY0i+e6rq/WyEgbne/nBauOENlzpCUp2lPnTncH6LqnOYu2NUA37GYqC
HVFvrAqGy8JzuW/M4vtipCxFIVGkepy/RlnyWJZas4PYgES2bZES+f8or4UwtrcG
g7b8cLkvBhfZkB2X/NgqBqhxBFVA
-----END CERTIFICATE-----
    EOF
    private_key = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDkgLoMrxpPpoYHs08VAVp56xA7/DNeVawkbYC693tovvO5HkzI
JmXAq/QwDRSLyhanIwzGckTnIFsTS/lD0yF9Wq7wdhWKum87YAqu/2/FtqIxmDQl
4Ma55LDtg0GHhLIqLUrCAIvmoZMUuSrh31PrJDGFc/bDalB6ioCXzlgMwQIDAQAB
AoGBAJG0u+5mocJ0jzbN0fm0+TqQ97MoaKEYxEIeSV3vfZQXX1aFybQ/N9caTwVs
8dMJtFQzd2v7ZZB0A19UrMfhE5KlD/pNaNmsZBz3qTqAs9zpU0ZNExGsWyVQgOFw
MyfgOj18q7ZU6Hujc6XknRj3pOYFtLnv0bmaiKObzuxqtsShAkEA8/Fm7ItRmbsu
k3qrKzhYNFx1beT61DTwr6g/BSaczsmIOnFFnAvnqya8L5/kMNRLzx88QrhXo32K
xea8i17T/QJBAO/L9iDOF1RxlpiA9OHc2mH86boyjkeINQZK12Hce0XHtsVgQtem
KCD6VdDXitdx1DgnGpI36BfqrfkxxtDRHRUCQAeAEZkOQ4kFf04bhG3EwrmBaj7h
vnCN3CSaeK2Q3VtiSOT7HJfKqenSPBD+yoZR0K7il/i5MECfmIezK3LhjIUCQEdT
3gRoCRx/JRJ72VuNvA/FkShnfVbdtxgGDwb29FwPSdhwB7HppKoajIgwdQYcv8ls
KEUyCAGFNvaWzdKzQPkCQQDgg2NQlICyPAxtgrQc02UcqQVh7EEGP8+qSAWM+jE8
vQwcdUsdzh8pdICOLc25yRvlXf3uQ/9ycUncnrS0aJJG
-----END RSA PRIVATE KEY-----

    EOF
}

resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
	datacenter    = "dal09"
	security_certificate_id = "${ibm_compute_ssl_certificate.test-cert.id}"
	ssl_offload = true
}

resource "ibm_lb_service_group" "test_service_group5" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTPS"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
	allocation = 50
}
`
