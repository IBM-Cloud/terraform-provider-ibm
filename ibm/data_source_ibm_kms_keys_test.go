// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMKMSDataSource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMKmsDataSourceConfig(instanceName, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
				),
			},
		},
	})
}
func TestAccIBMKMSHPCSDataSource_basic(t *testing.T) {
	t.Skip()
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMKmsDataSourceHpcsConfig(hpcsInstanceID, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
				),
			},
		},
	})
}

func TestAccIBMKmsDataSourceKeysPolicy_basic(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	interval_month := 3
	enabled := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMKmsDataSourceKeysPolicyConfig(instanceName, keyName, interval_month, enabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("data.ibm_kms_keys.test", "keys.0.policies.0.rotation.0.interval_month", "3"),
					resource.TestCheckResourceAttr("data.ibm_kms_keys.test", "keys.0.policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}
func testAccCheckIBMKmsDataSourceConfig(instanceName, keyName string) string {
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
	data "ibm_kms_keys" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}" 
	}
`, instanceName, keyName)
}

func testAccCheckIBMKmsDataSourceHpcsConfig(hpcsInstanceID, KeyName string) string {
	return fmt.Sprintf(`
	  resource "ibm_kms_key" "test" {
		instance_id = "%s"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	data "ibm_kms_keys" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}" 
	}
	
`, hpcsInstanceID, KeyName)
}

func testAccCheckIBMKmsDataSourceKeysPolicyConfig(instanceName, keyName string, interval_month int, enabled bool) string {
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
		policies {
			rotation {
				interval_month = %d
			}
			dual_auth_delete {
				enabled = %t
			}
		}
	}
	data "ibm_kms_keys" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}" 
	}
`, instanceName, keyName, interval_month, enabled)
}
