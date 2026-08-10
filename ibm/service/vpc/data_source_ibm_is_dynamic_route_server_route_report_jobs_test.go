// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerRouteReportJobsDataSourceBasic(t *testing.T) {
	dynamicRouteServerRouteReportJobDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobsDataSourceConfigBasic(dynamicRouteServerRouteReportJobDynamicRouteServerID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "dynamic_route_server_id"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerRouteReportJobsDataSourceAllArgs(t *testing.T) {
	dynamicRouteServerRouteReportJobDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerRouteReportJobFormat := "json"
	dynamicRouteServerRouteReportJobName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobsDataSourceConfig(dynamicRouteServerRouteReportJobDynamicRouteServerID, dynamicRouteServerRouteReportJobFormat, dynamicRouteServerRouteReportJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.created_at"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.format", dynamicRouteServerRouteReportJobFormat),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.name", dynamicRouteServerRouteReportJobName),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance", "route_report_jobs.0.storage_href"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobsDataSourceConfigBasic(dynamicRouteServerRouteReportJobDynamicRouteServerID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = "%s"
			storage_bucket {
				crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
				name = "bucket-27200-lwx4cfvcue"
			}
		}

		data "ibm_is_dynamic_route_server_route_report_jobs" "is_dynamic_route_server_route_report_jobs_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_id
			sort = "name"
		}
	`, dynamicRouteServerRouteReportJobDynamicRouteServerID)
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobsDataSourceConfig(dynamicRouteServerRouteReportJobDynamicRouteServerID string, dynamicRouteServerRouteReportJobFormat string, dynamicRouteServerRouteReportJobName string) string {
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

		data "ibm_is_dynamic_route_server_route_report_jobs" "is_dynamic_route_server_route_report_jobs_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_id
			sort = "name"
		}
	`, dynamicRouteServerRouteReportJobDynamicRouteServerID, dynamicRouteServerRouteReportJobFormat, dynamicRouteServerRouteReportJobName)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dynamicRouteServerRouteReportSchemaModel := make(map[string]interface{})
		dynamicRouteServerRouteReportSchemaModel["href"] = "https://github.ibm.com/cloud-content/downloads/blob/publish/drs/schema/v1/drs-route-report.json"
		dynamicRouteServerRouteReportSchemaModel["type"] = "json"
		dynamicRouteServerRouteReportSchemaModel["version"] = "v1"

		dynamicRouteServerRouteReportJobStatusReasonModel := make(map[string]interface{})
		dynamicRouteServerRouteReportJobStatusReasonModel["code"] = "cannot_access_storage_bucket"
		dynamicRouteServerRouteReportJobStatusReasonModel["message"] = "A failure occurred"
		dynamicRouteServerRouteReportJobStatusReasonModel["more_info"] = "https://cloud.ibm.com/docs/vpc?topic=vpc-object-storage-prereq"

		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		dynamicRouteServerCloudObjectStorageBucketReferenceModel := make(map[string]interface{})
		dynamicRouteServerCloudObjectStorageBucketReferenceModel["crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
		dynamicRouteServerCloudObjectStorageBucketReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerCloudObjectStorageBucketReferenceModel["name"] = "bucket-27200-lwx4cfvcue"
		dynamicRouteServerCloudObjectStorageBucketReferenceModel["resource_type"] = "cos_bucket"

		cloudObjectStorageObjectReferenceModel := make(map[string]interface{})
		cloudObjectStorageObjectReferenceModel["name"] = "my-dynamic-route-server-route-report.json"

		model := make(map[string]interface{})
		model["completed_at"] = "2026-01-02T03:04:05.006Z"
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["format"] = "json"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/route_report_job/r006-095e9baf-01d4-4e29-986e-20d26606b82a"
		model["id"] = "r006-095e9baf-01d4-4e29-986e-20d26606b82a"
		model["json_schema"] = []map[string]interface{}{dynamicRouteServerRouteReportSchemaModel}
		model["name"] = "my-dynamic-route-server-route-report-1"
		model["resource_type"] = "dynamic_route_server_route_report_job"
		model["started_at"] = "2026-01-02T03:04:05.006Z"
		model["status"] = "deleting"
		model["status_reasons"] = []map[string]interface{}{dynamicRouteServerRouteReportJobStatusReasonModel}
		model["storage_bucket"] = []map[string]interface{}{dynamicRouteServerCloudObjectStorageBucketReferenceModel}
		model["storage_href"] = "cos://us-south/bucket-27200-lwx4cfvcue/my-dynamic-route-server-route-report.json"
		model["storage_object"] = []map[string]interface{}{cloudObjectStorageObjectReferenceModel}

		assert.Equal(t, result, model)
	}

	dynamicRouteServerRouteReportSchemaModel := new(vpcv1.DynamicRouteServerRouteReportSchema)
	dynamicRouteServerRouteReportSchemaModel.Href = core.StringPtr("https://github.ibm.com/cloud-content/downloads/blob/publish/drs/schema/v1/drs-route-report.json")
	dynamicRouteServerRouteReportSchemaModel.Type = core.StringPtr("json")
	dynamicRouteServerRouteReportSchemaModel.Version = core.StringPtr("v1")

	dynamicRouteServerRouteReportJobStatusReasonModel := new(vpcv1.DynamicRouteServerRouteReportJobStatusReason)
	dynamicRouteServerRouteReportJobStatusReasonModel.Code = core.StringPtr("cannot_access_storage_bucket")
	dynamicRouteServerRouteReportJobStatusReasonModel.Message = core.StringPtr("A failure occurred")
	dynamicRouteServerRouteReportJobStatusReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-object-storage-prereq")

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	dynamicRouteServerCloudObjectStorageBucketReferenceModel := new(vpcv1.DynamicRouteServerCloudObjectStorageBucketReference)
	dynamicRouteServerCloudObjectStorageBucketReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue")
	dynamicRouteServerCloudObjectStorageBucketReferenceModel.Deleted = deletedModel
	dynamicRouteServerCloudObjectStorageBucketReferenceModel.Name = core.StringPtr("bucket-27200-lwx4cfvcue")
	dynamicRouteServerCloudObjectStorageBucketReferenceModel.ResourceType = core.StringPtr("cos_bucket")

	cloudObjectStorageObjectReferenceModel := new(vpcv1.CloudObjectStorageObjectReference)
	cloudObjectStorageObjectReferenceModel.Name = core.StringPtr("my-dynamic-route-server-route-report.json")

	model := new(vpcv1.DynamicRouteServerRouteReportJob)
	model.CompletedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Format = core.StringPtr("json")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/route_report_job/r006-095e9baf-01d4-4e29-986e-20d26606b82a")
	model.ID = core.StringPtr("r006-095e9baf-01d4-4e29-986e-20d26606b82a")
	model.JSONSchema = dynamicRouteServerRouteReportSchemaModel
	model.Name = core.StringPtr("my-dynamic-route-server-route-report-1")
	model.ResourceType = core.StringPtr("dynamic_route_server_route_report_job")
	model.StartedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Status = core.StringPtr("deleting")
	model.StatusReasons = []vpcv1.DynamicRouteServerRouteReportJobStatusReason{*dynamicRouteServerRouteReportJobStatusReasonModel}
	model.StorageBucket = dynamicRouteServerCloudObjectStorageBucketReferenceModel
	model.StorageHref = core.StringPtr("cos://us-south/bucket-27200-lwx4cfvcue/my-dynamic-route-server-route-report.json")
	model.StorageObject = cloudObjectStorageObjectReferenceModel

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportSchemaToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportSchemaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobStatusReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerCloudObjectStorageBucketReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerCloudObjectStorageBucketReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsCloudObjectStorageObjectReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "my-object"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CloudObjectStorageObjectReference)
	model.Name = core.StringPtr("my-object")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsCloudObjectStorageObjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
