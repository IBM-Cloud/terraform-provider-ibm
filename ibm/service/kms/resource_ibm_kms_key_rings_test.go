package kms_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSResource_Key_Ring_Name(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsResourceKeyRingConfig(instanceName, keyRing),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_rings.test", "key_ring_id", keyRing),
				),
			},
		},
	})
}

func TestAccIBMKMSResource_Key_Ring_Key(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsResourceKeyRingKeyConfig(instanceName, keyRing, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
				),
			},
		},
	})
}

func TestAccIBMKMSResource_Key_Ring_Not_Exist(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMKmsResourceKeyRingExistConfig(instanceName, keyName, keyRing),
				ExpectError: regexp.MustCompile("KEY_RING_NOT_FOUND_ERR:"),
			},
		},
	})
}

func testAccCheckIBMKmsResourceKeyRingConfig(instanceName, keyRing string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	}
	resource "ibm_kms_key_rings" "test" {
		instance_id = ibm_resource_instance.kms_instance.guid
		key_ring_id = "%s"
	}
`, instanceName, keyRing)
}

func testAccCheckIBMKmsResourceKeyRingKeyConfig(instanceName, keyRing, keyName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	}
	resource "ibm_kms_key_rings" "key_ring" {
		instance_id = ibm_resource_instance.kms_instance.guid
		key_ring_id   = "%s"
	}
	resource "ibm_kms_key" "test" {
		instance_id = ibm_resource_instance.kms_instance.guid
		key_name = "%s"
		key_ring_id = ibm_kms_key_rings.key_ring.key_ring_id
		standard_key =  true
		force_delete = true
	}
`, instanceName, keyRing, keyName)
}

func testAccCheckIBMKmsResourceKeyRingExistConfig(instanceName, keyName, keyRing string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	}
	resource "ibm_kms_key" "test" {
		instance_id = ibm_resource_instance.kms_instance.guid
		key_name = "%s"
		key_ring_id = "%s"
		standard_key =  true
		force_delete = true
	}
`, instanceName, keyRing, keyName)
}
