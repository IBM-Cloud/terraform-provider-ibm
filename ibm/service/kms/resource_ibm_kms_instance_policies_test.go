package kms_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMKMSInstancePolicy_basic_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	rotation_interval := 3
	dual_auth_delete := false
	rotation_interval_new := 5
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsInstancePolicyStandardConfigCheck(instanceName, rotation_interval, dual_auth_delete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "rotation.0.interval_month", "3"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "dual_auth_delete.0.enabled", "false"),
				),
			},
			{
				Config: testAccCheckIBMKmsInstancePolicyStandardConfigCheck(instanceName, rotation_interval_new, dual_auth_delete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "rotation.0.interval_month", "5"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMKMSInstancePolicy_rotation_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	rotation_interval := 3
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsInstancePolicyRotationCheck(instanceName, rotation_interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "rotation.0.interval_month", "3"),
				),
			},
		},
	})
}

func TestAccIBMKMSInstancePolicy_dualAuth_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	dual_auth_delete := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsInstancePolicyDualAuthCheck(instanceName, dual_auth_delete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMKMSInstancePolicy_metrics_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	metrics := true
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsInstancePolicyMetricCheck(instanceName, metrics),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "metrics.0.enabled", "true"),
				),
			},
		},
	})
}

func TestAccIBMKMSInstancePolicy_kcia_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	enable := true
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsInstancePolicyKciaCheck(instanceName, enable),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.enabled", "true"),
					// defaults below
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.create_root_key", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.create_standard_key", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.import_root_key", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.import_standard_key", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.enforce_token", "false"),
				),
			},
		},
	})
}

func TestAccIBMKMSInstancePolicy_kcia_attributes_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	enable := true
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsInstancePolicyKciaWithAtttributesCheck(instanceName, enable),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.enabled", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.create_root_key", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.create_standard_key", "false"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.import_root_key", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.import_standard_key", "false"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "key_create_import_access.0.enforce_token", "false"),
				),
			},
		},
	})
}

func TestAccIBMKMSInstancePolicyWithKey(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	rotation_interval := 3
	metrics := true
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsInstancePolicyWithKey(instanceName, metrics, rotation_interval, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "metrics.0.enabled", "true"),
					resource.TestCheckResourceAttr("ibm_kms_instance_policies.test", "rotation.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_kms_key_policies.test2", "policies.0.rotation.0.interval_month", "3"),
				),
			},
		},
	})
}

func TestAccIBMKMSInstancePolicy_invalid_interval_check(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	rotation_interval := 13
	dual_auth_delete := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMKmsInstancePolicyStandardConfigCheck(instanceName, rotation_interval, dual_auth_delete),
				ExpectError: regexp.MustCompile(`.*must contain a valid int value should be in range\(1, 12\).*`),
			},
		},
	})
}

func testAccCheckIBMKmsInstancePolicyStandardConfigCheck(instanceName string, rotation_interval int, dual_auth_delete bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  resource "ibm_kms_instance_policies" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		rotation {
			enabled = true
			interval_month = %d
		  }
		  dual_auth_delete {
			enabled = %t
		  }
	  }
`, addPrefixToResourceName(instanceName), rotation_interval, dual_auth_delete)
}

func testAccCheckIBMKmsInstancePolicyDualAuthCheck(instanceName string, dual_auth_delete bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  resource "ibm_kms_instance_policies" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		  dual_auth_delete {
			enabled = %t
		  }
	  }
`, addPrefixToResourceName(instanceName), dual_auth_delete)
}

func testAccCheckIBMKmsInstancePolicyWithKey(instanceName string, metrics bool, rotation int, keyName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  resource "ibm_kms_instance_policies" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		  metrics {
			  enabled = %t
		  }
		  rotation {
			enabled = true
			interval_month = %d
		  }
		}
		resource "ibm_kms_key" "test" {
			instance_id = ibm_kms_instance_policies.test.instance_id
			key_name       = "%s"
			standard_key   = false
		}
		data "ibm_kms_key_policies" "test2" {
			instance_id = ibm_kms_key.test.instance_id
			key_id = ibm_kms_key.test.key_id
		}
`, addPrefixToResourceName(instanceName), metrics, rotation, keyName)

}

func testAccCheckIBMKmsInstancePolicyRotationCheck(instanceName string, rotation_interval int) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  resource "ibm_kms_instance_policies" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		  rotation {
			enabled = true
			interval_month = %d
		  }
	  }

`, addPrefixToResourceName(instanceName), rotation_interval)
}

func testAccCheckIBMKmsInstancePolicyMetricCheck(instanceName string, metric bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  resource "ibm_kms_instance_policies" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		  metrics {
			enabled = %t
		  }
	  }
`, addPrefixToResourceName(instanceName), metric)
}

func testAccCheckIBMKmsInstancePolicyKciaCheck(instanceName string, kcia bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  resource "ibm_kms_instance_policies" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		  key_create_import_access {
			enabled = %t
		  }
	  }
`, addPrefixToResourceName(instanceName), kcia)
}

func testAccCheckIBMKmsInstancePolicyKciaWithAtttributesCheck(instanceName string, kcia bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  resource "ibm_kms_instance_policies" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		  key_create_import_access {
			enabled = %t
			create_root_key     = true
			create_standard_key = false
			import_root_key     = true
			import_standard_key = false
			enforce_token       = false
		  }
	  }
`, addPrefixToResourceName(instanceName), kcia)
}
