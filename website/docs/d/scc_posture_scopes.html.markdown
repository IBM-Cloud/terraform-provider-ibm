---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scopes"
description: |-
  Get information about list_scopes
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scopes

Review information of Security and Compliance Center posture scopes see [Managing Scopes](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-scopes).

## Example usage

```terraform
data "ibm_scc_posture_scopes" "list_scopes" {
	scope_id = "1"
}
```

## Argument reference

Review information of Security and Compliance Center posture scopes

* `scope_id` - (Optional, String) An auto-generated unique identifier for the scope.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the list_scopes.
* `scopes` - (Optional, List) Scopes.
Nested scheme for **scopes**:
	* `created_by` - (Optional, String) The user who created the scope.
	* `created_time` - (Optional, String) The time that the scope was created in UTC.
	* `collectors_id` - (Optional, List) The unique IDs of the collectors that are attached to the scope.
	* `description` - (Optional, String) A detailed description of the scope.
	* `enabled` - (Optional, Boolean) Indicates whether scope is enabled/disabled.
	* `environment_type` - (Optional, String) The environment that the scope is targeted to.
		* Constraints: Supported values are: **ibm**, **aws**, **azure**, **on_premise**, **hosted**, **services**, **openstack**, **gcp**
	* `last_scan_status_updated_time` - (Optional, String) The last time that a scan status for a scope was updated in UTC.
	* `last_scan_type` - (Optional, String) The last type of scan that was run on the scope.
		* Constraints: Allowable values are: **discovery**, **validation**, **fact_collection**, **fact_validation**, **inventory**, **remediation**, **abort_tasks**, **evidence**, **script**
	* `last_scan_type_description` - (Optional, String) A description of the last scan type.
	* `modified_by` - (Optional, String) The user who most recently modified the scope.
	* `modified_time` - (Optional, String) The time that the scope was last modified in UTC.
	* `name` - (Optional, String) A unique name for your scope.
	* `scope_id` - (Optional, String) An auto-generated unique identifier for the scope.
	* `scans` - (Optional, List) A list of the scans that have been run on the scope.
	Nested scheme for **scans**:
		* `discover_id` - (Optional, String) An auto-generated unique identifier for discovery.
		* `scan_id` - (Optional, String) An auto-generated unique identifier for the scan.
		* `status` - (Optional, String) The status of the collector as it completes a scan.
		  * Constraints: Supported values are: **pending**, **discovery_started**, **discovery_completed**, **error_in_discovery**, **gateway_aborted**, **controller_aborted**, **not_accepted**, **waiting_for_refine**, **validation_started**, **validation_completed**, **sent_to_collector**, **discovery_in_progress**, **validation_in_progress**, **error_in_validation**, **discovery_result_posted_with_error**, **discovery_result_posted_no_error**, **validation_result_posted_with_error**, **validation_result_posted_no_error**, **fact_collection_started**, **fact_collection_in_progress**, **fact_collection_completed**, **error_in_fact_collection**, **fact_validation_started**, **fact_validation_in_progress**, **fact_validation_completed**, **error_in_fact_validation**, **abort_task_request_received**, **error_in_abort_task_request**, **abort_task_request_completed**, **user_aborted**, **abort_task_request_failed**, **remediation_started**, **remediation_in_progress**, **error_in_remediation**, **remediation_completed**, **inventory_started**, **inventory_in_progress**, **inventory_completed**, **error_in_inventory**, **inventory_completed_with_error**
		* `status_message` - (Optional, String) The current status of the collector.

