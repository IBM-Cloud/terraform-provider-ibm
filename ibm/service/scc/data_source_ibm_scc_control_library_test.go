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

func TestAccIbmSccControlLibraryDataSourceBasic(t *testing.T) {
	controlLibraryControlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryControlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryControlLibraryType := "predefined"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryDataSourceConfigBasic(controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_libraries_id"),
				),
			},
		},
	})
}

func TestAccIbmSccControlLibraryDataSourceAllArgs(t *testing.T) {
	controlLibraryControlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryControlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryControlLibraryType := "predefined"
	controlLibraryVersionGroupLabel := fmt.Sprintf("tf_version_group_label_%d", acctest.RandIntRange(10, 100))
	controlLibraryControlLibraryVersion := fmt.Sprintf("tf_control_library_version_%d", acctest.RandIntRange(10, 100))
	controlLibraryLatest := "true"
	controlLibraryControlsCount := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryDataSourceConfig(controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType, controlLibraryVersionGroupLabel, controlLibraryControlLibraryVersion, controlLibraryLatest, controlLibraryControlsCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_libraries_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_library_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_library_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_library_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "version_group_label"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_library_version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "latest"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "hierarchy_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_parents_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.control_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.control_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.control_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.control_category"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.control_parent"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.control_requirement"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.status"),
				),
			},
		},
	})
}

func testAccCheckIbmSccControlLibraryDataSourceConfigBasic(controlLibraryControlLibraryName string, controlLibraryControlLibraryDescription string, controlLibraryControlLibraryType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			control_library_name = "control_library_name"
			control_library_description = "control_library_description"
			control_library_type = "predefined"
			controls {
				control_name = "control_name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_tags = [ "control_tags" ]
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
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}

		resource "ibm_scc_control_library" "scc_control_library_instance" {
			control_library_name = "%s"
			control_library_description = "%s"
			control_library_type = "%s"
			controls {
				control_name = "control_name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_tags = [ "control_tags" ]
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
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}

		data "ibm_scc_control_library" "scc_control_library_instance" {
			control_libraries_id = ibm_scc_control_library.scc_control_library_instance.controlLibrary_id
		}
	`, controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType)
}

func testAccCheckIbmSccControlLibraryDataSourceConfig(controlLibraryControlLibraryName string, controlLibraryControlLibraryDescription string, controlLibraryControlLibraryType string, controlLibraryVersionGroupLabel string, controlLibraryControlLibraryVersion string, controlLibraryLatest string, controlLibraryControlsCount string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			control_library_name = "control_library_name"
			control_library_description = "control_library_description"
			control_library_type = "predefined"
			controls {
				control_name = "control_name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_tags = [ "control_tags" ]
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
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}

		resource "ibm_scc_control_library" "scc_control_library_instance" {
			control_library_name = "%s"
			control_library_description = "%s"
			control_library_type = "%s"
			version_group_label = "%s"
			control_library_version = "%s"
			latest = %s
			controls_count = %s
			controls {
				control_name = "control_name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
				control_description = "control_description"
				control_category = "control_category"
				control_parent = "control_parent"
				control_tags = [ "control_tags" ]
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
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}

		data "ibm_scc_control_library" "scc_control_library_instance" {
			control_libraries_id = ibm_scc_control_library.scc_control_library_instance.controlLibrary_id
		}
	`, controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType, controlLibraryVersionGroupLabel, controlLibraryControlLibraryVersion, controlLibraryLatest, controlLibraryControlsCount)
}
