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

// Package dnssvcsv1 : Operations and models for the DnsSvcsV1 service
package dnssvcsv1

import (
	"github.com/IBM/go-sdk-core/core"
)

// DnsSvcsV1 : DNS svcs
//
// Version: 1.0.0
type DnsSvcsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.dns-svcs.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "dns-svcs"

// DnsSvcsV1Options : Service options
type DnsSvcsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDnsSvcsV1UsingExternalConfig : constructs an instance of DnsSvcsV1 with passed in options and external configuration.
func NewDnsSvcsV1UsingExternalConfig(options *DnsSvcsV1Options) (dnsSvcs *DnsSvcsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	dnsSvcs, err = NewDnsSvcsV1(options)
	if err != nil {
		return
	}

	err = dnsSvcs.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = dnsSvcs.Service.SetServiceURL(options.URL)
	}
	return
}

// NewDnsSvcsV1 : constructs an instance of DnsSvcsV1 with passed in options.
func NewDnsSvcsV1(options *DnsSvcsV1Options) (service *DnsSvcsV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &DnsSvcsV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (dnsSvcs *DnsSvcsV1) SetServiceURL(url string) error {
	return dnsSvcs.Service.SetServiceURL(url)
}
