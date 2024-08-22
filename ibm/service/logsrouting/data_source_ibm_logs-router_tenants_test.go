// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.90.1-64fd3296-20240515-180710
 */

package logsrouting_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMLogsRouterTenantsDataSourceBasic(t *testing.T) {
	tenantName := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTenantsDataSourceConfigBasic(tenantName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_tenants.logs_router_tenants_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMLogsRouterTenantsDataSourceConfigBasic(tenantName string) string {
	return fmt.Sprintf(`
		 resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
			 name = "%s"
			 targets {
				 log_sink_crn = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
				 name = "my-log-sink"
				 type = "logdna"
				 parameters {
					 host = "www.example.com"
					 port = 1
					 access_credential = "%s"
				 }
			 }
		 }
 
		 data "ibm_logs_router_tenants" "logs_router_tenants_instance" {
			 name = ibm_logs_router_tenant.logs_router_tenant_instance.name
		 }
	 `, tenantName, acc.IngestionKey)
}

func TestDataSourceIBMLogsRouterTenantsTenantToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetParametersTypeLogDnaModel := make(map[string]interface{})
		targetParametersTypeLogDnaModel["host"] = "www.example.com"
		targetParametersTypeLogDnaModel["port"] = int(8080)

		targetTypeModel := make(map[string]interface{})
		targetTypeModel["id"] = "C1C1C838-A4AC-4BD7-8BC6-3173B272429D"
		targetTypeModel["log_sink_crn"] = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
		targetTypeModel["name"] = "my-logdna-log-sink"
		targetTypeModel["etag"] = "c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6"
		targetTypeModel["type"] = "logdna"
		targetTypeModel["created_at"] = "2024-06-20T18:30:00.143156Z"
		targetTypeModel["updated_at"] = "2024-06-20T18:30:00.143156Z"
		targetTypeModel["parameters"] = []map[string]interface{}{targetParametersTypeLogDnaModel}

		model := make(map[string]interface{})
		model["id"] = "8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		model["created_at"] = "2024-06-20T18:30:00.143156Z"
		model["updated_at"] = "2024-06-20T18:30:00.143156Z"
		model["crn"] = "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
		model["name"] = "my-logging-tenant"
		model["etag"] = "822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049"
		model["targets"] = []map[string]interface{}{targetTypeModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogDnaModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDna)
	targetParametersTypeLogDnaModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogDnaModel.Port = core.Int64Ptr(int64(8080))

	targetTypeModel := new(ibmcloudlogsroutingv0.TargetTypeLogDna)
	targetTypeModel.ID = CreateMockUUID("C1C1C838-A4AC-4BD7-8BC6-3173B272429D")
	targetTypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
	targetTypeModel.Name = core.StringPtr("my-logdna-log-sink")
	targetTypeModel.Etag = core.StringPtr("c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6")
	targetTypeModel.Type = core.StringPtr("logdna")
	targetTypeModel.CreatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	targetTypeModel.UpdatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	targetTypeModel.Parameters = targetParametersTypeLogDnaModel

	model := new(ibmcloudlogsroutingv0.Tenant)
	model.ID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
	model.CreatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.UpdatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
	model.Name = core.StringPtr("my-logging-tenant")
	model.Etag = core.StringPtr("822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049")
	model.Targets = []ibmcloudlogsroutingv0.TargetTypeIntf{targetTypeModel}

	result, err := logsrouting.DataSourceIBMLogsRouterTenantsTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMLogsRouterTenantsTargetTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetParametersTypeLogDnaModel := make(map[string]interface{})
		targetParametersTypeLogDnaModel["host"] = "www.example.com"
		targetParametersTypeLogDnaModel["port"] = int(1)

		model := make(map[string]interface{})
		model["id"] = "8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		model["log_sink_crn"] = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
		model["name"] = "my-log-sink"
		model["etag"] = "c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6"
		model["type"] = "logdna"
		model["created_at"] = "2024-06-20T18:30:00.143156Z"
		model["updated_at"] = "2024-06-20T18:30:00.143156Z"
		model["parameters"] = []map[string]interface{}{targetParametersTypeLogDnaModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogDnaModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDna)
	targetParametersTypeLogDnaModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogDnaModel.Port = core.Int64Ptr(int64(1))

	model := new(ibmcloudlogsroutingv0.TargetType)
	model.ID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
	model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
	model.Name = core.StringPtr("my-log-sink")
	model.Etag = core.StringPtr("c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6")
	model.Type = core.StringPtr("logdna")
	model.CreatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.UpdatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.Parameters = targetParametersTypeLogDnaModel

	result, err := logsrouting.DataSourceIBMLogsRouterTenantsTargetTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMLogsRouterTenantsTargetParametersTypeLogDnaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDna)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(1))

	result, err := logsrouting.DataSourceIBMLogsRouterTenantsTargetParametersTypeLogDnaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMLogsRouterTenantsTargetTypeLogDnaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetParametersTypeLogDnaModel := make(map[string]interface{})
		targetParametersTypeLogDnaModel["host"] = "www.example.com"
		targetParametersTypeLogDnaModel["port"] = int(8080)

		model := make(map[string]interface{})
		model["id"] = "8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		model["log_sink_crn"] = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
		model["name"] = "my-log-sink"
		model["etag"] = "c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6"
		model["type"] = "logdna"
		model["created_at"] = "2024-06-20T18:30:00.143156Z"
		model["updated_at"] = "2024-06-20T18:30:00.143156Z"
		model["parameters"] = []map[string]interface{}{targetParametersTypeLogDnaModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogDnaModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDna)
	targetParametersTypeLogDnaModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogDnaModel.Port = core.Int64Ptr(int64(8080))

	model := new(ibmcloudlogsroutingv0.TargetTypeLogDna)
	model.ID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
	model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
	model.Name = core.StringPtr("my-log-sink")
	model.Etag = core.StringPtr("c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6")
	model.Type = core.StringPtr("logdna")
	model.CreatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.UpdatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.Parameters = targetParametersTypeLogDnaModel

	result, err := logsrouting.DataSourceIBMLogsRouterTenantsTargetTypeLogDnaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMLogsRouterTenantsTargetTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetParametersTypeLogsModel := make(map[string]interface{})
		targetParametersTypeLogsModel["host"] = "www.example.com"
		targetParametersTypeLogsModel["port"] = int(8080)

		model := make(map[string]interface{})
		model["id"] = "8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
		model["name"] = "my-log-sink"
		model["etag"] = "c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6"
		model["type"] = "logs"
		model["created_at"] = "2024-06-20T18:30:00.143156Z"
		model["updated_at"] = "2024-06-20T18:30:00.143156Z"
		model["parameters"] = []map[string]interface{}{targetParametersTypeLogsModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogs)
	targetParametersTypeLogsModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogsModel.Port = core.Int64Ptr(int64(8080))

	model := new(ibmcloudlogsroutingv0.TargetTypeLogs)
	model.ID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
	model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
	model.Name = core.StringPtr("my-log-sink")
	model.Etag = core.StringPtr("c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6")
	model.Type = core.StringPtr("logs")
	model.CreatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.UpdatedAt = core.StringPtr("2024-06-20T18:30:00.143156Z")
	model.Parameters = targetParametersTypeLogsModel

	result, err := logsrouting.DataSourceIBMLogsRouterTenantsTargetTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMLogsRouterTenantsTargetParametersTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogs)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(1))

	result, err := logsrouting.DataSourceIBMLogsRouterTenantsTargetParametersTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
