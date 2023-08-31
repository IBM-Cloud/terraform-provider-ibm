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

func TestAccIbmSccControlLibraryBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ControlLibrary
	controlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryType := "predefined"
	controlLibraryNameUpdate := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescriptionUpdate := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryTypeUpdate := "custom"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccControlLibraryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfigBasic(controlLibraryName, controlLibraryDescription, controlLibraryType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccControlLibraryExists("ibm_scc_control_library.scc_control_library", conf),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_name", controlLibraryName),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_description", controlLibraryDescription),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_type", controlLibraryType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfigBasic(controlLibraryNameUpdate, controlLibraryDescriptionUpdate, controlLibraryTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_name", controlLibraryNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_description", controlLibraryDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_type", controlLibraryTypeUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccControlLibraryAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ControlLibrary
	xCorrelationID := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestID := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	controlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryType := "predefined"
	versionGroupLabel := fmt.Sprintf("tf_version_group_label_%d", acctest.RandIntRange(10, 100))
	controlLibraryVersion := fmt.Sprintf("tf_control_library_version_%d", acctest.RandIntRange(10, 100))
	latest := "true"
	controlsCount := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	xCorrelationIDUpdate := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestIDUpdate := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	controlLibraryNameUpdate := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescriptionUpdate := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryTypeUpdate := "custom"
	versionGroupLabelUpdate := fmt.Sprintf("tf_version_group_label_%d", acctest.RandIntRange(10, 100))
	controlLibraryVersionUpdate := fmt.Sprintf("tf_control_library_version_%d", acctest.RandIntRange(10, 100))
	latestUpdate := "false"
	controlsCountUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccControlLibraryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfig(xCorrelationID, xRequestID, controlLibraryName, controlLibraryDescription, controlLibraryType, versionGroupLabel, controlLibraryVersion, latest, controlsCount),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccControlLibraryExists("ibm_scc_control_library.scc_control_library", conf),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "x_correlation_id", xCorrelationID),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "x_request_id", xRequestID),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_name", controlLibraryName),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_description", controlLibraryDescription),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_type", controlLibraryType),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "version_group_label", versionGroupLabel),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_version", controlLibraryVersion),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "latest", latest),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "controls_count", controlsCount),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfig(xCorrelationIDUpdate, xRequestIDUpdate, controlLibraryNameUpdate, controlLibraryDescriptionUpdate, controlLibraryTypeUpdate, versionGroupLabelUpdate, controlLibraryVersionUpdate, latestUpdate, controlsCountUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "x_correlation_id", xCorrelationIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "x_request_id", xRequestIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_name", controlLibraryNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_description", controlLibraryDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_type", controlLibraryTypeUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "version_group_label", versionGroupLabelUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "control_library_version", controlLibraryVersionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "latest", latestUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library", "controls_count", controlsCountUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_control_library.scc_control_library",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccControlLibraryConfigBasic(controlLibraryName string, controlLibraryDescription string, controlLibraryType string) string {
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
	`, controlLibraryName, controlLibraryDescription, controlLibraryType)
}

func testAccCheckIbmSccControlLibraryConfig(xCorrelationID string, xRequestID string, controlLibraryName string, controlLibraryDescription string, controlLibraryType string, versionGroupLabel string, controlLibraryVersion string, latest string, controlsCount string) string {
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
			x_correlation_id = "%s"
			x_request_id = "%s"
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
	`, xCorrelationID, xRequestID, controlLibraryName, controlLibraryDescription, controlLibraryType, versionGroupLabel, controlLibraryVersion, latest, controlsCount)
}

func testAccCheckIbmSccControlLibraryExists(n string, obj securityandcompliancecenterapiv3.ControlLibrary) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getControlLibraryOptions := &securityandcompliancecenterapiv3.GetControlLibraryOptions{}

		getControlLibraryOptions.SetControlLibrariesID(rs.Primary.ID)

		controlLibrary, _, err := securityandcompliancecenterapiClient.GetControlLibrary(getControlLibraryOptions)
		if err != nil {
			return err
		}

		obj = *controlLibrary
		return nil
	}
}

func testAccCheckIbmSccControlLibraryDestroy(s *terraform.State) error {
	securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_control_library" {
			continue
		}

		getControlLibraryOptions := &securityandcompliancecenterapiv3.GetControlLibraryOptions{}

		getControlLibraryOptions.SetControlLibrariesID(rs.Primary.ID)

		// Try to find the key
		_, response, err := securityandcompliancecenterapiClient.GetControlLibrary(getControlLibraryOptions)

		if err == nil {
			return fmt.Errorf("scc_control_library still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_control_library (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
