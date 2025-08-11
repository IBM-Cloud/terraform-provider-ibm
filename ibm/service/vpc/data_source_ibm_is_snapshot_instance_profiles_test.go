// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsSnapshotInstanceProfilesDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotInstanceProfilesDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_instance_profiles.is_snapshot_instance_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_instance_profiles.is_snapshot_instance_profiles", "instance_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_instance_profiles.is_snapshot_instance_profiles", "instance_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_instance_profiles.is_snapshot_instance_profiles", "instance_profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_instance_profiles.is_snapshot_instance_profiles", "instance_profiles.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotInstanceProfilesDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1 string) string {
	return testAccCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1) + fmt.Sprintf(` 
	data "ibm_is_snapshot_instance_profiles" "is_snapshot_instance_profiles" {
			identifier = ibm_is_snapshot.testacc_snapshot.id
	}
	`)
}
