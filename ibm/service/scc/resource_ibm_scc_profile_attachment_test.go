// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccProfileAttachmentBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ProfileAttachment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileAttachmentExists("ibm_scc_profile_attachment.scc_profile_attachment_instance", conf),
				),
			},
		},
	})
}

func TestAccIbmSccProfileAttachmentAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ProfileAttachment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfig(acc.SccInstanceID, acc.SccResourceGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileAttachmentExists("ibm_scc_profile_attachment.scc_profile_attachment_instance", conf),
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.#", "6"),
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "scope.0.properties.#", "3"),
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "scope.0.properties.2.name", "exclusions"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfigChange(acc.SccInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileAttachmentExists("ibm_scc_profile_attachment.scc_profile_attachment_instance", conf),
					// verify if all attachment_parameters are stored in the state
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.#", "6"),
					// verify the changes to rule-7c5f6385-67e4-4edf-bec8-c722558b2dec
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.0.assessment_id", "rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"),
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.0.parameter_value", "23"),
					// verify the changes to rule-9653d2c7-6290-4128-a5a3-65487ba40370
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.2.assessment_id", "rule-9653d2c7-6290-4128-a5a3-65487ba40370"),
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.2.parameter_value", "1234"),
					// verify the changes to rule-e16fcfea-fe21-4d30-a721-423611481fea
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.3.assessment_id", "rule-e16fcfea-fe21-4d30-a721-423611481fea"),
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.3.parameter_value", "['1.2', '1.3']"),
					// verify the changes to rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.4.assessment_id", "rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"),
					resource.TestCheckResourceAttr(
						"ibm_scc_profile_attachment.scc_profile_attachment_instance", "attachment_parameters.4.parameter_value", "4000"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_profile_attachment.scc_profile_attachment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccProfileAttachmentConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "control_library_name"
			control_library_description = "control_library_description"
			control_library_type = "custom"
			latest = true
			controls {
				control_id = "0d4624f5-f5f6-44ed-9e09-6662e2f4106c"
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
							parameter_type         = "numeric"
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

		resource "ibm_scc_scope" "scc_scope_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			name = "Scope-Terraform-Test"
			description = "Scope Made by Terraform Testing"
			environment = "ibm-cloud"
			properties  = {
				scope_id    = resource.ibm_scc_control_library.scc_control_library_instance.account_id
				scope_type  = "account"
			}
		}

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			profile_id = ibm_scc_profile.scc_profile_instance.profile_id
			name = "profile_attachment_name"
			description = "scc_profile_attachment_description"
			scope {
				id = ibm_scc_scope.scc_scope_instance.scope_id
			}
			schedule = "every_30_days"
			status = "disabled"
			notifications {
				enabled = false
				controls {
					failed_control_ids = []
				 	threshold_limit = 14
				}
			}
		}
  `, instanceID)
}

func testAccCheckIbmSccProfileAttachmentConfig(instanceID string, resGroupID string) string {
	return fmt.Sprintf(`
		locals {
			scc_profiles_map = tomap(merge([
				for i , cl in data.ibm_scc_profiles.scc_profiles.profiles :
				{(cl.profile_name) = "${cl.id}"}  if cl.latest == true && cl.profile_type == "predefined"
			]...))
		}

		data "ibm_scc_profiles" "scc_profiles" {
			instance_id = "%s"
		}

		data "ibm_iam_account_settings" "iam_account_settings" {
		}

		resource "ibm_scc_scope" "scc_scope_instance" {
			instance_id = data.ibm_scc_profiles.scc_profiles.instance_id
			name = "Scope-Terraform-Test"
			description = "Scope Made by Terraform Testing"
			environment = "ibm-cloud"
			properties  = {
				scope_id    = data.ibm_iam_account_settings.iam_account_settings.account_id
				scope_type  = "account"
			}
			exclusions {
				scope_id   = "%s"
				scope_type = "account.resource_group"
			}
		}

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			instance_id = data.ibm_scc_profiles.scc_profiles.instance_id
			profile_id = local.scc_profiles_map["CIS IBM Cloud Foundations Benchmark v1.1.0"]
			name = "terraform_ac_profile_attachment_name"
			description = "scc_profile_attachment_description"
			scope {
				id = ibm_scc_scope.scc_scope_instance.scope_id
			}
			schedule = "every_30_days"
			status = "disabled"
			notifications {
				enabled = false
				controls {
					failed_control_ids = []
					threshold_limit = 14
				}
			}
			attachment_parameters {
				parameter_value = "['1.2', '1.3']"
				assessment_id = "rule-e16fcfea-fe21-4d30-a721-423611481fea"
				assessment_type = "automated"
				parameter_display_name = "IBM Cloud Internet Services TLS version"
				parameter_name = "tls_version"
				parameter_type = "string_list"
			}
			attachment_parameters {
				parameter_value = "22"
				assessment_id = "rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"
				assessment_type = "automated"
				parameter_display_name = "Network ACL rule for allowed IPs to SSH port"
				parameter_name = "ssh_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "3389"
				assessment_id = "rule-9653d2c7-6290-4128-a5a3-65487ba40370"
				assessment_type = "automated"
				parameter_display_name = "Security group rule RDP allow port number"
				parameter_name = "rdp_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "22"
				assessment_id = "rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"
				assessment_type = "automated"
				parameter_display_name = "Security group rule SSH allow port number"
				parameter_name = "ssh_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "3389"
				assessment_id = "rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"
				assessment_type = "automated"
				parameter_display_name = "Disallowed IPs for ingress to RDP port"
				parameter_name = "rdp_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "['default']"
				assessment_id = "rule-96527f89-1867-4581-b923-1400e04661e0"
				assessment_type = "automated"
				parameter_display_name = "Exclude the default security groups"
				parameter_name = "exclude_default_security_groups"
				parameter_type = "string_list"
			}
  }
  `, instanceID, resGroupID)
}

// Returns a terraform change where the attachment_parameters are modified slightly.
func testAccCheckIbmSccProfileAttachmentConfigChange(instanceID string) string {
	return fmt.Sprintf(`
		locals {
			scc_profiles_map = tomap(merge([
			for i , cl in data.ibm_scc_profiles.scc_profiles.profiles :
				{(cl.profile_name) = "${cl.id}"}  if cl.latest == true && cl.profile_type == "predefined"
			]...))
		}

		data "ibm_scc_profiles" "scc_profiles" {
			instance_id = "%s"
		}

		data "ibm_iam_account_settings" "iam_account_settings" {}

		resource "ibm_scc_scope" "scc_scope_instance" {
			instance_id = data.ibm_scc_profiles.scc_profiles.instance_id
			name = "Scope-Terraform-Test"
			description = "Scope Made by Terraform Testing"
			environment = "ibm-cloud"
			properties  = {
				scope_id    = data.ibm_iam_account_settings.iam_account_settings.account_id
				scope_type  = "account"
			}
		}
		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			instance_id = "%s"
			profile_id = local.scc_profiles_map["CIS IBM Cloud Foundations Benchmark v1.1.0"]
			name = "profile_attachment_name"
			description = "scc_profile_attachment_description"
			scope {
				id = ibm_scc_scope.scc_scope_instance.scope_id
			}
			schedule = "every_30_days"
			status = "disabled"
			notifications {
				enabled = false
				controls {
					failed_control_ids = []
					threshold_limit = 14
				}
			}
			attachment_parameters {
				parameter_value = "['1.2', '1.3']"
				assessment_id = "rule-e16fcfea-fe21-4d30-a721-423611481fea"
				assessment_type = "automated"
				parameter_display_name = "IBM Cloud Internet Services TLS version"
				parameter_name = "tls_version"
				parameter_type = "string_list"
			}
			attachment_parameters {
				parameter_value = "8080"
				assessment_id = "rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"
				assessment_type = "automated"
				parameter_display_name = "Network ACL rule for allowed IPs to SSH port"
				parameter_name = "ssh_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "23"
				assessment_id = "rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"
				assessment_type = "automated"
				parameter_display_name = "Security group rule SSH allow port number"
				parameter_name = "ssh_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "4000"
				assessment_id = "rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"
				assessment_type = "automated"
				parameter_display_name = "Disallowed IPs for ingress to RDP port"
				parameter_name = "rdp_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "1234"
				assessment_id = "rule-9653d2c7-6290-4128-a5a3-65487ba40370"
				assessment_type = "automated"
				parameter_display_name = "Security group rule RDP allow port number"
				parameter_name = "rdp_port"
				parameter_type = "numeric"
			}
			attachment_parameters {
				parameter_value = "['default']"
				assessment_id = "rule-96527f89-1867-4581-b923-1400e04661e0"
				assessment_type = "automated"
				parameter_display_name = "Exclude the default security groups"
				parameter_name = "exclude_default_security_groups"
				parameter_type = "string_list"
			}
	}
`, instanceID, instanceID)
}

func testAccCheckIbmSccProfileAttachmentExists(n string, obj securityandcompliancecenterapiv3.ProfileAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getProfileAttachmentOptions := &securityandcompliancecenterapiv3.GetProfileAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProfileAttachmentOptions.SetInstanceID(parts[0])
		getProfileAttachmentOptions.SetProfileID(parts[1])
		getProfileAttachmentOptions.SetAttachmentID(parts[2])

		attachmentItem, _, err := securityandcompliancecenterapiClient.GetProfileAttachment(getProfileAttachmentOptions)
		if err != nil {
			return err
		}

		obj = *attachmentItem
		return nil
	}
}

func testAccCheckIbmSccProfileAttachmentDestroy(s *terraform.State) error {
	securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_profile_attachment" {
			continue
		}

		getProfileAttachmentOptions := &securityandcompliancecenterapiv3.GetProfileAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProfileAttachmentOptions.SetInstanceID(parts[0])
		getProfileAttachmentOptions.SetProfileID(parts[1])
		getProfileAttachmentOptions.SetAttachmentID(parts[2])

		// Try to find the key
		_, response, err := securityandcompliancecenterapiClient.GetProfileAttachment(getProfileAttachmentOptions)

		if err == nil {
			return flex.FmtErrorf("scc_profile_attachment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("Error checking for scc_profile_attachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
