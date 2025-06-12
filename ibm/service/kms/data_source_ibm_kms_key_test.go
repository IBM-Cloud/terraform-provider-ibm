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

func TestAccIBMKMSKeyDataSource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyDataSourceConfig(instanceName, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
				),
			},
		},
	})
}

func TestAccIBMKMSKeyDataSource_description(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	customDescription := "I am a custom description for the key"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyDataSourceConfigAndDescription(instanceName, keyName, customDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "description", customDescription),
				),
			},
		},
	})
}

func TestAccIBMKMSKeyDataSource_Key(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyDataSourceKeyConfig(instanceName, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_kms_key.test", "keys.0.name", keyName),
					resource.TestCheckResourceAttr("data.ibm_kms_key.test2", "keys.0.name", keyName),
					resource.TestCheckResourceAttr("data.ibm_kms_key.test2", "keys.1.name", keyName),
				),
			},
		},
	})
}

func TestAccIBMKMSKeyDataSourceHPCS_basic(t *testing.T) {
	t.Skip()
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsKeyDataSourceHpcsConfig(acc.HpcsInstanceID, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
				),
			},
		},
	})
}

func TestAccIBMKmsDataSourceKeyPolicy_basic(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	interval_month := 3
	enabled := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsDataSourceKeyPolicyConfig(instanceName, keyName, interval_month, enabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("data.ibm_kms_key.test", "keys.0.policies.0.rotation.0.interval_month", "3"),
					resource.TestCheckResourceAttr("data.ibm_kms_key.test", "keys.0.policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMKmsKeyDataSourceKeyConfig(instanceName, keyName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key" "test2" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "${ibm_kms_key.test.key_name}"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key" "test3" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "${ibm_kms_key.test.key_name}"
		standard_key =  true
		force_delete = true
	}
	data "ibm_kms_key" "test" {
		instance_id = "${ibm_kms_key.test3.instance_id}"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	data "ibm_kms_key" "test2" {
		instance_id = "${ibm_kms_key.test3.instance_id}"
		limit = 2
		key_name = "${ibm_kms_key.test.key_name}"
	}
`, addPrefixToResourceName(instanceName), keyName)
}

func testAccCheckIBMKmsKeyDataSourceConfig(instanceName, keyName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	data "ibm_kms_key" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}"
		key_name = "${ibm_kms_key.test.key_name}"
	}
`, addPrefixToResourceName(instanceName), keyName)
}

func testAccCheckIBMKmsKeyDataSourceConfigAndDescription(instanceName, keyName string, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		description  = "%s"
		force_delete = true
	}
	data "ibm_kms_key" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}"
		key_name = "${ibm_kms_key.test.key_name}"
	}
`, addPrefixToResourceName(instanceName), keyName, description)
}

func testAccCheckIBMKmsKeyDataSourceHpcsConfig(hpcsInstanceID string, KeyName string) string {
	return fmt.Sprintf(`
	  resource "ibm_kms_key" "test" {
		instance_id = "%s"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	data "ibm_kms_key" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}"
		key_name = "${ibm_kms_key.test.key_name}"
	}

`, acc.HpcsInstanceID, KeyName)
}

func testAccCheckIBMKmsDataSourceKeyPolicyConfig(instanceName, keyName string, interval_month int, enabled bool) string {
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
	resource "ibm_kms_key_policies" "testPolicy" {
		instance_id = ibm_kms_key.test.instance_id
		key_id = ibm_kms_key.test.key_id
		rotation {
			interval_month = %d
		}
		dual_auth_delete {
			enabled = %t
		}
	}
	data "ibm_kms_key" "test" {
		instance_id = "${ibm_kms_key_policies.testPolicy.instance_id}"
		key_id = "${ibm_kms_key.test.key_id}"
	}
`, addPrefixToResourceName(instanceName), keyName, interval_month, enabled)
}
