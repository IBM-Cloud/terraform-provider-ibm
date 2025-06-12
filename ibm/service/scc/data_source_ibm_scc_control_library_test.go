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
	controlLibraryControlLibraryType := "custom"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccControlLibraryDataSourceConfigBasic(acc.SccInstanceID, controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "control_library_id"),
				),
			},
		},
	})
}

func TestAccIbmSccControlLibraryDataSourceAllArgs(t *testing.T) {
	controlLibraryControlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryControlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryControlLibraryType := "custom"
	controlLibraryControlLibraryVersion := fmt.Sprintf("0.0.%d", acctest.RandIntRange(1, 100))
	controlLibraryLatest := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccControlLibraryDataSourceConfig(acc.SccInstanceID, controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType, controlLibraryControlLibraryVersion, controlLibraryLatest),
				Check: resource.ComposeTestCheckFunc(
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
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.control_requirement"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_library.scc_control_library", "controls.0.status"),
				),
			},
		},
	})
}

func testAccCheckIbmSccControlLibraryDataSourceConfigBasic(instanceID string, controlLibraryControlLibraryName string, controlLibraryControlLibraryDescription string, controlLibraryControlLibraryType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "%s"
			control_library_description = "%s"
			control_library_type = "%s"
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

		data "ibm_scc_control_library" "scc_control_library" {
			instance_id = ibm_scc_control_library.scc_control_library_instance.instance_id
			control_library_id = ibm_scc_control_library.scc_control_library_instance.control_library_id
		}

	`, instanceID, controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType)
}

func testAccCheckIbmSccControlLibraryDataSourceConfig(instanceID string, controlLibraryControlLibraryName string, controlLibraryControlLibraryDescription string, controlLibraryControlLibraryType string, controlLibraryControlLibraryVersion string, controlLibraryLatest string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "%s"
			control_library_description = "%s"
			control_library_type = "%s"
			control_library_version = "%s"
			latest = %s
			controls {
				control_name = "SC-7"
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

		data "ibm_scc_control_library" "scc_control_library" {
			instance_id = ibm_scc_control_library.scc_control_library_instance.instance_id
			control_library_id = ibm_scc_control_library.scc_control_library_instance.control_library_id
		}

	`, instanceID, controlLibraryControlLibraryName, controlLibraryControlLibraryDescription, controlLibraryControlLibraryType, controlLibraryControlLibraryVersion, controlLibraryLatest)
}
