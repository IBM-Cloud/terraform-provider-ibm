// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM-Cloud/container-services-go-sdk/satellitelinkv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmSatelliteEndpointBasic(t *testing.T) {
	var conf satellitelinkv1.Endpoint
	locationID := fmt.Sprintf("tf-location-%d", acctest.RandIntRange(10, 100))
	connType := "location"
	displayName := fmt.Sprintf("tf-display-name-%d", acctest.RandIntRange(10, 100))
	serverHost := "cloud.ibm.com"
	serverPort := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	clientProtocol := "tls"
	serverProtocol := "tls"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSatelliteEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteEndpointConfigBasic(locationID, connType, displayName, serverHost, serverPort, clientProtocol, serverProtocol),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSatelliteEndpointExists("ibm_satellite_endpoint.satellite_endpoint", conf),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "location", locationID),
				),
			},
		},
	})
}

func TestAccIbmSatelliteEndpointAllArgs(t *testing.T) {
	var conf satellitelinkv1.Endpoint
	locationID := fmt.Sprintf("tf-location-%d", acctest.RandIntRange(10, 100))
	connType := "location"
	displayName := fmt.Sprintf("tf-display-name-%d", acctest.RandIntRange(10, 100))
	serverHost := "cloud.ibm.com"
	serverPort := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	sni := fmt.Sprintf("tf-sni-%d", acctest.RandIntRange(10, 100))
	clientProtocol := "https"
	clientMutualAuth := "true"
	serverProtocol := "tls"
	serverMutualAuth := "true"
	rejectUnauth := "false"
	timeout := fmt.Sprintf("%d", acctest.RandIntRange(1, 180))
	createdBy := fmt.Sprintf("tf_created_by_%d", acctest.RandIntRange(10, 100))
	locationIDUpdate := fmt.Sprintf("tf-location-%d", acctest.RandIntRange(10, 100))
	connTypeUpdate := "location"
	displayNameUpdate := fmt.Sprintf("tf-display-name-%d", acctest.RandIntRange(10, 100))
	serverHostUpdate := "cloud.ibmcloud.com"
	serverPortUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	sniUpdate := fmt.Sprintf("tf-sni-%d", acctest.RandIntRange(10, 100))
	clientProtocolUpdate := "tls"
	clientMutualAuthUpdate := "false"
	serverProtocolUpdate := "tls"
	serverMutualAuthUpdate := "false"
	rejectUnauthUpdate := "true"
	timeoutUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1, 180))
	createdByUpdate := fmt.Sprintf("tf_created_by_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSatelliteEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteEndpointConfig(locationID, connType, displayName, serverHost, serverPort, sni, clientProtocol, clientMutualAuth, serverProtocol, serverMutualAuth, rejectUnauth, timeout, createdBy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSatelliteEndpointExists("ibm_satellite_endpoint.satellite_endpoint", conf),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "location", locationID),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "connection_type", connType),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "display_name", displayName),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_host", serverHost),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_port", serverPort),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "sni", sni),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "client_protocol", clientProtocol),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "client_mutual_auth", clientMutualAuth),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_protocol", serverProtocol),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_mutual_auth", serverMutualAuth),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "reject_unauth", rejectUnauth),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "timeout", timeout),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "created_by", createdBy),
				),
			},
			{
				Config: testAccCheckIbmSatelliteEndpointConfig(locationIDUpdate, connTypeUpdate, displayNameUpdate, serverHostUpdate, serverPortUpdate, sniUpdate, clientProtocolUpdate, clientMutualAuthUpdate, serverProtocolUpdate, serverMutualAuthUpdate, rejectUnauthUpdate, timeoutUpdate, createdByUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "location", locationIDUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "connection_type", connTypeUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "display_name", displayNameUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_host", serverHostUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_port", serverPortUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "sni", sniUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "client_protocol", clientProtocolUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "client_mutual_auth", clientMutualAuthUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_protocol", serverProtocolUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "server_mutual_auth", serverMutualAuthUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "reject_unauth", rejectUnauthUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "timeout", timeoutUpdate),
					resource.TestCheckResourceAttr("ibm_satellite_endpoint.satellite_endpoint", "created_by", createdByUpdate),
				),
			},
			{
				ResourceName:            "ibm_satellite_endpoint.satellite_endpoint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"certs"},
			},
		},
	})
}

func testAccCheckIbmSatelliteEndpointConfigBasic(locationID string, connType string, displayName string, serverHost string, serverPort string, clientProtocol, serverProtocol string) string {
	return fmt.Sprintf(`

		resource "ibm_satellite_endpoint" "satellite_endpoint" {
			location = "%s"
			connection_type = "%s"
			display_name = "%s"
			server_host = "%s"
			server_port = %s
			client_protocol = "%s"
			server_protocol = "%s"

			certs {
				client {
				  cert {
					filename = "satellite.pem"
					file_contents = "-----BEGIN CERTIFICATE-----\nMIIDmzCCAoOgAwIBAgIUMVdDe4IYAIMKRixT4v+bbvkXhxMwDQYJKoZIhvcNAQEL\nBQAwHTEbMBkGA1UEAxMScm9vdC1jYS0xNjI2OTYyMzg0MB4XDTIxMDcyMjEzNTYw\nMFoXDTIzMTAyNTEzNTYwMFowYjELMAkGA1UEBhMCVVMxFjAUBgNVBAgTDVNhbiBG\ncmFuY2lzY28xCzAJBgNVBAcTAkNBMRcwFQYDVQQKEw5zeXN0ZW06bWFzdGVyczEV\nMBMGA1UEAxMMc3lzdGVtOmFkbWluMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAuJ5VeZGEA10qgnqhSSeBtEFxrxCDRvXyY7NFw9NYvwy6ihTPpWAISSs6\nHsH1jlXuXgXtizru8QQ4HSGXyYNvFqVreZgrseUE08B3sL9FTzAehO3Q/azm5jTX\nFLTseUE2mSfSeWPMIYTWgvwMBl1DGRO3GATsp7Ep4uRN4zWIOC94JD76rk8HZSn6\nBVcvOyQT4XcyD9xYRcmztUoQjd92OAbI362OdjrJd/GaOFXSMYsK5MqhT/R/5p/I\nkDP+tHcp4m8ECI55//ztZvmxWAZsVymsvtU0HombLGSGO6WHcdqwenDyqX1toq+t\nVmgmkdzglV34dZkz7jxtS3j3GdzpDQIDAQABo4GNMIGKMA4GA1UdDwEB/wQEAwIF\noDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAd\nBgNVHQ4EFgQU4V3A7K5Fe2cO2XesFhvtiwe2QTowHwYDVR0jBBgwFoAUHpq8RvCW\n/vJ31q3868e5CyhhLiAwCwYDVR0RBAQwAoIAMA0GCSqGSIb3DQEBCwUAA4IBAQB0\n0H+QlV6QlALq3tKwPkYaBkYW2m403T/pBtziaxKI5d1GOx4AR4per3IASVWOQYJE\nmA6Iur5dmfKasUJXFZada7KseGFSdrEqG8bhghwn3O0TvCeINgwaiOSS2PRLsIg3\nGo9ErbdhiJgvZ7fSy9B3Gd9SlJpKMNiWrPwyhyv5yGdgxwnKO0hn994eMziag1AH\nfvuIBU0MRYCF3+lq6Nn+ZyUE4H3zv2bvMc9SzRlhDOg6cQUjGWYVTBszR9UmPnC9\nIAzW51A7d2MIEI5Nt7dsuCnnT/TLMOZG4mDnLXrCulvCnaRfk64jq9oUs+K9cYAM\ns3v73KPkn8OrZRh2CKbF\n-----END CERTIFICATE-----" 
				  }
				}
			}
		}
	`, locationID, connType, displayName, serverHost, serverPort, clientProtocol, serverProtocol)
}

func testAccCheckIbmSatelliteEndpointConfig(locationID string, connType string, displayName string, serverHost string, serverPort string, sni string, clientProtocol string, clientMutualAuth string, serverProtocol string, serverMutualAuth string, rejectUnauth string, timeout string, createdBy string) string {
	return fmt.Sprintf(`

		resource "ibm_satellite_endpoint" "satellite_endpoint" {
			location = "%s"
			connection_type = "%s"
			display_name = "%s"
			server_host = "%s"
			server_port = %s
			sni = "%s"
			client_protocol = "%s"
			client_mutual_auth = %s
			server_protocol = "%s"
			server_mutual_auth = %s
			reject_unauth = %s
			timeout = %s
			created_by = "%s"

		}
	`, locationID, connType, displayName, serverHost, serverPort, sni, clientProtocol, clientMutualAuth, serverProtocol, serverMutualAuth, rejectUnauth, timeout, createdBy)
}

func testAccCheckIbmSatelliteEndpointExists(n string, obj satellitelinkv1.Endpoint) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		satelliteLinkClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatellitLinkClientSession()
		if err != nil {
			return err
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getEndpointsOptions := &satellitelinkv1.GetEndpointsOptions{}

		getEndpointsOptions.SetLocationID(parts[0])
		getEndpointsOptions.SetEndpointID(parts[1])

		endpoint, _, err := satelliteLinkClient.GetEndpointsWithContext(nil, getEndpointsOptions)
		if err != nil {
			return err
		}

		obj = *endpoint
		return nil
	}
}

func testAccCheckIbmSatelliteEndpointDestroy(s *terraform.State) error {
	satelliteLinkClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatellitLinkClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_endpoint" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getEndpointsOptions := &satellitelinkv1.GetEndpointsOptions{}

		getEndpointsOptions.SetLocationID(parts[0])
		getEndpointsOptions.SetEndpointID(parts[1])

		// Try to find the key
		_, response, err := satelliteLinkClient.GetEndpointsWithContext(nil, getEndpointsOptions)

		if err == nil {
			return fmt.Errorf("satellite_endpoint still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for satellite_endpoint (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
