// Copyright IBM Corp. 2023 All Rights Reserved.
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
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccProfileBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Profile
	profileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileDescription := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileType := "predefined"
	profileNameUpdate := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileDescriptionUpdate := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileTypeUpdate := "custom"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfigBasic(profileName, profileDescription, profileType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileExists("ibm_scc_profile.scc_profile", conf),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_name", profileName),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_description", profileDescription),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_type", profileType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfigBasic(profileNameUpdate, profileDescriptionUpdate, profileTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_name", profileNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_description", profileDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_type", profileTypeUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccProfileAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Profile
	profileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileDescription := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileType := "predefined"
	profileNameUpdate := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileDescriptionUpdate := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileTypeUpdate := "custom"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfig(profileName, profileDescription, profileType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileExists("ibm_scc_profile.scc_profile", conf),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_name", profileName),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_description", profileDescription),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_type", profileType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileConfig(profileNameUpdate, profileDescriptionUpdate, profileTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_name", profileNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_description", profileDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile.scc_profile", "profile_type", profileTypeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_profile.scc_profile",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccProfileConfigBasic(profileName string, profileDescription string, profileType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_profile" "scc_profile_instance" {
			profile_name = "profile_name"
			profile_description = "profile_description"
			profile_type = "predefined"
			controls {
				control_library_id = "e98a56ff-dc24-41d4-9875-1e188e2da6cd"
				control_id = "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"
				control_library_version = "control_library_version"
				control_name = "control_name"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_requirement = true
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_specifications_count = 1
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					componenet_name = "componenet_name"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments_count = 1
					assessments {
						assessment_id = "assessment_id"
						assessment_method = "assessment_method"
						assessment_type = "assessment_type"
						assessment_description = "assessment_description"
						parameter_count = 1
						parameters {
							parameter_name = "parameter_name"
							parameter_display_name = "parameter_display_name"
							parameter_type = "string"
						}
					}
				}
			}
			default_parameters {
				assessment_type = "assessment_type"
				assessment_id = "assessment_id"
				parameter_name = "parameter_name"
				parameter_default_value = "parameter_default_value"
				parameter_display_name = "parameter_display_name"
				parameter_type = "string"
			}
		}
		resource "ibm_scc_profile" "scc_profile_instance" {
			profile_name = "%s"
			profile_description = "%s"
			profile_type = "%s"
			controls {
				control_library_id = "e98a56ff-dc24-41d4-9875-1e188e2da6cd"
				control_id = "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"
				control_library_version = "control_library_version"
				control_name = "control_name"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_requirement = true
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_specifications_count = 1
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					componenet_name = "componenet_name"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments_count = 1
					assessments {
						assessment_id = "assessment_id"
						assessment_method = "assessment_method"
						assessment_type = "assessment_type"
						assessment_description = "assessment_description"
						parameter_count = 1
						parameters {
							parameter_name = "parameter_name"
							parameter_display_name = "parameter_display_name"
							parameter_type = "string"
						}
					}
				}
			}
			default_parameters {
				assessment_type = "assessment_type"
				assessment_id = "assessment_id"
				parameter_name = "parameter_name"
				parameter_default_value = "parameter_default_value"
				parameter_display_name = "parameter_display_name"
				parameter_type = "string"
			}
		}
	`, profileName, profileDescription, profileType)
}

func testAccCheckIbmSccProfileConfig(profileName string, profileDescription string, profileType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_profile" "scc_profile_instance" {
			profile_name = "profile_name"
			profile_description = "profile_description"
			profile_type = "predefined"
			controls {
				control_library_id = "e98a56ff-dc24-41d4-9875-1e188e2da6cd"
				control_id = "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"
				control_library_version = "control_library_version"
				control_name = "control_name"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_requirement = true
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_specifications_count = 1
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					componenet_name = "componenet_name"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments_count = 1
					assessments {
						assessment_id = "assessment_id"
						assessment_method = "assessment_method"
						assessment_type = "assessment_type"
						assessment_description = "assessment_description"
						parameter_count = 1
						parameters {
							parameter_name = "parameter_name"
							parameter_display_name = "parameter_display_name"
							parameter_type = "string"
						}
					}
				}
			}
			default_parameters {
				assessment_type = "assessment_type"
				assessment_id = "assessment_id"
				parameter_name = "parameter_name"
				parameter_default_value = "parameter_default_value"
				parameter_display_name = "parameter_display_name"
				parameter_type = "string"
			}
		}

		resource "ibm_scc_profile" "scc_profile_instance" {
			profile_name = "%s"
			profile_description = "%s"
			profile_type = "%s"
			controls {
				control_library_id = "e98a56ff-dc24-41d4-9875-1e188e2da6cd"
				control_id = "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"
				control_library_version = "control_library_version"
				control_name = "control_name"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_requirement = true
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_specifications_count = 1
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					componenet_name = "componenet_name"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments_count = 1
					assessments {
						assessment_id = "assessment_id"
						assessment_method = "assessment_method"
						assessment_type = "assessment_type"
						assessment_description = "assessment_description"
						parameter_count = 1
						parameters {
							parameter_name = "parameter_name"
							parameter_display_name = "parameter_display_name"
							parameter_type = "string"
						}
					}
				}
			}
			default_parameters {
				assessment_type = "assessment_type"
				assessment_id = "assessment_id"
				parameter_name = "parameter_name"
				parameter_default_value = "parameter_default_value"
				parameter_display_name = "parameter_display_name"
				parameter_type = "string"
			}
		}
	`, profileName, profileDescription, profileType)
}

func testAccCheckIbmSccProfileExists(n string, obj securityandcompliancecenterapiv3.Profile) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getProfileOptions := &securityandcompliancecenterapiv3.GetProfileOptions{}

		getProfileOptions.SetProfileID(rs.Primary.ID)

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

		getProfileOptions.SetProfileID(rs.Primary.ID)

		// Try to find the key
		_, response, err := securityandcompliancecenterapiClient.GetProfile(getProfileOptions)

		if err == nil {
			return fmt.Errorf("scc_profile still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_profile (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
