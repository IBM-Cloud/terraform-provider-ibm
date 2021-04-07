---
subcategory: "API Gateway"
layout: "ibm"
page_title: "IBM: API Gateway Endpoint Subscription"
description: |-
  Create, Update,Manage and Delete Subscription of Endpoints of a API Gateway Service Instance.
---

# ibm\_api_gateway_endpoint_subscription

Provides Subscription for an Endpoint. This allows Subscription to be created, updated managed and deleted.Configuration of an api gateway data resource requires the region parameter to be set for the IBM provider in the provider.tf.  If not specified, endpoint to which subscription is being created will default to us-south. 

## Example Usage

```hcl
data "ibm_api_gateway" "endpoint"{
    service_instance_crn =ibm_resource_instance.apigateway.id
}

resource "ibm_api_gateway_endpoint_subscription" "subs" {
  artifact_id = data.ibm_api_gateway.endpoint.endpoints[0].endpoint_id
  client_id       = "testid"
  name            = "testname"
  type            = "external"
  generate_secret = true
}
```

## Argument Reference

The following arguments are supported:

* `artifact_id` - (Required,string) The ID of an Endpoint
* `client_id` - (Optional,string)Api Key for generating an API Key. The ID of subscription. Client ID is auto-generated if it is not provided by user.
* `name` - (Required,string) The display name for the API key.
* `type` - (Required,string) Type of sharing of API Key. Allowable values-[`external`],[`internal`]
* `client_secret` - (Optional,string) Secret key of the API key.
* `generate_secret` - (Optional,bool) It indicates if the secret key has to be auto generated or not. If it is true client secret (Secret key) is auto-generated.

**NOTE:** Subscriptions can be performed only if the Endpoint is Online i.e Manged attribute of an endpoint should be true.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The Id of the Subscription resource. It is a combination of <artifact_id>//<client_id>
* `secret_provided` - It indicates if client secret is provided or not. i.e if client secret is provided secret_provided will be true,else false.

## Import

The `ibm_api_gateway_endpoint_subscription` resource can be imported using the `id`. The ID is formed from the `Endpoint ID` and the `Client ID` concatentated using a `//` character.  

* **Endpoint ID** is a string of the form: `b2f2c5b1-a29d-4e0b-ae2f-a0313c3ea2d3`. The id of an existing Endpoint is not avaiable via the UI. It can be retrieved programmatically via the API gateway Endpoint API.

* **Client ID** is a user defined string or a auto generated string of the form `3969e1c0-3571-4134-b14c-bebbb0d56b21`. To View the Client ID `Application authentication` should be enabled on the **Define and secure** page of API Gateway Service. Client ID of a particular subscription will be available as `API key` in the **Manage and Sharing** page of API Gateway Service.

```
$ terraform import ibm_api_gateway_endpoint_subscription.subscription <artifact_id>//<client_id>

$ terraform import ibm_api_gateway_endpoint_subscription.subscription 705fd456-224e-412d-833f-51ff46e27fc8//testID