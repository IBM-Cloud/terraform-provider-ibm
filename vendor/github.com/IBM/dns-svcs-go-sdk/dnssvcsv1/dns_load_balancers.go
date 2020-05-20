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
	"fmt"

	common "github.com/IBM/dns-svcs-go-sdk/common"
	"github.com/IBM/go-sdk-core/core"
)

// ListLoadBalancers : List load balancers
// List the Global Load Balancers for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions) (result *ListLoadBalancers, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listLoadBalancersOptions, "listLoadBalancersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listLoadBalancersOptions, "listLoadBalancersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*listLoadBalancersOptions.InstanceID, *listLoadBalancersOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLoadBalancersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListLoadBalancers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listLoadBalancersOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listLoadBalancersOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalListLoadBalancers(m)
		response.Result = result
	}

	return
}

// CreateLoadBalancer : Create a load balancer
// Create a load balancer for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createLoadBalancerOptions, "createLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createLoadBalancerOptions, "createLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*createLoadBalancerOptions.InstanceID, *createLoadBalancerOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createLoadBalancerOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createLoadBalancerOptions.Name != nil {
		body["name"] = createLoadBalancerOptions.Name
	}
	if createLoadBalancerOptions.Description != nil {
		body["description"] = createLoadBalancerOptions.Description
	}
	if createLoadBalancerOptions.Enabled != nil {
		body["enabled"] = createLoadBalancerOptions.Enabled
	}
	if createLoadBalancerOptions.TTL != nil {
		body["ttl"] = createLoadBalancerOptions.TTL
	}
	if createLoadBalancerOptions.FallbackPool != nil {
		body["fallback_pool"] = createLoadBalancerOptions.FallbackPool
	}
	if createLoadBalancerOptions.DefaultPools != nil {
		body["default_pools"] = createLoadBalancerOptions.DefaultPools
	}
	if createLoadBalancerOptions.AzPools != nil {
		body["az_pools"] = createLoadBalancerOptions.AzPools
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalLoadBalancer(m)
		response.Result = result
	}

	return
}

// DeleteLoadBalancer : Delete a load balancer
// Delete a load balancer.
func (dnsSvcs *DnsSvcsV1) DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLoadBalancerOptions, "deleteLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLoadBalancerOptions, "deleteLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*deleteLoadBalancerOptions.InstanceID, *deleteLoadBalancerOptions.DnszoneID, *deleteLoadBalancerOptions.LbID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteLoadBalancerOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetLoadBalancer : Get a load balancer
// Get details of a load balancer.
func (dnsSvcs *DnsSvcsV1) GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLoadBalancerOptions, "getLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLoadBalancerOptions, "getLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*getLoadBalancerOptions.InstanceID, *getLoadBalancerOptions.DnszoneID, *getLoadBalancerOptions.LbID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getLoadBalancerOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalLoadBalancer(m)
		response.Result = result
	}

	return
}

// UpdateLoadBalancer : Update the properties of a load balancer
// Update the properties of a load balancer.
func (dnsSvcs *DnsSvcsV1) UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateLoadBalancerOptions, "updateLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateLoadBalancerOptions, "updateLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*updateLoadBalancerOptions.InstanceID, *updateLoadBalancerOptions.DnszoneID, *updateLoadBalancerOptions.LbID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateLoadBalancerOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateLoadBalancerOptions.Name != nil {
		body["name"] = updateLoadBalancerOptions.Name
	}
	if updateLoadBalancerOptions.Description != nil {
		body["description"] = updateLoadBalancerOptions.Description
	}
	if updateLoadBalancerOptions.Enabled != nil {
		body["enabled"] = updateLoadBalancerOptions.Enabled
	}
	if updateLoadBalancerOptions.TTL != nil {
		body["ttl"] = updateLoadBalancerOptions.TTL
	}
	if updateLoadBalancerOptions.FallbackPool != nil {
		body["fallback_pool"] = updateLoadBalancerOptions.FallbackPool
	}
	if updateLoadBalancerOptions.DefaultPools != nil {
		body["default_pools"] = updateLoadBalancerOptions.DefaultPools
	}
	if updateLoadBalancerOptions.AzPools != nil {
		body["az_pools"] = updateLoadBalancerOptions.AzPools
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalLoadBalancer(m)
		response.Result = result
	}

	return
}

// ListPools : List load balancer pools
// List the load balancer pools.
func (dnsSvcs *DnsSvcsV1) ListPools(listPoolsOptions *ListPoolsOptions) (result *ListPools, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPoolsOptions, "listPoolsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPoolsOptions, "listPoolsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*listPoolsOptions.InstanceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPoolsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListPools")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPoolsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listPoolsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalListPools(m)
		response.Result = result
	}

	return
}

// CreatePool : Create a load balancer pool
// Create a load balancer pool.
func (dnsSvcs *DnsSvcsV1) CreatePool(createPoolOptions *CreatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPoolOptions, "createPoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPoolOptions, "createPoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*createPoolOptions.InstanceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreatePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createPoolOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createPoolOptions.Name != nil {
		body["name"] = createPoolOptions.Name
	}
	if createPoolOptions.Description != nil {
		body["description"] = createPoolOptions.Description
	}
	if createPoolOptions.Enabled != nil {
		body["enabled"] = createPoolOptions.Enabled
	}
	if createPoolOptions.MinimumOrigins != nil {
		body["minimum_origins"] = createPoolOptions.MinimumOrigins
	}
	if createPoolOptions.Origins != nil {
		body["origins"] = createPoolOptions.Origins
	}
	if createPoolOptions.Monitor != nil {
		body["monitor"] = createPoolOptions.Monitor
	}
	if createPoolOptions.NotificationType != nil {
		body["notification_type"] = createPoolOptions.NotificationType
	}
	if createPoolOptions.NotificationChannel != nil {
		body["notification_channel"] = createPoolOptions.NotificationChannel
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalPool(m)
		response.Result = result
	}

	return
}

// DeletePool : Delete a load balancer pool
// Delete a load balancer pool.
func (dnsSvcs *DnsSvcsV1) DeletePool(deletePoolOptions *DeletePoolOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePoolOptions, "deletePoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePoolOptions, "deletePoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*deletePoolOptions.InstanceID, *deletePoolOptions.PoolID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeletePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deletePoolOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetPool : Get a load balancer pool
// Get details of a load balancer pool.
func (dnsSvcs *DnsSvcsV1) GetPool(getPoolOptions *GetPoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPoolOptions, "getPoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPoolOptions, "getPoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*getPoolOptions.InstanceID, *getPoolOptions.PoolID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetPool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getPoolOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalPool(m)
		response.Result = result
	}

	return
}

// UpdatePool : Update the properties of a load balancer pool
// Update the properties of a load balancer pool.
func (dnsSvcs *DnsSvcsV1) UpdatePool(updatePoolOptions *UpdatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePoolOptions, "updatePoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePoolOptions, "updatePoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*updatePoolOptions.InstanceID, *updatePoolOptions.PoolID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdatePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updatePoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updatePoolOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updatePoolOptions.Name != nil {
		body["name"] = updatePoolOptions.Name
	}
	if updatePoolOptions.Description != nil {
		body["description"] = updatePoolOptions.Description
	}
	if updatePoolOptions.Enabled != nil {
		body["enabled"] = updatePoolOptions.Enabled
	}
	if updatePoolOptions.MinimumOrigins != nil {
		body["minimum_origins"] = updatePoolOptions.MinimumOrigins
	}
	if updatePoolOptions.Origins != nil {
		body["origins"] = updatePoolOptions.Origins
	}
	if updatePoolOptions.Monitor != nil {
		body["monitor"] = updatePoolOptions.Monitor
	}
	if updatePoolOptions.NotificationType != nil {
		body["notification_type"] = updatePoolOptions.NotificationType
	}
	if updatePoolOptions.NotificationChannel != nil {
		body["notification_channel"] = updatePoolOptions.NotificationChannel
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalPool(m)
		response.Result = result
	}

	return
}

// ListMonitors : List load balancer monitors
// List the load balancer monitors.
func (dnsSvcs *DnsSvcsV1) ListMonitors(listMonitorsOptions *ListMonitorsOptions) (result *ListMonitors, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listMonitorsOptions, "listMonitorsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listMonitorsOptions, "listMonitorsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*listMonitorsOptions.InstanceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listMonitorsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListMonitors")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listMonitorsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listMonitorsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalListMonitors(m)
		response.Result = result
	}

	return
}

// CreateMonitor : Create a load balancer monitor
// Create a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) CreateMonitor(createMonitorOptions *CreateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createMonitorOptions, "createMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createMonitorOptions, "createMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*createMonitorOptions.InstanceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createMonitorOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createMonitorOptions.Description != nil {
		body["description"] = createMonitorOptions.Description
	}
	if createMonitorOptions.Type != nil {
		body["type"] = createMonitorOptions.Type
	}
	if createMonitorOptions.Port != nil {
		body["port"] = createMonitorOptions.Port
	}
	if createMonitorOptions.Interval != nil {
		body["interval"] = createMonitorOptions.Interval
	}
	if createMonitorOptions.Retries != nil {
		body["retries"] = createMonitorOptions.Retries
	}
	if createMonitorOptions.Timeout != nil {
		body["timeout"] = createMonitorOptions.Timeout
	}
	if createMonitorOptions.Method != nil {
		body["method"] = createMonitorOptions.Method
	}
	if createMonitorOptions.Path != nil {
		body["path"] = createMonitorOptions.Path
	}
	if createMonitorOptions.Header != nil {
		body["header"] = createMonitorOptions.Header
	}
	if createMonitorOptions.AllowInsecure != nil {
		body["allow_insecure"] = createMonitorOptions.AllowInsecure
	}
	if createMonitorOptions.ExpectedCodes != nil {
		body["expected_codes"] = createMonitorOptions.ExpectedCodes
	}
	if createMonitorOptions.ExpectedBody != nil {
		body["expected_body"] = createMonitorOptions.ExpectedBody
	}
	if createMonitorOptions.FollowRedirects != nil {
		body["follow_redirects"] = createMonitorOptions.FollowRedirects
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalMonitor(m)
		response.Result = result
	}

	return
}

// DeleteMonitor : Delete a load balancer monitor
// Delete a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) DeleteMonitor(deleteMonitorOptions *DeleteMonitorOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteMonitorOptions, "deleteMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteMonitorOptions, "deleteMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*deleteMonitorOptions.InstanceID, *deleteMonitorOptions.MonitorID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteMonitorOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetMonitor : Get a load balancer monitor
// Get details of a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) GetMonitor(getMonitorOptions *GetMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getMonitorOptions, "getMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getMonitorOptions, "getMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*getMonitorOptions.InstanceID, *getMonitorOptions.MonitorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getMonitorOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalMonitor(m)
		response.Result = result
	}

	return
}

// UpdateMonitor : Update the properties of a load balancer monitor
// Update the properties of a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateMonitorOptions, "updateMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateMonitorOptions, "updateMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*updateMonitorOptions.InstanceID, *updateMonitorOptions.MonitorID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateMonitorOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateMonitorOptions.Description != nil {
		body["description"] = updateMonitorOptions.Description
	}
	if updateMonitorOptions.Type != nil {
		body["type"] = updateMonitorOptions.Type
	}
	if updateMonitorOptions.Port != nil {
		body["port"] = updateMonitorOptions.Port
	}
	if updateMonitorOptions.Interval != nil {
		body["interval"] = updateMonitorOptions.Interval
	}
	if updateMonitorOptions.Retries != nil {
		body["retries"] = updateMonitorOptions.Retries
	}
	if updateMonitorOptions.Timeout != nil {
		body["timeout"] = updateMonitorOptions.Timeout
	}
	if updateMonitorOptions.Method != nil {
		body["method"] = updateMonitorOptions.Method
	}
	if updateMonitorOptions.Path != nil {
		body["path"] = updateMonitorOptions.Path
	}
	if updateMonitorOptions.Header != nil {
		body["header"] = updateMonitorOptions.Header
	}
	if updateMonitorOptions.AllowInsecure != nil {
		body["allow_insecure"] = updateMonitorOptions.AllowInsecure
	}
	if updateMonitorOptions.ExpectedCodes != nil {
		body["expected_codes"] = updateMonitorOptions.ExpectedCodes
	}
	if updateMonitorOptions.ExpectedBody != nil {
		body["expected_body"] = updateMonitorOptions.ExpectedBody
	}
	if updateMonitorOptions.FollowRedirects != nil {
		body["follow_redirects"] = updateMonitorOptions.FollowRedirects
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalMonitor(m)
		response.Result = result
	}

	return
}

// CreateLoadBalancerOptions : The CreateLoadBalancer options.
type CreateLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool ID's.
	AzPools *AzPools `json:"az_pools,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateLoadBalancerOptions : Instantiate CreateLoadBalancerOptions
func (*DnsSvcsV1) NewCreateLoadBalancerOptions(instanceID string, dnszoneID string) *CreateLoadBalancerOptions {
	return &CreateLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreateLoadBalancerOptions) SetInstanceID(instanceID string) *CreateLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *CreateLoadBalancerOptions) SetDnszoneID(dnszoneID string) *CreateLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateLoadBalancerOptions) SetName(name string) *CreateLoadBalancerOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateLoadBalancerOptions) SetDescription(description string) *CreateLoadBalancerOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *CreateLoadBalancerOptions) SetEnabled(enabled bool) *CreateLoadBalancerOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetTTL : Allow user to set TTL
func (options *CreateLoadBalancerOptions) SetTTL(ttl int64) *CreateLoadBalancerOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetFallbackPool : Allow user to set FallbackPool
func (options *CreateLoadBalancerOptions) SetFallbackPool(fallbackPool string) *CreateLoadBalancerOptions {
	options.FallbackPool = core.StringPtr(fallbackPool)
	return options
}

// SetDefaultPools : Allow user to set DefaultPools
func (options *CreateLoadBalancerOptions) SetDefaultPools(defaultPools []string) *CreateLoadBalancerOptions {
	options.DefaultPools = defaultPools
	return options
}

// SetAzPools : Allow user to set AzPools
func (options *CreateLoadBalancerOptions) SetAzPools(azPools *AzPools) *CreateLoadBalancerOptions {
	options.AzPools = azPools
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreateLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *CreateLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLoadBalancerOptions) SetHeaders(param map[string]string) *CreateLoadBalancerOptions {
	options.Headers = param
	return options
}

// CreateMonitorOptions : The CreateMonitor options.
type CreateMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	Header interface{} `json:"header,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS
	// monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// Follow redirects if returned by the origin. This parameter is only valid for HTTP and HTTPS monitors.
	FollowRedirects *bool `json:"follow_redirects,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateMonitorOptions.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	CreateMonitorOptions_Method_Get  = "GET"
	CreateMonitorOptions_Method_Head = "HEAD"
)

// NewCreateMonitorOptions : Instantiate CreateMonitorOptions
func (*DnsSvcsV1) NewCreateMonitorOptions(instanceID string) *CreateMonitorOptions {
	return &CreateMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreateMonitorOptions) SetInstanceID(instanceID string) *CreateMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateMonitorOptions) SetDescription(description string) *CreateMonitorOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetType : Allow user to set Type
func (options *CreateMonitorOptions) SetType(typeVar string) *CreateMonitorOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetPort : Allow user to set Port
func (options *CreateMonitorOptions) SetPort(port int64) *CreateMonitorOptions {
	options.Port = core.Int64Ptr(port)
	return options
}

// SetInterval : Allow user to set Interval
func (options *CreateMonitorOptions) SetInterval(interval int64) *CreateMonitorOptions {
	options.Interval = core.Int64Ptr(interval)
	return options
}

// SetRetries : Allow user to set Retries
func (options *CreateMonitorOptions) SetRetries(retries int64) *CreateMonitorOptions {
	options.Retries = core.Int64Ptr(retries)
	return options
}

// SetTimeout : Allow user to set Timeout
func (options *CreateMonitorOptions) SetTimeout(timeout int64) *CreateMonitorOptions {
	options.Timeout = core.Int64Ptr(timeout)
	return options
}

// SetMethod : Allow user to set Method
func (options *CreateMonitorOptions) SetMethod(method string) *CreateMonitorOptions {
	options.Method = core.StringPtr(method)
	return options
}

// SetPath : Allow user to set Path
func (options *CreateMonitorOptions) SetPath(path string) *CreateMonitorOptions {
	options.Path = core.StringPtr(path)
	return options
}

// SetHeader : Allow user to set Header
func (options *CreateMonitorOptions) SetHeader(header interface{}) *CreateMonitorOptions {
	options.Header = header
	return options
}

// SetAllowInsecure : Allow user to set AllowInsecure
func (options *CreateMonitorOptions) SetAllowInsecure(allowInsecure bool) *CreateMonitorOptions {
	options.AllowInsecure = core.BoolPtr(allowInsecure)
	return options
}

// SetExpectedCodes : Allow user to set ExpectedCodes
func (options *CreateMonitorOptions) SetExpectedCodes(expectedCodes string) *CreateMonitorOptions {
	options.ExpectedCodes = core.StringPtr(expectedCodes)
	return options
}

// SetExpectedBody : Allow user to set ExpectedBody
func (options *CreateMonitorOptions) SetExpectedBody(expectedBody string) *CreateMonitorOptions {
	options.ExpectedBody = core.StringPtr(expectedBody)
	return options
}

// SetFollowRedirects : Allow user to set FollowRedirects
func (options *CreateMonitorOptions) SetFollowRedirects(followRedirects bool) *CreateMonitorOptions {
	options.FollowRedirects = core.BoolPtr(followRedirects)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreateMonitorOptions) SetXCorrelationID(xCorrelationID string) *CreateMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateMonitorOptions) SetHeaders(param map[string]string) *CreateMonitorOptions {
	options.Headers = param
	return options
}

// CreatePoolOptions : The CreatePool options.
type CreatePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	MinimumOrigins *int64 `json:"minimum_origins,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []Origin `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The type of the notification channel.
	NotificationType *string `json:"notification_type,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePoolOptions.NotificationType property.
// The type of the notification channel.
const (
	CreatePoolOptions_NotificationType_Email   = "email"
	CreatePoolOptions_NotificationType_Webhook = "webhook"
)

// NewCreatePoolOptions : Instantiate CreatePoolOptions
func (*DnsSvcsV1) NewCreatePoolOptions(instanceID string) *CreatePoolOptions {
	return &CreatePoolOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreatePoolOptions) SetInstanceID(instanceID string) *CreatePoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetName : Allow user to set Name
func (options *CreatePoolOptions) SetName(name string) *CreatePoolOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreatePoolOptions) SetDescription(description string) *CreatePoolOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *CreatePoolOptions) SetEnabled(enabled bool) *CreatePoolOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetMinimumOrigins : Allow user to set MinimumOrigins
func (options *CreatePoolOptions) SetMinimumOrigins(minimumOrigins int64) *CreatePoolOptions {
	options.MinimumOrigins = core.Int64Ptr(minimumOrigins)
	return options
}

// SetOrigins : Allow user to set Origins
func (options *CreatePoolOptions) SetOrigins(origins []Origin) *CreatePoolOptions {
	options.Origins = origins
	return options
}

// SetMonitor : Allow user to set Monitor
func (options *CreatePoolOptions) SetMonitor(monitor string) *CreatePoolOptions {
	options.Monitor = core.StringPtr(monitor)
	return options
}

// SetNotificationType : Allow user to set NotificationType
func (options *CreatePoolOptions) SetNotificationType(notificationType string) *CreatePoolOptions {
	options.NotificationType = core.StringPtr(notificationType)
	return options
}

// SetNotificationChannel : Allow user to set NotificationChannel
func (options *CreatePoolOptions) SetNotificationChannel(notificationChannel string) *CreatePoolOptions {
	options.NotificationChannel = core.StringPtr(notificationChannel)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreatePoolOptions) SetXCorrelationID(xCorrelationID string) *CreatePoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePoolOptions) SetHeaders(param map[string]string) *CreatePoolOptions {
	options.Headers = param
	return options
}

// DeleteLoadBalancerOptions : The DeleteLoadBalancer options.
type DeleteLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a load balancer.
	LbID *string `json:"lb_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLoadBalancerOptions : Instantiate DeleteLoadBalancerOptions
func (*DnsSvcsV1) NewDeleteLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *DeleteLoadBalancerOptions {
	return &DeleteLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		LbID:       core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteLoadBalancerOptions) SetInstanceID(instanceID string) *DeleteLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *DeleteLoadBalancerOptions) SetDnszoneID(dnszoneID string) *DeleteLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetLbID : Allow user to set LbID
func (options *DeleteLoadBalancerOptions) SetLbID(lbID string) *DeleteLoadBalancerOptions {
	options.LbID = core.StringPtr(lbID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeleteLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *DeleteLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLoadBalancerOptions) SetHeaders(param map[string]string) *DeleteLoadBalancerOptions {
	options.Headers = param
	return options
}

// DeleteMonitorOptions : The DeleteMonitor options.
type DeleteMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"monitor_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteMonitorOptions : Instantiate DeleteMonitorOptions
func (*DnsSvcsV1) NewDeleteMonitorOptions(instanceID string, monitorID string) *DeleteMonitorOptions {
	return &DeleteMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID:  core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteMonitorOptions) SetInstanceID(instanceID string) *DeleteMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetMonitorID : Allow user to set MonitorID
func (options *DeleteMonitorOptions) SetMonitorID(monitorID string) *DeleteMonitorOptions {
	options.MonitorID = core.StringPtr(monitorID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeleteMonitorOptions) SetXCorrelationID(xCorrelationID string) *DeleteMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteMonitorOptions) SetHeaders(param map[string]string) *DeleteMonitorOptions {
	options.Headers = param
	return options
}

// DeletePoolOptions : The DeletePool options.
type DeletePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"pool_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePoolOptions : Instantiate DeletePoolOptions
func (*DnsSvcsV1) NewDeletePoolOptions(instanceID string, poolID string) *DeletePoolOptions {
	return &DeletePoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID:     core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeletePoolOptions) SetInstanceID(instanceID string) *DeletePoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPoolID : Allow user to set PoolID
func (options *DeletePoolOptions) SetPoolID(poolID string) *DeletePoolOptions {
	options.PoolID = core.StringPtr(poolID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeletePoolOptions) SetXCorrelationID(xCorrelationID string) *DeletePoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePoolOptions) SetHeaders(param map[string]string) *DeletePoolOptions {
	options.Headers = param
	return options
}

// GetLoadBalancerOptions : The GetLoadBalancer options.
type GetLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a load balancer.
	LbID *string `json:"lb_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLoadBalancerOptions : Instantiate GetLoadBalancerOptions
func (*DnsSvcsV1) NewGetLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *GetLoadBalancerOptions {
	return &GetLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		LbID:       core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetLoadBalancerOptions) SetInstanceID(instanceID string) *GetLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *GetLoadBalancerOptions) SetDnszoneID(dnszoneID string) *GetLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetLbID : Allow user to set LbID
func (options *GetLoadBalancerOptions) SetLbID(lbID string) *GetLoadBalancerOptions {
	options.LbID = core.StringPtr(lbID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *GetLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetLoadBalancerOptions) SetHeaders(param map[string]string) *GetLoadBalancerOptions {
	options.Headers = param
	return options
}

// GetMonitorOptions : The GetMonitor options.
type GetMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"monitor_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMonitorOptions : Instantiate GetMonitorOptions
func (*DnsSvcsV1) NewGetMonitorOptions(instanceID string, monitorID string) *GetMonitorOptions {
	return &GetMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID:  core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetMonitorOptions) SetInstanceID(instanceID string) *GetMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetMonitorID : Allow user to set MonitorID
func (options *GetMonitorOptions) SetMonitorID(monitorID string) *GetMonitorOptions {
	options.MonitorID = core.StringPtr(monitorID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetMonitorOptions) SetXCorrelationID(xCorrelationID string) *GetMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMonitorOptions) SetHeaders(param map[string]string) *GetMonitorOptions {
	options.Headers = param
	return options
}

// GetPoolOptions : The GetPool options.
type GetPoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"pool_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPoolOptions : Instantiate GetPoolOptions
func (*DnsSvcsV1) NewGetPoolOptions(instanceID string, poolID string) *GetPoolOptions {
	return &GetPoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID:     core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetPoolOptions) SetInstanceID(instanceID string) *GetPoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPoolID : Allow user to set PoolID
func (options *GetPoolOptions) SetPoolID(poolID string) *GetPoolOptions {
	options.PoolID = core.StringPtr(poolID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetPoolOptions) SetXCorrelationID(xCorrelationID string) *GetPoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPoolOptions) SetHeaders(param map[string]string) *GetPoolOptions {
	options.Headers = param
	return options
}

// ListLoadBalancersOptions : The ListLoadBalancers options.
type ListLoadBalancersOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLoadBalancersOptions : Instantiate ListLoadBalancersOptions
func (*DnsSvcsV1) NewListLoadBalancersOptions(instanceID string, dnszoneID string) *ListLoadBalancersOptions {
	return &ListLoadBalancersOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListLoadBalancersOptions) SetInstanceID(instanceID string) *ListLoadBalancersOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *ListLoadBalancersOptions) SetDnszoneID(dnszoneID string) *ListLoadBalancersOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListLoadBalancersOptions) SetXCorrelationID(xCorrelationID string) *ListLoadBalancersOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLoadBalancersOptions) SetHeaders(param map[string]string) *ListLoadBalancersOptions {
	options.Headers = param
	return options
}

// ListMonitorsOptions : The ListMonitors options.
type ListMonitorsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListMonitorsOptions : Instantiate ListMonitorsOptions
func (*DnsSvcsV1) NewListMonitorsOptions(instanceID string) *ListMonitorsOptions {
	return &ListMonitorsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListMonitorsOptions) SetInstanceID(instanceID string) *ListMonitorsOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListMonitorsOptions) SetXCorrelationID(xCorrelationID string) *ListMonitorsOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListMonitorsOptions) SetHeaders(param map[string]string) *ListMonitorsOptions {
	options.Headers = param
	return options
}

// ListPoolsOptions : The ListPools options.
type ListPoolsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPoolsOptions : Instantiate ListPoolsOptions
func (*DnsSvcsV1) NewListPoolsOptions(instanceID string) *ListPoolsOptions {
	return &ListPoolsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListPoolsOptions) SetInstanceID(instanceID string) *ListPoolsOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListPoolsOptions) SetXCorrelationID(xCorrelationID string) *ListPoolsOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListPoolsOptions) SetHeaders(param map[string]string) *ListPoolsOptions {
	options.Headers = param
	return options
}

// UpdateLoadBalancerOptions : The UpdateLoadBalancer options.
type UpdateLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a load balancer.
	LbID *string `json:"lb_id" validate:"required"`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool ID's.
	AzPools *AzPools `json:"az_pools,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateLoadBalancerOptions : Instantiate UpdateLoadBalancerOptions
func (*DnsSvcsV1) NewUpdateLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *UpdateLoadBalancerOptions {
	return &UpdateLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		LbID:       core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateLoadBalancerOptions) SetInstanceID(instanceID string) *UpdateLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *UpdateLoadBalancerOptions) SetDnszoneID(dnszoneID string) *UpdateLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetLbID : Allow user to set LbID
func (options *UpdateLoadBalancerOptions) SetLbID(lbID string) *UpdateLoadBalancerOptions {
	options.LbID = core.StringPtr(lbID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateLoadBalancerOptions) SetName(name string) *UpdateLoadBalancerOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateLoadBalancerOptions) SetDescription(description string) *UpdateLoadBalancerOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *UpdateLoadBalancerOptions) SetEnabled(enabled bool) *UpdateLoadBalancerOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetTTL : Allow user to set TTL
func (options *UpdateLoadBalancerOptions) SetTTL(ttl int64) *UpdateLoadBalancerOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetFallbackPool : Allow user to set FallbackPool
func (options *UpdateLoadBalancerOptions) SetFallbackPool(fallbackPool string) *UpdateLoadBalancerOptions {
	options.FallbackPool = core.StringPtr(fallbackPool)
	return options
}

// SetDefaultPools : Allow user to set DefaultPools
func (options *UpdateLoadBalancerOptions) SetDefaultPools(defaultPools []string) *UpdateLoadBalancerOptions {
	options.DefaultPools = defaultPools
	return options
}

// SetAzPools : Allow user to set AzPools
func (options *UpdateLoadBalancerOptions) SetAzPools(azPools *AzPools) *UpdateLoadBalancerOptions {
	options.AzPools = azPools
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdateLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *UpdateLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateLoadBalancerOptions) SetHeaders(param map[string]string) *UpdateLoadBalancerOptions {
	options.Headers = param
	return options
}

// UpdateMonitorOptions : The UpdateMonitor options.
type UpdateMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"monitor_id" validate:"required"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	Header interface{} `json:"header,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS
	// monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// Follow redirects if returned by the origin. This parameter is only valid for HTTP and HTTPS monitors.
	FollowRedirects *bool `json:"follow_redirects,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateMonitorOptions.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	UpdateMonitorOptions_Method_Get  = "GET"
	UpdateMonitorOptions_Method_Head = "HEAD"
)

// NewUpdateMonitorOptions : Instantiate UpdateMonitorOptions
func (*DnsSvcsV1) NewUpdateMonitorOptions(instanceID string, monitorID string) *UpdateMonitorOptions {
	return &UpdateMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID:  core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateMonitorOptions) SetInstanceID(instanceID string) *UpdateMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetMonitorID : Allow user to set MonitorID
func (options *UpdateMonitorOptions) SetMonitorID(monitorID string) *UpdateMonitorOptions {
	options.MonitorID = core.StringPtr(monitorID)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateMonitorOptions) SetDescription(description string) *UpdateMonitorOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetType : Allow user to set Type
func (options *UpdateMonitorOptions) SetType(typeVar string) *UpdateMonitorOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetPort : Allow user to set Port
func (options *UpdateMonitorOptions) SetPort(port int64) *UpdateMonitorOptions {
	options.Port = core.Int64Ptr(port)
	return options
}

// SetInterval : Allow user to set Interval
func (options *UpdateMonitorOptions) SetInterval(interval int64) *UpdateMonitorOptions {
	options.Interval = core.Int64Ptr(interval)
	return options
}

// SetRetries : Allow user to set Retries
func (options *UpdateMonitorOptions) SetRetries(retries int64) *UpdateMonitorOptions {
	options.Retries = core.Int64Ptr(retries)
	return options
}

// SetTimeout : Allow user to set Timeout
func (options *UpdateMonitorOptions) SetTimeout(timeout int64) *UpdateMonitorOptions {
	options.Timeout = core.Int64Ptr(timeout)
	return options
}

// SetMethod : Allow user to set Method
func (options *UpdateMonitorOptions) SetMethod(method string) *UpdateMonitorOptions {
	options.Method = core.StringPtr(method)
	return options
}

// SetPath : Allow user to set Path
func (options *UpdateMonitorOptions) SetPath(path string) *UpdateMonitorOptions {
	options.Path = core.StringPtr(path)
	return options
}

// SetHeader : Allow user to set Header
func (options *UpdateMonitorOptions) SetHeader(header interface{}) *UpdateMonitorOptions {
	options.Header = header
	return options
}

// SetAllowInsecure : Allow user to set AllowInsecure
func (options *UpdateMonitorOptions) SetAllowInsecure(allowInsecure bool) *UpdateMonitorOptions {
	options.AllowInsecure = core.BoolPtr(allowInsecure)
	return options
}

// SetExpectedCodes : Allow user to set ExpectedCodes
func (options *UpdateMonitorOptions) SetExpectedCodes(expectedCodes string) *UpdateMonitorOptions {
	options.ExpectedCodes = core.StringPtr(expectedCodes)
	return options
}

// SetExpectedBody : Allow user to set ExpectedBody
func (options *UpdateMonitorOptions) SetExpectedBody(expectedBody string) *UpdateMonitorOptions {
	options.ExpectedBody = core.StringPtr(expectedBody)
	return options
}

// SetFollowRedirects : Allow user to set FollowRedirects
func (options *UpdateMonitorOptions) SetFollowRedirects(followRedirects bool) *UpdateMonitorOptions {
	options.FollowRedirects = core.BoolPtr(followRedirects)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdateMonitorOptions) SetXCorrelationID(xCorrelationID string) *UpdateMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMonitorOptions) SetHeaders(param map[string]string) *UpdateMonitorOptions {
	options.Headers = param
	return options
}

// UpdatePoolOptions : The UpdatePool options.
type UpdatePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"pool_id" validate:"required"`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	MinimumOrigins *int64 `json:"minimum_origins,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []Origin `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The type of the notification channel.
	NotificationType *string `json:"notification_type,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdatePoolOptions.NotificationType property.
// The type of the notification channel.
const (
	UpdatePoolOptions_NotificationType_Email   = "email"
	UpdatePoolOptions_NotificationType_Webhook = "webhook"
)

// NewUpdatePoolOptions : Instantiate UpdatePoolOptions
func (*DnsSvcsV1) NewUpdatePoolOptions(instanceID string, poolID string) *UpdatePoolOptions {
	return &UpdatePoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID:     core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdatePoolOptions) SetInstanceID(instanceID string) *UpdatePoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPoolID : Allow user to set PoolID
func (options *UpdatePoolOptions) SetPoolID(poolID string) *UpdatePoolOptions {
	options.PoolID = core.StringPtr(poolID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdatePoolOptions) SetName(name string) *UpdatePoolOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdatePoolOptions) SetDescription(description string) *UpdatePoolOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *UpdatePoolOptions) SetEnabled(enabled bool) *UpdatePoolOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetMinimumOrigins : Allow user to set MinimumOrigins
func (options *UpdatePoolOptions) SetMinimumOrigins(minimumOrigins int64) *UpdatePoolOptions {
	options.MinimumOrigins = core.Int64Ptr(minimumOrigins)
	return options
}

// SetOrigins : Allow user to set Origins
func (options *UpdatePoolOptions) SetOrigins(origins []Origin) *UpdatePoolOptions {
	options.Origins = origins
	return options
}

// SetMonitor : Allow user to set Monitor
func (options *UpdatePoolOptions) SetMonitor(monitor string) *UpdatePoolOptions {
	options.Monitor = core.StringPtr(monitor)
	return options
}

// SetNotificationType : Allow user to set NotificationType
func (options *UpdatePoolOptions) SetNotificationType(notificationType string) *UpdatePoolOptions {
	options.NotificationType = core.StringPtr(notificationType)
	return options
}

// SetNotificationChannel : Allow user to set NotificationChannel
func (options *UpdatePoolOptions) SetNotificationChannel(notificationChannel string) *UpdatePoolOptions {
	options.NotificationChannel = core.StringPtr(notificationChannel)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdatePoolOptions) SetXCorrelationID(xCorrelationID string) *UpdatePoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePoolOptions) SetHeaders(param map[string]string) *UpdatePoolOptions {
	options.Headers = param
	return options
}

// AzPools : Map availability zones to pool ID's.
type AzPools struct {
	// us-south-1.
	UsSouth1 []string `json:"us-south-1,omitempty"`

	// us-south-2.
	UsSouth2 []string `json:"us-south-2,omitempty"`

	// us-south-3.
	UsSouth3 []string `json:"us-south-3,omitempty"`

	// us-east-1.
	UsEast1 []string `json:"us-east-1,omitempty"`

	// us-east-2.
	UsEast2 []string `json:"us-east-2,omitempty"`

	// us-east-3.
	UsEast3 []string `json:"us-east-3,omitempty"`

	// eu-gb-1.
	EuGb1 []string `json:"eu-gb-1,omitempty"`

	// eu-gb-2.
	EuGb2 []string `json:"eu-gb-2,omitempty"`

	// eu-gb-3.
	EuGb3 []string `json:"eu-gb-3,omitempty"`

	// eu-de-1.
	EuDe1 []string `json:"eu-de-1,omitempty"`

	// eu-de-2.
	EuDe2 []string `json:"eu-de-2,omitempty"`

	// eu-de-3.
	EuDe3 []string `json:"eu-de-3,omitempty"`

	// au-syd-1.
	AuSyd1 []string `json:"au-syd-1,omitempty"`

	// au-syd-2.
	AuSyd2 []string `json:"au-syd-2,omitempty"`

	// au-syd-3.
	AuSyd3 []string `json:"au-syd-3,omitempty"`

	// jp-tok-1.
	JpTok1 []string `json:"jp-tok-1,omitempty"`

	// jp-tok-2.
	JpTok2 []string `json:"jp-tok-2,omitempty"`

	// jp-tok-3.
	JpTok3 []string `json:"jp-tok-3,omitempty"`
}

// UnmarshalAzPools constructs an instance of AzPools from the specified map.
func UnmarshalAzPools(m map[string]interface{}) (result *AzPools, err error) {
	obj := new(AzPools)
	obj.UsSouth1, err = core.UnmarshalStringSlice(m, "us-south-1")
	if err != nil {
		return
	}
	obj.UsSouth2, err = core.UnmarshalStringSlice(m, "us-south-2")
	if err != nil {
		return
	}
	obj.UsSouth3, err = core.UnmarshalStringSlice(m, "us-south-3")
	if err != nil {
		return
	}
	obj.UsEast1, err = core.UnmarshalStringSlice(m, "us-east-1")
	if err != nil {
		return
	}
	obj.UsEast2, err = core.UnmarshalStringSlice(m, "us-east-2")
	if err != nil {
		return
	}
	obj.UsEast3, err = core.UnmarshalStringSlice(m, "us-east-3")
	if err != nil {
		return
	}
	obj.EuGb1, err = core.UnmarshalStringSlice(m, "eu-gb-1")
	if err != nil {
		return
	}
	obj.EuGb2, err = core.UnmarshalStringSlice(m, "eu-gb-2")
	if err != nil {
		return
	}
	obj.EuGb3, err = core.UnmarshalStringSlice(m, "eu-gb-3")
	if err != nil {
		return
	}
	obj.EuDe1, err = core.UnmarshalStringSlice(m, "eu-de-1")
	if err != nil {
		return
	}
	obj.EuDe2, err = core.UnmarshalStringSlice(m, "eu-de-2")
	if err != nil {
		return
	}
	obj.EuDe3, err = core.UnmarshalStringSlice(m, "eu-de-3")
	if err != nil {
		return
	}
	obj.AuSyd1, err = core.UnmarshalStringSlice(m, "au-syd-1")
	if err != nil {
		return
	}
	obj.AuSyd2, err = core.UnmarshalStringSlice(m, "au-syd-2")
	if err != nil {
		return
	}
	obj.AuSyd3, err = core.UnmarshalStringSlice(m, "au-syd-3")
	if err != nil {
		return
	}
	obj.JpTok1, err = core.UnmarshalStringSlice(m, "jp-tok-1")
	if err != nil {
		return
	}
	obj.JpTok2, err = core.UnmarshalStringSlice(m, "jp-tok-2")
	if err != nil {
		return
	}
	obj.JpTok3, err = core.UnmarshalStringSlice(m, "jp-tok-3")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAzPoolsAsProperty unmarshals an instance of AzPools that is stored as a property
// within the specified map.
func UnmarshalAzPoolsAsProperty(m map[string]interface{}, propertyName string) (result *AzPools, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AzPools'", propertyName)
			return
		}
		result, err = UnmarshalAzPools(objMap)
	}
	return
}

// ListLoadBalancers : List Global Load Balancers response.
type ListLoadBalancers struct {
	// An array of Global Load Balancers.
	LoadBalancers []LoadBalancer `json:"load_balancers" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of Global Load Balancers per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of Global Load Balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of Global Load Balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}

// UnmarshalListLoadBalancers constructs an instance of ListLoadBalancers from the specified map.
func UnmarshalListLoadBalancers(m map[string]interface{}) (result *ListLoadBalancers, err error) {
	obj := new(ListLoadBalancers)
	obj.LoadBalancers, err = UnmarshalLoadBalancerSliceAsProperty(m, "load_balancers")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Count, err = core.UnmarshalInt64(m, "count")
	if err != nil {
		return
	}
	obj.TotalCount, err = core.UnmarshalInt64(m, "total_count")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalFirstHrefAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalNextHrefAsProperty(m, "next")
	if err != nil {
		return
	}
	result = obj
	return
}

// ListMonitors : List load balancer monitors response.
type ListMonitors struct {
	// An array of load balancer monitors.
	Monitors []Monitor `json:"monitors" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of load balancer monitors per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of load balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of load balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}

// UnmarshalListMonitors constructs an instance of ListMonitors from the specified map.
func UnmarshalListMonitors(m map[string]interface{}) (result *ListMonitors, err error) {
	obj := new(ListMonitors)
	obj.Monitors, err = UnmarshalMonitorSliceAsProperty(m, "monitors")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Count, err = core.UnmarshalInt64(m, "count")
	if err != nil {
		return
	}
	obj.TotalCount, err = core.UnmarshalInt64(m, "total_count")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalFirstHrefAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalNextHrefAsProperty(m, "next")
	if err != nil {
		return
	}
	result = obj
	return
}

// ListPools : List load balancer pools response.
type ListPools struct {
	// An array of load balancer pools.
	Pools []Pool `json:"pools" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of load balancer pools per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of load balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of load balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}

// UnmarshalListPools constructs an instance of ListPools from the specified map.
func UnmarshalListPools(m map[string]interface{}) (result *ListPools, err error) {
	obj := new(ListPools)
	obj.Pools, err = UnmarshalPoolSliceAsProperty(m, "pools")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Count, err = core.UnmarshalInt64(m, "count")
	if err != nil {
		return
	}
	obj.TotalCount, err = core.UnmarshalInt64(m, "total_count")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalFirstHrefAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalNextHrefAsProperty(m, "next")
	if err != nil {
		return
	}
	result = obj
	return
}

// LoadBalancer : Load balancer details.
type LoadBalancer struct {
	// Identifier of the load balancer.
	ID *string `json:"id,omitempty"`

	// the time when a load balancer is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Healthy state of the load balancer.
	Health *bool `json:"health,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool ID's.
	AzPools *AzPools `json:"az_pools,omitempty"`
}

// UnmarshalLoadBalancer constructs an instance of LoadBalancer from the specified map.
func UnmarshalLoadBalancer(m map[string]interface{}) (result *LoadBalancer, err error) {
	obj := new(LoadBalancer)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.CreatedOn, err = core.UnmarshalString(m, "created_on")
	if err != nil {
		return
	}
	obj.ModifiedOn, err = core.UnmarshalString(m, "modified_on")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Enabled, err = core.UnmarshalBool(m, "enabled")
	if err != nil {
		return
	}
	obj.TTL, err = core.UnmarshalInt64(m, "ttl")
	if err != nil {
		return
	}
	obj.Health, err = core.UnmarshalBool(m, "health")
	if err != nil {
		return
	}
	obj.FallbackPool, err = core.UnmarshalString(m, "fallback_pool")
	if err != nil {
		return
	}
	obj.DefaultPools, err = core.UnmarshalStringSlice(m, "default_pools")
	if err != nil {
		return
	}
	obj.AzPools, err = UnmarshalAzPoolsAsProperty(m, "az_pools")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalLoadBalancerSlice unmarshals a slice of LoadBalancer instances from the specified list of maps.
func UnmarshalLoadBalancerSlice(s []interface{}) (slice []LoadBalancer, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'LoadBalancer'")
			return
		}
		obj, e := UnmarshalLoadBalancer(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalLoadBalancerSliceAsProperty unmarshals a slice of LoadBalancer instances that are stored as a property
// within the specified map.
func UnmarshalLoadBalancerSliceAsProperty(m map[string]interface{}, propertyName string) (slice []LoadBalancer, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'LoadBalancer'", propertyName)
			return
		}
		slice, err = UnmarshalLoadBalancerSlice(vSlice)
	}
	return
}

// Monitor : Load balancer monitor details.
type Monitor struct {
	// Identifier of the load balancer monitor.
	ID *string `json:"id,omitempty"`

	// the time when a load balancer monitor is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer monitor is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	Header interface{} `json:"header,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS
	// monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// Follow redirects if returned by the origin. This parameter is only valid for HTTP and HTTPS monitors.
	FollowRedirects *bool `json:"follow_redirects,omitempty"`
}

// Constants associated with the Monitor.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	Monitor_Method_Get  = "GET"
	Monitor_Method_Head = "HEAD"
)

// UnmarshalMonitor constructs an instance of Monitor from the specified map.
func UnmarshalMonitor(m map[string]interface{}) (result *Monitor, err error) {
	obj := new(Monitor)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.CreatedOn, err = core.UnmarshalString(m, "created_on")
	if err != nil {
		return
	}
	obj.ModifiedOn, err = core.UnmarshalString(m, "modified_on")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Port, err = core.UnmarshalInt64(m, "port")
	if err != nil {
		return
	}
	obj.Interval, err = core.UnmarshalInt64(m, "interval")
	if err != nil {
		return
	}
	obj.Retries, err = core.UnmarshalInt64(m, "retries")
	if err != nil {
		return
	}
	obj.Timeout, err = core.UnmarshalInt64(m, "timeout")
	if err != nil {
		return
	}
	obj.Method, err = core.UnmarshalString(m, "method")
	if err != nil {
		return
	}
	obj.Path, err = core.UnmarshalString(m, "path")
	if err != nil {
		return
	}
	obj.Header, err = core.UnmarshalAny(m, "header")
	if err != nil {
		return
	}
	obj.AllowInsecure, err = core.UnmarshalBool(m, "allow_insecure")
	if err != nil {
		return
	}
	obj.ExpectedCodes, err = core.UnmarshalString(m, "expected_codes")
	if err != nil {
		return
	}
	obj.ExpectedBody, err = core.UnmarshalString(m, "expected_body")
	if err != nil {
		return
	}
	obj.FollowRedirects, err = core.UnmarshalBool(m, "follow_redirects")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalMonitorSlice unmarshals a slice of Monitor instances from the specified list of maps.
func UnmarshalMonitorSlice(s []interface{}) (slice []Monitor, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Monitor'")
			return
		}
		obj, e := UnmarshalMonitor(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalMonitorSliceAsProperty unmarshals a slice of Monitor instances that are stored as a property
// within the specified map.
func UnmarshalMonitorSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Monitor, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Monitor'", propertyName)
			return
		}
		slice, err = UnmarshalMonitorSlice(vSlice)
	}
	return
}

// Origin : Origin server.
type Origin struct {
	// The name of the origin server.
	Name *string `json:"name,omitempty"`

	// Description of the origin server.
	Description *string `json:"description,omitempty"`

	// The address of the origin server. It can be a hostname or an IP address.
	Address *string `json:"address,omitempty"`

	// Whether the origin server is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Weight for traffic distribution.
	Weight *int64 `json:"weight,omitempty"`
}

// UnmarshalOrigin constructs an instance of Origin from the specified map.
func UnmarshalOrigin(m map[string]interface{}) (result *Origin, err error) {
	obj := new(Origin)
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.Address, err = core.UnmarshalString(m, "address")
	if err != nil {
		return
	}
	obj.Enabled, err = core.UnmarshalBool(m, "enabled")
	if err != nil {
		return
	}
	obj.Weight, err = core.UnmarshalInt64(m, "weight")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalOriginSlice unmarshals a slice of Origin instances from the specified list of maps.
func UnmarshalOriginSlice(s []interface{}) (slice []Origin, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Origin'")
			return
		}
		obj, e := UnmarshalOrigin(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalOriginSliceAsProperty unmarshals a slice of Origin instances that are stored as a property
// within the specified map.
func UnmarshalOriginSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Origin, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Origin'", propertyName)
			return
		}
		slice, err = UnmarshalOriginSlice(vSlice)
	}
	return
}

// Pool : Load balancer pool details.
type Pool struct {
	// Identifier of the load balancer pool.
	ID *string `json:"id,omitempty"`

	// the time when a load balancer pool is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer pool is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	MinimumOrigins *int64 `json:"minimum_origins,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []Origin `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The type of the notification channel.
	NotificationType *string `json:"notification_type,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`
}

// Constants associated with the Pool.NotificationType property.
// The type of the notification channel.
const (
	Pool_NotificationType_Email   = "email"
	Pool_NotificationType_Webhook = "webhook"
)

// UnmarshalPool constructs an instance of Pool from the specified map.
func UnmarshalPool(m map[string]interface{}) (result *Pool, err error) {
	obj := new(Pool)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.CreatedOn, err = core.UnmarshalString(m, "created_on")
	if err != nil {
		return
	}
	obj.ModifiedOn, err = core.UnmarshalString(m, "modified_on")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.Enabled, err = core.UnmarshalBool(m, "enabled")
	if err != nil {
		return
	}
	obj.MinimumOrigins, err = core.UnmarshalInt64(m, "minimum_origins")
	if err != nil {
		return
	}
	obj.Origins, err = UnmarshalOriginSliceAsProperty(m, "origins")
	if err != nil {
		return
	}
	obj.Monitor, err = core.UnmarshalString(m, "monitor")
	if err != nil {
		return
	}
	obj.NotificationType, err = core.UnmarshalString(m, "notification_type")
	if err != nil {
		return
	}
	obj.NotificationChannel, err = core.UnmarshalString(m, "notification_channel")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPoolSlice unmarshals a slice of Pool instances from the specified list of maps.
func UnmarshalPoolSlice(s []interface{}) (slice []Pool, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Pool'")
			return
		}
		obj, e := UnmarshalPool(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPoolSliceAsProperty unmarshals a slice of Pool instances that are stored as a property
// within the specified map.
func UnmarshalPoolSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Pool, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Pool'", propertyName)
			return
		}
		slice, err = UnmarshalPoolSlice(vSlice)
	}
	return
}
