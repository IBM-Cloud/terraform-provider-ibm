// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func TestAccIBMIsSnapshotConsistencyGroupsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	deleteSnapshotsOnDelete := "true"
	scgname := fmt.Sprintf("tf-snap-cons-grp-name-%d", acctest.RandIntRange(10, 100))
	snapname := fmt.Sprintf("tf-snap-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotConsistencyGroupsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, scgname, snapname, deleteSnapshotsOnDelete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.backup_policy_job.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.delete_snapshots_on_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.snapshots.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.snapshots.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.snapshots.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.snapshots.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.0.snapshots.0.href"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotConsistencyGroupsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, snapname, scgname, deleteSnapshotsOnDelete string) string {
	return testAccCheckIBMIsSnapshotConsistencyGroupConfig(vpcname, subnetname, sshname, publicKey, name, scgname, snapname, deleteSnapshotsOnDelete) + fmt.Sprintf(`
	data "ibm_is_snapshot_consistency_groups" "is_snapshot_consistency_groups" {
	}
`)
}

// TestAccIBMIsSnapshotConsistencyGroupsBackupPolicyJob tests the complete backup policy job flow
// This test creates infrastructure, waits for backup job execution, and validates SCG creation
func TestAccIBMIsSnapshotConsistencyGroupsBackupPolicyJob(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-bp-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-bp-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-bp-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-bp-%d", acctest.RandIntRange(10, 100))
	volume1Name := fmt.Sprintf("tf-vol1-bp-%d", acctest.RandIntRange(10, 100))
	volume2Name := fmt.Sprintf("tf-vol2-bp-%d", acctest.RandIntRange(10, 100))
	backupPolicyName := fmt.Sprintf("tf-bp-%d", acctest.RandIntRange(10, 100))
	backupPlanName := fmt.Sprintf("tf-bp-plan-%d", acctest.RandIntRange(10, 100))
	tagName := fmt.Sprintf("tf-bp-tag-%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	// Calculate cron spec for 3 minutes from now in UTC
	cronSpec := calculateCronSpec(3)

	var scgID, scgName, jobID string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyJobSCGDestroy(&jobID),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsSnapshotConsistencyGroupsBackupPolicyJobConfig(
					vpcname, subnetname, sshname, publicKey, instanceName,
					volume1Name, volume2Name, backupPolicyName, backupPlanName, tagName, cronSpec,
				),
				Check: resource.ComposeTestCheckFunc(
					// Verify backup policy and plan are created
					resource.TestCheckResourceAttr("ibm_is_backup_policy.test_bp", "name", backupPolicyName),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.test_bp_plan", "name", backupPlanName),

					// Custom check function to wait for backup job and validate SCG
					testAccCheckBackupPolicyJobAndSCG("ibm_is_backup_policy.test_bp", &scgID, &scgName, &jobID),
				),
			},
			{
				Config: testAccCheckIBMIsSnapshotConsistencyGroupsBackupPolicyJobConfigWithDataSources(
					vpcname, subnetname, sshname, publicKey, instanceName,
					volume1Name, volume2Name, backupPolicyName, backupPlanName, tagName, cronSpec,
				),
				Check: resource.ComposeTestCheckFunc(
					// Verify data source lists SCGs
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.test_scgs_all", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.test_scgs_all", "snapshot_consistency_groups.#"),

					// Verify data source with backup_policy_job filter
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.test_scgs_filtered", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.test_scgs_filtered", "snapshot_consistency_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.test_scgs_filtered", "snapshot_consistency_groups.0.backup_policy_job.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.test_scgs_filtered", "snapshot_consistency_groups.0.backup_policy_job.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.test_scgs_filtered", "snapshot_consistency_groups.0.backup_policy_job.0.href"),

					// Verify individual SCG data source by ID
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_id", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_id", "backup_policy_job.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_id", "backup_policy_job.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_id", "backup_policy_job.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_id", "backup_policy_job.0.resource_type"),

					// Verify individual SCG data source by name
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_name", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_name", "backup_policy_job.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_name", "backup_policy_job.0.id"),

					// Verify snapshots exist
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.test_scg_by_id", "snapshots.#"),
				),
			},
		},
	})
}

// calculateCronSpec generates a cron spec for X minutes from now in UTC
func calculateCronSpec(minutesFromNow int) string {
	targetTime := time.Now().UTC().Add(time.Duration(minutesFromNow) * time.Minute)
	// Format: "minute hour * * *" (every day at the specified time)
	return fmt.Sprintf("%d %d * * *", targetTime.Minute(), targetTime.Hour())
}

// testAccCheckIBMIsBackupPolicyJobSCGDestroy cleans up all SCGs and snapshots created by backup policy jobs
func testAccCheckIBMIsBackupPolicyJobSCGDestroy(jobID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if jobID == nil || *jobID == "" {
			fmt.Println("No job ID to clean up SCGs")
			return nil
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return fmt.Errorf("failed to get VPC client: %s", err)
		}

		// List all SCGs created by this backup policy job
		listSCGOptions := &vpcv1.ListSnapshotConsistencyGroupsOptions{
			BackupPolicyJobID: jobID,
		}

		scgCollection, _, err := vpcClient.ListSnapshotConsistencyGroups(listSCGOptions)
		if err != nil {
			return fmt.Errorf("error listing SCGs for job %s: %s", *jobID, err)
		}

		if len(scgCollection.SnapshotConsistencyGroups) == 0 {
			fmt.Printf("No SCGs found for job %s\n", *jobID)
			return nil
		}

		// Delete all SCGs created by this job
		deletedCount := 0
		for _, scg := range scgCollection.SnapshotConsistencyGroups {
			deleteSCGOptions := &vpcv1.DeleteSnapshotConsistencyGroupOptions{
				ID: scg.ID,
			}

			_, response, err := vpcClient.DeleteSnapshotConsistencyGroup(deleteSCGOptions)
			if err != nil {
				if response != nil && response.StatusCode == 404 {
					fmt.Printf("SCG %s already deleted\n", *scg.ID)
					continue
				}
				return fmt.Errorf("error deleting SCG %s: %s", *scg.ID, err)
			}

			fmt.Printf("Successfully deleted SCG %s (name: %s) with %d snapshots\n",
				*scg.ID, *scg.Name, len(scg.Snapshots))
			deletedCount++
		}

		fmt.Printf("Cleanup complete: deleted %d SCG(s) for job %s\n", deletedCount, *jobID)
		return nil
	}
}

// testAccCheckBackupPolicyJobAndSCG waits for backup job execution and validates SCG creation
func testAccCheckBackupPolicyJobAndSCG(backupPolicyResource string, scgID, scgName, jobID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Get backup policy ID from state
		rs, ok := s.RootModule().Resources[backupPolicyResource]
		if !ok {
			return fmt.Errorf("backup policy resource not found: %s", backupPolicyResource)
		}
		backupPolicyID := rs.Primary.ID

		// Get VPC client
		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return fmt.Errorf("failed to get VPC client: %s", err)
		}

		// Poll for backup job with timeout
		timeout := time.Now().Add(90 * time.Minute)
		pollInterval := 30 * time.Second

		fmt.Printf("Waiting for backup policy job to execute (timeout: 90 minutes)...\n")

		for time.Now().Before(timeout) {
			// List backup policy jobs
			listJobsOptions := &vpcv1.ListBackupPolicyJobsOptions{
				BackupPolicyID: &backupPolicyID,
			}

			jobsCollection, _, err := vpcClient.ListBackupPolicyJobs(listJobsOptions)
			if err != nil {
				fmt.Printf("Error listing backup policy jobs: %s\n", err)
				time.Sleep(pollInterval)
				continue
			}

			if len(jobsCollection.Jobs) == 0 {
				fmt.Printf("No backup jobs found yet, waiting %v...\n", pollInterval)
				time.Sleep(pollInterval)
				continue
			}

			// Get the first job
			foundJobID := *jobsCollection.Jobs[0].ID
			fmt.Printf("Found backup policy job: %s\n", foundJobID)

			// List snapshot consistency groups filtered by this job
			listSCGOptions := &vpcv1.ListSnapshotConsistencyGroupsOptions{
				BackupPolicyJobID: &foundJobID,
			}

			scgCollection, _, err := vpcClient.ListSnapshotConsistencyGroups(listSCGOptions)
			if err != nil {
				fmt.Printf("Error listing snapshot consistency groups: %s\n", err)
				time.Sleep(pollInterval)
				continue
			}

			if len(scgCollection.SnapshotConsistencyGroups) == 0 {
				fmt.Printf("No snapshot consistency groups found yet, waiting %v...\n", pollInterval)
				time.Sleep(pollInterval)
				continue
			}

			// Validate SCG
			scg := scgCollection.SnapshotConsistencyGroups[0]
			fmt.Printf("Found snapshot consistency group: %s\n", *scg.ID)

			// Verify backup_policy_job field is populated
			if scg.BackupPolicyJob == nil {
				return fmt.Errorf("snapshot consistency group %s has nil backup_policy_job", *scg.ID)
			}

			if scg.BackupPolicyJob.ID == nil || *scg.BackupPolicyJob.ID != foundJobID {
				return fmt.Errorf("snapshot consistency group %s has incorrect backup_policy_job.id: expected %s, got %v",
					*scg.ID, foundJobID, scg.BackupPolicyJob.ID)
			}

			// Verify snapshots exist
			if len(scg.Snapshots) == 0 {
				return fmt.Errorf("snapshot consistency group %s has no snapshots", *scg.ID)
			}

			// Store IDs for use in next test step
			*scgID = *scg.ID
			*scgName = *scg.Name
			*jobID = foundJobID

			fmt.Printf("Successfully validated snapshot consistency group with %d snapshots\n", len(scg.Snapshots))
			return nil
		}

		return fmt.Errorf("timeout waiting for backup policy job and snapshot consistency group creation")
	}
}

func testAccCheckIBMIsSnapshotConsistencyGroupsBackupPolicyJobConfig(
	vpcname, subnetname, sshname, publicKey, instanceName,
	volume1Name, volume2Name, backupPolicyName, backupPlanName, tagName, cronSpec string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "test_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "test_subnet" {
	name            = "%s"
	vpc             = ibm_is_vpc.test_vpc.id
	zone            = "%s"
	ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "test_ssh_key" {
	name       = "%s"
	public_key = "%s"
}

resource "ibm_is_volume" "test_volume_1" {
	name     = "%s"
	profile  = "general-purpose"
	zone     = "%s"
	capacity = 100
}

resource "ibm_is_volume" "test_volume_2" {
	name     = "%s"
	profile  = "sdp"
	zone     = "%s"
	capacity = 100
}

resource "ibm_is_instance" "test_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
	
	primary_network_interface {
		subnet = ibm_is_subnet.test_subnet.id
	}
	
	vpc  = ibm_is_vpc.test_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.test_ssh_key.id]
	
	volumes = [
		ibm_is_volume.test_volume_1.id,
		ibm_is_volume.test_volume_2.id,
	]
}

resource "ibm_resource_tag" "test_instance_tag" {
	resource_id = ibm_is_instance.test_instance.crn
	tags        = ["%s"]
}

resource "ibm_is_backup_policy" "test_bp" {
	match_user_tags     = ["%s"]
	name                = "%s"
	match_resource_type = "instance"
}

resource "ibm_is_backup_policy_plan" "test_bp_plan" {
	backup_policy_id = ibm_is_backup_policy.test_bp.id
	name             = "%s"
	cron_spec        = "%s"
	depends_on       = [ibm_resource_tag.test_instance_tag]
}

data "ibm_is_snapshot_consistency_groups" "test_scgs" {
	depends_on = [ibm_is_backup_policy_plan.test_bp_plan]
}
`,
		vpcname, subnetname, acc.ISZoneName, sshname, publicKey,
		volume1Name, acc.ISZoneName,
		volume2Name, acc.ISZoneName,
		instanceName, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName,
		tagName, tagName, backupPolicyName,
		backupPlanName, cronSpec,
	)
}

func testAccCheckIBMIsSnapshotConsistencyGroupsBackupPolicyJobConfigWithDataSources(
	vpcname, subnetname, sshname, publicKey, instanceName,
	volume1Name, volume2Name, backupPolicyName, backupPlanName, tagName, cronSpec string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "test_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "test_subnet" {
	name            = "%s"
	vpc             = ibm_is_vpc.test_vpc.id
	zone            = "%s"
	ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "test_ssh_key" {
	name       = "%s"
	public_key = "%s"
}

resource "ibm_is_volume" "test_volume_1" {
	name     = "%s"
	profile  = "general-purpose"
	zone     = "%s"
	capacity = 100
}

resource "ibm_is_volume" "test_volume_2" {
	name     = "%s"
	profile  = "sdp"
	zone     = "%s"
	capacity = 100
}

resource "ibm_is_instance" "test_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
	
	primary_network_interface {
		subnet = ibm_is_subnet.test_subnet.id
	}
	
	vpc  = ibm_is_vpc.test_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.test_ssh_key.id]
	
	volumes = [
		ibm_is_volume.test_volume_1.id,
		ibm_is_volume.test_volume_2.id,
	]
}

resource "ibm_resource_tag" "test_instance_tag" {
	resource_id = ibm_is_instance.test_instance.crn
	tags        = ["%s"]
}

resource "ibm_is_backup_policy" "test_bp" {
	match_user_tags     = ["%s"]
	name                = "%s"
	match_resource_type = "instance"
}

resource "ibm_is_backup_policy_plan" "test_bp_plan" {
	backup_policy_id = ibm_is_backup_policy.test_bp.id
	name             = "%s"
	cron_spec        = "%s"
	depends_on       = [ibm_resource_tag.test_instance_tag]
}

data "ibm_is_backup_policy_jobs" "test_bp_jobs" {
	backup_policy_id = ibm_is_backup_policy.test_bp.id
	depends_on       = [ibm_is_backup_policy_plan.test_bp_plan]
}

# Data source: List all SCGs
data "ibm_is_snapshot_consistency_groups" "test_scgs_all" {
	depends_on = [data.ibm_is_backup_policy_jobs.test_bp_jobs]
}

# Data source: List SCGs filtered by backup_policy_job
data "ibm_is_snapshot_consistency_groups" "test_scgs_filtered" {
	backup_policy_job = data.ibm_is_backup_policy_jobs.test_bp_jobs.jobs[0].id
	depends_on        = [data.ibm_is_backup_policy_jobs.test_bp_jobs]
}

# Data source: Get individual SCG by ID
data "ibm_is_snapshot_consistency_group" "test_scg_by_id" {
	identifier = data.ibm_is_snapshot_consistency_groups.test_scgs_filtered.snapshot_consistency_groups[0].id
	depends_on = [data.ibm_is_snapshot_consistency_groups.test_scgs_filtered]
}

# Data source: Get individual SCG by name
data "ibm_is_snapshot_consistency_group" "test_scg_by_name" {
	name       = data.ibm_is_snapshot_consistency_groups.test_scgs_filtered.snapshot_consistency_groups[0].name
	depends_on = [data.ibm_is_snapshot_consistency_groups.test_scgs_filtered]
}
`,
		vpcname, subnetname, acc.ISZoneName, sshname, publicKey,
		volume1Name, acc.ISZoneName,
		volume2Name, acc.ISZoneName,
		instanceName, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName,
		tagName, tagName, backupPolicyName,
		backupPlanName, cronSpec,
	)
}
