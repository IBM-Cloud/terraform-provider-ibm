// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsShareSnapshotBasic(t *testing.T) {
	var conf vpcv1.ShareSnapshot
	shareName := fmt.Sprintf("tf-name-share%d", acctest.RandIntRange(10, 100))
	shareSnapshotName := fmt.Sprintf("tf-name-share-snap%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsShareSnapshotDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsShareSnapshotConfigBasic(shareName, shareSnapshotName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsShareSnapshotExists("ibm_is_share_snapshot.is_share_snapshot_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_is_share_snapshot.is_share_snapshot_instance", "share"),
					resource.TestCheckResourceAttrSet("ibm_is_share_snapshot.is_share_snapshot_instance", "tags.#"),
					resource.TestCheckResourceAttr("ibm_is_share_snapshot.is_share_snapshot_instance", "name", shareSnapshotName),
				),
			},
		},
	})
}

func TestAccIBMIsShareSnapshotAllArgs(t *testing.T) {
	var conf vpcv1.ShareSnapshot
	name := fmt.Sprintf("tf-name-share-snapshot%d", acctest.RandIntRange(10, 100))
	shareName := fmt.Sprintf("tf-name-share%d", acctest.RandIntRange(10, 100))
	// nameUpdate := fmt.Sprintf("tf-name-share-snapshot%d", acctest.RandIntRange(10, 100))
	userTag1 := fmt.Sprintf("tfp-share-tag-%d", acctest.RandIntRange(10, 100))
	userTag2 := fmt.Sprintf("tfp-share-tag-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsShareSnapshotDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsShareSnapshotConfig(shareName, name, userTag1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsShareSnapshotExists("ibm_is_share_snapshot.is_share_snapshot_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_share_snapshot.is_share_snapshot_instance", "name", name),
					resource.TestCheckResourceAttrSet("ibm_is_share_snapshot.is_share_snapshot_instance", "tags.#"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsShareSnapshotConfig(shareName, name, userTag2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_share_snapshot.is_share_snapshot_instance", "name", name),
					resource.TestCheckResourceAttrSet("ibm_is_share_snapshot.is_share_snapshot_instance", "tags.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsShareSnapshotConfigBasic(shareName, shareSnapshotName string) string {
	return testAccCheckIbmIsShareConfigBasic(shareName) + fmt.Sprintf(`
		resource "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share = ibm_is_share.is_share.id
			name = "%s"
		}
	`, shareSnapshotName)
}

func testAccCheckIBMIsShareSnapshotConfig(shareName string, name, tags string) string {
	return testAccCheckIbmIsShareConfigBasic(shareName) + fmt.Sprintf(`
		resource "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share = ibm_is_share.is_share.id
			name = "%s"
			tags = ["%s"]
		}
	`, name, tags)
}

func testAccCheckIBMIsShareSnapshotExists(n string, obj vpcv1.ShareSnapshot) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getShareSnapshotOptions := &vpcv1.GetShareSnapshotOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getShareSnapshotOptions.SetShareID(parts[0])
		getShareSnapshotOptions.SetID(parts[1])

		shareSnapshot, _, err := vpcClient.GetShareSnapshot(getShareSnapshotOptions)
		if err != nil {
			return err
		}

		obj = *shareSnapshot
		return nil
	}
}

func testAccCheckIBMIsShareSnapshotDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_share_snapshot" {
			continue
		}

		getShareSnapshotOptions := &vpcv1.GetShareSnapshotOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getShareSnapshotOptions.SetShareID(parts[0])
		getShareSnapshotOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetShareSnapshot(getShareSnapshotOptions)

		if err == nil {
			return fmt.Errorf("ShareSnapshot still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ShareSnapshot (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsShareSnapshotBackupPolicyPlanReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		regionReferenceModel := make(map[string]interface{})
		regionReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		regionReferenceModel["name"] = "us-south"

		backupPolicyPlanRemoteModel := make(map[string]interface{})
		backupPolicyPlanRemoteModel["region"] = []map[string]interface{}{regionReferenceModel}

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/backup_policies/r134-076191ba-49c2-4763-94fd-c70de73ee2e6/plans/r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178"
		model["id"] = "r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178"
		model["name"] = "my-policy-plan"
		model["remote"] = []map[string]interface{}{backupPolicyPlanRemoteModel}
		model["resource_type"] = "backup_policy_plan"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	regionReferenceModel := new(vpcv1.RegionReference)
	regionReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	regionReferenceModel.Name = core.StringPtr("us-south")

	backupPolicyPlanRemoteModel := new(vpcv1.BackupPolicyPlanRemote)
	backupPolicyPlanRemoteModel.Region = regionReferenceModel

	model := new(vpcv1.BackupPolicyPlanReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/backup_policies/r134-076191ba-49c2-4763-94fd-c70de73ee2e6/plans/r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178")
	model.ID = core.StringPtr("r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178")
	model.Name = core.StringPtr("my-policy-plan")
	model.Remote = backupPolicyPlanRemoteModel
	model.ResourceType = core.StringPtr("backup_policy_plan")

	result, err := vpc.ResourceIBMIsShareSnapshotBackupPolicyPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsShareSnapshotDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsShareSnapshotDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsShareSnapshotBackupPolicyPlanRemoteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		regionReferenceModel := make(map[string]interface{})
		regionReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		regionReferenceModel["name"] = "us-south"

		model := make(map[string]interface{})
		model["region"] = []map[string]interface{}{regionReferenceModel}

		assert.Equal(t, result, model)
	}

	regionReferenceModel := new(vpcv1.RegionReference)
	regionReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	regionReferenceModel.Name = core.StringPtr("us-south")

	model := new(vpcv1.BackupPolicyPlanRemote)
	model.Region = regionReferenceModel

	result, err := vpc.ResourceIBMIsShareSnapshotBackupPolicyPlanRemoteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsShareSnapshotRegionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		model["name"] = "us-south"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.RegionReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	model.Name = core.StringPtr("us-south")

	result, err := vpc.ResourceIBMIsShareSnapshotRegionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsShareSnapshotResourceGroupReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
		model["id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		model["name"] = "my-resource-group"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ResourceGroupReference)
	model.Href = core.StringPtr("https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345")
	model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.Name = core.StringPtr("my-resource-group")

	result, err := vpc.ResourceIBMIsShareSnapshotResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsShareSnapshotShareSnapshotStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "encryption_key_deleted"
		model["message"] = "testString"
		model["more_info"] = "https://cloud.ibm.com/docs/key-protect?topic=key-protect-restore-keys"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ShareSnapshotStatusReason)
	model.Code = core.StringPtr("encryption_key_deleted")
	model.Message = core.StringPtr("testString")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/key-protect?topic=key-protect-restore-keys")

	result, err := vpc.ResourceIBMIsShareSnapshotShareSnapshotStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsShareSnapshotZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.ResourceIBMIsShareSnapshotZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
