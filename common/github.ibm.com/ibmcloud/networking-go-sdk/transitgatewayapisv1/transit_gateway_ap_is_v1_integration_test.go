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

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
	"os"
	"time"
)

var GATEWAY_INSTANCE_ID string
var CONN_INSTANCE_ID string

var _ = Describe(`TransitGatewayApIsV1`, func() {
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
	options := &transitgatewayapisv1.TransitGatewayApIsV1Options{
		ServiceName:   "TransitGatewayApIsV1_Mocking",
		Authenticator: authenticator,
		URL:           serviceURL,
		Version:       &version,
	}
	service, err := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(options)
	It(`Successfully created TransitGatewayApIsV1 service instance`, func() {
		Expect(err).To(BeNil())
	})

	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions)`, func() {
		Context(`Success: create Transit Gateway`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			name := "SDK-INT-gateway"
			location := os.Getenv("LOCATION")
			createTransitGatewayOptions := service.NewCreateTransitGatewayOptions(location, name)
			createTransitGatewayOptions.SetHeaders(header)
			//createTransitGatewayOptions.SetResourceGroup(service.NewResourceGroupIdentity(resourceGroup))

			It(`Successfully created new gateway`, func() {
				result, detailedResponse, err := service.CreateTransitGateway(createTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))
				Expect(*result.Name).To(Equal("SDK-INT-gateway"))
				Expect(*result.Location).To(Equal(os.Getenv("LOCATION")))
				os.Setenv("GATEWAY_INSTANCE_ID", *result.ID)
				//time.Sleep(30 * time.Second)
			})
			It("Successfully waits for gateway to report as available", func() {
				detailTransitGatewayOptions := service.NewDetailTransitGatewayOptions(os.Getenv("GATEWAY_INSTANCE_ID"))

				// Gateway creation might not be instantaneous.  Poll the Gateway looking for 'available' status.  Fail after 2 min
				timer := 0
				for {
					response, _, _ := service.DetailTransitGateway(detailTransitGatewayOptions)

					// if available then we are done
					if *response.Status == "available" {
						Expect(*response.Status).To(Equal("available")) // response is available, exit success
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

				firstResource := result.TransitGateways[0]
				Expect(*firstResource.ID).ToNot(BeNil())
				Expect(*firstResource.Name).ToNot(BeNil())
			})
		})
	})

	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions)`, func() {
		Context(`Success: update Gateway name by instanceID`, func() {

			name := "update-SDK-INT-gateway"
			It(`Successfully update name of gateway`, func() {
				gateway_id := os.Getenv("GATEWAY_INSTANCE_ID")
				fmt.Println("Gateway to update ", gateway_id)
				updateTransitGatewayOptions := service.NewUpdateTransitGatewayOptions(gateway_id).
					SetName(name)

				result, detailedResponse, err := service.UpdateTransitGateway(updateTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Name).To(Equal(name))
			})
		})

		Context(`Failed to update resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			instanceName := "To Update"
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

	Describe(`DetailTransitGateway(detailTransitGatewayOptions *DetailTransitGatewayOptions)`, func() {
		Context(`Successfully get gateway by instanceID`, func() {

			It(`Successfully get resource by instanceID`, func() {
				gateway_id := os.Getenv("GATEWAY_INSTANCE_ID")
				detailTransitGatewayOptions := service.NewDetailTransitGatewayOptions(gateway_id)

				result, detailedResponse, err := service.DetailTransitGateway(detailTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(gateway_id))
			})
		})

		Context(`Failed to get resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			detailTransitGatewayOptions := &transitgatewayapisv1.DetailTransitGatewayOptions{}
			detailTransitGatewayOptions.SetID(badinstanceID)
			detailTransitGatewayOptions.SetHeaders(header)
			It(`Failed to get resource by instanceID`, func() {
				result, detailedResponse, err := service.DetailTransitGateway(detailTransitGatewayOptions)
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
				name := "SDK-INT-gateway-conn"
				crn := os.Getenv("VPC_CRN")
				createTransitGatewayConnectionOptions := service.NewCreateTransitGatewayConnectionOptions(gatewayID, network_type)
				createTransitGatewayConnectionOptions.SetHeaders(header)
				createTransitGatewayConnectionOptions.SetName(name)
				createTransitGatewayConnectionOptions.SetNetworkID(crn)

				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))
				Expect(*result.Name).To(Equal(name))
				os.Setenv("CONN_INSTANCE_ID", *result.ID)
				//time.Sleep(60 * time.Second)
			})
			It("Successfully waits for connection to report as attached", func() {
				detailTransitGatewayConnectionOptions := service.NewDetailTransitGatewayConnectionOptions(os.Getenv("GATEWAY_INSTANCE_ID"), os.Getenv("CONN_INSTANCE_ID"))

				// Connection creation might not be instantaneous.  Poll the Conn looking for 'attached' status.  Fail after 2 min
				timer := 0
				for {
					response, _, _ := service.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptions)

					// if attached then we are done
					if *response.Status == "attached" {
						Expect(*response.Status).To(Equal("attached")) // response is attached, exit success
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

				firstResource := result.Connections[0]
				Expect(*firstResource.ID).ToNot(BeNil())
				Expect(*firstResource.Name).ToNot(BeNil())
			})
		})
	})

	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions)`, func() {
		Context(`Successfully update GatewayConnection name by instanceID`, func() {
			It(`Successfully update resource by instanceID`, func() {
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CONN_INSTANCE_ID")
				name := "update-SDK-INT-gateway"
				updateTransitGatewayConnectionOptions := service.NewUpdateTransitGatewayConnectionOptions(gatewayID, instanceID).
					SetName(name)

				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Name).To(Equal(name))
			})
		})

		Context(`Failed to update resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			instanceName := "To Update"
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

	Describe(`DetailTransitGatewayConnection(detailTransitGatewayConnectionOptions *DetailTransitGatewayConnectionOptions)`, func() {
		Context(`Successfully get gateway by instanceID`, func() {
			It(`Successfully get resource by instanceID`, func() {
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CONN_INSTANCE_ID")
				detailTransitGatewayConnectionOptions := service.NewDetailTransitGatewayConnectionOptions(gatewayID, instanceID)

				result, detailedResponse, err := service.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(instanceID))
			})
		})

		Context(`Failed to get resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			detailTransitGatewayConnectionOptions := &transitgatewayapisv1.DetailTransitGatewayConnectionOptions{}
			detailTransitGatewayConnectionOptions.SetTransitGatewayID(badinstanceID)
			detailTransitGatewayConnectionOptions.SetID(badinstanceID)
			detailTransitGatewayConnectionOptions.SetHeaders(header)
			It(`Failed to get resource by instanceID`, func() {
				result, detailedResponse, err := service.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptions)
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
				detailTransitGatewayConnectionOptions := service.NewDetailTransitGatewayConnectionOptions(os.Getenv("GATEWAY_INSTANCE_ID"), os.Getenv("CONN_INSTANCE_ID"))

				// Connection delete might not be instantaneous.  Poll the Conn looking for a not found.  Fail after 4 min
				timer := 0
				for {
					// Get the current rc for the VC
					_, detailedResponse, _ := service.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptions)

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

				firstResource := result.Locations[0]
				Expect(*firstResource.Name).ToNot(BeNil())
			})
		})
	})

	Describe(`DetailGatewayLocation(detailGatewayLocationOptions *DetailGatewayLocationOptions)`, func() {
		Context(`Successfully get location by ID`, func() {
			instanceID := "us-south"
			detailGatewayLocationOptions := service.NewDetailGatewayLocationOptions(instanceID)
			It(`Successfully get location by instanceID`, func() {
				result, detailedResponse, err := service.DetailGatewayLocation(detailGatewayLocationOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Name).To(Equal(instanceID))
			})
		})

		Context(`Failed to get location by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			detailGatewayLocationOptions := &transitgatewayapisv1.DetailGatewayLocationOptions{}
			detailGatewayLocationOptions.SetName(badinstanceID)
			detailGatewayLocationOptions.SetHeaders(header)
			It(`Failed to get resource by instanceID`, func() {
				result, detailedResponse, err := service.DetailGatewayLocation(detailGatewayLocationOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
