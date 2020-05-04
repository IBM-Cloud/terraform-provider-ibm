---
layout: "ibm"
page_title: "IBM: API Gateway"
sidebar_current: "docs-ibm-resource-api-gateway"
description: |-
  Reads Endpoint and its subscriptions.
---

# ibm\_api_gateway

Import the details of an existing IBM Cloud API Gateway instance as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.Configuration of an api gateway data source requires the region parameter to be set for the IBM provider in the provider.tf.  If not specified, endpoint will default to us-south.

## Example Usage for a single Api doc as input

```hcl

data "ibm_api_gateway" "apigateway" {
    service_instance_crn = ibm_resource_instance.apigateway.id
    
}

```

## Argument Reference

The following arguments are supported:

* `service_instance_crn` - (Required,string) The CRN-based API Gateway service instance ID.


## Attribute Reference

The following attributes are exported:

* `endpoints` - List of all endpoints of the API Gateway Instance. Each endpoint block in a list has following structure.
  * `endpoint_id` - The Id of the Endpoint
  * `name` - The display name for the Endpoint.
  * `managed` - Managed indicates if endpoint is online or offline.
  * `shared` -  The Shared status of an endpoint
  * `base_path` - Base path of an endpoint
  * `routes` -Invokable routes for an endpoint
  * `provider_id` - Provider ID of an endpoint.
  * `managed_url` -Managed URL for an endpoint
  * `alias_url` - Alias URL of an endpoint.
  * `open_api_doc` - APi Document of the endpoint
  * `subscriptions` - List of all subscriptions of an endpoint. Each subscription block in a list has following structure.
    * `client_id` -ClientID or ID of a subscription
    * `name` - Name of Subscription
    * `type` - Type of subscription. [internal],[external]
    * `secret_provided` - It denotes if client secret is provided to Subscription or not.