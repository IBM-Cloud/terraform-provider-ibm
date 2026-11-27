---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_server_ca"
description: |-
  Get information about sm_hsm_server_ca
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_server_ca

Provides a read-only data source to retrieve information about a sm_hsm_server_ca. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_hsm_server_ca" "sm_hsm_server_ca" {
	destination_id = ibm_sm_hsm_server_ca.sm_hsm_server_ca_instance.destination_id
	instance_id = ibm_sm_hsm_server_ca.sm_hsm_server_ca_instance.instance_id
	path_id = ibm_sm_hsm_server_ca.sm_hsm_server_ca_instance.path_id
	server_ca_id = ibm_sm_hsm_server_ca.sm_hsm_server_ca_instance.server_ca_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `path_id` - (Required, Forces new resource, String) Path ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `server_ca_id` - (Required, Forces new resource, String) The identifier of the Server CA to operate on. Matches the `id` field in the `ServerCa` response schema.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `region` - (Optional, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the sm_hsm_server_ca.
* `certificate` - (Forces new resource, String) PEM-encoded Server CA certificate. Immutable.
  * Constraints: The maximum length is `65536` characters. The minimum length is `1` character. The value must match regular expression `/^-----BEGIN CERTIFICATE-----[\\s\\S]+-----END CERTIFICATE-----\\s*$/`.
* `created_at` - (String) Computed field.
* `created_by` - (String) Computed field.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
* `description` - (String) Optional description. Omitted if not set.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `expires_at` - (String) Expiry parsed from certificate NotAfter field. System-managed. Populated on creation.
* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
* `health_state` - (String) AS OF health state of the Server CA certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation. Health check not yet performed.- `evaluating`: A health check is in progress for this Server CA.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.
  * Constraints: Allowable values are: `unknown`, `evaluating`, `valid`, `expired`.
* `name` - (String) Human-readable display name.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `status` - (String) Only one Server CA can be active per path. Starts inactive on creation.
  * Constraints: Allowable values are: `active`, `inactive`.
* `updated_at` - (String) Computed field. Omitted until first update.
* `updated_by` - (String) Computed field. Omitted until first update.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.

