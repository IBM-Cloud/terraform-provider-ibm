package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMIAMRoleDataSourceAction_basic(t *testing.T) {
	serviceName := "kms"
	name := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMRoleActionConfig(name, displayName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_role_actions.test", "service", serviceName),
				),
			},
		},
	})
}

func testAccCheckIBMIAMRoleActionConfig(name, displayName, serviceName string) string {
	return fmt.Sprintf(`

data "ibm_iam_role_actions" "test" {
  service = "%s"
}

resource "ibm_iam_custom_role" "customrole" {
    name         = "%s"
    display_name = "%s"
    description  = "Custom Role for test scenario2"
    service = "kms"
    actions      = [data.ibm_iam_role_actions.test.manager.18]
}
`, serviceName, name, displayName)
}
