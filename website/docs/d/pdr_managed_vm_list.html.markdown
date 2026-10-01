---
layout: "ibm"
page_title: "IBM : ibm_pdr_managed_vm_list"
description: |-
  Get information about pdr_managed_vm_list
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_managed_vm_list

Retrieves the list of managed vms for the specified service instance.

## Example Usage

```hcl
data "ibm_pdr_managed_vm_list" "pdr_managed_vm_list" {
	accept_language = "en-US"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```
### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Unique identifier of the API key.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `managed_vm_list` - (Map) A map where the key is the VM ID and the value is the corresponding ManagedVmDetails object.

Nested schema for **managed_vm_list**:
	* `core` - (String) The Number of cores assigned to the managed vitual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^\\d+(\\.\\d+)?$/`.
	* `dr_average_time` - (String) The DR operation average time(in minutes) for the managed virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^\\d+$/`.
	* `memory` - (String) The amount of memory (in GB) assigned to the managed virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^\\d+(\\.\\d+)?$/`.
	* `region` - (String) The source region where the managed virtual machine is deployed.
	  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `workgroup_name` - (String) The name of the workgroup where the managed virtual machine is added for disaster recovery.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `workspace_name` - (String) The Name of the power virtual server workspace.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `dr_region` - (String) The name of the region where the virtual machine is recovered.
	  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `vm_name` - (String) The name of the managed virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `vm_id` - (String) The id of the irtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.