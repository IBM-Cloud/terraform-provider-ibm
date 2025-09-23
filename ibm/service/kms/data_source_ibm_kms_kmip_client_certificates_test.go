// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSKMIPClientCertsDataSource_basic(t *testing.T) {
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
			// Create List Certs Data Source
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
					WithDataSourceKMSKMIPClientCerts(
						"cert_list",
						"null",
						wrapQuotes("myadapter"),
						100,
						0,
						false,
					),
				),
			},
		},
	})
}

func WithDataSourceKMSKMIPClientCerts(resourceName, adapterID, adapterName string, limit, offset int, totalCount bool) CreateResourceOption {
	// Use null for optional args like adapter name
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		data "ibm_kms_kmip_client_certs" "%s" {
			instance_id = ibm_resource_instance.kms_instance.guid
			adapter_id = %s
			adapter_name = %s
			limit = %d
			offset = %d
			show_total_count = %t
		}`, resourceName, adapterID, adapterName, limit, offset, totalCount)
	}
}
