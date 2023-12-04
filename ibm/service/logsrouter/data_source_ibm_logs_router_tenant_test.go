// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsRouterTenantDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_tenant.logs_router_tenant_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_tenant.logs_router_tenant_instance", "tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsRouterTenantDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
			target_type = "logdna"
			target_host = "tf-target-host-01"
			target_port = 42
			target_instance_crn = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbca::"
			access_credential = "test-credential"
		}
		data "ibm_logs_router_tenant" "logs_router_tenant_instance" {
			tenant_id = ibm_logs_router_tenant.logs_router_tenant_instance.id
		}
	`)
}
