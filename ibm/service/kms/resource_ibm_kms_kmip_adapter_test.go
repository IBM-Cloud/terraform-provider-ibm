// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSKMIPAdapterResource_basic(t *testing.T) {
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
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_kmip_adapter.test_adapter", "name", "myadapter"),
				),
			},
		},
	})
}

func WithResourceKMSKMIPAdapter(resourceName, profile, profileData, adapterName, description string) CreateResourceOption {
	// Use null for optional args like adapter name
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		resource "ibm_kms_kmip_adapter" "%s" {
			instance_id = ibm_resource_instance.kms_instance.guid
			profile = "%s"
			profile_data = %s
			name = %s
			description = %s
		}`, resourceName, profile, profileData, adapterName, description)
	}
}

func WithResourceKMSRootKey(resourceName, keyName string) CreateResourceOption {
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		resource "ibm_kms_key" "%s" {
			instance_id = ibm_resource_instance.kms_instance.guid
			key_name = "%s"
			standard_key = false
			force_delete = true
		}`, resourceName, keyName)
	}
}
