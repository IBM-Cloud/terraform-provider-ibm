// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms_test

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"testing"

	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSKeyWithPolicyOverridesResource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	payload := "LqMWNtSi3Snr4gFNO0PsFFLFRNs57mSXCQE7O2oE+g0="
	resourceName := "ibm_kms_key_with_policy_overrides"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Root Key
				Config: testAccCheckIBMKmsResourceConfig(instanceName, resourceName, keyName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
				),
			},
			{
				// Test Standard Key
				Config: testAccCheckIBMKmsResourceConfig(instanceName, resourceName, keyName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
				),
			},
			{
				// Test Imported Root Key
				Config: testAccCheckIBMKmsResourceImportConfig(instanceName, resourceName, keyName, false, payload),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
				),
			},
			{
				// Test Imported Standard Key
				Config: testAccCheckIBMKmsResourceImportConfig(instanceName, resourceName, keyName, true, payload),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
				),
			},
		},
	})
}

// Test for valid expiration date for create key operation
func TestAccIBMKMSKeyWithPolicyOverridesResource_ValidExpDate(t *testing.T) {

	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	hours := time.Duration(rand.Intn(24) + 1)
	mins := time.Duration(rand.Intn(60) + 1)
	sec := time.Duration(rand.Intn(60) + 1)
	loc, _ := time.LoadLocation("UTC")
	expirationDateValid := ((time.Now().In(loc).Add(time.Hour*hours + time.Minute*mins + time.Second*sec)).Format(time.RFC3339))
	resourceName := "ibm_kms_key_with_policy_overrides"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCreateStandardKeyConfig(instanceName, resourceName, keyName, expirationDateValid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName+".test", "key_name", keyName),
					resource.TestCheckResourceAttr(resourceName+".test", "expiration_date", expirationDateValid),
				),
			},
			{
				Config: testAccCheckIBMKmsCreateRootKeyConfig(instanceName, resourceName, keyName, expirationDateValid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName+".test", "key_name", keyName),
					resource.TestCheckResourceAttr(resourceName+".test", "expiration_date", expirationDateValid),
				),
			},
		},
	})
}

// Test for invalid expiration date for create key operation
func TestAccIBMKMSKeyWithPolicyOverridesResource_InvalidExpDate(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	hours := time.Duration(rand.Intn(24) + 1)
	mins := time.Duration(rand.Intn(60) + 1)
	sec := time.Duration(rand.Intn(60) + 1)
	expirationDateInvalid := (time.Now().Add(time.Hour*hours + time.Minute*mins + time.Second*sec)).String()
	resourceName := "ibm_kms_key_with_policy_overrides"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMKmsCreateStandardKeyConfig(instanceName, resourceName, keyName, expirationDateInvalid),
				ExpectError: regexp.MustCompile("Invalid time format"),
			},
			{
				Config:      testAccCheckIBMKmsCreateRootKeyConfig(instanceName, resourceName, keyName, expirationDateInvalid),
				ExpectError: regexp.MustCompile("Invalid time format"),
			},
		},
	})
}

// Test for Valid/Invalid policy for create key operation
func TestAccIBMKMSKeyWithPolicyOverridesResource_Policies(t *testing.T) {

	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	enabled_rotation := true
	valid_interval := rand.Intn(12) + 1
	invalid_interval := valid_interval + 12
	enabled_dual_auth := false

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Valid Interval
				Config: testAccCheckIBMKmsKeyWithPolicyOverridesAllPolicies(instanceName, keyName, false, enabled_rotation, valid_interval, enabled_dual_auth),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "rotation.0.interval_month", fmt.Sprint(valid_interval)),
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "rotation.0.enabled", strconv.FormatBool(enabled_rotation)),
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "dual_auth_delete.0.enabled", strconv.FormatBool(enabled_dual_auth)),
				),
			},
			{
				// Invalid Interval
				Config:      testAccCheckIBMKmsKeyWithPolicyOverridesAllPolicies(instanceName, keyName, false, enabled_rotation, invalid_interval, enabled_dual_auth),
				ExpectError: regexp.MustCompile("must contain a valid int value should be in range"),
			},
			{
				// Invalid(Rotation) Policy on Standard Key
				Config:      testAccCheckIBMKmsKeyWithPolicyOverridesAllPolicies(instanceName, keyName, true, enabled_rotation, valid_interval, enabled_dual_auth),
				ExpectError: regexp.MustCompile("Error while creating key"),
			},
		},
	})
}

func TestAccIBMKMSKeyWithPolicyOverridesResource_update(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	enabled_rotation := true
	rotation_interval := rand.Intn(12) + 1
	enabled_dual_auth := false
	resourceName := "ibm_kms_key_with_policy_overrides"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsResourceConfig(instanceName, resourceName, keyName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
				),
			},
			{
				Config: testAccCheckIBMKmsKeyWithPolicyOverridesAllPolicies(instanceName, keyName, false, enabled_rotation, rotation_interval, enabled_dual_auth),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "rotation.0.interval_month", fmt.Sprint(rotation_interval)),
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "rotation.0.enabled", strconv.FormatBool(enabled_rotation)),
					resource.TestCheckResourceAttr("ibm_kms_key_with_policy_overrides.test", "dual_auth_delete.0.enabled", strconv.FormatBool(enabled_dual_auth)),
				),
			},
		},
	})
}

func testAccCheckIBMKmsKeyWithPolicyOverridesAllPolicies(instanceName string, keyName string, standard_key bool, enabled_rotation bool, rotation_interval int, enabled_dual_auth bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }

	  resource "ibm_kms_key_with_policy_overrides" "test" {
		instance_id = ibm_resource_instance.kp_instance.guid
		key_name       = "%s"
		standard_key   = %t
		rotation {
			enabled = %t
			interval_month = %d
		}
		dual_auth_delete {
			enabled = %t
		}
	  }
`, instanceName, keyName, standard_key, enabled_rotation, rotation_interval, enabled_dual_auth)
}
