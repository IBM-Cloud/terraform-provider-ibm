// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsSnapshotConsistencyGroupDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsSnapshotConsistencyGroupDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, scgname, snapname, deleteSnapshotsOnDelete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "delete_snapshots_on_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshots.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshots.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshots.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshots.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshots.0.href"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotConsistencyGroupDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, snapname, scgname, deleteSnapshotsOnDelete string) string {
	return testAccCheckIBMIsSnapshotConsistencyGroupConfig(vpcname, subnetname, sshname, publicKey, name, scgname, snapname, deleteSnapshotsOnDelete) + fmt.Sprintf(`
	data "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group" {
		identifier = ibm_is_snapshot_consistency_group.is_snapshot_consistency_group.id
	  }
`)
}
