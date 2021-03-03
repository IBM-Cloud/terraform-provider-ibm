// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMIAMDynamicRule_Basic(t *testing.T) {
	agname := fmt.Sprintf("ag_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("rule_%d", acctest.RandIntRange(10, 100))
	expiration := 10
	identityProvider := "test-idp.com"
	claim := "blueGroups"
	operator := "CONTAINS"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMDynamicRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMDynamicRuleBasic(agname, name, identityProvider, claim, operator, expiration),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group_dynamic_rule.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_dynamic_rule.accgroup", "identity_provider", identityProvider),
				),
			},
		},
	})
}

func TestAccIBMIAMDynamicRuleimport(t *testing.T) {
	agname := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	expiration := 10
	identityProvider := "test-idp.com"
	claim := "blueGroups"
	operator := "CONTAINS"
	resourceName := "ibm_iam_access_group_dynamic_rule.accgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMDynamicRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMDynamicRuleMultiple(agname, name, identityProvider, claim, operator, expiration),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity_provider", identityProvider),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIAMDynamicRuleDestroy(s *terraform.State) error {
	accClient, err := testAccProvider.Meta().(ClientSession).IAMUUMAPIV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_dynamic_rule" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		grpID := parts[0]
		ruleID := parts[1]

		// Try to find the key
		_, _, err = accClient.DynamicRule().Get(grpID, ruleID)

		if err == nil {
			return fmt.Errorf("Dynamic rule still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for Dynamic rule (%s) to be destroyed: %s", ruleID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMDynamicRuleBasic(agname, name, identityProvider, claim, operator string, expiration int) string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_access_group" "newgroup" {
		name = "%s"
		tags = ["tag1", "tag2"]
	  }
	  
	  resource "ibm_iam_access_group_dynamic_rule" "accgroup" {
		name              = "%s"
		access_group_id   = ibm_iam_access_group.newgroup.id
		expiration        = %d
		identity_provider = "%s"
		conditions {
		  claim    = "%s"
		  operator = "%s"
		  value    = "test-bluegroup-saml"
		}
	  
	  }`, agname, name, expiration, identityProvider, claim, operator)
}

func testAccCheckIBMIAMDynamicRuleMultiple(agname, name, identityProvider, claim, operator string, expiration int) string {
	return fmt.Sprintf(`
	resource "ibm_iam_access_group" "newgroup" {
		name = "%s"
		tags = ["tag1", "tag2"]
	  }
	  
	  resource "ibm_iam_access_group_dynamic_rule" "accgroup" {
		name              = "%s"
		access_group_id   = ibm_iam_access_group.newgroup.id
		expiration        = %d
		identity_provider = "%s"
		conditions {
		  claim    = "%s"
		  operator = "%s"
		  value    = "test-bluegroup-saml"
		}
		conditions {
		  claim    = "%s"
		  operator = "%s"
		  value    = "test-bluegroup-saml"
		}
	  
	  }
	`, agname, name, expiration, identityProvider, claim, operator, claim, operator)
}
