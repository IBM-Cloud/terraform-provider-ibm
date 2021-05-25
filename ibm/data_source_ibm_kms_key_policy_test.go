package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKmsDataSourceKeyPolicy_basicNew(t *testing.T) {
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
				Config: testAccCheckIBMKmsDataSourceKeyPolicyConfigNew(instanceName, keyName, interval_month, enabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("data.ibm_kms_key_policy.test", "keys.0.policies.0.rotation.0.interval_month", "3"),
					resource.TestCheckResourceAttr("data.ibm_kms_key_policy.test", "keys.0.policies.0.dual_auth_delete.0.enabled", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMKmsDataSourceKeyPolicyConfigNew(instanceName, keyName string, interval_month int, enabled bool) string {
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
	resource "ibm_kms_key_policy" "test2" {
		instance_id = "${ibm_kms_key.test.instance_id}"
		key_id = "ibm_kms_key.test.key_id"
		policies {
			rotation {
				interval_month = %d
			}
			dual_auth_delete {
				enabled = %t
			}
		}
	}
	data "ibm_kms_key_policy" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}"
		key_id = "${ibm_kms_key.test.key_id}"
	}
`, instanceName, keyName, interval_month, enabled)
}
