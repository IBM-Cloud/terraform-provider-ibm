package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccIBMIAMUserPolicy_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "ibm_id", IAMUser),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "roles.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "resources.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_update(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "ibm_id", IAMUser),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "roles.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "resources.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicy_basic() string {
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
`, cfOrganization, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_update() string {
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
        roles   = ["viewer","administrator"]
        resources =  [{"service_name" = "All Identity and Access enabled services"}]
}
`, cfOrganization, IAMUser)
}
