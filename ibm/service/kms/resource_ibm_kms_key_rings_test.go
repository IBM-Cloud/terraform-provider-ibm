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

func TestAccIBMKMSResource_Key_Ring_ForceDeleteFalse(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Create a Key Ring and check force_delete is false
			{
				Config: testAccCheckIBMKmsResourceKeyRingForceDeleteConfig(instanceName, keyName, keyRing, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "force_delete", "false"),
				),
			},
			// Attempt to delete the key ring and remove associated with key
			{
				Config:      testAccCheckIBMKmsResourceDeleteKeyRing(instanceName, keyName),
				ExpectError: regexp.MustCompile("KEY_RING_KEYS_NOT_DELETED_ERR:"),
			},
			// Attempt to delete keys and key rings
			// Check keys are deleted but also expect error with
			{
				Config: testAccCheckIBMKmsResourceNoKeyRingNoKeys(instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("ibm_kms_key.test", "key_name"),
				),
				ExpectError: regexp.MustCompile("KEY_RING_NOT_EMPTY_ERR:"),
			},
		},
	})
}

func TestAccIBMKMSResource_Key_Ring_ForceDeleteTrue(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Create a Key Ring and check force_delete is true
			{
				Config: testAccCheckIBMKmsResourceKeyRingForceDeleteConfig(instanceName, keyName, keyRing, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "force_delete", "true"),
				),
			},
			// Attempt to delete the key ring and remove associated with key
			{
				Config:      testAccCheckIBMKmsResourceDeleteKeyRing(instanceName, keyName),
				ExpectError: regexp.MustCompile("KEY_RING_KEYS_NOT_DELETED_ERR:"),
			},
			// Attempt to delete keys and key rings
			// Check keys are deleted
			{
				Config: testAccCheckIBMKmsResourceNoKeyRingNoKeys(instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("ibm_kms_key.test", "key_name"),
				),
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

func testAccCheckIBMKmsResourceKeyRingForceDeleteConfig(instanceName, keyRing, keyName string, forceDelete bool) string {
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
		force_delete = %t
	}
	resource "ibm_kms_key" "test" {
		instance_id = ibm_resource_instance.kms_instance.guid
		key_name = "%s"
		key_ring_id = ibm_kms_key_rings.key_ring.key_ring_id
		standard_key =  true
		force_delete = true
	}
`, instanceName, keyRing, forceDelete, keyName)
}

func testAccCheckIBMKmsResourceDeleteKeyRing(instanceName, keyName string) string {
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
		standard_key =  true
		force_delete = true
	}
`, instanceName, keyName)
}

func testAccCheckIBMKmsResourceNoKeyRingNoKeys(instanceName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	}
`, instanceName)
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
