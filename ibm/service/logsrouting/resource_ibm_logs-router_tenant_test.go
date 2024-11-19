// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouting_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMLogsRouterTenantBasic(t *testing.T) {
	var conf ibmcloudlogsroutingv0.Tenant
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	host := fmt.Sprintf("www.example.%d.com", acctest.RandIntRange(10, 100))
	crn := "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
	nameUpdate := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	hostUpdate := fmt.Sprintf("www.example.%d.com", acctest.RandIntRange(10, 100))
	crnUpdate := "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterTenantDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTenantConfigBasic(name, crn, host),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterTenantExists("ibm_logs_router_tenant.logs_router_tenant_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTenantConfigBasic(nameUpdate, crnUpdate, hostUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_router_tenant.logs_router_tenant_instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["ibm_logs_router_tenant.logs_router_tenant_instance"]
					if !ok {
						return "", fmt.Errorf("Not found: %s", "ibm_logs_router_tenant.logs_router_tenant_instance")
					}
					return fmt.Sprintf("%s/%s", rs.Primary.ID, rs.Primary.Attributes["region"]), nil
				},
				ImportStateVerifyIgnore: []string{"targets.0.parameters.0.access_credential", "targets.1.parameters.0.access_credential"},
			},
		},
	})
}

func TestAccIBMLogsRouterTenantAllArgs(t *testing.T) {
	var conf ibmcloudlogsroutingv0.Tenant

	name := fmt.Sprintf("tenant-name-%d", acctest.RandIntRange(10, 100))
	host0 := fmt.Sprintf("www.example.%d.com", acctest.RandIntRange(10, 100))
	port0 := acctest.RandIntRange(1, 9999)
	target0Name := fmt.Sprintf("target-%s", acctest.RandString(4))
	accessCredential := fmt.Sprintf("access-%s", acctest.RandString(4))

	nameUpdate := fmt.Sprintf("tenant-name-%d", acctest.RandIntRange(10, 100))
	host0Update := fmt.Sprintf("www.example.%d.com", acctest.RandIntRange(10, 100))
	port0Update := acctest.RandIntRange(1, 9999)
	target0NameUpdate := fmt.Sprintf("target-%s", acctest.RandString(4))
	accessCredentialUpdate := fmt.Sprintf("access-%s", acctest.RandString(4))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterTenantDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTenantConfigAllArgs(name, target0Name, host0, port0, accessCredential),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterTenantExists("ibm_logs_router_tenant.logs_router_tenant_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTenantConfigAllArgs(nameUpdate, target0NameUpdate, host0Update, port0Update, accessCredentialUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_router_tenant.logs_router_tenant_instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["ibm_logs_router_tenant.logs_router_tenant_instance"]
					if !ok {
						return "", fmt.Errorf("Not found: %s", "ibm_logs_router_tenant.logs_router_tenant_instance")
					}
					return fmt.Sprintf("%s/%s", rs.Primary.ID, rs.Primary.Attributes["region"]), nil
				},
				ImportStateVerifyIgnore: []string{"targets.0.parameters.0.access_credential", "targets.1.parameters.0.access_credential"},
			},
		},
	})
}

func testAccCheckIBMLogsRouterTenantConfigBasic(name string, crn string, host string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
			name = "%s"
			region = "br-sao"
			targets {
				log_sink_crn = "%s"
				name = "my-log-sink"
				parameters {
					host = "%s"
					port = 1
					access_credential = "%s"
				}
			}
		}
		`, name, crn, host, acc.IngestionKey)
}

func testAccCheckIBMLogsRouterTenantConfigAllArgs(name string, target0Name string, host0 string, port0 int, accessCredential string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
			name = "%s"
			region = "br-sao"
			targets {
				log_sink_crn = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
				name = "%s"
				parameters {
					host = "%s"
					port = %d
					access_credential = "%s"
				}
			}
		}
		`, name, target0Name, host0, port0, accessCredential)
}

func testAccCheckIBMLogsRouterTenantExists(n string, obj ibmcloudlogsroutingv0.Tenant) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmCloudLogsRoutingClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMCloudLogsRoutingV0()
		if err != nil {
			return err
		}

		getTenantDetailOptions := &ibmcloudlogsroutingv0.GetTenantDetailOptions{}

		tenantId := strfmt.UUID(rs.Primary.ID)
		getTenantDetailOptions.SetTenantID(&tenantId)
		getTenantDetailOptions.SetRegion(rs.Primary.Attributes["region"])

		tenant, _, err := ibmCloudLogsRoutingClient.GetTenantDetail(getTenantDetailOptions)
		if err != nil {
			return err
		}

		obj = *tenant
		return nil
	}
}

func testAccCheckIBMLogsRouterTenantDestroy(s *terraform.State) error {
	ibmCloudLogsRoutingClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMCloudLogsRoutingV0()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_router_tenant" {
			continue
		}

		getTenantDetailOptions := &ibmcloudlogsroutingv0.GetTenantDetailOptions{}

		tenantId := strfmt.UUID(rs.Primary.ID)
		getTenantDetailOptions.SetTenantID(&tenantId)
		getTenantDetailOptions.SetRegion(rs.Primary.Attributes["region"])

		// Try to find the key
		_, response, err := ibmCloudLogsRoutingClient.GetTenantDetail(getTenantDetailOptions)

		if err == nil {
			return fmt.Errorf("logs_router_tenant still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_router_tenant (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMLogsRouterTenantTargetTypeToMap(t *testing.T) {
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

	result, err := logsrouting.ResourceIBMLogsRouterTenantTargetTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantTargetParametersTypeLogDnaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDna)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(1))

	result, err := logsrouting.ResourceIBMLogsRouterTenantTargetParametersTypeLogDnaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantTargetTypeLogDnaToMap(t *testing.T) {
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

	result, err := logsrouting.ResourceIBMLogsRouterTenantTargetTypeLogDnaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantTargetTypeLogsToMap(t *testing.T) {
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

	result, err := logsrouting.ResourceIBMLogsRouterTenantTargetTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantTargetParametersTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogs)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(1))

	result, err := logsrouting.ResourceIBMLogsRouterTenantTargetParametersTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantMapToTargetTypePrototype(t *testing.T) {
	checkResult := func(result ibmcloudlogsroutingv0.TargetTypePrototypeIntf) {
		targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
		targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
		targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
		targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

		model := new(ibmcloudlogsroutingv0.TargetTypePrototype)
		model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
		model.Name = core.StringPtr("my-log-sink")
		model.Parameters = targetParametersTypeLogDnaPrototypeModel

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogDnaPrototypeModel := make(map[string]interface{})
	targetParametersTypeLogDnaPrototypeModel["host"] = "www.example.com"
	targetParametersTypeLogDnaPrototypeModel["port"] = int(1)
	targetParametersTypeLogDnaPrototypeModel["access_credential"] = "ingestion-secret"

	model := make(map[string]interface{})
	model["log_sink_crn"] = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
	model["name"] = "my-log-sink"
	model["parameters"] = []interface{}{targetParametersTypeLogDnaPrototypeModel}

	result, err := logsrouting.ResourceIBMLogsRouterTenantMapToTargetTypePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantMapToTargetParametersTypeLogDnaPrototype(t *testing.T) {
	checkResult := func(result *ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype) {
		model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
		model.Host = core.StringPtr("www.example.com")
		model.Port = core.Int64Ptr(int64(1))
		model.AccessCredential = core.StringPtr("ingestion-secret")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["host"] = "www.example.com"
	model["port"] = int(1)
	model["access_credential"] = "ingestion-secret"

	result, err := logsrouting.ResourceIBMLogsRouterTenantMapToTargetParametersTypeLogDnaPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogDnaPrototype(t *testing.T) {
	checkResult := func(result *ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype) {
		targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
		targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
		targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(8080))
		targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("an-ingestion-secret")

		model := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
		model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
		model.Name = core.StringPtr("my-log-sink")
		model.Parameters = targetParametersTypeLogDnaPrototypeModel

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogDnaPrototypeModel := make(map[string]interface{})
	targetParametersTypeLogDnaPrototypeModel["host"] = "www.example.com"
	targetParametersTypeLogDnaPrototypeModel["port"] = int(8080)
	targetParametersTypeLogDnaPrototypeModel["access_credential"] = "an-ingestion-secret"

	model := make(map[string]interface{})
	model["log_sink_crn"] = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
	model["name"] = "my-log-sink"
	model["parameters"] = []interface{}{targetParametersTypeLogDnaPrototypeModel}

	result, err := logsrouting.ResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogDnaPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogsPrototype(t *testing.T) {
	checkResult := func(result *ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogsPrototype) {
		targetParametersTypeLogsPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype)
		targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
		targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(8080))

		model := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogsPrototype)
		model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
		model.Name = core.StringPtr("my-log-sink")
		model.Parameters = targetParametersTypeLogsPrototypeModel

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsPrototypeModel := make(map[string]interface{})
	targetParametersTypeLogsPrototypeModel["host"] = "www.example.com"
	targetParametersTypeLogsPrototypeModel["port"] = int(8080)

	model := make(map[string]interface{})
	model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
	model["name"] = "my-log-sink"
	model["parameters"] = []interface{}{targetParametersTypeLogsPrototypeModel}

	result, err := logsrouting.ResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogsPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterTenantMapToTargetParametersTypeLogsPrototype(t *testing.T) {
	checkResult := func(result *ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype) {
		model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype)
		model.Host = core.StringPtr("www.example.com")
		model.Port = core.Int64Ptr(int64(1))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["host"] = "www.example.com"
	model["port"] = int(1)

	result, err := logsrouting.ResourceIBMLogsRouterTenantMapToTargetParametersTypeLogsPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
