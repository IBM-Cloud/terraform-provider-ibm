//go:build examples

/**
 * (C) Copyright IBM Corp. 2021.
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

package globalcatalogv1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Global Catalog service.
//
// The following configuration properties are assumed to be defined:
//
// GLOBAL_CATALOG_URL=<service url>
// GLOBAL_CATALOG_AUTH_TYPE=iam
// GLOBAL_CATALOG_APIKEY=<IAM apikey>
// GLOBAL_CATALOG_AUTH_URL=<IAM token service URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`GlobalCatalogV1 Examples Tests`, func() {
	const externalConfigFile = "../global_catalog.env"

	var (
		globalCatalogService *globalcatalogv1.GlobalCatalogV1
		config               map[string]string
		configLoaded         bool = false
		catalogEntryID       string
		fetchedEntry         *globalcatalogv1.CatalogEntry
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
			config, err = core.GetServiceProperties(globalcatalogv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
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

			globalCatalogServiceOptions := &globalcatalogv1.GlobalCatalogV1Options{}

			globalCatalogService, err = globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(globalCatalogServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(globalCatalogService).ToNot(BeNil())

			core.SetLogger(core.NewLogger(core.LevelInfo, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			globalCatalogService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GlobalCatalogV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalogEntry request example`, func() {
			fmt.Println("\nCreateCatalogEntry() result:")
			// begin-create_catalog_entry

			displayName := "Example Web Starter"
			description := "Use the Example service in your applications"
			longDescription := "This is a starter that helps you use the Example service within your applications."
			overviewModelEN := &globalcatalogv1.Overview{
				DisplayName:     &displayName,
				Description:     &description,
				LongDescription: &longDescription,
			}
			overviewUIModel := make(map[string]globalcatalogv1.Overview)
			overviewUIModel["en"] = *overviewModelEN

			smallImageURL := "https://somehost.com/examplewebstarter/cachedIcon/small/0"
			mediumImageURL := "https://somehost.com/examplewebstarter/cachedIcon/medium/0"
			largeImageURL := "https://somehost.com/examplewebstarter/cachedIcon/large/0"
			imageModel := &globalcatalogv1.Image{
				Image:        &largeImageURL,
				SmallImage:   &smallImageURL,
				MediumImage:  &mediumImageURL,
				FeatureImage: &largeImageURL,
			}

			providerModel := &globalcatalogv1.Provider{
				Email:        core.StringPtr("info@examplestarter.com"),
				Name:         core.StringPtr("Example Starter Co., Inc."),
				Contact:      core.StringPtr("Example Starter Developer Relations"),
				SupportEmail: core.StringPtr("support@examplestarter.com"),
				Phone:        core.StringPtr("800-555-1234"),
			}

			metadataModel := &globalcatalogv1.ObjectMetadataSet{
				Version: core.StringPtr("1.0.0"),
			}

			catalogEntryID = uuid.New().String()

			createCatalogEntryOptions := globalCatalogService.NewCreateCatalogEntryOptions(
				"exampleWebStarter123",
				globalcatalogv1.CreateCatalogEntryOptionsKindTemplateConst,
				overviewUIModel,
				imageModel,
				false,
				[]string{"example-tag-1", "example-tag-2"},
				providerModel,
				catalogEntryID,
			)
			createCatalogEntryOptions.SetActive(true)
			createCatalogEntryOptions.SetMetadata(metadataModel)

			catalogEntry, response, err := globalCatalogService.CreateCatalogEntry(createCatalogEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogEntry, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogEntry).ToNot(BeNil())

		})
		It(`GetCatalogEntry request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nGetCatalogEntry() result:")
			// begin-get_catalog_entry

			getCatalogEntryOptions := globalCatalogService.NewGetCatalogEntryOptions(
				catalogEntryID,
			)
			getCatalogEntryOptions.SetComplete(true)

			catalogEntry, response, err := globalCatalogService.GetCatalogEntry(getCatalogEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogEntry, "", "  ")
			fmt.Println(string(b))

			fetchedEntry = catalogEntry

			// end-get_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogEntry).ToNot(BeNil())

		})
		It(`UpdateCatalogEntry request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nUpdateCatalogEntry() result:")
			// begin-update_catalog_entry

			displayName := "Example Web Starter V2"
			description := "Use the Example V2 service in your applications"
			longDescription := "This is a starter that helps you use the Example V2 service within your applications."
			overviewModelEN := &globalcatalogv1.Overview{
				DisplayName:     &displayName,
				Description:     &description,
				LongDescription: &longDescription,
			}
			overviewUI := make(map[string]globalcatalogv1.Overview)
			overviewUI["en"] = *overviewModelEN

			smallImageURL := "https://somehost.com/examplewebstarter/cachedIcon/small/0"
			mediumImageURL := "https://somehost.com/examplewebstarter/cachedIcon/medium/0"
			largeImageURL := "https://somehost.com/examplewebstarter/cachedIcon/large/0"
			imageModel := &globalcatalogv1.Image{
				Image:        &largeImageURL,
				SmallImage:   &smallImageURL,
				MediumImage:  &mediumImageURL,
				FeatureImage: &largeImageURL,
			}

			providerModel := &globalcatalogv1.Provider{
				Email:        core.StringPtr("info@examplestarter.com"),
				Name:         core.StringPtr("Example Starter Co., Inc."),
				Contact:      core.StringPtr("Example Starter Developer Relations"),
				SupportEmail: core.StringPtr("support@examplestarter.com"),
				Phone:        core.StringPtr("800-555-1234"),
			}

			metadataModel := &globalcatalogv1.ObjectMetadataSet{
				Version: core.StringPtr("2.0.0"),
			}

			updateCatalogEntryOptions := globalCatalogService.NewUpdateCatalogEntryOptions(
				catalogEntryID,
				"exampleWebStarter123",
				globalcatalogv1.UpdateCatalogEntryOptionsKindTemplateConst,
				overviewUI,
				imageModel,
				false,
				[]string{"example-tag-1", "example-tag-2", "new-example-tag-3"},
				providerModel,
			)
			updateCatalogEntryOptions.SetActive(true)
			updateCatalogEntryOptions.SetMetadata(metadataModel)
			updateCatalogEntryOptions.SetURL(*fetchedEntry.URL)

			catalogEntry, response, err := globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogEntry, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogEntry).ToNot(BeNil())

		})
		It(`ListCatalogEntries request example`, func() {
			fmt.Println("\nListCatalogEntries() result:")
			// begin-list_catalog_entries

			listCatalogEntriesOptions := globalCatalogService.NewListCatalogEntriesOptions()
			listCatalogEntriesOptions.SetOffset(0)
			listCatalogEntriesOptions.SetLimit(10)
			listCatalogEntriesOptions.SetQ("kind:template tag:example-tag-1")
			listCatalogEntriesOptions.SetComplete(true)

			entrySearchResult, response, err := globalCatalogService.ListCatalogEntries(listCatalogEntriesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entrySearchResult, "", "  ")
			fmt.Println(string(b))

			// end-list_catalog_entries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entrySearchResult).ToNot(BeNil())

		})
		It(`GetChildObjects request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nGetChildObjects() result:")
			// begin-get_child_objects

			getChildObjectsOptions := globalCatalogService.NewGetChildObjectsOptions(
				catalogEntryID,
				"*",
			)
			getChildObjectsOptions.SetOffset(0)
			getChildObjectsOptions.SetLimit(10)
			getChildObjectsOptions.SetComplete(true)

			entrySearchResult, response, err := globalCatalogService.GetChildObjects(getChildObjectsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entrySearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_child_objects

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entrySearchResult).ToNot(BeNil())
		})
		It(`RestoreCatalogEntry request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			// begin-restore_catalog_entry

			restoreCatalogEntryOptions := globalCatalogService.NewRestoreCatalogEntryOptions(
				catalogEntryID,
			)

			response, err := globalCatalogService.RestoreCatalogEntry(restoreCatalogEntryOptions)
			if err != nil {
				panic(err)
			}

			// end-restore_catalog_entry
			fmt.Printf("\nRestoreCatalogEntry() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetVisibility request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nGetVisibility() result:")
			// begin-get_visibility

			getVisibilityOptions := globalCatalogService.NewGetVisibilityOptions(
				catalogEntryID,
			)

			visibility, response, err := globalCatalogService.GetVisibility(getVisibilityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(visibility, "", "  ")
			fmt.Println(string(b))

			// end-get_visibility

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(visibility).ToNot(BeNil())

		})
		It(`UpdateVisibility request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			// begin-update_visibility

			updateVisibilityOptions := globalCatalogService.NewUpdateVisibilityOptions(
				catalogEntryID,
			)
			updateVisibilityOptions.SetRestrictions("private")

			response, err := globalCatalogService.UpdateVisibility(updateVisibilityOptions)
			if err != nil {
				fmt.Println("UpdateVisibility() returned the following error: ", err.Error())
			}

			// end-update_visibility
			fmt.Printf("\nUpdateVisibility() response status code: %d\n: ", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(response).ToNot(BeNil())
		})
		It(`GetPricing request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nGetPricing() result:")
			// begin-get_pricing

			getPricingOptions := globalCatalogService.NewGetPricingOptions(
				catalogEntryID,
			)

			pricingGet, response, err := globalCatalogService.GetPricing(getPricingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pricingGet, "", "  ")
			fmt.Println(string(b))

			// end-get_pricing

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pricingGet).ToNot(BeNil())

		})
		It(`GetAuditLogs request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nGetAuditLogs() result:")
			// begin-get_audit_logs

			getAuditLogsOptions := globalCatalogService.NewGetAuditLogsOptions(
				catalogEntryID,
			)
			getAuditLogsOptions.SetOffset(0)
			getAuditLogsOptions.SetLimit(10)

			auditSearchResult, response, err := globalCatalogService.GetAuditLogs(getAuditLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(auditSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_audit_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditSearchResult).ToNot(BeNil())
		})
		It(`UploadArtifact request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			// begin-upload_artifact

			artifactContents := "This is an example artifact associated with a catalog entry."

			uploadArtifactOptions := globalCatalogService.NewUploadArtifactOptions(
				catalogEntryID,
				"artifact.txt",
			)
			uploadArtifactOptions.SetArtifact(io.NopCloser(strings.NewReader(artifactContents)))
			uploadArtifactOptions.SetContentType("text/plain")

			response, err := globalCatalogService.UploadArtifact(uploadArtifactOptions)
			if err != nil {
				panic(err)
			}

			// end-upload_artifact
			fmt.Printf("\nUploadArtifact() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
		It(`GetArtifact request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nGetArtifact() result:")
			// begin-get_artifact

			getArtifactOptions := globalCatalogService.NewGetArtifactOptions(
				catalogEntryID,
				"artifact.txt",
			)

			result, response, err := globalCatalogService.GetArtifact(getArtifactOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				buf := new(bytes.Buffer)
				_, _ = buf.ReadFrom(result)
				fmt.Println(buf.String())
			}

			// end-get_artifact

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`ListArtifacts request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			fmt.Println("\nListArtifacts() result:")
			// begin-list_artifacts

			listArtifactsOptions := globalCatalogService.NewListArtifactsOptions(
				catalogEntryID,
			)

			artifacts, response, err := globalCatalogService.ListArtifacts(listArtifactsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(artifacts, "", "  ")
			fmt.Println(string(b))

			// end-list_artifacts

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(artifacts).ToNot(BeNil())

		})
		It(`DeleteArtifact request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			// begin-delete_artifact

			deleteArtifactOptions := globalCatalogService.NewDeleteArtifactOptions(
				catalogEntryID,
				"artifact.txt",
			)

			response, err := globalCatalogService.DeleteArtifact(deleteArtifactOptions)
			if err != nil {
				panic(err)
			}
			// end-delete_artifact
			fmt.Printf("\nDeleteArtifact() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteCatalogEntry request example`, func() {
			Expect(catalogEntryID).ToNot(BeEmpty())

			// begin-delete_catalog_entry

			deleteCatalogEntryOptions := globalCatalogService.NewDeleteCatalogEntryOptions(
				catalogEntryID,
			)

			response, err := globalCatalogService.DeleteCatalogEntry(deleteCatalogEntryOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_catalog_entry
			fmt.Printf("\nDeleteCatalogEntry() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})
