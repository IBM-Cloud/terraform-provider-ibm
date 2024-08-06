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
				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, false)),
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
				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, false), WithResourceKMSKey(keyName, "ibm_kms_key_rings.test.key_ring_id")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
				),
			},
			// Cleanup: Change force_delete to true to allow for cleanup
			{
				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, true), WithResourceKMSKey(keyName, "ibm_kms_key_rings.test.key_ring_id")),
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
				Config:      buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKey(keyName, keyRing)),
				ExpectError: regexp.MustCompile("KEY_RING_NOT_FOUND_ERR:"),
			},
		},
	})
}

// Developer note: Test is disabled as a bug exists where this is not properly testable
// func TestAccIBMKMSResource_Key_Ring_ForceDeleteFalse(t *testing.T) {
// 	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
// 	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
// 	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { acc.TestAccPreCheck(t) },
// 		Providers: acc.TestAccProviders,
// 		Steps: []resource.TestStep{
// 			// Create a Key Ring and check force_delete is false
// 			{
// 				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, false), WithResourceKMSKey(keyName, "ibm_kms_key_rings.test.key_ring_id")),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
// 					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
// 					resource.TestCheckResourceAttr("ibm_kms_key_rings.test", "force_delete", "false"),
// 				),
// 			},
// 			// Developer note: We cannot move key rings to default key ring as we have not implemented that PATCH endpoint in terraform. Therefore we must depend on the force_delete flag to clean up test cases
// 			// Attempt to delete the key ring and key
// 			{
// 				Config:      buildResourceSet(WithResourceKMSInstance(instanceName)),
// 				ExpectError: regexp.MustCompile("KEY_RING_NOT_EMPTY_ERR:"),
// 			},
// 			// Update key ring to force_delete for cleanup
// 			{
// 				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, true)),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_kms_key_rings.test", "force_delete", "true"),
// 				),
// 			},
// 			// Delete Key Ring
// 			{
// 				Config:      buildResourceSet(WithResourceKMSInstance(instanceName), WithDataKMSKeys()),
// 				ExpectError: regexp.MustCompile(`\[ERROR\] No keys in instance`),
// 			},
// 			// Developer note: There is no support for listing keys under a certain key state so we cannot verify deleted key is now in default key ring
// 		},
// 	})
// }

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
				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, true), WithResourceKMSKey(keyName, "ibm_kms_key_rings.test.key_ring_id")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
					resource.TestCheckResourceAttr("ibm_kms_key_rings.test", "force_delete", "true"),
				),
			},
			// Attempt to delete the key ring and key
			{
				Config: buildResourceSet(WithResourceKMSInstance(instanceName)),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config:      buildResourceSet(WithResourceKMSInstance(instanceName), WithDataKMSKeys()),
				ExpectError: regexp.MustCompile(`\[ERROR\] No keys in instance`),
			},
			{
				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithDataKMSKeyRings()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_kms_key_rings.test_key_rings", "key_rings.0.id", "default"),
				),
			},
		},
	})
}

// Developer note: Test is disabled as a bug exists where this is not properly testable
// func TestAccIBMKMSResource_Key_Ring_ForceDeleteTrueContainsActiveKeys(t *testing.T) {
// 	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
// 	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
// 	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { acc.TestAccPreCheck(t) },
// 		Providers: acc.TestAccProviders,
// 		Steps: []resource.TestStep{
// 			// Create a Key Ring and check force_delete is true
// 			{
// 				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, true), WithResourceKMSKey(keyName, "ibm_kms_key_rings.test.key_ring_id")),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
// 					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
// 					resource.TestCheckResourceAttr("ibm_kms_key_rings.test", "force_delete", "true"),
// 				),
// 			},
// 			// Attempt to delete the key ring while active key exists
// 			// We must specify key ring ID and not reference here as the resource is removed
// 			{
// 				Config:      buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKey(keyName, keyRing)),
// 				ExpectError: regexp.MustCompile("KEY_RING_KEYS_NOT_DELETED_ERR:"),
// 			},
// 			// Attempt to delete keys
// 			{
// 				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, true)),
// 			},
// 			// Attempt to delete key ring and check no more keys
// 			{
// 				Config:      buildResourceSet(WithResourceKMSInstance(instanceName), WithDataKMSKeys()),
// 				ExpectError: regexp.MustCompile(`\[ERROR\] No keys in instance`),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("data.ibm_kms_key_rings.test_key_rings", "key_rings.0.id", "default"),
// 				),
// 			},
// 		},
// 	})
// }

type CreateResourceOption func(resourceText *string)

func buildResourceSet(options ...CreateResourceOption) string {
	var fullResourceSet *string
	emptyString := ""
	fullResourceSet = &emptyString
	for _, opt := range options {
		opt(fullResourceSet)
		*fullResourceSet += "\n"
	}
	return *fullResourceSet
}

func WithResourceKMSInstance(instanceName string) CreateResourceOption {
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		resource "ibm_resource_instance" "kms_instance" {
			name              = "%s"
			service           = "kms"
			plan              = "tiered-pricing"
			location          =  "us-south"
		}`, addPrefixToResourceName(instanceName))
	}
}

func WithResourceKMSKeyRing(keyRing string, forceDelete bool) CreateResourceOption {
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		resource "ibm_kms_key_rings" "test" {
			instance_id = ibm_resource_instance.kms_instance.guid
			key_ring_id = "%s"
			force_delete = %t
		}`, keyRing, forceDelete)
	}
}

func WithResourceKMSKey(keyName string, keyRing string) CreateResourceOption {
	if keyRing != "ibm_kms_key_rings.test.key_ring_id" {
		keyRing = `"` + keyRing + `"`
	}
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		resource "ibm_kms_key" "test" {
			instance_id = ibm_resource_instance.kms_instance.guid
			key_name = "%s"
			key_ring_id = %s
			standard_key =  true
			force_delete = true
		}`, keyName, keyRing)
	}
}

func WithDataKMSKeys() CreateResourceOption {
	return func(resources *string) {
		*resources += `
		data "ibm_kms_keys" "test_keys" {
			instance_id = "${ibm_resource_instance.kms_instance.guid}"
   		}`
	}
}

func WithDataKMSKeyRings() CreateResourceOption {
	return func(resources *string) {
		*resources += `
		data "ibm_kms_key_rings" "test_key_rings" {
			instance_id = "${ibm_resource_instance.kms_instance.guid}"
   		}`
	}
}

func WithResourceKMSKeyAlias(tfConfigId string, alias string, tfConfigKeyId string) CreateResourceOption {
	return func(resources *string) {
		*resources += fmt.Sprintf(`
		resource "ibm_kms_key_alias" "%s" {
			instance_id = "${ibm_resource_instance.kms_instance.guid}"
			alias = "%s"
			key_id = "${%s}"
		}`, tfConfigId, alias, tfConfigKeyId)
	}
}
