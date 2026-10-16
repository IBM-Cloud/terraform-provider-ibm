// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package registry_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccIBMCrPlanUpgradeActionBasic triggers the action after create and checks
// the account ends on the Standard plan. Running it twice is safe: the action is a
// no-op when the plan already matches.
func TestAccIBMCrPlanUpgradeActionBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccCrPlanUpgradeActionConfig(),
				Check:  testAccCheckCrPlan("Standard"),
			},
		},
	})
}

func testAccCrPlanUpgradeActionConfig() string {
	return `
		action "ibm_cr_plan_upgrade" "test" {
			config {
				plan = "Standard"
			}
		}

		resource "terraform_data" "trigger" {
			input = "cr-plan-upgrade"

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_cr_plan_upgrade.test]
				}
			}
		}
	`
}

func testAccCheckCrPlan(want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerRegistryV1()
		if err != nil {
			return err
		}
		plan, _, err := client.GetPlans(client.NewGetPlansOptions())
		if err != nil {
			return err
		}
		if plan == nil || plan.Plan == nil || *plan.Plan != want {
			return fmt.Errorf("container registry plan is %v, want %s", plan, want)
		}
		return nil
	}
}
