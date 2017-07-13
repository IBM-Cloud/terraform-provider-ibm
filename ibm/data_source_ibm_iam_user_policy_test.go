package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMIAMUserPolicyDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMIAMUserPolicyCheck("data.ibm_iam_user_policy.testacc_iam_policies"),
				),
			},
		},
	})
}

func testAccIBMIAMUserPolicyCheck(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		var (
			policySize int
			err        error
		)

		if policySize, err = strconv.Atoi(a["policies.#"]); err != nil {
			return err
		}
		if policySize < 1 {
			return fmt.Errorf("No policies found")
		}
		return nil
	}
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
