package ibm

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccIBMLbShared_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbSharedConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "250"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMLbDedicated_Basic(t *testing.T) {
	t.SkipNow()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbDedicatedConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "connections", "15000"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ha_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "dedicated", "true"),
					resource.TestCheckResourceAttr(
						"ibm_lb.testacc_foobar_lb", "ssl_enabled", "true"),
				),
			},
		},
	})
}

const testAccCheckIBMLbSharedConfig_basic = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "dal09"
    ha_enabled  = false
}`

const testAccCheckIBMLbDedicatedConfig_basic = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 15000
    datacenter    = "dal09"
    ha_enabled  = false
    dedicated = true	
}`
