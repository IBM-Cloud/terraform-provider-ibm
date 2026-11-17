---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_ha_group"
description: |-
  Manages sm_hsm_ha_group.
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_ha_group

Create, update, and delete sm_hsm_ha_groups with this resource.

## Example Usage

```hcl
resource "ibm_sm_hsm_ha_group" "sm_hsm_ha_group_instance" {
  description = "Primary HA Group for prod"
  destination_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
  group_label = "prod_ha_group"
  group_serial = "11323465009866"
  instance_id = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
  name = "production-ha-group"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) Optional description. Omitted if not set.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `group_label` - (Required, Forces new resource, String) HA Group virtual token label. Immutable after creation.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~]+$/`.
* `group_serial` - (Required, Forces new resource, String) HA Group serial number. Immutable after creation.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `member_serials` - (Required, List) Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members.
  * Constraints: The list items must match regular expression `/^[0-9]+$/`. The maximum length is `100` items. The minimum length is `1` item.
* `name` - (Required, String) Human-readable display name.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sm_hsm_ha_group.
* `created_at` - (String) Computed field.
* `created_by` - (String) Computed field.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
* `ha_group_id` - (String) HA Group ID. Immutable. System-managed.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
* `health_state` - (String) AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. No health check has been performed yet.- `evaluating`: A health check is in progress for this resource.- `in_progress`: An apply operation is running across the destination.- `healthy`: Configured AND all member HSM nodes are healthy.- `degraded`: Configured AND any one member is unhealthy (but not all).- `unhealthy`: Configured AND all members are unhealthy or unknown.- `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active non-expired client certificate).
  * Constraints: Allowable values are: `unknown`, `evaluating`, `in_progress`, `healthy`, `degraded`, `unhealthy`, `not_configured`.
* `href` - (String) URL of the HA Group resource.
  * Constraints: The maximum length is `2048` characters.
* `provisioning_state` - (String) Lifecycle state of the HA Group:- `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only state.**Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values gracefully and MUST NOT fail on unrecognised states.
  * Constraints: Allowable values are: `completed`.
* `updated_at` - (String) Computed field. Omitted until first update.
* `updated_by` - (String) Computed field. Omitted until first update.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.


## Import

You can import the `ibm_sm_hsm_ha_group` resource by using `id`.
The `id` property can be formed from `instance_id`, `destination_id`, and `ha_group_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;destination_id&gt;/&lt;ha_group_id&gt;
</pre>
* `instance_id`: A string in the format `fe5c5c27-3e44-474a-9484-55892e12d6d1`. Secrets Manager instance ID.
* `destination_id`: A string in the format `b2c3d4e5-f6a7-8901-bcde-f12345678901`. Destination ID.
* `ha_group_id`: A string in the format `ee52ebb6-1728-4580-8ede-13f6504e3ae0`. HA Group ID. Immutable. System-managed.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager).

# Syntax
<pre>
$ terraform import ibm_sm_hsm_ha_group.sm_hsm_ha_group &lt;instance_id&gt;/&lt;destination_id&gt;/&lt;ha_group_id&gt;
</pre>
