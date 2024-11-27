package kms_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

func TestAccIBMKMSResource_Key_Ring_AlwaysForceDeleteTrue(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	keyRing := fmt.Sprintf("keyRing%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Create a Key Ring and check force_delete is false
			{
				Config: buildResourceSet(WithResourceKMSInstance(instanceName), WithResourceKMSKeyRing(keyRing, false), WithResourceKMSKey(keyName, "ibm_kms_key_rings.test.key_ring_id")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_ring_id", keyRing),
					resource.TestCheckResourceAttr("ibm_kms_key_rings.test", "force_delete", "false"),
				),
			},
			// Attempt to delete the key ring and key
			{
				Config:      buildResourceSet(WithResourceKMSInstance(instanceName), WithDataKMSKeys()),
				ExpectError: regexp.MustCompile(`\[ERROR\] No keys in instance`),
			},
		},
	})
}

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
