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

// Package ibmcloudfunctionsnamespaceapiv1 : Operations and models for the ibmcloudfunctionsnamespaceapiv1 service
package ibmcloudfunctionsnamespaceapiv1

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v3/core"
)

// common headers
const (
	ContentType         = "Content-Type"
	FormEncodeHeaderURL = "application/x-www-form-urlencoded"
)

// GetCloudFoundaryNamespaces : Retrieve all IBM Cloud Functions namespaces (Only classic)
// Compatibility: If passing basic authorization instead of an IAM access token the classic namespace associated with
// these authorization credentials is returned.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) GetCloudFoundaryNamespaces(getNamespacesOptions *GetNamespacesOptions) (result *NamespaceResponseList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getNamespacesOptions, "getNamespacesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"bluemix", "v2", "authenticate"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	builder.AddHeader(ContentType, FormEncodeHeaderURL)
	for k, v := range getNamespacesOptions.Headers {
		builder.AddFormData(k, "", "", v)
	}

	if getNamespacesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getNamespacesOptions.Limit))
	}
	if getNamespacesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getNamespacesOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceResponseList))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*NamespaceResponseList)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}
