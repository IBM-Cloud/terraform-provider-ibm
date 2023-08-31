// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProfileAttachmentDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProfileAttachmentDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "profiles_id"),
				),
			},
		},
	})
}

func TestAccIbmSccProfileAttachmentDataSourceAllArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProfileAttachmentDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "profiles_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_item_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "scope.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "scope.0.environment"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "schedule"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "notifications.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_parameters.0.assessment_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_parameters.0.assessment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_parameters.0.parameter_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_parameters.0.parameter_value"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_parameters.0.parameter_display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "attachment_parameters.0.parameter_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "last_scan.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "next_scan_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment", "description"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProfileAttachmentDataSourceConfigBasic() string {
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

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			profiles_id = ibm_scc_profile.scc_profile_instance.id
		}

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			profiles_id = ibm_scc_profile.scc_profile_instance.id
		}

		data "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			attachment_id = "attachment_id"
			profiles_id = ibm_scc_profile_attachment.scc_profile_attachment.profiles_id
		}
	`)
}

func testAccCheckIbmSccProfileAttachmentDataSourceConfig() string {
	return fmt.Sprint(`
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

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			profiles_id = ibm_scc_profile.scc_profile_instance.id
		}

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			profiles_id = ibm_scc_profile.scc_profile_instance.id
			profile_id = ibm_scc_profile.scc_profile_instance.id
		}

		data "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			attachment_id = "attachment_id"
			profiles_id = ibm_scc_profile_attachment.scc_profile_attachment.profiles_id
		}
	`)
}
