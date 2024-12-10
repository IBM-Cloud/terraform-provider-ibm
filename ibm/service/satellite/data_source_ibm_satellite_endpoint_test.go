// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmSatelliteEndpointDataSourceBasic(t *testing.T) {
	endpointLocationID := fmt.Sprintf("tf-satellite-loc-%d", acctest.RandIntRange(10, 100))
	connType := "location"
	displayName := fmt.Sprintf("tf-display-name-%d", acctest.RandIntRange(10, 100))
	serverHost := "cloud.ibm.com"
	serverPort := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	serverProtocol := "tls"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteEndpointDataSourceConfigBasic(endpointLocationID, connType, displayName, serverHost, serverPort, serverProtocol),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "endpoint_id"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "connection_type"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_host"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_port"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_mutual_auth"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_mutual_auth"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "reject_unauth"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "sources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "connector_port"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "service_name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_host"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_port"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "certs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "last_change"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "performance.#"),
				),
			},
		},
	})
}

func TestAccIbmSatelliteEndpointDataSourceAllArgs(t *testing.T) {
	endpointLocationID := fmt.Sprintf("tf-satellite-loc-%d", acctest.RandIntRange(10, 100))
	endpointConnType := "location"
	endpointDisplayName := fmt.Sprintf("tf-display-name-%d", acctest.RandIntRange(10, 100))
	endpointServerHost := "cloud.ibm.com"
	endpointServerPort := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	endpointSni := fmt.Sprintf("tf-sni-%d", acctest.RandIntRange(10, 100))
	endpointClientProtocol := "https"
	endpointClientMutualAuth := "true"
	endpointServerProtocol := "tls"
	endpointServerMutualAuth := "true"
	endpointRejectUnauth := "false"
	endpointTimeout := fmt.Sprintf("%d", acctest.RandIntRange(1, 180))
	endpointCreatedBy := fmt.Sprintf("tf_created_by_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteEndpointDataSourceConfig(endpointLocationID, endpointConnType, endpointDisplayName, endpointServerHost, endpointServerPort, endpointSni, endpointClientProtocol, endpointClientMutualAuth, endpointServerProtocol, endpointServerMutualAuth, endpointRejectUnauth, endpointTimeout, endpointCreatedBy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "endpoint_id"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "connection_type"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_host"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_port"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_mutual_auth"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "server_mutual_auth"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "reject_unauth"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "connector_port"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "service_name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_host"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "client_port"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "certs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "last_change"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_endpoint.satellite_endpoint", "performance.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSatelliteEndpointDataSourceConfigBasic(locationID string, connType string, displayName string, serverHost string, serverPort string, serverProtocol string) string {
	return fmt.Sprintf(`
		resource "ibm_satellite_endpoint" "satellite_endpoint" {
			location  = "%s"
			connection_type = "%s"
			display_name = "%s"
			server_host = "%s"
			server_port = %s
			server_protocol = "%s"
			client_protocol = "https"
		}

		data "ibm_satellite_endpoint" "satellite_endpoint" {
			location = ibm_satellite_endpoint.satellite_endpoint.location
			endpoint_id = ibm_satellite_endpoint.satellite_endpoint.endpoint_id
		}
	`, locationID, connType, displayName, serverHost, serverPort, serverProtocol)
}

func testAccCheckIbmSatelliteEndpointDataSourceConfig(endpointLocationID string, endpointConnType string, endpointDisplayName string, endpointServerHost string, endpointServerPort string, endpointSni string, endpointClientProtocol string, endpointClientMutualAuth string, endpointServerProtocol string, endpointServerMutualAuth string, endpointRejectUnauth string, endpointTimeout string, endpointCreatedBy string) string {
	return fmt.Sprintf(`
		resource "ibm_satellite_endpoint" "satellite_endpoint" {
			location  = "%s"
			connection_type = "%s"
			display_name = "%s"
			server_host = "%s"
			server_port = %s
			client_protocol = "%s"
			client_mutual_auth = %s
			server_protocol = "%s"
			server_mutual_auth = %s
			reject_unauth = %s
			timeout = %s
			created_by = "%s"
		}

		data "ibm_satellite_endpoint" "satellite_endpoint" {
			location = ibm_satellite_endpoint.satellite_endpoint.location
			endpoint_id = ibm_satellite_endpoint.satellite_endpoint.endpoint_id
		}
	`, endpointLocationID, endpointConnType, endpointDisplayName, endpointServerHost, endpointServerPort, endpointClientProtocol, endpointClientMutualAuth, endpointServerProtocol, endpointServerMutualAuth, endpointRejectUnauth, endpointTimeout, endpointCreatedBy)
}
