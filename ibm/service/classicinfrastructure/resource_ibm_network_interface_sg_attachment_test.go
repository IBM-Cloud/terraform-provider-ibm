// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMNetworkInterfaceSGAttachment(t *testing.T) {
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckNetworkInterfaceSGAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccTestAccIBMNetworkInterfaceSGAttachmentConfig(hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_compute_vm_instance.tfuatvm", "hostname", hostname),
					testAccCheckNetworkInterfaceSGAttachmentExists("ibm_network_interface_sg_attachment.ssh"),
					testAccCheckNetworkInterfaceSGAttachmentExists("ibm_network_interface_sg_attachment.http"),
				),
			},
		},
	})
}
func decomposeNetworkSGAttachmentID(attachmentID string) (sgID, interfaceID int, err error) {
	ids := strings.Split(attachmentID, "_")
	if len(ids) != 2 {
		return -1, -1, fmt.Errorf("[ERROR] The ibm_network_interface_sg_attachment id must be of the form <sg_id>_<network_interface_id> but it is %s", attachmentID)
	}
	sgID, err = strconv.Atoi(ids[0])
	if err != nil {
		return -1, -1, fmt.Errorf("[ERROR] Not  a valid security group ID, must be an integer: %s", err)
	}

	interfaceID, err = strconv.Atoi(ids[1])
	if err != nil {
		return -1, -1, fmt.Errorf("[ERROR] Not  a valid network interface ID, must be an integer: %s", err)
	}
	return
}
func testAccCheckNetworkInterfaceSGAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sgID, interfaceID, err := decomposeNetworkSGAttachmentID(rs.Primary.ID)
		if err != nil {
			return err
		}

		sess := acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession()
		service := services.GetNetworkSecurityGroupService(sess)
		bindings, err := service.Id(sgID).GetNetworkComponentBindings()
		if err != nil {
			return err
		}
		for _, b := range bindings {
			if *b.NetworkComponentId == interfaceID {
				return nil
			}
		}
		return fmt.Errorf("[ERROR] No association found between security group %d and network interface %d", sgID, interfaceID)
	}
}

func testAccCheckNetworkInterfaceSGAttachmentDestroy(s *terraform.State) error {
	service := services.GetNetworkSecurityGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_network_interface_sg_attachment" {
			continue
		}

		sgID, interfaceID, err := decomposeNetworkSGAttachmentID(rs.Primary.ID)
		if err != nil {
			return err
		}

		bindings, err := service.Id(sgID).GetNetworkComponentBindings()
		if err != nil {
			return err
		}
		for _, b := range bindings {
			if *b.NetworkComponentId == interfaceID {
				return fmt.Errorf("Association still exists between security group %d and network interface %d", sgID, interfaceID)
			}
		}
		return nil
	}

	return nil
}

func testAccTestAccIBMNetworkInterfaceSGAttachmentConfig(hostname string) string {
	v := fmt.Sprintf(`
		data "ibm_security_group" "allowssh" {
			name = "allow_ssh"
		}
		data "ibm_security_group" "allowhttp" {
			name = "allow_http"
		}
		resource "ibm_compute_vm_instance" "tfuatvm" {
			hostname                 = "%s"
			domain                   = "tfvmuatsg.com"
			os_reference_code        = "DEBIAN_9_64"
			datacenter               = "wdc07"
			network_speed            = 10
			hourly_billing           = true
			private_network_only     = false
			cores                    = 1
			memory                   = 1024
			disks                    = [25, 10, 20]
			dedicated_acct_host_only = true
			local_disk               = false
			ipv6_enabled             = true
			secondary_ip_count       = 4
			notes                    = "VM notes"
		}
		resource "ibm_network_interface_sg_attachment" "ssh" {
			security_group_id    = "${data.ibm_security_group.allowssh.id}"
			network_interface_id = "${ibm_compute_vm_instance.tfuatvm.public_interface_id}"
		}
		resource "ibm_network_interface_sg_attachment" "http" {
			security_group_id    = "${data.ibm_security_group.allowhttp.id}"
			network_interface_id = "${ibm_compute_vm_instance.tfuatvm.public_interface_id}"
		}
		  `, hostname)
	return v
}
