---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_destination"
description: |-
  Manages sm_hsm_destination.
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_destination

Create, update, and delete sm_hsm_destinations with this resource.

## Example Usage

```hcl
resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
  description = "Primary Thales Luna destination"
  instance_id = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
  name = "prod-hsm-destination"
  vendor = "thales_luna"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) Optional description. Omitted if not set.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `name` - (Required, String) Human-readable display name.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `vendor` - (Required, Forces new resource, String) HSM vendor. Immutable after creation.
  * Constraints: Allowable values are: `thales_luna`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sm_hsm_destination.
* `account_id` - (String) IBM Cloud account ID. System-managed. Required for account policy setup.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\/\\-_]+$/`.
* `created_at` - (String) Computed field. When the resource was created.
* `created_by` - (String) IAM identity of creator. Computed field.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
* `destination_id` - (String) Destination ID. Immutable. System-managed.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
* `health_state` - (String) AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.- `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.- `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result once apply completes (or set to `unknown` on failure).- `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1 active non-expired client certificate).- `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.- `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.- `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.
  * Constraints: Allowable values are: `unknown`, `evaluating`, `in_progress`, `healthy`, `degraded`, `unhealthy`, `not_configured`.
* `href` - (String) URL of the destination resource.
  * Constraints: The maximum length is `2048` characters.
* `provisioning_state` - (String) Lifecycle state of the destination:- `completed`: Destination created and available for configuration. Initial state on POST.- `deleting`: Hard delete initiated; asynchronous cleanup in progress.- `deletion_failed`: Asynchronous delete failed; the destination still exists. Retry DELETE to re-initiate.
  * Constraints: Allowable values are: `completed`, `deleting`, `deletion_failed`.
* `updated_at` - (String) Computed field. Omitted until first update.
* `updated_by` - (String) IAM identity of last updater. Computed field. Omitted until first update.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.


## Import

You can import the `ibm_sm_hsm_destination` resource by using `id`.
The `id` property can be formed from `instance_id`, and `destination_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;destination_id&gt;
</pre>
* `instance_id`: A string in the format `fe5c5c27-3e44-474a-9484-55892e12d6d1`. Secrets Manager instance ID.
* `destination_id`: A string in the format `b49ad385-0cb7-3694-a71c-2579d6e94b02`. Destination ID. Immutable. System-managed.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager).

# Syntax
<pre>
$ terraform import ibm_sm_hsm_destination.sm_hsm_destination &lt;instance_id&gt;/&lt;destination_id&gt;
</pre>
