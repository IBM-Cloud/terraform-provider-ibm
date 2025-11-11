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

* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_deployment_status.
* `orch_ext_connectivity_status` - (String) External connectivity status of the orchestrator cluster.
* `orch_standby_node_addition_status` - (String) Status of the standby node addition process.
* `orchestrator_cluster_message` - (String) Cluster status message.
* `orchestrator_cluster_type` - (String) Type of orchestrator cluster.
* `orchestrator_config_status` - (String) Configuration status of the orchestrator.
* `orchestrator_group_leader` - (String) Name of the orchestrator acting as the cluster leader.
* `orchestrator_name` - (String) Name of the primary orchestrator.
* `orchestrator_status` - (String) Status of the primary orchestrator.
* `schematic_workspace_name` - (String) Name of the schematic workspace.
* `schematic_workspace_status` - (String) Status of the schematic workspace.
* `ssh_key_name` - (String) Name of the SSH key associated with the orchestrator.
* `standby_orchestrator_name` - (String) Name of the standby orchestrator.
* `standby_orchestrator_status` - (String) Status of the standby orchestrator.

