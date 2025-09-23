// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIBMIsPrivatePathServiceGatewayBasic(t *testing.T) {
	var conf vpcv1.PrivatePathServiceGateway
	accessPolicy := "deny"
	accessPolicyUpdate := "review"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))
	nameUpdated := fmt.Sprintf("tf-test-ppsg-updated%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsPrivatePathServiceGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsPrivatePathServiceGatewayExists("ibm_is_private_path_service_gateway.is_private_path_service_gateway", conf),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway.is_private_path_service_gateway", "default_access_policy", accessPolicy),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway.is_private_path_service_gateway", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicyUpdate, nameUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway.is_private_path_service_gateway", "default_access_policy", accessPolicyUpdate),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway.is_private_path_service_gateway", "name", nameUpdated),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_private_path_service_gateway.is_private_path_service_gateway",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name string) string {
	return testAccCheckIBMISPPNLB(vpcname, subnetname, zone, cidr, lbname) + fmt.Sprintf(`
		resource "ibm_is_private_path_service_gateway" "is_private_path_service_gateway" {
			default_access_policy = "%s"
			name = "%s"
			load_balancer = ibm_is_lb.testacc_LB.id
			zonal_affinity = true
			service_endpoints = ["mytestfqdn.internal"]
		}
	`, accessPolicy, name)
}

func testAccCheckIBMIsPrivatePathServiceGatewayExists(n string, obj vpcv1.PrivatePathServiceGateway) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getPrivatePathServiceGatewayOptions := &vpcv1.GetPrivatePathServiceGatewayOptions{}
		getPrivatePathServiceGatewayOptions.SetID(rs.Primary.ID)

		privatePathServiceGateway, _, err := vpcClient.GetPrivatePathServiceGateway(getPrivatePathServiceGatewayOptions)
		if err != nil {
			return err
		}

		obj = *privatePathServiceGateway
		return nil
	}
}

func testAccCheckIBMIsPrivatePathServiceGatewayDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_private_path_service_gateway" {
			continue
		}

		getPrivatePathServiceGatewayOptions := &vpcv1.GetPrivatePathServiceGatewayOptions{}
		getPrivatePathServiceGatewayOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetPrivatePathServiceGateway(getPrivatePathServiceGatewayOptions)

		if err == nil {
			return fmt.Errorf("PrivatePathServiceGateway still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivatePathServiceGateway (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
