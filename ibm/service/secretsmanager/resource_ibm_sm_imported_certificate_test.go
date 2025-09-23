// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

var importedCertName = "terraform-test-imported-cert"
var modifiedImportedCertName = "modified-terraform-test--imported-cert"

func TestAccIbmSmImportedCertificateBasic(t *testing.T) {
	resourceName := "ibm_sm_imported_certificate.sm_imported_certificate_basic"

	resource.Test(t, resource.TestCase{

		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmImportedCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: importedCertificateConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "retrieved_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer"),
					resource.TestCheckResourceAttrSet(resourceName, "key_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "expiration_date"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_after"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
					resource.TestCheckResourceAttr(resourceName, "private_key_included", "false"),
					resource.TestCheckResourceAttr(resourceName, "intermediate_included", "false"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "retrieved_at"},
			},
		},
	})
}

func TestAccIbmSmImportedCertificateAllArgs(t *testing.T) {
	resourceName := "ibm_sm_imported_certificate.sm_imported_certificate"

	resource.Test(t, resource.TestCase{

		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmImportedCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: importedCertificateConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmImportedCertificateCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer"),
					resource.TestCheckResourceAttrSet(resourceName, "key_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "expiration_date"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_after"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
					resource.TestCheckResourceAttr(resourceName, "private_key_included", "false"),
					resource.TestCheckResourceAttr(resourceName, "intermediate_included", "false"),
				),
			},
			{
				Config: importedCertificateConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmImportedCertificateUpdated(resourceName),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "2"),
					resource.TestCheckResourceAttr(resourceName, "private_key_included", "true"),
					resource.TestCheckResourceAttr(resourceName, "intermediate_included", "true"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "retrieved_at"},
			},
		},
	})
}

func TestAccIbmSmImportedCertificateManagedCSR(t *testing.T) {
	resourceName := "ibm_sm_imported_certificate.managed_csr"

	resource.Test(t, resource.TestCase{

		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmImportedCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: importedCertificateConfigManagedCSR(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmManagedCsrCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttr(resourceName, "state", "0"),
					resource.TestCheckResourceAttr(resourceName, "state_description", "pre_activation"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "managed_csr.0.csr"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.common_name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.alt_names", "alt1"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.client_flag", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.code_signing_flag", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.organization.0", "IBM"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.country.0", "IL"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.ou.0", "ILSL"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.postal_code.0", "5320047"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.province.0", "DAN"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.locality.0", "Givatayim"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.email_protection_flag", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.exclude_cn_from_sans", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.ext_key_usage", "timestamping"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.ip_sans", "127.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.key_bits", "2048"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.key_type", "rsa"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.key_usage", "DigitalSignature"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.other_sans", "1.3.6.1.4.1.311.21.2.3;utf8:*.example.com"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.policy_identifiers", ""),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.uri_sans", "https://www.example.com/test"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.user_ids", "user"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.require_cn", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.rotate_keys", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.server_flag", "true"),
				),
			},
			{
				Config: mangedCSRConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmManagedCsrUpdated(resourceName),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.common_name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.alt_names", "alt2"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.client_flag", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.code_signing_flag", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.country.0", "USA"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.email_protection_flag", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.exclude_cn_from_sans", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.ext_key_usage", ""),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.ip_sans", ""),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.key_bits", "2048"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.key_type", "rsa"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.key_usage", ""),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.other_sans", ""),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.policy_identifiers", ""),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.uri_sans", ""),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.user_ids", "user"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.require_cn", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.rotate_keys", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_csr.0.server_flag", "false"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "retrieved_at"},
			},
		},
	})
}

var importedCertBasicConfigFormat = `
		resource "ibm_sm_imported_certificate" "sm_imported_certificate_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
			certificate = "-----BEGIN CERTIFICATE-----\r\nMIICsDCCAhmgAwIBAgIJALrogcLQxAOqMA0GCSqGSIb3DQEBCwUAMHExCzAJBgNV\r\nBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQwwCgYD\r\nVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2VydG1n\r\nbXQtZGV2LmNvbTAeFw0xODA0MjUwODM5NTlaFw00NTA5MTAwODM5NTlaMHExCzAJ\r\nBgNVBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQww\r\nCgYDVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2Vy\r\ndG1nbXQtZGV2LmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAmy/4uEEw\r\nAn75rBuAIv5zi+1b2ycUnlw94x3QzYtY3QHQysFu73U3rczVHOsQNd9VIoC0z8py\r\npMZZu7W6dv6cjOSXlpiLfd7Y9TWzO43mNUH0qrnFpSgXM9ZXN3PJWjmTH3yxAsdK\r\nd5wtRdSv9AwrHWo8hHoTumoXYNMDuehyVJ8CAwEAAaNQME4wHQYDVR0OBBYEFMNC\r\nbcvQ+Smn8ikBDrMKhPc4C+f5MB8GA1UdIwQYMBaAFMNCbcvQ+Smn8ikBDrMKhPc4\r\nC+f5MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADgYEAFe2fCmzTcmCHeijV\r\nq0+EOvMRVNF/FTYyjb24gUGTbouZOkfv7JK94lAt/u5mPhpftYX+b1wUlkz0Kyl5\r\n4IgM0XXpcPYDdxQ87c0l/nAUF7Pi++u7CVmJBlclyDOL6AmBpUE0HyquQT4rSp/K\r\n+5qcqSxVjznd5XgQrWQGHLI2tnY=\r\n-----END CERTIFICATE-----"
        }`

var importedCertFullConfigFormat = `
		resource "ibm_sm_imported_certificate" "sm_imported_certificate" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			custom_metadata = %s
			secret_group_id = "default"`

var importedCertWithManagedCsrFormat = `
		resource "ibm_sm_imported_certificate" "managed_csr" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			custom_metadata = %s
			secret_group_id = "default"
			managed_csr %s
        }`

var managedCsrConfigFormat = `{
		alt_names = "alt1"
		client_flag = true
		code_signing_flag = false
		common_name = "example.com"
		country = ["IL"]
		email_protection_flag = false
		exclude_cn_from_sans = false
		ext_key_usage = "timestamping"
		ext_key_usage_oids = "1.3.6.1.5.5.7.3.67"
		ip_sans = "127.0.0.1"
		key_bits = 2048
		key_type = "rsa"
		key_usage = "DigitalSignature"
		locality = ["Givatayim"]
		policy_identifiers = ""
		organization = ["IBM"]
		other_sans = "1.3.6.1.4.1.311.21.2.3;utf8:*.example.com"
		ou = ["ILSL"]
		postal_code = ["5320047"]
		province = ["DAN"]
		require_cn = true
		rotate_keys = false
		server_flag = true
		street_address = ["Ariel Sharon 4"]
		uri_sans = "https://www.example.com/test"
		user_ids = "user"
	}`

var managedCsrModified = `{
		alt_names = "alt2"
		client_flag = false
		code_signing_flag = true
		common_name = "example.com"
		country = ["USA"]
		email_protection_flag = true
		exclude_cn_from_sans = true
		key_bits = 2048
		key_type = "rsa"
		locality = ["Givatayim"]
		policy_identifiers = ""
		organization = ["IBM"]
		ou = ["ILSL"]
		postal_code = ["5320047"]
		province = ["DAN"]
		require_cn = false
		rotate_keys = true
		server_flag = false
		street_address = ["Ariel Sharon 4"]
		user_ids = "user"
	}`

var firstImportedCertData = `
			certificate = "-----BEGIN CERTIFICATE-----\r\nMIICsDCCAhmgAwIBAgIJALrogcLQxAOqMA0GCSqGSIb3DQEBCwUAMHExCzAJBgNV\r\nBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQwwCgYD\r\nVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2VydG1n\r\nbXQtZGV2LmNvbTAeFw0xODA0MjUwODM5NTlaFw00NTA5MTAwODM5NTlaMHExCzAJ\r\nBgNVBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQww\r\nCgYDVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2Vy\r\ndG1nbXQtZGV2LmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAmy/4uEEw\r\nAn75rBuAIv5zi+1b2ycUnlw94x3QzYtY3QHQysFu73U3rczVHOsQNd9VIoC0z8py\r\npMZZu7W6dv6cjOSXlpiLfd7Y9TWzO43mNUH0qrnFpSgXM9ZXN3PJWjmTH3yxAsdK\r\nd5wtRdSv9AwrHWo8hHoTumoXYNMDuehyVJ8CAwEAAaNQME4wHQYDVR0OBBYEFMNC\r\nbcvQ+Smn8ikBDrMKhPc4C+f5MB8GA1UdIwQYMBaAFMNCbcvQ+Smn8ikBDrMKhPc4\r\nC+f5MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADgYEAFe2fCmzTcmCHeijV\r\nq0+EOvMRVNF/FTYyjb24gUGTbouZOkfv7JK94lAt/u5mPhpftYX+b1wUlkz0Kyl5\r\n4IgM0XXpcPYDdxQ87c0l/nAUF7Pi++u7CVmJBlclyDOL6AmBpUE0HyquQT4rSp/K\r\n+5qcqSxVjznd5XgQrWQGHLI2tnY=\r\n-----END CERTIFICATE-----"
		`
var secondImportedCertData = `
			certificate = "-----BEGIN CERTIFICATE-----\r\nMIICsDCCAhmgAwIBAgIJALrogcLQxAOqMA0GCSqGSIb3DQEBCwUAMHExCzAJBgNV\r\nBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQwwCgYD\r\nVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2VydG1n\r\nbXQtZGV2LmNvbTAeFw0xODA0MjUwODM5NTlaFw00NTA5MTAwODM5NTlaMHExCzAJ\r\nBgNVBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQww\r\nCgYDVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2Vy\r\ndG1nbXQtZGV2LmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAmy/4uEEw\r\nAn75rBuAIv5zi+1b2ycUnlw94x3QzYtY3QHQysFu73U3rczVHOsQNd9VIoC0z8py\r\npMZZu7W6dv6cjOSXlpiLfd7Y9TWzO43mNUH0qrnFpSgXM9ZXN3PJWjmTH3yxAsdK\r\nd5wtRdSv9AwrHWo8hHoTumoXYNMDuehyVJ8CAwEAAaNQME4wHQYDVR0OBBYEFMNC\r\nbcvQ+Smn8ikBDrMKhPc4C+f5MB8GA1UdIwQYMBaAFMNCbcvQ+Smn8ikBDrMKhPc4\r\nC+f5MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADgYEAFe2fCmzTcmCHeijV\r\nq0+EOvMRVNF/FTYyjb24gUGTbouZOkfv7JK94lAt/u5mPhpftYX+b1wUlkz0Kyl5\r\n4IgM0XXpcPYDdxQ87c0l/nAUF7Pi++u7CVmJBlclyDOL6AmBpUE0HyquQT4rSp/K\r\n+5qcqSxVjznd5XgQrWQGHLI2tnY=\r\n-----END CERTIFICATE-----"
			intermediate = "-----BEGIN CERTIFICATE-----\r\nMIICsDCCAhmgAwIBAgIJALrogcLQxAOqMA0GCSqGSIb3DQEBCwUAMHExCzAJBgNV\r\nBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQwwCgYD\r\nVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2VydG1n\r\nbXQtZGV2LmNvbTAeFw0xODA0MjUwODM5NTlaFw00NTA5MTAwODM5NTlaMHExCzAJ\r\nBgNVBAYTAnVzMREwDwYDVQQIDAh1cy1zb3V0aDEPMA0GA1UEBwwGRGFsLTEwMQww\r\nCgYDVQQKDANJQk0xEzARBgNVBAsMCkNsb3VkQ2VydHMxGzAZBgNVBAMMEiouY2Vy\r\ndG1nbXQtZGV2LmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAmy/4uEEw\r\nAn75rBuAIv5zi+1b2ycUnlw94x3QzYtY3QHQysFu73U3rczVHOsQNd9VIoC0z8py\r\npMZZu7W6dv6cjOSXlpiLfd7Y9TWzO43mNUH0qrnFpSgXM9ZXN3PJWjmTH3yxAsdK\r\nd5wtRdSv9AwrHWo8hHoTumoXYNMDuehyVJ8CAwEAAaNQME4wHQYDVR0OBBYEFMNC\r\nbcvQ+Smn8ikBDrMKhPc4C+f5MB8GA1UdIwQYMBaAFMNCbcvQ+Smn8ikBDrMKhPc4\r\nC+f5MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADgYEAFe2fCmzTcmCHeijV\r\nq0+EOvMRVNF/FTYyjb24gUGTbouZOkfv7JK94lAt/u5mPhpftYX+b1wUlkz0Kyl5\r\n4IgM0XXpcPYDdxQ87c0l/nAUF7Pi++u7CVmJBlclyDOL6AmBpUE0HyquQT4rSp/K\r\n+5qcqSxVjznd5XgQrWQGHLI2tnY=\r\n-----END CERTIFICATE-----"
			private_key = "-----BEGIN PRIVATE KEY-----\r\nMIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJsv+LhBMAJ++awb\r\ngCL+c4vtW9snFJ5cPeMd0M2LWN0B0MrBbu91N63M1RzrEDXfVSKAtM/KcqTGWbu1\r\nunb+nIzkl5aYi33e2PU1szuN5jVB9Kq5xaUoFzPWVzdzyVo5kx98sQLHSnecLUXU\r\nr/QMKx1qPIR6E7pqF2DTA7noclSfAgMBAAECgYBsFjd3rf+QXXvsQaM3vF4iIYoO\r\n0+NqgPihzUx3PQ0BsZgJAD0SD2ReawIsCBTcUNbtFxPYfjrnRTeOo/5hjujdq0ei\r\nx1PDh4qzDDPRxOdkCHjfMQb/FBNQvhSh+nQsylCm1qZeaOwgqiM8johDvQ8XLaql\r\n/uNcc1kGXHHd7hKQkQJBAMv04YfjtDxdfanrVtjz8Nm3QGklnAgmddRfY9AZB1Vw\r\nT4hpfvmRi0zOXn2KTaVjAcdqp0Irg+IyTQzd+q9dFG0CQQDCyVOEzUfLHotITqPy\r\nzN2EQ/e/YNnfsElBgNbL44V0Gy2vclLBt6hsvJrD0lSXHCo8aWplIvs2cRM/8uv3\r\nim27AkBrgcQTrgoGO72OgJeBumv9RuPzyLhLb4JylGl3eonsFkxF+l3MzVQhAzK5\r\nd9pf0CVS6TwK3AcjhyIoIyYNo8GtAkBUyi6A8Jr/4BvhLdpQJr2Ghc+ijxZIOQSq\r\nbtsRhcjh8bLBXJKJoNi//JmiBDyuSqRYB8s4mzGfUTl/7M6qwqdhAkEAnZEM+ZUV\r\nV0lZA18QsbwYHY1GVmaOi/dpZjS4ECl+7hbqhHfry88bgXzRKaITxe5Tss+lwQQ7\r\ncfLx+EZh+XOvRw==\r\n-----END PRIVATE KEY-----"
		`

func importedCertificateConfigBasic() string {
	return fmt.Sprintf(importedCertBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		importedCertName)
}

func importedCertificateConfigAllArgs() string {
	return fmt.Sprintf(importedCertFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		importedCertName, description, label, customMetadata) +
		firstImportedCertData + `}`
}

// Update metadata + update secret data (rotate) at once
func importedCertificateConfigUpdated() string {
	return fmt.Sprintf(importedCertFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		modifiedImportedCertName, modifiedDescription, modifiedLabel, modifiedCustomMetadata) +
		secondImportedCertData + `}`
}

func importedCertificateConfigManagedCSR() string {
	return fmt.Sprintf(importedCertWithManagedCsrFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		importedCertName, description, label, customMetadata, managedCsrConfigFormat)
}

func mangedCSRConfigUpdated() string {
	return fmt.Sprintf(importedCertWithManagedCsrFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		importedCertName, description, label, customMetadata, managedCsrModified)
}

func testAccCheckIbmSmImportedCertificateCreated(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		importedCertIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := importedCertIntf.(*secretsmanagerv2.ImportedCertificate)

		if err := verifyAttr(*secret.Name, importedCertName, "secret name"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, description, "secret description"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], label, "label"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, customMetadata, "custom metadata"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmImportedCertificateUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		importedCerificateIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := importedCerificateIntf.(*secretsmanagerv2.ImportedCertificate)
		if err := verifyAttr(*secret.Name, modifiedImportedCertName, "secret name after update"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, modifiedDescription, "secret description after update"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels after update: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], modifiedLabel, "label after update"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, modifiedCustomMetadata, "custom metadata after update"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmManagedCsrCreated(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		importedCertIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := importedCertIntf.(*secretsmanagerv2.ImportedCertificate)
		managedCsr := secret.ManagedCsr
		if err := verifyAttr(*managedCsr.CommonName, "example.com", "common_name"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.AltNames, "alt1", "alt_names"); err != nil {
			return err
		}
		if err := verifyAttr(managedCsr.Organization[0], "IBM", "organization"); err != nil {
			return err
		}
		if err := verifyAttr(managedCsr.Ou[0], "ILSL", "ou"); err != nil {
			return err
		}
		if err := verifyAttr(managedCsr.Country[0], "IL", "country"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.ExtKeyUsage, "timestamping", "ext_key_usage"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.IpSans, "127.0.0.1", "ip_sans"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.KeyType, "rsa", "key_type"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*managedCsr.KeyBits), 2048, "key_bits"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.KeyUsage, "DigitalSignature", "key_usage"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.OtherSans, "1.3.6.1.4.1.311.21.2.3;utf8:*.example.com", "other_sans"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.UriSans, "https://www.example.com/test", "uri_sans"); err != nil {
			return err
		}
		if err := verifyAttr(*managedCsr.UserIds, "user", "user_ids"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.ClientFlag, true, "client_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.ServerFlag, true, "server_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.CodeSigningFlag, false, "code_signing_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.EmailProtectionFlag, false, "email_protection_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.ExcludeCnFromSans, false, "exclude_cn_from_sans"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.RequireCn, true, "require_cn"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.RotateKeys, false, "rotate_keys"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmManagedCsrUpdated(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		importedCertIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := importedCertIntf.(*secretsmanagerv2.ImportedCertificate)
		managedCsr := secret.ManagedCsr
		if err := verifyAttr(*managedCsr.AltNames, "alt2", "alt_names"); err != nil {
			return err
		}
		if err := verifyAttr(managedCsr.Country[0], "USA", "country"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.ClientFlag, false, "client_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.ServerFlag, false, "server_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.CodeSigningFlag, true, "code_signing_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.EmailProtectionFlag, true, "email_protection_flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.ExcludeCnFromSans, true, "exclude_cn_from_sans"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.RequireCn, false, "require_cn"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*managedCsr.RotateKeys, true, "rotate_keys"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmImportedCertificateDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_imported_certificate" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("ImportedCertificate still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ImportedCertificate (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
