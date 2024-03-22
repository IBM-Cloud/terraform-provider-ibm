// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerALBDataSource_Basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerALBDataSourceBasic(clusterName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.ibm_container_alb.alb", "alb_id", regexp.MustCompile("^(public|private)-cr.*-alb[1-10]$")),
					resource.TestMatchResourceAttr(
						"data.ibm_container_alb.alb", "enable", regexp.MustCompile("^true|false$")),
					resource.TestMatchResourceAttr(
						"data.ibm_container_alb.alb", "alb_type", regexp.MustCompile("^public|private$")),
					resource.TestMatchResourceAttr(
						"data.ibm_container_alb.alb", "state", regexp.MustCompile("^enabled|disabled$")),
					resource.TestCheckResourceAttr(
						"data.ibm_container_alb.alb", "zone", acc.Zone),
					resource.TestMatchResourceAttr(
						"data.ibm_container_alb.alb", "status", regexp.MustCompile("^healthy|warning|disabled$")),
				),
			},
		},
	})
}

func testAccCheckIBMContainerALBDataSourceBasic(clusterName string, enable bool) string {
	config := fmt.Sprintf(`resource "ibm_container_cluster" "testacc_cluster" {
		name       = "%s"
		datacenter = "%s"
		default_pool_size = 1
		machine_type    = "%s"
		hardware        = "shared"
		public_vlan_id  = "%s"
		private_vlan_id = "%s"
		timeouts {
		  create = "120m"
		  update = "120m"
		}
	  }
	  
	  data "ibm_container_alb" "alb" {
		alb_id = ibm_container_cluster.testacc_cluster.albs[0].id
	}`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)

	return config
}
