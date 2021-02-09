/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerClusterFeature_Basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterFeature_basic(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "public_service_endpoint", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "private_service_endpoint", "true"),
				),
			},
			{
				Config: testAccCheckIBMContainerClusterFeature_update(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "public_service_endpoint", "false"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "private_service_endpoint", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterFeature_basic(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name              = "%s"
  datacenter        = "%s"
  default_pool_size = 1
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  timeouts {
	  create = "120m"
  }
}

resource "ibm_container_cluster_feature" "feature" {
  cluster                  = ibm_container_cluster.testacc_cluster.id
  private_service_endpoint = "true"
  timeouts {
    create = "720m"
  }
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterFeature_update(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name              = "%s"
  datacenter        = "%s"
  default_pool_size = 1
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  timeouts {
	create = "120m"
  }
}

resource "ibm_container_cluster_feature" "feature" {
  cluster                 = ibm_container_cluster.testacc_cluster.id
  public_service_endpoint = "false"
  timeouts {
    update = "720m"
  }
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}
