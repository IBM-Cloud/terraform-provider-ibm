---
layout: "ibm"
page_title: "IBM: API Gateway Endpoint"
sidebar_current: "docs-ibm-resource-api-gateway-endpoint"
description: |-
  Create, Update,Manage and Delete Endpoints.
---

# ibm\_api_gateway_endpoint

Provides an Endpoint. This allows endpoint to be created, updated managed and deleted.

## Example Usage for a single Api doc as input

```hcl

resource "ibm_resource_instance" "apigateway"{
    name     = "testname"
    location = "global"
    service  = "api-gateway"
    plan     = "lite"
 }

resource "ibm_api_gateway_endpoint" "endpoint"{
    service_instance_crn = ibm_resource_instance.apigateway.id
    name="test-endpoint"
    managed="true"
    open_api_doc_name = var.file_path
    type="share" //optional while creating endpoint  required during updating actions of an endpoint
}
```
## Example Usage for a directory of Api docs as input

```hcl

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

## Argument Reference

The following arguments are supported:

* `service_instance_crn` - (Required,string) The CRN-based service instance ID.
* `name` - (Required,string) The display name for the Endpoint. Name is optional for creating an endpoint.But,it is required attribute for updating an endpoint.
* `managed` - (Optional,bool)(Default - false) Managed indicates if endpoint is online or offline.
* `open_api_doc_name` - (Required,string) API document that represents endpoint
* `routes` - (Optional,list) Invokable routes for an endpoint
* `provider_id` - (Optional,string)(Default - [`user-defined`]) Provider ID of an endpoint. Allowable values-[`user-defined`],[`whisk`]
* `type` - (Optional,string) (Default - unshare)Type of the action that is to be performed on endpoint. Allowable values-[`share`],[`unshare`],[`manage`],[`unmanage`]

**NOTE:** 
1. Endpoint actions are performed using 'type' argument to manage the actions, only after the endpoint is created .There fore endpoint actions are invoked during endpoint update function.

2. Basepath of the endpoint provided should be unique.

3. Endpoint cannot be shared if manage attribute is false i.e API cannot be shared if it is offline.

## Attribute Reference

The following attributes are exported:

* `id` - The Id of the Endpoint resource. It is a combination of <service_instance_crn>//<endpoint_ID>
* `endpoint_id` - It is also called as Artifact ID. The Id of an Endpoint
* `shared` - The Shared status of an endpoint
* `base_path` - Base path of an endpoint

## Import

The `ibm_api_gateway_endpoint` resource can be imported using the `id`. The ID is formed from the `Service Instance CRN` and the `Endpoint ID` concatentated using a `//` character.  

The CRN will be located on the **Overview** page of the API Gateway Service. 

* **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

* **Endpoint ID** is a string of the form: `b2f2c5b1-a29d-4e0b-ae2f-a0313c3ea2d3`. The id of an existing Endpoint is not avaiable via the UI. It can be retrieved programmatically via the API gateway Endpoint API 


```
$ terraform import ibm_api_gateway_endpoint.endpoint <service_instance_crn>//<endpoint_id>

$ terraform import ibm_api_gateway_endpoint.endpoint crn:v1:bluemix:public:api-gateway:global:a/4448261269a14562b839e0a3019ed980:a608ff36-7037-4da1-b7e9-663e3e1c6254:://b2f2c5b1-a29d-4e0b-ae2f-a0313c3ea2d3