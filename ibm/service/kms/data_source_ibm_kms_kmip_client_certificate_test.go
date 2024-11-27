// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMKMSKMIPClientCertDataSource_basic(t *testing.T) {
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
					WithDataSourceKMSKMIPClientCert(
						"cert_data",
						"null",
						wrapQuotes("myadapter"),
						"ibm_kms_kmip_client_cert.test_cert.cert_id",
						"null",
					),
				),
			},
			// Create Cert Data Source by Name
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
					WithDataSourceKMSKMIPClientCert(
						"cert_data",
						"null",
						wrapQuotes("myadapter"),
						"null",
						"ibm_kms_kmip_client_cert.test_cert.name",
					),
				),
			},
		},
	})
}

func WithDataSourceKMSKMIPClientCert(resourceName, adapterID, adapterName, certID, certName string) CreateResourceOption {
	// Use null for optional args like adapter name
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		data "ibm_kms_kmip_client_cert" "%s" {
			instance_id = ibm_resource_instance.kms_instance.guid
			adapter_id = %s
			adapter_name = %s
			cert_id = %s
			name = %s
		}`, resourceName, adapterID, adapterName, certID, certName)
	}
}
