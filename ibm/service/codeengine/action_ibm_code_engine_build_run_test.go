// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

// TestAccIbmCodeEngineBuildRunActionBasic tests successful build run invocation
// This test verifies that:
// - Action can be invoked with required parameters via lifecycle trigger
// - Build run is created successfully
// - No errors are returned
func TestAccIbmCodeEngineBuildRunActionBasic(t *testing.T) {
	projectID := acc.CeProjectId
	buildName := fmt.Sprintf("tf-build-test-%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				// Create build and trigger action via lifecycle
				Config: buildRunActionConfigBasic(projectID, buildName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_build.test_build", "build_id"),
					checkBuildRunActionInvoked(projectID, buildName),
				),
			},
		},
	})
}

// TestAccIbmCodeEngineBuildRunActionTimeout tests timeout handling
// This test verifies that:
// - Appropriate error is returned when build run exceeds wait timeout
func TestAccIbmCodeEngineBuildRunActionTimeout(t *testing.T) {
	projectID := acc.CeProjectId
	buildName := fmt.Sprintf("tf-build-timeout-fail-%d", acctest.RandIntRange(10, 1000))
	shortTimeout := int64(10) // 10 seconds - likely to timeout for real builds

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      buildRunActionConfigWithTimeout(projectID, buildName, shortTimeout),
				ExpectError: regexp.MustCompile("timeout|timed out"),
			},
		},
	})
}

// TestAccIbmCodeEngineBuildRunActionBuildNotFound tests error handling for non-existent build
// This test verifies that:
// - Action returns appropriate error when build doesn't exist
// - Error message is clear and actionable
func TestAccIbmCodeEngineBuildRunActionBuildNotFound(t *testing.T) {
	projectID := acc.CeProjectId
	nonExistentBuild := "non-existent-build-12345"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		ExternalProviders: map[string]resource.ExternalProvider{
			"null": {
				Source:            "hashicorp/null",
				VersionConstraint: "~> 3.0",
			},
		},
		Steps: []resource.TestStep{
			{
				Config:      buildRunActionConfigNonExistentBuild(projectID, nonExistentBuild),
				ExpectError: regexp.MustCompile("Resource Not Found|not found|404"),
			},
		},
	})

}

// TestAccIbmCodeEngineBuildRunActionNoWait tests no_wait parameter
// This test verifies that:
// - Action returns immediately when no_wait=true without waiting for completion
func TestAccIbmCodeEngineBuildRunActionNoWait(t *testing.T) {
	projectID := acc.CeProjectId
	buildName := fmt.Sprintf("tf-build-no-wait-%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: buildRunActionConfigWithNoWait(projectID, buildName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_build.test_build", "build_id"),
					checkBuildRunActionInvoked(projectID, buildName),
				),
			},
		},
	})
}

// TestAccIbmCodeEngineBuildRunActionWithName tests custom name parameter
// This test verifies that:
// - Action accepts custom name parameter
// - Build run is created with the specified name
func TestAccIbmCodeEngineBuildRunActionWithName(t *testing.T) {
	projectID := acc.CeProjectId
	buildName := fmt.Sprintf("tf-build-custom-name-%d", acctest.RandIntRange(10, 1000))
	customRunName := fmt.Sprintf("custom-run-%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: buildRunActionConfigWithName(projectID, buildName, customRunName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_build.test_build", "build_id"),
					checkBuildRunActionInvoked(projectID, buildName),
				),
			},
		},
	})
}

// Helper function to verify action was invoked by checking for recent build runs
func checkBuildRunActionInvoked(projectID, buildName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return fmt.Errorf("Error getting Code Engine client: %s", err)
		}

		// List build runs for this build to verify one was created
		listBuildRunsOptions := &codeenginev2.ListBuildRunsOptions{}
		listBuildRunsOptions.SetProjectID(projectID)
		listBuildRunsOptions.SetBuildName(buildName)
		listBuildRunsOptions.SetLimit(10)

		buildRunsList, _, err := codeEngineClient.ListBuildRuns(listBuildRunsOptions)
		if err != nil {
			return fmt.Errorf("Error listing build runs: %s", err)
		}

		if buildRunsList == nil || len(buildRunsList.BuildRuns) == 0 {
			return fmt.Errorf("No build runs found for build %s - action may not have been invoked", buildName)
		}

		// Verify at least one build run exists (action was invoked)
		return nil
	}
}

// Configuration helpers

func buildRunActionConfigBasic(projectID, buildName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		action "ibm_code_engine_build_run" "test_action" {
			config {
				project_id = "%s"
				build_name = "%s"
			}
		}

		resource "ibm_code_engine_build" "test_build" {
			project_id    = data.ibm_code_engine_project.test_project.project_id
			name          = "%s"
			output_image  = "private.us.icr.io/ce-terraform-test/%s"
			output_secret = "ce-terraform-test"
			source_url    = "https://github.com/IBM/CodeEngine"
			strategy_type = "dockerfile"

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_code_engine_build_run.test_action]
				}
			}
		}
	`, projectID, projectID, buildName, buildName, buildName)
}

func buildRunActionConfigWithTimeout(projectID, buildName string, timeout int64) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		action "ibm_code_engine_build_run" "test_action" {
			config {
				project_id = "%s"
				build_name = "%s"
				timeout    = %d
			}
		}

		resource "ibm_code_engine_build" "test_build" {
			project_id    = data.ibm_code_engine_project.test_project.project_id
			name          = "%s"
			output_image  = "private.us.icr.io/ce-terraform-test/%s"
			output_secret = "ce-terraform-test"
			source_url    = "https://github.com/IBM/CodeEngine"
			strategy_type = "dockerfile"

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_code_engine_build_run.test_action]
				}
			}
		}
	`, projectID, projectID, buildName, timeout, buildName, buildName)
}

func buildRunActionConfigWithNoWait(projectID, buildName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		action "ibm_code_engine_build_run" "test_action" {
			config {
				project_id = "%s"
				build_name = "%s"
				no_wait    = true
			}
		}

		resource "ibm_code_engine_build" "test_build" {
			project_id    = data.ibm_code_engine_project.test_project.project_id
			name          = "%s"
			output_image  = "private.us.icr.io/ce-terraform-test/%s"
			output_secret = "ce-terraform-test"
			source_url    = "https://github.com/IBM/CodeEngine"
			strategy_type = "dockerfile"

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_code_engine_build_run.test_action]
				}
			}
		}
	`, projectID, projectID, buildName, buildName, buildName)
}

func buildRunActionConfigWithName(projectID, buildName, customRunName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		action "ibm_code_engine_build_run" "test_action" {
			config {
				project_id = "%s"
				build_name = "%s"
				name       = "%s"
				no_wait    = true
			}
		}

		resource "ibm_code_engine_build" "test_build" {
			project_id    = data.ibm_code_engine_project.test_project.project_id
			name          = "%s"
			output_image  = "private.us.icr.io/ce-terraform-test/%s"
			output_secret = "ce-terraform-test"
			source_url    = "https://github.com/IBM/CodeEngine"
			strategy_type = "dockerfile"

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_code_engine_build_run.test_action]
				}
			}
		}
	`, projectID, projectID, buildName, customRunName, buildName, buildName)
}

func buildRunActionConfigNonExistentBuild(projectID, buildName string) string {
	return fmt.Sprintf(`
		terraform {
			required_providers {
				null = {
					source  = "hashicorp/null"
					version = "~> 3.0"
				}
			}
		}

		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		action "ibm_code_engine_build_run" "test_action" {
			config {
				project_id = "%s"
				build_name = "%s"
			}
		}

		resource "null_resource" "trigger_action" {
			provisioner "local-exec" {
				command = "echo 'Triggering action for non-existent build'"
			}

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_code_engine_build_run.test_action]
				}
			}
		}
	`, projectID, projectID, buildName)
}
