/*
 * (C) Copyright IBM Corp. 2020.
 */

package rangeapplicationsv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/rangeapplicationsv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`rangeapplicationsv1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zone_id := os.Getenv("ZONE_ID")
	url := os.Getenv("URL")
	globalOptions := &RangeApplicationsV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewRangeApplicationsV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`create/update/get/delete origin direct range app`, func() {
		Context(`create/update/get/delete origin direct range app`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewListRangeAppsOptions()
				listResult, listResp, listErr := service.ListRangeApps(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, app := range listResult.Result {
					// delete range app
					delOpt := service.NewDeleteRangeAppOptions(*app.ID)
					delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewListRangeAppsOptions()
				listResult, listResp, listErr := service.ListRangeApps(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, app := range listResult.Result {
					// delete range app
					delOpt := service.NewDeleteRangeAppOptions(*app.ID)
					delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`create/update/get/delete origin direct range app`, func() {
				shouldSkipTest()
				protocol := "tcp/22"
				dnsName := "example." + url
				origin_direct := []string{"tcp://12.1.1.1:22"}

				dnsOpt := &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}
				edgeIPs := &RangeAppReqEdgeIps{
					Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
					Connectivity: core.StringPtr(RangeAppReqEdgeIps_Connectivity_All),
				}
				// create range app
				createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
				createOpt.SetEdgeIps(edgeIPs)
				createOpt.SetIpFirewall(true)
				createOpt.SetOriginDirect(origin_direct)
				createOpt.SetOriginPort(22)
				createOpt.SetProxyProtocol(CreateRangeAppOptions_ProxyProtocol_Off)
				createOpt.SetTls(CreateRangeAppOptions_Tls_Off)
				createOpt.SetTrafficType(CreateRangeAppOptions_TrafficType_Direct)

				createResult, createResp, createErr := service.CreateRangeApp(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				// update range app
				origin_direct = []string{"tcp://12.12.12.12:25"}
				updateOpt := service.NewUpdateRangeAppOptions(*createResult.Result.ID, protocol, dnsOpt)
				updateOpt.SetEdgeIps(edgeIPs)
				updateOpt.SetIpFirewall(true)
				updateOpt.SetOriginDirect(origin_direct)
				updateOpt.SetOriginPort(22)
				updateOpt.SetProxyProtocol(UpdateRangeAppOptions_ProxyProtocol_Off)
				updateOpt.SetTls(UpdateRangeAppOptions_Tls_Off)
				updateOpt.SetTrafficType(UpdateRangeAppOptions_TrafficType_Direct)

				updateResult, updateResp, updateErr := service.UpdateRangeApp(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// get range app
				getOpt := service.NewGetRangeAppOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetRangeApp(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete range app
				delOpt := service.NewDeleteRangeAppOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/get/delete origin dns range app`, func() {
				shouldSkipTest()
				protocol := "tcp/22"
				dnsName := "example." + url
				origin_direct := []string{"tcp://12.1.1.1:22"}

				dnsOpt := &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}
				edgeIPs := &RangeAppReqEdgeIps{
					Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
					Connectivity: core.StringPtr(RangeAppReqEdgeIps_Connectivity_All),
				}
				// create range app
				createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
				createOpt.SetEdgeIps(edgeIPs)
				createOpt.SetIpFirewall(true)
				createOpt.SetOriginDirect(origin_direct)
				createOpt.SetOriginPort(22)
				createOpt.SetProxyProtocol(CreateRangeAppOptions_ProxyProtocol_Off)
				createOpt.SetTls(CreateRangeAppOptions_Tls_Off)
				createOpt.SetTrafficType(CreateRangeAppOptions_TrafficType_Direct)

				createResult, createResp, createErr := service.CreateRangeApp(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				// update range app
				dnsName = "origin-example.net"
				dnsOpt = &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}
				updateOpt := service.NewUpdateRangeAppOptions(*createResult.Result.ID, protocol, dnsOpt)
				updateOpt.SetEdgeIps(edgeIPs)
				updateOpt.SetIpFirewall(true)
				updateOpt.SetOriginDirect(origin_direct)
				updateOpt.SetOriginPort(22)
				updateOpt.SetProxyProtocol(UpdateRangeAppOptions_ProxyProtocol_Off)
				updateOpt.SetTls(UpdateRangeAppOptions_Tls_Off)
				updateOpt.SetTrafficType(UpdateRangeAppOptions_TrafficType_Direct)

				updateResult, updateResp, updateErr := service.UpdateRangeApp(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// get range app
				getOpt := service.NewGetRangeAppOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetRangeApp(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete range app
				delOpt := service.NewDeleteRangeAppOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`list range app`, func() {
				shouldSkipTest()
				for i := 1; i < 5; i++ {
					protocol := "tcp/22"
					dnsName := fmt.Sprintf("example%d.%s", i, url)
					direct := fmt.Sprintf("tcp://12.1.1.%d:2%d", i, i)
					origin_direct := []string{direct}

					dnsOpt := &RangeAppReqDns{
						Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
						Name: core.StringPtr(dnsName),
					}
					edgeIPs := &RangeAppReqEdgeIps{
						Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
						Connectivity: core.StringPtr(RangeAppReqEdgeIps_Connectivity_All),
					}
					// create range app
					createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
					createOpt.SetEdgeIps(edgeIPs)
					createOpt.SetIpFirewall(true)
					createOpt.SetOriginDirect(origin_direct)
					createOpt.SetOriginPort(22)
					createOpt.SetProxyProtocol(CreateRangeAppOptions_ProxyProtocol_Off)
					createOpt.SetTls(CreateRangeAppOptions_Tls_Off)
					createOpt.SetTrafficType(CreateRangeAppOptions_TrafficType_Direct)

					createResult, createResp, createErr := service.CreateRangeApp(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())
				}

				listOpt := service.NewListRangeAppsOptions()
				listOpt.SetPage(2)
				listOpt.SetPerPage(2)
				listOpt.SetDirection(ListRangeAppsOptions_Direction_Asc)
				listOpt.SetOrder(ListRangeAppsOptions_Order_Dns)
				listResult, listResp, listErr := service.ListRangeApps(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

			})
			It(`create/update/get/delete ip firewall range app`, func() {
				shouldSkipTest()
				protocol := "tcp/22"
				dnsName := "example." + url
				origin_direct := []string{"tcp://12.1.1.1:22"}

				dnsOpt := &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}
				edgeIPs := &RangeAppReqEdgeIps{
					Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
					Connectivity: core.StringPtr(RangeAppReqEdgeIps_Connectivity_All),
				}
				// create range app
				createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
				createOpt.SetEdgeIps(edgeIPs)
				createOpt.SetIpFirewall(false)
				createOpt.SetOriginDirect(origin_direct)
				createOpt.SetOriginPort(22)
				createOpt.SetProxyProtocol(CreateRangeAppOptions_ProxyProtocol_Off)
				createOpt.SetTls(CreateRangeAppOptions_Tls_Off)
				createOpt.SetTrafficType(CreateRangeAppOptions_TrafficType_Direct)

				createResult, createResp, createErr := service.CreateRangeApp(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				// update range app
				updateOpt := service.NewUpdateRangeAppOptions(*createResult.Result.ID, protocol, dnsOpt)
				updateOpt.SetEdgeIps(edgeIPs)
				updateOpt.SetIpFirewall(true)
				updateOpt.SetOriginDirect(origin_direct)
				updateOpt.SetOriginPort(22)
				updateOpt.SetProxyProtocol(UpdateRangeAppOptions_ProxyProtocol_Off)
				updateOpt.SetTls(UpdateRangeAppOptions_Tls_Off)
				updateOpt.SetTrafficType(UpdateRangeAppOptions_TrafficType_Direct)

				updateResult, updateResp, updateErr := service.UpdateRangeApp(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// get range app
				getOpt := service.NewGetRangeAppOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetRangeApp(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete range app
				delOpt := service.NewDeleteRangeAppOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/get/delete proxy protocol range app`, func() {
				shouldSkipTest()
				protocol := "tcp/22"
				dnsName := "example." + url
				origin_direct := []string{"tcp://12.1.1.1:22"}

				dnsOpt := &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}
				edgeIPs := &RangeAppReqEdgeIps{
					Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
					Connectivity: core.StringPtr(RangeAppReqEdgeIps_Connectivity_All),
				}

				createProxies := []string{
					CreateRangeAppOptions_ProxyProtocol_Off,
					CreateRangeAppOptions_ProxyProtocol_V1,
					CreateRangeAppOptions_ProxyProtocol_V2,
				}
				updateProxies := []string{
					UpdateRangeAppOptions_ProxyProtocol_V1,
					UpdateRangeAppOptions_ProxyProtocol_V2,
					UpdateRangeAppOptions_ProxyProtocol_Off,
				}
				for i, proxy := range createProxies {
					// create range app
					createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
					createOpt.SetEdgeIps(edgeIPs)
					createOpt.SetIpFirewall(false)
					createOpt.SetOriginDirect(origin_direct)
					createOpt.SetOriginPort(22)
					createOpt.SetProxyProtocol(proxy)
					createOpt.SetTls(CreateRangeAppOptions_Tls_Off)
					createOpt.SetTrafficType(CreateRangeAppOptions_TrafficType_Direct)

					createResult, createResp, createErr := service.CreateRangeApp(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())

					// update range app
					updateOpt := service.NewUpdateRangeAppOptions(*createResult.Result.ID, protocol, dnsOpt)
					updateOpt.SetEdgeIps(edgeIPs)
					updateOpt.SetIpFirewall(true)
					updateOpt.SetOriginDirect(origin_direct)
					updateOpt.SetOriginPort(22)
					updateOpt.SetProxyProtocol(updateProxies[i])
					updateOpt.SetTls(UpdateRangeAppOptions_Tls_Off)
					updateOpt.SetTrafficType(UpdateRangeAppOptions_TrafficType_Direct)

					updateResult, updateResp, updateErr := service.UpdateRangeApp(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					// get range app
					getOpt := service.NewGetRangeAppOptions(*createResult.Result.ID)
					getResult, getResp, getErr := service.GetRangeApp(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					// delete range app
					delOpt := service.NewDeleteRangeAppOptions(*createResult.Result.ID)
					delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`create/update/get/delete traffic type range app`, func() {
				shouldSkipTest()
				protocol := "tcp/22"
				dnsName := "example." + url
				origin_direct := []string{"tcp://12.1.1.1:22"}

				dnsOpt := &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}
				edgeIPs := &RangeAppReqEdgeIps{
					Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
					Connectivity: core.StringPtr(RangeAppReqEdgeIps_Connectivity_All),
				}

				createTraffics := []string{
					CreateRangeAppOptions_TrafficType_Direct,
					CreateRangeAppOptions_TrafficType_Http,
					CreateRangeAppOptions_TrafficType_Https,
				}
				updateTraffics := []string{
					CreateRangeAppOptions_TrafficType_Http,
					CreateRangeAppOptions_TrafficType_Https,
					CreateRangeAppOptions_TrafficType_Direct,
				}
				for i, traffic := range createTraffics {
					// create range app
					createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
					createOpt.SetEdgeIps(edgeIPs)
					createOpt.SetIpFirewall(false)
					createOpt.SetOriginDirect(origin_direct)
					createOpt.SetOriginPort(22)
					createOpt.SetProxyProtocol(CreateRangeAppOptions_ProxyProtocol_Off)
					createOpt.SetTls(CreateRangeAppOptions_Tls_Off)
					createOpt.SetTrafficType(traffic)

					createResult, createResp, createErr := service.CreateRangeApp(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())

					// update range app
					updateOpt := service.NewUpdateRangeAppOptions(*createResult.Result.ID, protocol, dnsOpt)
					updateOpt.SetEdgeIps(edgeIPs)
					updateOpt.SetIpFirewall(false)
					updateOpt.SetOriginDirect(origin_direct)
					updateOpt.SetOriginPort(22)
					updateOpt.SetProxyProtocol(UpdateRangeAppOptions_ProxyProtocol_Off)
					updateOpt.SetTls(UpdateRangeAppOptions_Tls_Off)
					updateOpt.SetTrafficType(updateTraffics[i])

					updateResult, updateResp, updateErr := service.UpdateRangeApp(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					// get range app
					getOpt := service.NewGetRangeAppOptions(*createResult.Result.ID)
					getResult, getResp, getErr := service.GetRangeApp(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					// delete range app
					delOpt := service.NewDeleteRangeAppOptions(*createResult.Result.ID)
					delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`create/update/get/delete tls range app`, func() {
				shouldSkipTest()
				protocol := "tcp/5678"
				dnsName := "example." + url
				origin_direct := []string{"tcp://12.1.1.1:5678"}

				dnsOpt := &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}
				edgeIPs := &RangeAppReqEdgeIps{
					Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
					Connectivity: core.StringPtr(RangeAppReqEdgeIps_Connectivity_All),
				}

				createTLS := []string{
					CreateRangeAppOptions_Tls_Off,
					CreateRangeAppOptions_Tls_Flexible,
					CreateRangeAppOptions_Tls_Strict,
					CreateRangeAppOptions_Tls_Full,
				}
				updateTLS := []string{
					UpdateRangeAppOptions_Tls_Flexible,
					UpdateRangeAppOptions_Tls_Strict,
					UpdateRangeAppOptions_Tls_Full,
					UpdateRangeAppOptions_Tls_Off,
				}
				for i, tls := range createTLS {
					// create range app
					createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
					createOpt.SetEdgeIps(edgeIPs)
					createOpt.SetIpFirewall(false)
					createOpt.SetOriginDirect(origin_direct)
					createOpt.SetOriginPort(22)
					createOpt.SetProxyProtocol(CreateRangeAppOptions_ProxyProtocol_V1)
					createOpt.SetTls(tls)
					createOpt.SetTrafficType(CreateRangeAppOptions_TrafficType_Direct)

					createResult, createResp, createErr := service.CreateRangeApp(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())

					// update range app
					updateOpt := service.NewUpdateRangeAppOptions(*createResult.Result.ID, protocol, dnsOpt)
					updateOpt.SetEdgeIps(edgeIPs)
					updateOpt.SetIpFirewall(true)
					updateOpt.SetOriginDirect(origin_direct)
					updateOpt.SetOriginPort(22)
					updateOpt.SetProxyProtocol(UpdateRangeAppOptions_ProxyProtocol_V1)
					updateOpt.SetTls(updateTLS[i])
					updateOpt.SetTrafficType(UpdateRangeAppOptions_TrafficType_Direct)

					updateResult, updateResp, updateErr := service.UpdateRangeApp(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					// get range app
					getOpt := service.NewGetRangeAppOptions(*createResult.Result.ID)
					getResult, getResp, getErr := service.GetRangeApp(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					// delete range app
					delOpt := service.NewDeleteRangeAppOptions(*createResult.Result.ID)
					delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`create/update/get/delete connectivity range app`, func() {
				shouldSkipTest()
				protocol := "tcp/22"
				dnsName := "example." + url
				origin_direct := []string{"tcp://12.1.1.1:22"}

				dnsOpt := &RangeAppReqDns{
					Type: core.StringPtr(RangeApplicationObjectDns_Type_Cname),
					Name: core.StringPtr(dnsName),
				}

				createConnectivity := []string{
					RangeApplicationObjectEdgeIps_Connectivity_All,
					RangeApplicationObjectEdgeIps_Connectivity_Ipv4,
					RangeApplicationObjectEdgeIps_Connectivity_Ipv6,
				}
				updateConnectivity := []string{
					RangeApplicationObjectEdgeIps_Connectivity_Ipv4,
					RangeApplicationObjectEdgeIps_Connectivity_Ipv6,
					RangeApplicationObjectEdgeIps_Connectivity_All,
				}
				for i, conn := range createConnectivity {
					// create range app
					edgeIPs := &RangeAppReqEdgeIps{
						Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
						Connectivity: core.StringPtr(conn),
					}
					createOpt := service.NewCreateRangeAppOptions(protocol, dnsOpt)
					createOpt.SetEdgeIps(edgeIPs)
					createOpt.SetIpFirewall(false)
					createOpt.SetOriginDirect(origin_direct)
					createOpt.SetOriginPort(22)
					createOpt.SetProxyProtocol(CreateRangeAppOptions_ProxyProtocol_V2)
					createOpt.SetTls(CreateRangeAppOptions_Tls_Flexible)
					createOpt.SetTrafficType(CreateRangeAppOptions_TrafficType_Direct)

					createResult, createResp, createErr := service.CreateRangeApp(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())

					// update range app
					edgeIPs = &RangeAppReqEdgeIps{
						Type:         core.StringPtr(RangeAppReqEdgeIps_Type_Dynamic),
						Connectivity: core.StringPtr(updateConnectivity[i]),
					}
					updateOpt := service.NewUpdateRangeAppOptions(*createResult.Result.ID, protocol, dnsOpt)
					updateOpt.SetEdgeIps(edgeIPs)
					updateOpt.SetIpFirewall(true)
					updateOpt.SetOriginDirect(origin_direct)
					updateOpt.SetOriginPort(22)
					updateOpt.SetProxyProtocol(UpdateRangeAppOptions_ProxyProtocol_V2)
					updateOpt.SetTls(CreateRangeAppOptions_Tls_Flexible)
					updateOpt.SetTrafficType(UpdateRangeAppOptions_TrafficType_Direct)

					updateResult, updateResp, updateErr := service.UpdateRangeApp(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					// get range app
					getOpt := service.NewGetRangeAppOptions(*createResult.Result.ID)
					getResult, getResp, getErr := service.GetRangeApp(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					// delete range app
					delOpt := service.NewDeleteRangeAppOptions(*createResult.Result.ID)
					delResult, delResp, delErr := service.DeleteRangeApp(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
		})
	})
})
