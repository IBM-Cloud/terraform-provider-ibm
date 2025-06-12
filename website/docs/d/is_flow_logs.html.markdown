---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_flow_logs"
description: |-
  Manages IBM Cloud Infrastructure Flow Logs.
---

# ibm_is_flow_logs
Retrieve an information of an existing IBM Cloud Infrastructure flow logs as a read-only data source. For more information, about VPC flow log, see [creating a flow log collector](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-flow-log-collector).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```


## Example usage

```terraform

data "ibm_is_flow_logs" "example" {
}

data "ibm_is_flow_logs" "example1" {
  name = "my-flow-log"
  vpc_name = ibm_is_vpc.example.name
  vpc_crn = ibm_is_vpc.example.crn
  vpc = ibm_is_vpc.example.id
  resource_group = ibm_is_vpc.example.resource_group
  target = ibm_is_vpc.example.id
  target_resource_type = "vpc"
}
```
## Argument reference

Review the argument references that you can specify for your data source. 

- `name` - (String) The name of the flow log collector
- `vpc_name` - (String) The name of the VPC this flow log collector resides in
- `vpc_crn` - (String) The CRN of the VPC this flow log collector resides in
- `vpc` - (String) The ID of the VPC this flow log collector resides in
- `resource_group` - (String) The ID of the Resource group this flow log collector belongs to
- `target` - (String) The ID of the target this collector is collecting flow logs for.
- `target_resource_type` - (String) The target resource type for this flow log collector. Available options are `instance`, `instance_network_attachment`, `network_interface`, `subnet`, `vpc`, `virtual_network_interface`
 
## Attribute reference
Review the attribute references that you can access after you retrieve your data source. 

- `flow_log_collectors` - (List) Lists all the flow logs in the IBM Cloud.

  Nested scheme for `flow_log_collectors`:
    - `access_tags` - (String) Access management tags associated for flow log.
	- `active` - (String) Indicates whether the collector is active.
	- `created_at` - (Timestamp) The date and time the flow log created.
	- `crn` - (String) The CRN of the flow log collector.
	- `href` - (String) The URL of the flow log collector.
	- `id` - (String) The unique identifier of the flow log collector.
	- `lifecycle_state` - (String) The lifecycle state of the flow log collector.
	- `name` - (String) The flow log collector name.
	- `resource_group` - (String) The resource group Id of the flow log.
	- `storage_bucket` - (String) The IBM Cloud Object Storage bucket name where the flow logs are logged.
	- `target` - (String) The target ID that the flow log collector collects the flow logs.
	- `vpc` - (String) The VPC of the flow log collector that are associated.



