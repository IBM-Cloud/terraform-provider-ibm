// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProfileAttachmentDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProfileAttachmentDataSourceConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "profile_id"),
				),
			},
		},
	})
}

func TestAccIbmSccProfileAttachmentDataSourceAllArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProfileAttachmentDataSourceConfig(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "scope.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "scope.0.environment"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "schedule"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "notifications.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "last_scan.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "next_scan_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_profile_attachment.scc_profile_attachment_instance", "description"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProfileAttachmentDataSourceConfigBasic(instanceID string) string {
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
			profile_name = "profile_name"
			profile_description = "profile_description"
			profile_type = "custom"
			controls {
				control_library_id = resource.ibm_scc_control_library.scc_control_library_instance.control_library_id
				control_id = resource.ibm_scc_control_library.scc_control_library_instance.controls[0].control_id
			}
			default_parameters {
			}
		}

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			profile_id = ibm_scc_profile.scc_profile_instance.profile_id
			name = "profile_attachment_name"
			description = "profile_attachment_description"
			scope {
				environment = "ibm-cloud"	
				properties {
					name = "scope_id"
					value = resource.ibm_scc_control_library.scc_control_library_instance.account_id
				}
				properties {
					name = "scope_type"
					value = "account"
				}
			}
			schedule = "every_30_days"
			status = "enabled"
			notifications {
				enabled = false
				controls {
					failed_control_ids = []
					threshold_limit = 14
				}
			}
		}

		data "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			attachment_id = ibm_scc_profile_attachment.scc_profile_attachment_instance.attachment_id
			profile_id = ibm_scc_profile_attachment.scc_profile_attachment_instance.profile_id
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
		}
	`, instanceID)
}

func testAccCheckIbmSccProfileAttachmentDataSourceConfig(instanceID string) string {
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
			profile_name = "profile_name"
			profile_description = "profile_description"
			profile_type = "custom"
			controls {
				control_library_id = resource.ibm_scc_control_library.scc_control_library_instance.control_library_id
				control_id = resource.ibm_scc_control_library.scc_control_library_instance.controls[0].control_id
			}
			default_parameters {
			}
		}

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			profile_id = ibm_scc_profile.scc_profile_instance.profile_id
			name = "profile_attachment_name"
			description = "profile_attachment_description"
			scope {
				environment = "ibm-cloud"	
				properties {
					name = "scope_id"
					value = resource.ibm_scc_control_library.scc_control_library_instance.account_id
				}
				properties {
					name = "scope_type"
					value = "account"
				}
			}
			schedule = "every_30_days"
			status = "enabled"
			notifications {
				enabled = false
				controls {
					failed_control_ids = []
					threshold_limit = 14
				}
			}
		}

		data "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			attachment_id = ibm_scc_profile_attachment.scc_profile_attachment_instance.attachment_id
			profile_id = ibm_scc_profile_attachment.scc_profile_attachment_instance.profile_id
		}
	`, instanceID)
}
