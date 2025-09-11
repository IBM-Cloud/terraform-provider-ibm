// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMPINetworkPeerBasic(t *testing.T) {
	customerAsn := "64512"
	customerCIDR := "192.168.91.5/30"
	ibmAsn := "64513"
	ibmCIDR := "192.168.91.6/30"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(65, 70))
	vlan := fmt.Sprintf("%d", acctest.RandIntRange(1000, 1500))
	customerAsnUpdate := "64514"
	customerCIDRUpdate := "192.168.91.6/30"
	ibmAsnUpdate := "64515"
	ibmCIDRUpdate := "192.168.91.5/30"
	nameUpdate := fmt.Sprintf("tf_name_updated_%d", acctest.RandIntRange(15, 45))
	vlanUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1500, 3500))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkPeerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeerConfigBasic(customerAsn, customerCIDR, ibmAsn, ibmCIDR, name, vlan),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkPeerExists("ibm_pi_network_peer.network_peer"),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_asn", customerAsn),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_cidr", customerCIDR),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_asn", ibmAsn),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_cidr", ibmCIDR),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_name", name),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_vlan", vlan),
				),
			},
			{
				Config: testAccCheckIBMPINetworkPeerConfigBasic(customerAsnUpdate, customerCIDRUpdate, ibmAsnUpdate, ibmCIDRUpdate, nameUpdate, vlanUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_asn", customerAsnUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_cidr", customerCIDRUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_asn", ibmAsnUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_cidr", ibmCIDRUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_vlan", vlanUpdate),
				),
			},
		},
	})
}

func TestAccIBMPINetworkPeerAllArgs(t *testing.T) {
	customerAsn := "64512"
	customerCIDR := "192.168.91.5/30"
	defaultExportRouteFilter := "allow"
	defaultImportRouteFilter := "allow"
	ibmAsn := "64513"
	ibmCIDR := "192.168.91.6/30"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVar := "dcnetwork_bgp"
	vlan := fmt.Sprintf("%d", acctest.RandIntRange(1000, 3000))
	customerAsnUpdate := "64514"
	customerCIDRUpdate := "192.168.91.6/30"
	defaultExportRouteFilterUpdate := "deny"
	defaultImportRouteFilterUpdate := "deny"
	ibmAsnUpdate := "64515"
	ibmCIDRUpdate := "192.168.91.5/30"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "dcnetwork_bgp"
	vlanUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1500, 3500))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkPeerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeerConfig(customerAsn, customerCIDR, defaultExportRouteFilter, defaultImportRouteFilter, ibmAsn, ibmCIDR, name, typeVar, vlan),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkPeerExists("ibm_pi_network_peer.network_peer"),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_asn", customerAsn),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_cidr", customerCIDR),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_default_export_route_filter", defaultExportRouteFilter),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_default_import_route_filter", defaultImportRouteFilter),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_asn", ibmAsn),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_cidr", ibmCIDR),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_name", name),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_type", typeVar),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_vlan", vlan),
				),
			},
			{
				Config: testAccCheckIBMPINetworkPeerConfig(customerAsnUpdate, customerCIDRUpdate, defaultExportRouteFilterUpdate, defaultImportRouteFilterUpdate, ibmAsnUpdate, ibmCIDRUpdate, nameUpdate, typeVarUpdate, vlanUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_asn", customerAsnUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_customer_cidr", customerCIDRUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_default_export_route_filter", defaultExportRouteFilterUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_default_import_route_filter", defaultImportRouteFilterUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_asn", ibmAsnUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_ibm_cidr", ibmCIDRUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_pi_network_peer.network_peer", "pi_vlan", vlanUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkPeerConfigBasic(customerAsn string, customerCIDR string, ibmAsn string, ibmCIDR string, name string, vlan string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_peer" "network_peer" {
			pi_cloud_instance_id = "%[1]s"
			pi_customer_asn = "%[2]s"
			pi_customer_cidr = "%[3]s"
			pi_ibm_asn = "%[4]s"
			pi_ibm_cidr = "%[5]s"
			pi_name = "%[6]s"
			pi_peer_interface_id = "%[8]s"
			pi_vlan = "%[7]s"
		}
	`, acc.Pi_cloud_instance_id, customerAsn, customerCIDR, ibmAsn, ibmCIDR, name, vlan, acc.Pi_peer_interface_id)
}

func testAccCheckIBMPINetworkPeerConfig(customerAsn string, customerCIDR string, defaultExportRouteFilter string, defaultImportRouteFilter string, ibmAsn string, ibmCIDR string, name string, typeVar string, vlan string) string {
	return fmt.Sprintf(`

		resource "ibm_pi_network_peer" "network_peer" {
			pi_cloud_instance_id = "%[1]s"
			pi_customer_asn = "%[2]s"
			pi_customer_cidr = "%[3]s"
			pi_default_export_route_filter = "%[4]s"
			pi_default_import_route_filter = "%[5]s"
			pi_ibm_asn = "%[6]s"
			pi_ibm_cidr = "%[7]s"
			pi_name = "%[8]s"
			pi_peer_interface_id = "%[11]s"
			pi_type = "%[9]s"
			pi_vlan = %[10]s
		}
	`, acc.Pi_cloud_instance_id, customerAsn, customerCIDR, defaultExportRouteFilter, defaultImportRouteFilter, ibmAsn, ibmCIDR, name, typeVar, vlan, acc.Pi_peer_interface_id)
}

func testAccCheckIBMPINetworkPeerExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		networkC := instance.NewIBMPINetworkPeerClient(context.Background(), sess, parts[0])
		_, err = networkC.GetNetworkPeer(parts[1])
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMPINetworkPeerDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network_peer" {
			continue
		}
		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		networkC := instance.NewIBMPINetworkPeerClient(context.Background(), sess, parts[0])
		_, err = networkC.GetNetworkPeer(parts[1])
		if err == nil {
			return fmt.Errorf("Network peer still exists: %s", parts[1])
		}
	}

	return nil
}
