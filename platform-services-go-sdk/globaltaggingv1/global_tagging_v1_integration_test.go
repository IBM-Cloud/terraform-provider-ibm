//go:build integration

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

package globaltaggingv1_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the globaltaggingv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`GlobalTaggingV1 Integration Tests`, func() {

	const (
		externalConfigFile = "../global_tagging.env"
		sdkLabel           = "go-sdk"
	)

	var resourceCRN string
	var (
		err                  error
		globalTaggingService *globaltaggingv1.GlobalTaggingV1
		serviceURL           string
		config               map[string]string

		userTag1   string = fmt.Sprintf("%s-user-test1", sdkLabel)
		userTag2   string = fmt.Sprintf("%s-user-test2", sdkLabel)
		accessTag1 string = fmt.Sprintf("env:%s-public", sdkLabel)
		accessTag2 string = fmt.Sprintf("region:%s-us-south", sdkLabel)
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
			config, err = core.GetServiceProperties(globaltaggingv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			resourceCRN = config["RESOURCE_CRN"]
			if resourceCRN == "" {
				Skip("Unable to load RESOURCE_CRN configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "\nService URL: %s\n", serviceURL)
			fmt.Fprintf(GinkgoWriter, "Resource ID: %s\n", resourceCRN)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			globalTaggingServiceOptions := &globaltaggingv1.GlobalTaggingV1Options{}

			globalTaggingService, err = globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(globalTaggingServiceOptions)

			Expect(err).To(BeNil())
			Expect(globalTaggingService).ToNot(BeNil())
			Expect(globalTaggingService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			globalTaggingService.EnableRetries(4, 30*time.Second)
		})

		It("Successfully setup the environment for tests", func() {
			fmt.Fprintln(GinkgoWriter, "Setup...")
			cleanTags(globalTaggingService, resourceCRN, sdkLabel)
			fmt.Fprintln(GinkgoWriter, "Finished setup.")
		})
	})

	Describe(`CreateTag - Create an access tag`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTag(createTagOptions *CreateTagOptions)`, func() {

			createTagOptions := &globaltaggingv1.CreateTagOptions{
				TagNames: []string{accessTag1, accessTag2},
				TagType:  core.StringPtr("access"),
			}

			createTagResults, response, err := globalTaggingService.CreateTag(createTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTagResults).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nCreateTag() response:\n%s", common.ToJSON(createTagResults))

			for _, elem := range createTagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}
		})
	})

	Describe(`AttachTag - Attach tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AttachTag(user tags)`, func() {

			resourceModel := &globaltaggingv1.Resource{
				ResourceID: &resourceCRN,
			}

			attachTagOptions := &globaltaggingv1.AttachTagOptions{
				Resources: []globaltaggingv1.Resource{*resourceModel},
				TagNames:  []string{userTag1, userTag2},
				TagType:   core.StringPtr("user"),
			}

			tagResults, response, err := globalTaggingService.AttachTag(attachTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nAttachTag(user) response:\n%s\n", common.ToJSON(tagResults))

			Expect(tagResults.Results).ToNot(BeEmpty())
			for _, elem := range tagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// Make sure the tag was in fact attached.
			tagNames := getTagNamesForResource(globalTaggingService, resourceCRN, "user")
			fmt.Fprintln(GinkgoWriter, "\nResource now has these user tags: ", tagNames)
			Expect(tagNames).ToNot(BeEmpty())
			Expect(tagNames).To(ContainElement(userTag1))
			Expect(tagNames).To(ContainElement(userTag2))
		})

		It(`AttachTag(access tags)`, func() {

			resourceModel := &globaltaggingv1.Resource{
				ResourceID: &resourceCRN,
			}

			attachTagOptions := &globaltaggingv1.AttachTagOptions{
				Resources: []globaltaggingv1.Resource{*resourceModel},
				TagNames:  []string{accessTag1, accessTag2},
				TagType:   core.StringPtr("access"),
			}

			tagResults, response, err := globalTaggingService.AttachTag(attachTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nAttachTag(access) response:\n%s\n", common.ToJSON(tagResults))

			Expect(tagResults.Results).ToNot(BeEmpty())
			for _, elem := range tagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// Make sure the tag was in fact attached.
			tagNames := getTagNamesForResource(globalTaggingService, resourceCRN, "access")
			fmt.Fprintln(GinkgoWriter, "\nResource now has these access tags: ", tagNames)
			Expect(tagNames).ToNot(BeEmpty())
			Expect(tagNames).To(ContainElement(accessTag1))
			Expect(tagNames).To(ContainElement(accessTag2))
		})

	})

	Describe(`ListTags - Get all tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTags(user tags)`, func() {

			// Retrieve the search results 1 page at a time to test the pagination.

			var results []globaltaggingv1.Tag = make([]globaltaggingv1.Tag, 0)
			var offset int64 = 0
			var moreResults bool = true

			listTagsOptions := &globaltaggingv1.ListTagsOptions{
				Limit:   core.Int64Ptr(500),
				TagType: core.StringPtr("user"),
			}

			for moreResults {
				listTagsOptions.SetOffset(offset)
				tagList, response, err := globalTaggingService.ListTags(listTagsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(tagList).ToNot(BeNil())

				if len(tagList.Items) > 0 {
					results = append(results, tagList.Items...)
					offset += int64(len(tagList.Items))
				} else {
					moreResults = false
				}
			}

			fmt.Fprintf(GinkgoWriter, "\nRetrieved a total of %d user tags.\n", len(results))
			var matches []string = make([]string, 0)
			for _, tag := range results {
				if strings.Contains(*tag.Name, sdkLabel) {
					matches = append(matches, *tag.Name)
				}
			}
			fmt.Fprintf(GinkgoWriter, "Found %d user tags containing our label: ", len(matches))
			fmt.Fprintln(GinkgoWriter, matches)
		})

		It(`ListTags(access tags)`, func() {

			// Retrieve the search results 1 page at a time to test the pagination.

			var results []globaltaggingv1.Tag = make([]globaltaggingv1.Tag, 0)
			var offset int64 = 0
			var moreResults bool = true

			listTagsOptions := &globaltaggingv1.ListTagsOptions{
				Limit:   core.Int64Ptr(500),
				TagType: core.StringPtr("access"),
			}

			for moreResults {
				listTagsOptions.SetOffset(offset)
				tagList, response, err := globalTaggingService.ListTags(listTagsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(tagList).ToNot(BeNil())

				if len(tagList.Items) > 0 {
					results = append(results, tagList.Items...)
					offset += int64(len(tagList.Items))
				} else {
					moreResults = false
				}
			}

			fmt.Fprintf(GinkgoWriter, "\nRetrieved a total of %d access tags.\n", len(results))
			var matches []string = make([]string, 0)
			for _, tag := range results {
				if strings.Contains(*tag.Name, sdkLabel) {
					matches = append(matches, *tag.Name)
				}
			}
			fmt.Fprintf(GinkgoWriter, "Found %d access tags containing our label: ", len(matches))
			fmt.Fprintln(GinkgoWriter, matches)
		})
	})

	Describe(`DetachTag - Detach tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DetachTag(user tag)`, func() {

			resourceModel := &globaltaggingv1.Resource{
				ResourceID: &resourceCRN,
			}

			detachTagOptions := &globaltaggingv1.DetachTagOptions{
				Resources: []globaltaggingv1.Resource{*resourceModel},
				TagNames:  []string{userTag1, userTag2},
				TagType:   core.StringPtr("user"),
			}

			tagResults, response, err := globalTaggingService.DetachTag(detachTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nDetachTag(user) response:\n%s\n", common.ToJSON(tagResults))

			Expect(tagResults.Results).ToNot(BeEmpty())
			for _, elem := range tagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// Make sure the tags were detached.
			tagNames := getTagNamesForResource(globalTaggingService, resourceCRN, "user")
			fmt.Fprintln(GinkgoWriter, "\nResource now has these user tags: ", tagNames)
			Expect(tagNames).ToNot(ContainElement(userTag1))
			Expect(tagNames).ToNot(ContainElement(userTag2))
		})

		It(`DetachTag(access tag)`, func() {

			resourceModel := &globaltaggingv1.Resource{
				ResourceID: &resourceCRN,
			}

			detachTagOptions := &globaltaggingv1.DetachTagOptions{
				Resources: []globaltaggingv1.Resource{*resourceModel},
				TagNames:  []string{accessTag1, accessTag2},
				TagType:   core.StringPtr("access"),
			}

			tagResults, response, err := globalTaggingService.DetachTag(detachTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nDetachTag(access) response:\n%s\n", common.ToJSON(tagResults))

			Expect(tagResults.Results).ToNot(BeEmpty())
			for _, elem := range tagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// Make sure the first tag was in fact detached and the second tag remains attached.
			tagNames := getTagNamesForResource(globalTaggingService, resourceCRN, "access")
			fmt.Fprintln(GinkgoWriter, "\nResource now has these access tags: ", tagNames)
			Expect(tagNames).ToNot(ContainElement(accessTag1))
			Expect(tagNames).ToNot(ContainElement(accessTag2))
		})
	})

	Describe(`DeleteTag - Delete an unused tag`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTag(user tag)`, func() {

			deleteTagOptions := &globaltaggingv1.DeleteTagOptions{
				TagName: &userTag1,
				TagType: core.StringPtr("user"),
			}

			deleteTagResults, response, err := globalTaggingService.DeleteTag(deleteTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagResults).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nDeleteTag(user) response:\n%s\n", common.ToJSON(deleteTagResults))

			for _, elem := range deleteTagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}
		})

		It(`DeleteTag(access tag)`, func() {

			deleteTagOptions := &globaltaggingv1.DeleteTagOptions{
				TagName: &accessTag1,
				TagType: core.StringPtr("access"),
			}

			deleteTagResults, response, err := globalTaggingService.DeleteTag(deleteTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagResults).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nDeleteTag(access) response:\n%s\n", common.ToJSON(deleteTagResults))

			for _, elem := range deleteTagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}
		})
	})

	Describe(`DeleteTagAll - Delete all unused tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTagAll(user tags)`, func() {
			deleteTagAllOptions := &globaltaggingv1.DeleteTagAllOptions{
				TagType: core.StringPtr("user"),
			}
			deleteTagsResult, response, err := globalTaggingService.DeleteTagAll(deleteTagAllOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagsResult).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "\nDeleteTagAll(user) response:\n%s\n", common.ToJSON(deleteTagsResult))
			for _, elem := range deleteTagsResult.Items {
				Expect(*elem.IsError).To(Equal(false))
			}
		})

		It(`DeleteTagAll(access tags)`, func() {
			deleteTagAllOptions := &globaltaggingv1.DeleteTagAllOptions{
				TagType: core.StringPtr("access"),
			}
			deleteTagsResult, response, err := globalTaggingService.DeleteTagAll(deleteTagAllOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagsResult).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "\nDeleteTagAll(access) response:\n%s\n", common.ToJSON(deleteTagsResult))
			for _, elem := range deleteTagsResult.Items {
				Expect(*elem.IsError).To(Equal(false))
			}
		})
	})

	Describe(`Teardown - clean up test data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Clean rules`, func() {
			fmt.Fprintln(GinkgoWriter, "Teardown...")
			cleanTags(globalTaggingService, resourceCRN, sdkLabel)
			fmt.Fprintln(GinkgoWriter, "Finished teardown.")
		})
	})
})

func getTagNamesForResource(service *globaltaggingv1.GlobalTaggingV1, resourceID string, tagType string) []string {
	listTagsOptions := &globaltaggingv1.ListTagsOptions{
		AttachedTo: &resourceID,
		TagType:    &tagType,
	}
	tagList, response, err := service.ListTags(listTagsOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	tagNames := []string{}
	for _, tag := range tagList.Items {
		tagNames = append(tagNames, *tag.Name)
	}

	return tagNames
}

func deleteTag(service *globaltaggingv1.GlobalTaggingV1, tag string, tagType string) {
	deleteTagOptions := &globaltaggingv1.DeleteTagOptions{
		TagName: &tag,
		TagType: &tagType,
	}
	deleteTagResults, response, err := service.DeleteTag(deleteTagOptions)

	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))
	Expect(deleteTagResults).ToNot(BeNil())
	for _, elem := range deleteTagResults.Results {
		Expect(*elem.IsError).To(Equal(false))
	}
}

func detachTag(service *globaltaggingv1.GlobalTaggingV1, resourceID string, tag string, tagType string) {
	resourceModel := &globaltaggingv1.Resource{
		ResourceID: &resourceID,
	}

	detachTagOptions := &globaltaggingv1.DetachTagOptions{
		Resources: []globaltaggingv1.Resource{*resourceModel},
		TagNames:  []string{tag},
		TagType:   &tagType,
	}

	tagResults, response, err := service.DetachTag(detachTagOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))
	Expect(tagResults).ToNot(BeNil())
	for _, elem := range tagResults.Results {
		Expect(*elem.IsError).To(Equal(false))
	}
}

func listTagsWithLabel(service *globaltaggingv1.GlobalTaggingV1, tagType string, label string) []string {
	var tagNames []string = make([]string, 0)
	var offset int64 = 0
	var moreResults bool = true
	listTagsOptions := &globaltaggingv1.ListTagsOptions{
		Limit:   core.Int64Ptr(1000),
		TagType: &tagType,
	}

	for moreResults {
		listTagsOptions.SetOffset(offset)
		tagList, response, err := service.ListTags(listTagsOptions)

		Expect(err).To(BeNil())
		Expect(response.StatusCode).To(Equal(200))
		Expect(tagList).ToNot(BeNil())

		if len(tagList.Items) > 0 {
			for _, tag := range tagList.Items {
				if strings.Contains(*tag.Name, label) {
					tagNames = append(tagNames, *tag.Name)
				}
			}
			offset += int64(len(tagList.Items))
		} else {
			moreResults = false
		}
	}
	return tagNames
}

func cleanTags(service *globaltaggingv1.GlobalTaggingV1, resourceCRN string, sdkLabel string) {
	// Detach all user and access tags that contain our label.
	userTags := getTagNamesForResource(service, resourceCRN, "user")
	for _, tagName := range userTags {
		if strings.Contains(tagName, sdkLabel) {
			detachTag(service, resourceCRN, tagName, "user")
			fmt.Fprintf(GinkgoWriter, "Detached user tag %s from resource %s\n", tagName, resourceCRN)
		}
	}
	userTags = getTagNamesForResource(service, resourceCRN, "user")
	fmt.Fprintln(GinkgoWriter, "\nResource now has these user tags: ", userTags)

	accessTags := getTagNamesForResource(service, resourceCRN, "access")
	for _, tagName := range accessTags {
		if strings.Contains(tagName, sdkLabel) {
			detachTag(service, resourceCRN, tagName, "access")
			fmt.Fprintf(GinkgoWriter, "Detached access tag %s from resource %s\n", tagName, resourceCRN)
		}
	}
	accessTags = getTagNamesForResource(service, resourceCRN, "access")
	fmt.Fprintln(GinkgoWriter, "\nResource now has these access tags: ", accessTags)

	// Delete all user and access tags that contain our label.
	userTags = listTagsWithLabel(service, "user", sdkLabel)
	fmt.Fprintf(GinkgoWriter, "Found %d user tag(s) that contain our label.\n", len(userTags))
	for _, tagName := range userTags {
		deleteTag(service, tagName, "user")
		fmt.Fprintf(GinkgoWriter, "Deleted user tag: %s\n", tagName)
	}

	accessTags = listTagsWithLabel(service, "access", sdkLabel)
	fmt.Fprintf(GinkgoWriter, "Found %d access tag(s) that contain our label.\n", len(accessTags))
	for _, tagName := range accessTags {
		deleteTag(service, tagName, "access")
		fmt.Fprintf(GinkgoWriter, "Deleted access tag: %s\n", tagName)
	}
}
