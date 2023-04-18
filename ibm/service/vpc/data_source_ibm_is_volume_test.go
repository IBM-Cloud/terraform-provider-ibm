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

func TestAccIBMISVolumeDatasource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	zone := "us-south-1"
	resName := "data.ibm_is_volume.testacc_dsvol"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeDataSourceConfig(name, zone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name),
					resource.TestCheckResourceAttr(
						resName, "zone", zone),
					resource.TestCheckResourceAttrSet(
						resName, "active"),
					resource.TestCheckResourceAttrSet(
						resName, "attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "busy"),
					resource.TestCheckResourceAttrSet(
						resName, "created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "resource_group"),
					resource.TestCheckResourceAttrSet(
						resName, "profile"),
				),
			},
		},
	})
}
func TestAccIBMISVolumeDatasource_from_snapshot(t *testing.T) {

	resName := "data.ibm_is_volume.testacc_dsvol"
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
			{
				Config: testAccCheckIBMISVolumeDataSourceFromSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "active"),
					resource.TestCheckResourceAttrSet(
						resName, "attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "busy"),
					resource.TestCheckResourceAttrSet(
						resName, "created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "resource_group"),
					resource.TestCheckResourceAttrSet(
						resName, "profile"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.#"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.architecture"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.display_name"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.dedicated_host_only"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.family"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.name"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.vendor"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.version"),
				),
			},
		},
	})
}
func testAccCheckIBMISVolumeDataSourceFromSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1 string) string {
	return testAccCheckIBMISVolumeConfigSnapshot(vpcname, subnetname, sshname, publicKey, volname, name, name1) + fmt.Sprintf(`
	
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_volume.storage.name
	}`)
}
func testAccCheckIBMISVolumeDataSourceConfig(name, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_volume" "testacc_volume"{
		name = "%s"
		profile = "10iops-tier"
		zone = "%s"
	}
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_volume.testacc_volume.name
	}`, name, zone)
}
