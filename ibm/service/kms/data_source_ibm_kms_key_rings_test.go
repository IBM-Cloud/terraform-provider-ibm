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

func TestAccIBMKMSKeyRingDataSource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyRingDataSourceConfig(instanceName, keyRing),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_kms_key_rings.test2", "key_rings.1.id", keyRing),
				),
			},
		},
	})
}

func testAccCheckIBMKmsKeyRingDataSourceConfig(instanceName, keyRing string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key_rings" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_ring_id = "%s"
	}
	data "ibm_kms_key_rings" "test2" {
		instance_id = "${ibm_kms_key_rings.test.instance_id}"
	}
`, addPrefixToResourceName(instanceName), keyRing)
}
