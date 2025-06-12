// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccControlLibraryBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ControlLibrary
	controlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryType := "custom"
	controlLibraryNameUpdate := controlLibraryName
	controlLibraryDescriptionUpdate := controlLibraryDescription
	controlLibraryTypeUpdate := controlLibraryType

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccControlLibraryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfigBasic(acc.SccInstanceID, controlLibraryName, controlLibraryDescription, controlLibraryType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccControlLibraryExists("ibm_scc_control_library.scc_control_library_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryName),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescription),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfigBasic(acc.SccInstanceID, controlLibraryNameUpdate, controlLibraryDescriptionUpdate, controlLibraryTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryTypeUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccControlLibraryAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ControlLibrary
	controlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryType := "custom"
	controlLibraryVersion := "0.0.0"
	latest := "true"
	controlsCount := "1"

	controlLibraryNameUpdate := controlLibraryName
	controlLibraryDescriptionUpdate := controlLibraryDescription
	controlLibraryTypeUpdate := "custom"
	controlLibraryVersionUpdate := "0.0.1"
	latestUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccControlLibraryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfigBasic(acc.SccInstanceID, controlLibraryName, controlLibraryDescription, controlLibraryType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccControlLibraryExists("ibm_scc_control_library.scc_control_library_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryName),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescription),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryType),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_version", controlLibraryVersion),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "latest", latest),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "controls_count", controlsCount),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfig(acc.SccInstanceID, controlLibraryNameUpdate, controlLibraryDescriptionUpdate, controlLibraryTypeUpdate, controlLibraryVersionUpdate, latestUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryTypeUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_version", controlLibraryVersionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "latest", latestUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "controls_count", controlsCount),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_control_library.scc_control_library_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccControlLibraryConfigBasic(instanceID string, controlLibraryName string, controlLibraryDescription string, controlLibraryType string) string {
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
						assessment_description = "test 1"
						parameters {
							parameter_display_name = "Sign out due to inactivity in seconds"
							parameter_name         = "session_invalidation_in_seconds"
							parameter_type         = "numeric"
						}
					}
					assessments {
						assessment_id = "rule-f88e215f-bb33-4bd8-bd1c-d8a065e9aa70"
						assessment_method = "ibm-cloud-rule"
						assessment_type = "automated"
						assessment_description = "test 2"
						parameters {
							parameter_display_name = "Maximum length of netmask bit that is considered as wide flow"
							parameter_name         = "netmask_bits_length"
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
	`, instanceID, controlLibraryName, controlLibraryDescription, controlLibraryType)
}

func testAccCheckIbmSccControlLibraryConfig(instanceID string, controlLibraryName string, controlLibraryDescription string, controlLibraryType string, controlLibraryVersion string, latest string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_control_library" "scc_control_library_instance" {
			instance_id = "%s"
			control_library_name = "%s"
			control_library_description = "%s"
			control_library_type = "%s"
			control_library_version = "%s"
			latest = %s
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
						assessment_id = "rule-f88e215f-bb33-4bd8-bd1c-d8a065e9aa70"
						assessment_method = "ibm-cloud-rule"
						assessment_type = "automated"
						assessment_description = "test 2"
						parameters {
							parameter_display_name = "Maximum length of netmask bit that is considered as wide flow"
							parameter_name         = "netmask_bits_length"
							parameter_type         = "numeric"
						}
					}
					assessments {
						assessment_id = "rule-a637949b-7e51-46c4-afd4-b96619001bf1"
						assessment_method = "ibm-cloud-rule"
						assessment_type = "automated"
						assessment_description = "test 1"
						parameters {
							parameter_display_name  = "Sign out due to inactivity in seconds"
							parameter_name          = "session_invalidation_in_seconds"
							parameter_type          = "numeric"
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
	`, instanceID, controlLibraryName, controlLibraryDescription, controlLibraryType, controlLibraryVersion, latest)
}

func testAccCheckIbmSccControlLibraryExists(n string, obj securityandcompliancecenterapiv3.ControlLibrary) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getControlLibraryOptions := &securityandcompliancecenterapiv3.GetControlLibraryOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		getControlLibraryOptions.SetInstanceID(id[0])
		getControlLibraryOptions.SetControlLibraryID(id[1])

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

		id := strings.Split(rs.Primary.ID, "/")
		getControlLibraryOptions.SetInstanceID(id[0])
		getControlLibraryOptions.SetControlLibraryID(id[1])

		// Try to find the key
		_, response, err := securityandcompliancecenterapiClient.GetControlLibrary(getControlLibraryOptions)

		if err == nil {
			return flex.FmtErrorf("scc_control_library still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("Error checking for scc_control_library (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
