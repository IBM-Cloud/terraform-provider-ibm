---
subcategory: "API Gateway"
layout: "ibm"
page_title: "IBM: API Gateway"
description: |-
  Reads Endpoint and its subscriptions.
---

# `ibm_api_gateway`

Retrieve information about an existing API Gateway instance. For more information, about API Gateway, see [Getting started with API Gateway](https://cloud.ibm.com/docs/api-gateway?topic=api-gateway-getting-started).

**Note**

Configuration of an API Gateway data source requires the region parameter to be set for the IBM provider in the `provider.tf.`  If not specified, endpoint will default to `us-south`.

## Example usage

```
data "ibm_api_gateway" "apigateway" {
    service_instance_crn = ibm_resource_instance.apigateway.id
    
}
```

## Argument reference
Review the input parameters that you can specify for your resource. 

- `service_instance_crn` - (Required, String) The CRN of the API Gateway service instance.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `endpoints`- (List) A list of API Gateway endpoints that are associated with the service instance.
	- `alias_url` - (String) The alias URL of an endpoint.
	- `base_path` - (String) The base path of the endpoint.
	- `endpoint_id` - (String) The ID of the endpoint.
	- `managed`- (Boolean) If set to **true**, the endpoint is online. If set to **false**, the endpoint is offline.
	- `managed_url` - (String) The managed URL of an endpoint.
	- `name` - (String) The name of the endpoint.
	- `open_api_doc` - (String) The Open API document of the endpoint.
	- `provider_id` - (String) The provider ID of the endpoint.
	- `routes` - (Strings) Invokable routes for an endpoint.
	- `shared` - (String) The shared status of the endpoint.
    	- `subscriptions`- (List of endpoint subscriptions) A list of subscriptions that you created for your endpoint.
		- `client_id` - (String) The client ID of a subscription.
	   	- `name` - (String) The name of the subscription.
		- `secret_provided`- (Boolean) If set to **true**, the client secret is provided. If set to **false**, the client secret is not provided.
		- `type` - (String) The type of subscription. Supported values are `bluemix` and `external`.
