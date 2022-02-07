// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

var IBMComputeAutoScalePolicyObjectMask = []string{
	"cooldown",
	"id",
	"name",
	"scaleActions",
	"scaleGroupId",
	"oneTimeTriggers",
	"repeatingTriggers",
	"resourceUseTriggers.watches",
	"triggers",
}

func TestAccIBMComputeAutoScalePolicy_Basic(t *testing.T) {
	var scalepolicy datatypes.Scale_Policy
	groupname := fmt.Sprintf("terraformuat_%d", acctest.RandIntRange(10, 100))
	hostname := acctest.RandString(16)
	policyname := acctest.RandString(16)
	updatedpolicyname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputeAutoScalePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeAutoScalePolicyConfig_basic(groupname, hostname, policyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					testAccCheckIBMComputeAutoScalePolicyAttributes(&scalepolicy, policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_type", "RELATIVE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_amount", "1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "cooldown", "0"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "triggers.#", "3"),
					testAccCheckIBMComputeAutoScalePolicyContainsRepeatingTriggers(&scalepolicy, 2, "0 1 ? * MON,WED *"),
					testAccCheckIBMComputeAutoScalePolicyContainsResourceUseTriggers(&scalepolicy, 120, "80"),
					testAccCheckIBMComputeAutoScalePolicyContainsOneTimeTriggers(&scalepolicy, testOnetimeTriggerDate),
				),
			},

			{
				Config: testAccCheckIBMComputeAutoScalePolicyConfig_updated(groupname, hostname, updatedpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", updatedpolicyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_type", "ABSOLUTE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_amount", "2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "cooldown", "35"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "triggers.#", "3"),
					testAccCheckIBMComputeAutoScalePolicyContainsRepeatingTriggers(&scalepolicy, 2, "0 1 ? * MON,WED,SAT *"),
					testAccCheckIBMComputeAutoScalePolicyContainsResourceUseTriggers(&scalepolicy, 130, "90"),
					testAccCheckIBMComputeAutoScalePolicyContainsOneTimeTriggers(&scalepolicy, testOnetimeTriggerUpdatedDate),
				),
			},
		},
	})
}

func TestAccIBMComputeAutoScaleWithTag(t *testing.T) {
	var scalepolicy datatypes.Scale_Policy
	groupname := fmt.Sprintf("terraformuat_%d", acctest.RandIntRange(10, 100))
	hostname := acctest.RandString(16)
	policyname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputeAutoScalePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeAutoScalePolicyWithTag(groupname, hostname, policyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					testAccCheckIBMComputeAutoScalePolicyAttributes(&scalepolicy, policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "tags.#", "2"),
				),
			},

			{
				Config: testAccCheckIBMComputeAutoScalePolicyWithUpdatedTag(groupname, hostname, policyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeAutoScalePolicyDestroy(s *terraform.State) error {
	service := services.GetScalePolicyService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_autoscale_policy" {
			continue
		}

		scalepolicyId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := service.Id(scalepolicyId).GetObject()

		if err == nil {
			return fmt.Errorf("Auto Scale Policy still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for Auto Scale Policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMComputeAutoScalePolicyContainsResourceUseTriggers(scalePolicy *datatypes.Scale_Policy, period int, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found := false

		for _, scaleResourceUseTrigger := range scalePolicy.ResourceUseTriggers {
			for _, scaleResourceUseWatch := range scaleResourceUseTrigger.Watches {
				if *scaleResourceUseWatch.Metric == "host.cpu.percent" && *scaleResourceUseWatch.Operator == ">" &&
					*scaleResourceUseWatch.Period == period && *scaleResourceUseWatch.Value == value {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("Resource use trigger not found in scale policy")

		}

		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyContainsRepeatingTriggers(scalePolicy *datatypes.Scale_Policy, typeId int, schedule string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found := false

		for _, scaleRepeatingTrigger := range scalePolicy.RepeatingTriggers {
			if *scaleRepeatingTrigger.TypeId == typeId && *scaleRepeatingTrigger.Schedule == schedule {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("Repeating trigger %d with schedule %s not found in scale policy", typeId, schedule)

		}

		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyContainsOneTimeTriggers(scalePolicy *datatypes.Scale_Policy, testOnetimeTriggerDate string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found := false
		const IBMComputeTimeFormat = "2006-01-02T15:04:05-07:00"
		utcLoc, _ := time.LoadLocation("UTC")

		for _, scaleOneTimeTrigger := range scalePolicy.OneTimeTriggers {
			if scaleOneTimeTrigger.Date.In(utcLoc).Format(IBMComputeTimeFormat) == testOnetimeTriggerDate {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("One time trigger with date %s not found in scale policy", testOnetimeTriggerDate)
		}

		return nil

	}
}

func testAccCheckIBMComputeAutoScalePolicyAttributes(scalepolicy *datatypes.Scale_Policy, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *scalepolicy.Name != policyname {
			return fmt.Errorf("Bad name: %s", *scalepolicy.Name)
		}

		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyExists(n string, scalepolicy *datatypes.Scale_Policy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		scalepolicyId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetScalePolicyService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		foundScalePolicy, err := service.Id(scalepolicyId).Mask(strings.Join(IBMComputeAutoScalePolicyObjectMask, ",")).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundScalePolicy.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*scalepolicy = foundScalePolicy
		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyConfig_basic(groupname, hostname, policyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 0
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
        os_reference_code = "DEBIAN_9_64"
        local_disk = false
        datacenter = "dal09"
    }
}

resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "RELATIVE"
    scale_amount = 1
    cooldown = 0
    scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "80"
                    period = 120
        }
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED *"
    }

}`, groupname, hostname, policyname, testOnetimeTriggerDate)
}

const IBMComputeTestTimeFormat = string("2006-01-02T15:04:05-07:00")

var utcLoc, _ = time.LoadLocation("UTC")

var testOnetimeTriggerDate = time.Now().In(utcLoc).AddDate(0, 0, 1).Format(IBMComputeTestTimeFormat)

func testAccCheckIBMComputeAutoScalePolicyConfig_updated(groupname, hostname, updatedpolicyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
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
        os_reference_code = "DEBIAN_9_64"
        local_disk = false
        datacenter = "dal09"
    }
}
resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "ABSOLUTE"
    scale_amount = 2
    cooldown = 35
    scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "90"
                    period = 130
        }
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED,SAT *"
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
    }
}`, groupname, hostname, updatedpolicyname, testOnetimeTriggerUpdatedDate)
}

func testAccCheckIBMComputeAutoScalePolicyWithTag(groupname, hostname, policyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
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
        os_reference_code = "DEBIAN_9_64"
        local_disk = false
        datacenter = "dal09"
    }
}

resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "RELATIVE"
    scale_amount = 1
    cooldown = 30
    scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "80"
                    period = 120
        }
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED *"
	}
	tags = ["one", "two"]

}`, groupname, hostname, policyname, testOnetimeTriggerDate)
}

func testAccCheckIBMComputeAutoScalePolicyWithUpdatedTag(groupname, hostname, policyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
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
        os_reference_code = "DEBIAN_9_64"
        local_disk = false
        datacenter = "dal09"
    }
}
resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "ABSOLUTE"
    scale_amount = 2
    cooldown = 35
    scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "90"
                    period = 130
        }
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED,SAT *"
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
	}
	tags = ["one", "two", "three"]
}`, groupname, hostname, policyname, testOnetimeTriggerUpdatedDate)
}

var testOnetimeTriggerUpdatedDate = time.Now().In(utcLoc).AddDate(0, 0, 2).Format(IBMComputeTestTimeFormat)
