package ibm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMComputeMonitor_Basic(t *testing.T) {
	var basicMonitor datatypes.Network_Monitor_Version1_Query_Host

	hostname := acctest.RandString(16)
	domain := "terraformmonitoruat.ibm.com"

	queryTypeID1 := "1"
	responseActionID1 := "1"
	waitCycles1 := "5"

	queryTypeID2 := "17"
	responseActionID2 := "2"
	waitCycles2 := "10"

	notifiedUsers := []int{6575505}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeMonitorConfigBasic(hostname, domain, queryTypeID1, responseActionID1, waitCycles1, notifiedUsers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeMonitorExists("ibm_compute_monitor.testacc_basic_monitor", &basicMonitor),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "guest_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "ip_address"),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "query_type_id", queryTypeID1),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "response_action_id", responseActionID1),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "wait_cycles", waitCycles1),
					resource.TestCheckFunc(testAccCheckIBMComputeMonitorNotifiedUsers),
				),
				Destroy: false,
			},

			{
				Config: testAccCheckIBMComputeMonitorConfigBasic(hostname, domain, queryTypeID2, responseActionID2, waitCycles2, notifiedUsers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeMonitorExists("ibm_compute_monitor.testacc_basic_monitor", &basicMonitor),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "guest_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "ip_address"),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "query_type_id", queryTypeID2),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "response_action_id", responseActionID2),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "wait_cycles", waitCycles2),
				),
				Destroy: false,
			},
		},
	})
}

func TestAccIBMComputeMonitorWithTag(t *testing.T) {
	var basicMonitor datatypes.Network_Monitor_Version1_Query_Host

	hostname := acctest.RandString(16)
	domain := "terraformmonitoruat.ibm.com"

	queryTypeID1 := "1"
	responseActionID1 := "1"
	waitCycles1 := "5"

	notifiedUsers := []int{6575505}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeMonitorWithTag(hostname, domain, queryTypeID1, responseActionID1, waitCycles1, notifiedUsers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeMonitorExists("ibm_compute_monitor.testacc_basic_monitor", &basicMonitor),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "guest_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "ip_address"),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "tags.#", "2"),
				),
				Destroy: false,
			},

			{
				Config: testAccCheckIBMComputeMonitorWithUpdatedTag(hostname, domain, queryTypeID1, responseActionID1, waitCycles1, notifiedUsers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeMonitorExists("ibm_compute_monitor.testacc_basic_monitor", &basicMonitor),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "guest_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_monitor.testacc_basic_monitor", "ip_address"),
					resource.TestCheckResourceAttr(
						"ibm_compute_monitor.testacc_basic_monitor", "tags.#", "3"),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIBMComputeMonitorDestroy(s *terraform.State) error {
	service := services.GetNetworkMonitorVersion1QueryHostService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_monitor" {
			continue
		}

		basicMonitorId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the basic monitor
		_, err := service.Id(basicMonitorId).GetObject()

		if err == nil {
			return errors.New("Basic Monitor still exists")
		}
	}

	return nil
}

func testAccCheckIBMComputeMonitorExists(n string, basicMonitor *datatypes.Network_Monitor_Version1_Query_Host) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		basicMonitorId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetNetworkMonitorVersion1QueryHostService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		foundBasicMonitor, err := service.Id(basicMonitorId).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundBasicMonitor.Id)) != rs.Primary.ID {
			return errors.New("Record not found")
		}

		*basicMonitor = foundBasicMonitor

		return nil
	}

}
func testAccCheckIBMComputeMonitorNotifiedUsers(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_monitor" {
			continue
		}

		if n, ok := rs.Primary.Attributes["notified_users.#"]; ok && n != "" && n != "0" {
			return nil
		}
		break
	}
	return errors.New("Basic monitor has no notified users")
}

func testAccCheckIBMComputeMonitorConfigBasic(hostname, domain, queryTypeID, responseActionID, waitCycles string, notifiedUsers []int) string {
	users := []string{}
	for _, v := range notifiedUsers {
		text := strconv.Itoa(v)
		users = append(users, text)
	}
	formattedUser := strings.Join(users, ",")

	config := fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vg-basic-monitor-test" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    dedicated_acct_host_only = true
    local_disk = false
    ipv6_enabled = true
    secondary_ip_count = 4
}
resource "ibm_compute_monitor" "testacc_basic_monitor" {
    guest_id = "${ibm_compute_vm_instance.vg-basic-monitor-test.id}"
    ip_address = "${ibm_compute_vm_instance.vg-basic-monitor-test.ipv4_address}"
    query_type_id = %s
    response_action_id = %s
    wait_cycles = %s     
	notified_users = [%s]
}`, hostname, domain, queryTypeID, responseActionID, waitCycles, formattedUser)
	return config
}

func testAccCheckIBMComputeMonitorWithTag(hostname, domain, queryTypeID, responseActionID, waitCycles string, notifiedUsers []int) string {
	users := []string{}
	for _, v := range notifiedUsers {
		text := strconv.Itoa(v)
		users = append(users, text)
	}
	formattedUser := strings.Join(users, ",")

	config := fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vg-basic-monitor-test" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    dedicated_acct_host_only = true
    local_disk = false
    ipv6_enabled = true
    secondary_ip_count = 4
}
resource "ibm_compute_monitor" "testacc_basic_monitor" {
    guest_id = "${ibm_compute_vm_instance.vg-basic-monitor-test.id}"
    ip_address = "${ibm_compute_vm_instance.vg-basic-monitor-test.ipv4_address}"
    query_type_id = %s
    response_action_id = %s
    wait_cycles = %s     
	notified_users = [%s]
	tags = ["one", "two"]
}`, hostname, domain, queryTypeID, responseActionID, waitCycles, formattedUser)
	return config
}

func testAccCheckIBMComputeMonitorWithUpdatedTag(hostname, domain, queryTypeID, responseActionID, waitCycles string, notifiedUsers []int) string {
	users := []string{}
	for _, v := range notifiedUsers {
		text := strconv.Itoa(v)
		users = append(users, text)
	}
	formattedUser := strings.Join(users, ",")

	config := fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vg-basic-monitor-test" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    dedicated_acct_host_only = true
    local_disk = false
    ipv6_enabled = true
    secondary_ip_count = 4
}
resource "ibm_compute_monitor" "testacc_basic_monitor" {
    guest_id = "${ibm_compute_vm_instance.vg-basic-monitor-test.id}"
    ip_address = "${ibm_compute_vm_instance.vg-basic-monitor-test.ipv4_address}"
    query_type_id = %s
    response_action_id = %s
    wait_cycles = %s     
	notified_users = [%s]
	tags = ["one", "two", "three"]
}`, hostname, domain, queryTypeID, responseActionID, waitCycles, formattedUser)
	return config
}
