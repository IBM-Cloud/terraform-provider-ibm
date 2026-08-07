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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISIPSecPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-name-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfipsecc-updated-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.example"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "disabled"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "3600"), // Testing default value
					resource.TestCheckResourceAttrSet(resourceKey, "encapsulation_mode"),
					resource.TestCheckResourceAttrSet(resourceKey, "transform_protocol"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyConfigUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes256"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "7200"), // Testing updated value
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyConfigUpdateAll(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", updatedName),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha384"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "10800"),
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
func TestAccIBMISIPSecPolicy_withResourceGroup(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-name-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.example_with_rg"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyWithResourceGroupConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "group_17"),
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

func checkPolicyDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_ipsec_policy" {
			continue
		}

		getipsecpoptions := &vpcv1.GetIpsecPolicyOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetIpsecPolicy(getipsecpoptions)
		if err == nil {
			return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISIpSecPolicyExists(n, policy string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getipsecpoptions := &vpcv1.GetIpsecPolicyOptions{
			ID: &rs.Primary.ID,
		}
		ipSecPolicy, _, err := sess.GetIpsecPolicy(getipsecpoptions)
		if err != nil {
			return err
		}
		policy = *ipSecPolicy.ID

		return nil
	}
}

func testAccCheckIBMISIPSecPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			pfs = "disabled"
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyConfigUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha512"
			encryption_algorithm = "aes256"
			pfs = "group_14"
			key_lifetime = 7200
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyConfigUpdateAll(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha384"
			encryption_algorithm = "aes128"
			pfs = "group_14" 
			key_lifetime = 10800
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyWithResourceGroupConfig(name string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "group" {
			name = "Default"
		}
		
		resource "ibm_is_ipsec_policy" "example_with_rg" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			pfs = "group_17"
			resource_group = data.ibm_resource_group.group.id
		}
	`, name)
}

func TestAccIBMISIPSecPolicy_Algorithms(t *testing.T) {
	namePrefix := "tfipsecc-alg-"

	// Test different algorithm combinations
	testCases := []struct {
		name                string
		authAlgorithm       string
		encryptionAlgorithm string
		pfs                 string
	}{
		{
			name:                "sha384-aes128-group_15",
			authAlgorithm:       "sha384",
			encryptionAlgorithm: "aes128",
			pfs:                 "group_15",
		},
		{
			name:                "disabled-aes128gcm16-group19",
			authAlgorithm:       "disabled",
			encryptionAlgorithm: "aes128gcm16",
			pfs:                 "group_19",
		},
		{
			name:                "sha512-aes256-group20",
			authAlgorithm:       "sha512",
			encryptionAlgorithm: "aes256",
			pfs:                 "group_20",
		},
		{
			name:                "disabled-aes256gcm16-group_31",
			authAlgorithm:       "disabled",
			encryptionAlgorithm: "aes256gcm16",
			pfs:                 "group_31",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			name := fmt.Sprintf("%s%d", namePrefix, acctest.RandIntRange(10, 100))
			resourceKey := "ibm_is_ipsec_policy.algorithm_test"

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { acc.TestAccPreCheck(t) },
				Providers:    acc.TestAccProviders,
				CheckDestroy: checkPolicyDestroy,
				Steps: []resource.TestStep{
					{
						Config: testAccCheckIBMISIPSecPolicyAlgorithmConfig(name, tc.authAlgorithm, tc.encryptionAlgorithm, tc.pfs),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(resourceKey, "name", name),
							resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", tc.authAlgorithm),
							resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", tc.encryptionAlgorithm),
							resource.TestCheckResourceAttr(resourceKey, "pfs", tc.pfs),
						),
					},
				},
			})
		})
	}
}

func testAccCheckIBMISIPSecPolicyAlgorithmConfig(name, authAlg, encAlg, pfs string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "algorithm_test" {
			name = "%s"
			authentication_algorithm = "%s"
			encryption_algorithm = "%s"
			pfs = "%s"
		}
	`, name, authAlg, encAlg, pfs)
}

func TestAccIBMISIPSecPolicy_MultipleAlgorithms(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-multi-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.multi_algorithm_test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyMultipleAlgorithmsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.0", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.1", "sha384"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.0", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.1", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.0", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.1", "group_15"),
				),
			},
		},
	})
}

func TestAccIBMISIPSecPolicy_MigrateSingularToMultipleAlgorithms(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-migrate-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.migration_test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicySingularConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes256"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "group_17"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyMigrationConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.0", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.1", "sha384"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.0", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.1", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.0", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.1", "group_15"),
				),
			},
		},
	})
}

func TestAccIBMISIPSecPolicy_UpdateMultipleAlgorithms(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-updmulti-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.update_multi_test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyUpdateMultiInitialConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.0", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.1", "sha384"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.0", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.1", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.0", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.1", "group_15"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyUpdateMultiUpdatedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.#", "1"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.0", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.#", "3"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.0", "aes256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.1", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.2", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.#", "3"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.0", "group_15"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.1", "group_16"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.2", "group_17"),
				),
			},
		},
	})
}

func testAccCheckIBMISIPSecPolicyUpdateMultiInitialConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "update_multi_test" {
			name = "%s"
			authentication_algorithms = ["sha512", "sha384"]
			encryption_algorithms = ["aes128", "aes192"]
			pfs_groups = ["group_14", "group_15"]
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyUpdateMultiUpdatedConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "update_multi_test" {
			name = "%s"
			authentication_algorithms = ["sha256"]
			encryption_algorithms = ["aes256", "aes192", "aes128"]
			pfs_groups = ["group_15", "group_16", "group_17"]
		}
	`, name)
}

func TestAccIBMISIPSecPolicy_ComputedFields(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-computed-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.computed_test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					// Check computed fields are set
					resource.TestCheckResourceAttrSet(resourceKey, "encapsulation_mode"),
					resource.TestCheckResourceAttrSet(resourceKey, "transform_protocol"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_controller_url"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_name"),
				),
			},
		},
	})
}

func testAccCheckIBMISIPSecPolicyConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "computed_test" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			pfs = "group_17"
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyMultipleAlgorithmsConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "multi_algorithm_test" {
			name = "%s"
			authentication_algorithms = ["sha512", "sha384"]
			encryption_algorithms = ["aes128", "aes192"]
			pfs_groups = ["group_14", "group_15"]
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicySingularConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "migration_test" {
			name = "%s"
			authentication_algorithm = "sha512"
			encryption_algorithm = "aes256"
			pfs = "group_17"
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyMigrationConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "migration_test" {
			name = "%s"
			authentication_algorithms = ["sha512", "sha384"]
			encryption_algorithms = ["aes128", "aes192"]
			pfs_groups = ["group_14", "group_15"]
		}
	`, name)
}

func TestAccIBMISIPSecPolicy_Scenario1_LegacyNoRegression(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-s1-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfipsecc-s1-upd-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.s1_legacy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyS1LegacyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "disabled"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "3600"),
					resource.TestCheckResourceAttrSet(resourceKey, "encapsulation_mode"),
					resource.TestCheckResourceAttrSet(resourceKey, "transform_protocol"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyS1LegacyUpdateLifetimeConfig(name, 7200),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "disabled"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "7200"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyS1LegacyUpdateAlgoConfig(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", updatedName),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes256"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "14400"),
				),
			},
			{
				ResourceName:      resourceKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISIPSecPolicy_Scenario2_MigrationWithIdempotency(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-s2-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.s2_migration"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyS2SingularConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "3600"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyS2MigratedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.0", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.1", "sha384"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.0", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.1", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.0", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.1", "group_15"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "3600"),
				),
			},
			{
				Config:   testAccCheckIBMISIPSecPolicyS2MigratedConfig(name),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISIPSecPolicy_Scenario3_NewUserFullLifecycle(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-s3-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.s3_newuser"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyS3NewUserConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.0", "sha384"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.1", "sha512"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.0", "aes256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.1", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.#", "3"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.0", "group_14"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.1", "group_15"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.2", "group_16"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "7200"),
					resource.TestCheckResourceAttrSet(resourceKey, "encapsulation_mode"),
					resource.TestCheckResourceAttrSet(resourceKey, "transform_protocol"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyS3NewUserUpdatedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceKey, "name", name),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.#", "1"),
					resource.TestCheckResourceAttr(resourceKey, "authentication_algorithms.0", "sha256"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.#", "3"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.0", "aes128"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.1", "aes192"),
					resource.TestCheckResourceAttr(resourceKey, "encryption_algorithms.2", "aes256"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.0", "group_19"),
					resource.TestCheckResourceAttr(resourceKey, "pfs_groups.1", "group_20"),
					resource.TestCheckResourceAttr(resourceKey, "key_lifetime", "14400"),
				),
			},
			{
				Config:   testAccCheckIBMISIPSecPolicyS3NewUserUpdatedConfig(name),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMISIPSecPolicyS1LegacyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "s1_legacy" {
			name                     = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm     = "aes128"
			pfs                      = "disabled"
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyS1LegacyUpdateLifetimeConfig(name string, lifetime int) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "s1_legacy" {
			name                     = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm     = "aes128"
			pfs                      = "disabled"
			key_lifetime             = %d
		}
	`, name, lifetime)
}

func testAccCheckIBMISIPSecPolicyS1LegacyUpdateAlgoConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "s1_legacy" {
			name                     = "%s"
			authentication_algorithm = "sha512"
			encryption_algorithm     = "aes256"
			pfs                      = "group_14"
			key_lifetime             = 14400
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyS2SingularConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "s2_migration" {
			name                     = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm     = "aes128"
			pfs                      = "group_14"
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyS2MigratedConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "s2_migration" {
			name                      = "%s"
			authentication_algorithms = ["sha256", "sha384"]
			encryption_algorithms     = ["aes128", "aes192"]
			pfs_groups                = ["group_14", "group_15"]
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyS3NewUserConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "s3_newuser" {
			name                      = "%s"
			authentication_algorithms = ["sha384", "sha512"]
			encryption_algorithms     = ["aes256", "aes128"]
			pfs_groups                = ["group_14", "group_15", "group_16"]
			key_lifetime              = 7200
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyS3NewUserUpdatedConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "s3_newuser" {
			name                      = "%s"
			authentication_algorithms = ["sha256"]
			encryption_algorithms     = ["aes128", "aes192", "aes256"]
			pfs_groups                = ["group_19", "group_20"]
			key_lifetime              = 14400
		}
	`, name)
}
