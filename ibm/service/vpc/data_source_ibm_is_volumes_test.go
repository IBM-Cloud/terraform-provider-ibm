// Copyright IBM Corp. 2021 All Rights Reserved.
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

func TestAccIBMIsVolumesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVolumesDataSourceConfigFilterByZone(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVolumesDataSourceConfigFilterByName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
		},
	})
}
func TestAccIBMIsVolumesFromSnapshotDataSourceBasic(t *testing.T) {
	resName := "data.ibm_is_volumes.is_volumes"
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
				Config: testAccCheckIBMIsVolumesFromSnapshotDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.active"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.busy"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.profile.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.architecture"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.display_name"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.dedicated_host_only"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.family"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.name"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.vendor"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.version"),
				),
			},
		},
	})
}
func testAccCheckIBMIsVolumesFromSnapshotDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1 string) string {
	return testAccCheckIBMISVolumeConfigSnapshot(vpcname, subnetname, sshname, publicKey, volname, name, name1) + fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			volume_name = ibm_is_volume.storage.name
			attachment_state = "unattached"
			encryption = "provider_managed"
			operating_system_family = "not:null"
			operating_system_architecture = "not:null"
		}
	`)
}
func testAccCheckIBMIsVolumesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
		}
	`)
}

func testAccCheckIBMIsVolumesDataSourceConfigFilterByZone() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			zone_name = "us-south-1"
		}
	`)
}

func testAccCheckIBMIsVolumesDataSourceConfigFilterByName() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			volume_name = "worrier-mailable-timpani-scowling"
		}
	`)
}
