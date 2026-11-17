---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_ha_group"
description: |-
  Get information about sm_hsm_ha_group
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_ha_group

Provides a read-only data source to retrieve information about a sm_hsm_ha_group. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_hsm_ha_group" "sm_hsm_ha_group" {
	destination_id = ibm_sm_hsm_ha_group.sm_hsm_ha_group_instance.destination_id
	ha_group_id = ibm_sm_hsm_ha_group.sm_hsm_ha_group_instance.ha_group_id
	instance_id = ibm_sm_hsm_ha_group.sm_hsm_ha_group_instance.instance_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `ha_group_id` - (Required, Forces new resource, String) The identifier of the HA Group to operate on. Matches the `id` field in the `HSMHAGroup` response schema.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the sm_hsm_ha_group.
* `created_at` - (String) Computed field.
* `created_by` - (String) Computed field.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
* `description` - (String) Optional description. Omitted if not set.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `group_label` - (Forces new resource, String) HA Group virtual token label. Immutable after creation.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~]+$/`.
* `group_serial` - (Forces new resource, String) HA Group serial number. Immutable after creation.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.
* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
* `health_state` - (String) AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. No health check has been performed yet.- `evaluating`: A health check is in progress for this resource.- `in_progress`: An apply operation is running across the destination.- `healthy`: Configured AND all member HSM nodes are healthy.- `degraded`: Configured AND any one member is unhealthy (but not all).- `unhealthy`: Configured AND all members are unhealthy or unknown.- `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active non-expired client certificate).
  * Constraints: Allowable values are: `unknown`, `evaluating`, `in_progress`, `healthy`, `degraded`, `unhealthy`, `not_configured`.
* `href` - (String) URL of the HA Group resource.
  * Constraints: The maximum length is `2048` characters.
* `member_serials` - (List) Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members.
  * Constraints: The list items must match regular expression `/^[0-9]+$/`. The maximum length is `100` items. The minimum length is `1` item.
* `name` - (String) Human-readable display name.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `provisioning_state` - (String) Lifecycle state of the HA Group:- `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only state.**Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values gracefully and MUST NOT fail on unrecognised states.
  * Constraints: Allowable values are: `completed`.
* `updated_at` - (String) Computed field. Omitted until first update.
* `updated_by` - (String) Computed field. Omitted until first update.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.

