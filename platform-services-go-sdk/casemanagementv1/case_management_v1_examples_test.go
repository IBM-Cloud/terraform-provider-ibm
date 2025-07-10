//go:build examples

/**
 * (C) Copyright IBM Corp. 2021, 2022.
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

package casemanagementv1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/casemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Case Management service.
//
// The following configuration properties are assumed to be defined:
//
// CASE_MANAGEMENT_URL=<service url>
// CASE_MANAGEMENT_AUTH_TYPE=iam
// CASE_MANAGEMENT_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// CASE_MANAGEMENT_APIKEY=<IAM apikey>
// CASE_MANAGEMENT_RESOURCE_CRN=<CRN of resource to use in examples>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`CaseManagementV1 Examples Tests`, func() {
	const externalConfigFile = "../case_management.env"

	var (
		caseManagementService *casemanagementv1.CaseManagementV1
		config                map[string]string
		configLoaded          bool = false
		caseNumber            string
		attachmentID          string
		resourceCRN           string
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping tests...")
		}
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(casemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			resourceCRN = config["RESOURCE_CRN"]
			if resourceCRN == "" {
				Skip("RESOURCE_CRN configuration property not found, skipping...")
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			caseManagementServiceOptions := &casemanagementv1.CaseManagementV1Options{}

			caseManagementService, err = casemanagementv1.NewCaseManagementV1UsingExternalConfig(caseManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(caseManagementService).ToNot(BeNil())
		})
	})

	Describe(`CaseManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCase request example`, func() {
			fmt.Println("\nCreateCase() result:")
			// begin-createCase

			offeringType, _ := caseManagementService.NewOfferingType(
				casemanagementv1.OfferingTypeGroupCRNServiceNameConst,
				"cloud-object-storage",
			)
			offeringPayload, _ := caseManagementService.NewOffering(
				"Cloud Object Storage",
				offeringType,
			)

			createCaseOptions := caseManagementService.NewCreateCaseOptions(
				"technical",
				"Example technical case",
				"This is an example case description. This is where the problem would be described.",
			)
			createCaseOptions.SetSeverity(4)
			createCaseOptions.SetOffering(offeringPayload)

			caseVar, response, err := caseManagementService.CreateCase(createCaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(caseVar, "", "  ")
			fmt.Println(string(b))

			// end-createCase

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(caseVar).ToNot(BeNil())
			Expect(caseVar.Number).ToNot(BeNil())

			caseNumber = *caseVar.Number

		})
		It(`GetCase request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())

			fmt.Println("\nGetCase() result:")
			// begin-getCase

			getCaseOptions := caseManagementService.NewGetCaseOptions(
				caseNumber,
			)
			getCaseOptions.SetFields([]string{
				casemanagementv1.GetCaseOptionsFieldsDescriptionConst,
				casemanagementv1.GetCaseOptionsFieldsStatusConst,
				casemanagementv1.GetCaseOptionsFieldsSeverityConst,
				casemanagementv1.GetCaseOptionsFieldsCreatedByConst,
			})

			caseVar, response, err := caseManagementService.GetCase(getCaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(caseVar, "", "  ")
			fmt.Println(string(b))

			// end-getCase

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(caseVar).ToNot(BeNil())

		})
		It(`GetCases request example`, func() {
			fmt.Println("\nGetCases() result:")
			// begin-getCases
			getCasesOptions := &casemanagementv1.GetCasesOptions{
				Search: core.StringPtr("Example"),
			}

			pager, err := caseManagementService.NewGetCasesPager(getCasesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []casemanagementv1.Case
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-getCases
		})
		It(`AddComment request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())

			fmt.Println("\nAddComment() result:")
			// begin-addComment

			addCommentOptions := caseManagementService.NewAddCommentOptions(
				caseNumber,
				"This is an example comment.",
			)

			comment, response, err := caseManagementService.AddComment(addCommentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(comment, "", "  ")
			fmt.Println(string(b))

			// end-addComment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(comment).ToNot(BeNil())

		})
		It(`AddWatchlist request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())

			fmt.Println("\nAddWatchlist() result:")
			// begin-addWatchlist

			watchListUser, _ := caseManagementService.NewUser("IBMid", "abc@ibm.com")

			addWatchlistOptions := caseManagementService.NewAddWatchlistOptions(
				caseNumber,
			)
			addWatchlistOptions.SetWatchlist([]casemanagementv1.User{*watchListUser})

			watchlistAddResponse, response, err := caseManagementService.AddWatchlist(addWatchlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(watchlistAddResponse, "", "  ")
			fmt.Println(string(b))

			// end-addWatchlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(watchlistAddResponse).ToNot(BeNil())

		})
		It(`RemoveWatchlist request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())

			fmt.Println("\nRemoveWatchlist() result:")
			// begin-removeWatchlist

			watchListUser, _ := caseManagementService.NewUser("IBMid", "abc@ibm.com")

			removeWatchlistOptions := caseManagementService.NewRemoveWatchlistOptions(
				caseNumber,
			)
			removeWatchlistOptions.SetWatchlist([]casemanagementv1.User{*watchListUser})

			watchlist, response, err := caseManagementService.RemoveWatchlist(removeWatchlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(watchlist, "", "  ")
			fmt.Println(string(b))

			// end-removeWatchlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(watchlist).ToNot(BeNil())

		})
		It(`AddResource request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())

			fmt.Println("\nAddResource() result:")
			// begin-addResource

			addResourceOptions := caseManagementService.NewAddResourceOptions(
				caseNumber,
			)
			addResourceOptions.SetCRN(resourceCRN)
			addResourceOptions.SetNote("This resource is the service that is having the problem.")

			resource, response, err := caseManagementService.AddResource(addResourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resource, "", "  ")
			fmt.Println(string(b))

			// end-addResource

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())

		})
		It(`UploadFile request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())

			fmt.Println("\nNewUploadFileOptions() result:")
			// begin-uploadFile

			exampleFileContent := "This is the content of the file to upload."

			exampleFile, _ := caseManagementService.NewFileWithMetadata(io.NopCloser(strings.NewReader(exampleFileContent)))
			exampleFile.Filename = core.StringPtr("example.log")
			exampleFile.ContentType = core.StringPtr("application/octet-stream")

			filePayload := []casemanagementv1.FileWithMetadata{*exampleFile}

			uploadFileOptions := caseManagementService.NewUploadFileOptions(
				caseNumber,
				filePayload,
			)

			attachment, response, err := caseManagementService.UploadFile(uploadFileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachment, "", "  ")
			fmt.Println(string(b))

			// end-uploadFile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
			Expect(attachment.ID).ToNot(BeNil())

			attachmentID = *attachment.ID
		})
		It(`DownloadFile request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())
			Expect(attachmentID).ToNot(BeEmpty())

			fmt.Println("\nDownloadFile() result:")
			// begin-downloadFile

			downloadFileOptions := caseManagementService.NewDownloadFileOptions(
				caseNumber,
				attachmentID,
			)

			result, response, err := caseManagementService.DownloadFile(downloadFileOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				buf := new(bytes.Buffer)
				_, _ = buf.ReadFrom(result)
				fmt.Println(buf.String())
			}

			// end-downloadFile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`DeleteFile request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())
			Expect(attachmentID).ToNot(BeEmpty())

			fmt.Println("\nDeleteFile() result:")
			// begin-deleteFile

			deleteFileOptions := caseManagementService.NewDeleteFileOptions(
				caseNumber,
				attachmentID,
			)

			attachmentList, response, err := caseManagementService.DeleteFile(deleteFileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentList, "", "  ")
			fmt.Println(string(b))

			// end-deleteFile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentList).ToNot(BeNil())

		})
		It(`UpdateCaseStatus request example`, func() {
			Expect(caseNumber).ToNot(BeEmpty())

			fmt.Println("\nUpdateCaseStatus() result:")
			// begin-updateCaseStatus

			statusPayloadModel := &casemanagementv1.ResolvePayload{
				Action:         core.StringPtr("resolve"),
				Comment:        core.StringPtr("The problem has been resolved."),
				ResolutionCode: core.Int64Ptr(int64(1)),
			}

			updateCaseStatusOptions := caseManagementService.NewUpdateCaseStatusOptions(
				caseNumber,
				statusPayloadModel,
			)

			caseVar, response, err := caseManagementService.UpdateCaseStatus(updateCaseStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(caseVar, "", "  ")
			fmt.Println(string(b))

			// end-updateCaseStatus

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(caseVar).ToNot(BeNil())

		})
	})
})
