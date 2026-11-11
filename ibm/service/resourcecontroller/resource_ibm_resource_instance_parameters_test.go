// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Full test: creates a resource instance with ibm_resource_instance, then
// manages its parameters with ibm_resource_instance_parameters.
//
// Defaults: an SCC Workload Protection (sysdig-secure) free-trial instance in
// acc.RegionName, with parameters {"enable_cspm": false}. Override with
// IBM_RESOURCE_INSTANCE_PARAMETERS_SERVICE, _PLAN and IBM_RESOURCE_INSTANCE_PARAMETERS
// for another service whose broker accepts updates.
//
// Steps:
//  1. create the instance and apply the parameters (waits for async updates)
//  2. same config again: empty plan (only managed keys are compared)
//  3. change destroy_parameters only: in-place update, no replacement
//  4. import by instance GUID
//  5. delete the instance outside Terraform: the next refresh drops the
//     parameters resource from state (404/410/removed handling)
//
// The instance is deleted at the end (CheckDestroy).
func TestAccIBMResourceInstanceParametersFull(t *testing.T) {
	service := envOr("IBM_RESOURCE_INSTANCE_PARAMETERS_SERVICE", "sysdig-secure")
	plan := envOr("IBM_RESOURCE_INSTANCE_PARAMETERS_PLAN", "free-trial")
	params := envOr("IBM_RESOURCE_INSTANCE_PARAMETERS", `{"enable_cspm":false}`)
	name := fmt.Sprintf("tf-rip-%d", acctest.RandIntRange(10000, 99999))
	res := "ibm_resource_instance_parameters.params"
	instance := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		CheckDestroy:             testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInstanceParametersFullConfig(name, service, plan, params, `{}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(instance),
					resource.TestCheckResourceAttrPair(res, "id", instance, "guid"),
					resource.TestCheckResourceAttrPair(res, "resource_instance_id", instance, "guid"),
					resource.TestCheckResourceAttr(res, "parameters", params),
					resource.TestCheckResourceAttr(res, "state", "active"),
					testAccCheckResourceInstanceHasParameters(instance, params),
				),
			},
			{
				Config:   testAccResourceInstanceParametersFullConfig(name, service, plan, params, `{}`),
				PlanOnly: true,
			},
			{
				Config: testAccResourceInstanceParametersFullConfig(name, service, plan, params, params),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(res, "destroy_parameters", params),
					resource.TestCheckResourceAttr(res, "state", "active"),
				),
			},
			{
				ResourceName:            res,
				ImportState:             true,
				ImportStateIdFunc:       testAccResourceInstanceGUID(instance),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameters", "destroy_parameters", "timeouts"},
			},
			{
				Config:             testAccResourceInstanceParametersFullConfig(name, service, plan, params, params),
				PreConfig:          func() { testAccDeleteResourceInstanceOutside(t, name) },
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func testAccResourceInstanceParametersFullConfig(name, service, plan, params, destroyParams string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "group" {
			is_default = true
		}

		resource "ibm_resource_instance" "instance" {
			name              = "%s"
			service           = "%s"
			plan              = "%s"
			location          = "%s"
			resource_group_id = data.ibm_resource_group.group.id
		}

		resource "ibm_resource_instance_parameters" "params" {
			resource_instance_id = ibm_resource_instance.instance.guid
			parameters           = %q
			destroy_parameters   = %q
		}
	`, name, service, plan, acc.RegionName, params, destroyParams)
}

func testAccResourceInstanceGUID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("Not found: %s", n)
		}
		return rs.Primary.Attributes["guid"], nil
	}
}

// testAccCheckResourceInstanceHasParameters reads the instance from the Resource
// Controller and checks every key the test sent, when the service echoes it back.
func testAccCheckResourceInstanceHasParameters(n, params string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		client, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		instance, _, err := client.GetResourceInstance(client.NewGetResourceInstanceOptions(rs.Primary.Attributes["guid"]))
		if err != nil {
			return err
		}
		var want map[string]any
		if err := json.Unmarshal([]byte(params), &want); err != nil {
			return err
		}
		for k, v := range want {
			if got, ok := instance.Parameters[k]; ok && fmt.Sprint(got) != fmt.Sprint(v) {
				return fmt.Errorf("parameter %s: got %v, want %v", k, got, v)
			}
		}
		return nil
	}
}

// testAccDeleteResourceInstanceOutside deletes the test instance directly, to
// check that the next refresh handles a gone instance.
func testAccDeleteResourceInstanceOutside(t *testing.T, name string) {
	client, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		t.Fatal(err)
	}
	list, _, err := client.ListResourceInstances(&rc.ListResourceInstancesOptions{Name: &name})
	if err != nil || len(list.Resources) != 1 {
		t.Fatalf("find instance %s: %v (found %d)", name, err, len(list.Resources))
	}
	opts := client.NewDeleteResourceInstanceOptions(*list.Resources[0].GUID)
	opts.SetRecursive(true)
	if _, err := client.DeleteResourceInstance(opts); err != nil {
		t.Fatalf("delete instance %s: %v", name, err)
	}
}

// Basic test on an existing instance. Set IBM_RESOURCE_INSTANCE_PARAMETERS_INSTANCE_ID to an instance that accepts the
// JSON in IBM_RESOURCE_INSTANCE_PARAMETERS (for example an SCC Workload Protection
// instance with {"enable_cspm":false}). The test is skipped when either is unset.
func TestAccIBMResourceInstanceParametersBasic(t *testing.T) {
	instanceID := os.Getenv("IBM_RESOURCE_INSTANCE_PARAMETERS_INSTANCE_ID")
	params := os.Getenv("IBM_RESOURCE_INSTANCE_PARAMETERS")
	if instanceID == "" || params == "" {
		t.Skip("IBM_RESOURCE_INSTANCE_PARAMETERS_INSTANCE_ID and IBM_RESOURCE_INSTANCE_PARAMETERS must be set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInstanceParametersConfig(instanceID, params),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_instance_parameters.test", "id", instanceID),
					resource.TestCheckResourceAttr("ibm_resource_instance_parameters.test", "parameters", params),
					resource.TestCheckResourceAttrSet("ibm_resource_instance_parameters.test", "state"),
				),
			},
			{
				// A second plan with the same config must be empty: server-side
				// fields outside the managed keys never show as drift.
				Config:   testAccResourceInstanceParametersConfig(instanceID, params),
				PlanOnly: true,
			},
			{
				ResourceName:      "ibm_resource_instance_parameters.test",
				ImportState:       true,
				ImportStateId:     instanceID,
				ImportStateVerify: true,
				// Parameters are not read back as a whole: only configured keys are managed.
				ImportStateVerifyIgnore: []string{"parameters", "timeouts"},
			},
		},
	})
}

func TestAccIBMResourceInstanceParametersInvalidJSON(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceInstanceParametersConfig("00000000-0000-0000-0000-000000000000", "not-json"),
				ExpectError: regexp.MustCompile("parameters must be a JSON object"),
			},
		},
	})
}

// destroy_parameters is checked at plan time, not first during destroy.
func TestAccIBMResourceInstanceParametersInvalidDestroyJSON(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `
		resource "ibm_resource_instance_parameters" "test" {
			resource_instance_id = "00000000-0000-0000-0000-000000000000"
			parameters           = jsonencode({ enable_cspm = true })
			destroy_parameters   = "{enable_cspm: false"
		}`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("destroy_parameters must be a JSON object"),
			},
		},
	})
}

func testAccResourceInstanceParametersConfig(instanceID, params string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance_parameters" "test" {
			resource_instance_id = "%s"
			parameters           = %q
		}
	`, instanceID, params)
}
