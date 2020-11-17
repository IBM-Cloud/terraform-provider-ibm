package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMPrivateDNSGlbLoadBalancer_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnspn%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbLoadBalancerBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbLoadBalancerExists("ibm_dns_glb_load_balancer.test-pdns-lb", resultprivatedns),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr("ibm_dns_glb_load_balancer.test-pdns-lb", "name", "Test-load-balancer"),
					resource.TestCheckResourceAttr("ibm_dns_glb_load_balancer.test-pdns-lb", "description", "new  lb"),
					resource.TestCheckResourceAttr("ibm_dns_glb_load_balancer.test-pdns-lb", "healthy_origins_threshold", "1"),
				),
			},
			{
				Config: testAccCheckIBMPrivateDNSGlbUpdateLoadBalancerBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbLoadBalancerExists("ibm_dns_glb_load_balancer.test-pdns-monitor", resultprivatedns),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr("ibm_dns_glb_load_balancer.test-pdns-lb", "name", "test-pdns-glb-monitor-update"),
					resource.TestCheckResourceAttr("ibm_dns_glb_load_balancer.test-pdns-lb", "description", "UpdatedMonitordescription"),
					resource.TestCheckResourceAttr("ibm_dns_glb_load_balancer.test-pdns-lb", "healthy_origins_threshold", "1"),
				),
			},
		},
	})

}

func TestAccIBMPrivateDNSGlboadBalancerImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbLoadBalancerBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbLoadBalancerExists("ibm_dns_glb_load_balancer.test-pdns-lb", resultprivatedns),
				),
			},
			{
				ResourceName:      "ibm_dns_glb_load_balancer.test-pdns-lb",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}

func testAccCheckIBMPrivateDNSGlbLoadBalancerBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "default"
    }

    resource "ibm_is_vpc" "test-pdns-vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "test-pdns-glb-monitor-vpc"
		resource_group = data.ibm_resource_group.rg.id
    }

    resource "ibm_resource_instance" "test-pdns-instance" {
		depends_on = [ibm_is_vpc.test-pdns-vpc]
		name = "test-pdns-glb-monitor-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
    }

    resource "ibm_dns_zone" "test-pdns-zone" {
		depends_on = [ibm_resource_instance.test-pdns-instance]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription"
		label = "testlabel-updated"
    }

	resource "ibm_dns_glb_pool" "test-pdns-pool-nw" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "testpool"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "test pool"
		enabled=true
		healthy_origins_threshold=1
		origins {
				name    = "example-1"
				address = "www.google.com"
				enabled = true
				description="origin pool"
		}
    }	
	resource "ibm_dns_glb_load_balancer" "test-pdns-lb" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "Test-load-balancer"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.guid
		description = "new  lb"
		ttl=120
		fallback_pool = ibm_dns_glb_pool.test-pdns-pool.id
		default_pools = [ibm_dns_glb_pool.test-pdns-pool.id]
		az_pools{
		  availability_zone="WEU"
		  pools = [ibm_dns_glb_pool.test-pdns-pool.id]
		}
    }             
	  `, name)

}

func testAccCheckIBMPrivateDNSGlbUpdateLoadBalancerBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "default"
    }

    resource "ibm_is_vpc" "test-pdns-vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "test-pdns-glb-monitor-vpc"
		resource_group = data.ibm_resource_group.rg.id
    }

    resource "ibm_resource_instance" "test-pdns-instance" {
		depends_on = [ibm_is_vpc.test-pdns-vpc]
		name = "test-pdns-glb-monitor-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
    }

    resource "ibm_dns_zone" "test-pdns-zone" {
		depends_on = [ibm_resource_instance.test-pdns-instance]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription"
		label = "testlabel-updated"
    }

	resource "ibm_dns_glb_pool" "test-pdns-pool-nw" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "testpool"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "test pool"
		enabled=true
		healthy_origins_threshold=1
		origins {
				name    = "example-1"
				address = "www.google.com"
				enabled = true
				description="origin pool"
		}
    }	
	resource "ibm_dns_glb_load_balancer" "test-pdns-lb" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "Update load balancer"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.guid
		description = "update lb"
		ttl=120
		fallback_pool = ibm_dns_glb_pool.test-pdns-pool.id
		default_pools = [ibm_dns_glb_pool.test-pdns-pool.id]
		az_pools{
		  availability_zone="WEU"
		  pools = [ibm_dns_glb_pool.test-pdns-pool.id]
		}
    }             
	  `, name)

}

func testAccCheckIBMPrivateDNSGlbMonitorDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_glb_load_balancer" {
			continue
		}
		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getLbOptions := pdnsClient.NewGetLoadBalancerOptions(partslist[0], partslist[1], partslist[2])
		r, res, err := pdnsClient.GetLoadBalancer(getLbOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSGlbLoadBalancerExists(n string, result string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getLbOptions := pdnsClient.NewGetLoadBalancerOptions(partslist[0], partslist[1], partslist[2])
		r, res, err := pdnsClient.GetLoadBalancer(getLbOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
		result = *r.ID
		return nil
	}
}
