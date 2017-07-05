package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMIAMUserPolicyDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_iam_policies", "policies.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_iam_policies", "policies.0.roles.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_iam_policies", "policies.0.resources.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_iam_policies", "policies.0.roles.0.name", "viewer"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicyDataSourceConfig() string {
	return fmt.Sprintf(`
    
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}

data "ibm_account" "testacc_acc" {
    org_guid = "${data.ibm_org.testacc_ds_org.id}"
}
resource "ibm_iam_user_policy" "testacc_iam_policy" {
        account_guid = "${data.ibm_account.testacc_acc.id}"
        ibm_id  = "%s"
        roles   = ["viewer"]
        resources = [{"service_name" = "All Identity and Access enabled services"}]
}
data "ibm_iam_user_policy" "testacc_iam_policies" {
        account_guid = "${data.ibm_account.testacc_acc.id}"
        ibm_id = "${ibm_iam_user_policy.testacc_iam_policy.ibm_id}"
}
`, cfOrganization, IAMUser)

}
