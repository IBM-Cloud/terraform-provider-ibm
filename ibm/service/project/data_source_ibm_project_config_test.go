// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmProjectConfigDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "needs_attention_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "outputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "references.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "deployment_model"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectConfigDataSourceConfigBasic() string {
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

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
            definition {
                name = "stage-environment"
                authorizations {
                    method = "api_key"
                    api_key = "%s"
                }
                locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global"
                inputs = {
                    app_repo_name = "grit-repo-name"
                }
            }
            lifecycle {
                ignore_changes = [
                    definition[0].authorizations[0].api_key,
                ]
            }
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project_config.project_config_instance.project_id
			project_config_id = ibm_project_config.project_config_instance.project_config_id
		}
	`, acc.ProjectsConfigApiKey)
}

func TestDataSourceIbmProjectConfigProjectConfigNeedsAttentionStateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["event_id"] = "testString"
		model["event"] = "testString"
		model["severity"] = "INFO"
		model["action_url"] = "testString"
		model["target"] = "testString"
		model["triggered_by"] = "testString"
		model["timestamp"] = "2019-01-01T12:00:00.000Z"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigNeedsAttentionState)
	model.EventID = core.StringPtr("testString")
	model.Event = core.StringPtr("testString")
	model.Severity = core.StringPtr("INFO")
	model.ActionURL = core.StringPtr("testString")
	model.Target = core.StringPtr("testString")
	model.TriggeredBy = core.StringPtr("testString")
	model.Timestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")

	result, err := project.DataSourceIbmProjectConfigProjectConfigNeedsAttentionStateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigOutputValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["description"] = "testString"
		model["value"] = "testString"
		model["sensitive"] = true

		assert.Equal(t, result, model)
	}

	model := new(projectv1.OutputValue)
	model.Name = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.Value = "testString"
	model.Sensitive = core.BoolPtr(true)

	result, err := project.DataSourceIbmProjectConfigOutputValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigReferenceValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ReferenceValue)

	result, err := project.DataSourceIbmProjectConfigReferenceValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigErrorToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectConfigErrorDetailsModel := make(map[string]interface{})

		model := make(map[string]interface{})
		model["message"] = "testString"
		model["details"] = []map[string]interface{}{projectConfigErrorDetailsModel}

		assert.Equal(t, result, model)
	}

	projectConfigErrorDetailsModel := new(projectv1.ProjectConfigErrorDetails)

	model := new(projectv1.ProjectConfigError)
	model.Message = core.StringPtr("testString")
	model.Details = projectConfigErrorDetailsModel

	result, err := project.DataSourceIbmProjectConfigProjectConfigErrorToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigErrorDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigErrorDetails)

	result, err := project.DataSourceIbmProjectConfigProjectConfigErrorDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectReferenceToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectConfigProjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectDefinitionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectDefinitionReference)
	model.Name = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigProjectDefinitionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigSchematicsMetadataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		scriptModel := make(map[string]interface{})
		scriptModel["type"] = "ansible"
		scriptModel["path"] = "scripts/validate-post-ansible-playbook.yaml"
		scriptModel["short_description"] = "testString"

		model := make(map[string]interface{})
		model["workspace_crn"] = "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		model["validate_pre_script"] = []map[string]interface{}{scriptModel}
		model["validate_post_script"] = []map[string]interface{}{scriptModel}
		model["deploy_pre_script"] = []map[string]interface{}{scriptModel}
		model["deploy_post_script"] = []map[string]interface{}{scriptModel}
		model["undeploy_pre_script"] = []map[string]interface{}{scriptModel}
		model["undeploy_post_script"] = []map[string]interface{}{scriptModel}

		assert.Equal(t, result, model)
	}

	scriptModel := new(projectv1.Script)
	scriptModel.Type = core.StringPtr("ansible")
	scriptModel.Path = core.StringPtr("scripts/validate-post-ansible-playbook.yaml")
	scriptModel.ShortDescription = core.StringPtr("testString")

	model := new(projectv1.SchematicsMetadata)
	model.WorkspaceCrn = core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
	model.ValidatePreScript = scriptModel
	model.ValidatePostScript = scriptModel
	model.DeployPreScript = scriptModel
	model.DeployPostScript = scriptModel
	model.UndeployPreScript = scriptModel
	model.UndeployPostScript = scriptModel

	result, err := project.DataSourceIbmProjectConfigSchematicsMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigScriptToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["type"] = "ansible"
		model["path"] = "scripts/validate-post-ansible-playbook.yaml"
		model["short_description"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.Script)
	model.Type = core.StringPtr("ansible")
	model.Path = core.StringPtr("scripts/validate-post-ansible-playbook.yaml")
	model.ShortDescription = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigScriptToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigMemberOfDefinitionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		stackMemberModel := make(map[string]interface{})
		stackMemberModel["name"] = "testString"
		stackMemberModel["config_id"] = "testString"

		stackConfigDefinitionSummaryModel := make(map[string]interface{})
		stackConfigDefinitionSummaryModel["name"] = "testString"
		stackConfigDefinitionSummaryModel["members"] = []map[string]interface{}{stackMemberModel}

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["definition"] = []map[string]interface{}{stackConfigDefinitionSummaryModel}
		model["version"] = int(0)
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	stackMemberModel := new(projectv1.StackMember)
	stackMemberModel.Name = core.StringPtr("testString")
	stackMemberModel.ConfigID = core.StringPtr("testString")

	stackConfigDefinitionSummaryModel := new(projectv1.StackConfigDefinitionSummary)
	stackConfigDefinitionSummaryModel.Name = core.StringPtr("testString")
	stackConfigDefinitionSummaryModel.Members = []projectv1.StackMember{*stackMemberModel}

	model := new(projectv1.MemberOfDefinition)
	model.ID = core.StringPtr("testString")
	model.Definition = stackConfigDefinitionSummaryModel
	model.Version = core.Int64Ptr(int64(0))
	model.Href = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigMemberOfDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigStackConfigDefinitionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		stackMemberModel := make(map[string]interface{})
		stackMemberModel["name"] = "testString"
		stackMemberModel["config_id"] = "testString"

		model := make(map[string]interface{})
		model["name"] = "testString"
		model["members"] = []map[string]interface{}{stackMemberModel}

		assert.Equal(t, result, model)
	}

	stackMemberModel := new(projectv1.StackMember)
	stackMemberModel.Name = core.StringPtr("testString")
	stackMemberModel.ConfigID = core.StringPtr("testString")

	model := new(projectv1.StackConfigDefinitionSummary)
	model.Name = core.StringPtr("testString")
	model.Members = []projectv1.StackMember{*stackMemberModel}

	result, err := project.DataSourceIbmProjectConfigStackConfigDefinitionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigStackMemberToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["config_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.StackMember)
	model.Name = core.StringPtr("testString")
	model.ConfigID = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigStackMemberToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigDefinitionResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectComplianceProfileModel := make(map[string]interface{})

		stackMemberModel := make(map[string]interface{})
		stackMemberModel["name"] = "testString"
		stackMemberModel["config_id"] = "testString"

		projectConfigUsesModel := make(map[string]interface{})
		projectConfigUsesModel["config_id"] = "testString"
		projectConfigUsesModel["project_id"] = "testString"

		projectConfigAuthModel := make(map[string]interface{})
		projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
		projectConfigAuthModel["method"] = "trusted_profile"
		projectConfigAuthModel["api_key"] = "testString"

		model := make(map[string]interface{})
		model["compliance_profile"] = []map[string]interface{}{projectComplianceProfileModel}
		model["locator_id"] = "testString"
		model["members"] = []map[string]interface{}{stackMemberModel}
		model["uses"] = []map[string]interface{}{projectConfigUsesModel}
		model["description"] = "testString"
		model["name"] = "testString"
		model["authorizations"] = []map[string]interface{}{projectConfigAuthModel}
		model["inputs"] = map[string]interface{}{"anyKey": "anyValue"}
		model["settings"] = map[string]interface{}{"anyKey": "anyValue"}
		model["environment_id"] = "testString"
		model["resource_crns"] = []string{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}

		assert.Equal(t, result, model)
	}

	projectComplianceProfileModel := new(projectv1.ProjectComplianceProfileNullableObject)

	stackMemberModel := new(projectv1.StackMember)
	stackMemberModel.Name = core.StringPtr("testString")
	stackMemberModel.ConfigID = core.StringPtr("testString")

	projectConfigUsesModel := new(projectv1.ProjectConfigUses)
	projectConfigUsesModel.ConfigID = core.StringPtr("testString")
	projectConfigUsesModel.ProjectID = core.StringPtr("testString")

	projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
	projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
	projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
	projectConfigAuthModel.ApiKey = core.StringPtr("testString")

	model := new(projectv1.ProjectConfigDefinitionResponse)
	model.ComplianceProfile = projectComplianceProfileModel
	model.LocatorID = core.StringPtr("testString")
	model.Members = []projectv1.StackMember{*stackMemberModel}
	model.Uses = []projectv1.ProjectConfigUses{*projectConfigUsesModel}
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Authorizations = projectConfigAuthModel
	model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
	model.Settings = map[string]interface{}{"anyKey": "anyValue"}
	model.EnvironmentID = core.StringPtr("testString")
	model.ResourceCrns = []string{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}

	result, err := project.DataSourceIbmProjectConfigProjectConfigDefinitionResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectComplianceProfileToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectConfigProjectComplianceProfileToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectComplianceProfileNullableObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectComplianceProfileNullableObject)

	result, err := project.DataSourceIbmProjectConfigProjectComplianceProfileNullableObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectComplianceProfileV1ToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectConfigProjectComplianceProfileV1ToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigUsesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["config_id"] = "testString"
		model["project_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigUses)
	model.ConfigID = core.StringPtr("testString")
	model.ProjectID = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigProjectConfigUsesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigAuthToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectConfigProjectConfigAuthToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectComplianceProfileModel := make(map[string]interface{})

		stackMemberModel := make(map[string]interface{})
		stackMemberModel["name"] = "testString"
		stackMemberModel["config_id"] = "testString"

		projectConfigUsesModel := make(map[string]interface{})
		projectConfigUsesModel["config_id"] = "testString"
		projectConfigUsesModel["project_id"] = "testString"

		projectConfigAuthModel := make(map[string]interface{})
		projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
		projectConfigAuthModel["method"] = "trusted_profile"
		projectConfigAuthModel["api_key"] = "testString"

		model := make(map[string]interface{})
		model["compliance_profile"] = []map[string]interface{}{projectComplianceProfileModel}
		model["locator_id"] = "testString"
		model["members"] = []map[string]interface{}{stackMemberModel}
		model["uses"] = []map[string]interface{}{projectConfigUsesModel}
		model["description"] = "testString"
		model["name"] = "testString"
		model["authorizations"] = []map[string]interface{}{projectConfigAuthModel}
		model["inputs"] = map[string]interface{}{"anyKey": "anyValue"}
		model["settings"] = map[string]interface{}{"anyKey": "anyValue"}
		model["environment_id"] = "testString"

		assert.Equal(t, result, model)
	}

	projectComplianceProfileModel := new(projectv1.ProjectComplianceProfileNullableObject)

	stackMemberModel := new(projectv1.StackMember)
	stackMemberModel.Name = core.StringPtr("testString")
	stackMemberModel.ConfigID = core.StringPtr("testString")

	projectConfigUsesModel := new(projectv1.ProjectConfigUses)
	projectConfigUsesModel.ConfigID = core.StringPtr("testString")
	projectConfigUsesModel.ProjectID = core.StringPtr("testString")

	projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
	projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
	projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
	projectConfigAuthModel.ApiKey = core.StringPtr("testString")

	model := new(projectv1.ProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponse)
	model.ComplianceProfile = projectComplianceProfileModel
	model.LocatorID = core.StringPtr("testString")
	model.Members = []projectv1.StackMember{*stackMemberModel}
	model.Uses = []projectv1.ProjectConfigUses{*projectConfigUsesModel}
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Authorizations = projectConfigAuthModel
	model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
	model.Settings = map[string]interface{}{"anyKey": "anyValue"}
	model.EnvironmentID = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectConfigAuthModel := make(map[string]interface{})
		projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
		projectConfigAuthModel["method"] = "trusted_profile"
		projectConfigAuthModel["api_key"] = "testString"

		model := make(map[string]interface{})
		model["resource_crns"] = []string{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
		model["description"] = "testString"
		model["name"] = "testString"
		model["authorizations"] = []map[string]interface{}{projectConfigAuthModel}
		model["inputs"] = map[string]interface{}{"anyKey": "anyValue"}
		model["settings"] = map[string]interface{}{"anyKey": "anyValue"}
		model["environment_id"] = "testString"

		assert.Equal(t, result, model)
	}

	projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
	projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
	projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
	projectConfigAuthModel.ApiKey = core.StringPtr("testString")

	model := new(projectv1.ProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponse)
	model.ResourceCrns = []string{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Authorizations = projectConfigAuthModel
	model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
	model.Settings = map[string]interface{}{"anyKey": "anyValue"}
	model.EnvironmentID = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigVersionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectConfigVersionDefinitionSummaryModel := make(map[string]interface{})
		projectConfigVersionDefinitionSummaryModel["environment_id"] = "testString"
		projectConfigVersionDefinitionSummaryModel["locator_id"] = "testString"

		model := make(map[string]interface{})
		model["definition"] = []map[string]interface{}{projectConfigVersionDefinitionSummaryModel}
		model["container_state"] = "approved"
		model["state"] = "approved"
		model["version"] = int(0)
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	projectConfigVersionDefinitionSummaryModel := new(projectv1.ProjectConfigVersionDefinitionSummary)
	projectConfigVersionDefinitionSummaryModel.EnvironmentID = core.StringPtr("testString")
	projectConfigVersionDefinitionSummaryModel.LocatorID = core.StringPtr("testString")

	model := new(projectv1.ProjectConfigVersionSummary)
	model.Definition = projectConfigVersionDefinitionSummaryModel
	model.ContainerState = core.StringPtr("approved")
	model.State = core.StringPtr("approved")
	model.Version = core.Int64Ptr(int64(0))
	model.Href = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigProjectConfigVersionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectConfigProjectConfigVersionDefinitionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["environment_id"] = "testString"
		model["locator_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigVersionDefinitionSummary)
	model.EnvironmentID = core.StringPtr("testString")
	model.LocatorID = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectConfigProjectConfigVersionDefinitionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
