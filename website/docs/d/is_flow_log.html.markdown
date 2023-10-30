---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_flow_log"
description: |-
  Get information about Flow Log Collector
---

# ibm_is_flow_log
Retrieve an information of VPC flow log. For more information, about VPC flow log, see [about IBM Cloud flow logs for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-flow-logs).

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

data "ibm_is_flow_log" "example" {
	identifier = ibm_is_flow_log.example.id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `identifier` - (Optional, String) The ID of the flow log collector, This is required when `name` is not specified.
- `name` - (Optional, String) The name of the flow log collector,  This is required when `identifier` is not specified.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `identifier` - The unique identifier of the FlowLogCollector.
- `access_tags` - (String) Access management tags associated for flow log.
- `active` - (Boolean) Indicates whether this collector is active.
- `auto_delete` - (Boolean) If set to `true`, this flow log collector will be automatically deleted when the target is deleted.
- `created_at` - (String) The date and time that the flow log collector was created.
- `crn` - (String) The CRN for this flow log collector.
- `href` - (String) The URL for this flow log collector.
- `lifecycle_state` - (String) The lifecycle state of the flow log collector.
- `name` - (String) The unique user-defined name for this flow log collector.
- `resource_group` - (List) The resource group object, for this flow log collector.

	Nested scheme for `resource_group`:
    - `href` - (Required, String) The URL for this resource group.
    - `id` - (Required, String) The unique identifier for this resource group.
    - `name` - (Required, String) The user-defined name for this resource group.

- `storage_bucket` - (Required, List) The Cloud Object Storage bucket where the collected flows are logged.
  
	Nested scheme for `storage_bucket`:
    - `name` - (Required, String) The globally unique name of this Cloud Object Storage bucket.

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

  -> **Note:**
  **&#x2022;** If the target is an instance network attachment, flow logs will be collected  for that instance network attachment.</br>
  **&#x2022;** If the target is an instance network interface, flow logs will be collected  for that instance network interface.</br>
  **&#x2022;** If the target is a virtual network interface, flow logs will be collected for the  the virtual network interface's `target` resource if the resource is:  - an instance network attachment.</br>
  **&#x2022;** If the target is a virtual server instance, flow logs will be collected  for all network attachments or network interfaces on that instance.- If the target is a subnet, flow logs will be collected  for all instance network interfaces and virtual network interfaces  attached to that subnet.</br>
  **&#x2022;** If the target is a VPC, flow logs will be collected for all instance network  interfaces and virtual network interfaces  attached to all subnets within that VPC. If the target is an instance, subnet, or VPC, flow logs will not be collectedfor any instance network attachments or instance network interfaces within the targetthat are themselves the target of a more specific flow log collector.</br>
  
- `vpc` - (List) The VPC this flow log collector is associated with.
	
	Nested scheme for `vpc`:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.

		Nested scheme for `deleted`:
		 - `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The unique user-defined name for this VPC.
