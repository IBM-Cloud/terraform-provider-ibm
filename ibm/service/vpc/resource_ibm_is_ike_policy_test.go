// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISIKEPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("tfike-name-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfike-updated-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ike_policy.example"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkIKEPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIKEPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "dh_group", "14"),
					resource.TestCheckResourceAttr(resourceKey, "ike_version", "1"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "28800"), // Default value
					resource.TestCheckResourceAttrSet(resourceKey, "negotiation_mode"),
					resource.TestCheckResourceAttrSet(resourceKey, "href"),
				),
			},
			{
				Config: testAccCheckIBMISIKEPolicyConfigUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha384"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes256"),
					resource.TestCheckResourceAttr(resourceKey, "dh_group", "15"),
					resource.TestCheckResourceAttr(resourceKey, "ike_version", "2"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "1800"), // Updated value
				),
			},
			{
				Config: testAccCheckIBMISIKEPolicyConfigUpdateAll(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", updatedName),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "dh_group", "16"),
					resource.TestCheckResourceAttr(resourceKey, "ike_version", "2"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "3600"),
				),
			},
			// Test importing the resource
			{
				ResourceName:      resourceKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Test with resource group
func TestAccIBMISIKEPolicy_withResourceGroup(t *testing.T) {
	name := fmt.Sprintf("tfike-rg-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ike_policy.example_with_rg"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkIKEPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIKEPolicyWithResourceGroupConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "dh_group", "19"),
					resource.TestCheckResourceAttr(resourceKey, "ike_version", "1"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_group"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_group_name"),
				),
			},
			// Test importing the resource
			{
				ResourceName:      resourceKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Test computed fields
func TestAccIBMISIKEPolicy_ComputedFields(t *testing.T) {
	name := fmt.Sprintf("tfike-computed-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ike_policy.computed_test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkIKEPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIKEPolicyConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					// Check computed fields are set
					resource.TestCheckResourceAttrSet(resourceKey, "negotiation_mode"),
					resource.TestCheckResourceAttrSet(resourceKey, "href"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_controller_url"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_name"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_group_name"),

					// Additional check for vpn_connections if any exist
					// Note: Since this is just a policy, there might not be connections yet
					resource.TestCheckResourceAttr(resourceKey, "vpn_connections.#", "0"),
				),
			},
		},
	})
}

// Test various algorithm and dh_group combinations
func TestAccIBMISIKEPolicy_Algorithms(t *testing.T) {
	namePrefix := "tfike-alg-"

	// Test different algorithm combinations
	testCases := []struct {
		name                string
		authAlgorithm       string
		encryptionAlgorithm string
		dhGroup             int
		ikeVersion          int
	}{
		{
			name:                "sha384-aes128-group16-ike1",
			authAlgorithm:       "sha384",
			encryptionAlgorithm: "aes128",
			dhGroup:             16,
			ikeVersion:          1,
		},
		{
			name:                "sha256-aes192-group15-ike2",
			authAlgorithm:       "sha256",
			encryptionAlgorithm: "aes192",
			dhGroup:             15,
			ikeVersion:          2,
		},
		{
			name:                "sha384-aes256-group14-ike1",
			authAlgorithm:       "sha384",
			encryptionAlgorithm: "aes256",
			dhGroup:             14,
			ikeVersion:          1,
		},
		{
			name:                "sha512-aes128-group19-ike2",
			authAlgorithm:       "sha512",
			encryptionAlgorithm: "aes128",
			dhGroup:             19,
			ikeVersion:          2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			name := fmt.Sprintf("%s%d", namePrefix, acctest.RandIntRange(10, 100))
			resourceKey := "ibm_is_ike_policy.algorithm_test"

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { acc.TestAccPreCheck(t) },
				Providers:    acc.TestAccProviders,
				CheckDestroy: checkIKEPolicyDestroy,
				Steps: []resource.TestStep{
					{
						Config: testAccCheckIBMISIKEPolicyAlgorithmConfig(name, tc.authAlgorithm, tc.encryptionAlgorithm, tc.dhGroup, tc.ikeVersion),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(resourceKey, "name", name),
							resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", tc.authAlgorithm),
							resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", tc.encryptionAlgorithm),
							resource.TestCheckResourceAttr(resourceKey, "dh_group", fmt.Sprintf("%d", tc.dhGroup)),
							resource.TestCheckResourceAttr(resourceKey, "ike_version", fmt.Sprintf("%d", tc.ikeVersion)),
						),
					},
				},
			})
		})
	}
}

// Test key_lifetime values and validation
func TestAccIBMISIKEPolicy_KeyLifetime(t *testing.T) {
	name := fmt.Sprintf("tfike-lifetime-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ike_policy.lifetime_test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkIKEPolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Use default lifetime value
				Config: testAccCheckIBMISIKEPolicyLifetimeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "28800"), // Default value
				),
			},
			{
				// Update with minimum lifetime value
				Config: testAccCheckIBMISIKEPolicyLifetimeConfigUpdate(name, 1800), // 30 minutes
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "1800"),
				),
			},
			{
				// Update with maximum lifetime value
				Config: testAccCheckIBMISIKEPolicyLifetimeConfigUpdate(name, 86400), // 24 hours
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "86400"),
				),
			},
		},
	})
}

func checkIKEPolicyDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_ike_policy" {
			continue
		}

		getikepoptions := &vpcv1.GetIkePolicyOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetIkePolicy(getikepoptions)
		if err == nil {
			return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISIKEPolicyExists(n, policy string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getikepoptions := &vpcv1.GetIkePolicyOptions{
			ID: &rs.Primary.ID,
		}
		ikePolicy, _, err := sess.GetIkePolicy(getikepoptions)
		if err != nil {
			return err
		}
		policy = *ikePolicy.ID
		return nil
	}
}

func testAccCheckIBMISIKEPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			dh_group = 14
			ike_version = 1
			# key_lifetime defaults to 28800
		}
	`, name)
}

func testAccCheckIBMISIKEPolicyConfigUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha384"
			encryption_algorithm = "aes256"
			dh_group = 15
			ike_version = 2
			key_lifetime = 1800
		}
	`, name)
}

func testAccCheckIBMISIKEPolicyConfigUpdateAll(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha512"
			encryption_algorithm = "aes192"
			dh_group = 16
			ike_version = 2
			key_lifetime = 3600
		}
	`, name)
}

func testAccCheckIBMISIKEPolicyConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "computed_test" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			dh_group = 14
			ike_version = 1
		}
	`, name)
}

func testAccCheckIBMISIKEPolicyWithResourceGroupConfig(name string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "group" {
			name = "Default" # Using Default resource group, change if needed
		}
		
		resource "ibm_is_ike_policy" "example_with_rg" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes192"
			dh_group = 19
			ike_version = 1
			resource_group = data.ibm_resource_group.group.id
		}
	`, name)
}

func testAccCheckIBMISIKEPolicyAlgorithmConfig(name, authAlg, encAlg string, dhGroup, ikeVersion int) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "algorithm_test" {
			name = "%s"
			authentication_algorithm = "%s"
			encryption_algorithm = "%s"
			dh_group = %d
			ike_version = %d
		}
	`, name, authAlg, encAlg, dhGroup, ikeVersion)
}

func testAccCheckIBMISIKEPolicyLifetimeConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "lifetime_test" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			dh_group = 14
			ike_version = 1
		}
	`, name)
}
func testAccCheckIBMISIKEPolicyLifetimeConfigUpdate(name string, lifetime int) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "lifetime_test" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			dh_group = 14
			ike_version = 1
			key_lifetime = %d
		}
	`, name, lifetime)
}
