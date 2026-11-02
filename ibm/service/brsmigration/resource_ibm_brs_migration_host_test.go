// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationHostBasic(t *testing.T) {
	var conf brsmigrationv1.Host
	migrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	typeVar := "virtual_server"
	env := "classic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBrsMigrationHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationHostConfigBasic(migrationID, typeVar, env),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBrsMigrationHostExists("ibm_brs_migration_host.brs_migration_host_instance", conf),
					resource.TestCheckResourceAttr("ibm_brs_migration_host.brs_migration_host_instance", "migration_id", migrationID),
					resource.TestCheckResourceAttr("ibm_brs_migration_host.brs_migration_host_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_brs_migration_host.brs_migration_host_instance", "env", env),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_brs_migration_host.brs_migration_host_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBrsMigrationHostConfigBasic(migrationID string, typeVar string, env string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_host" "brs_migration_host_instance" {
			migration_id = "%s"
			type = "%s"
			env = "%s"
		}
	`, migrationID, typeVar, env)
}

func testAccCheckIbmBrsMigrationHostExists(n string, obj brsmigrationv1.Host) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV1()
		if err != nil {
			return err
		}

		getHostOptions := &brsmigrationv1.GetHostOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getHostOptions.SetMigrationID(parts[0])
		getHostOptions.SetHostID(parts[1])

		host, _, err := brsMigrationClient.GetHost(getHostOptions)
		if err != nil {
			return err
		}

		obj = *host
		return nil
	}
}

func testAccCheckIbmBrsMigrationHostDestroy(s *terraform.State) error {
	brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_brs_migration_host" {
			continue
		}

		getHostOptions := &brsmigrationv1.GetHostOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getHostOptions.SetMigrationID(parts[0])
		getHostOptions.SetHostID(parts[1])

		// Try to find the key
		_, response, err := brsMigrationClient.GetHost(getHostOptions)

		if err == nil {
			return fmt.Errorf("brs_migration_host still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for brs_migration_host (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBrsMigrationHostHostComputeToMap(t *testing.T) {
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

	result, err := brsmigration.ResourceIbmBrsMigrationHostHostComputeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationHostHostComputeClassicComputeDetailsToMap(t *testing.T) {
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

	result, err := brsmigration.ResourceIbmBrsMigrationHostHostComputeClassicComputeDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationHostHostComputeVPCComputeDetailsToMap(t *testing.T) {
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

	result, err := brsmigration.ResourceIbmBrsMigrationHostHostComputeVPCComputeDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationHostVolumeAttachmentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.VolumeAttachment)
	model.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")

	result, err := brsmigration.ResourceIbmBrsMigrationHostVolumeAttachmentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
