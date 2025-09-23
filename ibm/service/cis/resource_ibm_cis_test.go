// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"strings"
	"testing"

	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/bmxerror"
)

const (
	CisInstanceSuccessStatus      = "active"
	CisInstanceProgressStatus     = "in progress"
	cisInstanceProvisioningStatus = "provisioning"
	CisInstanceInactiveStatus     = "inactive"
	CisInstanceFailStatus         = "failed"
	CisInstanceRemovedStatus      = "removed"
	cisInstanceReclamation        = "pending_reclamation"
)

func TestAccIBMCisInstance_Basic(t *testing.T) {
	t.Skip()
	var cisInstanceOne string
	testName := "test_acc"
	name := "ibm_cis.cis"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisInstance_basic(acc.CisResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(name, &cisInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "internet-svcs"),
					resource.TestCheckResourceAttr(name, "plan", "standard-next"),
					resource.TestCheckResourceAttr(name, "location", "global"),
				),
			},
		},
	})
}

func TestAccIBMCisInstance_CreateAfterManualDestroy(t *testing.T) {
	//t.Parallel()
	t.Skip()
	var cisInstanceOne, cisInstanceTwo string
	testName := "test_acc"
	name := "ibm_cis.cis"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisInstance_basic(acc.CisResourceGroup, testName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisInstanceExists(name, &cisInstanceOne),
					testAccCisInstanceManuallyDelete(&cisInstanceOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisInstance_basic(acc.CisResourceGroup, testName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisInstanceExists(name, &cisInstanceTwo),
					func(state *terraform.State) error {
						if cisInstanceOne == cisInstanceTwo {
							return fmt.Errorf("Cis instance id is unchanged even after we thought we deleted it ( %s )", cisInstanceTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccIBMCisInstance_import(t *testing.T) {
	t.Skip()
	var cisInstanceOne string
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_cis.cis"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisInstance_basic(acc.CisResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(resourceName, &cisInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "internet-svcs"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard-next"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

func testAccCheckIBMCisInstanceDestroy(s *terraform.State) error {
	rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis" {
			continue
		}

		instanceID := rs.Primary.ID
		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		_, response, err := rsConClient.GetResourceInstance(&rsInst)

		if err == nil {
			return fmt.Errorf("Instance still exists: %s", rs.Primary.ID)
		} else if strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s %s", rs.Primary.ID, err, response)
		}
	}
	return nil
}

func testAccCisInstanceManuallyDelete(tfCisId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_ = testAccCisInstanceManuallyDeleteUnwrapped(s, tfCisId)
		return nil
	}
}

func testAccCisInstanceManuallyDeleteUnwrapped(s *terraform.State, tfCisId *string) error {
	rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	instance := *tfCisId
	var instanceId string
	// if Id does not start with CRN, then zoneId/Pool/HealthCheckId passed. Extract InstanceId
	if strings.HasPrefix(instance, "crn") {
		instanceId = instance
	} else {
		_, instanceId, _ = flex.ConvertTftoCisTwoVar(instance)
	}
	recursive := true
	deleteReq := rc.DeleteResourceInstanceOptions{
		ID:        &instanceId,
		Recursive: &recursive,
	}
	response, err := rsConClient.DeleteResourceInstance(&deleteReq)
	if err != nil {
		return fmt.Errorf("[ERROR] Error deleting resource instance: %s %s", err, response)
	}

	_ = &resource.StateChangeConf{
		Pending: []string{CisInstanceProgressStatus, CisInstanceInactiveStatus, CisInstanceSuccessStatus},
		Target:  []string{CisInstanceRemovedStatus},
		Refresh: func() (interface{}, string, error) {
			rsInst := rc.GetResourceInstanceOptions{
				ID: &instanceId,
			}
			instance, response, err := rsConClient.GetResourceInstance(&rsInst)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return instance, CisInstanceSuccessStatus, nil
				}
				return nil, "", err
			}
			if *instance.State == CisInstanceFailStatus {
				return instance, *instance.State, fmt.Errorf("[ERROR] The resource instance %s failed to delete: %v %s", instanceId, err, response)
			}
			return instance, *instance.State, nil
		},
		Timeout:    90 * time.Second,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	if err != nil {
		return fmt.Errorf("[ERROR] Error waiting for resource instance (%s) to be deleted: %s", instanceId, err)
	}
	return nil
}

func testAccCheckIBMCisInstanceExists(n string, tfCisId *string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID

		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		instance, response, err := rsConClient.GetResourceInstance(&rsInst)
		if err != nil {
			if strings.Contains(err.Error(), "Object not found") ||
				strings.Contains(err.Error(), "status code: 404") {
				*tfCisId = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response)
		}
		if strings.Contains(*instance.State, "removed") {
			*tfCisId = ""
			return nil
		}

		*tfCisId = instanceID
		return nil
	}
}

func testAccCheckIBMCisInstance_basic(CisResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	  }
	  
	  resource "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		plan              = "standar-next"
		location          = "global"
	  }
				`, CisResourceGroup, name)
}
