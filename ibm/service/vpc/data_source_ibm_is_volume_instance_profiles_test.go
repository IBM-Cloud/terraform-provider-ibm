// Copyright IBM Corp. 2024 All Rights Reserved.
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

func TestAccIBMIsVolumeInstanceProfilesDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsVolumeInstanceProfilesDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_instance_profiles.is_volume_instance_profiles_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_instance_profiles.is_volume_instance_profiles_instance", "instance_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_instance_profiles.is_volume_instance_profiles_instance", "instance_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_instance_profiles.is_volume_instance_profiles_instance", "instance_profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_instance_profiles.is_volume_instance_profiles_instance", "instance_profiles.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeInstanceProfilesDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1 string) string {
	return testAccCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1) + fmt.Sprintf(` 
	data "ibm_is_volume_instance_profiles" "is_volume_instance_profiles_instance" {
			identifier = ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
		}
	`)
}
