// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmProjectConfigBasic(t *testing.T) {
	var conf projectv1.ProjectConfig

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectConfigDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectConfigExists("ibm_project_config.project_config_instance", conf),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_project_config.project_config_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func testAccCheckIbmProjectConfigConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			location = "ca-tor"
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
	`, acc.ProjectsConfigApiKey)
}

func testAccCheckIbmProjectConfigExists(n string, obj projectv1.ProjectConfig) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getConfigOptions := &projectv1.GetConfigOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigOptions.SetProjectID(parts[0])
		getConfigOptions.SetID(parts[1])

		projectConfig, _, err := projectClient.GetConfig(getConfigOptions)
		if err != nil {
			return err
		}

		obj = *projectConfig
		return nil
	}
}

func testAccCheckIbmProjectConfigDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project_config" {
			continue
		}

		getConfigOptions := &projectv1.GetConfigOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigOptions.SetProjectID(parts[0])
		getConfigOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := projectClient.GetConfig(getConfigOptions)

		if err == nil {
			return fmt.Errorf("project_config still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project_config (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmProjectConfigSchematicsMetadataToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigSchematicsMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigScriptToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigScriptToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectComplianceProfileToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigProjectComplianceProfileToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectComplianceProfileNullableObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectComplianceProfileNullableObject)

	result, err := project.ResourceIbmProjectConfigProjectComplianceProfileNullableObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectComplianceProfileV1ToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigProjectComplianceProfileV1ToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigStackMemberToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["config_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.StackMember)
	model.Name = core.StringPtr("testString")
	model.ConfigID = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectConfigStackMemberToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigAuthToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigProjectConfigAuthToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectComplianceProfileModel := make(map[string]interface{})

		stackMemberModel := make(map[string]interface{})
		stackMemberModel["name"] = "testString"
		stackMemberModel["config_id"] = "testString"

		projectConfigAuthModel := make(map[string]interface{})
		projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
		projectConfigAuthModel["method"] = "trusted_profile"
		projectConfigAuthModel["api_key"] = "testString"

		model := make(map[string]interface{})
		model["compliance_profile"] = []map[string]interface{}{projectComplianceProfileModel}
		model["locator_id"] = "testString"
		model["members"] = []map[string]interface{}{stackMemberModel}
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

	projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
	projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
	projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
	projectConfigAuthModel.ApiKey = core.StringPtr("testString")

	model := new(projectv1.ProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponse)
	model.ComplianceProfile = projectComplianceProfileModel
	model.LocatorID = core.StringPtr("testString")
	model.Members = []projectv1.StackMember{*stackMemberModel}
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Authorizations = projectConfigAuthModel
	model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
	model.Settings = map[string]interface{}{"anyKey": "anyValue"}
	model.EnvironmentID = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectConfigProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponseToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigNeedsAttentionStateToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigProjectConfigNeedsAttentionStateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigOutputValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["description"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.OutputValue)
	model.Name = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.Value = "testString"

	result, err := project.ResourceIbmProjectConfigOutputValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigReferenceValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ReferenceValue)

	result, err := project.ResourceIbmProjectConfigReferenceValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigErrorToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigProjectConfigErrorToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigErrorDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigErrorDetails)

	result, err := project.ResourceIbmProjectConfigProjectConfigErrorDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectReferenceToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigProjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectDefinitionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectDefinitionReference)
	model.Name = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectConfigProjectDefinitionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMemberOfDefinitionToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigMemberOfDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigStackConfigDefinitionSummaryToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectConfigStackConfigDefinitionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigVersionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectConfigVersionDefinitionSummaryModel := make(map[string]interface{})
		projectConfigVersionDefinitionSummaryModel["environment_id"] = "testString"
		projectConfigVersionDefinitionSummaryModel["locator_id"] = "testString"

		model := make(map[string]interface{})
		model["definition"] = []map[string]interface{}{projectConfigVersionDefinitionSummaryModel}
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
	model.State = core.StringPtr("approved")
	model.Version = core.Int64Ptr(int64(0))
	model.Href = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectConfigProjectConfigVersionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigProjectConfigVersionDefinitionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["environment_id"] = "testString"
		model["locator_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigVersionDefinitionSummary)
	model.EnvironmentID = core.StringPtr("testString")
	model.LocatorID = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectConfigProjectConfigVersionDefinitionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToProjectComplianceProfile(t *testing.T) {
	checkResult := func(result projectv1.ProjectComplianceProfileIntf) {
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

		assert.Equal(t, result, model)
	}

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

	result, err := project.ResourceIbmProjectConfigMapToProjectComplianceProfile(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToProjectComplianceProfileNullableObject(t *testing.T) {
	checkResult := func(result *projectv1.ProjectComplianceProfileNullableObject) {
		model := new(projectv1.ProjectComplianceProfileNullableObject)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})

	result, err := project.ResourceIbmProjectConfigMapToProjectComplianceProfileNullableObject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToProjectComplianceProfileV1(t *testing.T) {
	checkResult := func(result *projectv1.ProjectComplianceProfileV1) {
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

		assert.Equal(t, result, model)
	}

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

	result, err := project.ResourceIbmProjectConfigMapToProjectComplianceProfileV1(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToStackMember(t *testing.T) {
	checkResult := func(result *projectv1.StackMember) {
		model := new(projectv1.StackMember)
		model.Name = core.StringPtr("testString")
		model.ConfigID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["config_id"] = "testString"

	result, err := project.ResourceIbmProjectConfigMapToStackMember(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToProjectConfigAuth(t *testing.T) {
	checkResult := func(result *projectv1.ProjectConfigAuth) {
		model := new(projectv1.ProjectConfigAuth)
		model.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
		model.Method = core.StringPtr("trusted_profile")
		model.ApiKey = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
	model["method"] = "trusted_profile"
	model["api_key"] = "testString"

	result, err := project.ResourceIbmProjectConfigMapToProjectConfigAuth(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype(t *testing.T) {
	checkResult := func(result *projectv1.ProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype) {
		projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
		projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
		projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
		projectConfigAuthModel.ApiKey = core.StringPtr("testString")

		model := new(projectv1.ProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype)
		model.ResourceCrns = []string{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
		model.Description = core.StringPtr("testString")
		model.Name = core.StringPtr("testString")
		model.Authorizations = projectConfigAuthModel
		model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
		model.Settings = map[string]interface{}{"anyKey": "anyValue"}
		model.EnvironmentID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	projectConfigAuthModel := make(map[string]interface{})
	projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
	projectConfigAuthModel["method"] = "trusted_profile"
	projectConfigAuthModel["api_key"] = "testString"

	model := make(map[string]interface{})
	model["resource_crns"] = []interface{}{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
	model["description"] = "testString"
	model["name"] = "testString"
	model["authorizations"] = []interface{}{projectConfigAuthModel}
	model["inputs"] = map[string]interface{}{"anyKey": "anyValue"}
	model["settings"] = map[string]interface{}{"anyKey": "anyValue"}
	model["environment_id"] = "testString"

	result, err := project.ResourceIbmProjectConfigMapToProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToSchematicsWorkspace(t *testing.T) {
	checkResult := func(result *projectv1.SchematicsWorkspace) {
		model := new(projectv1.SchematicsWorkspace)
		model.WorkspaceCrn = core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["workspace_crn"] = "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"

	result, err := project.ResourceIbmProjectConfigMapToSchematicsWorkspace(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectConfigMapToProjectConfigDefinitionPatchResourceConfigDefinitionPropertiesPatch(t *testing.T) {
	checkResult := func(result *projectv1.ProjectConfigDefinitionPatchResourceConfigDefinitionPropertiesPatch) {
		projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
		projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
		projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
		projectConfigAuthModel.ApiKey = core.StringPtr("testString")

		model := new(projectv1.ProjectConfigDefinitionPatchResourceConfigDefinitionPropertiesPatch)
		model.ResourceCrns = []string{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
		model.Description = core.StringPtr("testString")
		model.Name = core.StringPtr("testString")
		model.Authorizations = projectConfigAuthModel
		model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
		model.Settings = map[string]interface{}{"anyKey": "anyValue"}
		model.EnvironmentID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	projectConfigAuthModel := make(map[string]interface{})
	projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
	projectConfigAuthModel["method"] = "trusted_profile"
	projectConfigAuthModel["api_key"] = "testString"

	model := make(map[string]interface{})
	model["resource_crns"] = []interface{}{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
	model["description"] = "testString"
	model["name"] = "testString"
	model["authorizations"] = []interface{}{projectConfigAuthModel}
	model["inputs"] = map[string]interface{}{"anyKey": "anyValue"}
	model["settings"] = map[string]interface{}{"anyKey": "anyValue"}
	model["environment_id"] = "testString"

	result, err := project.ResourceIbmProjectConfigMapToProjectConfigDefinitionPatchResourceConfigDefinitionPropertiesPatch(model)
	assert.Nil(t, err)
	checkResult(result)
}
