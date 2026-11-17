---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_paths"
description: |-
  Get information about sm_hsm_paths
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_paths

Provides a read-only data source to retrieve information about sm_hsm_paths. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_hsm_paths" "sm_hsm_paths" {
	destination_id = ibm_sm_hsm_path.sm_hsm_path_instance.destination_id
	instance_id = ibm_sm_hsm_path.sm_hsm_path_instance.instance_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the sm_hsm_paths.
* `paths` - (List) Paths for this destination. At most 20 paths may exist per destination.
  * Constraints: The maximum length is `20` items. The minimum length is `0` items.
Nested schema for **paths**:
	* `created_at` - (String) Computed field.
	* `created_by` - (String) Computed field.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
	* `crn` - (Forces new resource, String) CRN of the gateway through which the HSM is reached. Immutable.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^crn:v[0-9]:[a-z0-9\\-]+:[a-z0-9\\-]+:[a-z0-9\\-]+:[a-z0-9\\-]+:[a-zA-Z0-9\/\\-_:@]+::[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-_:]+$/`.
	* `description` - (String) Optional description. Omitted if not set.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
	* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
	* `health_state` - (String) AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.- `evaluating`: A health check is in progress for this path.- `in_progress`: An apply operation is running across the destination.- `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.- `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval, approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is missing or expired, no_valid_server_ca takes priority.- `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.- `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.
	  * Constraints: Allowable values are: `unknown`, `evaluating`, `in_progress`, `healthy`, `unhealthy`, `no_valid_server_ca`, `not_configured`.
	* `href` - (String) URL of the path resource.
	  * Constraints: The maximum length is `2048` characters.
	* `id` - (String) Path ID. Immutable. System-managed.
	  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `name` - (String) Human-readable display name.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
	* `provisioning_state` - (String) Lifecycle state of the path:- `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.- `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy for the gateway before proceeding.- `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.- `pending_approval`: Connectivity created. Waiting for customer approval.- `completed`: Connectivity established. Path fully operational.- `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.- `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.- `approval_declined`: Customer declined the pending approval. Call POST /reset.- `deleting`: Delete initiated; async cleanup in progress.- `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.
	  * Constraints: Allowable values are: `pending_server_ca`, `pending_account_policy`, `provisioning`, `pending_approval`, `completed`, `vpe_provisioning_failed`, `crn_already_in_use`, `approval_declined`, `deleting`, `deletion_failed`.
	* `updated_at` - (String) Computed field. Omitted until first update.
	* `updated_by` - (String) Computed field. Omitted until first update.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.

