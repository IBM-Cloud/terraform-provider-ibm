// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

func TestAccIBMContainer_ClusterConfigDataSourceVpcBasic(t *testing.T) {
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
				Config: testAccCheckIBMContainerClusterDataSourceVpcConfig(clusterName),
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
					resource.TestMatchResourceAttr(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "host", regexp.MustCompile("^https://.*private.*")),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDataSourceConfig(clustername string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  wait_till       = "normal"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}

data "ibm_container_cluster_config" "testacc_ds_cluster" {
  cluster_name_id = ibm_container_cluster.testacc_cluster.id
  admin           = %s
}`, clustername, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.IsAdminConfig)
}

func testAccCheckIBMContainerClusterDataSourceVpcConfig(clustername string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "testacc_cluster" {
		name              = "%[1]s"
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 1
		resource_group_id = "%[3]s"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		wait_till = "Normal"
	}

data "ibm_container_cluster_config" "testacc_ds_cluster" {
  cluster_name_id = ibm_container_vpc_cluster.testacc_cluster.id
}`, clustername, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID)
}

func testAccCheckIBMContainerClusterCalicoConfigDataSource(clustername string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  wait_till       = "Normal"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  private_service_endpoint = true
}

data "ibm_container_cluster_config" "testacc_ds_cluster" {
  cluster_name_id = ibm_container_cluster.testacc_cluster.id
  network         = true
  endpoint_type   = "private"
}`, clustername, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)
}
