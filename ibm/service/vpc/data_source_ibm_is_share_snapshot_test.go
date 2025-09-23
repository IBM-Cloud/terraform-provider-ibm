// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.91.0-d9755c53-20240605-153412
 */

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/vpc"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsShareSnapshotDataSourceBasic(t *testing.T) {
	shareSnapshotShareID := fmt.Sprintf("tf_share_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsShareSnapshotDataSourceConfigBasic(shareSnapshotShareID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "share"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "share_snapshot"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "fingerprint"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "minimum_size"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "zone.#"),
				),
			},
		},
	})
}

func TestAccIBMIsShareSnapshotDataSourceAllArgs(t *testing.T) {
	shareSnapshotShareID := fmt.Sprintf("tf_share_id_%d", acctest.RandIntRange(10, 100))
	shareSnapshotName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsShareSnapshotDataSourceConfig(shareSnapshotShareID, shareSnapshotName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "share"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "share_snapshot"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "backup_policy_plan.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "captured_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "fingerprint"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "minimum_size"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "status_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "status_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "status_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshot.is_share_snapshot_instance", "zone.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsShareSnapshotDataSourceConfigBasic(shareSnapshotShareID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share = "%s"
		}

		data "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share = ibm_is_share_snapshot.is_share_snapshot_instance.share
			share_snapshot = ibm_is_share_snapshot.is_share_snapshot_instance.share_snapshot
		}
	`, shareSnapshotShareID)
}

func testAccCheckIBMIsShareSnapshotDataSourceConfig(shareSnapshotShareID string, shareSnapshotName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share = "%s"
			name = "%s"
			user_tags = "FIXME"
		}

		data "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share = ibm_is_share_snapshot.is_share_snapshot_instance.share
			share_snapshot = ibm_is_share_snapshot.is_share_snapshot_instance.share_snapshot
		}
	`, shareSnapshotShareID, shareSnapshotName)
}

func TestDataSourceIBMIsShareSnapshotBackupPolicyPlanReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotBackupPolicyPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsShareSnapshotDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotBackupPolicyPlanRemoteToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotBackupPolicyPlanRemoteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotRegionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		model["name"] = "us-south"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.RegionReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	model.Name = core.StringPtr("us-south")

	result, err := vpc.DataSourceIBMIsShareSnapshotRegionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotResourceGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotShareSnapshotStatusReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotShareSnapshotStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.DataSourceIBMIsShareSnapshotZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
