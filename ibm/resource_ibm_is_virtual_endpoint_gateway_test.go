package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVirtualEndpointGateway_Basic(t *testing.T) {
	var endpointGateway string
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", "my-endpoint-gateway-1"),
				),
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_Import(t *testing.T) {
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "my-endpoint-gateway-1"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_FullySpecified(t *testing.T) {
	var monitor string
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckisVirtualEndpointGatewayConfigFullySpecified(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &monitor),
					resource.TestCheckResourceAttr(name, "name", "my-endpoint-gateway-1"),
				),
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_CreateAfterManualDestroy(t *testing.T) {
	var monitorOne, monitorTwo string
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &monitorOne),
					testAccisVirtualEndpointGatewayManuallyDelete(&monitorOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &monitorTwo),
					func(state *terraform.State) error {
						if monitorOne == monitorTwo {
							return fmt.Errorf("load balancer monitor id is unchanged even after we thought we deleted it ( %s )",
								monitorTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccisVirtualEndpointGatewayManuallyDelete(tfEndpointGwID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		sess, err := testAccProvider.Meta().(ClientSession).VpcV1APIScoped()
		if err != nil {
			return err
		}
		tfEndpointGw := *tfEndpointGwID
		opt := sess.NewDeleteEndpointGatewayOptions(tfEndpointGw)
		response, err := sess.DeleteEndpointGateway(opt)
		if err != nil {
			return fmt.Errorf("Delete Endpoint Gateway failed: %v", response)
		}
		return nil
	}
}

func testAccCheckisVirtualEndpointGatewayDestroy(s *terraform.State) error {
	sess, err := testAccProvider.Meta().(ClientSession).VpcV1APIScoped()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_endpoint_gateway" {
			continue
		}
		opt := sess.NewGetEndpointGatewayOptions(rs.Primary.ID)
		_, response, err := sess.GetEndpointGateway(opt)
		if err == nil {
			return fmt.Errorf("Endpoint Gateway still exists: %v", response)
		}
	}

	return nil
}

func testAccCheckisVirtualEndpointGatewayExists(n string, tfEndpointGwID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No endpoint gateway ID is set")
		}

		sess, err := testAccProvider.Meta().(ClientSession).VpcV1APIScoped()
		if err != nil {
			return err
		}

		opt := sess.NewGetEndpointGatewayOptions(rs.Primary.ID)
		result, response, err := sess.GetEndpointGateway(opt)
		if err != nil {
			return fmt.Errorf("Endpoint Gateway does not exist: %s", response)
		}
		*tfEndpointGwID = *result.ID
		return nil
	}
}

func testAccCheckisVirtualEndpointGatewayConfigBasic() string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "test-vpe-network-vpc"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "test-vpe-subnet1"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%[2]s"
		ipv4_cidr_block = "%[3]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "my-endpoint-gateway-1"
		target {
		  name          = "ibm-dns-server2"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		resource_group = data.ibm_resource_group.test_acc.id
	}`, cisResourceGroup, ISZoneName, ISCIDR)
}

func testAccCheckisVirtualEndpointGatewayConfigFullySpecified() string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "test-vpe-network-vpc"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "test-vpe-subnet1"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%[2]s"
		ipv4_cidr_block = "%[3]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "my-endpoint-gateway-1"
		target {
		  name          = "ibm-dns-server2"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		ips {
		  subnet_id   = ibm_is_subnet.testacc_subnet.id
		  name        = "test-reserved-ip1"
		}
		resource_group = data.ibm_resource_group.test_acc.id
	}`, cisResourceGroup, ISZoneName, ISCIDR)
}
