// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/logs-router-go-sdk/ibmlogsrouteropenapi30v0"
)

func TestAccIbmLogsRouterTenantBasic(t *testing.T) {
	var conf ibmlogsrouteropenapi30v0.TenantDetailsResponse
	targetType := fmt.Sprintf("logdna")
	targetHost := fmt.Sprintf("tf-target-host-%d", acctest.RandIntRange(10, 100))
	targetPort := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	targetInstanceCrn := fmt.Sprintf("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
	targetTypeUpdate := fmt.Sprintf("logdna")
	targetHostUpdate := fmt.Sprintf("tf-target-host-%d", acctest.RandIntRange(10, 100))
	targetPortUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	targetInstanceCrnUpdate := fmt.Sprintf("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsRouterTenantDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantConfigBasic(targetType, targetHost, targetPort, targetInstanceCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsRouterTenantExists("ibm_logs_router_tenant.logs_router_tenant_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "target_type", targetType),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "target_host", targetHost),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "target_instance_crn", targetInstanceCrn),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantConfigBasic(targetTypeUpdate, targetHostUpdate, targetPortUpdate, targetInstanceCrnUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "target_type", targetTypeUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "target_host", targetHostUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "target_port", targetPortUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "target_instance_crn", targetInstanceCrnUpdate),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_logs_router_tenant.logs_router_tenant_instance", //"ibm_logs_router_tenant.logs_router_tenant",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"access_credential"},
			},
		},
	})
}

func testAccCheckIbmLogsRouterTenantConfigBasic(targetType string, targetHost string, targetPort string, targetInstanceCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
			target_type = "%s"
			target_host = "%s"
			target_port = %s
			target_instance_crn = "%s"
			access_credential = "test-credential"
		}
	`, targetType, targetHost, targetPort, targetInstanceCrn)
}

func testAccCheckIbmLogsRouterTenantExists(n string, obj ibmlogsrouteropenapi30v0.TenantDetailsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmLogsRouterOpenApi30Client, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmLogsRouterOpenApi30V0()
		if err != nil {
			return err
		}

		getTenantDetailOptions := &ibmlogsrouteropenapi30v0.GetTenantDetailOptions{}
		tenantId := strfmt.UUID(rs.Primary.ID)
		getTenantDetailOptions.SetTenantID(&tenantId)

		tenantDetailsResponse, _, err := ibmLogsRouterOpenApi30Client.GetTenantDetail(getTenantDetailOptions)
		if err != nil {
			return err
		}

		obj = *tenantDetailsResponse
		return nil
	}
}

func testAccCheckIbmLogsRouterTenantDestroy(s *terraform.State) error {
	ibmLogsRouterOpenApi30Client, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmLogsRouterOpenApi30V0()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_router_tenant" {
			continue
		}

		getTenantDetailOptions := &ibmlogsrouteropenapi30v0.GetTenantDetailOptions{}
		tenantId := strfmt.UUID(rs.Primary.ID)
		getTenantDetailOptions.SetTenantID(&tenantId)

		// Try to find the key
		_, response, err := ibmLogsRouterOpenApi30Client.GetTenantDetail(getTenantDetailOptions)

		if err == nil {
			return fmt.Errorf("logs_router_tenant still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_router_tenant (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
