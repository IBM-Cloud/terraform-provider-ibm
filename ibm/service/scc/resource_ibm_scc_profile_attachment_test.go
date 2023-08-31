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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccProfileAttachmentBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.AttachmentItem

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileAttachmentExists("ibm_scc_profile_attachment.scc_profile_attachment", conf),
				),
			},
		},
	})
}

func TestAccIbmSccProfileAttachmentAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.AttachmentItem
	xCorrelationID := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestID := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	xCorrelationIDUpdate := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestIDUpdate := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProfileAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfig(xCorrelationID, xRequestID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProfileAttachmentExists("ibm_scc_profile_attachment.scc_profile_attachment", conf),
					resource.TestCheckResourceAttr("ibm_scc_profile_attachment.scc_profile_attachment", "x_correlation_id", xCorrelationID),
					resource.TestCheckResourceAttr("ibm_scc_profile_attachment.scc_profile_attachment", "x_request_id", xRequestID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProfileAttachmentConfig(xCorrelationIDUpdate, xRequestIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_profile_attachment.scc_profile_attachment", "x_correlation_id", xCorrelationIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_profile_attachment.scc_profile_attachment", "x_request_id", xRequestIDUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_profile_attachment.scc_profile_attachment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccProfileAttachmentConfigBasic() string {
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
	`)
}

func testAccCheckIbmSccProfileAttachmentConfig(xCorrelationID string, xRequestID string) string {
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
			profiles_id = "profiles_id"
		}

		resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
			profiles_id = ibm_scc_profile.scc_profile_instance.id
			x_correlation_id = "%s"
			x_request_id = "%s"
			profile_id = ibm_scc_profile.scc_profile_instance.id
		}
	`, xCorrelationID, xRequestID)
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

		getProfileAttachmentOptions.SetProfileID(parts[0])
		getProfileAttachmentOptions.SetAttachmentID(parts[1])

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

		getProfileAttachmentOptions.SetProfileID(parts[0])
		getProfileAttachmentOptions.SetAttachmentID(parts[1])

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
