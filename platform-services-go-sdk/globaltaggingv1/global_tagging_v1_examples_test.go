//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Global Tagging service.
//
// The following configuration properties are assumed to be defined:
//
// GLOBAL_TAGGING_URL=<service url>
// GLOBAL_TAGGING_AUTHTYPE=iam
// GLOBAL_TAGGING_APIKEY=<IAM api key>
// GLOBAL_TAGGING_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// GLOBAL_TAGGING_RESOURCE_CRN=<the crn of the resource to be used in the examples>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`GlobalTaggingV1 Examples Tests`, func() {
	const externalConfigFile = "../global_tagging.env"

	var (
		globalTaggingService *globaltaggingv1.GlobalTaggingV1
		config               map[string]string
		configLoaded         bool

		resourceCRN string
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
			config, err = core.GetServiceProperties(globaltaggingv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			resourceCRN = config["RESOURCE_CRN"]
			if resourceCRN == "" {
				Skip("Unable to load RESOURCE_CRN configuration property, skipping tests")
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

			globalTaggingServiceOptions := &globaltaggingv1.GlobalTaggingV1Options{}

			globalTaggingService, err = globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(globalTaggingServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(globalTaggingService).ToNot(BeNil())
		})
	})

	Describe(`GlobalTaggingV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTag request example`, func() {
			fmt.Println("\nCreateTag() result:")
			// begin-create_tag

			createTagOptions := globalTaggingService.NewCreateTagOptions(
				[]string{"env:example-access-tag"},
			)
			createTagOptions.SetTagType("access")

			createTagResults, response, err := globalTaggingService.CreateTag(createTagOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTagResults, "", "  ")
			fmt.Println(string(b))

			// end-create_tag

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTagResults).ToNot(BeNil())
		})
		It(`ListTags request example`, func() {
			fmt.Println("\nListTags() result:")
			// begin-list_tags

			listTagsOptions := globalTaggingService.NewListTagsOptions()
			listTagsOptions.SetTagType("user")
			listTagsOptions.SetAttachedOnly(true)
			listTagsOptions.SetFullData(true)
			listTagsOptions.SetProviders([]string{"ghost"})
			listTagsOptions.SetOrderByName("asc")

			tagList, response, err := globalTaggingService.ListTags(listTagsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tagList, "", "  ")
			fmt.Println(string(b))

			// end-list_tags

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagList).ToNot(BeNil())
		})
		It(`AttachTag request example`, func() {
			fmt.Println("\nAttachTag() result:")
			// begin-attach_tag

			attachTagOptions := globalTaggingService.NewAttachTagOptions()
			attachTagOptions.SetTagNames([]string{"tag_test_1", "tag_test_2"})
			attachTagOptions.SetTagType("user")

			tagResults, response, err := globalTaggingService.AttachTag(attachTagOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tagResults, "", "  ")
			fmt.Println(string(b))

			// end-attach_tag

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())

		})
		It(`DetachTag request example`, func() {
			fmt.Println("\nDetachTag() result:")
			// begin-detach_tag

			detachTagOptions := globalTaggingService.NewDetachTagOptions()
			detachTagOptions.SetTagNames([]string{"tag_test_1", "tag_test_2"})
			detachTagOptions.SetTagType("user")

			tagResults, response, err := globalTaggingService.DetachTag(detachTagOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tagResults, "", "  ")
			fmt.Println(string(b))

			// end-detach_tag

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())

		})
		It(`DeleteTag request example`, func() {
			fmt.Println("\nDeleteTag() result:")
			// begin-delete_tag

			deleteTagOptions := globalTaggingService.NewDeleteTagOptions("env:example-access-tag")
			deleteTagOptions.SetTagType("access")

			deleteTagResults, response, err := globalTaggingService.DeleteTag(deleteTagOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteTagResults, "", "  ")
			fmt.Println(string(b))

			// end-delete_tag

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagResults).ToNot(BeNil())
		})
		It(`DeleteTagAll request example`, func() {
			fmt.Println("\nDeleteTagAll() result:")
			// begin-delete_tag_all

			deleteTagAllOptions := globalTaggingService.NewDeleteTagAllOptions()
			deleteTagAllOptions.SetTagType("user")

			deleteTagsResult, response, err := globalTaggingService.DeleteTagAll(deleteTagAllOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteTagsResult, "", "  ")
			fmt.Println(string(b))

			// end-delete_tag_all

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagsResult).ToNot(BeNil())
		})
	})
})
