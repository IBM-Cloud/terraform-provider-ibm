---
subcategory: "API Gateway"
layout: "ibm"
page_title: "IBM: API Gateway Endpoint Subscription"
description: |-
  Create, Update,Manage and Delete Subscription of Endpoints of a API Gateway service instance.
---

# `ibm_api_gateway_endpoint_subscription`

Create, update, or delete a subscription for an API Gateway endpoint. For more information, about API Gateway subscription, see [Managing access for API Gateway](https://cloud.ibm.com/docs/api-gateway?topic=api-gateway-iam).

Endpoint subscriptions can be added only if the endpoint is online. Make sure to set the `managed` input parameter to **true** in the `ibm_api_gateway_endpoint` resource.

Configuration of an api gateway data resource requires the region parameter to be set for the IBM provider in the `provider.tf`.  If not specified, endpoint to which subscription is being created will default to `us-south`. 

## Example usage
```
data "ibm_api_gateway" "endpoint"{
    service_instance_crn =ibm_resource_instance.apigateway.id
}

resource "ibm_api_gateway_endpoint_subscription" "subs" {
  artifact_id = data.ibm_api_gateway.endpoint.endpoints[0].endpoint_id
  client_id   = "testid"
  name        = "testname"
  type        = "external"
  generate_secret = true
}
```
data "ibm_api_gateway" "endpoint"{
    service_instance_crn =ibm_resource_instance.apigateway.id
}


## Argument reference 
Review the input parameters that you can specify for your resource. 

- `artifact_id` - (Required, String) The ID of an API endpoint. 
- `client_id` - (Optional, String) The API key to generate an API key for the subscription. The generated API key represents the ID of a subscription.
- `client_secret` - (Optional, String) The secret of the API key.
- `generate_secret` - (Optional, Boolean) If set to **true**, the secret key is auto-generated. If set to **false**, the secret key is not auto-generated.
- `name` - (Required, String) The name for an API key.
- `type` - (Required, String) The type of API key sharing. Supported values are `External`, and `Bluemix`.


**Note**
Subscriptions can be performed only if the endpoint is online that is managed attribute of an endpoint should be true.

## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `id` - (String) The ID of the subscription resource. The ID is composed of `<artifact_id>//<client_id>`.
- `secret_provided`- (Boolean) Indicates your secrete if provided or not. If set to **true**, the client secret is provided. If set to **false**, the client secret is not provided.


### Import
The `ibm_api_gateway_endpoint_subscription` resource can be imported by using the ID. The ID is composed of `<artifact_id>//<client_id>`.

- **Endpoint ID**: The Endpoint ID can be retrieved programmatically via the API Gateway endpoint API.
- **Client ID**: The Client ID is an auto-generated string. To view the client ID in the IBM Cloud console, you must enable **Application authentication** on the **Define and secure** page of the API Gateway service. The client ID of a particular subscription is available as an API key in the **Manage and Sharing** page of the API Gateway service.

**Syntax**

```
terraform import ibm_api_gateway_endpoint_subscription.subscription <artifact_id>//<client_id>
```

**Example**

```
terraform import ibm_api_gateway_endpoint_subscription.subscription 705fd456-224e-412d-833f-51ff46e27fc8//testID
```






