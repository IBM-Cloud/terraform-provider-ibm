package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMLbaasHealthMonitor_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasHealthMonitorConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "timeout", "4"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "interval", "5"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "max_retries", "2"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "url_path", "/"),
				),
			},
			{
				Config: testAccCheckIBMLbaasHealthMonitorConfig_update(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "timeout", "5"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "interval", "8"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "max_retries", "6"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_hm", "url_path", "/abc")),
			},
		},
	})
}

func TestAccIBMLbaasHealthMonitor_tcp(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasHealthMonitorConfig_tcp(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_tcp", "timeout", "4"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_tcp", "interval", "5"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_tcp", "max_retries", "2"),
				),
			},
			{
				Config: testAccCheckIBMLbaasHealthMonitorConfig_tcp_update(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_tcp", "timeout", "5"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_tcp", "interval", "8"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_health_monitor.lbaas_tcp", "max_retries", "6")),
			},
		},
	})
}

func TestAccIBMLbaasHealthMonitor_InvalidInterval(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaasHealthMonitor_InvalidInterval,
				ExpectError: regexp.MustCompile("must be between 2 and 60"),
			},
		},
	})
}

func TestAccIBMLbaasHealthMonitor_InvalidTimeout(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaasHealthMonitor_Timeout,
				ExpectError: regexp.MustCompile("must be between 1 and 59"),
			},
		},
	})
}

func TestAccIBMLbaasHealthMonitor_InvalidMaxRetries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaasHealthMonitor_MaxRetries,
				ExpectError: regexp.MustCompile("must be between 1 and 10"),
			},
		},
	})
}

func TestAccIBMLbaasHealthMonitor_InvalidURLPath(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaasHealthMonitor_URLPath,
				ExpectError: regexp.MustCompile("should start with"),
			},
		},
	})
}

const testAccCheckIBMLbaasHealthMonitor_InvalidInterval = `
resource "ibm_lbaas_health_monitor" "lbaas_hm" {
    	protocol = "HTTP"
        port = 80
        timeout = 5
        interval = 1
        max_retries = 6
        url_path = "/abc"
        lbaas_id = "90528c28-1516-4e71-8612-42d1602eb006"
        monitor_id = "90528c28-1516-4e71-8612-42d1602eb6"
}
`

const testAccCheckIBMLbaasHealthMonitor_MaxRetries = `
resource "ibm_lbaas_health_monitor" "lbaas_hm" {
	    protocol = "HTTP"
        port = 80
        timeout = 5
        interval = 8
        max_retries = 12
        url_path = "/abc"
        lbaas_id = "90528c28-1516-4e71-8612-42d1602eb006"
        monitor_id = "90528c28-1516-4e71-8612-42d1602eb6"
}
`

const testAccCheckIBMLbaasHealthMonitor_Timeout = `
resource "ibm_lbaas_health_monitor" "lbaas_hm" {
	   protocol = "HTTP"
        port = 80
        timeout = 60
        interval = 8
        max_retries = 6
        url_path = "/abc"
        lbaas_id = "90528c28-1516-4e71-8612-42d1602eb006"
        monitor_id = "90528c28-1516-4e71-8612-42d1602eb6"
}
`

const testAccCheckIBMLbaasHealthMonitor_URLPath = `
resource "ibm_lbaas_health_monitor" "lbaas_hm" {
	    protocol = "HTTP"
        port = 80
        timeout = 60
        interval = 8
        max_retries = 6
        url_path = "abc"
        lbaas_id = "90528c28-1516-4e71-8612-42d1602eb006"
        monitor_id = "90528c28-1516-4e71-8612-42d1602eb6"
}
`

func testAccCheckIBMLbaasHealthMonitorConfig_basic(name string) string {
	return fmt.Sprintf(`

resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "weighted_round_robin"
  }]
}
data "ibm_lbaas" "ds_lbaas" {
    name = "${ibm_lbaas.lbaas.name}"
}
resource "ibm_lbaas_health_monitor" "lbaas_hm" {
	    protocol = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_protocol}"
        port = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_port}"
        timeout = 4
        lbaas_id = "${data.ibm_lbaas.ds_lbaas.id}"
        monitor_id = "${data.ibm_lbaas.ds_lbaas.health_monitors.0.monitor_id}"
}
`, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasHealthMonitorConfig_update(name string) string {
	return fmt.Sprintf(`

resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "weighted_round_robin"
  },
  ]
}
data "ibm_lbaas" "ds_lbaas" {
    name = "${ibm_lbaas.lbaas.name}"
}
resource "ibm_lbaas_health_monitor" "lbaas_hm" {
	    protocol = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_protocol}"
        port = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_port}"
        timeout = 5
        interval = 8
        max_retries = 6
        url_path = "/abc"
        lbaas_id = "${data.ibm_lbaas.ds_lbaas.id}"
        monitor_id = "${data.ibm_lbaas.ds_lbaas.health_monitors.0.monitor_id}"
}
`, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasHealthMonitorConfig_tcp(name string) string {
	return fmt.Sprintf(`

resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
  protocols = [{
    "frontend_protocol" = "TCP"
    "frontend_port" = 9443
    "backend_protocol" = "TCP"
    "backend_port" = 9443
    "load_balancing_method" = "weighted_round_robin"
  }]
}
data "ibm_lbaas" "ds_lbaas" {
    name = "${ibm_lbaas.lbaas.name}"
}
resource "ibm_lbaas_health_monitor" "lbaas_tcp" {
	    protocol = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_protocol}"
        port = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_port}"
        timeout = 4
        lbaas_id = "${data.ibm_lbaas.ds_lbaas.id}"
        monitor_id = "${data.ibm_lbaas.ds_lbaas.health_monitors.0.monitor_id}"
}
`, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasHealthMonitorConfig_tcp_update(name string) string {
	return fmt.Sprintf(`

resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
  protocols = [
  {
    "frontend_protocol" = "TCP"
    "frontend_port" = 9443
    "backend_protocol" = "TCP"
    "backend_port" = 9443
    "load_balancing_method" = "weighted_round_robin"
  }]
}
data "ibm_lbaas" "ds_lbaas" {
    name = "${ibm_lbaas.lbaas.name}"
}
resource "ibm_lbaas_health_monitor" "lbaas_tcp" {
	    protocol = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_protocol}"
        port = "${data.ibm_lbaas.ds_lbaas.protocols.0.backend_port}"
        timeout = 5
        interval = 8
        max_retries = 6
        lbaas_id = "${data.ibm_lbaas.ds_lbaas.id}"
        monitor_id = "${data.ibm_lbaas.ds_lbaas.health_monitors.0.monitor_id}"
}
`, name, lbaasSubnetId)
}
