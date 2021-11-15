---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_flow_log"
description: |-
  Get information about FlowLogCollector
subcategory: "Virtual Private Cloud API"
---

# ibm_is_flow_log

Provides a read-only data source for FlowLogCollector. 
[creating a flow log collector](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-flow-log-collector).

## Example Usage

```terraform

data "ibm_is_flow_log" "is_flow_log" {
	identifier = ibm_is_flow_log.test_flow_log.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `identifier` - (Optional, String) The ID of the subnet, This is required when name is not specified.
- `name` - (Optional, String) The name of the subnet,  This is required when identifier is not specified.
## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `identifier` - The unique identifier of the FlowLogCollector.
- `active` - (Boolean) Indicates whether this collector is active.
- `auto_delete` - (Boolean) If set to `true`, this flow log collector will be automatically deleted when the target is deleted.
- `created_at` - (String) The date and time that the flow log collector was created.
- `crn` - (String) The CRN for this flow log collector.
- `href` - (String) The URL for this flow log collector.
- `lifecycle_state` - (String) The lifecycle state of the flow log collector.
- `name` - (String) The unique user-defined name for this flow log collector.
- `resource_group` - (List) The resource group for this flow log collector.

	Nested scheme for `resource_group`:
    	- `href` - (Required, String) The URL for this resource group.
    	- `id` - (Required, String) The unique identifier for this resource group.
    	- `name` - (Required, String) The user-defined name for this resource group.

- `storage_bucket` - (Required, List) The Cloud Object Storage bucket where the collected flows are logged.
  
	Nested scheme for `storage_bucket`:
    	- `name` - (Required, String) The globally unique name of this COS bucket.

- `target` - (List) The target this collector is collecting flow logs for. If the target is an instance,subnet, or VPC, flow logs will not be collected for any network interfaces within the target that are themselves the target of a more specific flow log collector.

	Nested scheme for `target`:
    	- `crn` - (String) The CRN for this virtual server instance.
    	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
			Nested scheme for `deleted`:
      		- `more_info` - (String) Link to documentation about deleted resources.
    	- `href` - (String) The URL for this network interface.
    	- `id` - (String) The unique identifier for this network interface.
    	- `name` - (String) The user-defined name for this network interface.
    	- `resource_type` - (String) The resource type. Allowable values are: `network_interface`.

- `vpc` - (List) The VPC this flow log collector is associated with.
	
	Nested scheme for `vpc`:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		Nested scheme for `deleted`:
  			-`more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The unique user-defined name for this VPC.

