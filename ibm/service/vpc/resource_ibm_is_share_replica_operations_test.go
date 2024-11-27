// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmIsShareReplicaOperationsFailover(t *testing.T) {
	var conf vpcv1.Share
	shareName := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	replicaName := fmt.Sprintf("tf-fsrep-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareReplicaOperationsFailover(shareName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.share", "name", shareName),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "replication_role", "replica"),
				),
			},
			{
				Config: testAccCheckIbmIsShareReplicaOperationsFailover(shareName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.share", "name", shareName),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "replication_role", "source"),
					resource.TestCheckResourceAttr("ibm_is_share.share", "replication_role", "replica"),
				),
			},
			{
				Config: testAccCheckIbmIsShareReplicaOperationsFailoverStepTwo(shareName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.share", "name", shareName),
				),
			},
			{
				Config: testAccCheckIbmIsShareReplicaOperationsFailoverStepTwo(shareName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.share", "name", shareName),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "replication_role", "replica"),
				),
			},
		},
	})
}

func TestAccIbmIsShareReplicaOperationsSplit(t *testing.T) {
	var conf vpcv1.Share

	shareName := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	replicaName := fmt.Sprintf("tf-fsrep-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareReplicaOperationsSplit(shareName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.share", "name", shareName),
				),
			},
			{
				Config: testAccCheckIbmIsShareReplicaOperationsSplit(shareName, replicaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.share", "name", shareName),
					resource.TestCheckResourceAttr("ibm_is_share.replica", "replication_role", "none"),
					resource.TestCheckResourceAttr("ibm_is_share.share", "replication_role", "none"),
				),
			},
		},
	})
}

func testAccCheckIbmIsShareReplicaOperationsFailover(shareName, replicaName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "share" {
			zone = "us-south-1"
			size = 200
			name = "%s"
			profile = "%s"
		}
		resource "ibm_is_share" "replica" {
			zone = "us-south-3"
			name = "%s"
			profile = "%s"
			replication_cron_spec = "0 */5 * * *"
			source_share = ibm_is_share.share.id
		}

		resource "ibm_is_share_replica_operations" "test" {
			share_replica = ibm_is_share.replica.id
			fallback_policy = "split"
			timeout = 500
		}
	`, shareName, acc.ShareProfileName, replicaName, acc.ShareProfileName)
}
func testAccCheckIbmIsShareReplicaOperationsFailoverStepTwo(shareName, replicaName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "share" {
			zone = "us-south-1"
			size = 200
			name = "%s"
			profile = "%s"
		}
		resource "ibm_is_share" "replica" {
			zone = "us-south-3"
			name = "%s"
			profile = "%s"
			replication_cron_spec = "0 */5 * * *"
			source_share = ibm_is_share.share.id
		}

		resource "ibm_is_share_replica_operations" "test" {
			share_replica = ibm_is_share.share.id
			fallback_policy = "split"
			timeout = 500
		}
	`, shareName, acc.ShareProfileName, replicaName, acc.ShareProfileName)
}

func testAccCheckIbmIsShareReplicaOperationsSplit(shareName, replicaName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "share" {
			zone = "us-south-1"
			size = 200
			name = "%s"
			profile = "%s"
		}
		resource "ibm_is_share" "replica" {
			zone = "us-south-3"
			name = "%s"
			profile = "%s"
			replication_cron_spec = "0 */5 * * *"
			source_share = ibm_is_share.share.id
		}

		resource "ibm_is_share_replica_operations" "test" {
			share_replica = ibm_is_share.replica.id
			split_share = true
		}
	`, shareName, acc.ShareProfileName, replicaName, acc.ShareProfileName)
}
