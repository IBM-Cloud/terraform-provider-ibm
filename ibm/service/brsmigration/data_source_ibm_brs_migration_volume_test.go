// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationVolumeDataSourceBasic(t *testing.T) {
	volumeMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	volumeEnv := "classic"
	volumeStorageType := "block"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationVolumeDataSourceConfigBasic(volumeMigrationID, volumeEnv, volumeStorageType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "env"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "storage_type"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "attachment_state"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "storage.#"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "host_attachments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "migrated"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_volume.brs_migration_volume_instance", "registered_at"),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationVolumeDataSourceConfigBasic(volumeMigrationID string, volumeEnv string, volumeStorageType string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_volume" "brs_migration_volume_instance" {
			migration_id = "%s"
			env = "%s"
			storage_type = "%s"
		}

		data "ibm_brs_migration_volume" "brs_migration_volume_instance" {
			migration_id = ibm_brs_migration_volume.brs_migration_volume_instance.migration_id
			volume_id = ibm_brs_migration_volume.brs_migration_volume_instance.volume_id
		}
	`, volumeMigrationID, volumeEnv, volumeStorageType)
}


func TestDataSourceIbmBrsMigrationVolumeVolumeStorageToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sourcePathModel := make(map[string]interface{})
		sourcePathModel["path"] = "10.240.1.23:/nxg_s_vol_r006-aaa"
		sourcePathModel["vpc_id"] = "r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee"
		sourcePathModel["mount_target_id"] = "r006-11223344-5566-7788-99aa-bbccddeeff00"

		model := make(map[string]interface{})
		model["global_identifier"] = "r134-abcdef01-2345-6789-abcd-ef0123456789"
		model["name"] = "prod-data-vol-01"
		model["capacity_gib"] = int(0)
		model["iops"] = int(0)
		model["profile"] = "general-purpose"
		model["lifecycle_state"] = "stable"
		model["encryption"] = "provider_managed"
		model["throughput_mbps"] = int(0)
		model["source_paths"] = []map[string]interface{}{sourcePathModel}
		model["datacenter"] = "dal10"
		model["iscsi_target_ips"] = []string{"161.26.98.5", "161.26.98.6"}
		model["region"] = "us-east"
		model["zone"] = "us-east-1"
		model["crn"] = "crn:v1:bluemix:public:is:us-east-1:a/123456::volume:r134-abcdef01-2345-6789-abcd-ef0123456789"
		model["resource_group_id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		model["replication_role"] = "none"
		model["access_control_mode"] = "none"
		model["availability_mode"] = "none"

		assert.Equal(t, result, model)
	}

	sourcePathModel := new(brsmigrationv1.SourcePath)
	sourcePathModel.Path = core.StringPtr("10.240.1.23:/nxg_s_vol_r006-aaa")
	sourcePathModel.VpcID = core.StringPtr("r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee")
	sourcePathModel.MountTargetID = core.StringPtr("r006-11223344-5566-7788-99aa-bbccddeeff00")

	model := new(brsmigrationv1.VolumeStorage)
	model.GlobalIdentifier = core.StringPtr("r134-abcdef01-2345-6789-abcd-ef0123456789")
	model.Name = core.StringPtr("prod-data-vol-01")
	model.CapacityGib = core.Int64Ptr(int64(0))
	model.Iops = core.Int64Ptr(int64(0))
	model.Profile = core.StringPtr("general-purpose")
	model.LifecycleState = core.StringPtr("stable")
	model.Encryption = core.StringPtr("provider_managed")
	model.ThroughputMbps = core.Int64Ptr(int64(0))
	model.SourcePaths = []brsmigrationv1.SourcePath{*sourcePathModel}
	model.Datacenter = core.StringPtr("dal10")
	model.IscsiTargetIps = []string{"161.26.98.5", "161.26.98.6"}
	model.Region = core.StringPtr("us-east")
	model.Zone = core.StringPtr("us-east-1")
	model.Crn = core.StringPtr("crn:v1:bluemix:public:is:us-east-1:a/123456::volume:r134-abcdef01-2345-6789-abcd-ef0123456789")
	model.ResourceGroupID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.ReplicationRole = core.StringPtr("none")
	model.AccessControlMode = core.StringPtr("none")
	model.AvailabilityMode = core.StringPtr("none")

	result, err := brsmigration.DataSourceIbmBrsMigrationVolumeVolumeStorageToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationVolumeSourcePathToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["path"] = "10.240.1.23:/nxg_s_vol_r006-aaa"
		model["vpc_id"] = "r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee"
		model["mount_target_id"] = "r006-11223344-5566-7788-99aa-bbccddeeff00"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.SourcePath)
	model.Path = core.StringPtr("10.240.1.23:/nxg_s_vol_r006-aaa")
	model.VpcID = core.StringPtr("r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee")
	model.MountTargetID = core.StringPtr("r006-11223344-5566-7788-99aa-bbccddeeff00")

	result, err := brsmigration.DataSourceIbmBrsMigrationVolumeSourcePathToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationVolumeVolumeStorageClassicVolumeStorageDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sourcePathModel := make(map[string]interface{})
		sourcePathModel["path"] = "10.240.1.23:/nxg_s_vol_r006-aaa"
		sourcePathModel["vpc_id"] = "r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee"
		sourcePathModel["mount_target_id"] = "r006-11223344-5566-7788-99aa-bbccddeeff00"

		model := make(map[string]interface{})
		model["global_identifier"] = "r134-abcdef01-2345-6789-abcd-ef0123456789"
		model["name"] = "prod-data-vol-01"
		model["capacity_gib"] = int(0)
		model["iops"] = int(0)
		model["profile"] = "general-purpose"
		model["lifecycle_state"] = "stable"
		model["encryption"] = "provider_managed"
		model["throughput_mbps"] = int(0)
		model["source_paths"] = []map[string]interface{}{sourcePathModel}
		model["datacenter"] = "dal10"
		model["iscsi_target_ips"] = []string{"161.26.98.5", "161.26.98.6"}

		assert.Equal(t, result, model)
	}

	sourcePathModel := new(brsmigrationv1.SourcePath)
	sourcePathModel.Path = core.StringPtr("10.240.1.23:/nxg_s_vol_r006-aaa")
	sourcePathModel.VpcID = core.StringPtr("r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee")
	sourcePathModel.MountTargetID = core.StringPtr("r006-11223344-5566-7788-99aa-bbccddeeff00")

	model := new(brsmigrationv1.VolumeStorageClassicVolumeStorageDetails)
	model.GlobalIdentifier = core.StringPtr("r134-abcdef01-2345-6789-abcd-ef0123456789")
	model.Name = core.StringPtr("prod-data-vol-01")
	model.CapacityGib = core.Int64Ptr(int64(0))
	model.Iops = core.Int64Ptr(int64(0))
	model.Profile = core.StringPtr("general-purpose")
	model.LifecycleState = core.StringPtr("stable")
	model.Encryption = core.StringPtr("provider_managed")
	model.ThroughputMbps = core.Int64Ptr(int64(0))
	model.SourcePaths = []brsmigrationv1.SourcePath{*sourcePathModel}
	model.Datacenter = core.StringPtr("dal10")
	model.IscsiTargetIps = []string{"161.26.98.5", "161.26.98.6"}

	result, err := brsmigration.DataSourceIbmBrsMigrationVolumeVolumeStorageClassicVolumeStorageDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationVolumeVolumeStorageVPCVolumeStorageDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sourcePathModel := make(map[string]interface{})
		sourcePathModel["path"] = "10.240.1.23:/nxg_s_vol_r006-aaa"
		sourcePathModel["vpc_id"] = "r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee"
		sourcePathModel["mount_target_id"] = "r006-11223344-5566-7788-99aa-bbccddeeff00"

		model := make(map[string]interface{})
		model["global_identifier"] = "r134-abcdef01-2345-6789-abcd-ef0123456789"
		model["name"] = "prod-data-vol-01"
		model["capacity_gib"] = int(0)
		model["iops"] = int(0)
		model["profile"] = "general-purpose"
		model["lifecycle_state"] = "stable"
		model["encryption"] = "provider_managed"
		model["throughput_mbps"] = int(0)
		model["source_paths"] = []map[string]interface{}{sourcePathModel}
		model["region"] = "us-east"
		model["zone"] = "us-east-1"
		model["crn"] = "crn:v1:bluemix:public:is:us-east-1:a/123456::volume:r134-abcdef01-2345-6789-abcd-ef0123456789"
		model["resource_group_id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		model["replication_role"] = "none"
		model["access_control_mode"] = "none"
		model["availability_mode"] = "none"

		assert.Equal(t, result, model)
	}

	sourcePathModel := new(brsmigrationv1.SourcePath)
	sourcePathModel.Path = core.StringPtr("10.240.1.23:/nxg_s_vol_r006-aaa")
	sourcePathModel.VpcID = core.StringPtr("r006-aaa11111-bbbb-cccc-dddd-eeeeeeeeeeee")
	sourcePathModel.MountTargetID = core.StringPtr("r006-11223344-5566-7788-99aa-bbccddeeff00")

	model := new(brsmigrationv1.VolumeStorageVPCVolumeStorageDetails)
	model.GlobalIdentifier = core.StringPtr("r134-abcdef01-2345-6789-abcd-ef0123456789")
	model.Name = core.StringPtr("prod-data-vol-01")
	model.CapacityGib = core.Int64Ptr(int64(0))
	model.Iops = core.Int64Ptr(int64(0))
	model.Profile = core.StringPtr("general-purpose")
	model.LifecycleState = core.StringPtr("stable")
	model.Encryption = core.StringPtr("provider_managed")
	model.ThroughputMbps = core.Int64Ptr(int64(0))
	model.SourcePaths = []brsmigrationv1.SourcePath{*sourcePathModel}
	model.Region = core.StringPtr("us-east")
	model.Zone = core.StringPtr("us-east-1")
	model.Crn = core.StringPtr("crn:v1:bluemix:public:is:us-east-1:a/123456::volume:r134-abcdef01-2345-6789-abcd-ef0123456789")
	model.ResourceGroupID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.ReplicationRole = core.StringPtr("none")
	model.AccessControlMode = core.StringPtr("none")
	model.AvailabilityMode = core.StringPtr("none")

	result, err := brsmigration.DataSourceIbmBrsMigrationVolumeVolumeStorageVPCVolumeStorageDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationVolumeHostAttachmentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host_id"] = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["mount_path"] = "/mnt/data"
		model["type"] = "ext4"
		model["block_device"] = "/dev/vdb"
		model["device_id"] = "0717-80b3e36e-41f4-40e9-bd56-beae81792a68-679qb"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.HostAttachment)
	model.HostID = core.StringPtr("host-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.MountPath = core.StringPtr("/mnt/data")
	model.Type = core.StringPtr("ext4")
	model.BlockDevice = core.StringPtr("/dev/vdb")
	model.DeviceID = core.StringPtr("0717-80b3e36e-41f4-40e9-bd56-beae81792a68-679qb")

	result, err := brsmigration.DataSourceIbmBrsMigrationVolumeHostAttachmentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
