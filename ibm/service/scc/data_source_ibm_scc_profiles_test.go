package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfilesDataSourceConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "profiles.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "profiles.0.profile_name"),
				),
			},
		},
	})
}

func TestAccIbmSccProfilesDataSourceAllArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfilesDataSourceConfigAllArgs(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "profiles.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profiles.scc_profiles_instance", "profiles.0.profile_name"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProfilesDataSourceConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_profiles" "scc_profiles_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIbmSccProfilesDataSourceConfigAllArgs(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_profiles" "scc_profiles_instance" {
			instance_id = "%s"
      profile_type = "predefined"
		}
	`, instanceID)
}
