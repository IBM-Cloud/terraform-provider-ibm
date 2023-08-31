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
	profileProfileType := "predefined"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileDataSourceConfigBasic(profileProfileName, profileProfileDescription, profileProfileType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "profiles_id"),
				),
			},
		},
	})
}

func TestAccIbmSccProfileDataSourceAllArgs(t *testing.T) {
	profileXCorrelationID := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	profileXRequestID := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	profileProfileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))
	profileProfileDescription := fmt.Sprintf("tf_profile_description_%d", acctest.RandIntRange(10, 100))
	profileProfileType := "predefined"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileDataSourceConfig(profileXCorrelationID, profileXRequestID, profileProfileName, profileProfileDescription, profileProfileType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "profiles_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "x_correlation_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "x_request_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "profile_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "profile_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "profile_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "profile_version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "version_group_label"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "latest"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "hierarchy_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "control_parents_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "attachments_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_library_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_library_version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_category"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_parent"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_requirement"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "controls.0.control_specifications_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "default_parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "default_parameters.0.assessment_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "default_parameters.0.assessment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "default_parameters.0.parameter_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "default_parameters.0.parameter_default_value"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "default_parameters.0.parameter_display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile.scc_profile", "default_parameters.0.parameter_type"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProfileDataSourceConfigBasic(profileProfileName string, profileProfileDescription string, profileProfileType string) string {
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

		data "ibm_scc_profile" "scc_profile_instance" {
			profiles_id = ibm_scc_profile.scc_profile_instance.profile_id
			X-Correlation-ID = ibm_scc_profile.scc_profile.x_correlation_id
			X-Request-ID = ibm_scc_profile.scc_profile.x_request_id
		}
	`, profileProfileName, profileProfileDescription, profileProfileType)
}

func testAccCheckIbmSccProfileDataSourceConfig(profileXCorrelationID string, profileXRequestID string, profileProfileName string, profileProfileDescription string, profileProfileType string) string {
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
			x_correlation_id = "%s"
			x_request_id = "%s"
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

		data "ibm_scc_profile" "scc_profile_instance" {
			profiles_id = ibm_scc_profile.scc_profile_instance.profile_id
			X-Correlation-ID = ibm_scc_profile.scc_profile.x_correlation_id
			X-Request-ID = ibm_scc_profile.scc_profile.x_request_id
		}
	`, profileXCorrelationID, profileXRequestID, profileProfileName, profileProfileDescription, profileProfileType)
}
