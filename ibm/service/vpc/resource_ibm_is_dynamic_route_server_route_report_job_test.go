// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerRouteReportJobBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerRouteReportJob
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerRouteReportJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobConfigBasic(dynamicRouteServerID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerRouteReportJobExists("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "dynamic_route_server_id", dynamicRouteServerID),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerRouteReportJobAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerRouteReportJob
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	format := "json"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	formatUpdate := "json"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerRouteReportJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(dynamicRouteServerID, format, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerRouteReportJobExists("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(dynamicRouteServerID, formatUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "format", formatUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobConfigBasic(dynamicRouteServerID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = "%s"
			storage_bucket {
				crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
				name = "bucket-27200-lwx4cfvcue"
			}
		}
	`, dynamicRouteServerID)
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(dynamicRouteServerID string, format string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = "%s"
			format = "%s"
			name = "%s"
			storage_bucket {
				crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
				name = "bucket-27200-lwx4cfvcue"
			}
		}
	`, dynamicRouteServerID, format, name)
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobExists(n string, obj vpcv1.DynamicRouteServerRouteReportJob) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

		dynamicRouteServerRouteReportJob, _, err := vpcClient.GetDynamicRouteServerRouteReportJob(getDynamicRouteServerRouteReportJobOptions)
		if err != nil {
			return err
		}

		obj = *dynamicRouteServerRouteReportJob
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server_route_report_job" {
			continue
		}

		getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetDynamicRouteServerRouteReportJob(getDynamicRouteServerRouteReportJobOptions)

		if err == nil {
			return fmt.Errorf("DynamicRouteServerRouteReportJob still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DynamicRouteServerRouteReportJob (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["name"] = "bucket-27200-lwx4cfvcue"
		model["resource_type"] = "cos_bucket"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.DynamicRouteServerCloudObjectStorageBucketReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue")
	model.Deleted = deletedModel
	model.Name = core.StringPtr("bucket-27200-lwx4cfvcue")
	model.ResourceType = core.StringPtr("cos_bucket")

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://github.ibm.com/cloud-content/downloads/blob/publish/drs/schema/v1/drs-route-report.json"
		model["type"] = "json"
		model["version"] = "v1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerRouteReportSchema)
	model.Href = core.StringPtr("https://github.ibm.com/cloud-content/downloads/blob/publish/drs/schema/v1/drs-route-report.json")
	model.Type = core.StringPtr("json")
	model.Version = core.StringPtr("v1")

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "cannot_access_storage_bucket"
		model["message"] = "A failure occurred"
		model["more_info"] = "https://cloud.ibm.com/docs/vpc?topic=vpc-object-storage-prereq"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerRouteReportJobStatusReason)
	model.Code = core.StringPtr("cannot_access_storage_bucket")
	model.Message = core.StringPtr("A failure occurred")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-object-storage-prereq")

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "my-object"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CloudObjectStorageObjectReference)
	model.Name = core.StringPtr("my-object")

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentity(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityIntf) {
		model := new(vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentity)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue")
		model.Name = core.StringPtr("bucket-27200-lwx4cfvcue")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
	model["name"] = "bucket-27200-lwx4cfvcue"

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByCRN) {
		model := new(vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName) {
		model := new(vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName)
		model.Name = core.StringPtr("bucket-27200-lwx4cfvcue")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "bucket-27200-lwx4cfvcue"

	result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName(model)
	assert.Nil(t, err)
	checkResult(result)
}
