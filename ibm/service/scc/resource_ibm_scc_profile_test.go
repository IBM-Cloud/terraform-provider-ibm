// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIbmSccProfileBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Profile
	profileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileDescription := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileType := "custom"
	profileNameUpdate := profileName
	profileDescriptionUpdate := profileDescription
	profileTypeUpdate := profileType

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfigBasic(acc.SccInstanceID, profileName, profileDescription, profileType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileExists("ibm_scc_profile.scc_profile_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_name", profileName),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_description", profileDescription),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_type", profileType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfigBasic(acc.SccInstanceID, profileNameUpdate, profileDescriptionUpdate, profileTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_name", profileNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_description", profileDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_type", profileTypeUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccProfileAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Profile
	profileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileDescription := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileType := "custom"
	profileNameUpdate := profileName
	profileDescriptionUpdate := profileDescription
	profileVersion := "0.0.1"
	profileTypeUpdate := profileType

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfig(acc.SccInstanceID, profileName, profileDescription, profileType, profileVersion),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileExists("ibm_scc_profile.scc_profile_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_name", profileName),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_description", profileDescription),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_type", profileType),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_version", profileVersion),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfig(acc.SccInstanceID, profileNameUpdate, profileDescriptionUpdate, profileTypeUpdate, profileVersion),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_name", profileNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_description", profileDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_type", profileTypeUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile_instance", "profile_version", profileVersion),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_profile.scc_profile_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccProfileConfigBasic(instanceID string, profileName string, profileDescription string, profileType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "control_library_name"
			control_library_description = "control_library_description"
			control_library_type = "custom"
			latest = true
			controls {
        control_id ="0d4624f5-f5f6-44ed-9e09-6662e2f4106c"
				control_name = "control-name"
				control_description = "control_description"
				control_category = "control_category"
				control_tags = [ "control_tags" ]
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					component_name = "f3517159-889e-4781-819a-89d89b747c85"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments {
						assessment_id = "rule-a637949b-7e51-46c4-afd4-b96619001bf1"
						assessment_method = "ibm-cloud-rule"
						assessment_type = "automated"
						assessment_description = "assessment_description"
						parameters {
							parameter_display_name = "Sign out due to inactivity in seconds"
              parameter_name         = "session_invalidation_in_seconds"
							parameter_type = "numeric"
						}
					}
				}
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}

		resource "ibm_scc_profile" "scc_profile_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			profile_name = "%s"
			profile_description = "%s"
			profile_type = "%s"
			controls {
				control_library_id = resource.ibm_scc_control_library.scc_control_library_instance.control_library_id
				control_id = resource.ibm_scc_control_library.scc_control_library_instance.controls[0].control_id
			}
		}

	`, instanceID, profileName, profileDescription, profileType)
}

func testAccCheckIbmSccProfileConfig(instanceID string, profileName string, profileDescription string, profileType string, profileVersion string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "control_library_name"
			control_library_description = "control_library_description"
			control_library_type = "custom"
			latest = true
			controls {
        control_id ="0d4624f5-f5f6-44ed-9e09-6662e2f4106c"
				control_name = "control-name"
				control_description = "control_description"
				control_category = "control_category"
				control_tags = [ "control_tags" ]
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					component_name = "f3517159-889e-4781-819a-89d89b747c85"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments {
						assessment_id = "rule-a637949b-7e51-46c4-afd4-b96619001bf1"
						assessment_method = "ibm-cloud-rule"
						assessment_type = "automated"
						assessment_description = "assessment_description"
						parameters {
							parameter_display_name = "Sign out due to inactivity in seconds"
              parameter_name         = "session_invalidation_in_seconds"
							parameter_type = "numeric"
						}
					}
				}
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}
		
		resource "ibm_scc_profile" "scc_profile_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			profile_name = "%s"
			profile_description = "%s"
			profile_type = "%s"
			profile_version = "%s"
			controls {
				control_library_id = resource.ibm_scc_control_library.scc_control_library_instance.control_library_id
				control_id = resource.ibm_scc_control_library.scc_control_library_instance.controls[0].control_id
			}
			default_parameters {
				assessment_type = "automated"
				assessment_id = "rule-a637949b-7e51-46c4-afd4-b96619001bf1"
				parameter_name = "session_invalidation_in_seconds"
				parameter_type = "numeric"
				parameter_default_value = "9"
				parameter_display_name = "Sign out due to inactivity in seconds"
			}
		}

	`, instanceID, profileName, profileDescription, profileType, profileVersion)
}

func testAccCheckIbmSccProfileExists(n string, obj securityandcompliancecenterapiv3.Profile) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getProfileOptions := &securityandcompliancecenterapiv3.GetProfileOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		getProfileOptions.SetInstanceID(id[0])
		getProfileOptions.SetProfileID(id[1])

		profile, _, err := securityandcompliancecenterapiClient.GetProfile(getProfileOptions)
		if err != nil {
			return err
		}

		obj = *profile
		return nil
	}
}

func testAccCheckIbmSccProfileDestroy(s *terraform.State) error {
	securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_profile" {
			continue
		}

		getProfileOptions := &securityandcompliancecenterapiv3.GetProfileOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		getProfileOptions.SetInstanceID(id[0])
		getProfileOptions.SetProfileID(id[1])

		// Try to find the key
		_, response, err := securityandcompliancecenterapiClient.GetProfile(getProfileOptions)

		if err == nil {
			return flex.FmtErrorf("scc_profile still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("Error checking for scc_profile (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
