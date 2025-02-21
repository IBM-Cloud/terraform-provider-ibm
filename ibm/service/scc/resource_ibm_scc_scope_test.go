// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccScopeBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Scope
	instanceID := acc.SccInstanceID
	scopeID := acc.SccAccountID
	resGrpID := acc.SccResourceGroupID
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	environment := "ibm-cloud"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccScopeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccScopeConfig(instanceID, name, description, environment, scopeID, resGrpID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccScopeExists("ibm_scc_scope.scc_scope_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIbmSccScopeAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Scope
	instanceID := acc.SccInstanceID
	scopeID := acc.SccAccountID
	resGrpID := acc.SccResourceGroupID
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	environment := "ibm-cloud"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccScopeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccScopeConfig(instanceID, name, description, environment, scopeID, resGrpID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccScopeExists("ibm_scc_scope.scc_scope_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "environment", environment),
				),
			},
			{
				Config: testAccCheckIBMSccScopeConfig(instanceID, nameUpdate, descriptionUpdate, environment, scopeID, resGrpID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_scope.scc_scope_instance", "environment", environment),
				),
			},
			{
				ResourceName:      "ibm_scc_scope.scc_scope_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccScopeConfig(instanceID string, name string, description string, environment string, scopeId string, resource_group_id string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_scope" "scc_scope_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			environment = "%s"
			properties  = {
				scope_id    = "%s"
				scope_type  = "account"
			}
			exclusions {
				scope_id   = "%s"
				scope_type = "account.resource_group"
			}
		}
	`, instanceID, name, description, environment, scopeId, resource_group_id)
}

func testAccCheckIBMSccScopeExists(n string, obj securityandcompliancecenterapiv3.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		securityAndComplianceCenterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getScopeOptions := &securityandcompliancecenterapiv3.GetScopeOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getScopeOptions.SetInstanceID(parts[0])
		getScopeOptions.SetScopeID(parts[1])

		scope, _, err := securityAndComplianceCenterClient.GetScope(getScopeOptions)
		if err != nil {
			return err
		}

		obj = *scope
		return nil
	}
}

func testAccCheckIBMSccScopeDestroy(s *terraform.State) error {
	securityAndComplianceCenterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_scope" {
			continue
		}

		getScopeOptions := &securityandcompliancecenterapiv3.GetScopeOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getScopeOptions.SetInstanceID(parts[0])
		getScopeOptions.SetScopeID(parts[1])

		// Try to find the key
		_, response, err := securityAndComplianceCenterClient.GetScope(getScopeOptions)

		if err == nil {
			return fmt.Errorf("scc_scope still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_scope (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
