// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMDLImportRouteFilterDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_import_route_filter.test_dl_import_route_filter"
	gatewayname := fmt.Sprintf("import-route-filter-gateway-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLImportRouteFilterDataSourceConfig(gatewayname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "action"),
				),
			},
		},
	})
}

func testAccCheckIBMDLImportRouteFilterDataSourceConfig(gatewayname string) string {
	return fmt.Sprintf(`
	data "ibm_dl_ports" "ds_dlports" {
	}

	resource ibm_dl_gateway test_dl_gateway {
			bgp_asn =  64999
			global = true
			metered = false
			name = "%s"
			speed_mbps = 1000
			type =  "connect"
			port = data.ibm_dl_ports.ds_dlports.ports[0].port_id
		    import_route_filters {
				action = "deny"
				prefix = "10.10.9.0/24"
				ge =25
				le = 26
			}				
	}
	data "ibm_dl_import_route_filters" "test_dl_import_route_filters" {
		gateway = ibm_dl_gateway.test_dl_gateway.id
    }
	data "ibm_dl_import_route_filter" "test_dl_import_route_filter" {
			gateway = ibm_dl_gateway.test_dl_gateway.id
			im_filter_id = data.ibm_dl_import_route_filters.test_dl_import_route_filters.import_route_filters[0].im_filter_id
	 }
	  `, gatewayname)
}
