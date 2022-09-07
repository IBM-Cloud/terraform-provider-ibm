package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/scc-go-sdk/v4/posturemanagementv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMSccPostureProfileImportBasic(t *testing.T) {
	name := "ibm_scc_posture_profile_import." + "profiles"
	file := "../../test-fixtures/import_profile.csv"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSccPostureProfileImportConfigBasic(file),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "file", file),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "type", "custom"),
					testAccCheckSccPostureProfileImportRemoveImportedRecords(name),
				),
			},
			{
				ResourceName: name,
				ImportState:  true,
			},
		},
	})
}

func testAccCheckSccPostureProfileImportConfigBasic(file string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_posture_profile_import" "profiles" {
			file      = "%[1]s"
		}`, file)
}

func testAccCheckSccPostureProfileImportRemoveImportedRecords(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		postureManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PostureManagementV2()
		if err != nil {
			return err
		}
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}
		deleteProfileOptions := &posturemanagementv2.DeleteProfileOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		accountID := userDetails.UserAccount
		deleteProfileOptions.SetAccountID(accountID)

		deleteProfileOptions.SetID(rs.Primary.ID)
		_, err = postureManagementClient.DeleteProfile(deleteProfileOptions)
		if err != nil {
			return err
		}
		return nil
	}
}
