// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/mitchellh/go-homedir"
)

func TestAccIBMContainer_ClusterConfigDataSourceBasic(t *testing.T) {
	homeDir, err := homedir.Dir()
	if err != nil {
		t.Fatalf("Error fetching homedir: %s", err)
	}
	clusterName := fmt.Sprintf("tf-cluster-config-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterDataSourceConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster_config.testacc_ds_cluster", "config_dir", homeDir),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "config_file_path"),
				),
			},
		},
	})
}

func TestAccIBMContainer_ClusterConfigCalicoDataSourceBasic(t *testing.T) {
	homeDir, err := homedir.Dir()
	if err != nil {
		t.Fatalf("Error fetching homedir: %s", err)
	}
	clusterName := fmt.Sprintf("tf-cluster-config-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterCalicoConfigDataSource(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster_config.testacc_ds_cluster", "config_dir", homeDir),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "config_file_path"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "calico_config_file_path"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDataSourceConfig(clustername string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name        	   = "%s"
  datacenter   	   = "%s"
  machine_type     = "%s"
  hardware         = "shared"
  wait_till        = "MasterNodeReady"
  public_vlan_id   = "%s"
  private_vlan_id  = "%s"
}

data "ibm_container_cluster_config" "testacc_ds_cluster" {
  cluster_name_id = ibm_container_cluster.testacc_cluster.id
}`, clustername, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)
}

func testAccCheckIBMContainerClusterCalicoConfigDataSource(clustername string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  wait_till        = "MasterNodeReady"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}

data "ibm_container_cluster_config" "testacc_ds_cluster" {
  cluster_name_id = ibm_container_cluster.testacc_cluster.id
  network         = true
}`, clustername, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)
}
