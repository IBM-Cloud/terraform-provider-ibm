// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmProjectEnvironmentBasic(t *testing.T) {
	var conf projectv1.Environment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectEnvironmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEnvironmentConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectEnvironmentExists("ibm_project_environment.project_environment_instance", conf),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_project_environment.project_environment_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func testAccCheckIbmProjectEnvironmentConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
            location = "us-south"
            resource_group = "Default"
            definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
                monitoring_enabled = true
                auto_deploy = true
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
	`, acc.ProjectsConfigApiKey)
}

func testAccCheckIbmProjectEnvironmentExists(n string, obj projectv1.Environment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getProjectEnvironmentOptions := &projectv1.GetProjectEnvironmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProjectEnvironmentOptions.SetProjectID(parts[0])
		getProjectEnvironmentOptions.SetID(parts[1])

		environment, _, err := projectClient.GetProjectEnvironment(getProjectEnvironmentOptions)
		if err != nil {
			return err
		}

		obj = *environment
		return nil
	}
}

func testAccCheckIbmProjectEnvironmentDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project_environment" {
			continue
		}

		getProjectEnvironmentOptions := &projectv1.GetProjectEnvironmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProjectEnvironmentOptions.SetProjectID(parts[0])
		getProjectEnvironmentOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := projectClient.GetProjectEnvironment(getProjectEnvironmentOptions)

		if err == nil {
			return fmt.Errorf("project_environment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project_environment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmProjectEnvironmentProjectConfigAuthToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectEnvironmentProjectConfigAuthToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentProjectComplianceProfileToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectEnvironmentProjectComplianceProfileToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentProjectComplianceProfileNullableObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectComplianceProfileNullableObject)

	result, err := project.ResourceIbmProjectEnvironmentProjectComplianceProfileNullableObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentProjectComplianceProfileV1ToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectEnvironmentProjectComplianceProfileV1ToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentProjectReferenceToMap(t *testing.T) {
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

	result, err := project.ResourceIbmProjectEnvironmentProjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentProjectDefinitionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectDefinitionReference)
	model.Name = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectEnvironmentProjectDefinitionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentMapToProjectConfigAuth(t *testing.T) {
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

	result, err := project.ResourceIbmProjectEnvironmentMapToProjectConfigAuth(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentMapToProjectComplianceProfile(t *testing.T) {
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

	result, err := project.ResourceIbmProjectEnvironmentMapToProjectComplianceProfile(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentMapToProjectComplianceProfileNullableObject(t *testing.T) {
	checkResult := func(result *projectv1.ProjectComplianceProfileNullableObject) {
		model := new(projectv1.ProjectComplianceProfileNullableObject)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})

	result, err := project.ResourceIbmProjectEnvironmentMapToProjectComplianceProfileNullableObject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectEnvironmentMapToProjectComplianceProfileV1(t *testing.T) {
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

	result, err := project.ResourceIbmProjectEnvironmentMapToProjectComplianceProfileV1(model)
	assert.Nil(t, err)
	checkResult(result)
}
