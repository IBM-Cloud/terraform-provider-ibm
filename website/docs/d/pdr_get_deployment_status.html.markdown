---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_deployment_status"
description: |-
  Get information about pdr_get_deployment_status
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_deployment_status

Provides a read-only data source to retrieve information about pdr_get_deployment_status. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_get_deployment_status" "pdr_get_deployment_status" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_deployment_status.
* `orch_ext_connectivity_status` - (String) 
* `orch_standby_node_addition_status` - (String) 
* `orchestrator_cluster_message` - (String) 
* `orchestrator_cluster_type` - (String) 
* `orchestrator_config_status` - (String) 
* `orchestrator_group_leader` - (String) 
* `orchestrator_name` - (String) 
* `orchestrator_status` - (String) 
* `schematic_workspace_name` - (String) 
* `schematic_workspace_status` - (String) 
* `ssh_key_name` - (String) 
* `standby_orchestrator_name` - (String) 
* `standby_orchestrator_status` - (String) 

