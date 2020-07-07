package dnsrecordsv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/dnsrecordsv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`DNSRecordsV1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
		fmt.Println("config is not loaded : ", err)
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zone_id := os.Getenv("ZONE_ID")
	globalOptions := &DnsRecordsV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}
	testService, testServiceErr := NewDnsRecordsV1(globalOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}
	Describe(`CIS_Frontend_API_Spec-DNSRecords.yaml`, func() {
		Context("DnsRecordsV1Options", func() {
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllDnsRecords(testService.NewListAllDnsRecordsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, rule := range result.Result {
					option := testService.NewDeleteDnsRecordOptions(*rule.ID)
					delResult, response, err := testService.DeleteDnsRecord(option)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())

				}
			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllDnsRecords(testService.NewListAllDnsRecordsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, rule := range result.Result {
					option := testService.NewDeleteDnsRecordOptions(*rule.ID)
					delResult, response, err := testService.DeleteDnsRecord(option)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
			It(`create/delete/get dns A type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_A
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.test-example.com")
				options.SetType(mode)
				options.SetContent("1.2.3.4")
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Record Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// update DNS Record Options
				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.testexample.com")
				updateOpt.SetContent("Test Text")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
		})
	})
	It(`create/delete/get dns Caa type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Caa
		options := testService.NewCreateDnsRecordOptions()
		options.SetName("host-1.test-example.com")
		options.SetType(mode)
		Data := map[string]interface{}{"tag": "http",
			"value": "domain.com"}
		options.SetData(Data)
		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Reord Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		//Update DNS Record Option
		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_Txt
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("Test Text")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
	It(`create/delete/get dns Cname type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Cname
		options := testService.NewCreateDnsRecordOptions()
		options.SetName("host-1.test-example.com")
		options.SetType(mode)
		options.SetContent("domain.com")
		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Reord Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_Txt
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("Test Text")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
	It(`create/delete/get dns Aaaa type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Aaaa
		options := testService.NewCreateDnsRecordOptions()
		options.SetName("host-1.test-example.com")
		options.SetType(mode)
		options.SetContent("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Record Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_Txt
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("Test Text")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
	It(`create/delete/get dns Mx type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Mx
		options := testService.NewCreateDnsRecordOptions()
		options.SetName("host-1.test-example.com")
		options.SetType(mode)
		//options.SetContent("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
		options.SetContent("domain.com")
		options.SetPriority(int64(1))
		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Reord Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_Txt
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("Test Text")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
	It(`create/delete/get dns Ns type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Ns
		options := testService.NewCreateDnsRecordOptions()
		options.SetName("host-1.test-example.com")
		options.SetType(mode)
		options.SetContent("domain.com")
		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Reord Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_Txt
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("Test Text")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
	It(`create/delete/get dns Spf type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Spf
		options := testService.NewCreateDnsRecordOptions()
		options.SetName("host-1.test-example.com")
		options.SetType(mode)
		options.SetContent("domain.com")
		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Reord Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_Txt
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("Test Text")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
	It(`create/delete/get dns Srv type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Srv
		options := testService.NewCreateDnsRecordOptions()
		options.SetType(mode)
		Data := map[string]interface{}{"name": "host-1.test-example.com",
			"priority": int64(1),
			"service":  "_sip.example.com",
			"proto":    "UDP",
			"weight":   10,
			"port":     1024,
			"target":   "domain.com"}
		options.SetData(Data)
		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Reord Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_Txt
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("Test Text")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
	It(`create/delete/get dns Txt type records`, func() {
		shouldSkipTest()
		mode := CreateDnsRecordOptions_Type_Txt
		options := testService.NewCreateDnsRecordOptions()
		options.SetName("host-1.test-example.com")
		options.SetType(mode)
		options.SetContent("Test Text")

		result, response, err := testService.CreateDnsRecord(options)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())

		// get DNS Reord Options
		getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
		result, response, err = testService.GetDnsRecord(getOption)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
		newModes := CreateDnsRecordOptions_Type_A
		updateOpt.SetType(newModes)
		updateOpt.SetName("host-1.testexample.com")
		updateOpt.SetContent("1.2.3.4")
		result, response, err = testService.UpdateDnsRecord(updateOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())

		// delete Dns Record Options
		deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
		delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(delResult).ToNot(BeNil())
		Expect(*result.Success).Should(BeTrue())
	})
})
