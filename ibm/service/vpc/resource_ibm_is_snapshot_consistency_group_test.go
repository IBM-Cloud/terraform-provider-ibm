// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsSnapshotConsistencyGroupBasic(t *testing.T) {
	var conf vpcv1.SnapshotConsistencyGroup
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	deleteSnapshotsOnDelete := "true"
	scgname := fmt.Sprintf("tf-snap-cons-grp-name-%d", acctest.RandIntRange(10, 100))
	deleteSnapshotsOnDeleteUpdate := "false"
	snapname := fmt.Sprintf("tf-snap-name-%d", acctest.RandIntRange(10, 100))
	scgnameUpdate := fmt.Sprintf("tf-snap-cons-grp-name-update-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsSnapshotConsistencyGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotConsistencyGroupConfig(vpcname, subnetname, sshname, publicKey, name, scgname, snapname, deleteSnapshotsOnDelete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsSnapshotConsistencyGroupExists("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", conf),
					resource.TestCheckResourceAttr("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "delete_snapshots_on_delete", deleteSnapshotsOnDelete),
					resource.TestCheckResourceAttr("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "name", scgname),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshot_reference.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshot_reference.0.crn"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshot_reference.0.name"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotConsistencyGroupConfig(vpcname, subnetname, sshname, publicKey, name, scgnameUpdate, snapname, deleteSnapshotsOnDeleteUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "delete_snapshots_on_delete", deleteSnapshotsOnDeleteUpdate),
					resource.TestCheckResourceAttr("ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "name", scgnameUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotConsistencyGroupConfig(vpcname, subnetname, sshname, publicKey, name, snapname, scgname, deleteSnapshotsOnDelete string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name           				= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  } 
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }

	resource "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group" {
		delete_snapshots_on_delete = "%s"
		snapshots {
		  name = "%s"
		  source_volume = ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
		}
		name = "%s"
	  }
	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, deleteSnapshotsOnDelete, scgname, snapname)
}

func testAccCheckIBMIsSnapshotConsistencyGroupExists(n string, obj vpcv1.SnapshotConsistencyGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getSnapshotConsistencyGroupOptions := &vpcv1.GetSnapshotConsistencyGroupOptions{}

		getSnapshotConsistencyGroupOptions.SetID(rs.Primary.ID)

		snapshotConsistencyGroup, _, err := vpcClient.GetSnapshotConsistencyGroup(getSnapshotConsistencyGroupOptions)
		if err != nil {
			return err
		}

		obj = *snapshotConsistencyGroup
		return nil
	}
}

func testAccCheckIBMIsSnapshotConsistencyGroupDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_snapshot_consistency_group" {
			continue
		}

		getSnapshotConsistencyGroupOptions := &vpcv1.GetSnapshotConsistencyGroupOptions{}

		getSnapshotConsistencyGroupOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetSnapshotConsistencyGroup(getSnapshotConsistencyGroupOptions)

		if err == nil {
			return fmt.Errorf("SnapshotConsistencyGroup still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for SnapshotConsistencyGroup (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
