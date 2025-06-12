// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccIBMContainerVPCClusterWorkerPoolDataSourceBase(cluster_name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[3]s"
	  vpc_id            = "%[1]s"
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = "%[2]s"
		name      = "us-south-1"
	  }
	}

	resource "ibm_container_vpc_worker_pool" "test_pool" {
		cluster           = ibm_container_vpc_cluster.cluster.id
		vpc_id            = "%[1]s"
		flavor            = "cx2.2x4"
		worker_count      = 1
		worker_pool_name  = "default"
		zones {
			subnet_id = "%[2]s"
			name      = "us-south-1"
		}
		import_on_create  = "true"
		orphan_on_delete = "true"
	}
		`, acc.IksClusterVpcID, acc.IksClusterSubnetID, cluster_name)
}

// TestAccIBMContainerVpcClusterWorkerPoolDataSourceBasic ...
func TestAccIBMContainerVpcClusterWorkerPoolDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-wp-ds-basic-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(name string) string {
	return testAccIBMContainerVPCClusterWorkerPoolDataSourceBase(name) + `
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_worker_pool" {
	    cluster = "${ibm_container_vpc_cluster.cluster.id}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`
}

// TestAccIBMContainerVpcClusterWorkerPoolDataSourceDedicatedHost ...
func TestAccIBMContainerVpcClusterWorkerPoolDataSourceDedicatedHost(t *testing.T) {
	if acc.HostPoolID == "" {
		fmt.Println("[WARN] Skipping TestAccIBMContainerVpcClusterWorkerPoolResourceDedicatedHost - IBM_CONTAINER_DEDICATEDHOST_POOL_ID is unset")
		return
	}

	name := fmt.Sprintf("tf-vpc-wp-ds-dhost-%d", acctest.RandIntRange(10, 100))
	hostpoolID := acc.HostPoolID
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfigDedicatedHost(name, hostpoolID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_worker_pool.vpc_worker_pool", "host_pool_id", hostpoolID),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfigDedicatedHost(name, hostpoolID string) string {
	return testAccCheckIBMVpcContainerWorkerPoolDedicatedHostCreate(
		acc.ClusterName, name, "bx2d.4x16", acc.IksClusterSubnetID, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, hostpoolID) + `
	data "ibm_container_vpc_worker_pool" "vpc_worker_pool" {
	    cluster = "${ibm_container_vpc_worker_pool.vpc_worker_pool.cluster}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.vpc_worker_pool.worker_pool_name}"
		depends_on = [
			ibm_container_vpc_worker_pool.vpc_worker_pool
		]
	}
`
}

// TestAccIBMContainerVpcClusterWorkerPoolDataSourceSecondaryStorage ...
func TestAccIBMContainerVpcClusterWorkerPoolDataSourceSecondaryStorage(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-wp-ds-secstorage-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceSecStorage(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "id"),
					resource.TestCheckResourceAttr(
						"data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "secondary_storage.0.name", acc.WorkerPoolSecondaryStorage)),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceSecStorage(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolSecStorage(name) + `
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_worker_pool" {
	    cluster = "${ibm_container_vpc_worker_pool.test_pool.cluster}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
	`
}

// TestAccIBMContainerVpcClusterWorkerPoolDataSourceKMS ...
func TestAccIBMContainerVpcClusterWorkerPoolDataSourceKMS(t *testing.T) {
	if acc.CrkID == "" {
		fmt.Println("[WARN] Skipping TestAccIBMContainerVpcClusterWorkerPoolDataSourceKMS - IBM_CRK_ID is unset")
		return
	}
	name := fmt.Sprintf("tf-vpc-wp-ds-kms-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceKMS(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "id"),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "autoscale_enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "crk", acc.CrkID),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "kms_instance_id", acc.KmsInstanceID),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceKMS(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolKMS(name) + `
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_worker_pool" {
	    cluster = "${ibm_container_vpc_worker_pool.test_pool.cluster}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`
}

// TestAccIBMContainerVpcClusterWorkerPoolDataSourceKmsAccount ...
func TestAccIBMContainerVpcClusterWorkerPoolDataSourceKmsAccount(t *testing.T) {
	if acc.KmsAccountID == "" {
		fmt.Println("[WARN] Skipping TestAccIBMContainerVpcClusterWorkerPoolDataSourceKmsAccount - IBM_KMS_ACCOUNT_ID is unset")
		return
	}
	name := fmt.Sprintf("tf-vpc-wp-ds-kmsacc-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceKmsAccount(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_kms_worker_pool", "id"),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_kms_worker_pool", "crk", acc.CrkID),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_kms_worker_pool", "kms_instance_id", acc.KmsInstanceID),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceKmsAccount(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolKmsAccount(name) + `
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_kms_worker_pool" {
	    cluster = "${ibm_container_vpc_worker_pool.test_pool.cluster}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`
}
func TestAccIBMContainerVpcOpenshiftClusterWorkerPoolDataSource(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	openshiftFlavour := "bx2.16x64"
	openShiftworkerCount := "2"
	operatingSystem := "REDHAT_8_64"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcOpenshiftClusterWorkerPoolDataSourceConfig(name, openshiftFlavour, openShiftworkerCount, operatingSystem),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.default_worker_pool", "operating_system", operatingSystem),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVpcOpenshiftClusterWorkerPoolDataSourceConfig(name, openshiftFlavour, openShiftworkerCount, operatingSystem string) string {
	return testAccCheckIBMContainerOcpClusterBasic(name, openshiftFlavour, openShiftworkerCount, operatingSystem) + `
	data "ibm_container_vpc_cluster_worker_pool" "default_worker_pool" {
	    cluster = ibm_container_vpc_cluster.cluster.id
	    worker_pool_name = "default"
	}
`
}
