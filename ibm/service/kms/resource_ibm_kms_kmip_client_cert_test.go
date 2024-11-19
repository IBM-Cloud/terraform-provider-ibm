// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"regexp"
	"strings"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSKMIPClientCertResource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	myCert, err := generateSelfSignedCertificate()
	if err != nil {
		t.Error(err)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Create a CRK and an adapter
			{
				Config: buildResourceSet(
					WithResourceKMSInstance(instanceName),
					WithResourceKMSRootKey("adapter_test_crk", "TestCRK"),
					WithResourceKMSKMIPAdapter(
						"test_adapter",
						"native_1.0",
						convertMapToTerraformConfigString(map[string]string{
							wrapQuotes("crk_id"): "ibm_kms_key.adapter_test_crk.key_id",
						}),
						wrapQuotes("myadapter"),
						"null",
					),
					WithResourceKMSKMIPClientCert(
						"test_cert",
						"ibm_kms_kmip_adapter.test_adapter.adapter_id",
						myCert,
						wrapQuotes("mycert"),
					),
				),
			},
		},
	})
}

func TestAccIBMKMSKMIPClientCertResource_DuplicateCertError(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	myCert, err := generateSelfSignedCertificate()
	if err != nil {
		t.Error(err)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Create a CRK and an adapter
			{
				Config: buildResourceSet(
					WithResourceKMSInstance(instanceName),
					WithResourceKMSRootKey("adapter_test_crk", "TestCRK"),
					WithResourceKMSKMIPAdapter(
						"test_adapter",
						"native_1.0",
						convertMapToTerraformConfigString(map[string]string{
							wrapQuotes("crk_id"): "ibm_kms_key.adapter_test_crk.key_id",
						}),
						wrapQuotes("myadapter"),
						"null",
					),
					WithResourceKMSKMIPClientCert(
						"test_cert",
						"ibm_kms_kmip_adapter.test_adapter.adapter_id",
						myCert,
						wrapQuotes("mycert"),
					),
					WithResourceKMSKMIPClientCert(
						"test_cert_dupe",
						"ibm_kms_kmip_adapter.test_adapter.adapter_id",
						myCert,
						"null",
					),
				),
				ExpectError: regexp.MustCompile("KMIP_CERT_DUPLICATE_ERR"),
			},
		},
	})
}

func TestAccIBMKMSKMIPClientCertResource_InvalidCert(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Create a CRK and an adapter
			{
				Config: buildResourceSet(
					WithResourceKMSInstance(instanceName),
					WithResourceKMSRootKey("adapter_test_crk", "TestCRK"),
					WithResourceKMSKMIPAdapter(
						"test_adapter",
						"native_1.0",
						convertMapToTerraformConfigString(map[string]string{
							wrapQuotes("crk_id"): "ibm_kms_key.adapter_test_crk.key_id",
						}),
						wrapQuotes("myadapter"),
						"null",
					),
					WithResourceKMSKMIPClientCert(
						"test_cert",
						"ibm_kms_kmip_adapter.test_adapter.adapter_id",
						"invalidPEM",
						wrapQuotes("mycert"),
					),
				),
				ExpectError: regexp.MustCompile("INVALID_FIELD_ERR"),
			},
		},
	})
}

func TestAccIBMKMSKMIPClientCertResource_DuplicateNameError(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	myCert, err := generateSelfSignedCertificate()
	if err != nil {
		t.Error(err)
	}
	myCert2, err := generateSelfSignedCertificate()
	if err != nil {
		t.Error(err)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Create a CRK and an adapter
			{
				Config: buildResourceSet(
					WithResourceKMSInstance(instanceName),
					WithResourceKMSRootKey("adapter_test_crk", "TestCRK"),
					WithResourceKMSKMIPAdapter(
						"test_adapter",
						"native_1.0",
						convertMapToTerraformConfigString(map[string]string{
							wrapQuotes("crk_id"): "ibm_kms_key.adapter_test_crk.key_id",
						}),
						wrapQuotes("myadapter"),
						"null",
					),
					WithResourceKMSKMIPClientCert(
						"test_cert",
						"ibm_kms_kmip_adapter.test_adapter.adapter_id",
						myCert,
						wrapQuotes("mycert"),
					),
					WithResourceKMSKMIPClientCert(
						"test_cert_dupe",
						"ibm_kms_kmip_adapter.test_adapter.adapter_id",
						myCert2,
						wrapQuotes("mycert"),
					),
				),
				ExpectError: regexp.MustCompile("KMIP_CERT_DUPLICATE_NAME_ERR"),
			},
		},
	})
}

func WithResourceKMSKMIPClientCert(resourceName, adapterID, certificate, certName string) CreateResourceOption {
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		resource "ibm_kms_kmip_client_cert" "%s" {
			instance_id = ibm_resource_instance.kms_instance.guid
			adapter_id = %s
			certificate = "%s"
			name = %s
		}`, resourceName, adapterID, certificate, certName)
	}
}

func generateSelfSignedCertificate() (string, error) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"IBM"},
			Country:      []string{"US"},
			Province:     []string{"TX"},
			Locality:     []string{"Austin"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0), // Expires 10 years from now
		SubjectKeyId: []byte{1, 2, 3, 4, 5, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &certPrivateKey.PublicKey, certPrivateKey)
	if err != nil {
		return "", err
	}

	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	certString := certPEM.String()
	certString = strings.Replace(certString, "\n", "\\n", -1)
	return certString, nil
}
