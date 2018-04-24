package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMIAMUserPolicy_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
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

func TestAccIBMIAMUserPolicy_Tag(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_tag(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "ibm_id", IAMUser),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "roles.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "resources.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "tags.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyUpdate_tag(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "ibm_id", IAMUser),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "roles.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "resources.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_policy.testacc_iam_policy", "tags.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_ServiceNameEmpty(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMIAMUserPolicy_ServiceNameEmpty(),
				ExpectError: regexp.MustCompile("service_name cannot be empty"),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_InvalidRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMIAMUserPolicy_InvalidRole(),
				ExpectError: regexp.MustCompile(fmt.Sprintf("The given role %q is not valid. Valid roles are", "viewerrole")),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_InvalidUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMIAMUserPolicy_InvalidUser(),
				ExpectError: regexp.MustCompile("does not exist in the account"),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicyDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
	if err != nil {
		return fmt.Errorf("Error checking IAM Policy %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_user_policy" {
			continue
		}

		userID, err := getIBMID(rs.Primary.Attributes["account_guid"], rs.Primary.Attributes["ibm_id"], testAccProvider.Meta())
		if err != nil {
			return err
		}

		_, err = client.IAMPolicy().Get(rs.Primary.Attributes["account_guid"], userID, rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Policy with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
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

func testAccCheckIBMIAMUserPolicy_InvalidRole() string {
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
        roles   = ["viewerrole"]
        resources = [{"service_name" = "All Identity and Access enabled services"}]
}
`, cfOrganization, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_InvalidUser() string {
	return fmt.Sprintf(`
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}

data "ibm_account" "testacc_acc" {
    org_guid = "${data.ibm_org.testacc_ds_org.id}"
}

resource "ibm_iam_user_policy" "testacc_iam_policy" {
        account_guid = "${data.ibm_account.testacc_acc.id}"
        ibm_id  = "sample@example.com"
        roles   = ["viewer"]
        resources = [{"service_name" = "All Identity and Access enabled services"}]
}
`, cfOrganization)
}

func testAccCheckIBMIAMUserPolicy_tag() string {
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
        tags = ["one"]
}
`, cfOrganization, IAMUser)
}

func testAccCheckIBMIAMUserPolicyUpdate_tag() string {
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
        tags = ["one", "two"]
}
`, cfOrganization, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_ServiceNameEmpty() string {
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
        resources = [{"service_name" = ""}]
}
`, cfOrganization, IAMUser)
}
