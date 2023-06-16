// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/vpc-beta-go-sdk/vpcbetav1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmIsShareBasic(t *testing.T) {
	var conf vpcbetav1.Share
	name := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.is_share", conf),
				),
			},
			{
				Config: testAccCheckIbmIsShareConfigBasic(name),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIbmIsShareAllArgs(t *testing.T) {
	var conf vpcbetav1.Share

	name := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	shareTargetName := fmt.Sprintf("tf-fs-tg-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	size := acctest.RandIntRange(10, 50)
	sizeUpdate := acctest.RandIntRange(51, 70)
	nameUpdate := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareConfig(vpcname, name, size, shareTargetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.is_share", conf),
					//resource.TestCheckResourceAttr("ibm_is_share.is_share", "iops", strconv.Itoa(iops)),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "name", name),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "size", strconv.Itoa(size)),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "tags.0", "sr01"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "tags.1", "sr02"),
				),
			},
			{
				Config: testAccCheckIbmIsShareConfig(vpcname, nameUpdate, sizeUpdate, shareTargetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					//resource.TestCheckResourceAttr("ibm_is_share.is_share", "iops", strconv.Itoa(iops)),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "size", strconv.Itoa(sizeUpdate)),
				),
			},
		},
	})
}

func TestAccIbmIsShareReplica(t *testing.T) {
	var conf vpcbetav1.Share

	name := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	replicaName := fmt.Sprintf("tf-fsrp-name-%d", acctest.RandIntRange(10, 100))
	shareTargetName := fmt.Sprintf("tf-fs-tg-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	size := acctest.RandIntRange(10, 50)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareConfigReplica(vpcname, name, size, shareTargetName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.replica", conf),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "name", replicaName),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "replication_role", "replica"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "replication_role", "source"),
					resource.TestCheckResourceAttrSet("ibm_is_share.replica", "id"),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "tags.0", "sr01"),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "tags.1", "sr02"),
				),
			},
		},
	})
}

func TestAccIbmIsShareReplicaInline(t *testing.T) {
	var conf vpcbetav1.Share

	name := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	replicaName := fmt.Sprintf("tf-fsrp-name-%d", acctest.RandIntRange(10, 100))
	shareTargetName := fmt.Sprintf("tf-fs-tg-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	size := acctest.RandIntRange(10, 50)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareConfigReplicaInline(vpcname, name, size, shareTargetName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.is_share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "name", replicaName),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "replication_role", "replica"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "replication_role", "source"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share", "id"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "tags.0", "sr01"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "tags.1", "sr02"),
				),
			},
		},
	})
}

func testAccCheckIbmIsShareConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "is_share" {
			zone = "us-south-2"
			size = 200
			name = "%s"
			profile = "%s"
		}
	`, name, acc.ShareProfileName)
}

func testAccCheckIbmIsShareConfig(vpcName, name string, size int, shareTergetName string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "group" {
		is_default = "true"
	}
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_share" "is_share" {
		name = "%s"
		profile = "%s"
		resource_group = data.ibm_resource_group.group.id
		size = %d
		mount_targets {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
		}
		zone = "us-south-2"
		tags = ["sr01", "sr02"]
	}
	`, vpcName, name, acc.ShareProfileName, size, shareTergetName)
}

func testAccCheckIbmIsShareConfigReplica(vpcName, name string, size int, shareTergetName, replicaName string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "group" {
		is_default = "true"
	}
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_share" "is_share" {
		name = "%s"
		profile = "%s"
		resource_group = data.ibm_resource_group.group.id
		size = %d
		mount_targets {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
		}
		zone = "us-south-1"
		tags = ["sr01", "sr02"]
	}
	resource "ibm_is_share" "replica" {
		name = "%s"
		profile = "%s"
		zone = "us-south-2"
    	source_share = ibm_is_share.is_share.id
    	replication_cron_spec = "0 */5 * * *"
		tags = ["sr01", "sr02"]
	}
	`, vpcName, name, acc.ShareProfileName, size, shareTergetName, replicaName, acc.ShareProfileName)
}

func testAccCheckIbmIsShareConfigReplicaInline(vpcName, name string, size int, shareTergetName, replicaName string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "group" {
		is_default = "true"
	}
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_share" "is_share" {
		name = "%s"
		profile = "%s"
		resource_group = data.ibm_resource_group.group.id
		size = %d
		mount_targets {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
		}
		replica_share {
			name = "%s"
			replication_cron_spec = "0 */5 * * *"
			profile = "%s"
			zone = "us-south-3"
		  }
		zone = "us-south-1"
		tags = ["sr01", "sr02"]
	}
	`, vpcName, name, acc.ShareProfileName, size, shareTergetName, replicaName, acc.ShareProfileName)
}

func testAccCheckIbmIsShareExists(n string, obj vpcbetav1.Share) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1BetaAPI()
		if err != nil {
			return err
		}

		getShareOptions := &vpcbetav1.GetShareOptions{}

		getShareOptions.SetID(rs.Primary.ID)

		share, _, err := vpcClient.GetShare(getShareOptions)
		if err != nil {
			return err
		}

		obj = *share
		return nil
	}
}

func testAccCheckIbmIsShareDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1BetaAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_share" {
			continue
		}

		getShareOptions := &vpcbetav1.GetShareOptions{}

		getShareOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetShare(getShareOptions)

		if err == nil {
			return fmt.Errorf("Share still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Share (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
