/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package transitgatewayapisv1_test

/*

How to run this test:

go test -v ./transitgatewayapisv1

*/

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
)

var _ = Describe(`TransitGatewayApisV1`, func() {
	err := godotenv.Load("../transit.env")
	It(`Successfully loading .env file`, func() {
		Expect(err).To(BeNil())
	})

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
		URL:    "https://iam.test.cloud.ibm.com/identity/token",
	}

	version := strfmt.Date(time.Now())
	serviceURL := os.Getenv("SERVICE_URL")
	options := &transitgatewayapisv1.TransitGatewayApisV1Options{
		ServiceName:   "TransitGatewayApisV1_Mocking",
		Authenticator: authenticator,
		URL:           serviceURL,
		Version:       &version,
	}
	service, err := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(options)
	It(`Successfully created TransitGatewayApisV1 service instance`, func() {
		Expect(err).To(BeNil())
	})

	timestamp := time.Now().Unix()
	name := "GO-SDK-INT-GATEWAY-" + strconv.FormatInt(timestamp, 10)
	updateName := "GO-SDK-INT-GATEWAY-UPDATE-" + strconv.FormatInt(timestamp, 10)
	connectionName := "GO-SDK-INT-CONNECTION-" + strconv.FormatInt(timestamp, 10)
	updateConnectionName := "GO-SDK-INT-CONNECTION-UPDATE-" + strconv.FormatInt(timestamp, 10)

	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions)`, func() {
		Context(`Success: create Transit Gateway`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}

			location := os.Getenv("LOCATION")
			createTransitGatewayOptions := service.NewCreateTransitGatewayOptions(location, name)
			createTransitGatewayOptions.SetHeaders(header)

			It(`Successfully created new gateway`, func() {
				result, detailedResponse, err := service.CreateTransitGateway(createTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))
				Expect(*result.Name).To(Equal(name))
				Expect(*result.Location).To(Equal(os.Getenv("LOCATION")))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Crn).NotTo(Equal(""))
				Expect(*result.Global).NotTo(BeNil())
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))

				os.Setenv("GATEWAY_INSTANCE_ID", *result.ID)
				//time.Sleep(30 * time.Second)
			})
			It("Successfully waits for gateway to report as available", func() {
				getTransitGatewayOptions := service.NewGetTransitGatewayOptions(os.Getenv("GATEWAY_INSTANCE_ID"))

				// Gateway creation might not be instantaneous.  Poll the Gateway looking for 'available' status.  Fail after 2 min
				timer := 0
				for {
					response, _, _ := service.GetTransitGateway(getTransitGatewayOptions)

					// if available then we are done
					if *response.Status == "available" {
						Expect(*response.Status).To(Equal("available")) // response is available, exit success
						Expect(*response.Name).To(Equal(name))
						Expect(*response.Location).To(Equal(os.Getenv("LOCATION")))
						Expect(*response.CreatedAt).NotTo(Equal(""))
						Expect(*response.ID).To(Equal(os.Getenv("GATEWAY_INSTANCE_ID")))
						Expect(*response.Crn).NotTo(Equal(""))
						Expect(*response.Global).NotTo(BeNil())
						Expect(*response.ResourceGroup.ID).NotTo(Equal(""))
						break
					}

					// other than available, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5sec)
						Expect(*response.Status).To(Equal("available")) // timed out fail if status is not available
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})
		})

		Context(`Fail: to create new gateway`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			createTransitGatewayOptions := &transitgatewayapisv1.CreateTransitGatewayOptions{}
			createTransitGatewayOptions.SetName("testString")
			createTransitGatewayOptions.SetLocation("testString")
			createTransitGatewayOptions.SetHeaders(header)
			It(`Fail to create new resource`, func() {
				result, detailedResponse, err := service.CreateTransitGateway(createTransitGatewayOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).ToNot(Equal(200))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`ListTransitGateways(listTransitGatewaysOptions *ListTransitGatewaysOptions)`, func() {
		Context(`Successfully list Transit Gateways`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			listTransitGatewaysOptions := service.NewListTransitGatewaysOptions().
				SetHeaders(header)

			It(`Successfully list all gateways`, func() {
				result, detailedResponse, err := service.ListTransitGateways(listTransitGatewaysOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(len(result.TransitGateways)).Should(BeNumerically(">", 0))
				found := false
				for _, gw := range result.TransitGateways {
					if *gw.ID == os.Getenv("GATEWAY_INSTANCE_ID") {
						Expect(*gw.Name).To(Equal(name))
						Expect(*gw.Status).To(Equal("available"))
						Expect(*gw.Location).To(Equal(os.Getenv("LOCATION")))
						Expect(*gw.CreatedAt).NotTo(Equal(""))
						Expect(*gw.UpdatedAt).NotTo(Equal(""))
						Expect(*gw.Crn).NotTo(Equal(""))
						Expect(*gw.Global).NotTo(BeNil())
						Expect(*gw.ResourceGroup.ID).NotTo(Equal(""))
						found = true
						break
					}
				}
				Expect(found).To(Equal(true))
			})
		})
	})

	Describe(`GetTransitGateway(getTransitGatewayOptions *GetTransitGatewayOptions)`, func() {
		Context(`Successfully get gateway by instanceID`, func() {

			It(`Successfully get resource by instanceID`, func() {
				gateway_id := os.Getenv("GATEWAY_INSTANCE_ID")
				getTransitGatewayOptions := service.NewGetTransitGatewayOptions(gateway_id)

				result, detailedResponse, err := service.GetTransitGateway(getTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(name))
				Expect(*result.Location).To(Equal(os.Getenv("LOCATION")))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				//				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("available"))
				Expect(*result.ID).To(Equal(gateway_id))
				Expect(*result.Crn).NotTo(Equal(""))
				Expect(*result.Global).NotTo(BeNil())
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
			})
		})

		Context(`Failed to get resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			getTransitGatewayOptions := &transitgatewayapisv1.GetTransitGatewayOptions{}
			getTransitGatewayOptions.SetID(badinstanceID)
			getTransitGatewayOptions.SetHeaders(header)
			It(`Failed to get resource by instanceID`, func() {
				result, detailedResponse, err := service.GetTransitGateway(getTransitGatewayOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions)`, func() {
		Context(`Success: update Gateway name by instanceID`, func() {

			It(`Successfully update name of gateway`, func() {
				gateway_id := os.Getenv("GATEWAY_INSTANCE_ID")
				fmt.Println("Gateway to update ", gateway_id)
				updateTransitGatewayOptions := service.NewUpdateTransitGatewayOptions(gateway_id).
					SetName(updateName)

				result, detailedResponse, err := service.UpdateTransitGateway(updateTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updateName))
				Expect(*result.Location).To(Equal(os.Getenv("LOCATION")))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("available"))
				Expect(*result.ID).To(Equal(gateway_id))
				Expect(*result.Crn).NotTo(Equal(""))
				Expect(*result.Global).NotTo(BeNil())
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
			})
		})

		Context(`Failed to update resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			instanceName := "To Update-" + strconv.FormatInt(timestamp, 10)
			updateTransitGatewayOptions := &transitgatewayapisv1.UpdateTransitGatewayOptions{}
			updateTransitGatewayOptions.SetID(badinstanceID)
			updateTransitGatewayOptions.SetName(instanceName)
			updateTransitGatewayOptions.SetHeaders(header)
			It(`Failed to update gateway by instanceID`, func() {
				result, detailedResponse, err := service.UpdateTransitGateway(updateTransitGatewayOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
	Describe(`CreateTransitGatewayConnection(createTransitGatewayConnectionOptions *CreateTransitGatewayConnectionOptions)`, func() {
		Context(`Success: create Transit Gateway Connection`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			It(`Successfully create new resource`, func() {
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				network_type := "vpc"
				crn := os.Getenv("VPC_CRN")
				createTransitGatewayConnectionOptions := service.NewCreateTransitGatewayConnectionOptions(gatewayID, network_type)
				createTransitGatewayConnectionOptions.SetHeaders(header)
				createTransitGatewayConnectionOptions.SetName(connectionName)
				createTransitGatewayConnectionOptions.SetNetworkID(crn)

				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("CONN_INSTANCE_ID", *result.ID)

				Expect(*result.Name).To(Equal(connectionName))
				Expect(*result.NetworkID).To(Equal(crn))
				Expect(*result.NetworkType).To(Equal(network_type))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
			})

			It("Successfully waits for connection to report as attached", func() {

				getTransitGatewayConnectionOptions := service.NewGetTransitGatewayConnectionOptions(os.Getenv("GATEWAY_INSTANCE_ID"), os.Getenv("CONN_INSTANCE_ID"))

				// Connection creation might not be instantaneous.  Poll the Conn looking for 'attached' status.  Fail after 2 min
				timer := 0
				for {
					response, _, _ := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)

					// if attached then we are done
					if *response.Status == "attached" {
						// response is attached, exit success
						Expect(*response.Name).To(Equal(connectionName))
						Expect(*response.NetworkID).To(Equal(os.Getenv("VPC_CRN")))
						Expect(*response.NetworkType).To(Equal("vpc"))
						Expect(*response.CreatedAt).NotTo(Equal(""))
						Expect(*response.UpdatedAt).NotTo(Equal(""))
						Expect(*response.ID).To(Equal(os.Getenv("CONN_INSTANCE_ID")))
						break
					}

					// other than attached, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5sec)
						Expect(*response.Status).To(Equal("attached")) // timed out fail if status is not attached
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})
		})

		Context(`Fail to create resource instance`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			createTransitGatewayConnectionOptions := &transitgatewayapisv1.CreateTransitGatewayConnectionOptions{}
			createTransitGatewayConnectionOptions.SetName("testString")
			createTransitGatewayConnectionOptions.SetTransitGatewayID("testString")
			createTransitGatewayConnectionOptions.SetNetworkType("testString")
			createTransitGatewayConnectionOptions.SetHeaders(header)
			It(`Fail to create new resource`, func() {
				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).ToNot(Equal(200))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`ListTransitGatewayConnections(listTransitGatewayConnectionsOptions *ListTransitGatewayConnectionsOptions)`, func() {
		Context(`Successfully list Transit GatewayConnections`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			It(`Successfully list all gateways`, func() {
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				listTransitGatewayConnectionsOptions := service.NewListTransitGatewayConnectionsOptions(gatewayID).
					SetTransitGatewayID(gatewayID).
					SetHeaders(header)

				result, detailedResponse, err := service.ListTransitGatewayConnections(listTransitGatewayConnectionsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(len(result.Connections)).Should(BeNumerically(">", 0))
				found := false
				for _, conn := range result.Connections {
					if *conn.ID == os.Getenv("CONN_INSTANCE_ID") {
						Expect(*conn.Name).To(Equal(connectionName))
						Expect(*conn.Status).To(Equal("attached"))
						Expect(*conn.NetworkID).To(Equal(os.Getenv("VPC_CRN")))
						Expect(*conn.NetworkType).To(Equal("vpc"))
						Expect(*conn.CreatedAt).NotTo(Equal(""))
						Expect(*conn.UpdatedAt).NotTo(Equal(""))

						found = true
						break
					}
				}
				Expect(found).To(Equal(true))
			})
		})
	})

	Describe(`GetTransitGatewayConnection(getTransitGatewayConnectionOptions *GetTransitGatewayConnectionOptions)`, func() {
		Context(`Successfully get gateway by instanceID`, func() {
			It(`Successfully get resource by instanceID`, func() {
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CONN_INSTANCE_ID")
				getTransitGatewayConnectionOptions := service.NewGetTransitGatewayConnectionOptions(gatewayID, instanceID)

				result, detailedResponse, err := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(instanceID))
				Expect(*result.Name).To(Equal(connectionName))
				Expect(*result.NetworkID).To(Equal(os.Getenv("VPC_CRN")))
				Expect(*result.NetworkType).To(Equal("vpc"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
			})
		})

		Context(`Failed to get resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			getTransitGatewayConnectionOptions := &transitgatewayapisv1.GetTransitGatewayConnectionOptions{}
			getTransitGatewayConnectionOptions.SetTransitGatewayID(badinstanceID)
			getTransitGatewayConnectionOptions.SetID(badinstanceID)
			getTransitGatewayConnectionOptions.SetHeaders(header)
			It(`Failed to get resource by instanceID`, func() {
				result, detailedResponse, err := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions)`, func() {
		Context(`Successfully update GatewayConnection name by instanceID`, func() {
			It(`Successfully update resource by instanceID`, func() {
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CONN_INSTANCE_ID")
				updateTransitGatewayConnectionOptions := service.NewUpdateTransitGatewayConnectionOptions(gatewayID, instanceID).
					SetName(updateConnectionName)

				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updateConnectionName))
				Expect(*result.NetworkID).To(Equal(os.Getenv("VPC_CRN")))
				Expect(*result.NetworkType).To(Equal("vpc"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.ID).To(Equal(os.Getenv("CONN_INSTANCE_ID")))
				Expect(*result.Status).To(Equal("attached"))
			})
		})

		Context(`Failed to update resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			instanceName := "To Update-" + strconv.FormatInt(timestamp, 10)
			updateTransitGatewayConnectionOptions := &transitgatewayapisv1.UpdateTransitGatewayConnectionOptions{}
			updateTransitGatewayConnectionOptions.SetTransitGatewayID(badinstanceID)
			updateTransitGatewayConnectionOptions.SetID(badinstanceID)
			updateTransitGatewayConnectionOptions.SetName(instanceName)
			updateTransitGatewayConnectionOptions.SetHeaders(header)
			It(`Failed to update gateway by instanceID`, func() {
				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions *DeleteTransitGatewayConnectionOptions)`, func() {
		Context(`Successfully delete resource by instanceID`, func() {
			It(`Successfully delete resource by instanceID`, func() {
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CONN_INSTANCE_ID")
				deleteTransitGatewayConnectionOptions := service.NewDeleteTransitGatewayConnectionOptions(gatewayID, instanceID)

				detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
				//time.Sleep(90 * time.Second)
			})
			It("Successfully waits for connection to report as deleted", func() {
				getTransitGatewayConnectionOptions := service.NewGetTransitGatewayConnectionOptions(os.Getenv("GATEWAY_INSTANCE_ID"), os.Getenv("CONN_INSTANCE_ID"))

				// Connection delete might not be instantaneous.  Poll the Conn looking for a not found.  Fail after 4 min
				timer := 0
				for {
					// Get the current rc for the VC
					_, detailedResponse, _ := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)

					// if 404 then we are done
					if detailedResponse.StatusCode == 404 {
						Expect(detailedResponse.StatusCode).To(Equal(404)) // response is 404, exit success
						break
					}

					// other than 404, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 4 min timer (24x10sec)
						Expect(detailedResponse.StatusCode).To(Equal(404)) // timed out fail if code is not 404
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(10) * time.Second)
						timer = timer + 1
					}
				}
			})
		})

		Context(`Failed to delete resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			deleteTransitGatewayConnectionOptions := &transitgatewayapisv1.DeleteTransitGatewayConnectionOptions{}
			deleteTransitGatewayConnectionOptions.SetTransitGatewayID(badinstanceID)
			deleteTransitGatewayConnectionOptions.SetID(badinstanceID)
			deleteTransitGatewayConnectionOptions.SetHeaders(header)
			It(`Failed to delete resource by instanceID`, func() {
				detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`DeleteTransitGateway(deleteTransitGatewayOptions *DeleteTransitGatewayOptions)`, func() {
		Context(`Successfully delete resource by instanceID`, func() {
			It(`Successfully delete resource by instanceID`, func() {
				instanceID := os.Getenv("GATEWAY_INSTANCE_ID")
				deleteTransitGatewayOptions := service.NewDeleteTransitGatewayOptions(instanceID)

				detailedResponse, err := service.DeleteTransitGateway(deleteTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		Context(`Failed to delete resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			deleteTransitGatewayOptions := &transitgatewayapisv1.DeleteTransitGatewayOptions{}
			deleteTransitGatewayOptions.SetID(badinstanceID)
			deleteTransitGatewayOptions.SetHeaders(header)
			It(`Failed to delete resource by instanceID`, func() {
				detailedResponse, err := service.DeleteTransitGateway(deleteTransitGatewayOptions)
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`ListGatewayLocations(listGatewayLocationsOptions *ListGatewayLocationsOptions)`, func() {
		Context(`Successfully list Locations`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			listGatewayLocationsOptions := service.NewListGatewayLocationsOptions().
				SetHeaders(header)

			It(`Successfully list all locations`, func() {
				result, detailedResponse, err := service.ListGatewayLocations(listGatewayLocationsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(len(result.Locations)).Should(BeNumerically(">", 0))

				firstResource := result.Locations[0]
				Expect(*firstResource.Name).ToNot(BeNil())
				Expect(*firstResource.BillingLocation).ToNot(BeNil())
				Expect(*firstResource.Type).ToNot(BeNil())
			})
		})
	})

	Describe(`GetGatewayLocation(getGatewayLocationOptions *GetGatewayLocationOptions)`, func() {
		Context(`Successfully get location by ID`, func() {
			instanceID := "us-south"
			getGatewayLocationOptions := service.NewGetGatewayLocationOptions(instanceID)
			It(`Successfully get location by instanceID`, func() {
				result, detailedResponse, err := service.GetGatewayLocation(getGatewayLocationOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Name).To(Equal(instanceID))
				Expect(*result.BillingLocation).ToNot(BeNil())
				Expect(*result.Type).ToNot(BeNil())
				Expect(len(result.LocalConnectionLocations)).Should(BeNumerically(">", 0))
			})
		})

		Context(`Failed to get location by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			getGatewayLocationOptions := &transitgatewayapisv1.GetGatewayLocationOptions{}
			getGatewayLocationOptions.SetName(badinstanceID)
			getGatewayLocationOptions.SetHeaders(header)
			It(`Failed to get resource by instanceID`, func() {
				result, detailedResponse, err := service.GetGatewayLocation(getGatewayLocationOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
