package directlinkapisv1_test

/*

How to run this test:

go test -v ./directlinkapisv1

*/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/directlinkapisv1"
)

var _ = Describe(`DirectLinkApisV1`, func() {
	err := godotenv.Load("../directlink.env")
	It(`Successfully loading .env file`, func() {
		Expect(err).To(BeNil())
	})

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
		URL:    "https://iam.test.cloud.ibm.com/identity/token",
	}

	//authenticator := &core.BearerTokenAuthenticator{
	//	BearerToken: os.Getenv("BEARER_TOKEN"),
	//}

	version := strfmt.Date(time.Now())
	serviceURL := os.Getenv("SERVICE_URL")
	options := &directlinkapisv1.DirectLinkApisV1Options{
		ServiceName:   "DirectLinkApisV1_Mocking",
		Authenticator: authenticator,
		URL:           serviceURL,
		Version:       &version,
	}
	service, err := directlinkapisv1.NewDirectLinkApisV1UsingExternalConfig(options)
	It(`Successfully created DirectLinkApisV1 service instance`, func() {
		Expect(err).To(BeNil())
	})

	Describe("Direct Link Gateways", func() {

		gatewayName := "GO-INT-SDK"
		updatedGatewayName := "GO-INT-SDK-PATCH"
		bgpAsn := int64(64999)
		bgpBaseCidr := "169.254.0.0/16"
		crossConnectRouter := "LAB-xcr01.dal09"
		global := true
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		metered := false
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"

		invalidGatewayId := "000000000000000000000000000000000000"

		Context("Get non existing gateway", func() {

			getGatewayOptions := service.NewGetGatewayOptions(invalidGatewayId)

			It(`Returns the http response with error code 404`, func() {
				result, detailedResponse, err := service.GetGateway(getGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Gateway"))
				Expect(detailedResponse.StatusCode).To(Equal(404))

			})
		})

		Context("Create gateway", func() {

			gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, bgpBaseCidr, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)

			createGatewayOptions := service.NewCreateGatewayOptions(gateway)

			It("Fails when Invalid BGP is provided", func() {
				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(65000, bgpBaseCidr, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("BGP AS Number is invalid."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It("Fails when invalid bgp_base_cidr is provided", func() {
				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, "169.254.0.0", global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("cidr is not 169.254.0.0/16 range, localIP and remoteIP must be manually assigned"))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It("Fails when invalid speed_mbps is provided", func() {
				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, bgpBaseCidr, global, metered, gatewayName, 10000000000, gatewayType, carrierName, crossConnectRouter, customerName, locationName)

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Location with provided 'linkSpeed' and 'OfferingType'."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It("Fails when invalid locations is provided", func() {
				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, bgpBaseCidr, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, "InvalidCity")

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Location with provided 'linkSpeed' and 'OfferingType'."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It("Successfully Creates a gateway", func() {
				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
			})

			It("Successfully fetches the created Gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				getGatewayOptions := service.NewGetGatewayOptions(gatewayId)

				result, detailedResponse, err := service.GetGateway(getGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
			})

			It("Throws an Error when creating a gateway with same name", func() {
				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("A gateway with the same name already exists"))
				Expect(detailedResponse.StatusCode).To(Equal(409))
			})

		})

		Context("Successfully fetch the gateways list", func() {

			listGatewaysOptions := service.NewListGatewaysOptions()

			It(`Successfully list all gateways`, func() {
				result, detailedResponse, err := service.ListGateways(listGatewaysOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				gateways := result.Gateways
				Expect(len(gateways)).Should(BeNumerically(">", 0))
			})
		})

		Context("Update the Gateway", func() {
			It("Fails if an invalid GatewayID is provided", func() {
				patchGatewayOptions := service.NewUpdateGatewayOptions(invalidGatewayId).SetOperationalStatus("loa_accepted")

				result, detailedResponse, err := service.UpdateGateway(patchGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Gateway"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
			})

			It("Successfully Updates the Gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				patchGatewayOptions := service.NewUpdateGatewayOptions(gatewayId)

				result, detailedResponse, err := service.UpdateGateway(patchGatewayOptions.SetGlobal(false).SetSpeedMbps(int64(1000)).SetName(updatedGatewayName))
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updatedGatewayName))
				Expect(*result.Global).To(Equal(false))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))

				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
			})

			It("Successfully fetches the updated Gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				getGatewayOptions := service.NewGetGatewayOptions(gatewayId)

				result, detailedResponse, err := service.GetGateway(getGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updatedGatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(false))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
			})
		})

		Context("Delete a gateway", func() {
			It("Fails if an invalid GatewayID is provided", func() {
				deteleGatewayOptions := service.NewDeleteGatewayOptions(invalidGatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Gateway"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
			})

			It("Successfully deletes a gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		Context("DirectLink connect gateway", func() {

			// to create a connect gateway, we need to have a port.  List the ports and save the id of the 1st one found
			portId := ""
			It("List ports and save the id of the first port", func() {
				listPortsOptions := service.NewListPortsOptions()
				result, detailedResponse, err := service.ListPorts(listPortsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				portId = *result.Ports[0].ID
			})

			It("create connect gateway", func() {
				sec := time.Now().Unix()
				timestamp := strconv.FormatInt(sec, 10)
				gatewayName = "GO-INT-SDK-CONNECT-" + timestamp
				portIdentity, _ := service.NewGatewayPortIdentity(portId)
				gateway, _ := service.NewGatewayTemplateGatewayTypeConnectTemplate(bgpAsn, bgpBaseCidr, global, metered, gatewayName,
					speedMbps, "connect", portIdentity)
				createGatewayOptions := service.NewCreateGatewayOptions(gateway)
				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				// Save the gateway id for deletion
				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal("connect"))
				Expect(*result.Port.ID).To(Equal(portId))
			})

			It("Successfully waits for connect gateway to be provisioned state", func() {
				getGatewayOptions := service.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

				// before a connect gateway can be deleted, it needs to have operational_status of provisioned.  We need to wait for
				// the new gateway to go to provisioned so we can delete it.
				timer := 0
				for {
					// Get the current status for the gateway
					result, detailedResponse, err := service.GetGateway(getGatewayOptions)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(200))

					// if operational status is "provisioned" then we are done
					if *result.OperationalStatus == "provisioned" {
						Expect(*result.OperationalStatus).To(Equal("provisioned"))
						break
					}

					// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5sec)
						Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully deletes connect gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)
				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})

	Describe("Offering Types", func() {

		Context("Locations", func() {
			It("should fetch the locations for the type dedicated", func() {
				listOfferingTypeLocationsOptions := service.NewListOfferingTypeLocationsOptions("dedicated")
				result, detailedResponse, err := service.ListOfferingTypeLocations(listOfferingTypeLocationsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Locations)).Should(BeNumerically(">", 0))
				os.Setenv("OT_DEDICATED_LOCATION_DISPLAY_NAME", *result.Locations[0].DisplayName)
				os.Setenv("OT_DEDICATED_LOCATION_NAME", *result.Locations[0].Name)
			})

			It("should fetch the locations for the type connect", func() {
				listOfferingTypeLocationsOptions := service.NewListOfferingTypeLocationsOptions("connect")

				result, detailedResponse, err := service.ListOfferingTypeLocations(listOfferingTypeLocationsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Locations)).Should(BeNumerically(">", 0))
				os.Setenv("OT_CONNECT_LOCATION_DISPLAY_NAME", *result.Locations[0].DisplayName)
				os.Setenv("OT_CONNECT_LOCATION_NAME", *result.Locations[0].Name)
			})

			It("should return an error for invalid location type", func() {
				listOfferingTypeLocationsOptions := service.NewListOfferingTypeLocationsOptions("RANDOM")

				result, detailedResponse, err := service.ListOfferingTypeLocations(listOfferingTypeLocationsOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("offering_type_location: RANDOM"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})
		})

		Context("Cross Connect Routers", func() {
			It("should list the location info for type dedicated and location short name", func() {
				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("dedicated", os.Getenv("OT_DEDICATED_LOCATION_NAME"))

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.CrossConnectRouters)).Should(BeNumerically(">", 0))
			})

			It("should list the location info for type dedicated and location display name", func() {
				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("dedicated", os.Getenv("OT_DEDICATED_LOCATION_DISPLAY_NAME"))

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.CrossConnectRouters)).Should(BeNumerically(">", 0))
			})

			It("should return proper error when unsupported offering type CONNECT is provided", func() {
				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("connect", os.Getenv("OT_CONNECT_LOCATION_NAME"))

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("The supplied OfferingType is not supported for this call"))
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(result).To(BeNil())
			})

			It("should return proper error when incorrect offering type is provided", func() {
				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("random", os.Getenv("OT_CONNECT_LOCATION_DISPLAY_NAME"))

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Invalid Direct Link Offering Type."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(result).To(BeNil())
			})

			It("should return proper error when incorrect location is provided", func() {
				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("dedicated", "florida")

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Classic Location not found: florida"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})
		})

		Context("Offering Speeds", func() {
			It("should fetch the offering speeds for the type dedicated", func() {
				listOfferingTypeSpeedsOptions := service.NewListOfferingTypeSpeedsOptions("dedicated")

				result, detailedResponse, err := service.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Speeds)).Should(BeNumerically(">", 0))
				fmt.Println("Offering Speeds ==== ", result.Speeds[0].LinkSpeed)
			})

			It("should fetch the offering speeds for the type connect", func() {
				listOfferingTypeSpeedsOptions := service.NewListOfferingTypeSpeedsOptions("connect")

				result, detailedResponse, err := service.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Speeds)).Should(BeNumerically(">", 0))
				fmt.Println("Offering Speeds ==== ", result.Speeds[0].LinkSpeed)
			})

			It("should proper error for invalid offering type", func() {
				listOfferingTypeSpeedsOptions := service.NewListOfferingTypeSpeedsOptions("random")

				result, detailedResponse, err := service.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find OfferingType"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("Ports", func() {
		It("should fetch the ports", func() {
			listPortsOptions := service.NewListPortsOptions()

			result, detailedResponse, err := service.ListPorts(listPortsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(len(result.Ports)).Should(BeNumerically(">", 0))
			fmt.Println("Ports ID ==== ", result.Ports[0].ID)
			fmt.Println("Ports LocationDisplayName ==== ", result.Ports[0].LocationDisplayName)
			os.Setenv("PORT_ID", *result.Ports[0].ID)
			os.Setenv("PORT_LOCATION_DISPLAY_NAME", *result.Ports[0].LocationDisplayName)
			os.Setenv("PORT_LOCATION_NAME", *result.Ports[0].LocationName)
			os.Setenv("PORT_LABEL", *result.Ports[0].Label)
		})

		It("should fetch the port by ID", func() {
			portId := os.Getenv("PORT_ID")
			locationDisplayName := os.Getenv("PORT_LOCATION_DISPLAY_NAME")
			locationName := os.Getenv("PORT_LOCATION_NAME")
			label := os.Getenv("PORT_LABEL")
			getPortOptions := service.NewGetPortOptions(portId)

			result, detailedResponse, err := service.GetPort(getPortOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(portId))
			Expect(*result.LocationDisplayName).To(Equal(locationDisplayName))
			Expect(*result.LocationName).To(Equal(locationName))
			Expect(*result.Label).To(Equal(label))
		})
	})

	Describe("Direct Link Virtual Connections", func() {
		gatewayName := "GO-INT-VC-SDK"
		bgpAsn := int64(64999)
		bgpBaseCidr := "169.254.0.0/16"
		crossConnectRouter := "LAB-xcr01.dal09"
		global := true
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		metered := false
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"

		Context("Create gateway", func() {

			gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, bgpBaseCidr, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)

			createGatewayOptions := service.NewCreateGatewayOptions(gateway)

			It("Successfully created a gateway", func() {
				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
			})

			It("Successfully create a CLASSIC virtual connection", func() {
				vcName := "GO-INT-CLASSIC-VC-SDK"
				createGatewayVCOptions := service.NewCreateGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), vcName, directlinkapisv1.CreateGatewayVirtualConnectionOptions_Type_Classic)
				result, detailedResponse, err := service.CreateGatewayVirtualConnection(createGatewayVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("CLASSIC_VC_ID", *result.ID)

				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.Type).To(Equal(directlinkapisv1.CreateGatewayVirtualConnectionOptions_Type_Classic))
			})

			It("Successfully get a CLASSIC virtual connection", func() {
				vcName := "GO-INT-CLASSIC-VC-SDK"
				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("CLASSIC_VC_ID"))
				result, detailedResponse, err := service.GetGatewayVirtualConnection(getGatewayVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("CLASSIC_VC_ID")))
				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.Type).To(Equal(directlinkapisv1.CreateGatewayVirtualConnectionOptions_Type_Classic))
			})

			It("Successfully create a Gen 2 VPC virtual connection", func() {
				vcName := "GO-INT-GEN2-VPC-VC-SDK"
				vpcCrn := os.Getenv("GEN2_VPC_CRN")
				createGatewayVCOptions := service.NewCreateGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), vcName, directlinkapisv1.CreateGatewayVirtualConnectionOptions_Type_Vpc)
				createGatewayVCOptionsWithNetworkID := createGatewayVCOptions.SetNetworkID(vpcCrn)
				result, detailedResponse, err := service.CreateGatewayVirtualConnection(createGatewayVCOptionsWithNetworkID)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				// save the id so it can be deleted later
				os.Setenv("GEN2_VPC_VC_ID", *result.ID)

				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.Type).To(Equal(directlinkapisv1.CreateGatewayVirtualConnectionOptions_Type_Vpc))
				Expect(*result.NetworkID).To(Equal(vpcCrn))
			})

			It("Successfully get a Gen 2 VPC virtual connection", func() {
				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("GEN2_VPC_VC_ID"))
				result, detailedResponse, err := service.GetGatewayVirtualConnection(getGatewayVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("GEN2_VPC_VC_ID")))
				Expect(*result.Name).To(Equal("GO-INT-GEN2-VPC-VC-SDK"))
				Expect(*result.Type).To(Equal(directlinkapisv1.CreateGatewayVirtualConnectionOptions_Type_Vpc))
				Expect(*result.NetworkID).To(Equal(os.Getenv("GEN2_VPC_CRN")))
			})

			It("Successfully list the virtual connections for a gateway", func() {
				listVcOptions := service.NewListGatewayVirtualConnectionsOptions(os.Getenv("GATEWAY_ID"))
				result, detailedResponse, err := service.ListGatewayVirtualConnections(listVcOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				vcs := result.VirtualConnections
				// two VCs were created for the GW, so we should expect 2
				Expect(len(vcs)).Should(BeNumerically("==", 2))
			})

			It("Successfully Update a virtual connection name", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("GEN2_VPC_VC_ID")
				vcName := "GO-INT-GEN2-VPC-VC-PATCH-SDK"
				patchGatewayOptions := service.NewUpdateGatewayVirtualConnectionOptions(gatewayId, vcId)
				patchGatewayOptions = patchGatewayOptions.SetName(vcName)

				result, detailedResponse, err := service.UpdateGatewayVirtualConnection(patchGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("GEN2_VPC_VC_ID")))
				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.NetworkID).To(Equal(os.Getenv("GEN2_VPC_CRN")))
			})

			It("Fail to Update a virtual connection status", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("GEN2_VPC_VC_ID")
				patchGatewayOptions := service.NewUpdateGatewayVirtualConnectionOptions(gatewayId, vcId)
				patchGatewayOptions = patchGatewayOptions.SetStatus(directlinkapisv1.UpdateGatewayVirtualConnectionOptions_Status_Rejected)

				result, detailedResponse, err := service.UpdateGatewayVirtualConnection(patchGatewayOptions)

				// GW owner is not allowed to change the status, but the test calls the API with the status parameter to valid it is allowed.
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("gateway owner can't patch vc status."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It("Successfully delete a CLASSIC virtual connection for a gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("CLASSIC_VC_ID")
				deleteClassicVCOptions := service.NewDeleteGatewayVirtualConnectionOptions(gatewayId, vcId)

				detailedResponse, err := service.DeleteGatewayVirtualConnection(deleteClassicVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for CLASSIC virtual connection to report as deleted", func() {
				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("CLASSIC_VC_ID"))

				// VC delete might not be instantaneous.  Poll the VC looking for a not found.  Fail after 2 min
				timer := 0
				for {
					// Get the current rc for the VC
					_, detailedResponse, _ := service.GetGatewayVirtualConnection(getGatewayVCOptions)

					// if 404 then we are done
					if detailedResponse.StatusCode == 404 {
						Expect(detailedResponse.StatusCode).To(Equal(404)) // response is 404, exit success
						break
					}

					// other than 404, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5sec)
						Expect(detailedResponse.StatusCode).To(Equal(404)) // timed out fail if code is not 404
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully deletes GEN 2 VPC virtual connection for a gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("GEN2_VPC_VC_ID")
				deleteVpcVcOptions := service.NewDeleteGatewayVirtualConnectionOptions(gatewayId, vcId)

				detailedResponse, err := service.DeleteGatewayVirtualConnection(deleteVpcVcOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for GEN 2 VPC virtual connection to report as deleted", func() {
				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("GEN2_VPC_VC_ID"))

				// VC delete might not be instantaneous.  Poll the VC looking for a not found.  Fail after 2 min
				timer := 0
				for {
					// Get the current rc for the VC
					_, detailedResponse, _ := service.GetGatewayVirtualConnection(getGatewayVCOptions)

					// if 404 then we are done
					if detailedResponse.StatusCode == 404 {
						Expect(detailedResponse.StatusCode).To(Equal(404)) // response is 404, exit success
						break
					}

					// other than 404, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5 sec)
						Expect(detailedResponse.StatusCode).To(Equal(404)) // timed out fail if code is not 404
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully deletes a gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})

	Describe("LOA and Completion Notice", func() {
		gatewayName := "GO-INT-LOA-SDK"
		bgpAsn := int64(64999)
		bgpBaseCidr := "169.254.0.0/16"
		crossConnectRouter := "LAB-xcr01.dal09"
		global := true
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		metered := false
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"

		// notes about LOA and CN testing.  When a GW is created, a github issue is also created by dl-rest.  The issue is used for managing the LOA and CN.  In normal operation,
		// an LOA is added to the issue via manual GH interaction.  After that occurs and the GH label changed, then CN upload is allowed.  Since we do not have the ability to
		// do the manual steps for integration testing, the test will only do the following
		//	- Issue GET LOA for a gateway.  It will expect a 404 error since no one has added the LOA to the GH issue
		//  - PUT a completion notice to the gw.  It will fail with a 412 error because the GH issue and GW status are in the wrong state due to no manual interaction
		//  - GET CN for a gw.  It will expect a 404 since the CN could not be uploaded
		//
		Context("Create gateway", func() {
			It("Successfully created a gateway", func() {
				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, bgpBaseCidr, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)
				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)
			})

			It("Successfully call loa", func() {
				listLOAOptions := service.NewListGatewayLetterOfAuthorizationOptions(os.Getenv("GATEWAY_ID"))
				result, detailedResponse, err := service.ListGatewayLetterOfAuthorization(listLOAOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Please check whether the resource you are requesting exists."))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})

			It("Successfully PUT completion notice", func() {
				buffer, err := ioutil.ReadFile("completion_notice.pdf")
				Expect(err).To(BeNil())
				r := ioutil.NopCloser(bytes.NewReader(buffer))

				createCNOptions := service.NewCreateGatewayCompletionNoticeOptions(os.Getenv("GATEWAY_ID"))
				createCNOptions.SetUpload(r)

				detailedResponse, err := service.CreateGatewayCompletionNotice(createCNOptions)

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Invalid gateway status to upload completion notice."))
				Expect(detailedResponse.StatusCode).To(Equal(412))
			})

			It("Successfully call completion notice", func() {
				listCNOptions := service.NewListGatewayCompletionNoticeOptions(os.Getenv("GATEWAY_ID"))
				result, detailedResponse, err := service.ListGatewayCompletionNotice(listCNOptions)

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Please check whether the resource you are requesting exists."))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})

			It("Successfully deletes a gateway", func() {
				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})
})
