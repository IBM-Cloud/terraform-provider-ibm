// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISSnapshotsDatasource_basic(t *testing.T) {
	var snapshot string
	snpName := "data.ibm_is_snapshots.ds_snapshot"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`yourkey`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSnapshotExists("ibm_is_snapshot.testacc_snapshot", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", name1),
				),
			},
			{
				Config: testDSCheckIBMISSnapshotsConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						snpName, "snapshots.0.name", name1),
					// resource.TestCheckResourceAttrSet(snpName, "snapshots.0.delatable"), // Commented as it is deprecated.
					resource.TestCheckResourceAttrSet(snpName, "snapshots.0.href"),
					resource.TestCheckResourceAttrSet(snpName, "snapshots.0.crn"),
					resource.TestCheckResourceAttrSet(snpName, "snapshots.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet(snpName, "snapshots.0.encryption"),
					resource.TestCheckResourceAttrSet(snpName, "snapshots.0.captured_at"),
				),
			},
		},
	})
}

func testDSCheckIBMISSnapshotsConfig(name1 string) string {
	return fmt.Sprintf(`
		data "ibm_is_snapshots" "ds_snapshot" {
			name = "%s"
		}`, name1)
}
