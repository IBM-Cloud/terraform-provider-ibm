// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProfileDataSourceBasic(t *testing.T) {
	profileProfileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileProfileDescription := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileProfileType := "custom"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileDataSourceConfigBasic(acc.SccInstanceID, profileProfileName, profileProfileDescription, profileProfileType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "profile_id"),
				),
			},
		},
	})
}

func TestAccIbmSccProfileDataSourceAllArgs(t *testing.T) {
	profileProfileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileProfileDescription := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileProfileType := "custom"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileDataSourceConfig(acc.SccInstanceID, profileProfileName, profileProfileDescription, profileProfileType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "profile_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "profile_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "profile_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "profile_version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "version_group_label"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "latest"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "hierarchy_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "controls_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "control_parents_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "attachments_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile_instance", "controls.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProfileDataSourceConfigBasic(instanceID string, profileProfileName string, profileProfileDescription string, profileProfileType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "control_library_name"
			control_library_description = "control_library_description"
			control_library_type = "custom"
			latest = true
			controls {
				control_name = "control-name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
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

		data "ibm_scc_profile" "scc_profile_instance" {
			profile_id = resource.ibm_scc_profile.scc_profile_instance.profile_id
			instance_id = resource.ibm_scc_profile.scc_profile_instance.instance_id
		}
	`, instanceID, profileProfileName, profileProfileDescription, profileProfileType)
}

func testAccCheckIbmSccProfileDataSourceConfig(instanceID string, profileProfileName string, profileProfileDescription string, profileProfileType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "control_library_name"
			control_library_description = "control_library_description"
			control_library_type = "custom"
			latest = true
			controls {
				control_name = "control-name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
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
			default_parameters {
			}
		}

		data "ibm_scc_profile" "scc_profile_instance" {
			profile_id = resource.ibm_scc_profile.scc_profile_instance.profile_id
			instance_id = resource.ibm_scc_profile.scc_profile_instance.instance_id
		}
	`, instanceID, profileProfileName, profileProfileDescription, profileProfileType)
}
