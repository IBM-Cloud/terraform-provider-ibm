// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSKMIPObjectDataSource_basic(t *testing.T) {
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
					WithDataSourceKMSKMIPObject(
						"object_data",
						"ibm_kms_kmip_adapter.test_adapter.adapter_id",
						"null",
						wrapQuotes("8d52da43-06af-4d18-8162-e77a6f60290e"),
					),
				),
				// No way to use API to create a KMIP object, so will just expect a failure
				ExpectError: regexp.MustCompile("KMIP_OBJECT_NOT_FOUND_ERR"),
			},
		},
	})
}

func WithDataSourceKMSKMIPObject(resourceName, adapterID, adapterName, objID string) CreateResourceOption {
	// Use null for optional args like adapter name
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		data "ibm_kms_kmip_object" "%s" {
			instance_id = ibm_resource_instance.kms_instance.guid
			adapter_id = %s
			adapter_name = %s 
			object_id = %s 
		}`, resourceName, adapterID, adapterName, objID)
	}
}
