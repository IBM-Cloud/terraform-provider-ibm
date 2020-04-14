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

package vpcclassicv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`VpcClassicV1`, func() {
	Describe(`ListFloatingIps(listFloatingIpsOptions *ListFloatingIpsOptions)`, func() {
		version := "testString"
		listFloatingIpsPath := "/floating_ips"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listFloatingIpsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips?limit=20"}, "floating_ips": [{"address": "203.0.113.1", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip", "status": "available", "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}], "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}}`)
			}))
			It(`Invoke ListFloatingIps successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListFloatingIps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListFloatingIpsOptions model
				listFloatingIpsOptionsModel := new(vpcclassicv1.ListFloatingIpsOptions)
				listFloatingIpsOptionsModel.Start = core.StringPtr("testString")
				listFloatingIpsOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListFloatingIps(listFloatingIpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ReserveFloatingIp(reserveFloatingIpOptions *ReserveFloatingIpOptions)`, func() {
		version := "testString"
		reserveFloatingIpPath := "/floating_ips"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(reserveFloatingIpPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"address": "203.0.113.1", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip", "status": "available", "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke ReserveFloatingIp successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReserveFloatingIp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(vpcclassicv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the FloatingIPPrototypeFloatingIPByZone model
				floatingIpPrototypeModel := new(vpcclassicv1.FloatingIPPrototypeFloatingIPByZone)
				floatingIpPrototypeModel.Name = core.StringPtr("my-floating-ip")
				floatingIpPrototypeModel.Zone = zoneIdentityModel

				// Construct an instance of the ReserveFloatingIpOptions model
				reserveFloatingIpOptionsModel := new(vpcclassicv1.ReserveFloatingIpOptions)
				reserveFloatingIpOptionsModel.FloatingIPPrototype = floatingIpPrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReserveFloatingIp(reserveFloatingIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ReleaseFloatingIp(releaseFloatingIpOptions *ReleaseFloatingIpOptions)`, func() {
		version := "testString"
		releaseFloatingIpPath := "/floating_ips/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(releaseFloatingIpPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke ReleaseFloatingIp successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.ReleaseFloatingIp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ReleaseFloatingIpOptions model
				releaseFloatingIpOptionsModel := new(vpcclassicv1.ReleaseFloatingIpOptions)
				releaseFloatingIpOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.ReleaseFloatingIp(releaseFloatingIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetFloatingIp(getFloatingIpOptions *GetFloatingIpOptions)`, func() {
		version := "testString"
		getFloatingIpPath := "/floating_ips/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getFloatingIpPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"address": "203.0.113.1", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip", "status": "available", "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetFloatingIp successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetFloatingIp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFloatingIpOptions model
				getFloatingIpOptionsModel := new(vpcclassicv1.GetFloatingIpOptions)
				getFloatingIpOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetFloatingIp(getFloatingIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateFloatingIp(updateFloatingIpOptions *UpdateFloatingIpOptions)`, func() {
		version := "testString"
		updateFloatingIpPath := "/floating_ips/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateFloatingIpPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"address": "203.0.113.1", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip", "status": "available", "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke UpdateFloatingIp successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateFloatingIp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NetworkInterfaceIdentityByID model
				networkInterfaceIdentityModel := new(vpcclassicv1.NetworkInterfaceIdentityByID)
				networkInterfaceIdentityModel.ID = core.StringPtr("10c02d81-0ecb-4dc5-897d-28392913b81e")

				// Construct an instance of the UpdateFloatingIpOptions model
				updateFloatingIpOptionsModel := new(vpcclassicv1.UpdateFloatingIpOptions)
				updateFloatingIpOptionsModel.ID = core.StringPtr("testString")
				updateFloatingIpOptionsModel.Name = core.StringPtr("my-floating-ip")
				updateFloatingIpOptionsModel.Target = networkInterfaceIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateFloatingIp(updateFloatingIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListRegions(listRegionsOptions *ListRegionsOptions)`, func() {
		version := "testString"
		listRegionsPath := "/regions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listRegionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"regions": [{"endpoint": "Endpoint", "href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south", "name": "us-south", "status": "available"}]}`)
			}))
			It(`Invoke ListRegions successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListRegions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRegionsOptions model
				listRegionsOptionsModel := new(vpcclassicv1.ListRegionsOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListRegions(listRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetRegion(getRegionOptions *GetRegionOptions)`, func() {
		version := "testString"
		getRegionPath := "/regions/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getRegionPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"endpoint": "Endpoint", "href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south", "name": "us-south", "status": "available"}`)
			}))
			It(`Invoke GetRegion successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetRegion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRegionOptions model
				getRegionOptionsModel := new(vpcclassicv1.GetRegionOptions)
				getRegionOptionsModel.Name = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetRegion(getRegionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListZones(listZonesOptions *ListZonesOptions)`, func() {
		version := "testString"
		listZonesPath := "/regions/testString/zones"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listZonesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"zones": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1", "region": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south", "name": "us-south"}, "status": "available"}]}`)
			}))
			It(`Invoke ListZones successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListZones(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(vpcclassicv1.ListZonesOptions)
				listZonesOptionsModel.RegionName = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListZones(listZonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetZone(getZoneOptions *GetZoneOptions)`, func() {
		version := "testString"
		getZonePath := "/regions/testString/zones/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getZonePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1", "region": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south", "name": "us-south"}, "status": "available"}`)
			}))
			It(`Invoke GetZone successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(vpcclassicv1.GetZoneOptions)
				getZoneOptionsModel.RegionName = core.StringPtr("testString")
				getZoneOptionsModel.ZoneName = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetZone(getZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListImages(listImagesOptions *ListImagesOptions)`, func() {
		version := "testString"
		listImagesPath := "/images"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listImagesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["resource_group.id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["visibility"]).To(Equal([]string{"private"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/images?limit=20"}, "images": [{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "file": {"size": 4}, "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "minimum_provisioned_size": 22, "name": "my-image", "operating_system": {"architecture": "amd64", "display_name": "Ubuntu Server 16.04 LTS amd64", "family": "Ubuntu Server", "href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-16-amd64", "name": "ubuntu-16-amd64", "vendor": "Canonical", "version": "16.04 LTS"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "visibility": "private"}], "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/images?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}}`)
			}))
			It(`Invoke ListImages successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(vpcclassicv1.ListImagesOptions)
				listImagesOptionsModel.Start = core.StringPtr("testString")
				listImagesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listImagesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listImagesOptionsModel.Visibility = core.StringPtr("private")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListImages(listImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateImage(createImageOptions *CreateImageOptions)`, func() {
		version := "testString"
		createImagePath := "/images"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createImagePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "file": {"size": 4}, "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "minimum_provisioned_size": 22, "name": "my-image", "operating_system": {"architecture": "amd64", "display_name": "Ubuntu Server 16.04 LTS amd64", "family": "Ubuntu Server", "href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-16-amd64", "name": "ubuntu-16-amd64", "vendor": "Canonical", "version": "16.04 LTS"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "visibility": "private"}`)
			}))
			It(`Invoke CreateImage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImageFilePrototype model
				imageFilePrototypeModel := new(vpcclassicv1.ImageFilePrototype)
				imageFilePrototypeModel.Href = core.StringPtr("cos://us-south/custom-image-vpc-bucket/customImage-0.vhd")

				// Construct an instance of the OperatingSystemIdentityByName model
				operatingSystemIdentityModel := new(vpcclassicv1.OperatingSystemIdentityByName)
				operatingSystemIdentityModel.Name = core.StringPtr("ubuntu-16-amd64")

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the ImagePrototypeImageByFile model
				imagePrototypeModel := new(vpcclassicv1.ImagePrototypeImageByFile)
				imagePrototypeModel.Name = core.StringPtr("my-image")
				imagePrototypeModel.ResourceGroup = resourceGroupIdentityModel
				imagePrototypeModel.File = imageFilePrototypeModel
				imagePrototypeModel.OperatingSystem = operatingSystemIdentityModel

				// Construct an instance of the CreateImageOptions model
				createImageOptionsModel := new(vpcclassicv1.CreateImageOptions)
				createImageOptionsModel.ImagePrototype = imagePrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateImage(createImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteImage(deleteImageOptions *DeleteImageOptions)`, func() {
		version := "testString"
		deleteImagePath := "/images/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteImagePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(202)
			}))
			It(`Invoke DeleteImage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(vpcclassicv1.DeleteImageOptions)
				deleteImageOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetImage(getImageOptions *GetImageOptions)`, func() {
		version := "testString"
		getImagePath := "/images/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getImagePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "file": {"size": 4}, "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "minimum_provisioned_size": 22, "name": "my-image", "operating_system": {"architecture": "amd64", "display_name": "Ubuntu Server 16.04 LTS amd64", "family": "Ubuntu Server", "href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-16-amd64", "name": "ubuntu-16-amd64", "vendor": "Canonical", "version": "16.04 LTS"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "visibility": "private"}`)
			}))
			It(`Invoke GetImage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetImageOptions model
				getImageOptionsModel := new(vpcclassicv1.GetImageOptions)
				getImageOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetImage(getImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateImage(updateImageOptions *UpdateImageOptions)`, func() {
		version := "testString"
		updateImagePath := "/images/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateImagePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "file": {"size": 4}, "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "minimum_provisioned_size": 22, "name": "my-image", "operating_system": {"architecture": "amd64", "display_name": "Ubuntu Server 16.04 LTS amd64", "family": "Ubuntu Server", "href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-16-amd64", "name": "ubuntu-16-amd64", "vendor": "Canonical", "version": "16.04 LTS"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "visibility": "private"}`)
			}))
			It(`Invoke UpdateImage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateImageOptions model
				updateImageOptionsModel := new(vpcclassicv1.UpdateImageOptions)
				updateImageOptionsModel.ID = core.StringPtr("testString")
				updateImageOptionsModel.Name = core.StringPtr("my-image")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateImage(updateImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListOperatingSystems(listOperatingSystemsOptions *ListOperatingSystemsOptions)`, func() {
		version := "testString"
		listOperatingSystemsPath := "/operating_systems"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listOperatingSystemsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "operating_systems": [{"architecture": "amd64", "display_name": "Ubuntu Server 16.04 LTS amd64", "family": "Ubuntu Server", "href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-16-amd64", "name": "ubuntu-16-amd64", "vendor": "Canonical", "version": "16.04 LTS"}]}`)
			}))
			It(`Invoke ListOperatingSystems successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListOperatingSystems(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOperatingSystemsOptions model
				listOperatingSystemsOptionsModel := new(vpcclassicv1.ListOperatingSystemsOptions)
				listOperatingSystemsOptionsModel.Start = core.StringPtr("testString")
				listOperatingSystemsOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListOperatingSystems(listOperatingSystemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetOperatingSystem(getOperatingSystemOptions *GetOperatingSystemOptions)`, func() {
		version := "testString"
		getOperatingSystemPath := "/operating_systems/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getOperatingSystemPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"architecture": "amd64", "display_name": "Ubuntu Server 16.04 LTS amd64", "family": "Ubuntu Server", "href": "https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-16-amd64", "name": "ubuntu-16-amd64", "vendor": "Canonical", "version": "16.04 LTS"}`)
			}))
			It(`Invoke GetOperatingSystem successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetOperatingSystem(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOperatingSystemOptions model
				getOperatingSystemOptionsModel := new(vpcclassicv1.GetOperatingSystemOptions)
				getOperatingSystemOptionsModel.Name = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetOperatingSystem(getOperatingSystemOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListInstanceProfiles(listInstanceProfilesOptions *ListInstanceProfilesOptions)`, func() {
		version := "testString"
		listInstanceProfilesPath := "/instance/profiles"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listInstanceProfilesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "profiles": [{"bandwidth": {"type": "fixed", "value": 20000}, "crn": "crn:v1:bluemix:public:is:us-south-1:::instance-profile:bc1-4x16", "family": "balanced", "href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bc1-4x16", "name": "bc1-4x16", "port_speed": {"type": "fixed", "value": 1000}}], "total_count": 132}`)
			}))
			It(`Invoke ListInstanceProfiles successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListInstanceProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceProfilesOptions model
				listInstanceProfilesOptionsModel := new(vpcclassicv1.ListInstanceProfilesOptions)
				listInstanceProfilesOptionsModel.Start = core.StringPtr("testString")
				listInstanceProfilesOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListInstanceProfiles(listInstanceProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetInstanceProfile(getInstanceProfileOptions *GetInstanceProfileOptions)`, func() {
		version := "testString"
		getInstanceProfilePath := "/instance/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getInstanceProfilePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"bandwidth": {"type": "fixed", "value": 20000}, "crn": "crn:v1:bluemix:public:is:us-south-1:::instance-profile:bc1-4x16", "family": "balanced", "href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bc1-4x16", "name": "bc1-4x16", "port_speed": {"type": "fixed", "value": 1000}}`)
			}))
			It(`Invoke GetInstanceProfile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetInstanceProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceProfileOptions model
				getInstanceProfileOptionsModel := new(vpcclassicv1.GetInstanceProfileOptions)
				getInstanceProfileOptionsModel.Name = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetInstanceProfile(getInstanceProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListInstances(listInstancesOptions *ListInstancesOptions)`, func() {
		version := "testString"
		listInstancesPath := "/instances"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listInstancesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["network_interfaces.subnet.id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["network_interfaces.subnet.crn"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["network_interfaces.subnet.name"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances?limit=20"}, "instances": [{"bandwidth": 1000, "boot_volume_attachment": {"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "image": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "name": "my-image"}, "memory": 8, "name": "my-instance", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}], "primary_network_interface": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}, "profile": {"crn": "crn:v1:bluemix:public:is:us-south-1:::instance-profile:bc1-4x16", "href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bc1-4x16", "name": "bc1-4x16"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "failed", "vcpu": {"architecture": "amd64", "count": 4}, "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}], "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "total_count": 132}`)
			}))
			It(`Invoke ListInstances successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstancesOptions model
				listInstancesOptionsModel := new(vpcclassicv1.ListInstancesOptions)
				listInstancesOptionsModel.Start = core.StringPtr("testString")
				listInstancesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listInstancesOptionsModel.NetworkInterfacesSubnetID = core.StringPtr("testString")
				listInstancesOptionsModel.NetworkInterfacesSubnetCrn = core.StringPtr("testString")
				listInstancesOptionsModel.NetworkInterfacesSubnetName = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListInstances(listInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateInstance(createInstanceOptions *CreateInstanceOptions)`, func() {
		version := "testString"
		createInstancePath := "/instances"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createInstancePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"bandwidth": 1000, "boot_volume_attachment": {"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "image": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "name": "my-image"}, "memory": 8, "name": "my-instance", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}], "primary_network_interface": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}, "profile": {"crn": "crn:v1:bluemix:public:is:us-south-1:::instance-profile:bc1-4x16", "href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bc1-4x16", "name": "bc1-4x16"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "failed", "vcpu": {"architecture": "amd64", "count": 4}, "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke CreateInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EncryptionKeyIdentityByCRN model
				encryptionKeyIdentityModel := new(vpcclassicv1.EncryptionKeyIdentityByCRN)
				encryptionKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179")

				// Construct an instance of the VolumeProfileIdentityByName model
				volumeProfileIdentityModel := new(vpcclassicv1.VolumeProfileIdentityByName)
				volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

				// Construct an instance of the SecurityGroupIdentityByID model
				securityGroupIdentityModel := new(vpcclassicv1.SecurityGroupIdentityByID)
				securityGroupIdentityModel.ID = core.StringPtr("be5df5ca-12a0-494b-907e-aa6ec2bfa271")

				// Construct an instance of the SubnetIdentityByID model
				subnetIdentityModel := new(vpcclassicv1.SubnetIdentityByID)
				subnetIdentityModel.ID = core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

				// Construct an instance of the VolumeAttachmentPrototypeInstanceContextVolumeVolumePrototypeInstanceContextVolumePrototypeInstanceContextVolumeByCapacity model
				volumeAttachmentPrototypeInstanceContextVolumeModel := new(vpcclassicv1.VolumeAttachmentPrototypeInstanceContextVolumeVolumePrototypeInstanceContextVolumePrototypeInstanceContextVolumeByCapacity)
				volumeAttachmentPrototypeInstanceContextVolumeModel.EncryptionKey = encryptionKeyIdentityModel
				volumeAttachmentPrototypeInstanceContextVolumeModel.Iops = core.Int64Ptr(int64(10000))
				volumeAttachmentPrototypeInstanceContextVolumeModel.Name = core.StringPtr("my-volume")
				volumeAttachmentPrototypeInstanceContextVolumeModel.Profile = volumeProfileIdentityModel
				volumeAttachmentPrototypeInstanceContextVolumeModel.Capacity = core.Int64Ptr(int64(100))

				// Construct an instance of the VolumePrototypeInstanceByImageContext model
				volumePrototypeInstanceByImageContextModel := new(vpcclassicv1.VolumePrototypeInstanceByImageContext)
				volumePrototypeInstanceByImageContextModel.Capacity = core.Int64Ptr(int64(100))
				volumePrototypeInstanceByImageContextModel.EncryptionKey = encryptionKeyIdentityModel
				volumePrototypeInstanceByImageContextModel.Iops = core.Int64Ptr(int64(10000))
				volumePrototypeInstanceByImageContextModel.Name = core.StringPtr("my-volume")
				volumePrototypeInstanceByImageContextModel.Profile = volumeProfileIdentityModel

				// Construct an instance of the ImageIdentityByID model
				imageIdentityModel := new(vpcclassicv1.ImageIdentityByID)
				imageIdentityModel.ID = core.StringPtr("72b27b5c-f4b0-48bb-b954-5becc7c1dcb8")

				// Construct an instance of the InstanceProfileIdentityByName model
				instanceProfileIdentityModel := new(vpcclassicv1.InstanceProfileIdentityByName)
				instanceProfileIdentityModel.Name = core.StringPtr("bc1-4x16")

				// Construct an instance of the KeyIdentityByID model
				keyIdentityModel := new(vpcclassicv1.KeyIdentityByID)
				keyIdentityModel.ID = core.StringPtr("a6b1a881-2ce8-41a3-80fc-36316a73f803")

				// Construct an instance of the NetworkInterfacePrototype model
				networkInterfacePrototypeModel := new(vpcclassicv1.NetworkInterfacePrototype)
				networkInterfacePrototypeModel.Name = core.StringPtr("my-network-interface")
				networkInterfacePrototypeModel.PrimaryIpv4Address = core.StringPtr("10.0.0.5")
				networkInterfacePrototypeModel.SecurityGroups = []vpcclassicv1.SecurityGroupIdentityIntf{securityGroupIdentityModel}
				networkInterfacePrototypeModel.Subnet = subnetIdentityModel

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the VPCIdentityByID model
				vpcIdentityModel := new(vpcclassicv1.VPCIdentityByID)
				vpcIdentityModel.ID = core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b")

				// Construct an instance of the VolumeAttachmentPrototypeInstanceByImageContext model
				volumeAttachmentPrototypeInstanceByImageContextModel := new(vpcclassicv1.VolumeAttachmentPrototypeInstanceByImageContext)
				volumeAttachmentPrototypeInstanceByImageContextModel.DeleteVolumeOnInstanceDelete = core.BoolPtr(true)
				volumeAttachmentPrototypeInstanceByImageContextModel.Name = core.StringPtr("my-volume-attachment")
				volumeAttachmentPrototypeInstanceByImageContextModel.Volume = volumePrototypeInstanceByImageContextModel

				// Construct an instance of the VolumeAttachmentPrototypeInstanceContext model
				volumeAttachmentPrototypeInstanceContextModel := new(vpcclassicv1.VolumeAttachmentPrototypeInstanceContext)
				volumeAttachmentPrototypeInstanceContextModel.DeleteVolumeOnInstanceDelete = core.BoolPtr(true)
				volumeAttachmentPrototypeInstanceContextModel.Name = core.StringPtr("my-volume-attachment")
				volumeAttachmentPrototypeInstanceContextModel.Volume = volumeAttachmentPrototypeInstanceContextVolumeModel

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(vpcclassicv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the InstancePrototypeInstanceByImage model
				instancePrototypeModel := new(vpcclassicv1.InstancePrototypeInstanceByImage)
				instancePrototypeModel.Keys = []vpcclassicv1.KeyIdentityIntf{keyIdentityModel}
				instancePrototypeModel.Name = core.StringPtr("my-instance")
				instancePrototypeModel.NetworkInterfaces = []vpcclassicv1.NetworkInterfacePrototype{*networkInterfacePrototypeModel}
				instancePrototypeModel.Profile = instanceProfileIdentityModel
				instancePrototypeModel.ResourceGroup = resourceGroupIdentityModel
				instancePrototypeModel.UserData = core.StringPtr("testString")
				instancePrototypeModel.VolumeAttachments = []vpcclassicv1.VolumeAttachmentPrototypeInstanceContext{*volumeAttachmentPrototypeInstanceContextModel}
				instancePrototypeModel.Vpc = vpcIdentityModel
				instancePrototypeModel.BootVolumeAttachment = volumeAttachmentPrototypeInstanceByImageContextModel
				instancePrototypeModel.Image = imageIdentityModel
				instancePrototypeModel.PrimaryNetworkInterface = networkInterfacePrototypeModel
				instancePrototypeModel.Zone = zoneIdentityModel

				// Construct an instance of the CreateInstanceOptions model
				createInstanceOptionsModel := new(vpcclassicv1.CreateInstanceOptions)
				createInstanceOptionsModel.InstancePrototype = instancePrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateInstance(createInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteInstance(deleteInstanceOptions *DeleteInstanceOptions)`, func() {
		version := "testString"
		deleteInstancePath := "/instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteInstancePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteInstanceOptions model
				deleteInstanceOptionsModel := new(vpcclassicv1.DeleteInstanceOptions)
				deleteInstanceOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteInstance(deleteInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetInstance(getInstanceOptions *GetInstanceOptions)`, func() {
		version := "testString"
		getInstancePath := "/instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getInstancePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"bandwidth": 1000, "boot_volume_attachment": {"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "image": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "name": "my-image"}, "memory": 8, "name": "my-instance", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}], "primary_network_interface": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}, "profile": {"crn": "crn:v1:bluemix:public:is:us-south-1:::instance-profile:bc1-4x16", "href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bc1-4x16", "name": "bc1-4x16"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "failed", "vcpu": {"architecture": "amd64", "count": 4}, "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(vpcclassicv1.GetInstanceOptions)
				getInstanceOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateInstance(updateInstanceOptions *UpdateInstanceOptions)`, func() {
		version := "testString"
		updateInstancePath := "/instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateInstancePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"bandwidth": 1000, "boot_volume_attachment": {"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "image": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "href": "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "id": "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8", "name": "my-image"}, "memory": 8, "name": "my-instance", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}], "primary_network_interface": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}, "profile": {"crn": "crn:v1:bluemix:public:is:us-south-1:::instance-profile:bc1-4x16", "href": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bc1-4x16", "name": "bc1-4x16"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "failed", "vcpu": {"architecture": "amd64", "count": 4}, "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke UpdateInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateInstanceOptions model
				updateInstanceOptionsModel := new(vpcclassicv1.UpdateInstanceOptions)
				updateInstanceOptionsModel.ID = core.StringPtr("testString")
				updateInstanceOptionsModel.Name = core.StringPtr("my-instance")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateInstance(updateInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetInstanceInitialization(getInstanceInitializationOptions *GetInstanceInitializationOptions)`, func() {
		version := "testString"
		getInstanceInitializationPath := "/instances/testString/initialization"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getInstanceInitializationPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"keys": [{"crn": "crn:v1:bluemix:public:is:us-south:a/123456::key:a6b1a881-2ce8-41a3-80fc-36316a73f803", "fingerprint": "SHA256:yxavE4CIOL2NlsqcurRO3xGjkP6m/0mp8ugojH5yxlY", "href": "https://us-south.iaas.cloud.ibm.com/v1/keys/a6b1a881-2ce8-41a3-80fc-36316a73f803", "id": "a6b1a881-2ce8-41a3-80fc-36316a73f803", "name": "my-key"}], "password": {"encrypted_password": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "encryption_key": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::key:a6b1a881-2ce8-41a3-80fc-36316a73f803", "fingerprint": "SHA256:yxavE4CIOL2NlsqcurRO3xGjkP6m/0mp8ugojH5yxlY", "href": "https://us-south.iaas.cloud.ibm.com/v1/keys/a6b1a881-2ce8-41a3-80fc-36316a73f803", "id": "a6b1a881-2ce8-41a3-80fc-36316a73f803", "name": "my-key"}}}`)
			}))
			It(`Invoke GetInstanceInitialization successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetInstanceInitialization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceInitializationOptions model
				getInstanceInitializationOptionsModel := new(vpcclassicv1.GetInstanceInitializationOptions)
				getInstanceInitializationOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetInstanceInitialization(getInstanceInitializationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateInstanceAction(createInstanceActionOptions *CreateInstanceActionOptions)`, func() {
		version := "testString"
		createInstanceActionPath := "/instances/testString/actions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createInstanceActionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"completed_at": "2019-01-01T12:00:00", "created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/actions/109a1b6e-1242-4de1-be44-38705e9474ed", "id": "109a1b6e-1242-4de1-be44-38705e9474ed", "started_at": "2019-01-01T12:00:00", "status": "completed", "type": "reboot"}`)
			}))
			It(`Invoke CreateInstanceAction successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateInstanceAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateInstanceActionOptions model
				createInstanceActionOptionsModel := new(vpcclassicv1.CreateInstanceActionOptions)
				createInstanceActionOptionsModel.InstanceID = core.StringPtr("testString")
				createInstanceActionOptionsModel.Type = core.StringPtr("reboot")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateInstanceAction(createInstanceActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListNetworkInterfaces(listNetworkInterfacesOptions *ListNetworkInterfacesOptions)`, func() {
		version := "testString"
		listNetworkInterfacesPath := "/instances/testString/network_interfaces"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listNetworkInterfacesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"network_interfaces": [{"created_at": "2019-01-01T12:00:00", "floating_ips": [{"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}], "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "port_speed": 1000, "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "security_groups": [{"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}], "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}, "type": "primary"}]}`)
			}))
			It(`Invoke ListNetworkInterfaces successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListNetworkInterfaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNetworkInterfacesOptions model
				listNetworkInterfacesOptionsModel := new(vpcclassicv1.ListNetworkInterfacesOptions)
				listNetworkInterfacesOptionsModel.InstanceID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListNetworkInterfaces(listNetworkInterfacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetNetworkInterface(getNetworkInterfaceOptions *GetNetworkInterfaceOptions)`, func() {
		version := "testString"
		getNetworkInterfacePath := "/instances/testString/network_interfaces/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getNetworkInterfacePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "floating_ips": [{"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}], "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "port_speed": 1000, "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "security_groups": [{"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}], "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}, "type": "primary"}`)
			}))
			It(`Invoke GetNetworkInterface successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetNetworkInterface(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNetworkInterfaceOptions model
				getNetworkInterfaceOptionsModel := new(vpcclassicv1.GetNetworkInterfaceOptions)
				getNetworkInterfaceOptionsModel.InstanceID = core.StringPtr("testString")
				getNetworkInterfaceOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetNetworkInterface(getNetworkInterfaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListNetworkInterfaceFloatingIps(listNetworkInterfaceFloatingIpsOptions *ListNetworkInterfaceFloatingIpsOptions)`, func() {
		version := "testString"
		listNetworkInterfaceFloatingIpsPath := "/instances/testString/network_interfaces/testString/floating_ips"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listNetworkInterfaceFloatingIpsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"floating_ips": [{"address": "203.0.113.1", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip", "status": "available", "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}]}`)
			}))
			It(`Invoke ListNetworkInterfaceFloatingIps successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListNetworkInterfaceFloatingIps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNetworkInterfaceFloatingIpsOptions model
				listNetworkInterfaceFloatingIpsOptionsModel := new(vpcclassicv1.ListNetworkInterfaceFloatingIpsOptions)
				listNetworkInterfaceFloatingIpsOptionsModel.InstanceID = core.StringPtr("testString")
				listNetworkInterfaceFloatingIpsOptionsModel.NetworkInterfaceID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListNetworkInterfaceFloatingIps(listNetworkInterfaceFloatingIpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteNetworkInterfaceFloatingIpBinding(deleteNetworkInterfaceFloatingIpBindingOptions *DeleteNetworkInterfaceFloatingIpBindingOptions)`, func() {
		version := "testString"
		deleteNetworkInterfaceFloatingIpBindingPath := "/instances/testString/network_interfaces/testString/floating_ips/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteNetworkInterfaceFloatingIpBindingPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteNetworkInterfaceFloatingIpBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteNetworkInterfaceFloatingIpBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNetworkInterfaceFloatingIpBindingOptions model
				deleteNetworkInterfaceFloatingIpBindingOptionsModel := new(vpcclassicv1.DeleteNetworkInterfaceFloatingIpBindingOptions)
				deleteNetworkInterfaceFloatingIpBindingOptionsModel.InstanceID = core.StringPtr("testString")
				deleteNetworkInterfaceFloatingIpBindingOptionsModel.NetworkInterfaceID = core.StringPtr("testString")
				deleteNetworkInterfaceFloatingIpBindingOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteNetworkInterfaceFloatingIpBinding(deleteNetworkInterfaceFloatingIpBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetNetworkInterfaceFloatingIp(getNetworkInterfaceFloatingIpOptions *GetNetworkInterfaceFloatingIpOptions)`, func() {
		version := "testString"
		getNetworkInterfaceFloatingIpPath := "/instances/testString/network_interfaces/testString/floating_ips/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getNetworkInterfaceFloatingIpPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"address": "203.0.113.1", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip", "status": "available", "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetNetworkInterfaceFloatingIp successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetNetworkInterfaceFloatingIp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNetworkInterfaceFloatingIpOptions model
				getNetworkInterfaceFloatingIpOptionsModel := new(vpcclassicv1.GetNetworkInterfaceFloatingIpOptions)
				getNetworkInterfaceFloatingIpOptionsModel.InstanceID = core.StringPtr("testString")
				getNetworkInterfaceFloatingIpOptionsModel.NetworkInterfaceID = core.StringPtr("testString")
				getNetworkInterfaceFloatingIpOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetNetworkInterfaceFloatingIp(getNetworkInterfaceFloatingIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateNetworkInterfaceFloatingIpBinding(createNetworkInterfaceFloatingIpBindingOptions *CreateNetworkInterfaceFloatingIpBindingOptions)`, func() {
		version := "testString"
		createNetworkInterfaceFloatingIpBindingPath := "/instances/testString/network_interfaces/testString/floating_ips/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createNetworkInterfaceFloatingIpBindingPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"address": "203.0.113.1", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip", "status": "available", "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke CreateNetworkInterfaceFloatingIpBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateNetworkInterfaceFloatingIpBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateNetworkInterfaceFloatingIpBindingOptions model
				createNetworkInterfaceFloatingIpBindingOptionsModel := new(vpcclassicv1.CreateNetworkInterfaceFloatingIpBindingOptions)
				createNetworkInterfaceFloatingIpBindingOptionsModel.InstanceID = core.StringPtr("testString")
				createNetworkInterfaceFloatingIpBindingOptionsModel.NetworkInterfaceID = core.StringPtr("testString")
				createNetworkInterfaceFloatingIpBindingOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateNetworkInterfaceFloatingIpBinding(createNetworkInterfaceFloatingIpBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVolumeAttachments(listVolumeAttachmentsOptions *ListVolumeAttachmentsOptions)`, func() {
		version := "testString"
		listVolumeAttachmentsPath := "/instances/testString/volume_attachments"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVolumeAttachmentsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"volume_attachments": [{"created_at": "2019-01-01T12:00:00", "delete_volume_on_instance_delete": true, "device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "status": "attached", "type": "boot", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}]}`)
			}))
			It(`Invoke ListVolumeAttachments successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVolumeAttachments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVolumeAttachmentsOptions model
				listVolumeAttachmentsOptionsModel := new(vpcclassicv1.ListVolumeAttachmentsOptions)
				listVolumeAttachmentsOptionsModel.InstanceID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVolumeAttachments(listVolumeAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVolumeAttachment(createVolumeAttachmentOptions *CreateVolumeAttachmentOptions)`, func() {
		version := "testString"
		createVolumeAttachmentPath := "/instances/testString/volume_attachments"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVolumeAttachmentPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "delete_volume_on_instance_delete": true, "device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "status": "attached", "type": "boot", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}`)
			}))
			It(`Invoke CreateVolumeAttachment successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateVolumeAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VolumeIdentityByID model
				volumeIdentityModel := new(vpcclassicv1.VolumeIdentityByID)
				volumeIdentityModel.ID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5")

				// Construct an instance of the CreateVolumeAttachmentOptions model
				createVolumeAttachmentOptionsModel := new(vpcclassicv1.CreateVolumeAttachmentOptions)
				createVolumeAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				createVolumeAttachmentOptionsModel.Volume = volumeIdentityModel
				createVolumeAttachmentOptionsModel.DeleteVolumeOnInstanceDelete = core.BoolPtr(true)
				createVolumeAttachmentOptionsModel.Name = core.StringPtr("my-volume-attachment")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateVolumeAttachment(createVolumeAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVolumeAttachment(deleteVolumeAttachmentOptions *DeleteVolumeAttachmentOptions)`, func() {
		version := "testString"
		deleteVolumeAttachmentPath := "/instances/testString/volume_attachments/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVolumeAttachmentPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteVolumeAttachment successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVolumeAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVolumeAttachmentOptions model
				deleteVolumeAttachmentOptionsModel := new(vpcclassicv1.DeleteVolumeAttachmentOptions)
				deleteVolumeAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				deleteVolumeAttachmentOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVolumeAttachment(deleteVolumeAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVolumeAttachment(getVolumeAttachmentOptions *GetVolumeAttachmentOptions)`, func() {
		version := "testString"
		getVolumeAttachmentPath := "/instances/testString/volume_attachments/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVolumeAttachmentPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "delete_volume_on_instance_delete": true, "device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "status": "attached", "type": "boot", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}`)
			}))
			It(`Invoke GetVolumeAttachment successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVolumeAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVolumeAttachmentOptions model
				getVolumeAttachmentOptionsModel := new(vpcclassicv1.GetVolumeAttachmentOptions)
				getVolumeAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				getVolumeAttachmentOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVolumeAttachment(getVolumeAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVolumeAttachment(updateVolumeAttachmentOptions *UpdateVolumeAttachmentOptions)`, func() {
		version := "testString"
		updateVolumeAttachmentPath := "/instances/testString/volume_attachments/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVolumeAttachmentPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "delete_volume_on_instance_delete": true, "device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "name": "my-volume-attachment", "status": "attached", "type": "boot", "volume": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "name": "my-volume"}}`)
			}))
			It(`Invoke UpdateVolumeAttachment successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateVolumeAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVolumeAttachmentOptions model
				updateVolumeAttachmentOptionsModel := new(vpcclassicv1.UpdateVolumeAttachmentOptions)
				updateVolumeAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				updateVolumeAttachmentOptionsModel.ID = core.StringPtr("testString")
				updateVolumeAttachmentOptionsModel.DeleteVolumeOnInstanceDelete = core.BoolPtr(true)
				updateVolumeAttachmentOptionsModel.Name = core.StringPtr("my-volume-attachment")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateVolumeAttachment(updateVolumeAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions)`, func() {
		version := "testString"
		listLoadBalancersPath := "/load_balancers"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLoadBalancersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"load_balancers": [{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::load-balancer:dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "hostname": "myloadbalancer-123456-us-south-1.lb.bluemix.net", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "id": "dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "is_public": true, "listeners": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer", "operating_status": "offline", "pools": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}], "private_ips": [{"address": "192.168.3.4"}], "provisioning_status": "active", "public_ips": [{"address": "192.168.3.4"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}]}`)
			}))
			It(`Invoke ListLoadBalancers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLoadBalancers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancersOptions model
				listLoadBalancersOptionsModel := new(vpcclassicv1.ListLoadBalancersOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions)`, func() {
		version := "testString"
		createLoadBalancerPath := "/load_balancers"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createLoadBalancerPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::load-balancer:dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "hostname": "myloadbalancer-123456-us-south-1.lb.bluemix.net", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "id": "dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "is_public": true, "listeners": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer", "operating_status": "offline", "pools": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}], "private_ips": [{"address": "192.168.3.4"}], "provisioning_status": "active", "public_ips": [{"address": "192.168.3.4"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke CreateLoadBalancer successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolMemberTargetPrototypeByAddress model
				loadBalancerPoolMemberTargetPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberTargetPrototypeByAddress)
				loadBalancerPoolMemberTargetPrototypeModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the LoadBalancerPoolHealthMonitorPrototype model
				loadBalancerPoolHealthMonitorPrototypeModel := new(vpcclassicv1.LoadBalancerPoolHealthMonitorPrototype)
				loadBalancerPoolHealthMonitorPrototypeModel.Delay = core.Int64Ptr(int64(5))
				loadBalancerPoolHealthMonitorPrototypeModel.MaxRetries = core.Int64Ptr(int64(2))
				loadBalancerPoolHealthMonitorPrototypeModel.Port = core.Int64Ptr(int64(22))
				loadBalancerPoolHealthMonitorPrototypeModel.Timeout = core.Int64Ptr(int64(2))
				loadBalancerPoolHealthMonitorPrototypeModel.Type = core.StringPtr("http")
				loadBalancerPoolHealthMonitorPrototypeModel.UrlPath = core.StringPtr("/")

				// Construct an instance of the LoadBalancerPoolIdentityByName model
				loadBalancerPoolIdentityByNameModel := new(vpcclassicv1.LoadBalancerPoolIdentityByName)
				loadBalancerPoolIdentityByNameModel.Name = core.StringPtr("my-load-balancer-pool")

				// Construct an instance of the LoadBalancerPoolMemberPrototype model
				loadBalancerPoolMemberPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberPrototype)
				loadBalancerPoolMemberPrototypeModel.Port = core.Int64Ptr(int64(80))
				loadBalancerPoolMemberPrototypeModel.Target = loadBalancerPoolMemberTargetPrototypeModel
				loadBalancerPoolMemberPrototypeModel.Weight = core.Int64Ptr(int64(50))

				// Construct an instance of the LoadBalancerPoolSessionPersistencePrototype model
				loadBalancerPoolSessionPersistencePrototypeModel := new(vpcclassicv1.LoadBalancerPoolSessionPersistencePrototype)
				loadBalancerPoolSessionPersistencePrototypeModel.Type = core.StringPtr("source_ip")

				// Construct an instance of the LoadBalancerListenerPrototypeLoadBalancerContext model
				loadBalancerListenerPrototypeLoadBalancerContextModel := new(vpcclassicv1.LoadBalancerListenerPrototypeLoadBalancerContext)
				loadBalancerListenerPrototypeLoadBalancerContextModel.ConnectionLimit = core.Int64Ptr(int64(2000))
				loadBalancerListenerPrototypeLoadBalancerContextModel.DefaultPool = loadBalancerPoolIdentityByNameModel
				loadBalancerListenerPrototypeLoadBalancerContextModel.Port = core.Int64Ptr(int64(443))
				loadBalancerListenerPrototypeLoadBalancerContextModel.Protocol = core.StringPtr("http")

				// Construct an instance of the LoadBalancerPoolPrototype model
				loadBalancerPoolPrototypeModel := new(vpcclassicv1.LoadBalancerPoolPrototype)
				loadBalancerPoolPrototypeModel.Algorithm = core.StringPtr("least_connections")
				loadBalancerPoolPrototypeModel.HealthMonitor = loadBalancerPoolHealthMonitorPrototypeModel
				loadBalancerPoolPrototypeModel.Members = []vpcclassicv1.LoadBalancerPoolMemberPrototype{*loadBalancerPoolMemberPrototypeModel}
				loadBalancerPoolPrototypeModel.Name = core.StringPtr("my-load-balancer-pool")
				loadBalancerPoolPrototypeModel.Protocol = core.StringPtr("http")
				loadBalancerPoolPrototypeModel.SessionPersistence = loadBalancerPoolSessionPersistencePrototypeModel

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the SubnetIdentityByID model
				subnetIdentityModel := new(vpcclassicv1.SubnetIdentityByID)
				subnetIdentityModel.ID = core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(vpcclassicv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.IsPublic = core.BoolPtr(true)
				createLoadBalancerOptionsModel.Subnets = []vpcclassicv1.SubnetIdentityIntf{subnetIdentityModel}
				createLoadBalancerOptionsModel.Listeners = []vpcclassicv1.LoadBalancerListenerPrototypeLoadBalancerContext{*loadBalancerListenerPrototypeLoadBalancerContextModel}
				createLoadBalancerOptionsModel.Name = core.StringPtr("my-load-balancer")
				createLoadBalancerOptionsModel.Pools = []vpcclassicv1.LoadBalancerPoolPrototype{*loadBalancerPoolPrototypeModel}
				createLoadBalancerOptionsModel.ResourceGroup = resourceGroupIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions)`, func() {
		version := "testString"
		deleteLoadBalancerPath := "/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLoadBalancerPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteLoadBalancer successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(vpcclassicv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions)`, func() {
		version := "testString"
		getLoadBalancerPath := "/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoadBalancerPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::load-balancer:dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "hostname": "myloadbalancer-123456-us-south-1.lb.bluemix.net", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "id": "dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "is_public": true, "listeners": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer", "operating_status": "offline", "pools": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}], "private_ips": [{"address": "192.168.3.4"}], "provisioning_status": "active", "public_ips": [{"address": "192.168.3.4"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke GetLoadBalancer successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerOptions model
				getLoadBalancerOptionsModel := new(vpcclassicv1.GetLoadBalancerOptions)
				getLoadBalancerOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions)`, func() {
		version := "testString"
		updateLoadBalancerPath := "/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateLoadBalancerPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::load-balancer:dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "hostname": "myloadbalancer-123456-us-south-1.lb.bluemix.net", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "id": "dd754295-e9e0-4c9d-bf6c-58fbc59e5727", "is_public": true, "listeners": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer", "operating_status": "offline", "pools": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}], "private_ips": [{"address": "192.168.3.4"}], "provisioning_status": "active", "public_ips": [{"address": "192.168.3.4"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke UpdateLoadBalancer successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateLoadBalancerOptions model
				updateLoadBalancerOptionsModel := new(vpcclassicv1.UpdateLoadBalancerOptions)
				updateLoadBalancerOptionsModel.ID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Name = core.StringPtr("my-load-balancer")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoadBalancerStatistics(getLoadBalancerStatisticsOptions *GetLoadBalancerStatisticsOptions)`, func() {
		version := "testString"
		getLoadBalancerStatisticsPath := "/load_balancers/testString/statistics"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoadBalancerStatisticsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"active_connections": 797, "connection_rate": 91.121, "data_processed_this_month": 10093173145, "throughput": 167.278}`)
			}))
			It(`Invoke GetLoadBalancerStatistics successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerStatistics(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerStatisticsOptions model
				getLoadBalancerStatisticsOptionsModel := new(vpcclassicv1.GetLoadBalancerStatisticsOptions)
				getLoadBalancerStatisticsOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerStatistics(getLoadBalancerStatisticsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLoadBalancerListeners(listLoadBalancerListenersOptions *ListLoadBalancerListenersOptions)`, func() {
		version := "testString"
		listLoadBalancerListenersPath := "/load_balancers/testString/listeners"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLoadBalancerListenersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"listeners": [{"certificate_instance": {"crn": "crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758"}, "connection_limit": 2000, "created_at": "2019-01-01T12:00:00", "default_pool": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "policies": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "port": 443, "protocol": "http", "provisioning_status": "active"}]}`)
			}))
			It(`Invoke ListLoadBalancerListeners successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLoadBalancerListeners(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancerListenersOptions model
				listLoadBalancerListenersOptionsModel := new(vpcclassicv1.ListLoadBalancerListenersOptions)
				listLoadBalancerListenersOptionsModel.LoadBalancerID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLoadBalancerListeners(listLoadBalancerListenersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateLoadBalancerListener(createLoadBalancerListenerOptions *CreateLoadBalancerListenerOptions)`, func() {
		version := "testString"
		createLoadBalancerListenerPath := "/load_balancers/testString/listeners"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createLoadBalancerListenerPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"certificate_instance": {"crn": "crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758"}, "connection_limit": 2000, "created_at": "2019-01-01T12:00:00", "default_pool": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "policies": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "port": 443, "protocol": "http", "provisioning_status": "active"}`)
			}))
			It(`Invoke CreateLoadBalancerListener successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateLoadBalancerListener(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID model
				loadBalancerListenerPolicyPrototypeTargetModel := new(vpcclassicv1.LoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID)
				loadBalancerListenerPolicyPrototypeTargetModel.ID = core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004")

				// Construct an instance of the LoadBalancerListenerPolicyRulePrototype model
				loadBalancerListenerPolicyRulePrototypeModel := new(vpcclassicv1.LoadBalancerListenerPolicyRulePrototype)
				loadBalancerListenerPolicyRulePrototypeModel.Condition = core.StringPtr("contains")
				loadBalancerListenerPolicyRulePrototypeModel.Field = core.StringPtr("MY-APP-HEADER")
				loadBalancerListenerPolicyRulePrototypeModel.Type = core.StringPtr("header")
				loadBalancerListenerPolicyRulePrototypeModel.Value = core.StringPtr("testString")

				// Construct an instance of the CertificateInstanceIdentityByCRN model
				certificateInstanceIdentityModel := new(vpcclassicv1.CertificateInstanceIdentityByCRN)
				certificateInstanceIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758")

				// Construct an instance of the LoadBalancerListenerPolicyPrototype model
				loadBalancerListenerPolicyPrototypeModel := new(vpcclassicv1.LoadBalancerListenerPolicyPrototype)
				loadBalancerListenerPolicyPrototypeModel.Action = core.StringPtr("forward")
				loadBalancerListenerPolicyPrototypeModel.Name = core.StringPtr("my-policy")
				loadBalancerListenerPolicyPrototypeModel.Priority = core.Int64Ptr(int64(5))
				loadBalancerListenerPolicyPrototypeModel.Rules = []vpcclassicv1.LoadBalancerListenerPolicyRulePrototype{*loadBalancerListenerPolicyRulePrototypeModel}
				loadBalancerListenerPolicyPrototypeModel.Target = loadBalancerListenerPolicyPrototypeTargetModel

				// Construct an instance of the LoadBalancerPoolIdentityByID model
				loadBalancerPoolIdentityModel := new(vpcclassicv1.LoadBalancerPoolIdentityByID)
				loadBalancerPoolIdentityModel.ID = core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004")

				// Construct an instance of the CreateLoadBalancerListenerOptions model
				createLoadBalancerListenerOptionsModel := new(vpcclassicv1.CreateLoadBalancerListenerOptions)
				createLoadBalancerListenerOptionsModel.LoadBalancerID = core.StringPtr("testString")
				createLoadBalancerListenerOptionsModel.Port = core.Int64Ptr(int64(443))
				createLoadBalancerListenerOptionsModel.Protocol = core.StringPtr("http")
				createLoadBalancerListenerOptionsModel.CertificateInstance = certificateInstanceIdentityModel
				createLoadBalancerListenerOptionsModel.ConnectionLimit = core.Int64Ptr(int64(2000))
				createLoadBalancerListenerOptionsModel.DefaultPool = loadBalancerPoolIdentityModel
				createLoadBalancerListenerOptionsModel.Policies = []vpcclassicv1.LoadBalancerListenerPolicyPrototype{*loadBalancerListenerPolicyPrototypeModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateLoadBalancerListener(createLoadBalancerListenerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLoadBalancerListener(deleteLoadBalancerListenerOptions *DeleteLoadBalancerListenerOptions)`, func() {
		version := "testString"
		deleteLoadBalancerListenerPath := "/load_balancers/testString/listeners/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLoadBalancerListenerPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteLoadBalancerListener successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteLoadBalancerListener(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerListenerOptions model
				deleteLoadBalancerListenerOptionsModel := new(vpcclassicv1.DeleteLoadBalancerListenerOptions)
				deleteLoadBalancerListenerOptionsModel.LoadBalancerID = core.StringPtr("testString")
				deleteLoadBalancerListenerOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteLoadBalancerListener(deleteLoadBalancerListenerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoadBalancerListener(getLoadBalancerListenerOptions *GetLoadBalancerListenerOptions)`, func() {
		version := "testString"
		getLoadBalancerListenerPath := "/load_balancers/testString/listeners/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoadBalancerListenerPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"certificate_instance": {"crn": "crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758"}, "connection_limit": 2000, "created_at": "2019-01-01T12:00:00", "default_pool": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "policies": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "port": 443, "protocol": "http", "provisioning_status": "active"}`)
			}))
			It(`Invoke GetLoadBalancerListener successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerListener(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerListenerOptions model
				getLoadBalancerListenerOptionsModel := new(vpcclassicv1.GetLoadBalancerListenerOptions)
				getLoadBalancerListenerOptionsModel.LoadBalancerID = core.StringPtr("testString")
				getLoadBalancerListenerOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerListener(getLoadBalancerListenerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateLoadBalancerListener(updateLoadBalancerListenerOptions *UpdateLoadBalancerListenerOptions)`, func() {
		version := "testString"
		updateLoadBalancerListenerPath := "/load_balancers/testString/listeners/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateLoadBalancerListenerPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"certificate_instance": {"crn": "crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758"}, "connection_limit": 2000, "created_at": "2019-01-01T12:00:00", "default_pool": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "policies": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "port": 443, "protocol": "http", "provisioning_status": "active"}`)
			}))
			It(`Invoke UpdateLoadBalancerListener successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateLoadBalancerListener(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CertificateInstanceIdentityByCRN model
				certificateInstanceIdentityModel := new(vpcclassicv1.CertificateInstanceIdentityByCRN)
				certificateInstanceIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758")

				// Construct an instance of the LoadBalancerPoolIdentityByID model
				loadBalancerPoolIdentityModel := new(vpcclassicv1.LoadBalancerPoolIdentityByID)
				loadBalancerPoolIdentityModel.ID = core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004")

				// Construct an instance of the UpdateLoadBalancerListenerOptions model
				updateLoadBalancerListenerOptionsModel := new(vpcclassicv1.UpdateLoadBalancerListenerOptions)
				updateLoadBalancerListenerOptionsModel.LoadBalancerID = core.StringPtr("testString")
				updateLoadBalancerListenerOptionsModel.ID = core.StringPtr("testString")
				updateLoadBalancerListenerOptionsModel.CertificateInstance = certificateInstanceIdentityModel
				updateLoadBalancerListenerOptionsModel.ConnectionLimit = core.Int64Ptr(int64(2000))
				updateLoadBalancerListenerOptionsModel.DefaultPool = loadBalancerPoolIdentityModel
				updateLoadBalancerListenerOptionsModel.Port = core.Int64Ptr(int64(443))
				updateLoadBalancerListenerOptionsModel.Protocol = core.StringPtr("http")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateLoadBalancerListener(updateLoadBalancerListenerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLoadBalancerListenerPolicies(listLoadBalancerListenerPoliciesOptions *ListLoadBalancerListenerPoliciesOptions)`, func() {
		version := "testString"
		listLoadBalancerListenerPoliciesPath := "/load_balancers/testString/listeners/testString/policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLoadBalancerListenerPoliciesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"policies": [{"action": "forward", "created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-policy", "priority": 5, "provisioning_status": "active", "rules": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}}]}`)
			}))
			It(`Invoke ListLoadBalancerListenerPolicies successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLoadBalancerListenerPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancerListenerPoliciesOptions model
				listLoadBalancerListenerPoliciesOptionsModel := new(vpcclassicv1.ListLoadBalancerListenerPoliciesOptions)
				listLoadBalancerListenerPoliciesOptionsModel.LoadBalancerID = core.StringPtr("testString")
				listLoadBalancerListenerPoliciesOptionsModel.ListenerID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLoadBalancerListenerPolicies(listLoadBalancerListenerPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateLoadBalancerListenerPolicy(createLoadBalancerListenerPolicyOptions *CreateLoadBalancerListenerPolicyOptions)`, func() {
		version := "testString"
		createLoadBalancerListenerPolicyPath := "/load_balancers/testString/listeners/testString/policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createLoadBalancerListenerPolicyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"action": "forward", "created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-policy", "priority": 5, "provisioning_status": "active", "rules": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}}`)
			}))
			It(`Invoke CreateLoadBalancerListenerPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateLoadBalancerListenerPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID model
				loadBalancerListenerPolicyPrototypeTargetModel := new(vpcclassicv1.LoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID)
				loadBalancerListenerPolicyPrototypeTargetModel.ID = core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004")

				// Construct an instance of the LoadBalancerListenerPolicyRulePrototype model
				loadBalancerListenerPolicyRulePrototypeModel := new(vpcclassicv1.LoadBalancerListenerPolicyRulePrototype)
				loadBalancerListenerPolicyRulePrototypeModel.Condition = core.StringPtr("contains")
				loadBalancerListenerPolicyRulePrototypeModel.Field = core.StringPtr("MY-APP-HEADER")
				loadBalancerListenerPolicyRulePrototypeModel.Type = core.StringPtr("header")
				loadBalancerListenerPolicyRulePrototypeModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateLoadBalancerListenerPolicyOptions model
				createLoadBalancerListenerPolicyOptionsModel := new(vpcclassicv1.CreateLoadBalancerListenerPolicyOptions)
				createLoadBalancerListenerPolicyOptionsModel.LoadBalancerID = core.StringPtr("testString")
				createLoadBalancerListenerPolicyOptionsModel.ListenerID = core.StringPtr("testString")
				createLoadBalancerListenerPolicyOptionsModel.Action = core.StringPtr("forward")
				createLoadBalancerListenerPolicyOptionsModel.Priority = core.Int64Ptr(int64(5))
				createLoadBalancerListenerPolicyOptionsModel.Name = core.StringPtr("my-policy")
				createLoadBalancerListenerPolicyOptionsModel.Rules = []vpcclassicv1.LoadBalancerListenerPolicyRulePrototype{*loadBalancerListenerPolicyRulePrototypeModel}
				createLoadBalancerListenerPolicyOptionsModel.Target = loadBalancerListenerPolicyPrototypeTargetModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateLoadBalancerListenerPolicy(createLoadBalancerListenerPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLoadBalancerListenerPolicy(deleteLoadBalancerListenerPolicyOptions *DeleteLoadBalancerListenerPolicyOptions)`, func() {
		version := "testString"
		deleteLoadBalancerListenerPolicyPath := "/load_balancers/testString/listeners/testString/policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLoadBalancerListenerPolicyPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteLoadBalancerListenerPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteLoadBalancerListenerPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerListenerPolicyOptions model
				deleteLoadBalancerListenerPolicyOptionsModel := new(vpcclassicv1.DeleteLoadBalancerListenerPolicyOptions)
				deleteLoadBalancerListenerPolicyOptionsModel.LoadBalancerID = core.StringPtr("testString")
				deleteLoadBalancerListenerPolicyOptionsModel.ListenerID = core.StringPtr("testString")
				deleteLoadBalancerListenerPolicyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteLoadBalancerListenerPolicy(deleteLoadBalancerListenerPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoadBalancerListenerPolicy(getLoadBalancerListenerPolicyOptions *GetLoadBalancerListenerPolicyOptions)`, func() {
		version := "testString"
		getLoadBalancerListenerPolicyPath := "/load_balancers/testString/listeners/testString/policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoadBalancerListenerPolicyPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"action": "forward", "created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-policy", "priority": 5, "provisioning_status": "active", "rules": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}}`)
			}))
			It(`Invoke GetLoadBalancerListenerPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerListenerPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerListenerPolicyOptions model
				getLoadBalancerListenerPolicyOptionsModel := new(vpcclassicv1.GetLoadBalancerListenerPolicyOptions)
				getLoadBalancerListenerPolicyOptionsModel.LoadBalancerID = core.StringPtr("testString")
				getLoadBalancerListenerPolicyOptionsModel.ListenerID = core.StringPtr("testString")
				getLoadBalancerListenerPolicyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerListenerPolicy(getLoadBalancerListenerPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateLoadBalancerListenerPolicy(updateLoadBalancerListenerPolicyOptions *UpdateLoadBalancerListenerPolicyOptions)`, func() {
		version := "testString"
		updateLoadBalancerListenerPolicyPath := "/load_balancers/testString/listeners/testString/policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateLoadBalancerListenerPolicyPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"action": "forward", "created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-policy", "priority": 5, "provisioning_status": "active", "rules": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "target": {"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "name": "my-load-balancer-pool"}}`)
			}))
			It(`Invoke UpdateLoadBalancerListenerPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateLoadBalancerListenerPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID model
				loadBalancerListenerPolicyPatchTargetModel := new(vpcclassicv1.LoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID)
				loadBalancerListenerPolicyPatchTargetModel.ID = core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004")

				// Construct an instance of the UpdateLoadBalancerListenerPolicyOptions model
				updateLoadBalancerListenerPolicyOptionsModel := new(vpcclassicv1.UpdateLoadBalancerListenerPolicyOptions)
				updateLoadBalancerListenerPolicyOptionsModel.LoadBalancerID = core.StringPtr("testString")
				updateLoadBalancerListenerPolicyOptionsModel.ListenerID = core.StringPtr("testString")
				updateLoadBalancerListenerPolicyOptionsModel.ID = core.StringPtr("testString")
				updateLoadBalancerListenerPolicyOptionsModel.Name = core.StringPtr("my-policy")
				updateLoadBalancerListenerPolicyOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateLoadBalancerListenerPolicyOptionsModel.Target = loadBalancerListenerPolicyPatchTargetModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateLoadBalancerListenerPolicy(updateLoadBalancerListenerPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLoadBalancerListenerPolicyRules(listLoadBalancerListenerPolicyRulesOptions *ListLoadBalancerListenerPolicyRulesOptions)`, func() {
		version := "testString"
		listLoadBalancerListenerPolicyRulesPath := "/load_balancers/testString/listeners/testString/policies/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLoadBalancerListenerPolicyRulesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rules": [{"condition": "contains", "created_at": "2019-01-01T12:00:00", "field": "MY-APP-HEADER", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "provisioning_status": "active", "type": "header", "value": "Value"}]}`)
			}))
			It(`Invoke ListLoadBalancerListenerPolicyRules successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLoadBalancerListenerPolicyRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancerListenerPolicyRulesOptions model
				listLoadBalancerListenerPolicyRulesOptionsModel := new(vpcclassicv1.ListLoadBalancerListenerPolicyRulesOptions)
				listLoadBalancerListenerPolicyRulesOptionsModel.LoadBalancerID = core.StringPtr("testString")
				listLoadBalancerListenerPolicyRulesOptionsModel.ListenerID = core.StringPtr("testString")
				listLoadBalancerListenerPolicyRulesOptionsModel.PolicyID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLoadBalancerListenerPolicyRules(listLoadBalancerListenerPolicyRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateLoadBalancerListenerPolicyRule(createLoadBalancerListenerPolicyRuleOptions *CreateLoadBalancerListenerPolicyRuleOptions)`, func() {
		version := "testString"
		createLoadBalancerListenerPolicyRulePath := "/load_balancers/testString/listeners/testString/policies/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createLoadBalancerListenerPolicyRulePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"condition": "contains", "created_at": "2019-01-01T12:00:00", "field": "MY-APP-HEADER", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "provisioning_status": "active", "type": "header", "value": "Value"}`)
			}))
			It(`Invoke CreateLoadBalancerListenerPolicyRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateLoadBalancerListenerPolicyRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLoadBalancerListenerPolicyRuleOptions model
				createLoadBalancerListenerPolicyRuleOptionsModel := new(vpcclassicv1.CreateLoadBalancerListenerPolicyRuleOptions)
				createLoadBalancerListenerPolicyRuleOptionsModel.LoadBalancerID = core.StringPtr("testString")
				createLoadBalancerListenerPolicyRuleOptionsModel.ListenerID = core.StringPtr("testString")
				createLoadBalancerListenerPolicyRuleOptionsModel.PolicyID = core.StringPtr("testString")
				createLoadBalancerListenerPolicyRuleOptionsModel.Condition = core.StringPtr("contains")
				createLoadBalancerListenerPolicyRuleOptionsModel.Type = core.StringPtr("header")
				createLoadBalancerListenerPolicyRuleOptionsModel.Value = core.StringPtr("testString")
				createLoadBalancerListenerPolicyRuleOptionsModel.Field = core.StringPtr("MY-APP-HEADER")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateLoadBalancerListenerPolicyRule(createLoadBalancerListenerPolicyRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLoadBalancerListenerPolicyRule(deleteLoadBalancerListenerPolicyRuleOptions *DeleteLoadBalancerListenerPolicyRuleOptions)`, func() {
		version := "testString"
		deleteLoadBalancerListenerPolicyRulePath := "/load_balancers/testString/listeners/testString/policies/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLoadBalancerListenerPolicyRulePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteLoadBalancerListenerPolicyRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteLoadBalancerListenerPolicyRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerListenerPolicyRuleOptions model
				deleteLoadBalancerListenerPolicyRuleOptionsModel := new(vpcclassicv1.DeleteLoadBalancerListenerPolicyRuleOptions)
				deleteLoadBalancerListenerPolicyRuleOptionsModel.LoadBalancerID = core.StringPtr("testString")
				deleteLoadBalancerListenerPolicyRuleOptionsModel.ListenerID = core.StringPtr("testString")
				deleteLoadBalancerListenerPolicyRuleOptionsModel.PolicyID = core.StringPtr("testString")
				deleteLoadBalancerListenerPolicyRuleOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteLoadBalancerListenerPolicyRule(deleteLoadBalancerListenerPolicyRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoadBalancerListenerPolicyRule(getLoadBalancerListenerPolicyRuleOptions *GetLoadBalancerListenerPolicyRuleOptions)`, func() {
		version := "testString"
		getLoadBalancerListenerPolicyRulePath := "/load_balancers/testString/listeners/testString/policies/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoadBalancerListenerPolicyRulePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"condition": "contains", "created_at": "2019-01-01T12:00:00", "field": "MY-APP-HEADER", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "provisioning_status": "active", "type": "header", "value": "Value"}`)
			}))
			It(`Invoke GetLoadBalancerListenerPolicyRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerListenerPolicyRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerListenerPolicyRuleOptions model
				getLoadBalancerListenerPolicyRuleOptionsModel := new(vpcclassicv1.GetLoadBalancerListenerPolicyRuleOptions)
				getLoadBalancerListenerPolicyRuleOptionsModel.LoadBalancerID = core.StringPtr("testString")
				getLoadBalancerListenerPolicyRuleOptionsModel.ListenerID = core.StringPtr("testString")
				getLoadBalancerListenerPolicyRuleOptionsModel.PolicyID = core.StringPtr("testString")
				getLoadBalancerListenerPolicyRuleOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerListenerPolicyRule(getLoadBalancerListenerPolicyRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateLoadBalancerListenerPolicyRule(updateLoadBalancerListenerPolicyRuleOptions *UpdateLoadBalancerListenerPolicyRuleOptions)`, func() {
		version := "testString"
		updateLoadBalancerListenerPolicyRulePath := "/load_balancers/testString/listeners/testString/policies/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateLoadBalancerListenerPolicyRulePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"condition": "contains", "created_at": "2019-01-01T12:00:00", "field": "MY-APP-HEADER", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/listeners/70294e14-4e61-11e8-bcf4-0242ac110004/policies/f3187486-7b27-4c79-990c-47d33c0e2278/rules/873a84b0-84d6-49c6-8948-1fa527b25762", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "provisioning_status": "active", "type": "header", "value": "Value"}`)
			}))
			It(`Invoke UpdateLoadBalancerListenerPolicyRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateLoadBalancerListenerPolicyRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateLoadBalancerListenerPolicyRuleOptions model
				updateLoadBalancerListenerPolicyRuleOptionsModel := new(vpcclassicv1.UpdateLoadBalancerListenerPolicyRuleOptions)
				updateLoadBalancerListenerPolicyRuleOptionsModel.LoadBalancerID = core.StringPtr("testString")
				updateLoadBalancerListenerPolicyRuleOptionsModel.ListenerID = core.StringPtr("testString")
				updateLoadBalancerListenerPolicyRuleOptionsModel.PolicyID = core.StringPtr("testString")
				updateLoadBalancerListenerPolicyRuleOptionsModel.ID = core.StringPtr("testString")
				updateLoadBalancerListenerPolicyRuleOptionsModel.Condition = core.StringPtr("contains")
				updateLoadBalancerListenerPolicyRuleOptionsModel.Field = core.StringPtr("MY-APP-HEADER")
				updateLoadBalancerListenerPolicyRuleOptionsModel.Type = core.StringPtr("header")
				updateLoadBalancerListenerPolicyRuleOptionsModel.Value = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateLoadBalancerListenerPolicyRule(updateLoadBalancerListenerPolicyRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLoadBalancerPools(listLoadBalancerPoolsOptions *ListLoadBalancerPoolsOptions)`, func() {
		version := "testString"
		listLoadBalancerPoolsPath := "/load_balancers/testString/pools"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLoadBalancerPoolsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"pools": [{"algorithm": "least_connections", "created_at": "2019-01-01T12:00:00", "health_monitor": {"delay": 5, "max_retries": 2, "port": 22, "timeout": 2, "type": "http", "url_path": "/"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "members": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer-pool", "protocol": "http", "provisioning_status": "active", "session_persistence": {"type": "source_ip"}}]}`)
			}))
			It(`Invoke ListLoadBalancerPools successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLoadBalancerPools(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancerPoolsOptions model
				listLoadBalancerPoolsOptionsModel := new(vpcclassicv1.ListLoadBalancerPoolsOptions)
				listLoadBalancerPoolsOptionsModel.LoadBalancerID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLoadBalancerPools(listLoadBalancerPoolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateLoadBalancerPool(createLoadBalancerPoolOptions *CreateLoadBalancerPoolOptions)`, func() {
		version := "testString"
		createLoadBalancerPoolPath := "/load_balancers/testString/pools"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createLoadBalancerPoolPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"algorithm": "least_connections", "created_at": "2019-01-01T12:00:00", "health_monitor": {"delay": 5, "max_retries": 2, "port": 22, "timeout": 2, "type": "http", "url_path": "/"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "members": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer-pool", "protocol": "http", "provisioning_status": "active", "session_persistence": {"type": "source_ip"}}`)
			}))
			It(`Invoke CreateLoadBalancerPool successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolMemberTargetPrototypeByAddress model
				loadBalancerPoolMemberTargetPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberTargetPrototypeByAddress)
				loadBalancerPoolMemberTargetPrototypeModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the LoadBalancerPoolHealthMonitorPrototype model
				loadBalancerPoolHealthMonitorPrototypeModel := new(vpcclassicv1.LoadBalancerPoolHealthMonitorPrototype)
				loadBalancerPoolHealthMonitorPrototypeModel.Delay = core.Int64Ptr(int64(5))
				loadBalancerPoolHealthMonitorPrototypeModel.MaxRetries = core.Int64Ptr(int64(2))
				loadBalancerPoolHealthMonitorPrototypeModel.Port = core.Int64Ptr(int64(22))
				loadBalancerPoolHealthMonitorPrototypeModel.Timeout = core.Int64Ptr(int64(2))
				loadBalancerPoolHealthMonitorPrototypeModel.Type = core.StringPtr("http")
				loadBalancerPoolHealthMonitorPrototypeModel.UrlPath = core.StringPtr("/")

				// Construct an instance of the LoadBalancerPoolMemberPrototype model
				loadBalancerPoolMemberPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberPrototype)
				loadBalancerPoolMemberPrototypeModel.Port = core.Int64Ptr(int64(80))
				loadBalancerPoolMemberPrototypeModel.Target = loadBalancerPoolMemberTargetPrototypeModel
				loadBalancerPoolMemberPrototypeModel.Weight = core.Int64Ptr(int64(50))

				// Construct an instance of the LoadBalancerPoolSessionPersistencePrototype model
				loadBalancerPoolSessionPersistencePrototypeModel := new(vpcclassicv1.LoadBalancerPoolSessionPersistencePrototype)
				loadBalancerPoolSessionPersistencePrototypeModel.Type = core.StringPtr("source_ip")

				// Construct an instance of the CreateLoadBalancerPoolOptions model
				createLoadBalancerPoolOptionsModel := new(vpcclassicv1.CreateLoadBalancerPoolOptions)
				createLoadBalancerPoolOptionsModel.LoadBalancerID = core.StringPtr("testString")
				createLoadBalancerPoolOptionsModel.Algorithm = core.StringPtr("least_connections")
				createLoadBalancerPoolOptionsModel.HealthMonitor = loadBalancerPoolHealthMonitorPrototypeModel
				createLoadBalancerPoolOptionsModel.Protocol = core.StringPtr("http")
				createLoadBalancerPoolOptionsModel.Members = []vpcclassicv1.LoadBalancerPoolMemberPrototype{*loadBalancerPoolMemberPrototypeModel}
				createLoadBalancerPoolOptionsModel.Name = core.StringPtr("my-load-balancer-pool")
				createLoadBalancerPoolOptionsModel.SessionPersistence = loadBalancerPoolSessionPersistencePrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateLoadBalancerPool(createLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLoadBalancerPool(deleteLoadBalancerPoolOptions *DeleteLoadBalancerPoolOptions)`, func() {
		version := "testString"
		deleteLoadBalancerPoolPath := "/load_balancers/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLoadBalancerPoolPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteLoadBalancerPool successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerPoolOptions model
				deleteLoadBalancerPoolOptionsModel := new(vpcclassicv1.DeleteLoadBalancerPoolOptions)
				deleteLoadBalancerPoolOptionsModel.LoadBalancerID = core.StringPtr("testString")
				deleteLoadBalancerPoolOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteLoadBalancerPool(deleteLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoadBalancerPool(getLoadBalancerPoolOptions *GetLoadBalancerPoolOptions)`, func() {
		version := "testString"
		getLoadBalancerPoolPath := "/load_balancers/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoadBalancerPoolPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"algorithm": "least_connections", "created_at": "2019-01-01T12:00:00", "health_monitor": {"delay": 5, "max_retries": 2, "port": 22, "timeout": 2, "type": "http", "url_path": "/"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "members": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer-pool", "protocol": "http", "provisioning_status": "active", "session_persistence": {"type": "source_ip"}}`)
			}))
			It(`Invoke GetLoadBalancerPool successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerPoolOptions model
				getLoadBalancerPoolOptionsModel := new(vpcclassicv1.GetLoadBalancerPoolOptions)
				getLoadBalancerPoolOptionsModel.LoadBalancerID = core.StringPtr("testString")
				getLoadBalancerPoolOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerPool(getLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateLoadBalancerPool(updateLoadBalancerPoolOptions *UpdateLoadBalancerPoolOptions)`, func() {
		version := "testString"
		updateLoadBalancerPoolPath := "/load_balancers/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateLoadBalancerPoolPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"algorithm": "least_connections", "created_at": "2019-01-01T12:00:00", "health_monitor": {"delay": 5, "max_retries": 2, "port": 22, "timeout": 2, "type": "http", "url_path": "/"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "members": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004"}], "name": "my-load-balancer-pool", "protocol": "http", "provisioning_status": "active", "session_persistence": {"type": "source_ip"}}`)
			}))
			It(`Invoke UpdateLoadBalancerPool successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolHealthMonitorPatch model
				loadBalancerPoolHealthMonitorPatchModel := new(vpcclassicv1.LoadBalancerPoolHealthMonitorPatch)
				loadBalancerPoolHealthMonitorPatchModel.Delay = core.Int64Ptr(int64(5))
				loadBalancerPoolHealthMonitorPatchModel.MaxRetries = core.Int64Ptr(int64(2))
				loadBalancerPoolHealthMonitorPatchModel.Port = core.Int64Ptr(int64(22))
				loadBalancerPoolHealthMonitorPatchModel.Timeout = core.Int64Ptr(int64(2))
				loadBalancerPoolHealthMonitorPatchModel.Type = core.StringPtr("http")
				loadBalancerPoolHealthMonitorPatchModel.UrlPath = core.StringPtr("/")

				// Construct an instance of the LoadBalancerPoolSessionPersistencePatch model
				loadBalancerPoolSessionPersistencePatchModel := new(vpcclassicv1.LoadBalancerPoolSessionPersistencePatch)
				loadBalancerPoolSessionPersistencePatchModel.Type = core.StringPtr("source_ip")

				// Construct an instance of the UpdateLoadBalancerPoolOptions model
				updateLoadBalancerPoolOptionsModel := new(vpcclassicv1.UpdateLoadBalancerPoolOptions)
				updateLoadBalancerPoolOptionsModel.LoadBalancerID = core.StringPtr("testString")
				updateLoadBalancerPoolOptionsModel.ID = core.StringPtr("testString")
				updateLoadBalancerPoolOptionsModel.Algorithm = core.StringPtr("least_connections")
				updateLoadBalancerPoolOptionsModel.HealthMonitor = loadBalancerPoolHealthMonitorPatchModel
				updateLoadBalancerPoolOptionsModel.Name = core.StringPtr("my-load-balancer-pool")
				updateLoadBalancerPoolOptionsModel.Protocol = core.StringPtr("http")
				updateLoadBalancerPoolOptionsModel.SessionPersistence = loadBalancerPoolSessionPersistencePatchModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateLoadBalancerPool(updateLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLoadBalancerPoolMembers(listLoadBalancerPoolMembersOptions *ListLoadBalancerPoolMembersOptions)`, func() {
		version := "testString"
		listLoadBalancerPoolMembersPath := "/load_balancers/testString/pools/testString/members"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLoadBalancerPoolMembersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"members": [{"created_at": "2019-01-01T12:00:00", "health": "faulted", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "port": 80, "provisioning_status": "active", "target": {"address": "192.168.3.4"}, "weight": 50}]}`)
			}))
			It(`Invoke ListLoadBalancerPoolMembers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLoadBalancerPoolMembers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancerPoolMembersOptions model
				listLoadBalancerPoolMembersOptionsModel := new(vpcclassicv1.ListLoadBalancerPoolMembersOptions)
				listLoadBalancerPoolMembersOptionsModel.LoadBalancerID = core.StringPtr("testString")
				listLoadBalancerPoolMembersOptionsModel.PoolID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLoadBalancerPoolMembers(listLoadBalancerPoolMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateLoadBalancerPoolMember(createLoadBalancerPoolMemberOptions *CreateLoadBalancerPoolMemberOptions)`, func() {
		version := "testString"
		createLoadBalancerPoolMemberPath := "/load_balancers/testString/pools/testString/members"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createLoadBalancerPoolMemberPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "health": "faulted", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "port": 80, "provisioning_status": "active", "target": {"address": "192.168.3.4"}, "weight": 50}`)
			}))
			It(`Invoke CreateLoadBalancerPoolMember successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateLoadBalancerPoolMember(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolMemberTargetPrototypeByAddress model
				loadBalancerPoolMemberTargetPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberTargetPrototypeByAddress)
				loadBalancerPoolMemberTargetPrototypeModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the CreateLoadBalancerPoolMemberOptions model
				createLoadBalancerPoolMemberOptionsModel := new(vpcclassicv1.CreateLoadBalancerPoolMemberOptions)
				createLoadBalancerPoolMemberOptionsModel.LoadBalancerID = core.StringPtr("testString")
				createLoadBalancerPoolMemberOptionsModel.PoolID = core.StringPtr("testString")
				createLoadBalancerPoolMemberOptionsModel.Port = core.Int64Ptr(int64(80))
				createLoadBalancerPoolMemberOptionsModel.Target = loadBalancerPoolMemberTargetPrototypeModel
				createLoadBalancerPoolMemberOptionsModel.Weight = core.Int64Ptr(int64(50))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateLoadBalancerPoolMember(createLoadBalancerPoolMemberOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateLoadBalancerPoolMembers(updateLoadBalancerPoolMembersOptions *UpdateLoadBalancerPoolMembersOptions)`, func() {
		version := "testString"
		updateLoadBalancerPoolMembersPath := "/load_balancers/testString/pools/testString/members"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateLoadBalancerPoolMembersPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"members": [{"created_at": "2019-01-01T12:00:00", "health": "faulted", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "port": 80, "provisioning_status": "active", "target": {"address": "192.168.3.4"}, "weight": 50}]}`)
			}))
			It(`Invoke UpdateLoadBalancerPoolMembers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateLoadBalancerPoolMembers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolMemberTargetPrototypeByAddress model
				loadBalancerPoolMemberTargetPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberTargetPrototypeByAddress)
				loadBalancerPoolMemberTargetPrototypeModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the LoadBalancerPoolMemberPrototype model
				loadBalancerPoolMemberPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberPrototype)
				loadBalancerPoolMemberPrototypeModel.Port = core.Int64Ptr(int64(80))
				loadBalancerPoolMemberPrototypeModel.Target = loadBalancerPoolMemberTargetPrototypeModel
				loadBalancerPoolMemberPrototypeModel.Weight = core.Int64Ptr(int64(50))

				// Construct an instance of the UpdateLoadBalancerPoolMembersOptions model
				updateLoadBalancerPoolMembersOptionsModel := new(vpcclassicv1.UpdateLoadBalancerPoolMembersOptions)
				updateLoadBalancerPoolMembersOptionsModel.LoadBalancerID = core.StringPtr("testString")
				updateLoadBalancerPoolMembersOptionsModel.PoolID = core.StringPtr("testString")
				updateLoadBalancerPoolMembersOptionsModel.Members = []vpcclassicv1.LoadBalancerPoolMemberPrototype{*loadBalancerPoolMemberPrototypeModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateLoadBalancerPoolMembers(updateLoadBalancerPoolMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLoadBalancerPoolMember(deleteLoadBalancerPoolMemberOptions *DeleteLoadBalancerPoolMemberOptions)`, func() {
		version := "testString"
		deleteLoadBalancerPoolMemberPath := "/load_balancers/testString/pools/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLoadBalancerPoolMemberPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteLoadBalancerPoolMember successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteLoadBalancerPoolMember(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerPoolMemberOptions model
				deleteLoadBalancerPoolMemberOptionsModel := new(vpcclassicv1.DeleteLoadBalancerPoolMemberOptions)
				deleteLoadBalancerPoolMemberOptionsModel.LoadBalancerID = core.StringPtr("testString")
				deleteLoadBalancerPoolMemberOptionsModel.PoolID = core.StringPtr("testString")
				deleteLoadBalancerPoolMemberOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteLoadBalancerPoolMember(deleteLoadBalancerPoolMemberOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoadBalancerPoolMember(getLoadBalancerPoolMemberOptions *GetLoadBalancerPoolMemberOptions)`, func() {
		version := "testString"
		getLoadBalancerPoolMemberPath := "/load_balancers/testString/pools/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoadBalancerPoolMemberPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "health": "faulted", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "port": 80, "provisioning_status": "active", "target": {"address": "192.168.3.4"}, "weight": 50}`)
			}))
			It(`Invoke GetLoadBalancerPoolMember successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerPoolMember(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerPoolMemberOptions model
				getLoadBalancerPoolMemberOptionsModel := new(vpcclassicv1.GetLoadBalancerPoolMemberOptions)
				getLoadBalancerPoolMemberOptionsModel.LoadBalancerID = core.StringPtr("testString")
				getLoadBalancerPoolMemberOptionsModel.PoolID = core.StringPtr("testString")
				getLoadBalancerPoolMemberOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerPoolMember(getLoadBalancerPoolMemberOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateLoadBalancerPoolMember(updateLoadBalancerPoolMemberOptions *UpdateLoadBalancerPoolMemberOptions)`, func() {
		version := "testString"
		updateLoadBalancerPoolMemberPath := "/load_balancers/testString/pools/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateLoadBalancerPoolMemberPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "health": "faulted", "href": "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004/members/80294e14-4e61-11e8-bcf4-0242ac110004", "id": "70294e14-4e61-11e8-bcf4-0242ac110004", "port": 80, "provisioning_status": "active", "target": {"address": "192.168.3.4"}, "weight": 50}`)
			}))
			It(`Invoke UpdateLoadBalancerPoolMember successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateLoadBalancerPoolMember(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolMemberTargetPrototypeByAddress model
				loadBalancerPoolMemberTargetPrototypeModel := new(vpcclassicv1.LoadBalancerPoolMemberTargetPrototypeByAddress)
				loadBalancerPoolMemberTargetPrototypeModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the UpdateLoadBalancerPoolMemberOptions model
				updateLoadBalancerPoolMemberOptionsModel := new(vpcclassicv1.UpdateLoadBalancerPoolMemberOptions)
				updateLoadBalancerPoolMemberOptionsModel.LoadBalancerID = core.StringPtr("testString")
				updateLoadBalancerPoolMemberOptionsModel.PoolID = core.StringPtr("testString")
				updateLoadBalancerPoolMemberOptionsModel.ID = core.StringPtr("testString")
				updateLoadBalancerPoolMemberOptionsModel.Port = core.Int64Ptr(int64(80))
				updateLoadBalancerPoolMemberOptionsModel.Target = loadBalancerPoolMemberTargetPrototypeModel
				updateLoadBalancerPoolMemberOptionsModel.Weight = core.Int64Ptr(int64(50))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateLoadBalancerPoolMember(updateLoadBalancerPoolMemberOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListNetworkAcls(listNetworkAclsOptions *ListNetworkAclsOptions)`, func() {
		version := "testString"
		listNetworkAclsPath := "/network_acls"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listNetworkAclsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls?limit=20"}, "limit": 20, "network_acls": [{"created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl", "rules": [{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}], "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}], "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}}`)
			}))
			It(`Invoke ListNetworkAcls successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListNetworkAcls(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNetworkAclsOptions model
				listNetworkAclsOptionsModel := new(vpcclassicv1.ListNetworkAclsOptions)
				listNetworkAclsOptionsModel.Start = core.StringPtr("testString")
				listNetworkAclsOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListNetworkAcls(listNetworkAclsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateNetworkAcl(createNetworkAclOptions *CreateNetworkAclOptions)`, func() {
		version := "testString"
		createNetworkAclPath := "/network_acls"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createNetworkAclPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl", "rules": [{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}], "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke CreateNetworkAcl successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateNetworkAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NetworkACLRuleReference model
				networkAclRuleReferenceModel := new(vpcclassicv1.NetworkACLRuleReference)
				networkAclRuleReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRuleReferenceModel.ID = core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRuleReferenceModel.Name = core.StringPtr("my-rule-1")

				// Construct an instance of the NetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolAll model
				networkAclRulePrototypeNetworkAclContextModel := new(vpcclassicv1.NetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolAll)
				networkAclRulePrototypeNetworkAclContextModel.Action = core.StringPtr("allow")
				networkAclRulePrototypeNetworkAclContextModel.Before = networkAclRuleReferenceModel
				networkAclRulePrototypeNetworkAclContextModel.CreatedAt = CreateMockDateTime()
				networkAclRulePrototypeNetworkAclContextModel.Destination = core.StringPtr("192.168.3.0/24")
				networkAclRulePrototypeNetworkAclContextModel.Direction = core.StringPtr("inbound")
				networkAclRulePrototypeNetworkAclContextModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRulePrototypeNetworkAclContextModel.ID = core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRulePrototypeNetworkAclContextModel.IpVersion = core.StringPtr("ipv4")
				networkAclRulePrototypeNetworkAclContextModel.Name = core.StringPtr("my-rule-2")
				networkAclRulePrototypeNetworkAclContextModel.Protocol = core.StringPtr("all")
				networkAclRulePrototypeNetworkAclContextModel.Source = core.StringPtr("192.168.3.0/24")

				// Construct an instance of the NetworkACLPrototypeNetworkACLByRules model
				networkAclPrototypeModel := new(vpcclassicv1.NetworkACLPrototypeNetworkACLByRules)
				networkAclPrototypeModel.Name = core.StringPtr("my-resource")
				networkAclPrototypeModel.Rules = []vpcclassicv1.NetworkACLRulePrototypeNetworkACLContextIntf{networkAclRulePrototypeNetworkAclContextModel}

				// Construct an instance of the CreateNetworkAclOptions model
				createNetworkAclOptionsModel := new(vpcclassicv1.CreateNetworkAclOptions)
				createNetworkAclOptionsModel.NetworkACLPrototype = networkAclPrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateNetworkAcl(createNetworkAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteNetworkAcl(deleteNetworkAclOptions *DeleteNetworkAclOptions)`, func() {
		version := "testString"
		deleteNetworkAclPath := "/network_acls/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteNetworkAclPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteNetworkAcl successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteNetworkAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNetworkAclOptions model
				deleteNetworkAclOptionsModel := new(vpcclassicv1.DeleteNetworkAclOptions)
				deleteNetworkAclOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteNetworkAcl(deleteNetworkAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetNetworkAcl(getNetworkAclOptions *GetNetworkAclOptions)`, func() {
		version := "testString"
		getNetworkAclPath := "/network_acls/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getNetworkAclPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl", "rules": [{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}], "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke GetNetworkAcl successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetNetworkAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNetworkAclOptions model
				getNetworkAclOptionsModel := new(vpcclassicv1.GetNetworkAclOptions)
				getNetworkAclOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetNetworkAcl(getNetworkAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateNetworkAcl(updateNetworkAclOptions *UpdateNetworkAclOptions)`, func() {
		version := "testString"
		updateNetworkAclPath := "/network_acls/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateNetworkAclPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl", "rules": [{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}], "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke UpdateNetworkAcl successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateNetworkAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateNetworkAclOptions model
				updateNetworkAclOptionsModel := new(vpcclassicv1.UpdateNetworkAclOptions)
				updateNetworkAclOptionsModel.ID = core.StringPtr("testString")
				updateNetworkAclOptionsModel.Name = core.StringPtr("my-resource")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateNetworkAcl(updateNetworkAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListNetworkAclRules(listNetworkAclRulesOptions *ListNetworkAclRulesOptions)`, func() {
		version := "testString"
		listNetworkAclRulesPath := "/network_acls/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listNetworkAclRulesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["direction"]).To(Equal([]string{"inbound"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "rules": [{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}]}`)
			}))
			It(`Invoke ListNetworkAclRules successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListNetworkAclRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNetworkAclRulesOptions model
				listNetworkAclRulesOptionsModel := new(vpcclassicv1.ListNetworkAclRulesOptions)
				listNetworkAclRulesOptionsModel.NetworkAclID = core.StringPtr("testString")
				listNetworkAclRulesOptionsModel.Start = core.StringPtr("testString")
				listNetworkAclRulesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listNetworkAclRulesOptionsModel.Direction = core.StringPtr("inbound")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListNetworkAclRules(listNetworkAclRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateNetworkAclRule(createNetworkAclRuleOptions *CreateNetworkAclRuleOptions)`, func() {
		version := "testString"
		createNetworkAclRulePath := "/network_acls/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createNetworkAclRulePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}`)
			}))
			It(`Invoke CreateNetworkAclRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateNetworkAclRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NetworkACLRuleReference model
				networkAclRuleReferenceModel := new(vpcclassicv1.NetworkACLRuleReference)
				networkAclRuleReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRuleReferenceModel.ID = core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRuleReferenceModel.Name = core.StringPtr("my-rule-1")

				// Construct an instance of the NetworkACLRulePrototypeNetworkACLRuleProtocolICMP model
				networkAclRulePrototypeModel := new(vpcclassicv1.NetworkACLRulePrototypeNetworkACLRuleProtocolICMP)
				networkAclRulePrototypeModel.Action = core.StringPtr("allow")
				networkAclRulePrototypeModel.Before = networkAclRuleReferenceModel
				networkAclRulePrototypeModel.CreatedAt = CreateMockDateTime()
				networkAclRulePrototypeModel.Destination = core.StringPtr("192.168.3.0/24")
				networkAclRulePrototypeModel.Direction = core.StringPtr("inbound")
				networkAclRulePrototypeModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRulePrototypeModel.ID = core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRulePrototypeModel.IpVersion = core.StringPtr("ipv4")
				networkAclRulePrototypeModel.Name = core.StringPtr("my-rule-2")
				networkAclRulePrototypeModel.Protocol = core.StringPtr("icmp")
				networkAclRulePrototypeModel.Source = core.StringPtr("192.168.3.0/24")
				networkAclRulePrototypeModel.Code = core.Int64Ptr(int64(0))
				networkAclRulePrototypeModel.Type = core.Int64Ptr(int64(8))

				// Construct an instance of the CreateNetworkAclRuleOptions model
				createNetworkAclRuleOptionsModel := new(vpcclassicv1.CreateNetworkAclRuleOptions)
				createNetworkAclRuleOptionsModel.NetworkAclID = core.StringPtr("testString")
				createNetworkAclRuleOptionsModel.NetworkACLRulePrototype = networkAclRulePrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateNetworkAclRule(createNetworkAclRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteNetworkAclRule(deleteNetworkAclRuleOptions *DeleteNetworkAclRuleOptions)`, func() {
		version := "testString"
		deleteNetworkAclRulePath := "/network_acls/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteNetworkAclRulePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteNetworkAclRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteNetworkAclRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNetworkAclRuleOptions model
				deleteNetworkAclRuleOptionsModel := new(vpcclassicv1.DeleteNetworkAclRuleOptions)
				deleteNetworkAclRuleOptionsModel.NetworkAclID = core.StringPtr("testString")
				deleteNetworkAclRuleOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteNetworkAclRule(deleteNetworkAclRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetNetworkAclRule(getNetworkAclRuleOptions *GetNetworkAclRuleOptions)`, func() {
		version := "testString"
		getNetworkAclRulePath := "/network_acls/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getNetworkAclRulePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}`)
			}))
			It(`Invoke GetNetworkAclRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetNetworkAclRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNetworkAclRuleOptions model
				getNetworkAclRuleOptionsModel := new(vpcclassicv1.GetNetworkAclRuleOptions)
				getNetworkAclRuleOptionsModel.NetworkAclID = core.StringPtr("testString")
				getNetworkAclRuleOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetNetworkAclRule(getNetworkAclRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateNetworkAclRule(updateNetworkAclRuleOptions *UpdateNetworkAclRuleOptions)`, func() {
		version := "testString"
		updateNetworkAclRulePath := "/network_acls/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateNetworkAclRulePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}`)
			}))
			It(`Invoke UpdateNetworkAclRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateNetworkAclRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NetworkACLRuleReference model
				networkAclRuleReferenceModel := new(vpcclassicv1.NetworkACLRuleReference)
				networkAclRuleReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRuleReferenceModel.ID = core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRuleReferenceModel.Name = core.StringPtr("my-rule-1")

				// Construct an instance of the NetworkACLRulePatchNetworkACLRuleProtocolICMP model
				networkAclRulePatchModel := new(vpcclassicv1.NetworkACLRulePatchNetworkACLRuleProtocolICMP)
				networkAclRulePatchModel.Action = core.StringPtr("allow")
				networkAclRulePatchModel.Before = networkAclRuleReferenceModel
				networkAclRulePatchModel.CreatedAt = CreateMockDateTime()
				networkAclRulePatchModel.Destination = core.StringPtr("192.168.3.0/24")
				networkAclRulePatchModel.Direction = core.StringPtr("inbound")
				networkAclRulePatchModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRulePatchModel.ID = core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9")
				networkAclRulePatchModel.IpVersion = core.StringPtr("ipv4")
				networkAclRulePatchModel.Name = core.StringPtr("my-rule-2")
				networkAclRulePatchModel.Protocol = core.StringPtr("icmp")
				networkAclRulePatchModel.Source = core.StringPtr("192.168.3.0/24")
				networkAclRulePatchModel.Code = core.Int64Ptr(int64(0))
				networkAclRulePatchModel.Type = core.Int64Ptr(int64(8))

				// Construct an instance of the UpdateNetworkAclRuleOptions model
				updateNetworkAclRuleOptionsModel := new(vpcclassicv1.UpdateNetworkAclRuleOptions)
				updateNetworkAclRuleOptionsModel.NetworkAclID = core.StringPtr("testString")
				updateNetworkAclRuleOptionsModel.ID = core.StringPtr("testString")
				updateNetworkAclRuleOptionsModel.NetworkACLRulePatch = networkAclRulePatchModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateNetworkAclRule(updateNetworkAclRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListPublicGateways(listPublicGatewaysOptions *ListPublicGatewaysOptions)`, func() {
		version := "testString"
		listPublicGatewaysPath := "/public_gateways"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listPublicGatewaysPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "public_gateways": [{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "floating_ip": {"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway", "status": "available", "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}]}`)
			}))
			It(`Invoke ListPublicGateways successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListPublicGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPublicGatewaysOptions model
				listPublicGatewaysOptionsModel := new(vpcclassicv1.ListPublicGatewaysOptions)
				listPublicGatewaysOptionsModel.Start = core.StringPtr("testString")
				listPublicGatewaysOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListPublicGateways(listPublicGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreatePublicGateway(createPublicGatewayOptions *CreatePublicGatewayOptions)`, func() {
		version := "testString"
		createPublicGatewayPath := "/public_gateways"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createPublicGatewayPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "floating_ip": {"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway", "status": "available", "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke CreatePublicGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreatePublicGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByID model
				publicGatewayPrototypeFloatingIpModel := new(vpcclassicv1.PublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByID)
				publicGatewayPrototypeFloatingIpModel.ID = core.StringPtr("39300233-9995-4806-89a5-3c1b6eb88689")

				// Construct an instance of the VPCIdentityByID model
				vpcIdentityModel := new(vpcclassicv1.VPCIdentityByID)
				vpcIdentityModel.ID = core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(vpcclassicv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreatePublicGatewayOptions model
				createPublicGatewayOptionsModel := new(vpcclassicv1.CreatePublicGatewayOptions)
				createPublicGatewayOptionsModel.Vpc = vpcIdentityModel
				createPublicGatewayOptionsModel.Zone = zoneIdentityModel
				createPublicGatewayOptionsModel.FloatingIp = publicGatewayPrototypeFloatingIpModel
				createPublicGatewayOptionsModel.Name = core.StringPtr("my-public-gateway")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreatePublicGateway(createPublicGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeletePublicGateway(deletePublicGatewayOptions *DeletePublicGatewayOptions)`, func() {
		version := "testString"
		deletePublicGatewayPath := "/public_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deletePublicGatewayPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeletePublicGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeletePublicGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePublicGatewayOptions model
				deletePublicGatewayOptionsModel := new(vpcclassicv1.DeletePublicGatewayOptions)
				deletePublicGatewayOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeletePublicGateway(deletePublicGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetPublicGateway(getPublicGatewayOptions *GetPublicGatewayOptions)`, func() {
		version := "testString"
		getPublicGatewayPath := "/public_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getPublicGatewayPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "floating_ip": {"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway", "status": "available", "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetPublicGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetPublicGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPublicGatewayOptions model
				getPublicGatewayOptionsModel := new(vpcclassicv1.GetPublicGatewayOptions)
				getPublicGatewayOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPublicGateway(getPublicGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdatePublicGateway(updatePublicGatewayOptions *UpdatePublicGatewayOptions)`, func() {
		version := "testString"
		updatePublicGatewayPath := "/public_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updatePublicGatewayPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "floating_ip": {"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway", "status": "available", "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke UpdatePublicGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdatePublicGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdatePublicGatewayOptions model
				updatePublicGatewayOptionsModel := new(vpcclassicv1.UpdatePublicGatewayOptions)
				updatePublicGatewayOptionsModel.ID = core.StringPtr("testString")
				updatePublicGatewayOptionsModel.Name = core.StringPtr("my-public-gateway")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdatePublicGateway(updatePublicGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListKeys(listKeysOptions *ListKeysOptions)`, func() {
		version := "testString"
		listKeysPath := "/keys"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listKeysPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/keys?limit=20"}, "keys": [{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::key:a6b1a881-2ce8-41a3-80fc-36316a73f803", "fingerprint": "SHA256:yxavE4CIOL2NlsqcurRO3xGjkP6m/0mp8ugojH5yxlY", "href": "https://us-south.iaas.cloud.ibm.com/v1/keys/a6b1a881-2ce8-41a3-80fc-36316a73f803", "id": "a6b1a881-2ce8-41a3-80fc-36316a73f803", "length": 2048, "name": "my-key", "public_key": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "type": "rsa"}], "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/keys?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "total_count": 132}`)
			}))
			It(`Invoke ListKeys successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListKeysOptions model
				listKeysOptionsModel := new(vpcclassicv1.ListKeysOptions)
				listKeysOptionsModel.Start = core.StringPtr("testString")
				listKeysOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListKeys(listKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateKey(createKeyOptions *CreateKeyOptions)`, func() {
		version := "testString"
		createKeyPath := "/keys"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createKeyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::key:a6b1a881-2ce8-41a3-80fc-36316a73f803", "fingerprint": "SHA256:yxavE4CIOL2NlsqcurRO3xGjkP6m/0mp8ugojH5yxlY", "href": "https://us-south.iaas.cloud.ibm.com/v1/keys/a6b1a881-2ce8-41a3-80fc-36316a73f803", "id": "a6b1a881-2ce8-41a3-80fc-36316a73f803", "length": 2048, "name": "my-key", "public_key": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "type": "rsa"}`)
			}))
			It(`Invoke CreateKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the CreateKeyOptions model
				createKeyOptionsModel := new(vpcclassicv1.CreateKeyOptions)
				createKeyOptionsModel.PublicKey = core.StringPtr("AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En")
				createKeyOptionsModel.Name = core.StringPtr("my-key")
				createKeyOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createKeyOptionsModel.Type = core.StringPtr("rsa")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateKey(createKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteKey(deleteKeyOptions *DeleteKeyOptions)`, func() {
		version := "testString"
		deleteKeyPath := "/keys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteKeyPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteKeyOptions model
				deleteKeyOptionsModel := new(vpcclassicv1.DeleteKeyOptions)
				deleteKeyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteKey(deleteKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetKey(getKeyOptions *GetKeyOptions)`, func() {
		version := "testString"
		getKeyPath := "/keys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getKeyPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::key:a6b1a881-2ce8-41a3-80fc-36316a73f803", "fingerprint": "SHA256:yxavE4CIOL2NlsqcurRO3xGjkP6m/0mp8ugojH5yxlY", "href": "https://us-south.iaas.cloud.ibm.com/v1/keys/a6b1a881-2ce8-41a3-80fc-36316a73f803", "id": "a6b1a881-2ce8-41a3-80fc-36316a73f803", "length": 2048, "name": "my-key", "public_key": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "type": "rsa"}`)
			}))
			It(`Invoke GetKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetKeyOptions model
				getKeyOptionsModel := new(vpcclassicv1.GetKeyOptions)
				getKeyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetKey(getKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateKey(updateKeyOptions *UpdateKeyOptions)`, func() {
		version := "testString"
		updateKeyPath := "/keys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateKeyPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::key:a6b1a881-2ce8-41a3-80fc-36316a73f803", "fingerprint": "SHA256:yxavE4CIOL2NlsqcurRO3xGjkP6m/0mp8ugojH5yxlY", "href": "https://us-south.iaas.cloud.ibm.com/v1/keys/a6b1a881-2ce8-41a3-80fc-36316a73f803", "id": "a6b1a881-2ce8-41a3-80fc-36316a73f803", "length": 2048, "name": "my-key", "public_key": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "type": "rsa"}`)
			}))
			It(`Invoke UpdateKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateKeyOptions model
				updateKeyOptionsModel := new(vpcclassicv1.UpdateKeyOptions)
				updateKeyOptionsModel.ID = core.StringPtr("testString")
				updateKeyOptionsModel.Name = core.StringPtr("my-key")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateKey(updateKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListSecurityGroups(listSecurityGroupsOptions *ListSecurityGroupsOptions)`, func() {
		version := "testString"
		listSecurityGroupsPath := "/security_groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listSecurityGroupsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["vpc.id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["vpc.crn"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["vpc.name"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"security_groups": [{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "rules": [{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}}]}`)
			}))
			It(`Invoke ListSecurityGroups successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListSecurityGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSecurityGroupsOptions model
				listSecurityGroupsOptionsModel := new(vpcclassicv1.ListSecurityGroupsOptions)
				listSecurityGroupsOptionsModel.VpcID = core.StringPtr("testString")
				listSecurityGroupsOptionsModel.VpcCrn = core.StringPtr("testString")
				listSecurityGroupsOptionsModel.VpcName = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListSecurityGroups(listSecurityGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateSecurityGroup(createSecurityGroupOptions *CreateSecurityGroupOptions)`, func() {
		version := "testString"
		createSecurityGroupPath := "/security_groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSecurityGroupPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "rules": [{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}}`)
			}))
			It(`Invoke CreateSecurityGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateSecurityGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteIP model
				securityGroupRulePrototypeSecurityGroupRuleProtocolIcmpRemoteModel := new(vpcclassicv1.SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteIP)
				securityGroupRulePrototypeSecurityGroupRuleProtocolIcmpRemoteModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMP model
				securityGroupRulePrototypeModel := new(vpcclassicv1.SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMP)
				securityGroupRulePrototypeModel.Direction = core.StringPtr("inbound")
				securityGroupRulePrototypeModel.IpVersion = core.StringPtr("ipv4")
				securityGroupRulePrototypeModel.Protocol = core.StringPtr("icmp")
				securityGroupRulePrototypeModel.Remote = securityGroupRulePrototypeSecurityGroupRuleProtocolIcmpRemoteModel
				securityGroupRulePrototypeModel.Code = core.Int64Ptr(int64(0))
				securityGroupRulePrototypeModel.Type = core.Int64Ptr(int64(8))

				// Construct an instance of the VPCIdentityByID model
				vpcIdentityModel := new(vpcclassicv1.VPCIdentityByID)
				vpcIdentityModel.ID = core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b")

				// Construct an instance of the CreateSecurityGroupOptions model
				createSecurityGroupOptionsModel := new(vpcclassicv1.CreateSecurityGroupOptions)
				createSecurityGroupOptionsModel.Vpc = vpcIdentityModel
				createSecurityGroupOptionsModel.Name = core.StringPtr("my-security-group")
				createSecurityGroupOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createSecurityGroupOptionsModel.Rules = []vpcclassicv1.SecurityGroupRulePrototypeIntf{securityGroupRulePrototypeModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateSecurityGroup(createSecurityGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSecurityGroup(deleteSecurityGroupOptions *DeleteSecurityGroupOptions)`, func() {
		version := "testString"
		deleteSecurityGroupPath := "/security_groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSecurityGroupPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteSecurityGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteSecurityGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSecurityGroupOptions model
				deleteSecurityGroupOptionsModel := new(vpcclassicv1.DeleteSecurityGroupOptions)
				deleteSecurityGroupOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteSecurityGroup(deleteSecurityGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSecurityGroup(getSecurityGroupOptions *GetSecurityGroupOptions)`, func() {
		version := "testString"
		getSecurityGroupPath := "/security_groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSecurityGroupPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "rules": [{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}}`)
			}))
			It(`Invoke GetSecurityGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSecurityGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSecurityGroupOptions model
				getSecurityGroupOptionsModel := new(vpcclassicv1.GetSecurityGroupOptions)
				getSecurityGroupOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSecurityGroup(getSecurityGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateSecurityGroup(updateSecurityGroupOptions *UpdateSecurityGroupOptions)`, func() {
		version := "testString"
		updateSecurityGroupPath := "/security_groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateSecurityGroupPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group", "network_interfaces": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface"}], "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "rules": [{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}}`)
			}))
			It(`Invoke UpdateSecurityGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateSecurityGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSecurityGroupOptions model
				updateSecurityGroupOptionsModel := new(vpcclassicv1.UpdateSecurityGroupOptions)
				updateSecurityGroupOptionsModel.ID = core.StringPtr("testString")
				updateSecurityGroupOptionsModel.Name = core.StringPtr("my-security-group")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateSecurityGroup(updateSecurityGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListSecurityGroupNetworkInterfaces(listSecurityGroupNetworkInterfacesOptions *ListSecurityGroupNetworkInterfacesOptions)`, func() {
		version := "testString"
		listSecurityGroupNetworkInterfacesPath := "/security_groups/testString/network_interfaces"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listSecurityGroupNetworkInterfacesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"network_interfaces": [{"created_at": "2019-01-01T12:00:00", "floating_ips": [{"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}], "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "port_speed": 1000, "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "security_groups": [{"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}], "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}, "type": "primary"}]}`)
			}))
			It(`Invoke ListSecurityGroupNetworkInterfaces successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListSecurityGroupNetworkInterfaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSecurityGroupNetworkInterfacesOptions model
				listSecurityGroupNetworkInterfacesOptionsModel := new(vpcclassicv1.ListSecurityGroupNetworkInterfacesOptions)
				listSecurityGroupNetworkInterfacesOptionsModel.SecurityGroupID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListSecurityGroupNetworkInterfaces(listSecurityGroupNetworkInterfacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSecurityGroupNetworkInterfaceBinding(deleteSecurityGroupNetworkInterfaceBindingOptions *DeleteSecurityGroupNetworkInterfaceBindingOptions)`, func() {
		version := "testString"
		deleteSecurityGroupNetworkInterfaceBindingPath := "/security_groups/testString/network_interfaces/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSecurityGroupNetworkInterfaceBindingPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteSecurityGroupNetworkInterfaceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteSecurityGroupNetworkInterfaceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSecurityGroupNetworkInterfaceBindingOptions model
				deleteSecurityGroupNetworkInterfaceBindingOptionsModel := new(vpcclassicv1.DeleteSecurityGroupNetworkInterfaceBindingOptions)
				deleteSecurityGroupNetworkInterfaceBindingOptionsModel.SecurityGroupID = core.StringPtr("testString")
				deleteSecurityGroupNetworkInterfaceBindingOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteSecurityGroupNetworkInterfaceBinding(deleteSecurityGroupNetworkInterfaceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions *GetSecurityGroupNetworkInterfaceOptions)`, func() {
		version := "testString"
		getSecurityGroupNetworkInterfacePath := "/security_groups/testString/network_interfaces/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSecurityGroupNetworkInterfacePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "floating_ips": [{"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}], "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "port_speed": 1000, "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "security_groups": [{"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}], "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}, "type": "primary"}`)
			}))
			It(`Invoke GetSecurityGroupNetworkInterface successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSecurityGroupNetworkInterface(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSecurityGroupNetworkInterfaceOptions model
				getSecurityGroupNetworkInterfaceOptionsModel := new(vpcclassicv1.GetSecurityGroupNetworkInterfaceOptions)
				getSecurityGroupNetworkInterfaceOptionsModel.SecurityGroupID = core.StringPtr("testString")
				getSecurityGroupNetworkInterfaceOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateSecurityGroupNetworkInterfaceBinding(createSecurityGroupNetworkInterfaceBindingOptions *CreateSecurityGroupNetworkInterfaceBindingOptions)`, func() {
		version := "testString"
		createSecurityGroupNetworkInterfaceBindingPath := "/security_groups/testString/network_interfaces/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSecurityGroupNetworkInterfaceBindingPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "floating_ips": [{"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}], "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e", "id": "10c02d81-0ecb-4dc5-897d-28392913b81e", "name": "my-network-interface", "port_speed": 1000, "primary_ipv4_address": "192.168.3.4", "resource_type": "network_interface", "security_groups": [{"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}], "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}, "type": "primary"}`)
			}))
			It(`Invoke CreateSecurityGroupNetworkInterfaceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateSecurityGroupNetworkInterfaceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSecurityGroupNetworkInterfaceBindingOptions model
				createSecurityGroupNetworkInterfaceBindingOptionsModel := new(vpcclassicv1.CreateSecurityGroupNetworkInterfaceBindingOptions)
				createSecurityGroupNetworkInterfaceBindingOptionsModel.SecurityGroupID = core.StringPtr("testString")
				createSecurityGroupNetworkInterfaceBindingOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateSecurityGroupNetworkInterfaceBinding(createSecurityGroupNetworkInterfaceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListSecurityGroupRules(listSecurityGroupRulesOptions *ListSecurityGroupRulesOptions)`, func() {
		version := "testString"
		listSecurityGroupRulesPath := "/security_groups/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listSecurityGroupRulesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rules": [{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}]}`)
			}))
			It(`Invoke ListSecurityGroupRules successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListSecurityGroupRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSecurityGroupRulesOptions model
				listSecurityGroupRulesOptionsModel := new(vpcclassicv1.ListSecurityGroupRulesOptions)
				listSecurityGroupRulesOptionsModel.SecurityGroupID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListSecurityGroupRules(listSecurityGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateSecurityGroupRule(createSecurityGroupRuleOptions *CreateSecurityGroupRuleOptions)`, func() {
		version := "testString"
		createSecurityGroupRulePath := "/security_groups/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSecurityGroupRulePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}`)
			}))
			It(`Invoke CreateSecurityGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateSecurityGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteIP model
				securityGroupRulePrototypeSecurityGroupRuleProtocolIcmpRemoteModel := new(vpcclassicv1.SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteIP)
				securityGroupRulePrototypeSecurityGroupRuleProtocolIcmpRemoteModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMP model
				securityGroupRulePrototypeModel := new(vpcclassicv1.SecurityGroupRulePrototypeSecurityGroupRuleProtocolICMP)
				securityGroupRulePrototypeModel.Direction = core.StringPtr("inbound")
				securityGroupRulePrototypeModel.IpVersion = core.StringPtr("ipv4")
				securityGroupRulePrototypeModel.Protocol = core.StringPtr("icmp")
				securityGroupRulePrototypeModel.Remote = securityGroupRulePrototypeSecurityGroupRuleProtocolIcmpRemoteModel
				securityGroupRulePrototypeModel.Code = core.Int64Ptr(int64(0))
				securityGroupRulePrototypeModel.Type = core.Int64Ptr(int64(8))

				// Construct an instance of the CreateSecurityGroupRuleOptions model
				createSecurityGroupRuleOptionsModel := new(vpcclassicv1.CreateSecurityGroupRuleOptions)
				createSecurityGroupRuleOptionsModel.SecurityGroupID = core.StringPtr("testString")
				createSecurityGroupRuleOptionsModel.SecurityGroupRulePrototype = securityGroupRulePrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateSecurityGroupRule(createSecurityGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSecurityGroupRule(deleteSecurityGroupRuleOptions *DeleteSecurityGroupRuleOptions)`, func() {
		version := "testString"
		deleteSecurityGroupRulePath := "/security_groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSecurityGroupRulePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteSecurityGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteSecurityGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSecurityGroupRuleOptions model
				deleteSecurityGroupRuleOptionsModel := new(vpcclassicv1.DeleteSecurityGroupRuleOptions)
				deleteSecurityGroupRuleOptionsModel.SecurityGroupID = core.StringPtr("testString")
				deleteSecurityGroupRuleOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteSecurityGroupRule(deleteSecurityGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSecurityGroupRule(getSecurityGroupRuleOptions *GetSecurityGroupRuleOptions)`, func() {
		version := "testString"
		getSecurityGroupRulePath := "/security_groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSecurityGroupRulePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}`)
			}))
			It(`Invoke GetSecurityGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSecurityGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSecurityGroupRuleOptions model
				getSecurityGroupRuleOptionsModel := new(vpcclassicv1.GetSecurityGroupRuleOptions)
				getSecurityGroupRuleOptionsModel.SecurityGroupID = core.StringPtr("testString")
				getSecurityGroupRuleOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSecurityGroupRule(getSecurityGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateSecurityGroupRule(updateSecurityGroupRuleOptions *UpdateSecurityGroupRuleOptions)`, func() {
		version := "testString"
		updateSecurityGroupRulePath := "/security_groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateSecurityGroupRulePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}`)
			}))
			It(`Invoke UpdateSecurityGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateSecurityGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteIP model
				securityGroupRulePatchSecurityGroupRuleProtocolIcmpRemoteModel := new(vpcclassicv1.SecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteIP)
				securityGroupRulePatchSecurityGroupRuleProtocolIcmpRemoteModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the SecurityGroupRulePatchSecurityGroupRuleProtocolICMP model
				securityGroupRulePatchModel := new(vpcclassicv1.SecurityGroupRulePatchSecurityGroupRuleProtocolICMP)
				securityGroupRulePatchModel.Direction = core.StringPtr("inbound")
				securityGroupRulePatchModel.IpVersion = core.StringPtr("ipv4")
				securityGroupRulePatchModel.Protocol = core.StringPtr("icmp")
				securityGroupRulePatchModel.Remote = securityGroupRulePatchSecurityGroupRuleProtocolIcmpRemoteModel
				securityGroupRulePatchModel.Code = core.Int64Ptr(int64(0))
				securityGroupRulePatchModel.Type = core.Int64Ptr(int64(8))

				// Construct an instance of the UpdateSecurityGroupRuleOptions model
				updateSecurityGroupRuleOptionsModel := new(vpcclassicv1.UpdateSecurityGroupRuleOptions)
				updateSecurityGroupRuleOptionsModel.SecurityGroupID = core.StringPtr("testString")
				updateSecurityGroupRuleOptionsModel.ID = core.StringPtr("testString")
				updateSecurityGroupRuleOptionsModel.SecurityGroupRulePatch = securityGroupRulePatchModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateSecurityGroupRule(updateSecurityGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListSubnets(listSubnetsOptions *ListSubnetsOptions)`, func() {
		version := "testString"
		listSubnetsPath := "/subnets"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listSubnetsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/subnets?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/subnets?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "subnets": [{"available_ipv4_address_count": 15, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "ipv4_cidr_block": "10.0.0.0/24", "name": "my-subnet", "network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "public_gateway": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway"}, "status": "available", "total_ipv4_address_count": 256, "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}]}`)
			}))
			It(`Invoke ListSubnets successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListSubnets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSubnetsOptions model
				listSubnetsOptionsModel := new(vpcclassicv1.ListSubnetsOptions)
				listSubnetsOptionsModel.Start = core.StringPtr("testString")
				listSubnetsOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListSubnets(listSubnetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateSubnet(createSubnetOptions *CreateSubnetOptions)`, func() {
		version := "testString"
		createSubnetPath := "/subnets"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSubnetPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"available_ipv4_address_count": 15, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "ipv4_cidr_block": "10.0.0.0/24", "name": "my-subnet", "network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "public_gateway": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway"}, "status": "available", "total_ipv4_address_count": 256, "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke CreateSubnet successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateSubnet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NetworkACLIdentityByID model
				networkAclIdentityModel := new(vpcclassicv1.NetworkACLIdentityByID)
				networkAclIdentityModel.ID = core.StringPtr("a4e28308-8ee7-46ab-8108-9f881f22bdbf")

				// Construct an instance of the PublicGatewayIdentityByID model
				publicGatewayIdentityModel := new(vpcclassicv1.PublicGatewayIdentityByID)
				publicGatewayIdentityModel.ID = core.StringPtr("dc5431ef-1fc6-4861-adc9-a59d077d1241")

				// Construct an instance of the VPCIdentityByID model
				vpcIdentityModel := new(vpcclassicv1.VPCIdentityByID)
				vpcIdentityModel.ID = core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(vpcclassicv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the SubnetPrototypeSubnetByTotalCount model
				subnetPrototypeModel := new(vpcclassicv1.SubnetPrototypeSubnetByTotalCount)
				subnetPrototypeModel.Name = core.StringPtr("my-subnet")
				subnetPrototypeModel.NetworkAcl = networkAclIdentityModel
				subnetPrototypeModel.PublicGateway = publicGatewayIdentityModel
				subnetPrototypeModel.Vpc = vpcIdentityModel
				subnetPrototypeModel.TotalIpv4AddressCount = core.Int64Ptr(int64(256))
				subnetPrototypeModel.Zone = zoneIdentityModel

				// Construct an instance of the CreateSubnetOptions model
				createSubnetOptionsModel := new(vpcclassicv1.CreateSubnetOptions)
				createSubnetOptionsModel.SubnetPrototype = subnetPrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateSubnet(createSubnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSubnet(deleteSubnetOptions *DeleteSubnetOptions)`, func() {
		version := "testString"
		deleteSubnetPath := "/subnets/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSubnetPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteSubnet successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteSubnet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSubnetOptions model
				deleteSubnetOptionsModel := new(vpcclassicv1.DeleteSubnetOptions)
				deleteSubnetOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteSubnet(deleteSubnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSubnet(getSubnetOptions *GetSubnetOptions)`, func() {
		version := "testString"
		getSubnetPath := "/subnets/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSubnetPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"available_ipv4_address_count": 15, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "ipv4_cidr_block": "10.0.0.0/24", "name": "my-subnet", "network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "public_gateway": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway"}, "status": "available", "total_ipv4_address_count": 256, "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetSubnet successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSubnet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSubnetOptions model
				getSubnetOptionsModel := new(vpcclassicv1.GetSubnetOptions)
				getSubnetOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSubnet(getSubnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateSubnet(updateSubnetOptions *UpdateSubnetOptions)`, func() {
		version := "testString"
		updateSubnetPath := "/subnets/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateSubnetPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"available_ipv4_address_count": 15, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "ipv4_cidr_block": "10.0.0.0/24", "name": "my-subnet", "network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "public_gateway": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway"}, "status": "available", "total_ipv4_address_count": 256, "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke UpdateSubnet successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateSubnet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NetworkACLIdentityByID model
				networkAclIdentityModel := new(vpcclassicv1.NetworkACLIdentityByID)
				networkAclIdentityModel.ID = core.StringPtr("a4e28308-8ee7-46ab-8108-9f881f22bdbf")

				// Construct an instance of the PublicGatewayIdentityByID model
				publicGatewayIdentityModel := new(vpcclassicv1.PublicGatewayIdentityByID)
				publicGatewayIdentityModel.ID = core.StringPtr("dc5431ef-1fc6-4861-adc9-a59d077d1241")

				// Construct an instance of the UpdateSubnetOptions model
				updateSubnetOptionsModel := new(vpcclassicv1.UpdateSubnetOptions)
				updateSubnetOptionsModel.ID = core.StringPtr("testString")
				updateSubnetOptionsModel.Name = core.StringPtr("my-subnet")
				updateSubnetOptionsModel.NetworkAcl = networkAclIdentityModel
				updateSubnetOptionsModel.PublicGateway = publicGatewayIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateSubnet(updateSubnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSubnetNetworkAcl(getSubnetNetworkAclOptions *GetSubnetNetworkAclOptions)`, func() {
		version := "testString"
		getSubnetNetworkAclPath := "/subnets/testString/network_acl"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSubnetNetworkAclPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl", "rules": [{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}], "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke GetSubnetNetworkAcl successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSubnetNetworkAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSubnetNetworkAclOptions model
				getSubnetNetworkAclOptionsModel := new(vpcclassicv1.GetSubnetNetworkAclOptions)
				getSubnetNetworkAclOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSubnetNetworkAcl(getSubnetNetworkAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`SetSubnetNetworkAclBinding(setSubnetNetworkAclBindingOptions *SetSubnetNetworkAclBindingOptions)`, func() {
		version := "testString"
		setSubnetNetworkAclBindingPath := "/subnets/testString/network_acl"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(setSubnetNetworkAclBindingPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl", "rules": [{"action": "allow", "before": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "name": "my-rule-1"}, "created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "direction": "inbound", "href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9", "id": "8daca77a-4980-4d33-8f3e-7038797be8f9", "ip_version": "ipv4", "name": "my-rule-2", "protocol": "udp", "source": "192.168.3.0/24", "port_max": 22, "port_min": 22, "source_port_max": 65535, "source_port_min": 49152}], "subnets": [{"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}]}`)
			}))
			It(`Invoke SetSubnetNetworkAclBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.SetSubnetNetworkAclBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NetworkACLIdentityByID model
				networkAclIdentityModel := new(vpcclassicv1.NetworkACLIdentityByID)
				networkAclIdentityModel.ID = core.StringPtr("a4e28308-8ee7-46ab-8108-9f881f22bdbf")

				// Construct an instance of the SetSubnetNetworkAclBindingOptions model
				setSubnetNetworkAclBindingOptionsModel := new(vpcclassicv1.SetSubnetNetworkAclBindingOptions)
				setSubnetNetworkAclBindingOptionsModel.ID = core.StringPtr("testString")
				setSubnetNetworkAclBindingOptionsModel.NetworkACLIdentity = networkAclIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.SetSubnetNetworkAclBinding(setSubnetNetworkAclBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSubnetPublicGatewayBinding(deleteSubnetPublicGatewayBindingOptions *DeleteSubnetPublicGatewayBindingOptions)`, func() {
		version := "testString"
		deleteSubnetPublicGatewayBindingPath := "/subnets/testString/public_gateway"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSubnetPublicGatewayBindingPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteSubnetPublicGatewayBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteSubnetPublicGatewayBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSubnetPublicGatewayBindingOptions model
				deleteSubnetPublicGatewayBindingOptionsModel := new(vpcclassicv1.DeleteSubnetPublicGatewayBindingOptions)
				deleteSubnetPublicGatewayBindingOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteSubnetPublicGatewayBinding(deleteSubnetPublicGatewayBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSubnetPublicGateway(getSubnetPublicGatewayOptions *GetSubnetPublicGatewayOptions)`, func() {
		version := "testString"
		getSubnetPublicGatewayPath := "/subnets/testString/public_gateway"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSubnetPublicGatewayPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "floating_ip": {"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway", "status": "available", "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetSubnetPublicGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSubnetPublicGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSubnetPublicGatewayOptions model
				getSubnetPublicGatewayOptionsModel := new(vpcclassicv1.GetSubnetPublicGatewayOptions)
				getSubnetPublicGatewayOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`SetSubnetPublicGatewayBinding(setSubnetPublicGatewayBindingOptions *SetSubnetPublicGatewayBindingOptions)`, func() {
		version := "testString"
		setSubnetPublicGatewayBindingPath := "/subnets/testString/public_gateway"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(setSubnetPublicGatewayBindingPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241", "floating_ip": {"address": "203.0.113.1", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689", "href": "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689", "id": "39300233-9995-4806-89a5-3c1b6eb88689", "name": "my-floating-ip"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241", "id": "dc5431ef-1fc6-4861-adc9-a59d077d1241", "name": "my-public-gateway", "resource_type": "public_gateway", "status": "available", "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke SetSubnetPublicGatewayBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.SetSubnetPublicGatewayBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PublicGatewayIdentityByID model
				publicGatewayIdentityModel := new(vpcclassicv1.PublicGatewayIdentityByID)
				publicGatewayIdentityModel.ID = core.StringPtr("dc5431ef-1fc6-4861-adc9-a59d077d1241")

				// Construct an instance of the SetSubnetPublicGatewayBindingOptions model
				setSubnetPublicGatewayBindingOptionsModel := new(vpcclassicv1.SetSubnetPublicGatewayBindingOptions)
				setSubnetPublicGatewayBindingOptionsModel.ID = core.StringPtr("testString")
				setSubnetPublicGatewayBindingOptionsModel.PublicGatewayIdentity = publicGatewayIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.SetSubnetPublicGatewayBinding(setSubnetPublicGatewayBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpcs(listVpcsOptions *ListVpcsOptions)`, func() {
		version := "testString"
		listVpcsPath := "/vpcs"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpcsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				// TODO: Add check for classic_access query parameter

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "vpcs": [{"classic_access": false, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "cse_source_ips": [{"ip": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}], "default_network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "default_security_group": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available"}]}`)
			}))
			It(`Invoke ListVpcs successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpcsOptions model
				listVpcsOptionsModel := new(vpcclassicv1.ListVpcsOptions)
				listVpcsOptionsModel.Start = core.StringPtr("testString")
				listVpcsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listVpcsOptionsModel.ClassicAccess = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpcs(listVpcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVpc(createVpcOptions *CreateVpcOptions)`, func() {
		version := "testString"
		createVpcPath := "/vpcs"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVpcPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"classic_access": false, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "cse_source_ips": [{"ip": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}], "default_network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "default_security_group": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available"}`)
			}))
			It(`Invoke CreateVpc successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateVpc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the CreateVpcOptions model
				createVpcOptionsModel := new(vpcclassicv1.CreateVpcOptions)
				createVpcOptionsModel.AddressPrefixManagement = core.StringPtr("manual")
				createVpcOptionsModel.ClassicAccess = core.BoolPtr(false)
				createVpcOptionsModel.Name = core.StringPtr("my-vpc")
				createVpcOptionsModel.ResourceGroup = resourceGroupIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateVpc(createVpcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVpc(deleteVpcOptions *DeleteVpcOptions)`, func() {
		version := "testString"
		deleteVpcPath := "/vpcs/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVpcPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteVpc successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVpc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVpcOptions model
				deleteVpcOptionsModel := new(vpcclassicv1.DeleteVpcOptions)
				deleteVpcOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVpc(deleteVpcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpc(getVpcOptions *GetVpcOptions)`, func() {
		version := "testString"
		getVpcPath := "/vpcs/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpcPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"classic_access": false, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "cse_source_ips": [{"ip": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}], "default_network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "default_security_group": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available"}`)
			}))
			It(`Invoke GetVpc successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVpc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVpcOptions model
				getVpcOptionsModel := new(vpcclassicv1.GetVpcOptions)
				getVpcOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVpc(getVpcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVpc(updateVpcOptions *UpdateVpcOptions)`, func() {
		version := "testString"
		updateVpcPath := "/vpcs/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVpcPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"classic_access": false, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "cse_source_ips": [{"ip": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}], "default_network_acl": {"href": "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf", "id": "a4e28308-8ee7-46ab-8108-9f881f22bdbf", "name": "my-network-acl"}, "default_security_group": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "my-security-group"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available"}`)
			}))
			It(`Invoke UpdateVpc successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateVpc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVpcOptions model
				updateVpcOptionsModel := new(vpcclassicv1.UpdateVpcOptions)
				updateVpcOptionsModel.ID = core.StringPtr("testString")
				updateVpcOptionsModel.Name = core.StringPtr("my-vpc")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateVpc(updateVpcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpcDefaultSecurityGroup(getVpcDefaultSecurityGroupOptions *GetVpcDefaultSecurityGroupOptions)`, func() {
		version := "testString"
		getVpcDefaultSecurityGroupPath := "/vpcs/testString/default_security_group"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpcDefaultSecurityGroupPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271", "href": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271", "id": "be5df5ca-12a0-494b-907e-aa6ec2bfa271", "name": "observant-chip-emphatic-engraver", "rules": [{"direction": "inbound", "id": "6f2a6efe-21e2-401c-b237-620aa26ba16a", "ip_version": "ipv4", "protocol": "udp", "remote": {"anyKey": "anyValue"}, "port_max": 22, "port_min": 22}], "vpc": {"crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b", "id": "4727d842-f94f-4a2d-824a-9bc9b02c523b", "name": "my-vpc"}}`)
			}))
			It(`Invoke GetVpcDefaultSecurityGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVpcDefaultSecurityGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVpcDefaultSecurityGroupOptions model
				getVpcDefaultSecurityGroupOptionsModel := new(vpcclassicv1.GetVpcDefaultSecurityGroupOptions)
				getVpcDefaultSecurityGroupOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVpcDefaultSecurityGroup(getVpcDefaultSecurityGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpcAddressPrefixes(listVpcAddressPrefixesOptions *ListVpcAddressPrefixesOptions)`, func() {
		version := "testString"
		listVpcAddressPrefixesPath := "/vpcs/testString/address_prefixes"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpcAddressPrefixesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"address_prefixes": [{"cidr": "192.168.3.0/24", "created_at": "2019-01-01T12:00:00", "has_subnets": true, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/address_prefixes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "is_default": false, "name": "my-address-prefix-2", "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}], "first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/a4e28308-8ee7-46ab-8108-9f881f22bdbf/address_prefixes?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/a4e28308-8ee7-46ab-8108-9f881f22bdbf/address_prefixes?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}}`)
			}))
			It(`Invoke ListVpcAddressPrefixes successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpcAddressPrefixes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpcAddressPrefixesOptions model
				listVpcAddressPrefixesOptionsModel := new(vpcclassicv1.ListVpcAddressPrefixesOptions)
				listVpcAddressPrefixesOptionsModel.VpcID = core.StringPtr("testString")
				listVpcAddressPrefixesOptionsModel.Start = core.StringPtr("testString")
				listVpcAddressPrefixesOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpcAddressPrefixes(listVpcAddressPrefixesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVpcAddressPrefix(createVpcAddressPrefixOptions *CreateVpcAddressPrefixOptions)`, func() {
		version := "testString"
		createVpcAddressPrefixPath := "/vpcs/testString/address_prefixes"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVpcAddressPrefixPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"cidr": "192.168.3.0/24", "created_at": "2019-01-01T12:00:00", "has_subnets": true, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/address_prefixes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "is_default": false, "name": "my-address-prefix-2", "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke CreateVpcAddressPrefix successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateVpcAddressPrefix(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(vpcclassicv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreateVpcAddressPrefixOptions model
				createVpcAddressPrefixOptionsModel := new(vpcclassicv1.CreateVpcAddressPrefixOptions)
				createVpcAddressPrefixOptionsModel.VpcID = core.StringPtr("testString")
				createVpcAddressPrefixOptionsModel.Cidr = core.StringPtr("10.0.0.0/24")
				createVpcAddressPrefixOptionsModel.Zone = zoneIdentityModel
				createVpcAddressPrefixOptionsModel.IsDefault = core.BoolPtr(true)
				createVpcAddressPrefixOptionsModel.Name = core.StringPtr("my-address-prefix-2")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateVpcAddressPrefix(createVpcAddressPrefixOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVpcAddressPrefix(deleteVpcAddressPrefixOptions *DeleteVpcAddressPrefixOptions)`, func() {
		version := "testString"
		deleteVpcAddressPrefixPath := "/vpcs/testString/address_prefixes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVpcAddressPrefixPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteVpcAddressPrefix successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVpcAddressPrefix(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVpcAddressPrefixOptions model
				deleteVpcAddressPrefixOptionsModel := new(vpcclassicv1.DeleteVpcAddressPrefixOptions)
				deleteVpcAddressPrefixOptionsModel.VpcID = core.StringPtr("testString")
				deleteVpcAddressPrefixOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVpcAddressPrefix(deleteVpcAddressPrefixOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpcAddressPrefix(getVpcAddressPrefixOptions *GetVpcAddressPrefixOptions)`, func() {
		version := "testString"
		getVpcAddressPrefixPath := "/vpcs/testString/address_prefixes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpcAddressPrefixPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"cidr": "192.168.3.0/24", "created_at": "2019-01-01T12:00:00", "has_subnets": true, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/address_prefixes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "is_default": false, "name": "my-address-prefix-2", "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetVpcAddressPrefix successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVpcAddressPrefix(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVpcAddressPrefixOptions model
				getVpcAddressPrefixOptionsModel := new(vpcclassicv1.GetVpcAddressPrefixOptions)
				getVpcAddressPrefixOptionsModel.VpcID = core.StringPtr("testString")
				getVpcAddressPrefixOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVpcAddressPrefix(getVpcAddressPrefixOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVpcAddressPrefix(updateVpcAddressPrefixOptions *UpdateVpcAddressPrefixOptions)`, func() {
		version := "testString"
		updateVpcAddressPrefixPath := "/vpcs/testString/address_prefixes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVpcAddressPrefixPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"cidr": "192.168.3.0/24", "created_at": "2019-01-01T12:00:00", "has_subnets": true, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/address_prefixes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "is_default": false, "name": "my-address-prefix-2", "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke UpdateVpcAddressPrefix successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateVpcAddressPrefix(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVpcAddressPrefixOptions model
				updateVpcAddressPrefixOptionsModel := new(vpcclassicv1.UpdateVpcAddressPrefixOptions)
				updateVpcAddressPrefixOptionsModel.VpcID = core.StringPtr("testString")
				updateVpcAddressPrefixOptionsModel.ID = core.StringPtr("testString")
				updateVpcAddressPrefixOptionsModel.IsDefault = core.BoolPtr(false)
				updateVpcAddressPrefixOptionsModel.Name = core.StringPtr("my-address-prefix-2")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateVpcAddressPrefix(updateVpcAddressPrefixOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpcRoutes(listVpcRoutesOptions *ListVpcRoutesOptions)`, func() {
		version := "testString"
		listVpcRoutesPath := "/vpcs/testString/routes"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpcRoutesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["zone.name"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"routes": [{"created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/routes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "lifecycle_state": "stable", "name": "my-route-1", "next_hop": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}]}`)
			}))
			It(`Invoke ListVpcRoutes successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpcRoutes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpcRoutesOptions model
				listVpcRoutesOptionsModel := new(vpcclassicv1.ListVpcRoutesOptions)
				listVpcRoutesOptionsModel.VpcID = core.StringPtr("testString")
				listVpcRoutesOptionsModel.ZoneName = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpcRoutes(listVpcRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVpcRoute(createVpcRouteOptions *CreateVpcRouteOptions)`, func() {
		version := "testString"
		createVpcRoutePath := "/vpcs/testString/routes"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVpcRoutePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/routes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "lifecycle_state": "stable", "name": "my-route-1", "next_hop": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke CreateVpcRoute successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateVpcRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RouteNextHopPrototypeRouteNextHopIP model
				routeNextHopPrototypeModel := new(vpcclassicv1.RouteNextHopPrototypeRouteNextHopIP)
				routeNextHopPrototypeModel.Address = core.StringPtr("192.168.3.4")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(vpcclassicv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreateVpcRouteOptions model
				createVpcRouteOptionsModel := new(vpcclassicv1.CreateVpcRouteOptions)
				createVpcRouteOptionsModel.VpcID = core.StringPtr("testString")
				createVpcRouteOptionsModel.Destination = core.StringPtr("192.168.3.0/24")
				createVpcRouteOptionsModel.Zone = zoneIdentityModel
				createVpcRouteOptionsModel.Name = core.StringPtr("my-route-2")
				createVpcRouteOptionsModel.NextHop = routeNextHopPrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateVpcRoute(createVpcRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVpcRoute(deleteVpcRouteOptions *DeleteVpcRouteOptions)`, func() {
		version := "testString"
		deleteVpcRoutePath := "/vpcs/testString/routes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVpcRoutePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteVpcRoute successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVpcRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVpcRouteOptions model
				deleteVpcRouteOptionsModel := new(vpcclassicv1.DeleteVpcRouteOptions)
				deleteVpcRouteOptionsModel.VpcID = core.StringPtr("testString")
				deleteVpcRouteOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVpcRoute(deleteVpcRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpcRoute(getVpcRouteOptions *GetVpcRouteOptions)`, func() {
		version := "testString"
		getVpcRoutePath := "/vpcs/testString/routes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpcRoutePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/routes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "lifecycle_state": "stable", "name": "my-route-1", "next_hop": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetVpcRoute successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVpcRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVpcRouteOptions model
				getVpcRouteOptionsModel := new(vpcclassicv1.GetVpcRouteOptions)
				getVpcRouteOptionsModel.VpcID = core.StringPtr("testString")
				getVpcRouteOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVpcRoute(getVpcRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVpcRoute(updateVpcRouteOptions *UpdateVpcRouteOptions)`, func() {
		version := "testString"
		updateVpcRoutePath := "/vpcs/testString/routes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVpcRoutePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "destination": "192.168.3.0/24", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/8e454ead-0db7-48ac-9a8b-2698d8c470a7/routes/1a15dca5-7e33-45e1-b7c5-bc690e569531", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "lifecycle_state": "stable", "name": "my-route-1", "next_hop": {"address": "192.168.3.4"}, "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke UpdateVpcRoute successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateVpcRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVpcRouteOptions model
				updateVpcRouteOptionsModel := new(vpcclassicv1.UpdateVpcRouteOptions)
				updateVpcRouteOptionsModel.VpcID = core.StringPtr("testString")
				updateVpcRouteOptionsModel.ID = core.StringPtr("testString")
				updateVpcRouteOptionsModel.Name = core.StringPtr("my-route-2")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateVpcRoute(updateVpcRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListIkePolicies(listIkePoliciesOptions *ListIkePoliciesOptions)`, func() {
		version := "testString"
		listIkePoliciesPath := "/ike_policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listIkePoliciesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/ike_policies?limit=20"}, "ike_policies": [{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "dh_group": 2, "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ike_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "ike_version": 1, "key_lifetime": 28800, "name": "my-ike-policy", "negotiation_mode": "main", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}}], "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/ike_policies?start=9d5a91a3e2cbd233b5a5b33436855ed&limit=20"}, "total_count": 132}`)
			}))
			It(`Invoke ListIkePolicies successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListIkePolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListIkePoliciesOptions model
				listIkePoliciesOptionsModel := new(vpcclassicv1.ListIkePoliciesOptions)
				listIkePoliciesOptionsModel.Start = core.StringPtr("testString")
				listIkePoliciesOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListIkePolicies(listIkePoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateIkePolicy(createIkePolicyOptions *CreateIkePolicyOptions)`, func() {
		version := "testString"
		createIkePolicyPath := "/ike_policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createIkePolicyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "dh_group": 2, "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ike_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "ike_version": 1, "key_lifetime": 28800, "name": "my-ike-policy", "negotiation_mode": "main", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}}`)
			}))
			It(`Invoke CreateIkePolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateIkePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the CreateIkePolicyOptions model
				createIkePolicyOptionsModel := new(vpcclassicv1.CreateIkePolicyOptions)
				createIkePolicyOptionsModel.AuthenticationAlgorithm = core.StringPtr("md5")
				createIkePolicyOptionsModel.DhGroup = core.Int64Ptr(int64(2))
				createIkePolicyOptionsModel.EncryptionAlgorithm = core.StringPtr("triple_des")
				createIkePolicyOptionsModel.IkeVersion = core.Int64Ptr(int64(1))
				createIkePolicyOptionsModel.KeyLifetime = core.Int64Ptr(int64(28800))
				createIkePolicyOptionsModel.Name = core.StringPtr("my-ike-policy")
				createIkePolicyOptionsModel.ResourceGroup = resourceGroupIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateIkePolicy(createIkePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteIkePolicy(deleteIkePolicyOptions *DeleteIkePolicyOptions)`, func() {
		version := "testString"
		deleteIkePolicyPath := "/ike_policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteIkePolicyPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteIkePolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteIkePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteIkePolicyOptions model
				deleteIkePolicyOptionsModel := new(vpcclassicv1.DeleteIkePolicyOptions)
				deleteIkePolicyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteIkePolicy(deleteIkePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetIkePolicy(getIkePolicyOptions *GetIkePolicyOptions)`, func() {
		version := "testString"
		getIkePolicyPath := "/ike_policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getIkePolicyPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "dh_group": 2, "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ike_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "ike_version": 1, "key_lifetime": 28800, "name": "my-ike-policy", "negotiation_mode": "main", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}}`)
			}))
			It(`Invoke GetIkePolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetIkePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIkePolicyOptions model
				getIkePolicyOptionsModel := new(vpcclassicv1.GetIkePolicyOptions)
				getIkePolicyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetIkePolicy(getIkePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateIkePolicy(updateIkePolicyOptions *UpdateIkePolicyOptions)`, func() {
		version := "testString"
		updateIkePolicyPath := "/ike_policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateIkePolicyPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "dh_group": 2, "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ike_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "ike_version": 1, "key_lifetime": 28800, "name": "my-ike-policy", "negotiation_mode": "main", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}}`)
			}))
			It(`Invoke UpdateIkePolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateIkePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateIkePolicyOptions model
				updateIkePolicyOptionsModel := new(vpcclassicv1.UpdateIkePolicyOptions)
				updateIkePolicyOptionsModel.ID = core.StringPtr("testString")
				updateIkePolicyOptionsModel.AuthenticationAlgorithm = core.StringPtr("md5")
				updateIkePolicyOptionsModel.DhGroup = core.Int64Ptr(int64(2))
				updateIkePolicyOptionsModel.EncryptionAlgorithm = core.StringPtr("triple_des")
				updateIkePolicyOptionsModel.IkeVersion = core.Int64Ptr(int64(1))
				updateIkePolicyOptionsModel.KeyLifetime = core.Int64Ptr(int64(28800))
				updateIkePolicyOptionsModel.Name = core.StringPtr("my-ike-policy")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateIkePolicy(updateIkePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpnGatewayIkePolicyConnections(listVpnGatewayIkePolicyConnectionsOptions *ListVpnGatewayIkePolicyConnectionsOptions)`, func() {
		version := "testString"
		listVpnGatewayIkePolicyConnectionsPath := "/ike_policies/testString/connections"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpnGatewayIkePolicyConnectionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"connections": [{"admin_state_up": true, "authentication_mode": "psk", "created_at": "2019-01-01T12:00:00", "dead_peer_detection": {"action": "restart", "interval": 30, "timeout": 120}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "ike_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "ipsec_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "local_cidrs": ["192.168.1.0/24"], "name": "my-vpn-connection", "peer_address": "169.21.50.5", "peer_cidrs": ["10.45.1.0/24"], "psk": "lkj14b1oi0alcniejkso", "route_mode": "policy", "status": "down"}]}`)
			}))
			It(`Invoke ListVpnGatewayIkePolicyConnections successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpnGatewayIkePolicyConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpnGatewayIkePolicyConnectionsOptions model
				listVpnGatewayIkePolicyConnectionsOptionsModel := new(vpcclassicv1.ListVpnGatewayIkePolicyConnectionsOptions)
				listVpnGatewayIkePolicyConnectionsOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpnGatewayIkePolicyConnections(listVpnGatewayIkePolicyConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListIpsecPolicies(listIpsecPoliciesOptions *ListIpsecPoliciesOptions)`, func() {
		version := "testString"
		listIpsecPoliciesPath := "/ipsec_policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listIpsecPoliciesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies?limit=20"}, "ipsec_policies": [{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "encapsulation_mode": "tunnel", "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "key_lifetime": 3600, "name": "my-ipsec-policy", "pfs": "disabled", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "transform_protocol": "esp"}], "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies?start=9d5a91a3e2cbd233b5a5b33436855ed&limit=20"}, "total_count": 132}`)
			}))
			It(`Invoke ListIpsecPolicies successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListIpsecPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListIpsecPoliciesOptions model
				listIpsecPoliciesOptionsModel := new(vpcclassicv1.ListIpsecPoliciesOptions)
				listIpsecPoliciesOptionsModel.Start = core.StringPtr("testString")
				listIpsecPoliciesOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListIpsecPolicies(listIpsecPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateIpsecPolicy(createIpsecPolicyOptions *CreateIpsecPolicyOptions)`, func() {
		version := "testString"
		createIpsecPolicyPath := "/ipsec_policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createIpsecPolicyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "encapsulation_mode": "tunnel", "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "key_lifetime": 3600, "name": "my-ipsec-policy", "pfs": "disabled", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "transform_protocol": "esp"}`)
			}))
			It(`Invoke CreateIpsecPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateIpsecPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the CreateIpsecPolicyOptions model
				createIpsecPolicyOptionsModel := new(vpcclassicv1.CreateIpsecPolicyOptions)
				createIpsecPolicyOptionsModel.AuthenticationAlgorithm = core.StringPtr("md5")
				createIpsecPolicyOptionsModel.EncryptionAlgorithm = core.StringPtr("triple_des")
				createIpsecPolicyOptionsModel.Pfs = core.StringPtr("disabled")
				createIpsecPolicyOptionsModel.KeyLifetime = core.Int64Ptr(int64(3600))
				createIpsecPolicyOptionsModel.Name = core.StringPtr("my-ipsec-policy")
				createIpsecPolicyOptionsModel.ResourceGroup = resourceGroupIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateIpsecPolicy(createIpsecPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteIpsecPolicy(deleteIpsecPolicyOptions *DeleteIpsecPolicyOptions)`, func() {
		version := "testString"
		deleteIpsecPolicyPath := "/ipsec_policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteIpsecPolicyPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteIpsecPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteIpsecPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteIpsecPolicyOptions model
				deleteIpsecPolicyOptionsModel := new(vpcclassicv1.DeleteIpsecPolicyOptions)
				deleteIpsecPolicyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteIpsecPolicy(deleteIpsecPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetIpsecPolicy(getIpsecPolicyOptions *GetIpsecPolicyOptions)`, func() {
		version := "testString"
		getIpsecPolicyPath := "/ipsec_policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getIpsecPolicyPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "encapsulation_mode": "tunnel", "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "key_lifetime": 3600, "name": "my-ipsec-policy", "pfs": "disabled", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "transform_protocol": "esp"}`)
			}))
			It(`Invoke GetIpsecPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetIpsecPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIpsecPolicyOptions model
				getIpsecPolicyOptionsModel := new(vpcclassicv1.GetIpsecPolicyOptions)
				getIpsecPolicyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetIpsecPolicy(getIpsecPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateIpsecPolicy(updateIpsecPolicyOptions *UpdateIpsecPolicyOptions)`, func() {
		version := "testString"
		updateIpsecPolicyPath := "/ipsec_policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateIpsecPolicyPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"authentication_algorithm": "md5", "connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "encapsulation_mode": "tunnel", "encryption_algorithm": "triple_des", "href": "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "key_lifetime": 3600, "name": "my-ipsec-policy", "pfs": "disabled", "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "transform_protocol": "esp"}`)
			}))
			It(`Invoke UpdateIpsecPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateIpsecPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateIpsecPolicyOptions model
				updateIpsecPolicyOptionsModel := new(vpcclassicv1.UpdateIpsecPolicyOptions)
				updateIpsecPolicyOptionsModel.ID = core.StringPtr("testString")
				updateIpsecPolicyOptionsModel.AuthenticationAlgorithm = core.StringPtr("md5")
				updateIpsecPolicyOptionsModel.EncryptionAlgorithm = core.StringPtr("triple_des")
				updateIpsecPolicyOptionsModel.KeyLifetime = core.Int64Ptr(int64(3600))
				updateIpsecPolicyOptionsModel.Name = core.StringPtr("my-ipsec-policy")
				updateIpsecPolicyOptionsModel.Pfs = core.StringPtr("disabled")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateIpsecPolicy(updateIpsecPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpnGatewayIpsecPolicyConnections(listVpnGatewayIpsecPolicyConnectionsOptions *ListVpnGatewayIpsecPolicyConnectionsOptions)`, func() {
		version := "testString"
		listVpnGatewayIpsecPolicyConnectionsPath := "/ipsec_policies/testString/connections"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpnGatewayIpsecPolicyConnectionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"connections": [{"admin_state_up": true, "authentication_mode": "psk", "created_at": "2019-01-01T12:00:00", "dead_peer_detection": {"action": "restart", "interval": 30, "timeout": 120}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "ike_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "ipsec_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "local_cidrs": ["192.168.1.0/24"], "name": "my-vpn-connection", "peer_address": "169.21.50.5", "peer_cidrs": ["10.45.1.0/24"], "psk": "lkj14b1oi0alcniejkso", "route_mode": "policy", "status": "down"}]}`)
			}))
			It(`Invoke ListVpnGatewayIpsecPolicyConnections successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpnGatewayIpsecPolicyConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpnGatewayIpsecPolicyConnectionsOptions model
				listVpnGatewayIpsecPolicyConnectionsOptionsModel := new(vpcclassicv1.ListVpnGatewayIpsecPolicyConnectionsOptions)
				listVpnGatewayIpsecPolicyConnectionsOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpnGatewayIpsecPolicyConnections(listVpnGatewayIpsecPolicyConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpnGateways(listVpnGatewaysOptions *ListVpnGatewaysOptions)`, func() {
		version := "testString"
		listVpnGatewaysPath := "/vpn_gateways"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpnGatewaysPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["resource_group.id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways?start=9d5a91a3e2cbd233b5a5b33436855ed&limit=20"}, "total_count": 132, "vpn_gateways": [{"connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpn:ddf51bec-3424-11e8-b467-0ed5f89f718b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "name": "my-vpn-gateway", "public_ip": {"address": "192.168.3.4"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}]}`)
			}))
			It(`Invoke ListVpnGateways successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpnGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpnGatewaysOptions model
				listVpnGatewaysOptionsModel := new(vpcclassicv1.ListVpnGatewaysOptions)
				listVpnGatewaysOptionsModel.Start = core.StringPtr("testString")
				listVpnGatewaysOptionsModel.Limit = core.Int64Ptr(int64(38))
				listVpnGatewaysOptionsModel.ResourceGroupID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpnGateways(listVpnGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVpnGateway(createVpnGatewayOptions *CreateVpnGatewayOptions)`, func() {
		version := "testString"
		createVpnGatewayPath := "/vpn_gateways"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVpnGatewayPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpn:ddf51bec-3424-11e8-b467-0ed5f89f718b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "name": "my-vpn-gateway", "public_ip": {"address": "192.168.3.4"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}`)
			}))
			It(`Invoke CreateVpnGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateVpnGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the SubnetIdentityByID model
				subnetIdentityModel := new(vpcclassicv1.SubnetIdentityByID)
				subnetIdentityModel.ID = core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

				// Construct an instance of the CreateVpnGatewayOptions model
				createVpnGatewayOptionsModel := new(vpcclassicv1.CreateVpnGatewayOptions)
				createVpnGatewayOptionsModel.Subnet = subnetIdentityModel
				createVpnGatewayOptionsModel.Name = core.StringPtr("my-vpn-gateway")
				createVpnGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateVpnGateway(createVpnGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVpnGateway(deleteVpnGatewayOptions *DeleteVpnGatewayOptions)`, func() {
		version := "testString"
		deleteVpnGatewayPath := "/vpn_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVpnGatewayPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(202)
			}))
			It(`Invoke DeleteVpnGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVpnGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVpnGatewayOptions model
				deleteVpnGatewayOptionsModel := new(vpcclassicv1.DeleteVpnGatewayOptions)
				deleteVpnGatewayOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVpnGateway(deleteVpnGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpnGateway(getVpnGatewayOptions *GetVpnGatewayOptions)`, func() {
		version := "testString"
		getVpnGatewayPath := "/vpn_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpnGatewayPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpn:ddf51bec-3424-11e8-b467-0ed5f89f718b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "name": "my-vpn-gateway", "public_ip": {"address": "192.168.3.4"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}`)
			}))
			It(`Invoke GetVpnGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVpnGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVpnGatewayOptions model
				getVpnGatewayOptionsModel := new(vpcclassicv1.GetVpnGatewayOptions)
				getVpnGatewayOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVpnGateway(getVpnGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVpnGateway(updateVpnGatewayOptions *UpdateVpnGatewayOptions)`, func() {
		version := "testString"
		updateVpnGatewayPath := "/vpn_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVpnGatewayPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"connections": [{"href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "name": "my-vpn-connection"}], "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south:a/123456::vpn:ddf51bec-3424-11e8-b467-0ed5f89f718b", "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b", "id": "ddf51bec-3424-11e8-b467-0ed5f89f718b", "name": "my-vpn-gateway", "public_ip": {"address": "192.168.3.4"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "subnet": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "id": "7ec86020-1c6e-4889-b3f0-a15f2e50f87e", "name": "my-subnet"}}`)
			}))
			It(`Invoke UpdateVpnGateway successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateVpnGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVpnGatewayOptions model
				updateVpnGatewayOptionsModel := new(vpcclassicv1.UpdateVpnGatewayOptions)
				updateVpnGatewayOptionsModel.ID = core.StringPtr("testString")
				updateVpnGatewayOptionsModel.Name = core.StringPtr("my-vpn-gateway")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateVpnGateway(updateVpnGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpnGatewayConnections(listVpnGatewayConnectionsOptions *ListVpnGatewayConnectionsOptions)`, func() {
		version := "testString"
		listVpnGatewayConnectionsPath := "/vpn_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpnGatewayConnectionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["status"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"connections": [{"admin_state_up": true, "authentication_mode": "psk", "created_at": "2019-01-01T12:00:00", "dead_peer_detection": {"action": "restart", "interval": 30, "timeout": 120}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "ike_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "ipsec_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "local_cidrs": ["192.168.1.0/24"], "name": "my-vpn-connection", "peer_address": "169.21.50.5", "peer_cidrs": ["10.45.1.0/24"], "psk": "lkj14b1oi0alcniejkso", "route_mode": "policy", "status": "down"}]}`)
			}))
			It(`Invoke ListVpnGatewayConnections successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpnGatewayConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpnGatewayConnectionsOptions model
				listVpnGatewayConnectionsOptionsModel := new(vpcclassicv1.ListVpnGatewayConnectionsOptions)
				listVpnGatewayConnectionsOptionsModel.VpnGatewayID = core.StringPtr("testString")
				listVpnGatewayConnectionsOptionsModel.Status = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpnGatewayConnections(listVpnGatewayConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVpnGatewayConnection(createVpnGatewayConnectionOptions *CreateVpnGatewayConnectionOptions)`, func() {
		version := "testString"
		createVpnGatewayConnectionPath := "/vpn_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVpnGatewayConnectionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"admin_state_up": true, "authentication_mode": "psk", "created_at": "2019-01-01T12:00:00", "dead_peer_detection": {"action": "restart", "interval": 30, "timeout": 120}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "ike_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "ipsec_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "local_cidrs": ["192.168.1.0/24"], "name": "my-vpn-connection", "peer_address": "169.21.50.5", "peer_cidrs": ["10.45.1.0/24"], "psk": "lkj14b1oi0alcniejkso", "route_mode": "policy", "status": "down"}`)
			}))
			It(`Invoke CreateVpnGatewayConnection successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateVpnGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IKEPolicyIdentityByID model
				ikePolicyIdentityModel := new(vpcclassicv1.IKEPolicyIdentityByID)
				ikePolicyIdentityModel.ID = core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b")

				// Construct an instance of the IPsecPolicyIdentityByID model
				iPsecPolicyIdentityModel := new(vpcclassicv1.IPsecPolicyIdentityByID)
				iPsecPolicyIdentityModel.ID = core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b")

				// Construct an instance of the VPNGatewayConnectionDPDPrototype model
				vpnGatewayConnectionDpdPrototypeModel := new(vpcclassicv1.VPNGatewayConnectionDPDPrototype)
				vpnGatewayConnectionDpdPrototypeModel.Action = core.StringPtr("restart")
				vpnGatewayConnectionDpdPrototypeModel.Interval = core.Int64Ptr(int64(30))
				vpnGatewayConnectionDpdPrototypeModel.Timeout = core.Int64Ptr(int64(120))

				// Construct an instance of the CreateVpnGatewayConnectionOptions model
				createVpnGatewayConnectionOptionsModel := new(vpcclassicv1.CreateVpnGatewayConnectionOptions)
				createVpnGatewayConnectionOptionsModel.VpnGatewayID = core.StringPtr("testString")
				createVpnGatewayConnectionOptionsModel.PeerAddress = core.StringPtr("169.21.50.5")
				createVpnGatewayConnectionOptionsModel.Psk = core.StringPtr("lkj14b1oi0alcniejkso")
				createVpnGatewayConnectionOptionsModel.AdminStateUp = core.BoolPtr(true)
				createVpnGatewayConnectionOptionsModel.DeadPeerDetection = vpnGatewayConnectionDpdPrototypeModel
				createVpnGatewayConnectionOptionsModel.IkePolicy = ikePolicyIdentityModel
				createVpnGatewayConnectionOptionsModel.IpsecPolicy = iPsecPolicyIdentityModel
				createVpnGatewayConnectionOptionsModel.LocalCidrs = []string{"192.168.1.0/24"}
				createVpnGatewayConnectionOptionsModel.Name = core.StringPtr("my-vpn-connection")
				createVpnGatewayConnectionOptionsModel.PeerCidrs = []string{"10.45.1.0/24"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateVpnGatewayConnection(createVpnGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVpnGatewayConnection(deleteVpnGatewayConnectionOptions *DeleteVpnGatewayConnectionOptions)`, func() {
		version := "testString"
		deleteVpnGatewayConnectionPath := "/vpn_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVpnGatewayConnectionPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(202)
			}))
			It(`Invoke DeleteVpnGatewayConnection successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVpnGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVpnGatewayConnectionOptions model
				deleteVpnGatewayConnectionOptionsModel := new(vpcclassicv1.DeleteVpnGatewayConnectionOptions)
				deleteVpnGatewayConnectionOptionsModel.VpnGatewayID = core.StringPtr("testString")
				deleteVpnGatewayConnectionOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVpnGatewayConnection(deleteVpnGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpnGatewayConnection(getVpnGatewayConnectionOptions *GetVpnGatewayConnectionOptions)`, func() {
		version := "testString"
		getVpnGatewayConnectionPath := "/vpn_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpnGatewayConnectionPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"admin_state_up": true, "authentication_mode": "psk", "created_at": "2019-01-01T12:00:00", "dead_peer_detection": {"action": "restart", "interval": 30, "timeout": 120}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "ike_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "ipsec_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "local_cidrs": ["192.168.1.0/24"], "name": "my-vpn-connection", "peer_address": "169.21.50.5", "peer_cidrs": ["10.45.1.0/24"], "psk": "lkj14b1oi0alcniejkso", "route_mode": "policy", "status": "down"}`)
			}))
			It(`Invoke GetVpnGatewayConnection successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVpnGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVpnGatewayConnectionOptions model
				getVpnGatewayConnectionOptionsModel := new(vpcclassicv1.GetVpnGatewayConnectionOptions)
				getVpnGatewayConnectionOptionsModel.VpnGatewayID = core.StringPtr("testString")
				getVpnGatewayConnectionOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVpnGatewayConnection(getVpnGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVpnGatewayConnection(updateVpnGatewayConnectionOptions *UpdateVpnGatewayConnectionOptions)`, func() {
		version := "testString"
		updateVpnGatewayConnectionPath := "/vpn_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVpnGatewayConnectionPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"admin_state_up": true, "authentication_mode": "psk", "created_at": "2019-01-01T12:00:00", "dead_peer_detection": {"action": "restart", "interval": 30, "timeout": 120}, "href": "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/93487806-7743-4c46-81d6-72869883ea0b", "id": "a10a5771-dc23-442c-8460-c3601d8542f7", "ike_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "ipsec_policy": {"id": "ddf51bec-3424-11e8-b467-0ed5f89f718b"}, "local_cidrs": ["192.168.1.0/24"], "name": "my-vpn-connection", "peer_address": "169.21.50.5", "peer_cidrs": ["10.45.1.0/24"], "psk": "lkj14b1oi0alcniejkso", "route_mode": "policy", "status": "down"}`)
			}))
			It(`Invoke UpdateVpnGatewayConnection successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateVpnGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IKEPolicyIdentityByID model
				ikePolicyIdentityModel := new(vpcclassicv1.IKEPolicyIdentityByID)
				ikePolicyIdentityModel.ID = core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b")

				// Construct an instance of the IPsecPolicyIdentityByID model
				iPsecPolicyIdentityModel := new(vpcclassicv1.IPsecPolicyIdentityByID)
				iPsecPolicyIdentityModel.ID = core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b")

				// Construct an instance of the VPNGatewayConnectionDPDPrototype model
				vpnGatewayConnectionDpdPrototypeModel := new(vpcclassicv1.VPNGatewayConnectionDPDPrototype)
				vpnGatewayConnectionDpdPrototypeModel.Action = core.StringPtr("restart")
				vpnGatewayConnectionDpdPrototypeModel.Interval = core.Int64Ptr(int64(30))
				vpnGatewayConnectionDpdPrototypeModel.Timeout = core.Int64Ptr(int64(120))

				// Construct an instance of the UpdateVpnGatewayConnectionOptions model
				updateVpnGatewayConnectionOptionsModel := new(vpcclassicv1.UpdateVpnGatewayConnectionOptions)
				updateVpnGatewayConnectionOptionsModel.VpnGatewayID = core.StringPtr("testString")
				updateVpnGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateVpnGatewayConnectionOptionsModel.AdminStateUp = core.BoolPtr(true)
				updateVpnGatewayConnectionOptionsModel.DeadPeerDetection = vpnGatewayConnectionDpdPrototypeModel
				updateVpnGatewayConnectionOptionsModel.IkePolicy = ikePolicyIdentityModel
				updateVpnGatewayConnectionOptionsModel.IpsecPolicy = iPsecPolicyIdentityModel
				updateVpnGatewayConnectionOptionsModel.Name = core.StringPtr("my-vpn-connection")
				updateVpnGatewayConnectionOptionsModel.PeerAddress = core.StringPtr("169.21.50.5")
				updateVpnGatewayConnectionOptionsModel.Psk = core.StringPtr("lkj14b1oi0alcniejkso")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateVpnGatewayConnection(updateVpnGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpnGatewayConnectionLocalCidrs(listVpnGatewayConnectionLocalCidrsOptions *ListVpnGatewayConnectionLocalCidrsOptions)`, func() {
		version := "testString"
		listVpnGatewayConnectionLocalCidrsPath := "/vpn_gateways/testString/connections/testString/local_cidrs"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpnGatewayConnectionLocalCidrsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"local_cidrs": ["192.168.1.0/24"]}`)
			}))
			It(`Invoke ListVpnGatewayConnectionLocalCidrs successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpnGatewayConnectionLocalCidrs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpnGatewayConnectionLocalCidrsOptions model
				listVpnGatewayConnectionLocalCidrsOptionsModel := new(vpcclassicv1.ListVpnGatewayConnectionLocalCidrsOptions)
				listVpnGatewayConnectionLocalCidrsOptionsModel.VpnGatewayID = core.StringPtr("testString")
				listVpnGatewayConnectionLocalCidrsOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpnGatewayConnectionLocalCidrs(listVpnGatewayConnectionLocalCidrsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVpnGatewayConnectionLocalCidr(deleteVpnGatewayConnectionLocalCidrOptions *DeleteVpnGatewayConnectionLocalCidrOptions)`, func() {
		version := "testString"
		deleteVpnGatewayConnectionLocalCidrPath := "/vpn_gateways/testString/connections/testString/local_cidrs/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVpnGatewayConnectionLocalCidrPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteVpnGatewayConnectionLocalCidr successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVpnGatewayConnectionLocalCidr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVpnGatewayConnectionLocalCidrOptions model
				deleteVpnGatewayConnectionLocalCidrOptionsModel := new(vpcclassicv1.DeleteVpnGatewayConnectionLocalCidrOptions)
				deleteVpnGatewayConnectionLocalCidrOptionsModel.VpnGatewayID = core.StringPtr("testString")
				deleteVpnGatewayConnectionLocalCidrOptionsModel.ID = core.StringPtr("testString")
				deleteVpnGatewayConnectionLocalCidrOptionsModel.PrefixAddress = core.StringPtr("testString")
				deleteVpnGatewayConnectionLocalCidrOptionsModel.PrefixLength = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVpnGatewayConnectionLocalCidr(deleteVpnGatewayConnectionLocalCidrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpnGatewayConnectionLocalCidr(getVpnGatewayConnectionLocalCidrOptions *GetVpnGatewayConnectionLocalCidrOptions)`, func() {
		version := "testString"
		getVpnGatewayConnectionLocalCidrPath := "/vpn_gateways/testString/connections/testString/local_cidrs/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpnGatewayConnectionLocalCidrPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke GetVpnGatewayConnectionLocalCidr successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetVpnGatewayConnectionLocalCidr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetVpnGatewayConnectionLocalCidrOptions model
				getVpnGatewayConnectionLocalCidrOptionsModel := new(vpcclassicv1.GetVpnGatewayConnectionLocalCidrOptions)
				getVpnGatewayConnectionLocalCidrOptionsModel.VpnGatewayID = core.StringPtr("testString")
				getVpnGatewayConnectionLocalCidrOptionsModel.ID = core.StringPtr("testString")
				getVpnGatewayConnectionLocalCidrOptionsModel.PrefixAddress = core.StringPtr("testString")
				getVpnGatewayConnectionLocalCidrOptionsModel.PrefixLength = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetVpnGatewayConnectionLocalCidr(getVpnGatewayConnectionLocalCidrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`SetVpnGatewayConnectionLocalCidr(setVpnGatewayConnectionLocalCidrOptions *SetVpnGatewayConnectionLocalCidrOptions)`, func() {
		version := "testString"
		setVpnGatewayConnectionLocalCidrPath := "/vpn_gateways/testString/connections/testString/local_cidrs/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(setVpnGatewayConnectionLocalCidrPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke SetVpnGatewayConnectionLocalCidr successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.SetVpnGatewayConnectionLocalCidr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SetVpnGatewayConnectionLocalCidrOptions model
				setVpnGatewayConnectionLocalCidrOptionsModel := new(vpcclassicv1.SetVpnGatewayConnectionLocalCidrOptions)
				setVpnGatewayConnectionLocalCidrOptionsModel.VpnGatewayID = core.StringPtr("testString")
				setVpnGatewayConnectionLocalCidrOptionsModel.ID = core.StringPtr("testString")
				setVpnGatewayConnectionLocalCidrOptionsModel.PrefixAddress = core.StringPtr("testString")
				setVpnGatewayConnectionLocalCidrOptionsModel.PrefixLength = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.SetVpnGatewayConnectionLocalCidr(setVpnGatewayConnectionLocalCidrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVpnGatewayConnectionPeerCidrs(listVpnGatewayConnectionPeerCidrsOptions *ListVpnGatewayConnectionPeerCidrsOptions)`, func() {
		version := "testString"
		listVpnGatewayConnectionPeerCidrsPath := "/vpn_gateways/testString/connections/testString/peer_cidrs"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVpnGatewayConnectionPeerCidrsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"peer_cidrs": ["10.45.1.0/24"]}`)
			}))
			It(`Invoke ListVpnGatewayConnectionPeerCidrs successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVpnGatewayConnectionPeerCidrs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVpnGatewayConnectionPeerCidrsOptions model
				listVpnGatewayConnectionPeerCidrsOptionsModel := new(vpcclassicv1.ListVpnGatewayConnectionPeerCidrsOptions)
				listVpnGatewayConnectionPeerCidrsOptionsModel.VpnGatewayID = core.StringPtr("testString")
				listVpnGatewayConnectionPeerCidrsOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVpnGatewayConnectionPeerCidrs(listVpnGatewayConnectionPeerCidrsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVpnGatewayConnectionPeerCidr(deleteVpnGatewayConnectionPeerCidrOptions *DeleteVpnGatewayConnectionPeerCidrOptions)`, func() {
		version := "testString"
		deleteVpnGatewayConnectionPeerCidrPath := "/vpn_gateways/testString/connections/testString/peer_cidrs/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVpnGatewayConnectionPeerCidrPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteVpnGatewayConnectionPeerCidr successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVpnGatewayConnectionPeerCidr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVpnGatewayConnectionPeerCidrOptions model
				deleteVpnGatewayConnectionPeerCidrOptionsModel := new(vpcclassicv1.DeleteVpnGatewayConnectionPeerCidrOptions)
				deleteVpnGatewayConnectionPeerCidrOptionsModel.VpnGatewayID = core.StringPtr("testString")
				deleteVpnGatewayConnectionPeerCidrOptionsModel.ID = core.StringPtr("testString")
				deleteVpnGatewayConnectionPeerCidrOptionsModel.PrefixAddress = core.StringPtr("testString")
				deleteVpnGatewayConnectionPeerCidrOptionsModel.PrefixLength = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVpnGatewayConnectionPeerCidr(deleteVpnGatewayConnectionPeerCidrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVpnGatewayConnectionPeerCidr(getVpnGatewayConnectionPeerCidrOptions *GetVpnGatewayConnectionPeerCidrOptions)`, func() {
		version := "testString"
		getVpnGatewayConnectionPeerCidrPath := "/vpn_gateways/testString/connections/testString/peer_cidrs/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVpnGatewayConnectionPeerCidrPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke GetVpnGatewayConnectionPeerCidr successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetVpnGatewayConnectionPeerCidr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetVpnGatewayConnectionPeerCidrOptions model
				getVpnGatewayConnectionPeerCidrOptionsModel := new(vpcclassicv1.GetVpnGatewayConnectionPeerCidrOptions)
				getVpnGatewayConnectionPeerCidrOptionsModel.VpnGatewayID = core.StringPtr("testString")
				getVpnGatewayConnectionPeerCidrOptionsModel.ID = core.StringPtr("testString")
				getVpnGatewayConnectionPeerCidrOptionsModel.PrefixAddress = core.StringPtr("testString")
				getVpnGatewayConnectionPeerCidrOptionsModel.PrefixLength = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetVpnGatewayConnectionPeerCidr(getVpnGatewayConnectionPeerCidrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`SetVpnGatewayConnectionPeerCidr(setVpnGatewayConnectionPeerCidrOptions *SetVpnGatewayConnectionPeerCidrOptions)`, func() {
		version := "testString"
		setVpnGatewayConnectionPeerCidrPath := "/vpn_gateways/testString/connections/testString/peer_cidrs/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(setVpnGatewayConnectionPeerCidrPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke SetVpnGatewayConnectionPeerCidr successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.SetVpnGatewayConnectionPeerCidr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SetVpnGatewayConnectionPeerCidrOptions model
				setVpnGatewayConnectionPeerCidrOptionsModel := new(vpcclassicv1.SetVpnGatewayConnectionPeerCidrOptions)
				setVpnGatewayConnectionPeerCidrOptionsModel.VpnGatewayID = core.StringPtr("testString")
				setVpnGatewayConnectionPeerCidrOptionsModel.ID = core.StringPtr("testString")
				setVpnGatewayConnectionPeerCidrOptionsModel.PrefixAddress = core.StringPtr("testString")
				setVpnGatewayConnectionPeerCidrOptionsModel.PrefixLength = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.SetVpnGatewayConnectionPeerCidr(setVpnGatewayConnectionPeerCidrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVolumeProfiles(listVolumeProfilesOptions *ListVolumeProfilesOptions)`, func() {
		version := "testString"
		listVolumeProfilesPath := "/volume/profiles"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVolumeProfilesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "profiles": [{"family": "tiered", "href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose", "name": "general-purpose"}], "total_count": 132}`)
			}))
			It(`Invoke ListVolumeProfiles successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVolumeProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVolumeProfilesOptions model
				listVolumeProfilesOptionsModel := new(vpcclassicv1.ListVolumeProfilesOptions)
				listVolumeProfilesOptionsModel.Start = core.StringPtr("testString")
				listVolumeProfilesOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVolumeProfiles(listVolumeProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVolumeProfile(getVolumeProfileOptions *GetVolumeProfileOptions)`, func() {
		version := "testString"
		getVolumeProfilePath := "/volume/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVolumeProfilePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"family": "tiered", "href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose", "name": "general-purpose"}`)
			}))
			It(`Invoke GetVolumeProfile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVolumeProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVolumeProfileOptions model
				getVolumeProfileOptionsModel := new(vpcclassicv1.GetVolumeProfileOptions)
				getVolumeProfileOptionsModel.Name = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVolumeProfile(getVolumeProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVolumes(listVolumesOptions *ListVolumesOptions)`, func() {
		version := "testString"
		listVolumesPath := "/volumes"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVolumesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["zone.name"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"first": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volumes?limit=20"}, "limit": 20, "next": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volumes?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=20"}, "volumes": [{"capacity": 100, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "encryption": "provider_managed", "encryption_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "iops": 10000, "name": "my-volume", "profile": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose", "name": "general-purpose"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "instance": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "name": "my-instance"}, "name": "my-volume-attachment", "type": "boot"}], "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}]}`)
			}))
			It(`Invoke ListVolumes successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListVolumes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVolumesOptions model
				listVolumesOptionsModel := new(vpcclassicv1.ListVolumesOptions)
				listVolumesOptionsModel.Start = core.StringPtr("testString")
				listVolumesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listVolumesOptionsModel.Name = core.StringPtr("testString")
				listVolumesOptionsModel.ZoneName = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListVolumes(listVolumesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVolume(createVolumeOptions *CreateVolumeOptions)`, func() {
		version := "testString"
		createVolumePath := "/volumes"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVolumePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"capacity": 100, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "encryption": "provider_managed", "encryption_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "iops": 10000, "name": "my-volume", "profile": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose", "name": "general-purpose"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "instance": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "name": "my-instance"}, "name": "my-volume-attachment", "type": "boot"}], "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke CreateVolume successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EncryptionKeyIdentityByCRN model
				encryptionKeyIdentityModel := new(vpcclassicv1.EncryptionKeyIdentityByCRN)
				encryptionKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179")

				// Construct an instance of the ResourceGroupIdentityByID model
				resourceGroupIdentityModel := new(vpcclassicv1.ResourceGroupIdentityByID)
				resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

				// Construct an instance of the VolumeProfileIdentityByName model
				volumeProfileIdentityModel := new(vpcclassicv1.VolumeProfileIdentityByName)
				volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(vpcclassicv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the VolumePrototypeVolumeByCapacity model
				volumePrototypeModel := new(vpcclassicv1.VolumePrototypeVolumeByCapacity)
				volumePrototypeModel.EncryptionKey = encryptionKeyIdentityModel
				volumePrototypeModel.Iops = core.Int64Ptr(int64(10000))
				volumePrototypeModel.Name = core.StringPtr("my-volume")
				volumePrototypeModel.Profile = volumeProfileIdentityModel
				volumePrototypeModel.ResourceGroup = resourceGroupIdentityModel
				volumePrototypeModel.Zone = zoneIdentityModel
				volumePrototypeModel.Capacity = core.Int64Ptr(int64(100))

				// Construct an instance of the CreateVolumeOptions model
				createVolumeOptionsModel := new(vpcclassicv1.CreateVolumeOptions)
				createVolumeOptionsModel.VolumePrototype = volumePrototypeModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateVolume(createVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVolume(deleteVolumeOptions *DeleteVolumeOptions)`, func() {
		version := "testString"
		deleteVolumePath := "/volumes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVolumePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteVolume successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVolumeOptions model
				deleteVolumeOptionsModel := new(vpcclassicv1.DeleteVolumeOptions)
				deleteVolumeOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteVolume(deleteVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVolume(getVolumeOptions *GetVolumeOptions)`, func() {
		version := "testString"
		getVolumePath := "/volumes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVolumePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"capacity": 100, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "encryption": "provider_managed", "encryption_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "iops": 10000, "name": "my-volume", "profile": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose", "name": "general-purpose"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "instance": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "name": "my-instance"}, "name": "my-volume-attachment", "type": "boot"}], "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke GetVolume successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVolumeOptions model
				getVolumeOptionsModel := new(vpcclassicv1.GetVolumeOptions)
				getVolumeOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVolume(getVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVolume(updateVolumeOptions *UpdateVolumeOptions)`, func() {
		version := "testString"
		updateVolumePath := "/volumes/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVolumePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["generation"]).To(Equal([]string{fmt.Sprint(int64(1))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"capacity": 100, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "encryption": "provider_managed", "encryption_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "id": "1a6b7274-678d-4dfb-8981-c71dd9d4daa5", "iops": 10000, "name": "my-volume", "profile": {"href": "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose", "name": "general-purpose"}, "resource_group": {"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345", "id": "fee82deba12e4c0fb69c3b09d1f12345"}, "status": "available", "volume_attachments": [{"device": {"id": "80b3e36e-41f4-40e9-bd56-beae81792a68"}, "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/volume_attachments/82cbf856-9cbb-45fb-b62f-d7bcef32399a", "id": "82cbf856-9cbb-45fb-b62f-d7bcef32399a", "instance": {"crn": "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:1e09281b-f177-46fb-baf1-bc152b2e391a", "href": "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a", "id": "1e09281b-f177-46fb-baf1-bc152b2e391a", "name": "my-instance"}, "name": "my-volume-attachment", "type": "boot"}], "zone": {"href": "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1", "name": "us-south-1"}}`)
			}))
			It(`Invoke UpdateVolume successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateVolume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVolumeOptions model
				updateVolumeOptionsModel := new(vpcclassicv1.UpdateVolumeOptions)
				updateVolumeOptionsModel.ID = core.StringPtr("testString")
				updateVolumeOptionsModel.Name = core.StringPtr("my-volume")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateVolume(updateVolumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			version := "testString"
			testService, _ := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
				URL:           "http://vpcclassicv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewImageFilePrototype successfully`, func() {
				href := "cos://us-south/custom-image-vpc-bucket/customImage-0.vhd"
				model, err := testService.NewImageFilePrototype(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPolicyPrototype successfully`, func() {
				action := "forward"
				priority := int64(5)
				model, err := testService.NewLoadBalancerListenerPolicyPrototype(action, priority)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPolicyRulePrototype successfully`, func() {
				condition := "contains"
				typeVar := "header"
				value := "testString"
				model, err := testService.NewLoadBalancerListenerPolicyRulePrototype(condition, typeVar, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPrototypeLoadBalancerContext successfully`, func() {
				port := int64(443)
				protocol := "http"
				model, err := testService.NewLoadBalancerListenerPrototypeLoadBalancerContext(port, protocol)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolHealthMonitorPatch successfully`, func() {
				delay := int64(5)
				maxRetries := int64(2)
				timeout := int64(2)
				typeVar := "http"
				model, err := testService.NewLoadBalancerPoolHealthMonitorPatch(delay, maxRetries, timeout, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolHealthMonitorPrototype successfully`, func() {
				delay := int64(5)
				maxRetries := int64(2)
				timeout := int64(2)
				typeVar := "http"
				model, err := testService.NewLoadBalancerPoolHealthMonitorPrototype(delay, maxRetries, timeout, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolIdentityByName successfully`, func() {
				name := "my-load-balancer-pool"
				model, err := testService.NewLoadBalancerPoolIdentityByName(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolMemberPrototype successfully`, func() {
				port := int64(80)
				var target vpcclassicv1.LoadBalancerPoolMemberTargetPrototypeIntf = nil
				_, err := testService.NewLoadBalancerPoolMemberPrototype(port, target)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolPrototype successfully`, func() {
				algorithm := "least_connections"
				var healthMonitor *vpcclassicv1.LoadBalancerPoolHealthMonitorPrototype = nil
				protocol := "http"
				_, err := testService.NewLoadBalancerPoolPrototype(algorithm, healthMonitor, protocol)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolSessionPersistencePatch successfully`, func() {
				typeVar := "source_ip"
				model, err := testService.NewLoadBalancerPoolSessionPersistencePatch(typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolSessionPersistencePrototype successfully`, func() {
				typeVar := "source_ip"
				model, err := testService.NewLoadBalancerPoolSessionPersistencePrototype(typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRuleReference successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				name := "my-rule-1"
				model, err := testService.NewNetworkACLRuleReference(href, id, name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkInterfacePrototype successfully`, func() {
				var subnet vpcclassicv1.SubnetIdentityIntf = nil
				_, err := testService.NewNetworkInterfacePrototype(subnet)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewVolumeAttachmentPrototypeInstanceByImageContext successfully`, func() {
				var volume *vpcclassicv1.VolumePrototypeInstanceByImageContext = nil
				_, err := testService.NewVolumeAttachmentPrototypeInstanceByImageContext(volume)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewVolumeAttachmentPrototypeInstanceContext successfully`, func() {
				var volume vpcclassicv1.VolumeAttachmentPrototypeInstanceContextVolumeIntf = nil
				_, err := testService.NewVolumeAttachmentPrototypeInstanceContext(volume)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewVolumePrototypeInstanceByImageContext successfully`, func() {
				var profile vpcclassicv1.VolumeProfileIdentityIntf = nil
				_, err := testService.NewVolumePrototypeInstanceByImageContext(profile)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCertificateInstanceIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758"
				model, err := testService.NewCertificateInstanceIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEncryptionKeyIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
				model, err := testService.NewEncryptionKeyIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFloatingIPPrototypeFloatingIPByTarget successfully`, func() {
				var target vpcclassicv1.NetworkInterfaceIdentityIntf = nil
				_, err := testService.NewFloatingIPPrototypeFloatingIPByTarget(target)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewFloatingIPPrototypeFloatingIPByZone successfully`, func() {
				var zone vpcclassicv1.ZoneIdentityIntf = nil
				_, err := testService.NewFloatingIPPrototypeFloatingIPByZone(zone)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewIKEPolicyIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/ike_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b"
				model, err := testService.NewIKEPolicyIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewIKEPolicyIdentityByID successfully`, func() {
				id := "ddf51bec-3424-11e8-b467-0ed5f89f718b"
				model, err := testService.NewIKEPolicyIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewIPsecPolicyIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies/ddf51bec-3424-11e8-b467-0ed5f89f718b"
				model, err := testService.NewIPsecPolicyIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewIPsecPolicyIdentityByID successfully`, func() {
				id := "ddf51bec-3424-11e8-b467-0ed5f89f718b"
				model, err := testService.NewIPsecPolicyIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImageIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::image:72b27b5c-f4b0-48bb-b954-5becc7c1dcb8"
				model, err := testService.NewImageIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImageIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/images/72b27b5c-f4b0-48bb-b954-5becc7c1dcb8"
				model, err := testService.NewImageIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImageIdentityByID successfully`, func() {
				id := "72b27b5c-f4b0-48bb-b954-5becc7c1dcb8"
				model, err := testService.NewImageIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImagePrototypeImageByFile successfully`, func() {
				var file *vpcclassicv1.ImageFilePrototype = nil
				var operatingSystem vpcclassicv1.OperatingSystemIdentityIntf = nil
				_, err := testService.NewImagePrototypeImageByFile(file, operatingSystem)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewInstanceProfileIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south-1:::instance-profile:bc1-4x16"
				model, err := testService.NewInstanceProfileIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewInstanceProfileIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bc1-4x16"
				model, err := testService.NewInstanceProfileIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewInstanceProfileIdentityByName successfully`, func() {
				name := "bc1-4x16"
				model, err := testService.NewInstanceProfileIdentityByName(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewInstancePrototypeInstanceByImage successfully`, func() {
				var image vpcclassicv1.ImageIdentityIntf = nil
				var primaryNetworkInterface *vpcclassicv1.NetworkInterfacePrototype = nil
				var zone vpcclassicv1.ZoneIdentityIntf = nil
				_, err := testService.NewInstancePrototypeInstanceByImage(image, primaryNetworkInterface, zone)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewKeyIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::key:a6b1a881-2ce8-41a3-80fc-36316a73f803"
				model, err := testService.NewKeyIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewKeyIdentityByFingerprint successfully`, func() {
				fingerprint := "SHA256:yxavE4CIOL2NlsqcurRO3xGjkP6m/0mp8ugojH5yxlY"
				model, err := testService.NewKeyIdentityByFingerprint(fingerprint)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewKeyIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/keys/a6b1a881-2ce8-41a3-80fc-36316a73f803"
				model, err := testService.NewKeyIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewKeyIdentityByID successfully`, func() {
				id := "a6b1a881-2ce8-41a3-80fc-36316a73f803"
				model, err := testService.NewKeyIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPolicyPrototypeTargetLoadBalancerListenerPolicyRedirectURLPrototype successfully`, func() {
				httpStatusCode := int64(301)
				url := "https://www.redirect.com"
				model, err := testService.NewLoadBalancerListenerPolicyPrototypeTargetLoadBalancerListenerPolicyRedirectURLPrototype(httpStatusCode, url)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004"
				model, err := testService.NewLoadBalancerPoolIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolIdentityByID successfully`, func() {
				id := "70294e14-4e61-11e8-bcf4-0242ac110004"
				model, err := testService.NewLoadBalancerPoolIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerPoolMemberTargetPrototypeByAddress successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewLoadBalancerPoolMemberTargetPrototypeByAddress(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf"
				model, err := testService.NewNetworkACLIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLIdentityByID successfully`, func() {
				id := "a4e28308-8ee7-46ab-8108-9f881f22bdbf"
				model, err := testService.NewNetworkACLIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLPrototypeNetworkACLBySourceNetworkACL successfully`, func() {
				var sourceNetworkAcl vpcclassicv1.NetworkACLIdentityIntf = nil
				_, err := testService.NewNetworkACLPrototypeNetworkACLBySourceNetworkACL(sourceNetworkAcl)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewNetworkACLRulePatchBeforeNetworkACLRuleIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				model, err := testService.NewNetworkACLRulePatchBeforeNetworkACLRuleIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePatchBeforeNetworkACLRuleIdentityByID successfully`, func() {
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				model, err := testService.NewNetworkACLRulePatchBeforeNetworkACLRuleIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePatchNetworkACLRuleProtocolAll successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePatchNetworkACLRuleProtocolAll(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePatchNetworkACLRuleProtocolICMP successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePatchNetworkACLRuleProtocolICMP(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePatchNetworkACLRuleProtocolTCPUDP successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePatchNetworkACLRuleProtocolTCPUDP(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeBeforeNetworkACLRuleIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				model, err := testService.NewNetworkACLRulePrototypeBeforeNetworkACLRuleIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeBeforeNetworkACLRuleIdentityByID successfully`, func() {
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				model, err := testService.NewNetworkACLRulePrototypeBeforeNetworkACLRuleIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolAll successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolAll(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolICMP successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolICMP(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolTCPUDP successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolTCPUDP(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeNetworkACLRuleProtocolAll successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePrototypeNetworkACLRuleProtocolAll(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeNetworkACLRuleProtocolICMP successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePrototypeNetworkACLRuleProtocolICMP(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkACLRulePrototypeNetworkACLRuleProtocolTCPUDP successfully`, func() {
				action := "allow"
				createdAt := CreateMockDateTime()
				destination := "192.168.3.0/24"
				direction := "inbound"
				href := "https://us-south.iaas.cloud.ibm.com/v1/network_acls/a4e28308-8ee7-46ab-8108-9f881f22bdbf/rules/8daca77a-4980-4d33-8f3e-7038797be8f9"
				id := "8daca77a-4980-4d33-8f3e-7038797be8f9"
				ipVersion := "ipv4"
				name := "my-rule-2"
				source := "192.168.3.0/24"
				model, err := testService.NewNetworkACLRulePrototypeNetworkACLRuleProtocolTCPUDP(action, createdAt, destination, direction, href, id, ipVersion, name, source)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkInterfaceIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e"
				model, err := testService.NewNetworkInterfaceIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNetworkInterfaceIdentityByID successfully`, func() {
				id := "10c02d81-0ecb-4dc5-897d-28392913b81e"
				model, err := testService.NewNetworkInterfaceIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewOperatingSystemIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-16-amd64"
				model, err := testService.NewOperatingSystemIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewOperatingSystemIdentityByName successfully`, func() {
				name := "ubuntu-16-amd64"
				model, err := testService.NewOperatingSystemIdentityByName(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPublicGatewayIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south-1:a/123456::public-gateway:dc5431ef-1fc6-4861-adc9-a59d077d1241"
				model, err := testService.NewPublicGatewayIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPublicGatewayIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/public_gateways/dc5431ef-1fc6-4861-adc9-a59d077d1241"
				model, err := testService.NewPublicGatewayIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPublicGatewayIdentityByID successfully`, func() {
				id := "dc5431ef-1fc6-4861-adc9-a59d077d1241"
				model, err := testService.NewPublicGatewayIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceGroupIdentityByID successfully`, func() {
				id := "fee82deba12e4c0fb69c3b09d1f12345"
				model, err := testService.NewResourceGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRouteNextHopPrototypeRouteNextHopIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewRouteNextHopPrototypeRouteNextHopIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePatchRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePatchRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePrototypeRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePrototypeRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteCIDR successfully`, func() {
				cidrBlock := "192.168.3.0/24"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteCIDR(cidrBlock)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteIP successfully`, func() {
				address := "192.168.3.4"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteIP(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAll successfully`, func() {
				direction := "inbound"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAll(direction)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMP successfully`, func() {
				direction := "inbound"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMP(direction)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDP successfully`, func() {
				direction := "inbound"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDP(direction)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubnetIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
				model, err := testService.NewSubnetIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubnetIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
				model, err := testService.NewSubnetIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubnetIdentityByID successfully`, func() {
				id := "7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
				model, err := testService.NewSubnetIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubnetPrototypeSubnetByCIDR successfully`, func() {
				var vpc vpcclassicv1.VPCIdentityIntf = nil
				ipv4CidrBlock := "10.0.0.0/24"
				_, err := testService.NewSubnetPrototypeSubnetByCIDR(vpc, ipv4CidrBlock)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewSubnetPrototypeSubnetByTotalCount successfully`, func() {
				var vpc vpcclassicv1.VPCIdentityIntf = nil
				totalIpv4AddressCount := int64(256)
				var zone vpcclassicv1.ZoneIdentityIntf = nil
				_, err := testService.NewSubnetPrototypeSubnetByTotalCount(vpc, totalIpv4AddressCount, zone)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewVPCIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b"
				model, err := testService.NewVPCIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVPCIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/vpcs/4727d842-f94f-4a2d-824a-9bc9b02c523b"
				model, err := testService.NewVPCIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVPCIdentityByID successfully`, func() {
				id := "4727d842-f94f-4a2d-824a-9bc9b02c523b"
				model, err := testService.NewVPCIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south-1:a/123456::volume:1a6b7274-678d-4dfb-8981-c71dd9d4daa5"
				model, err := testService.NewVolumeIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/volumes/1a6b7274-678d-4dfb-8981-c71dd9d4daa5"
				model, err := testService.NewVolumeIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeIdentityByID successfully`, func() {
				id := "1a6b7274-678d-4dfb-8981-c71dd9d4daa5"
				model, err := testService.NewVolumeIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeProfileIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"
				model, err := testService.NewVolumeProfileIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeProfileIdentityByName successfully`, func() {
				name := "general-purpose"
				model, err := testService.NewVolumeProfileIdentityByName(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumePrototypeVolumeByCapacity successfully`, func() {
				var profile vpcclassicv1.VolumeProfileIdentityIntf = nil
				var zone vpcclassicv1.ZoneIdentityIntf = nil
				capacity := int64(100)
				_, err := testService.NewVolumePrototypeVolumeByCapacity(profile, zone, capacity)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewZoneIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
				model, err := testService.NewZoneIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewZoneIdentityByName successfully`, func() {
				name := "us-south-1"
				model, err := testService.NewZoneIdentityByName(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004"
				model, err := testService.NewLoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID successfully`, func() {
				id := "70294e14-4e61-11e8-bcf4-0242ac110004"
				model, err := testService.NewLoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/load_balancers/dd754295-e9e0-4c9d-bf6c-58fbc59e5727/pools/70294e14-4e61-11e8-bcf4-0242ac110004"
				model, err := testService.NewLoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID successfully`, func() {
				id := "70294e14-4e61-11e8-bcf4-0242ac110004"
				model, err := testService.NewLoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByAddress successfully`, func() {
				address := "203.0.113.1"
				model, err := testService.NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByAddress(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south-1:a/123456::floating-ip:39300233-9995-4806-89a5-3c1b6eb88689"
				model, err := testService.NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/39300233-9995-4806-89a5-3c1b6eb88689"
				model, err := testService.NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByID successfully`, func() {
				id := "39300233-9995-4806-89a5-3c1b6eb88689"
				model, err := testService.NewPublicGatewayPrototypeFloatingIpFloatingIPIdentityFloatingIPIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePatchSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolAllRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolICMPRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN successfully`, func() {
				crn := "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByCRN(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref successfully`, func() {
				href := "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByHref(href)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByID successfully`, func() {
				id := "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				model, err := testService.NewSecurityGroupRulePrototypeSecurityGroupRuleProtocolTCPUDPRemoteSecurityGroupIdentitySecurityGroupIdentityByID(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVolumeAttachmentPrototypeInstanceContextVolumeVolumePrototypeInstanceContextVolumePrototypeInstanceContextVolumeByCapacity successfully`, func() {
				var profile vpcclassicv1.VolumeProfileIdentityIntf = nil
				capacity := int64(100)
				_, err := testService.NewVolumeAttachmentPrototypeInstanceContextVolumeVolumePrototypeInstanceContextVolumePrototypeInstanceContextVolumeByCapacity(profile, capacity)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
