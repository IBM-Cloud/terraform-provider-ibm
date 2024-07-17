// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccProfileAttachmentBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.AttachmentItem

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
	var conf securityandcompliancecenterapiv3.AttachmentItem

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfig(acc.SccInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileAttachmentExists("ibm_scc_profile_attachment.scc_profile_attachment_instance", conf),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfig(acc.SccInstanceID),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfigChange(acc.SccInstanceID),
				Check:  resource.ComposeAggregateTestCheckFunc(),
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
			version_group_label = "03354ab4-03be-41c0-a469-826fc0262e78"
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

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			instance_id = resource.ibm_scc_control_library.scc_control_library_instance.instance_id
			profile_id = ibm_scc_profile.scc_profile_instance.profile_id
			name = "profile_attachment_name"
			description = "scc_profile_attachment_description"
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
	`, instanceID)
}

func testAccCheckIbmSccProfileAttachmentConfig(instanceID string) string {
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

    resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
		instance_id = "%s"
		profile_id = local.scc_profiles_map["CIS IBM Cloud Foundations Benchmark v1.1.0"]
		name = "profile_attachment_name"
		description = "scc_profile_attachment_description"
			scope {
				environment = "ibm-cloud"	
				properties {
					name = "scope_id"
					value = data.ibm_iam_account_settings.iam_account_settings.account_id
				}
				properties {
					name = "scope_type"
					value = "account"
				}
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
	}
	`, instanceID, instanceID)
}

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

    data "ibm_iam_account_settings" "iam_account_settings" {
    }

    resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
		instance_id = "%s"
		profile_id = local.scc_profiles_map["CIS IBM Cloud Foundations Benchmark v1.1.0"]
		name = "profile_attachment_name"
		description = "scc_profile_attachment_description"
			scope {
				environment = "ibm-cloud"	
				properties {
					name = "scope_id"
					value = data.ibm_iam_account_settings.iam_account_settings.account_id
				}
				properties {
					name = "scope_type"
					value = "account"
				}
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
			parameter_value = "22"
			assessment_id = "rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"
			assessment_type = "automated"
			parameter_display_name = "Security group rule SSH allow port number"
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
			parameter_value = "3389"
			assessment_id = "rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"
			assessment_type = "automated"
			parameter_display_name = "Disallowed IPs for ingress to RDP port"
			parameter_name = "rdp_port"
			parameter_type = "numeric"
		}
	}
	`, instanceID, instanceID)
}

func testAccCheckIbmSccProfileAttachmentExists(n string, obj securityandcompliancecenterapiv3.AttachmentItem) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
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
			return fmt.Errorf("scc_profile_attachment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_profile_attachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
