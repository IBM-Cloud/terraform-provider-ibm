package kms_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSKeyPolicy_basic_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	rotation_interval := 3
	dual_auth_delete := false
	rotation_interval_new := 5
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyPolicyStandardConfigCheck(instanceName, keyName, rotation_interval, dual_auth_delete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "policies.0.rotation.0.interval_month", "3"),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
			{
				Config: testAccCheckIBMKmsKeyPolicyStandardConfigCheck(instanceName, keyName, rotation_interval_new, dual_auth_delete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "policies.0.rotation.0.interval_month", "5"),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMKMSKeyPolicy_rotation_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	rotation_interval := 3
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyPolicyRotationCheck(instanceName, keyName, rotation_interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "policies.0.rotation.0.interval_month", "3"),
				),
			},
		},
	})
}

func TestAccIBMKMSKeyPolicy_dualAuth_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	dual_auth_delete := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyPolicyDualAuthCheck(instanceName, keyName, dual_auth_delete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMKMSKeyPolicy_dualAuth_check_with_Alias(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	aliasName := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	dual_auth_delete := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyPolicyDualAuthCheckWithAlias(instanceName, keyName, aliasName, dual_auth_delete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "alias", aliasName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMKMSKeyPolicy_invalid_interval_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	rotation_interval := 13
	dual_auth_delete := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMKmsKeyPolicyStandardConfig(instanceName, keyName, rotation_interval, dual_auth_delete),
				ExpectError: regexp.MustCompile("must contain a valid int value should be in range(1, 12)"),
			},
		},
	})
}

func testAccCheckIBMKmsKeyPolicyStandardConfigCheck(instanceName, KeyName string, rotation_interval int, dual_auth_delete bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }

	  resource "ibm_kms_key" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_name       = "%s"
		standard_key   = false
	  }
	  resource "ibm_kms_key_policies" "Policy" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_id = ibm_kms_key.test.key_id
		rotation {
			interval_month = %d
		  }
		  dual_auth_delete {
			enabled = %t
		  }
	  }
`, instanceName, KeyName, rotation_interval, dual_auth_delete)
}

func testAccCheckIBMKmsKeyPolicyDualAuthCheck(instanceName, KeyName string, dual_auth_delete bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }

	  resource "ibm_kms_key" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_name       = "%s"
		standard_key   = false
	  }
	  resource "ibm_kms_key_policies" "Policy" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_id = ibm_kms_key.test.key_id
		  dual_auth_delete {
			enabled = %t
		  }
	  }
`, instanceName, KeyName, dual_auth_delete)
}

func testAccCheckIBMKmsKeyPolicyDualAuthCheckWithAlias(instanceName, KeyName string, alias string, dual_auth_delete bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }

	  resource "ibm_kms_key" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_name       = "%s"
		standard_key   = false
	  }

	  resource "ibm_kms_key_alias" "alias_test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		alias = "%s"
		key_id = ibm_kms_key.test.key_id
	  }

	  resource "ibm_kms_key_policies" "Policy" {
		instance_id = ibm_resource_instance.kp_instance.guid
		alias = ibm_kms_key_alias.alias_test.alias
		  dual_auth_delete {
			enabled = %t
		  }
	  }
`, instanceName, KeyName, alias, dual_auth_delete)
}

func testAccCheckIBMKmsKeyPolicyRotationCheck(instanceName, KeyName string, rotation_interval int) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }

	  resource "ibm_kms_key" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_name       = "%s"
		standard_key   = false
	  }
	  resource "ibm_kms_key_policies" "Policy" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_id = ibm_kms_key.test.key_id
		  rotation {
			interval_month = %d
		  }
	  }

`, instanceName, KeyName, rotation_interval)
}
