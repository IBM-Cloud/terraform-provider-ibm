---
layout: "ibm"
page_title: "IBM : ibm_cloudant"
description: |-
  Get information about Cloudant instance.
subcategory: "Cloud Databases"
---

# ibm_cloudant

Provides a read-only data source for an existing IBM Cloud Cloudant service. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
resource "ibm_resource_instance" "cloudant" {
    name     = "cloudant-service-name"
    service  = "cloudantnosqldb"
    plan     = "lite"
    location = "us-south"
 }

 data "ibm_cloudant" "instance" {
    name     = ibm_resource_instance.cloudant.name
 }
```

## Argument Reference

The following arguments are supported:

* `name` (Required, String) The name of the IBM Cloudant resource instance.
* `id` (Optional, String) The unique identifier of the Cloudant resource.
* `location` (Optional, String) The location or the environment in which instance exists.
* `resource_group_id` (Optional, String) The id of the resource group in which the instance is present. If not provided it takes the default resource group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `capacity` (Number) A number of blocks of throughput units. For more details please read about [`blocks`](https://cloud.ibm.com/apidocs/cloudant#putcapacitythroughputconfiguration) parameter.
* `cors_config` (List of Object) Configuration for CORS. (see [below for nested attributes](#nestedatt--cors_config))
    * `allow_credentials` (Boolean) - Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.
    * `origins` (List of String) - Contains the list of allowed origin domains with the full URL including the protocol. Subdomains count as separate domains, so all subdomains used have to be listed.
* `crn` (String) CRN of resource instance.
* `enable_cors` (Boolean) Boolean value to turn CORS on and off.
* `extensions` (Map of String) The extended metadata as a map associated with the resource instance.
* `features` (List of String) List of enabled optional features.
* `features_flags` (List of String) List of feature flags.
* `guid` (String) Guid of resource instance.
* `include_data_events` (Boolean) Include `data` event types in events sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. By default emitted events are only of `management` type.
* `plan` (String) The plan type of the instance.
* `resource_controller_url` (String) The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource.
* `resource_crn` (String) The crn of the resource.
* `resource_group_name` (String) The resource group name in which resource is provisioned.
* `resource_name` (String) The name of the resource.
* `resource_status` (String) The status of the resource.
* `service` (String) The service type of the instance.
* `status` (String) The resource instance status.
* `throughput` (Map of Number) Schema for detailed information about throughput capacity with breakdown by specific throughput requests classes.
* `version` (String) Vendor version.
