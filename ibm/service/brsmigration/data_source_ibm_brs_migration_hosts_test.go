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
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationHostsDataSourceBasic(t *testing.T) {
	hostMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	hostType := "virtual_server"
	hostEnv := "classic"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationHostsDataSourceConfigBasic(hostMigrationID, hostType, hostEnv),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_hosts.brs_migration_hosts_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_hosts.brs_migration_hosts_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_hosts.brs_migration_hosts_instance", "hosts.#"),
					resource.TestCheckResourceAttr("data.ibm_brs_migration_hosts.brs_migration_hosts_instance", "hosts.0.type", hostType),
					resource.TestCheckResourceAttr("data.ibm_brs_migration_hosts.brs_migration_hosts_instance", "hosts.0.env", hostEnv),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationHostsDataSourceConfigBasic(hostMigrationID string, hostType string, hostEnv string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_host" "brs_migration_host_instance" {
			migration_id = "%s"
			type = "%s"
			env = "%s"
		}

		data "ibm_brs_migration_hosts" "brs_migration_hosts_instance" {
			migration_id = ibm_brs_migration_host.brs_migration_host_instance.migration_id
			env = ibm_brs_migration_host.brs_migration_host_instance.env
			type = ibm_brs_migration_host.brs_migration_host_instance.type
			migrated = ibm_brs_migration_host.brs_migration_host_instance.migrated
		}
	`, hostMigrationID, hostType, hostEnv)
}


func TestDataSourceIbmBrsMigrationHostsHostToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		hostComputeModel := make(map[string]interface{})
		hostComputeModel["status"] = "pending"
		hostComputeModel["os_family"] = "linux"
		hostComputeModel["global_identifier"] = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		hostComputeModel["throughput_mbps"] = int(0)
		hostComputeModel["public_ips"] = []string{"testString"}
		hostComputeModel["name"] = "prod-migration-server-01"
		hostComputeModel["os_type"] = "UBUNTU_22_64"
		hostComputeModel["ip_address"] = "10.240.0.5"
		hostComputeModel["profile"] = "bx2-4x16"
		hostComputeModel["vcpu_count"] = int(0)
		hostComputeModel["memory_gib"] = int(0)
		hostComputeModel["image_id"] = "r134-f47ac10b-58cc-4372-a567-0e02b2c3d479"
		hostComputeModel["datacenter"] = "dal10"

		volumeAttachmentModel := make(map[string]interface{})
		volumeAttachmentModel["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"

		model := make(map[string]interface{})
		model["id"] = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["type"] = "virtual_server"
		model["env"] = "vpc"
		model["compute"] = []map[string]interface{}{hostComputeModel}
		model["volume_attachments"] = []map[string]interface{}{volumeAttachmentModel}
		model["migrated"] = false
		model["workload_id"] = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["registered_at"] = "2024-06-01T09:00:00.000Z"

		assert.Equal(t, result, model)
	}

	hostComputeModel := new(brsmigrationv1.HostComputeClassicComputeDetails)
	hostComputeModel.Status = core.StringPtr("pending")
	hostComputeModel.OsFamily = core.StringPtr("linux")
	hostComputeModel.GlobalIdentifier = core.StringPtr("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	hostComputeModel.ThroughputMbps = core.Int64Ptr(int64(0))
	hostComputeModel.PublicIps = []string{"testString"}
	hostComputeModel.Name = core.StringPtr("prod-migration-server-01")
	hostComputeModel.OsType = core.StringPtr("UBUNTU_22_64")
	hostComputeModel.IpAddress = core.StringPtr("10.240.0.5")
	hostComputeModel.Profile = core.StringPtr("bx2-4x16")
	hostComputeModel.VcpuCount = core.Int64Ptr(int64(0))
	hostComputeModel.MemoryGib = core.Int64Ptr(int64(0))
	hostComputeModel.ImageID = core.StringPtr("r134-f47ac10b-58cc-4372-a567-0e02b2c3d479")
	hostComputeModel.Datacenter = core.StringPtr("dal10")

	volumeAttachmentModel := new(brsmigrationv1.VolumeAttachment)
	volumeAttachmentModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")

	model := new(brsmigrationv1.Host)
	model.ID = core.StringPtr("host-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.Type = core.StringPtr("virtual_server")
	model.Env = core.StringPtr("vpc")
	model.Compute = hostComputeModel
	model.VolumeAttachments = []brsmigrationv1.VolumeAttachment{*volumeAttachmentModel}
	model.Migrated = core.BoolPtr(false)
	model.WorkloadID = core.StringPtr("wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.RegisteredAt = CreateMockDateTime("2024-06-01T09:00:00.000Z")

	result, err := brsmigration.DataSourceIbmBrsMigrationHostsHostToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationHostsHostComputeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "pending"
		model["os_family"] = "linux"
		model["global_identifier"] = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["throughput_mbps"] = int(0)
		model["public_ips"] = []string{"testString"}
		model["name"] = "prod-migration-server-01"
		model["os_type"] = "UBUNTU_22_64"
		model["ip_address"] = "10.240.0.5"
		model["profile"] = "bx2-4x16"
		model["vcpu_count"] = int(0)
		model["memory_gib"] = int(0)
		model["image_id"] = "r134-f47ac10b-58cc-4372-a567-0e02b2c3d479"
		model["datacenter"] = "dal10"
		model["region"] = "us-east"
		model["zone"] = "us-east-1"
		model["lifecycle_state"] = "deleting"
		model["health_state"] = "ok"
		model["cpu_architecture"] = "amd64"
		model["subnet_id"] = "0717-abcdef01-2345-6789-abcd-ef0123456789"
		model["security_groups"] = []string{"testString"}
		model["resource_group_id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		model["crn"] = "crn:v1:bluemix:public:is:us-east-1:a/123456::instance:0717-e2b9a867-4a23-4b2e-9f3c-d1234567890a"
		model["vpc_id"] = "r134-12345678-abcd-ef01-2345-678901234567"
		model["vpc_name"] = "my-migration-vpc"
		model["boot_volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.HostCompute)
	model.Status = core.StringPtr("pending")
	model.OsFamily = core.StringPtr("linux")
	model.GlobalIdentifier = core.StringPtr("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.ThroughputMbps = core.Int64Ptr(int64(0))
	model.PublicIps = []string{"testString"}
	model.Name = core.StringPtr("prod-migration-server-01")
	model.OsType = core.StringPtr("UBUNTU_22_64")
	model.IpAddress = core.StringPtr("10.240.0.5")
	model.Profile = core.StringPtr("bx2-4x16")
	model.VcpuCount = core.Int64Ptr(int64(0))
	model.MemoryGib = core.Int64Ptr(int64(0))
	model.ImageID = core.StringPtr("r134-f47ac10b-58cc-4372-a567-0e02b2c3d479")
	model.Datacenter = core.StringPtr("dal10")
	model.Region = core.StringPtr("us-east")
	model.Zone = core.StringPtr("us-east-1")
	model.LifecycleState = core.StringPtr("deleting")
	model.HealthState = core.StringPtr("ok")
	model.CpuArchitecture = core.StringPtr("amd64")
	model.SubnetID = core.StringPtr("0717-abcdef01-2345-6789-abcd-ef0123456789")
	model.SecurityGroups = []string{"testString"}
	model.ResourceGroupID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.Crn = core.StringPtr("crn:v1:bluemix:public:is:us-east-1:a/123456::instance:0717-e2b9a867-4a23-4b2e-9f3c-d1234567890a")
	model.VpcID = core.StringPtr("r134-12345678-abcd-ef01-2345-678901234567")
	model.VpcName = core.StringPtr("my-migration-vpc")
	model.BootVolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")

	result, err := brsmigration.DataSourceIbmBrsMigrationHostsHostComputeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationHostsHostComputeClassicComputeDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "pending"
		model["os_family"] = "linux"
		model["global_identifier"] = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["throughput_mbps"] = int(0)
		model["public_ips"] = []string{"testString"}
		model["name"] = "prod-migration-server-01"
		model["os_type"] = "UBUNTU_22_64"
		model["ip_address"] = "10.240.0.5"
		model["profile"] = "bx2-4x16"
		model["vcpu_count"] = int(0)
		model["memory_gib"] = int(0)
		model["image_id"] = "r134-f47ac10b-58cc-4372-a567-0e02b2c3d479"
		model["datacenter"] = "dal10"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.HostComputeClassicComputeDetails)
	model.Status = core.StringPtr("pending")
	model.OsFamily = core.StringPtr("linux")
	model.GlobalIdentifier = core.StringPtr("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.ThroughputMbps = core.Int64Ptr(int64(0))
	model.PublicIps = []string{"testString"}
	model.Name = core.StringPtr("prod-migration-server-01")
	model.OsType = core.StringPtr("UBUNTU_22_64")
	model.IpAddress = core.StringPtr("10.240.0.5")
	model.Profile = core.StringPtr("bx2-4x16")
	model.VcpuCount = core.Int64Ptr(int64(0))
	model.MemoryGib = core.Int64Ptr(int64(0))
	model.ImageID = core.StringPtr("r134-f47ac10b-58cc-4372-a567-0e02b2c3d479")
	model.Datacenter = core.StringPtr("dal10")

	result, err := brsmigration.DataSourceIbmBrsMigrationHostsHostComputeClassicComputeDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationHostsHostComputeVPCComputeDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "pending"
		model["os_family"] = "linux"
		model["global_identifier"] = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["throughput_mbps"] = int(0)
		model["public_ips"] = []string{"testString"}
		model["name"] = "prod-migration-server-01"
		model["os_type"] = "UBUNTU_22_64"
		model["ip_address"] = "10.240.0.5"
		model["profile"] = "bx2-4x16"
		model["vcpu_count"] = int(0)
		model["memory_gib"] = int(0)
		model["image_id"] = "r134-f47ac10b-58cc-4372-a567-0e02b2c3d479"
		model["region"] = "us-east"
		model["zone"] = "us-east-1"
		model["lifecycle_state"] = "deleting"
		model["health_state"] = "ok"
		model["cpu_architecture"] = "amd64"
		model["subnet_id"] = "0717-abcdef01-2345-6789-abcd-ef0123456789"
		model["security_groups"] = []string{"testString"}
		model["resource_group_id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		model["crn"] = "crn:v1:bluemix:public:is:us-east-1:a/123456::instance:0717-e2b9a867-4a23-4b2e-9f3c-d1234567890a"
		model["vpc_id"] = "r134-12345678-abcd-ef01-2345-678901234567"
		model["vpc_name"] = "my-migration-vpc"
		model["boot_volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.HostComputeVPCComputeDetails)
	model.Status = core.StringPtr("pending")
	model.OsFamily = core.StringPtr("linux")
	model.GlobalIdentifier = core.StringPtr("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.ThroughputMbps = core.Int64Ptr(int64(0))
	model.PublicIps = []string{"testString"}
	model.Name = core.StringPtr("prod-migration-server-01")
	model.OsType = core.StringPtr("UBUNTU_22_64")
	model.IpAddress = core.StringPtr("10.240.0.5")
	model.Profile = core.StringPtr("bx2-4x16")
	model.VcpuCount = core.Int64Ptr(int64(0))
	model.MemoryGib = core.Int64Ptr(int64(0))
	model.ImageID = core.StringPtr("r134-f47ac10b-58cc-4372-a567-0e02b2c3d479")
	model.Region = core.StringPtr("us-east")
	model.Zone = core.StringPtr("us-east-1")
	model.LifecycleState = core.StringPtr("deleting")
	model.HealthState = core.StringPtr("ok")
	model.CpuArchitecture = core.StringPtr("amd64")
	model.SubnetID = core.StringPtr("0717-abcdef01-2345-6789-abcd-ef0123456789")
	model.SecurityGroups = []string{"testString"}
	model.ResourceGroupID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.Crn = core.StringPtr("crn:v1:bluemix:public:is:us-east-1:a/123456::instance:0717-e2b9a867-4a23-4b2e-9f3c-d1234567890a")
	model.VpcID = core.StringPtr("r134-12345678-abcd-ef01-2345-678901234567")
	model.VpcName = core.StringPtr("my-migration-vpc")
	model.BootVolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")

	result, err := brsmigration.DataSourceIbmBrsMigrationHostsHostComputeVPCComputeDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationHostsVolumeAttachmentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.VolumeAttachment)
	model.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")

	result, err := brsmigration.DataSourceIbmBrsMigrationHostsVolumeAttachmentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
