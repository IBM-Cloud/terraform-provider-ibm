package ibm

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMComputeAutoScaleGroup_Basic(t *testing.T) {
	var scalegroup datatypes.Scale_Group
	groupname := fmt.Sprintf("terraformuat_%d", acctest.RandInt())
	hostname := acctest.RandString(16)
	updatedgroupname := fmt.Sprintf("terraformuat_%d", acctest.RandInt())
	updatedhostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeAutoScaleGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScaleGroupConfig_basic(groupname, hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScaleGroupExists("ibm_compute_autoscale_group.sample-http-cluster", &scalegroup),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "name", groupname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "regional_group", "na-usa-central-1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "cooldown", "30"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "minimum_member_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "maximum_member_count", "10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "termination_policy", "CLOSEST_TO_NEXT_CHARGE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "port", "8080"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "health_check.type", "HTTP"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.cores", "1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.memory", "4096"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.network_speed", "1000"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.os_reference_code", "DEBIAN_8_64"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.local_disk", "false"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.disks.0", "25"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.disks.1", "100"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.post_install_script_uri", ""),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.user_metadata", "#!/bin/bash"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScaleGroupConfig_updated(updatedgroupname, updatedhostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScaleGroupExists("ibm_compute_autoscale_group.sample-http-cluster", &scalegroup),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "name", updatedgroupname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "regional_group", "na-usa-central-1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "minimum_member_count", "2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "maximum_member_count", "12"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "termination_policy", "NEWEST"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "cooldown", "35"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "port", "9090"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "health_check.type", "HTTP-CUSTOM"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.hostname", updatedhostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.cores", "2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.memory", "8192"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.network_speed", "100"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.os_reference_code", "CENTOS_7_64"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "virtual_guest_member_template.0.post_install_script_uri", "https://www.google.com"),
				),
			},
		},
	})
}

func TestAccIBMComputeAutoScaleGroupWithTag(t *testing.T) {
	var scalegroup datatypes.Scale_Group
	groupname := fmt.Sprintf("terraformuat_%d", acctest.RandInt())
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeAutoScaleGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScaleGroupWithTag(groupname, hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScaleGroupExists("ibm_compute_autoscale_group.sample-http-cluster", &scalegroup),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "name", groupname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "regional_group", "na-usa-central-1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "cooldown", "30"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "minimum_member_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "maximum_member_count", "10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "termination_policy", "CLOSEST_TO_NEXT_CHARGE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "tags.#", "2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScaleGroupWithUpdatedTag(groupname, hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScaleGroupExists("ibm_compute_autoscale_group.sample-http-cluster", &scalegroup),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "name", groupname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "regional_group", "na-usa-central-1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "minimum_member_count", "2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "maximum_member_count", "12"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "termination_policy", "NEWEST"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "cooldown", "35"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_group.sample-http-cluster", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeAutoScaleGroupDestroy(s *terraform.State) error {
	service := services.GetScaleGroupService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_autoscale_group" {
			continue
		}

		scalegroupId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := service.Id(scalegroupId).GetObject()

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for Auto Scale (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMComputeAutoScaleGroupContainsNetworkVlan(scaleGroup *datatypes.Scale_Group, vlanNumber int, primaryRouterHostname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found := false

		for _, scaleVlan := range scaleGroup.NetworkVlans {
			vlan := *scaleVlan.NetworkVlan

			if *vlan.VlanNumber == vlanNumber && *vlan.PrimaryRouter.Hostname == primaryRouterHostname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf(
				"Vlan number %d with router hostname %s not found in scale group",
				vlanNumber,
				primaryRouterHostname)
		}

		return nil
	}
}

func testAccCheckIBMComputeAutoScaleGroupExists(n string, scalegroup *datatypes.Scale_Group) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		scalegroupId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetScaleGroupService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		foundScaleGroup, err := service.Id(scalegroupId).Mask(strings.Join(IBMComputeAutoScaleGroupObjectMask, ",")).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundScaleGroup.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record %s not found", rs.Primary.ID)
		}

		*scalegroup = foundScaleGroup

		return nil
	}
}

func testAccCheckIBMComputeAutoScaleGroupConfig_basic(groupname, hostname string) string {
	return fmt.Sprintf(`
resource "ibm_lb" "local_lb_01" {
    connections = 250
    datacenter = "dal09"
    ha_enabled = false
}

resource "ibm_lb_service_group" "http_sg" {
    load_balancer_id = "${ibm_lb.local_lb_01.id}"
    allocation = 100
    port = 80
    routing_method = "ROUND_ROBIN"
    routing_type = "HTTP"
}

resource "ibm_compute_autoscale_group" "sample-http-cluster" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 30
    minimum_member_count = 1
    maximum_member_count = 10
    termination_policy = "CLOSEST_TO_NEXT_CHARGE"
    virtual_server_id = "${ibm_lb_service_group.http_sg.id}"
    port = 8080
    health_check = {
        type = "HTTP"
    }
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 1
        memory = 4096
        network_speed = 1000
        hourly_billing = true
        os_reference_code = "DEBIAN_8_64"
        local_disk = false
        disks = [25,100]
        datacenter = "dal09"
        post_install_script_uri = ""
        user_metadata = "#!/bin/bash"
    }
}`, groupname, hostname)
}

func testAccCheckIBMComputeAutoScaleGroupConfig_updated(updatedgroupname, updatedhostname string) string {
	return fmt.Sprintf(`
resource "ibm_lb" "local_lb_01" {
    connections = 250
    datacenter = "dal09"
    ha_enabled = false
}

resource "ibm_lb_service_group" "http_sg" {
    load_balancer_id = "${ibm_lb.local_lb_01.id}"
    allocation = 100
    port = 80
    routing_method = "ROUND_ROBIN"
    routing_type = "HTTP"
}
resource "ibm_compute_autoscale_group" "sample-http-cluster" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 35
    minimum_member_count = 2
    maximum_member_count = 12
    termination_policy = "NEWEST"
    virtual_server_id = "${ibm_lb_service_group.http_sg.id}"
    port = 9090
    health_check = {
        type = "HTTP-CUSTOM"
        custom_method = "GET"
        custom_request = "/healthcheck"
        custom_response = 200
    }
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 2
        memory = 8192
        network_speed = 100
        hourly_billing = true
        os_reference_code = "CENTOS_7_64"
        local_disk = false
        disks = [25,100]
        datacenter = "dal09"
        post_install_script_uri = "https://www.google.com"
        user_metadata = "#!/bin/bash"
    }
}`, updatedgroupname, updatedhostname)
}

func testAccCheckIBMComputeAutoScaleGroupWithTag(groupname, hostname string) string {
	return fmt.Sprintf(`

resource "ibm_compute_autoscale_group" "sample-http-cluster" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 30
    minimum_member_count = 1
    maximum_member_count = 10
    termination_policy = "CLOSEST_TO_NEXT_CHARGE"
    
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 1
        memory = 4096
        network_speed = 1000
        hourly_billing = true
        os_reference_code = "DEBIAN_8_64"
        local_disk = false
        disks = [25,100]
        datacenter = "dal09"
        post_install_script_uri = ""
        user_metadata = "#!/bin/bash"
	}
	tags = ["one", "two"]
}`, groupname, hostname)
}

func testAccCheckIBMComputeAutoScaleGroupWithUpdatedTag(groupname, hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 35
    minimum_member_count = 2
    maximum_member_count = 12
	termination_policy = "NEWEST"
	
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 2
        memory = 8192
        network_speed = 100
        hourly_billing = true
        os_reference_code = "CENTOS_7_64"
        local_disk = false
        disks = [25,100]
        datacenter = "dal09"
        post_install_script_uri = "https://www.google.com"
        user_metadata = "#!/bin/bash"
	}
	tags = ["one", "two", "three"]
}`, groupname, hostname)
}
