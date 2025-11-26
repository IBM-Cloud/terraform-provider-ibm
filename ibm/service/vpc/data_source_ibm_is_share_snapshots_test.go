// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.91.0-d9755c53-20240605-153412
 */

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsShareSnapshotsDataSourceBasic(t *testing.T) {
	shareSnapshotShareID := fmt.Sprintf("tf_share_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsShareSnapshotsDataSourceConfigBasic(shareSnapshotShareID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "share_id"),
				),
			},
		},
	})
}

func TestAccIBMIsShareSnapshotsDataSourceAllArgs(t *testing.T) {
	shareSnapshotShareID := fmt.Sprintf("tf_share_id_%d", acctest.RandIntRange(10, 100))
	shareSnapshotName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsShareSnapshotsDataSourceConfig(shareSnapshotShareID, shareSnapshotName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "share_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "backup_policy_plan_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.captured_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.fingerprint"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.minimum_size"),
					resource.TestCheckResourceAttr("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.name", shareSnapshotName),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_snapshots.is_share_snapshots_instance", "snapshots.0.status"),
				),
			},
		},
	})
}

func testAccCheckIBMIsShareSnapshotsDataSourceConfigBasic(shareSnapshotShareID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share_id = "%s"
		}

		data "ibm_is_share_snapshots" "is_share_snapshots_instance" {
			share_id = ibm_is_share_snapshot.is_share_snapshot_instance.share_id
			backup_policy_plan_id = "backup_policy_plan_id"
			name = ibm_is_share_snapshot.is_share_snapshot_instance.name
			sort = "name"
		}
	`, shareSnapshotShareID)
}

func testAccCheckIBMIsShareSnapshotsDataSourceConfig(shareSnapshotShareID string, shareSnapshotName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share_snapshot" "is_share_snapshot_instance" {
			share_id = "%s"
			name = "%s"
			user_tags = "FIXME"
		}

		data "ibm_is_share_snapshots" "is_share_snapshots_instance" {
			share_id = ibm_is_share_snapshot.is_share_snapshot_instance.share_id
			backup_policy_plan_id = "backup_policy_plan_id"
			name = ibm_is_share_snapshot.is_share_snapshot_instance.name
			sort = "name"
		}
	`, shareSnapshotShareID, shareSnapshotName)
}

func TestDataSourceIBMIsShareSnapshotsShareSnapshotToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		regionReferenceModel := make(map[string]interface{})
		regionReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		regionReferenceModel["name"] = "us-south"

		backupPolicyPlanRemoteModel := make(map[string]interface{})
		backupPolicyPlanRemoteModel["region"] = []map[string]interface{}{regionReferenceModel}

		backupPolicyPlanReferenceModel := make(map[string]interface{})
		backupPolicyPlanReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		backupPolicyPlanReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/backup_policies/r134-076191ba-49c2-4763-94fd-c70de73ee2e6/plans/r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178"
		backupPolicyPlanReferenceModel["id"] = "r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178"
		backupPolicyPlanReferenceModel["name"] = "my-policy-plan"
		backupPolicyPlanReferenceModel["remote"] = []map[string]interface{}{backupPolicyPlanRemoteModel}
		backupPolicyPlanReferenceModel["resource_type"] = "backup_policy_plan"

		resourceGroupReferenceModel := make(map[string]interface{})
		resourceGroupReferenceModel["href"] = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
		resourceGroupReferenceModel["id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		resourceGroupReferenceModel["name"] = "Default"

		shareSnapshotStatusReasonModel := make(map[string]interface{})
		shareSnapshotStatusReasonModel["code"] = "encryption_key_deleted"
		shareSnapshotStatusReasonModel["message"] = "testString"
		shareSnapshotStatusReasonModel["more_info"] = "https://cloud.ibm.com/docs/key-protect?topic=key-protect-restore-keys"

		zoneReferenceModel := make(map[string]interface{})
		zoneReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		zoneReferenceModel["name"] = "us-south-1"

		model := make(map[string]interface{})
		model["backup_policy_plan"] = []map[string]interface{}{backupPolicyPlanReferenceModel}
		model["captured_at"] = "2019-01-01T12:00:00.000Z"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::share-snapshot:r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
		model["fingerprint"] = "7abc3aef-c2bc-4f65-a296-2928e534d498"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/shares/r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/snapshots/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
		model["id"] = "r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
		model["lifecycle_state"] = "stable"
		model["minimum_size"] = int(10)
		model["name"] = "my-share-snapshot"
		model["resource_group"] = []map[string]interface{}{resourceGroupReferenceModel}
		model["resource_type"] = "share_snapshot"
		model["status"] = "available"
		model["status_reasons"] = []map[string]interface{}{shareSnapshotStatusReasonModel}
		model["user_tags"] = []string{"testString"}
		model["zone"] = []map[string]interface{}{zoneReferenceModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	regionReferenceModel := new(vpcv1.RegionReference)
	regionReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	regionReferenceModel.Name = core.StringPtr("us-south")

	backupPolicyPlanRemoteModel := new(vpcv1.BackupPolicyPlanRemote)
	backupPolicyPlanRemoteModel.Region = regionReferenceModel

	backupPolicyPlanReferenceModel := new(vpcv1.BackupPolicyPlanReference)
	backupPolicyPlanReferenceModel.Deleted = deletedModel
	backupPolicyPlanReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/backup_policies/r134-076191ba-49c2-4763-94fd-c70de73ee2e6/plans/r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178")
	backupPolicyPlanReferenceModel.ID = core.StringPtr("r134-6da51cfe-6f7b-4638-a6ba-00e9c327b178")
	backupPolicyPlanReferenceModel.Name = core.StringPtr("my-policy-plan")
	backupPolicyPlanReferenceModel.Remote = backupPolicyPlanRemoteModel
	backupPolicyPlanReferenceModel.ResourceType = core.StringPtr("backup_policy_plan")

	resourceGroupReferenceModel := new(vpcv1.ResourceGroupReference)
	resourceGroupReferenceModel.Href = core.StringPtr("https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345")
	resourceGroupReferenceModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	resourceGroupReferenceModel.Name = core.StringPtr("Default")

	shareSnapshotStatusReasonModel := new(vpcv1.ShareSnapshotStatusReason)
	shareSnapshotStatusReasonModel.Code = core.StringPtr("encryption_key_deleted")
	shareSnapshotStatusReasonModel.Message = core.StringPtr("testString")
	shareSnapshotStatusReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/key-protect?topic=key-protect-restore-keys")

	zoneReferenceModel := new(vpcv1.ZoneReference)
	zoneReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	zoneReferenceModel.Name = core.StringPtr("us-south-1")

	model := new(vpcv1.ShareSnapshot)
	model.BackupPolicyPlan = backupPolicyPlanReferenceModel
	model.CapturedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::share-snapshot:r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4")
	model.Fingerprint = core.StringPtr("7abc3aef-c2bc-4f65-a296-2928e534d498")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/shares/r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/snapshots/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4")
	model.ID = core.StringPtr("r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4")
	model.LifecycleState = core.StringPtr("stable")
	model.MinimumSize = core.Int64Ptr(int64(10))
	model.Name = core.StringPtr("my-share-snapshot")
	model.ResourceGroup = resourceGroupReferenceModel
	model.ResourceType = core.StringPtr("share_snapshot")
	model.Status = core.StringPtr("available")
	model.StatusReasons = []vpcv1.ShareSnapshotStatusReason{*shareSnapshotStatusReasonModel}
	model.UserTags = []string{"testString"}
	model.Zone = zoneReferenceModel

	result, err := vpc.DataSourceIBMIsShareSnapshotsShareSnapshotToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotsBackupPolicyPlanReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotsBackupPolicyPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotsDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsShareSnapshotsDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotsBackupPolicyPlanRemoteToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotsBackupPolicyPlanRemoteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotsRegionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		model["name"] = "us-south"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.RegionReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	model.Name = core.StringPtr("us-south")

	result, err := vpc.DataSourceIBMIsShareSnapshotsRegionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotsResourceGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotsResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotsShareSnapshotStatusReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsShareSnapshotsShareSnapshotStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareSnapshotsZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.DataSourceIBMIsShareSnapshotsZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
