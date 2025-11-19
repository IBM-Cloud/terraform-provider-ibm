// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmProjectEnvironmentDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEnvironmentDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "project_environment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "project.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectEnvironmentDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
            location = "us-south"
            resource_group = "Default"
            definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
                monitoring_enabled = true
                auto_deploy = false
                auto_deploy_mode = "manual_approval"
            }
        }

        resource "ibm_project_environment" "project_environment_instance" {
            project_id = ibm_project.project_instance.id
            definition {
                name = "environment-stage"
                description = "environment for stage project"
                authorizations {
                    method = "api_key"
                    api_key = "%s"
               }
            }
            lifecycle {
                ignore_changes = [
                    definition[0].authorizations[0].api_key,
                ]
            }
        }

		data "ibm_project_environment" "project_environment_instance" {
			project_id = ibm_project_environment.project_environment_instance.project_id
			project_environment_id = ibm_project_environment.project_environment_instance.project_environment_id
		}
	`, acc.ProjectsConfigApiKey)
}

func TestDataSourceIbmProjectEnvironmentProjectReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectDefinitionReferenceModel := make(map[string]interface{})
		projectDefinitionReferenceModel["name"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["href"] = "testString"
		model["definition"] = []map[string]interface{}{projectDefinitionReferenceModel}
		model["crn"] = "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"

		assert.Equal(t, result, model)
	}

	projectDefinitionReferenceModel := new(projectv1.ProjectDefinitionReference)
	projectDefinitionReferenceModel.Name = core.StringPtr("testString")

	model := new(projectv1.ProjectReference)
	model.ID = core.StringPtr("testString")
	model.Href = core.StringPtr("testString")
	model.Definition = projectDefinitionReferenceModel
	model.Crn = core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")

	result, err := project.DataSourceIbmProjectEnvironmentProjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectEnvironmentProjectDefinitionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectDefinitionReference)
	model.Name = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectEnvironmentProjectDefinitionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectEnvironmentEnvironmentDefinitionRequiredPropertiesResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectConfigAuthModel := make(map[string]interface{})
		projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
		projectConfigAuthModel["method"] = "trusted_profile"
		projectConfigAuthModel["api_key"] = "testString"

		projectComplianceProfileModel := make(map[string]interface{})

		model := make(map[string]interface{})
		model["description"] = "testString"
		model["name"] = "testString"
		model["authorizations"] = []map[string]interface{}{projectConfigAuthModel}
		model["inputs"] = map[string]interface{}{"anyKey": "anyValue"}
		model["compliance_profile"] = []map[string]interface{}{projectComplianceProfileModel}

		assert.Equal(t, result, model)
	}

	projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
	projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
	projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
	projectConfigAuthModel.ApiKey = core.StringPtr("testString")

	projectComplianceProfileModel := new(projectv1.ProjectComplianceProfileNullableObject)

	model := new(projectv1.EnvironmentDefinitionRequiredPropertiesResponse)
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Authorizations = projectConfigAuthModel
	model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
	model.ComplianceProfile = projectComplianceProfileModel

	result, err := project.DataSourceIbmProjectEnvironmentEnvironmentDefinitionRequiredPropertiesResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectEnvironmentProjectConfigAuthToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
		model["method"] = "trusted_profile"
		model["api_key"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigAuth)
	model.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
	model.Method = core.StringPtr("trusted_profile")
	model.ApiKey = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectEnvironmentProjectConfigAuthToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectEnvironmentProjectComplianceProfileToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["instance_id"] = "testString"
		model["instance_location"] = "us-south"
		model["attachment_id"] = "someattachmentid"
		model["profile_name"] = "SCCProfilev1.0"
		model["wp_policy_id"] = "testString"
		model["wp_instance_id"] = "testString"
		model["wp_instance_name"] = "testString"
		model["wp_instance_location"] = "us-south"
		model["wp_zone_id"] = "testString"
		model["wp_zone_name"] = "testString"
		model["wp_policy_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectComplianceProfile)
	model.ID = core.StringPtr("testString")
	model.InstanceID = core.StringPtr("testString")
	model.InstanceLocation = core.StringPtr("us-south")
	model.AttachmentID = core.StringPtr("someattachmentid")
	model.ProfileName = core.StringPtr("SCCProfilev1.0")
	model.WpPolicyID = core.StringPtr("testString")
	model.WpInstanceID = core.StringPtr("testString")
	model.WpInstanceName = core.StringPtr("testString")
	model.WpInstanceLocation = core.StringPtr("us-south")
	model.WpZoneID = core.StringPtr("testString")
	model.WpZoneName = core.StringPtr("testString")
	model.WpPolicyName = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectEnvironmentProjectComplianceProfileToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectEnvironmentProjectComplianceProfileNullableObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectComplianceProfileNullableObject)

	result, err := project.DataSourceIbmProjectEnvironmentProjectComplianceProfileNullableObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectEnvironmentProjectComplianceProfileV1ToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["instance_id"] = "testString"
		model["instance_location"] = "us-south"
		model["attachment_id"] = "testString"
		model["profile_name"] = "testString"
		model["wp_policy_id"] = "testString"
		model["wp_instance_id"] = "testString"
		model["wp_instance_name"] = "testString"
		model["wp_instance_location"] = "us-south"
		model["wp_zone_id"] = "testString"
		model["wp_zone_name"] = "testString"
		model["wp_policy_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectComplianceProfileV1)
	model.ID = core.StringPtr("testString")
	model.InstanceID = core.StringPtr("testString")
	model.InstanceLocation = core.StringPtr("us-south")
	model.AttachmentID = core.StringPtr("testString")
	model.ProfileName = core.StringPtr("testString")
	model.WpPolicyID = core.StringPtr("testString")
	model.WpInstanceID = core.StringPtr("testString")
	model.WpInstanceName = core.StringPtr("testString")
	model.WpInstanceLocation = core.StringPtr("us-south")
	model.WpZoneID = core.StringPtr("testString")
	model.WpZoneName = core.StringPtr("testString")
	model.WpPolicyName = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectEnvironmentProjectComplianceProfileV1ToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
