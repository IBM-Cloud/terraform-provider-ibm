---
layout: "ibm"
page_title: "IBM : ibm_pdr_last_operation"
description: |-
  Get information about pdr_last_operation
subcategory: "DrAutomation Service"
---

# ibm_pdr_last_operation

Provides a read-only data source to retrieve information about a pdr_last_operation. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_last_operation" "pdr_last_operation" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_last_operation.
* `crn` - (String) 
* `deployment_name` - (String) 
* `is_ksys_ha` - (Boolean) 
* `orch_ext_connectivity_status` - (String) 
* `orch_standby_node_addtion_status` - (String) 
* `orchestrator_cluster_message` - (String) 
* `orchestrator_config_status` - (String) 
* `primary_description` - (String) 
* `primary_ip_address` - (String) 
* `primary_orchestrator_status` - (String) 
* `recovery_location` - (String) 
* `resource_group` - (String) 
* `standby_description` - (String) 
* `standby_ip_address` - (String) 
* `standby_status` - (String) 
* `status` - (String) 

