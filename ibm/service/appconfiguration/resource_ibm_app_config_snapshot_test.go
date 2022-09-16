package appconfiguration_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccIbmIbmAppConfigSnapshotBasic(t *testing.T) {
	var conf appconfigurationv1.SnapshotResponseGetApi
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	git_config_name := fmt.Sprintf("tf_git_config_name_%d", acctest.RandIntRange(10, 100))
	git_config_id := fmt.Sprintf("tf_git_config_id_%d", acctest.RandIntRange(10, 100))
	git_url := fmt.Sprintf("tf_git_url_%d", acctest.RandIntRange(10, 100))
	git_branch := fmt.Sprintf("tf_git_branch_%d", acctest.RandIntRange(10, 100))
	git_file_path := fmt.Sprintf("tf_git_file_path_%d", acctest.RandIntRange(10, 100))
	git_token := fmt.Sprintf("tf_git_token_%d", acctest.RandIntRange(10, 100))
	collection_id := fmt.Sprintf("tf_collection_id_%d", acctest.RandIntRange(10, 100))
	git_config_nameUpdate := fmt.Sprintf("tf_git_config_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigSnapshotConfigBasic(instanceName, git_config_name, git_config_id, git_url, git_branch, git_file_path,
					git_token, collection_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigSnapshotExists("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "environment_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "property_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "type"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "description"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "tags"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "created_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "updated_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "href"),
					resource.TestCheckResourceAttrSet("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "segment_exists"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigSnapshotConfigBasic(instanceName, git_config_nameUpdate, git_config_id, git_url, git_branch, git_file_path,
					git_token, collection_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_snapshot.ibm_app_config_snapshot_resource1", "name", git_config_nameUpdate),
				),
			},
		},
	})

}

func testAccCheckIbmAppConfigSnapshotConfigBasic(instanceName, git_config_id, git_config_name, git_url, git_branch, git_file_path,
	git_token, collection_id string) string {
	return fmt.Sprintf(`
    	resource "ibm_resource_instance" "app_config_terraform_test476" {
    		name     = "%s"
    		location = "us-south"
    		service  = "apprapp"
    		plan     = "lite"
    	}
    	resource "ibm_app_config_snapshot" "ibm_app_config_property_resource1" {
    		guid 		= ibm_resource_instance.app_config_terraform_test476.guid
    		git_config_id = "%s"
    		git_config_name = "%s"
    		git_url = "%s"
    		git_branch = "%s"
    		git_file_path = "%s"
            git_token = "%s"
    		collection_id = "%s"
    		environment_id = "dev"
    	}`, instanceName, git_config_id, git_config_name, git_url, git_branch, git_file_path,
		git_token, collection_id)
}

func testAccCheckIbmAppConfigSnapshotExists(n string, obj appconfigurationv1.SnapshotResponseGetApi) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}
		options := &appconfigurationv1.GetGitconfigOptions{}

		options.SetGitConfigID(parts[1])

		snapshot, _, err := appconfigClient.GetGitconfig(options)
		if err != nil {
			return err
		}

		obj = *snapshot
		return nil
	}
}

func testAccCheckIbmAppConfigSnapshotDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_config_snapshot" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}
		options := &appconfigurationv1.GetGitconfigOptions{}

		options.SetGitConfigID(parts[1])

		// Try to find the key
		_, response, err := appconfigClient.GetGitconfig(options)

		if err == nil {
			return fmt.Errorf("app_config_snapshot still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for app_config_snapshot (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
