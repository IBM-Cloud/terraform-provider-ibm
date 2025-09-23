// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamTrustedProfileBasic(t *testing.T) {
	var conf iamidentityv1.TrustedProfile
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileExists("ibm_iam_trusted_profile.iam_trusted_profile_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfileAllArgs(t *testing.T) {
	var conf iamidentityv1.TrustedProfile
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfig(name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileExists("ibm_iam_trusted_profile.iam_trusted_profile_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile_instance", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfig(nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile_instance", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_trusted_profile.iam_trusted_profile_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIBMIamTrustedProfileConfig(name string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
			name = "%s"
			description = "%s"
			lifecycle {
              ignore_changes = [history]
            }
		}
	`, name, description)
}

func testAccCheckIBMIamTrustedProfileExists(n string, obj iamidentityv1.TrustedProfile) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getProfileOptions := &iamidentityv1.GetProfileOptions{}

		getProfileOptions.SetProfileID(rs.Primary.ID)

		trustedProfile, _, err := iamIdentityClient.GetProfile(getProfileOptions)
		if err != nil {
			return err
		}

		obj = *trustedProfile
		return nil
	}
}

func testAccCheckIBMIamTrustedProfileDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile" {
			continue
		}

		getProfileOptions := &iamidentityv1.GetProfileOptions{}

		getProfileOptions.SetProfileID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetProfile(getProfileOptions)

		if err == nil {
			return fmt.Errorf("iam_trusted_profile still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_trusted_profile (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIamTrustedProfileEnityHistoryRecordToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["timestamp"] = "testString"
		model["iam_id"] = "testString"
		model["iam_id_account"] = "testString"
		model["action"] = "testString"
		model["params"] = []string{"testString"}
		model["message"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.EnityHistoryRecord)
	model.Timestamp = core.StringPtr("testString")
	model.IamID = core.StringPtr("testString")
	model.IamIDAccount = core.StringPtr("testString")
	model.Action = core.StringPtr("testString")
	model.Params = []string{"testString"}
	model.Message = core.StringPtr("testString")

	result, err := iamidentity.ResourceIBMIamTrustedProfileEnityHistoryRecordToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
