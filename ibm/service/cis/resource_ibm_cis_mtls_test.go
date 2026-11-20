// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccCisMtlsCert = "-----BEGIN CERTIFICATE-----\nMIIEFzCCAv+gAwIBAgIJAMhhsP5Ubtu2MA0GCSqGSIb3DQEBCwUAMIGhMQswCQYD\nVQQGEwJpbjESMBAGA1UECAwJa2FybmF0YWthMRIwEAYDVQQHDAliYW5nYWxvcmUx\nDDAKBgNVBAoMA2libTEMMAoGA1UECwwDY2lzMSowKAYDVQQDDCFtdGxzNy5hdXN0\nZXN0LTEwLmNpc3Rlc3QtbG9hZC5jb20xIjAgBgkqhkiG9w0BCQEWE2RhcnVueWEu\nZC5jQGlibS5jb20wHhcNMjIwNDIyMTEwMzU3WhcNMzIwNDE5MTEwMzU3WjCBoTEL\nMAkGA1UEBhMCaW4xEjAQBgNVBAgMCWthcm5hdGFrYTESMBAGA1UEBwwJYmFuZ2Fs\nb3JlMQwwCgYDVQQKDANpYm0xDDAKBgNVBAsMA2NpczEqMCgGA1UEAwwhbXRsczcu\nYXVzdGVzdC0xMC5jaXN0ZXN0LWxvYWQuY29tMSIwIAYJKoZIhvcNAQkBFhNkYXJ1\nbnlhLmQuY0BpYm0uY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA\n3tjgNpucsvwNFPNWl1DXkWGFLzvdMKDdk3PTAJ3AAYFG4jLVDtZurf3qCLZ8fcz+\nnukYdDKhRZYSP9QvGwDTS4mHOTV/6FAYsb7qfke+V8+v0okmCca07KgTUKFR5F9e\nw1NPYW9yRjoVpy/Kgs983WigDBRQeo50wcLYG7APml0ceqsBKZaXOiTVrf2xDSvd\nNn6Qchgd5dmxiP+drypt7BGIf9j8QlN5HvEETfUQQybwJfq9G6KhNKIKcw+IKGIy\nbI03RmItC+eVhwja/t1UldlXt/L3JduwEkq9QNQe080toAZyaQ/9Vymk80DTrffN\njb1YG224XLlflSSdzbUC0QIDAQABo1AwTjAdBgNVHQ4EFgQUs5QUMLmjPfNutr8U\n2zcjT/yH1pYwHwYDVR0jBBgwFoAUs5QUMLmjPfNutr8U2zcjT/yH1pYwDAYDVR0T\nBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAPCqm4rXm0ptf0iSp+u4X60A3U3ON\ntSpKq5BU1KGF0i5/ZB1ia1we2ORdOzeoNIhoffmRCg/a//Ba5fLRhktzXMcT/zwC\nDVxH9OAtFoj6/rfEko6s+NP/WtWMd7YF1w4wVvK189YWSUDKbE4MijeDLvEfBi3T\nStNu14p4gN8hkSLX/3Rn9ZmI2wDIpqsYRF5KPfvNZ0iIpvJoBWjS6bbVYGd3yNs+\nrXez+Q36oEFfMcM35EEt3qo2EGu4mljqZxhIae5Hy4sKe4c6s0AfpYA4wTQ97cAg\nQ0Sdw3p+PIqPMOcY1sjRLbvPDHGbzc60LvKhHgt/7Cc5ntvxIjJ9ZUt5Ng==\n-----END CERTIFICATE-----\n"

func TestAccIBMCisMtls_Basic(t *testing.T) {
	name := "ibm_cis_mtls." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisMtlsBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "certificate", testAccCisMtlsCert),
					resource.TestCheckResourceAttr(name, "name", "MTLS-Cert"),
					resource.TestCheckResourceAttr(name, "associated_hostnames.#", "0"),
				),
			},
		},
	})
}

func testAccCheckCisMtlsBasic1(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_mtls" "%[1]s" {
		cis_id                    = data.ibm_cis.cis.id
		domain_id                 = data.ibm_cis_domain.cis_domain.domain_id
		certificate               = %[2]q
		name                      = "MTLS-Cert"
		associated_hostnames      = []
	  }
`, id, testAccCisMtlsCert)
}

func TestAccIBMCisMtls_MultiZone(t *testing.T) {
	name := "ibm_cis_mtls.test"
	hostA := "mtls-acc." + acc.CisDomainStatic
	hostB := "mtls-acc." + acc.CisDomainTest
	cert := testAccCisMtlsSelfSignedCA(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisMtlsMultiZone(cert, hostA, hostB),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "zone.0.hostnames.*", hostA),
					resource.TestCheckTypeSetElemAttr(name, "zone.1.hostnames.*", hostB),
					resource.TestCheckResourceAttrSet(name, "fingerprint"),
				),
			},
			{
				Config: testAccCheckCisMtlsMultiZone(cert, hostA, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone.#", "1"),
					resource.TestCheckResourceAttr(name, "zone.0.hostnames.#", "1"),
				),
			},
		},
	})
}

func testAccCheckCisMtlsMultiZone(cert, hostA, hostB string) string {
	zoneB := ""
	if hostB != "" {
		zoneB = fmt.Sprintf(`
		zone {
			domain_id = data.ibm_cis_domain.cis_domain_b.domain_id
			hostnames = [%q]
		}`, hostB)
	}
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_domain" "cis_domain_b" {
		cis_id = data.ibm_cis.cis.id
		domain = %[1]q
	}
	resource "ibm_cis_mtls" "test" {
		cis_id      = data.ibm_cis.cis.id
		certificate = %[2]q
		name        = "MTLS-Cert-MultiZone"
		zone {
			domain_id = data.ibm_cis_domain.cis_domain.domain_id
			hostnames = [%[3]q]
		}%[4]s
	}
`, acc.CisDomainTest, cert, hostA, zoneB)
}

// testAccCisMtlsSelfSignedCA returns a fresh CA, since CIS rejects a PEM that is already uploaded to the instance.
func testAccCisMtlsSelfSignedCA(t *testing.T) string {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 64))
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{
		SerialNumber:          serial,
		Subject:               pkix.Name{CommonName: "tf-acc-mtls-" + serial.Text(16)},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
}
