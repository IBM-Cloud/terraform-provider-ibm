package kms_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMKMSDataSourceKeyPolicy_basicNew(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	rotationEnable := true
	interval_month := 3
	enabled := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsDataSourceKeyPolicyConfigNew(instanceName, keyName, rotationEnable, interval_month, enabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("data.ibm_kms_key_policies.test", "policies.0.rotation.0.interval_month", "3"),
					resource.TestCheckResourceAttr("data.ibm_kms_key_policies.test", "policies.0.rotation.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_kms_key_policies.test", "policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMKmsDataSourceKeyPolicyConfigNew(instanceName, keyName string, rotationEnable bool, interval_month int, enabled bool) string {
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
	resource "ibm_kms_key_policies" "test2" {
		instance_id = "${ibm_kms_key.test.instance_id}"
		key_id = "${ibm_kms_key.test.key_id}"
			rotation {
				enabled = %t
				interval_month = %d
			}
			dual_auth_delete {
				enabled = %t
			}
	}
	data "ibm_kms_key_policies" "test" {
		instance_id = "${ibm_kms_key_policies.test2.instance_id}"
		key_id = "${ibm_kms_key_policies.test2.key_id}"
	}
`, addPrefixToResourceName(instanceName), keyName, rotationEnable, interval_month, enabled)
}
