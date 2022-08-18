package kms_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKmsDataSourceInstancePolicy_basicNew(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	interval_month := 3
	dadenabled := true
	metricEnable := true
	kciaEnable := true
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsDataSourceInstancePolicyConfigNew(instanceName, interval_month, dadenabled, metricEnable, kciaEnable),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_kms_instance_policies.test", "rotation.0.interval_month", "3"),
					resource.TestCheckResourceAttr("data.ibm_kms_instance_policies.test", "dual_auth_delete.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_kms_instance_policies.test", "metrics.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_kms_instance_policies.test", "key_create_import_access.0.enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMKmsDataSourceInstancePolicyConfigNew(instanceName string, interval_month int, dadenabled, metricEnable, kciaEnable bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	}

	resource "ibm_kms_instance_policies" "test2" {
		instance_id = "${ibm_resource_instance.kp_instance.guid}"
			rotation {
				enabled = true
				interval_month = %d
			}
			dual_auth_delete {
				enabled = %t
			}
			metrics {
				enabled = %t
			}
			key_create_import_access {
				enabled = %t
			}

	}
	data "ibm_kms_instance_policies" "test" {
		instance_id = "${ibm_kms_instance_policies.test2.instance_id}"
	}
`, instanceName, interval_month, dadenabled, metricEnable, kciaEnable)
}
