---
subcategory: "API Gateway"
layout: "ibm"
page_title: "IBM: API Gateway Endpoint"
description: |-
  Create, update, manage and delete endpoints.
---


# `ibm_api_gateway_endpoint`

Create, update, or delete an API endpoint for an API gateway. For more information, about API Gateway custom endpoint, see [Customizing the domain for an API endpoint](https://cloud.ibm.com/docs/api-gateway?topic=api-gateway-getting-started).

## Example usage

#### Example for a single  API Documentation as input

```
resource "ibm_resource_instance" "apigateway"{
    name     = "testname"
    location = "global"
    service  = "api-gateway"
    plan     = "lite"
 }resource "ibm_api_gateway_endpoint" "endpoint"{
    service_instance_crn = ibm_resource_instance.apigateway.id
    name="test-endpoint"
    managed="true"
    open_api_doc_name = var.file_path
    type="share" //optional while creating endpoint  required during updating actions of an endpoint
}
```

#### Example for a directory of  API Documentation as input

```
resource "ibm_resource_instance" "apigateway"{
    name     = "testname"
    location = "global"
    service  = "api-gateway"
    plan     = "lite"
 }

resource "ibm_api_gateway_endpoint" "endpoint"{
    for_each=fileset(var.dir_path, "*.json")
    service_instance_crn = ibm_resource_instance.apigateway.id
    managed="true"
    name=replace("endpoint-${each.key}",".json","")
    open_api_doc_name=format("%s%s",var.dir_path,each.key)
    type="share" //optional while creating endpoint  required during updating actions of an endpoint
}
}
```

## Argument reference 
Review the input parameters that you can specify for your resource. 

- `managed` - (Optional, Boolean) If set to **true**, the endpoint is online. If set to **false**, the endpoint is offline. The default value is false. The API endpoint cannot be shared if this value is set to **false**.
- `name` - (Required, String) The name of the API Gateway endpoint. This value is optional when you create an API Gateway endpoint, but required when you update the endpoint.
- `open_api_doc_name` - (Required, String) The API document that represents the endpoint.
- `provider_id` - (Optional, String) The provider ID of an API endpoint. Supported values are `user-defined`, and `whisk`. The default value is `user-defined`.
- `routes` (List , Optional) The routes that you can invoke for an endpoint.
- `service_instance_crn` - (Required, String) The CRN of the service instance.
- `type` - (Optional, String) The type of action that is performed on the API endpoint. Supported values are `share`, `unshare`, `manage`, and `unmanage`. The default value is `unshare`. The endpoint actions are performed by using the `type` parameter after the endpoint is created. As a consequence, endpoint actions are invoked during an endpoint update only.

**Note**

* Endpoint actions are performed using `type` argument to manage the actions, only after the endpoint is created. The endpoint actions are invoked during endpoint update function.

* Basepath of the endpoint provided should be unique.

* Endpoint cannot be shared if manage attribute is set to `false` that means, the API cannot be shared if it is offline.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `base_path` - (String) The base path of an endpoint. The base paths must be unique.
- `endpoint_id` - (String) The ID of the endpoint, also referred to as the artifact ID.
- `id` - (String) The ID of the endpoint. The ID is composed of `<service_instance_crn>//<endpoint_ID>`.
- `shared` - (String) The shared status of an endpoint.


### Import
The `ibm_api_gateway_endpoint` resource can be imported by using the ID. The ID is composed of `<service_instance_crn>//<endpoint_ID>`.

The CRN will be located on the **Overview** page of the API Gateway Service.

* **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

* **Endpoint ID** is a string of the form: `b2f2c5b1-a29d-4e0b-ae2f-a0313c3ea2d3`. The ID of an existing endpoint is not avaiable through the console. It can be retrieved programmatically through the API Gateway Endpoint API.

**Syntax**

```
$ terraform import ibm_api_gateway_endpoint.endpoint <service_instance_crn>//<endpoint_id>

```

**Example**

```

$ terraform import ibm_api_gateway_endpoint.endpoint crn:v1:bluemix:public:api-gateway:global:a/4448261269a14562b839e0a3019ed980:a608ff36-7037-4da1-b7e9-663e3e1c6254:://b2f2c5b1-a29d-4e0b-ae2f-a0313c3ea2d3

```
