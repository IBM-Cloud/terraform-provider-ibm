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

package catalogmanagementv1_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Catalog Management service.
//
// The following configuration properties are assumed to be defined:
// CATALOG_MANAGEMENT_URL=<service base url>
// CATALOG_MANAGEMENT_AUTH_TYPE=iam
// CATALOG_MANAGEMENT_APIKEY=<IAM apikey>
// CATALOG_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// CATALOG_MANAGEMENT_CLUSTER_ID=<ID of the cluster>
// CATALOG_MANAGEMENT_ACCOUNT_ID=<ID of the Account>
// CATALOG_MANAGEMENT_GIT_TOKEN=<Token used in communication with Git repository>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`CatalogManagementV1 Examples Tests`, func() {
	const externalConfigFile = "../catalog_mgmt.env"

	var (
		err                           error
		catalogManagementService      *catalogmanagementv1.CatalogManagementV1
		catalogManagementAdminService *catalogmanagementv1.CatalogManagementV1
		serviceURL                    string
		config                        map[string]string
		accountID                     string
		bearerToken                   string
		gitAuthTokenForPublicRepo     string
		catalogID                     string
		objectCatalogID               string
		offeringID                    string
		kindID                        string
		clusterID                     string
		objectID                      string
		offeringInstanceID            string
		versionLocatorID              string
		planID                        string
		approverToken                 string
		offeringVersion               *catalogmanagementv1.Offering
		catalogAccount                *catalogmanagementv1.Account
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(catalogmanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).NotTo(BeNil())

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			catalogManagementServiceOptions := &catalogmanagementv1.CatalogManagementV1Options{
				ServiceName: catalogmanagementv1.DefaultServiceName,
			}

			catalogManagementService, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(catalogManagementServiceOptions)
			Expect(err).To(BeNil())
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(catalogManagementService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			catalogManagementService.EnableRetries(4, 30*time.Second)
		})
		It("Successfully construct the service client admin instance", func() {
			catalogManagementServiceOptions := &catalogmanagementv1.CatalogManagementV1Options{
				ServiceName: "CATALOG_MANAGEMENT_APPROVER",
			}
			catalogManagementAdminService, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(catalogManagementServiceOptions)
			Expect(err).To(BeNil())
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(catalogManagementService.Service.Options.URL).To(Equal(serviceURL))

			approverRequestToken, err := catalogManagementAdminService.Service.Options.Authenticator.(*core.IamAuthenticator).RequestToken()
			Expect(err).To(BeNil())
			approverToken = approverRequestToken.AccessToken
			Expect(approverToken).ToNot(BeNil())

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			catalogManagementService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CatalogManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`CreateCatalog for offerings request example`, func() {
			fmt.Println("\nCreateCatalog() result:")
			// begin-create_catalog

			createCatalogOptions := catalogManagementService.NewCreateCatalogOptions()
			createCatalogOptions.Label = core.StringPtr("Catalog Management Service")
			createCatalogOptions.Tags = []string{"go", "sdk"}
			createCatalogOptions.Kind = core.StringPtr("offering")
			createCatalogOptions.OwningAccount = &accountID

			catalog, response, err := catalogManagementService.CreateCatalog(createCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalog, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalog).ToNot(BeNil())

			catalogID = *catalog.ID
		})

		It(`CreateCatalog for objects request example`, func() {
			fmt.Println("\nCreateCatalog() result:")
			// begin-create_catalog

			createCatalogOptions := catalogManagementService.NewCreateCatalogOptions()
			createCatalogOptions.Label = core.StringPtr("Catalog Management Service")
			createCatalogOptions.Tags = []string{"go", "sdk"}
			createCatalogOptions.Kind = core.StringPtr("vpe")
			createCatalogOptions.OwningAccount = &accountID

			catalog, response, err := catalogManagementService.CreateCatalog(createCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalog, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalog).ToNot(BeNil())

			objectCatalogID = *catalog.ID
		})

		It(`GetCatalog request example`, func() {
			fmt.Println("\nGetCatalog() result:")
			// begin-get_catalog

			getCatalogOptions := catalogManagementService.NewGetCatalogOptions(
				catalogID,
			)

			catalog, response, err := catalogManagementService.GetCatalog(getCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalog, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})

		It(`ReplaceCatalog request example`, func() {
			fmt.Println("\nReplaceCatalog() result:")
			// begin-replace_catalog

			replaceCatalogOptions := catalogManagementService.NewReplaceCatalogOptions(
				catalogID,
			)
			replaceCatalogOptions.ID = &catalogID
			replaceCatalogOptions.Tags = []string{"python", "sdk", "updated"}
			replaceCatalogOptions.OwningAccount = &accountID
			replaceCatalogOptions.Kind = core.StringPtr("vpe")

			catalog, response, err := catalogManagementService.ReplaceCatalog(replaceCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalog, "", "  ")
			fmt.Println(string(b))

			// end-replace_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})

		It(`ListCatalogs request example`, func() {
			fmt.Println("\nListCatalogs() result:")
			// begin-list_catalogs

			listCatalogsOptions := catalogManagementService.NewListCatalogsOptions()

			catalogSearchResult, response, err := catalogManagementService.ListCatalogs(listCatalogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-list_catalogs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogSearchResult).ToNot(BeNil())
		})

		It(`CreateOffering request example`, func() {
			fmt.Println("\nCreateOffering() result:")
			// begin-create_offering

			createOfferingOptions := catalogManagementService.NewCreateOfferingOptions(
				catalogID,
			)
			createOfferingOptions.Name = core.StringPtr("offering-name")

			offering, response, err := catalogManagementService.CreateOffering(createOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-create_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

			offeringID = *offering.ID
		})

		It(`GetOffering request example`, func() {
			fmt.Println("\nGetOffering() result:")
			// begin-get_offering

			getOfferingOptions := catalogManagementService.NewGetOfferingOptions(
				catalogID,
				offeringID,
			)

			offering, response, err := catalogManagementService.GetOffering(getOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-get_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())
		})

		It(`GetOfferingStats request example`, func() {
			fmt.Println("\nGetOfferingStats() result:")
			// begin-get_offering_stats

			getOfferingStatsOptions := catalogManagementService.NewGetOfferingStatsOptions(
				catalogID,
				offeringID,
			)

			offering, response, err := catalogManagementService.GetOfferingStats(getOfferingStatsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_stats

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())
		})

		It(`ReplaceOffering request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nReplaceOffering() result:")
			// begin-replace_offering

			replaceOfferingOptions := catalogManagementService.NewReplaceOfferingOptions(
				catalogID,
				offeringID,
			)

			offering, response, err := catalogManagementService.ReplaceOffering(replaceOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-replace_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())
		})

		It(`ListOfferings request example`, func() {
			fmt.Println("\nListOfferings() result:")
			// begin-list_offerings

			listOfferingsOptions := catalogManagementService.NewListOfferingsOptions(
				catalogID,
			)
			listOfferingsOptions.Limit = core.Int64Ptr(100)
			listOfferingsOptions.Offset = core.Int64Ptr(0)

			offeringSearchResult, response, err := catalogManagementService.ListOfferings(listOfferingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-list_offerings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringSearchResult).ToNot(BeNil())
		})

		It(`ImportOffering request example`, func() {
			fmt.Println("\nImportOffering() result:")
			// begin-import_offering

			flavorModel := &catalogmanagementv1.Flavor{
				Name:      core.StringPtr("testString"),
				Label:     core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Index:     core.Int64Ptr(int64(38)),
			}

			importOfferingOptions := catalogManagementService.NewImportOfferingOptions(
				catalogID,
			)
			importOfferingOptions.Tags = []string{"go", "sdk"}
			importOfferingOptions.TargetKinds = []string{"terraform"}
			importOfferingOptions.Zipurl = core.StringPtr("https://github.com/IBM-Cloud/terraform-sample/archive/refs/tags/v1.1.0.tar.gz")
			importOfferingOptions.OfferingID = &offeringID
			importOfferingOptions.TargetVersion = core.StringPtr("0.0.2")
			importOfferingOptions.Repotype = core.StringPtr("git_public")
			importOfferingOptions.ProductKind = core.StringPtr("solution")
			importOfferingOptions.Flavor = flavorModel
			importOfferingOptions.InstallType = core.StringPtr("fullstack")
			importOfferingOptions.XAuthToken = &gitAuthTokenForPublicRepo

			offering, response, err := catalogManagementService.ImportOffering(importOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-import_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

			versionLocatorID = *offering.Kinds[0].Versions[0].VersionLocator
			offeringID = *offering.ID
			kindID = *offering.Kinds[0].ID
		})

		It(`GetOfferingChangeNotices request example`, func() {
			Skip("Skip by design.")
			fmt.Println("\nGetOfferingChangeNotices() result:")
			// begin-get_offering_change_notices

			getOfferingChangeNoticesOptionsModel := new(catalogmanagementv1.GetOfferingChangeNoticesOptions)
			getOfferingChangeNoticesOptionsModel.CatalogIdentifier = &catalogID
			getOfferingChangeNoticesOptionsModel.OfferingID = &offeringID
			getOfferingChangeNoticesOptionsModel.Kind = core.StringPtr("terraform")
			getOfferingChangeNoticesOptionsModel.Version = core.StringPtr("1.0.0")

			result, response, err := catalogManagementService.GetOfferingChangeNotices(getOfferingChangeNoticesOptionsModel)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_change_notices

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})

		It(`ReloadOffering request example`, func() {
			Skip("Skip by design.")
			fmt.Println("\nReloadOffering() result:")
			// begin-reload_offering

			reloadOfferingOptions := catalogManagementService.NewReloadOfferingOptions(
				catalogID,
				offeringID,
				"0.0.2",
			)
			reloadOfferingOptions.Tags = []string{"go", "sdk"}
			reloadOfferingOptions.TargetKinds = []string{"roks"}
			reloadOfferingOptions.Zipurl = core.StringPtr("https://github.com/rhm-samples/node-red-operator/blob/master/node-red-operator/bundle/0.0.2/node-red-operator.v0.0.2.clusterserviceversion.yaml")
			reloadOfferingOptions.RepoType = core.StringPtr("git_public")

			offering, response, err := catalogManagementService.ReloadOffering(reloadOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-reload_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())
		})

		It(`CreateObject request example`, func() {
			fmt.Println("\nCreateObject() result:")
			// begin-create_object

			publishObjectModel := &catalogmanagementv1.PublishObject{
				PermitIBMPublicPublish: core.BoolPtr(true),
				IBMApproved:            core.BoolPtr(true),
				PublicApproved:         core.BoolPtr(true),
			}

			stateModel := &catalogmanagementv1.State{
				Current: core.StringPtr("new"),
			}

			createObjectOptions := catalogManagementService.NewCreateObjectOptions(
				objectCatalogID,
			)
			createObjectOptions.CatalogID = &objectCatalogID
			createObjectOptions.Name = core.StringPtr("object_in_ibm_cloud")
			createObjectOptions.CRN = core.StringPtr("crn:v1:bluemix:public:iam-global-endpoint:global:::endpoint:private.iam.cloud.ibm.com")
			createObjectOptions.ParentID = core.StringPtr("us-south")
			createObjectOptions.Kind = core.StringPtr("vpe")
			createObjectOptions.Publish = publishObjectModel
			createObjectOptions.State = stateModel

			catalogObject, response, err := catalogManagementService.CreateObject(createObjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogObject, "", "  ")
			fmt.Println(string(b))

			// end-create_object

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogObject).ToNot(BeNil())

			objectID = *catalogObject.ID
		})

		It(`ListRegions request example`, func() {
			fmt.Println("\nListRegions() result:")
			// begin-list_regions

			listRegionOptions := catalogManagementService.NewListRegionsOptions()

			regions, response, err := catalogManagementService.ListRegions(listRegionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(regions, "", "  ")
			fmt.Println(string(b))

			// end-list_regions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(regions).ToNot(BeNil())
		})

		It(`PreviewRegions request example`, func() {
			fmt.Println("\nPreviewRegions() result:")
			// begin-preview_regions

			previewRegionOptions := catalogManagementService.NewPreviewRegionsOptions()

			regions, response, err := catalogManagementService.PreviewRegions(previewRegionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(regions, "", "  ")
			fmt.Println(string(b))

			// end-preview_regions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(regions).ToNot(BeNil())
		})

		It(`GetCatalogAccount request example`, func() {
			fmt.Println("\nGetCatalogAccount() result:")
			// begin-get_catalog_account

			getCatalogAccountOptions := catalogManagementService.NewGetCatalogAccountOptions()

			account, response, err := catalogManagementService.GetCatalogAccount(getCatalogAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(account, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())

			catalogAccount = account
		})

		It(`UpdateCatalogAccount request example`, func() {
			// Skip("Skipped bby design.")
			// begin-update_catalog_account

			includeAllFilter := &catalogmanagementv1.Filters{
				IncludeAll: core.BoolPtr(true),
			}
			updateCatalogAccountOptions := catalogManagementService.NewUpdateCatalogAccountOptions()
			updateCatalogAccountOptions.Rev = catalogAccount.Rev
			updateCatalogAccountOptions.AccountFilters = includeAllFilter
			updateCatalogAccountOptions.ID = &accountID
			updateCatalogAccountOptions.RegionFilter = core.StringPtr("geo:na")

			_, response, err := catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptions)
			if err != nil {
				panic(err)
			}

			// end-update_catalog_account
			fmt.Printf("\nUpdateCatalogAccount() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`GetCatalogAccountFilters request example`, func() {
			fmt.Println("\nGetCatalogAccountFilters() result:")
			// begin-get_catalog_account_filters

			getCatalogAccountFiltersOptions := catalogManagementService.NewGetCatalogAccountFiltersOptions()

			accumulatedFilters, response, err := catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accumulatedFilters, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_account_filters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accumulatedFilters).ToNot(BeNil())
		})

		It(`AddShareApprovalList request example`, func() {
			fmt.Println("\nAddShareApprovalList() result:")
			// begin-add_share_approval_list

			addShareApprovalListOptions := catalogManagementService.NewAddShareApprovalListOptions(
				"offering",
				[]string{"-acct-testString"},
			)

			accessListBulkResponse, response, err := catalogManagementService.AddShareApprovalList(addShareApprovalListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessListBulkResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_share_approval_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})
		It(`GetShareApprovalList request example`, func() {
			fmt.Println("\nGetShareApprovalList() result:")
			// begin-get_share_approval_list
			getShareApprovalListOptions := &catalogmanagementv1.GetShareApprovalListOptions{
				ObjectType: core.StringPtr("offering"),
				Limit:      core.Int64Ptr(int64(10)),
			}

			pager, err := catalogManagementService.NewGetShareApprovalListPager(getShareApprovalListOptions)
			if err != nil {
				panic(err)
			}

			var allResults []catalogmanagementv1.ShareApprovalAccess
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_share_approval_list
		})
		It(`UpdateShareApprovalListAsSource request example`, func() {
			fmt.Println("\nUpdateShareApprovalListAsSource() result:")
			// begin-update_share_approval_list_as_source

			updateShareApprovalListAsSourceOptions := catalogManagementService.NewUpdateShareApprovalListAsSourceOptions(
				"offering",
				"approved",
				[]string{"-acct-testString"},
			)

			accessListBulkResponse, response, err := catalogManagementService.UpdateShareApprovalListAsSource(updateShareApprovalListAsSourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessListBulkResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_share_approval_list_as_source

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})
		It(`GetShareApprovalListAsSource request example`, func() {
			fmt.Println("\nGetShareApprovalListAsSource() result:")
			// begin-get_share_approval_list_as_source
			getShareApprovalListAsSourceOptions := &catalogmanagementv1.GetShareApprovalListAsSourceOptions{
				ObjectType:              core.StringPtr("offering"),
				ApprovalStateIdentifier: core.StringPtr("approved"),
				Limit:                   core.Int64Ptr(int64(10)),
			}

			pager, err := catalogManagementService.NewGetShareApprovalListAsSourcePager(getShareApprovalListAsSourceOptions)
			if err != nil {
				panic(err)
			}

			var allResults []catalogmanagementv1.ShareApprovalAccess
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_share_approval_list_as_source
		})
		It(`DeleteShareApprovalList request example`, func() {
			fmt.Println("\nDeleteShareApprovalList() result:")
			// begin-delete_share_approval_list

			deleteShareApprovalListOptions := catalogManagementService.NewDeleteShareApprovalListOptions(
				"offering",
				[]string{"-acct-testString"},
			)

			accessListBulkResponse, response, err := catalogManagementService.DeleteShareApprovalList(deleteShareApprovalListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessListBulkResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_share_approval_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})

		It(`GetOfferingSourceArchive request example`, func() {
			fmt.Println("\nGetOfferingSourceArchive() result:")
			// begin-delete_share_approval_list

			getOfferingSourceArchiveOptions := &catalogmanagementv1.GetOfferingSourceArchiveOptions{
				Version:   core.StringPtr("0.0.2"),
				CatalogID: core.StringPtr(catalogID),
				ID:        core.StringPtr(offeringID),
			}

			source, response, err := catalogManagementService.GetOfferingSourceArchive(getOfferingSourceArchiveOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(source, "", "  ")
			fmt.Println(string(b))

			// end-delete_share_approval_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(source).ToNot(BeNil())
		})

		It(`GetConsumptionOfferings request example`, func() {
			fmt.Println("\nGetConsumptionOfferings() result:")
			// begin-get_consumption_offerings

			getConsumptionOfferingsOptions := catalogManagementService.NewGetConsumptionOfferingsOptions()

			offeringSearchResult, response, err := catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_consumption_offerings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringSearchResult).ToNot(BeNil())
		})

		It(`ImportOfferingVersion request example`, func() {
			fmt.Println("\nImportOfferingVersion() result:")
			// begin-import_offering_version

			importOfferingVersionOptions := catalogManagementService.NewImportOfferingVersionOptions(
				catalogID,
				offeringID,
			)
			importOfferingVersionOptions.TargetKinds = []string{"roks"}
			importOfferingVersionOptions.Zipurl = core.StringPtr("https://github.com/rhm-samples/node-red-operator/blob/master/node-red-operator/bundle/0.0.2/node-red-operator.v0.0.2.clusterserviceversion.yaml")
			importOfferingVersionOptions.TargetVersion = core.StringPtr("0.0.3")
			importOfferingVersionOptions.Repotype = core.StringPtr("git_public")

			offering, response, err := catalogManagementService.ImportOfferingVersion(importOfferingVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-import_offering_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())
		})

		It(`GetVersions request example`, func() {
			fmt.Println("\nGetVersions() result:")
			// begin-get_versions

			getVersionsOptions := catalogManagementService.NewGetVersionsOptions(
				catalogID,
				offeringID,
				kindID,
			)

			versions, response, err := catalogManagementService.GetVersions(getVersionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(versions, "", "  ")
			fmt.Println(string(b))

			// end-get_versions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versions).ToNot(BeNil())
		})

		It(`GetVersion request example`, func() {
			fmt.Println("\nGetVersion() result:")
			// begin-get_version

			getVersionOptions := catalogManagementService.NewGetVersionOptions(
				versionLocatorID,
			)

			offering, response, err := catalogManagementService.GetVersion(getVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-get_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

			offeringVersion = offering
		})

		It(`GetVersionDependencies request example`, func() {
			fmt.Println("\nGetVersionDependencies() result:")
			// begin-get_version_dependencies

			getVersionDependenciesOptions := catalogManagementService.NewGetVersionDependenciesOptions(
				versionLocatorID,
			)

			version, response, err := catalogManagementService.GetVersionDependencies(getVersionDependenciesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(version, "", "  ")
			fmt.Println(string(b))

			// end-get_version_dependencies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(version).ToNot(BeNil())
		})

		It(`ValidateInputs request example`, func() {
			fmt.Println("\nValidateInputs() result:")
			// begin-validate_inputs

			validateInputsOptions := catalogManagementService.NewValidateInputsOptions(
				versionLocatorID,
			)
			validateInputsOptions.SetInput1("testString1")
			validateInputsOptions.SetInput2("testString2")

			resp, response, err := catalogManagementService.ValidateInputs(validateInputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp, "", "  ")
			fmt.Println(string(b))

			// end-validate_inputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp).ToNot(BeNil())
		})

		It(`UpdateVersion request example`, func() {
			fmt.Println("\nUpdateVersion() result:")
			// begin-update_version

			updateVersionOptions := catalogManagementService.NewUpdateVersionOptions(
				versionLocatorID,
			)

			updateVersionOptions.ID = offeringVersion.ID
			updateVersionOptions.CatalogID = offeringVersion.CatalogID
			updateVersionOptions.Rev = offeringVersion.Rev
			updateVersionOptions.URL = offeringVersion.URL
			updateVersionOptions.CRN = offeringVersion.CRN
			updateVersionOptions.Label = offeringVersion.Label
			updateVersionOptions.Kinds = offeringVersion.Kinds
			updateVersionOptions.Kinds[0].Versions[0].SolutionInfo.ArchitectureDiagrams = []catalogmanagementv1.ArchitectureDiagram{
				{
					Diagram: &catalogmanagementv1.MediaItem{
						URL:     core.StringPtr("data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTAwIiBoZWlnaHQ9IjEwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8Y2lyY2xlIGN4PSI1MCIgY3k9IjUwIiByPSI0MCIgZmlsbD0icmVkIiAvPgo8L3N2Zz4="),
						Type:    core.StringPtr("image/svg+xml"),
						Caption: core.StringPtr("caption"),
					},
					Description: core.StringPtr("Simple red circle diagram"),
				},
			}

			offering, response, err := catalogManagementService.UpdateVersion(updateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-update_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

			offeringVersion = offering
		})

		It(`PatchUpdateVersion request example`, func() {
			fmt.Println("\nPatchUpdateVersion() result:")
			// begin-patch_update_version

			jsonPatchOperationModel := &catalogmanagementv1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/kinds/0/versions/0/long_description"),
				Value: core.StringPtr("testString"),
			}

			patchUpdateVersionOptions := catalogManagementService.NewPatchUpdateVersionOptions(
				versionLocatorID,
				fmt.Sprintf("\"%s\"", *offeringVersion.Rev),
			)

			patchUpdateVersionOptions.Updates = []catalogmanagementv1.JSONPatchOperation{*jsonPatchOperationModel}

			offering, response, err := catalogManagementService.PatchUpdateVersion(patchUpdateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-patch_update_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

			offeringVersion = offering
		})

		// Offering must be "managed in Partner Center" before we can perform plan operations
		// Done with helper API call as we do not expose this route in our api definition
		It(`SetAllowPublishOffering`, func() {
			headers := map[string]string{
				"X-Approver-Token": approverToken,
			}

			response, err := catalogManagementService.SetAllowPublishOffering(catalogID, offeringID, "publish_approved", true, headers)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		// Must create a plan with an approver token because we only allow Partner Center to create plans
		// Done with helper API call because we do not expose this route in our api definition
		It(`AddPlan`, func() {
			featureModel := &catalogmanagementv1.Feature{
				Title:           core.StringPtr("testString"),
				TitleI18n:       make(map[string]string),
				Description:     core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
			}

			versionRangeModel := new(catalogmanagementv1.VersionRange)
			versionRangeModel.Kinds = []string{"terraform"}
			versionRangeModel.Version = core.StringPtr(">=1.0.0")

			planModel := new(catalogmanagementv1.Plan)
			planModel.Label = core.StringPtr("testString")
			planModel.Name = core.StringPtr("testString")
			planModel.ShortDescription = core.StringPtr("testString")
			planModel.PricingTags = []string{"free"}
			planModel.VersionRange = versionRangeModel
			planModel.Features = []catalogmanagementv1.Feature{*featureModel}
			planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}

			headers := map[string]string{
				"X-Approver-Token": approverToken,
			}

			plan, response, err := catalogManagementService.AddPlan(catalogID, offeringID, planModel, headers)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(plan).ToNot(BeNil())

			planID = *plan.ID

			b, _ := json.MarshalIndent(plan, "", "  ")
			fmt.Println(string(b))
		})

		It(`DeletePlan(deletePlanOptions *DeletePlanOptions)`, func() {
			deletePlanOptions := new(catalogmanagementv1.DeletePlanOptions)
			deletePlanOptions.PlanLocID = &planID

			response, err := catalogManagementService.DeletePlan(deletePlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			fmt.Printf("\nDeletePlan() response status code: %d\n", response.StatusCode)
		})

		// Must create a plan with an approver token because we only allow Partner Center to create plans
		// Done with helper API call because we do not expose this route in our api definition
		It(`AddPlan`, func() {
			featureModel := &catalogmanagementv1.Feature{
				Title:           core.StringPtr("testString"),
				TitleI18n:       make(map[string]string),
				Description:     core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
			}

			versionRangeModel := new(catalogmanagementv1.VersionRange)
			versionRangeModel.Kinds = []string{"terraform"}
			versionRangeModel.Version = core.StringPtr(">=1.0.0")

			planModel := new(catalogmanagementv1.Plan)
			planModel.Label = core.StringPtr("testString")
			planModel.Name = core.StringPtr("testString")
			planModel.ShortDescription = core.StringPtr("testString")
			planModel.PricingTags = []string{"free"}
			planModel.VersionRange = versionRangeModel
			planModel.Features = []catalogmanagementv1.Feature{*featureModel}
			planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}

			headers := map[string]string{
				"X-Approver-Token": approverToken,
			}

			plan, response, err := catalogManagementService.AddPlan(catalogID, offeringID, planModel, headers)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(plan).ToNot(BeNil())

			b, _ := json.MarshalIndent(plan, "", "  ")
			fmt.Println(string(b))

			planID = *plan.ID
		})

		// Must set plan to publish_approved using approver token before other plan operations will work
		// Done with helper API call because we do not expose this route in our api definition
		It(`SetValidatePlan()`, func() {
			headers := map[string]string{
				"X-Approver-Token": approverToken,
			}
			response, err := catalogManagementService.SetValidatePlan(planID, headers)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

			fmt.Printf("\nSetValidatePlan() response status code: %d\n", response.StatusCode)
		})

		// Must set plan to publish_approved using approver token before other plan operations will work
		// Done with helper API call because we do not expose this route in our api definition
		It(`SetAllowPublishPlan()`, func() {
			headers := map[string]string{
				"X-Approver-Token": approverToken,
			}
			response, err := catalogManagementService.SetAllowPublishPlan(planID, headers)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			fmt.Printf("\nSetAllowPublishPlan() response status code: %d\n", response.StatusCode)
		})

		It(`GetPlan(getPlanOptions *GetPlanOptions)`, func() {
			getPlanOptions := new(catalogmanagementv1.GetPlanOptions)
			getPlanOptions.PlanLocID = &planID

			plan, response, err := catalogManagementService.GetPlan(getPlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(plan).ToNot(BeNil())

			b, _ := json.MarshalIndent(plan, "", "  ")
			fmt.Println(string(b))
		})

		It(`ConsumablePlan(consumablePlanOptions *ConsumablePlanOptions)`, func() {
			consumablePlanOptions := new(catalogmanagementv1.ConsumablePlanOptions)
			consumablePlanOptions.PlanLocID = &planID

			response, err := catalogManagementService.ConsumablePlan(consumablePlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

			fmt.Printf("\nConsumablePlan() response status code: %d\n", response.StatusCode)
		})

		It(`SetDeprecatePlan(setDeprecatePlanOptions *SetDeprecatePlanOptions)`, func() {
			setDeprecatePlanOptions := new(catalogmanagementv1.SetDeprecatePlanOptions)
			setDeprecatePlanOptions.PlanLocID = &planID
			setDeprecatePlanOptions.Setting = core.StringPtr("true")

			response, err := catalogManagementService.SetDeprecatePlan(setDeprecatePlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

			fmt.Printf("\nSetDeprecatePlan() response status code: %d\n", response.StatusCode)
		})

		It(`GetOfferingUpdates request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetOfferingUpdates() result:")
			// begin-get_offering_updates

			getOfferingUpdatesOptions := catalogManagementService.NewGetOfferingUpdatesOptions(
				catalogID,
				offeringID,
				"roks",
				"",
			)
			getOfferingUpdatesOptions.Version = core.StringPtr("0.0.2")
			getOfferingUpdatesOptions.ClusterID = &clusterID
			getOfferingUpdatesOptions.Region = core.StringPtr("us-south")
			getOfferingUpdatesOptions.Namespace = core.StringPtr("application-development-namespace")

			versionUpdateDescriptor, response, err := catalogManagementService.GetOfferingUpdates(getOfferingUpdatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(versionUpdateDescriptor, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_updates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionUpdateDescriptor).ToNot(BeNil())
		})

		It(`GetOfferingAbout request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetOfferingAbout() result:")
			// begin-get_offering_about

			getOfferingAboutOptions := catalogManagementService.NewGetOfferingAboutOptions(
				versionLocatorID,
			)

			result, response, err := catalogManagementService.GetOfferingAbout(getOfferingAboutOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_about

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})

		It(`GetIamPermissions request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetIamPermissions() result:")
			// begin-get_iam_permissions

			getIamPermissionsOptions := catalogManagementService.NewGetIamPermissionsOptions(
				versionLocatorID,
			)

			result, response, err := catalogManagementService.GetIamPermissions(getIamPermissionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_iam_permissions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})

		It(`GetOfferingLicense request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetOfferingLicense() result:")
			// begin-get_offering_license

			getOfferingLicenseOptions := catalogManagementService.NewGetOfferingLicenseOptions(
				versionLocatorID,
				"license-id",
			)

			result, response, err := catalogManagementService.GetOfferingLicense(getOfferingLicenseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_license

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})

		It(`GetOfferingContainerImages request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetOfferingContainerImages() result:")
			// begin-get_offering_container_images

			getOfferingContainerImagesOptions := catalogManagementService.NewGetOfferingContainerImagesOptions(
				versionLocatorID,
			)

			imageManifest, response, err := catalogManagementService.GetOfferingContainerImages(getOfferingContainerImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageManifest, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_container_images

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageManifest).ToNot(BeNil())
		})

		It(`TestVersion request example`, func() {
			// begin-test_version

			testVersionOptions := catalogManagementService.NewTestVersionOptions(
				versionLocatorID,
			)

			response, err := catalogManagementService.TestVersion(testVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-test_version
			fmt.Printf("\nTestVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})

		It(`DeprecateVersion request example`, func() {
			// begin-deprecate_version

			deprecateVersionOptions := catalogManagementService.NewDeprecateVersionOptions(
				versionLocatorID,
			)

			response, err := catalogManagementService.DeprecateVersion(deprecateVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-deprecate_version
			fmt.Printf("\nDeprecateVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})

		It(`CommitVersion request example`, func() {
			Skip("Skipped by design.")
			// begin-commit_version

			commitVersionOptions := catalogManagementService.NewCommitVersionOptions(
				versionLocatorID,
			)

			response, err := catalogManagementService.CommitVersion(commitVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-commit_version
			fmt.Printf("\nCommitVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`CopyVersion request example`, func() {
			Skip("Skipped by design.")
			// begin-copy_version

			copyVersionOptions := catalogManagementService.NewCopyVersionOptions(
				versionLocatorID,
			)
			copyVersionOptions.TargetKinds = []string{"roks"}

			response, err := catalogManagementService.CopyVersion(copyVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-copy_version
			fmt.Printf("\nCopyVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`GetOfferingWorkingCopy request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetOfferingWorkingCopy() result:")
			// begin-get_offering_working_copy

			getOfferingWorkingCopyOptions := catalogManagementService.NewGetOfferingWorkingCopyOptions(
				versionLocatorID,
			)

			version, response, err := catalogManagementService.GetOfferingWorkingCopy(getOfferingWorkingCopyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(version, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_working_copy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(version).ToNot(BeNil())
		})

		It(`GetVersion request example`, func() {
			fmt.Println("\nGetVersion() result:")
			// begin-get_version

			getVersionOptions := catalogManagementService.NewGetVersionOptions(
				versionLocatorID,
			)

			offering, response, err := catalogManagementService.GetVersion(getVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-get_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())
		})

		It(`GetCluster request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetCluster() result:")
			// begin-get_cluster

			getClusterOptions := catalogManagementService.NewGetClusterOptions(
				clusterID,
				"us-south",
				bearerToken,
			)

			clusterInfo, response, err := catalogManagementService.GetCluster(getClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clusterInfo, "", "  ")
			fmt.Println(string(b))

			// end-get_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterInfo).ToNot(BeNil())
		})

		It(`GetNamespaces request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetNamespaces() result:")
			// begin-get_namespaces

			getNamespacesOptions := catalogManagementService.NewGetNamespacesOptions(
				clusterID,
				"us-south",
				bearerToken,
			)

			namespaceSearchResult, response, err := catalogManagementService.GetNamespaces(getNamespacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(namespaceSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_namespaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespaceSearchResult).ToNot(BeNil())
		})

		It(`DeployOperators request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nDeployOperators() result:")
			// begin-deploy_operators

			deployOperatorsOptions := catalogManagementService.NewDeployOperatorsOptions(
				bearerToken,
			)
			deployOperatorsOptions.ClusterID = &clusterID
			deployOperatorsOptions.Region = core.StringPtr("us-south")
			deployOperatorsOptions.AllNamespaces = core.BoolPtr(true)
			deployOperatorsOptions.VersionLocatorID = &versionLocatorID

			operatorDeployResult, response, err := catalogManagementService.DeployOperators(deployOperatorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatorDeployResult, "", "  ")
			fmt.Println(string(b))

			// end-deploy_operators

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).ToNot(BeNil())
		})

		It(`ListOperators request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nListOperators() result:")
			// begin-list_operators

			listOperatorsOptions := catalogManagementService.NewListOperatorsOptions(
				bearerToken,
				clusterID,
				"us-south",
				versionLocatorID,
			)

			operatorDeployResult, response, err := catalogManagementService.ListOperators(listOperatorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatorDeployResult, "", "  ")
			fmt.Println(string(b))

			// end-list_operators

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).ToNot(BeNil())
		})

		It(`ReplaceOperators request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nReplaceOperators() result:")
			// begin-replace_operators

			replaceOperatorsOptions := catalogManagementService.NewReplaceOperatorsOptions(
				bearerToken,
			)
			replaceOperatorsOptions.ClusterID = &clusterID
			replaceOperatorsOptions.Region = core.StringPtr("us-south")
			replaceOperatorsOptions.AllNamespaces = core.BoolPtr(true)
			replaceOperatorsOptions.VersionLocatorID = &versionLocatorID

			operatorDeployResult, response, err := catalogManagementService.ReplaceOperators(replaceOperatorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatorDeployResult, "", "  ")
			fmt.Println(string(b))

			// end-replace_operators

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).ToNot(BeNil())
		})

		It(`InstallVersion request example`, func() {
			Skip("Skipped by design.")
			// begin-install_version

			installVersionOptions := catalogManagementService.NewInstallVersionOptions(
				versionLocatorID,
				bearerToken,
			)

			response, err := catalogManagementService.InstallVersion(installVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-install_version
			fmt.Printf("\nInstallVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})

		It(`PreinstallVersion request example`, func() {
			Skip("Skipped by design.")
			// begin-preinstall_version

			preinstallVersionOptions := catalogManagementService.NewPreinstallVersionOptions(
				versionLocatorID,
				bearerToken,
			)

			response, err := catalogManagementService.PreinstallVersion(preinstallVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-preinstall_version
			fmt.Printf("\nPreinstallVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})

		It(`GetPreinstall request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetPreinstall() result:")
			// begin-get_preinstall

			getPreinstallOptions := catalogManagementService.NewGetPreinstallOptions(
				versionLocatorID,
				bearerToken,
			)

			installStatus, response, err := catalogManagementService.GetPreinstall(getPreinstallOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(installStatus, "", "  ")
			fmt.Println(string(b))

			// end-get_preinstall

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(installStatus).ToNot(BeNil())
		})

		It(`ValidateInstall request example`, func() {
			Skip("Skipped by design.")
			// begin-validate_install

			validateInstallOptions := catalogManagementService.NewValidateInstallOptions(
				versionLocatorID,
				bearerToken,
			)

			response, err := catalogManagementService.ValidateInstall(validateInstallOptions)
			if err != nil {
				panic(err)
			}

			// end-validate_install
			fmt.Printf("\nValidateInstall() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})

		It(`GetValidationStatus request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetValidationStatus() result:")
			// begin-get_validation_status

			getValidationStatusOptions := catalogManagementService.NewGetValidationStatusOptions(
				versionLocatorID,
				bearerToken,
			)

			validation, response, err := catalogManagementService.GetValidationStatus(getValidationStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(validation, "", "  ")
			fmt.Println(string(b))

			// end-get_validation_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validation).ToNot(BeNil())
		})

		It(`SearchObjects request example`, func() {
			fmt.Println("\nSearchObjects() result:")
			// begin-search_objects

			searchObjectsOptions := catalogManagementService.NewSearchObjectsOptions(
				"name: object*",
			)
			searchObjectsOptions.Collapse = core.BoolPtr(true)
			searchObjectsOptions.Digest = core.BoolPtr(true)
			searchObjectsOptions.Limit = core.Int64Ptr(100)
			searchObjectsOptions.Offset = core.Int64Ptr(0)

			objectSearchResult, response, err := catalogManagementService.SearchObjects(searchObjectsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-search_objects

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectSearchResult).ToNot(BeNil())
		})

		It(`ListObjects request example`, func() {
			fmt.Println("\nListObjects() result:")
			// begin-list_objects

			listObjectsOptions := catalogManagementService.NewListObjectsOptions(
				objectCatalogID,
			)
			listObjectsOptions.Limit = core.Int64Ptr(100)
			listObjectsOptions.Offset = core.Int64Ptr(0)

			objectListResult, response, err := catalogManagementService.ListObjects(listObjectsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectListResult, "", "  ")
			fmt.Println(string(b))

			// end-list_objects

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectListResult).ToNot(BeNil())
		})

		It(`ReplaceObject request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nReplaceObject() result:")
			// begin-replace_object

			replaceObjectOptions := catalogManagementService.NewReplaceObjectOptions(
				objectCatalogID,
				objectID,
			)
			replaceObjectOptions.ID = &objectID
			replaceObjectOptions.Name = core.StringPtr("updated-object-name")
			replaceObjectOptions.ParentID = core.StringPtr("us-south")
			replaceObjectOptions.Kind = core.StringPtr("vpe")
			replaceObjectOptions.CatalogID = &objectCatalogID

			catalogObject, response, err := catalogManagementService.ReplaceObject(replaceObjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogObject, "", "  ")
			fmt.Println(string(b))

			// end-replace_object

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogObject).ToNot(BeNil())
		})

		It(`GetObject request example`, func() {
			fmt.Println("\nGetObject() result:")
			// begin-get_object

			getObjectOptions := catalogManagementService.NewGetObjectOptions(
				objectCatalogID,
				objectID,
			)

			catalogObject, response, err := catalogManagementService.GetObject(getObjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogObject, "", "  ")
			fmt.Println(string(b))

			// end-get_object

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogObject).ToNot(BeNil())
		})

		It(`GetObjectAccess request example`, func() {
			fmt.Println("\nGetObjectAccess() result:")
			// begin-get_object_access

			getObjectAccessOptions := catalogManagementService.NewGetObjectAccessOptions(
				objectCatalogID,
				objectID,
				accountID,
			)

			objectAccess, response, err := catalogManagementService.GetObjectAccess(getObjectAccessOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectAccess, "", "  ")
			fmt.Println(string(b))

			// end-get_object_access

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectAccess).ToNot(BeNil())
		})

		It(`AddObjectAccessList request example`, func() {
			fmt.Println("\nAddObjectAccessList() result:")
			// begin-add_object_access_list

			addObjectAccessListOptions := catalogManagementService.NewAddObjectAccessListOptions(
				objectCatalogID,
				objectID,
				[]string{accountID},
			)

			accessListBulkResponse, response, err := catalogManagementService.AddObjectAccessList(addObjectAccessListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessListBulkResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_object_access_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})

		It(`GetObjectAccessList request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nGetObjectAccessList() result:")
			// begin-get_object_access_list

			getObjectAccessListOptions := catalogManagementService.NewGetObjectAccessListOptions(
				objectCatalogID,
				objectID,
			)

			objectAccessListResult, response, err := catalogManagementService.GetObjectAccessList(getObjectAccessListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectAccessListResult, "", "  ")
			fmt.Println(string(b))

			// end-get_object_access_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectAccessListResult).ToNot(BeNil())
		})

		It(`CreateOfferingInstance request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nCreateOfferingInstance() result:")
			// begin-create_offering_instance

			createOfferingInstanceOptions := catalogManagementService.NewCreateOfferingInstanceOptions(
				bearerToken,
			)
			createOfferingInstanceOptions.ID = &offeringID
			createOfferingInstanceOptions.CatalogID = &catalogID
			createOfferingInstanceOptions.OfferingID = &offeringID
			createOfferingInstanceOptions.KindFormat = core.StringPtr("vpe")
			createOfferingInstanceOptions.Version = core.StringPtr("0.0.2")
			createOfferingInstanceOptions.ClusterID = &clusterID
			createOfferingInstanceOptions.ClusterRegion = core.StringPtr("us-south")
			createOfferingInstanceOptions.ClusterAllNamespaces = core.BoolPtr(true)
			createOfferingInstanceOptions.ParentCRN = core.StringPtr("testString")
			createOfferingInstanceOptions.ClusterRegion = core.StringPtr("testString")

			offeringInstance, response, err := catalogManagementService.CreateOfferingInstance(createOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringInstance, "", "  ")
			fmt.Println(string(b))

			// end-create_offering_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offeringInstance).ToNot(BeNil())

			offeringInstanceID = *offeringInstance.ID
		})

		It(`GetOfferingInstance request example`, func() {
			fmt.Println("\nGetOfferingInstance() result:")
			Skip("Skipped by design.")
			// begin-get_offering_instance

			getOfferingInstanceOptions := catalogManagementService.NewGetOfferingInstanceOptions(
				offeringInstanceID,
			)

			offeringInstance, response, err := catalogManagementService.GetOfferingInstance(getOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringInstance, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringInstance).ToNot(BeNil())
		})

		It(`PutOfferingInstance request example`, func() {
			Skip("Skipped by design.")
			fmt.Println("\nPutOfferingInstance() result:")
			// begin-put_offering_instance

			putOfferingInstanceOptions := catalogManagementService.NewPutOfferingInstanceOptions(
				offeringInstanceID,
				bearerToken,
			)
			putOfferingInstanceOptions.ID = &offeringID
			putOfferingInstanceOptions.CatalogID = &catalogID
			putOfferingInstanceOptions.OfferingID = &offeringID
			putOfferingInstanceOptions.KindFormat = core.StringPtr("vpe")
			putOfferingInstanceOptions.Version = core.StringPtr("0.0.2")
			putOfferingInstanceOptions.ClusterID = &clusterID
			putOfferingInstanceOptions.ClusterRegion = core.StringPtr("us-south")
			putOfferingInstanceOptions.ClusterAllNamespaces = core.BoolPtr(true)
			putOfferingInstanceOptions.ParentCRN = core.StringPtr("testString")
			putOfferingInstanceOptions.ClusterRegion = core.StringPtr("testString")

			offeringInstance, response, err := catalogManagementService.PutOfferingInstance(putOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringInstance, "", "  ")
			fmt.Println(string(b))

			// end-put_offering_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringInstance).ToNot(BeNil())
		})

		It(`DeleteVersion request example`, func() {
			// begin-delete_version

			deleteVersionOptions := catalogManagementService.NewDeleteVersionOptions(
				versionLocatorID,
			)

			response, err := catalogManagementService.DeleteVersion(deleteVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_version
			fmt.Printf("\nDeleteVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`DeleteOperators request example`, func() {
			Skip("Skipped by design.")
			// begin-delete_operators

			deleteOperatorsOptions := catalogManagementService.NewDeleteOperatorsOptions(
				bearerToken,
				clusterID,
				"us-south",
				versionLocatorID,
			)

			response, err := catalogManagementService.DeleteOperators(deleteOperatorsOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_operators
			fmt.Printf("\nDeleteOperators() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`DeleteOfferingInstance request example`, func() {
			Skip("Skipped by design.")
			// begin-delete_offering_instance

			deleteOfferingInstanceOptions := catalogManagementService.NewDeleteOfferingInstanceOptions(
				offeringInstanceID,
				bearerToken,
			)

			response, err := catalogManagementService.DeleteOfferingInstance(deleteOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_offering_instance
			fmt.Printf("\nDeleteOfferingInstance() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`DeleteObjectAccessList request example`, func() {
			fmt.Println("\nDeleteObjectAccessList() result:")
			// begin-delete_object_access_list

			deleteObjectAccessListOptions := catalogManagementService.NewDeleteObjectAccessListOptions(
				objectCatalogID,
				objectID,
				[]string{accountID},
			)

			accessListBulkResponse, response, err := catalogManagementService.DeleteObjectAccessList(deleteObjectAccessListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessListBulkResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_object_access_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})

		It(`DeleteObject request example`, func() {
			// begin-delete_object

			deleteObjectOptions := catalogManagementService.NewDeleteObjectOptions(
				objectCatalogID,
				objectID,
			)

			response, err := catalogManagementService.DeleteObject(deleteObjectOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_object
			fmt.Printf("\nDeleteObject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		// Unset pc managed
		It(`SetAllowPublishOffering`, func() {
			headers := map[string]string{
				"X-Approver-Token": approverToken,
			}

			response, err := catalogManagementService.SetAllowPublishOffering(catalogID, offeringID, "pc_managed", false, headers)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`DeleteOffering request example`, func() {
			// begin-delete_offering

			deleteOfferingOptions := catalogManagementService.NewDeleteOfferingOptions(
				catalogID,
				offeringID,
			)

			response, err := catalogManagementService.DeleteOffering(deleteOfferingOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_offering
			fmt.Printf("\nDeleteOffering() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`DeleteCatalog for offerings request example`, func() {
			// begin-delete_catalog

			deleteCatalogOptions := catalogManagementService.NewDeleteCatalogOptions(
				catalogID,
			)

			response, err := catalogManagementService.DeleteCatalog(deleteCatalogOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_catalog
			fmt.Printf("\nDeleteCatalog() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`DeleteCatalog for objects request example`, func() {
			// begin-delete_catalog

			deleteCatalogOptions := catalogManagementService.NewDeleteCatalogOptions(
				objectCatalogID,
			)

			response, err := catalogManagementService.DeleteCatalog(deleteCatalogOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_catalog
			fmt.Printf("\nDeleteCatalog() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})
