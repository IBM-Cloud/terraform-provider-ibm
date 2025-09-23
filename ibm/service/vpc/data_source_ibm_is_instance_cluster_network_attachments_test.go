// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsInstanceClusterNetworkAttachmentsDataSourceBasic(t *testing.T) {
	instanceClusterNetworkAttachmentInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceClusterNetworkAttachmentsDataSourceConfigBasic(instanceClusterNetworkAttachmentInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.#"),
				),
			},
		},
	})
}

func TestAccIBMIsInstanceClusterNetworkAttachmentsDataSourceAllArgs(t *testing.T) {
	instanceClusterNetworkAttachmentInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	instanceClusterNetworkAttachmentName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceClusterNetworkAttachmentsDataSourceConfig(instanceClusterNetworkAttachmentInstanceID, instanceClusterNetworkAttachmentName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.cluster_network_interface.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.cluster_network_interface.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.cluster_network_interface.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.cluster_network_interface.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.cluster_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.cluster_network_interface.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.cluster_network_interface.0.subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_cluster_network_attachments.is_instance_cluster_network_attachments_instance", "cluster_network_attachments.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsInstanceClusterNetworkAttachmentsDataSourceConfigBasic(instanceClusterNetworkAttachmentInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance" {
			instance_id = "%s"
			cluster_network_interface {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
				id = "0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
				name = "my-cluster-network-interface"
				primary_ip {
					address = "10.1.0.6"
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
					id = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
					name = "my-cluster-network-subnet-reserved-ip"
					resource_type = "cluster_network_subnet_reserved_ip"
				}
				resource_type = "cluster_network_interface"
				subnet {
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
					id = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
					name = "my-cluster-network-subnet"
					resource_type = "cluster_network_subnet"
				}
			}
		}

		data "ibm_is_instance_cluster_network_attachments" "is_instance_cluster_network_attachments_instance" {
			instance_id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance.instance_id
		}
	`, instanceClusterNetworkAttachmentInstanceID)
}

func testAccCheckIBMIsInstanceClusterNetworkAttachmentsDataSourceConfig(instanceClusterNetworkAttachmentInstanceID string, instanceClusterNetworkAttachmentName string) string {
	return fmt.Sprintf(`
		data "ibm_is_instance_cluster_network_attachments" "is_instance_cluster_network_attachments_instance" {
			instance_id = "02c7_a8850825-23f1-43f5-92cc-8e97b1c86313"
		}
	`)
}
