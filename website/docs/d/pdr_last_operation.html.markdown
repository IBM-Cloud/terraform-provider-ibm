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
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_last_operation.
* `crn` - (String) Cloud Resource Name (CRN) of the service instance.
* `deployment_name` - (String) Name of the service instance deployment.
* `is_ksys_ha` - (Boolean) Indicates whether high availability (HA) is enabled for the orchestrator.
* `last_updated_orchestrator_deployment_time` - (String) Deployment time of primary orchestrator VM.
* `last_updated_standby_orchestrator_deployment_time` - (String) Deployment time of StandBy orchestrator VM.
* `mfa_enabled` - (String) Multiple Factor Authentication Enabled or not.
* `orch_ext_connectivity_status` - (String) Status of standby node addition to the orchestrator cluster.
* `orch_standby_node_addtion_status` - (String) Health or informational message about the orchestrator cluster.
* `orchestrator_cluster_message` - (String) Current status of the primary orchestrator VM.
* `orchestrator_config_status` - (String) Configuration status of the orchestrator cluster.
* `plan_name` - (String) Name of the Plan.
* `primary_description` - (String) Detailed status message for the primary orchestrator VM.
* `primary_ip_address` - (String) IP address of the primary orchestrator VM.
* `primary_orchestrator_status` - (String) Configuration status of the orchestrator cluster.
* `recovery_location` - (String) Disaster recovery location associated with the instance.
* `resource_group` - (String) Resource group to which the service instance belongs.
* `standby_description` - (String) Detailed status message for the standby orchestrator VM.
* `standby_ip_address` - (String) IP address of the standby orchestrator VM.
* `standby_status` - (String) Current state of the standby orchestrator VM.
* `status` - (String) Overall status of the service instance.

