---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_server_ca"
description: |-
  Manages sm_hsm_server_ca.
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_server_ca

Create, update, and delete sm_hsm_server_cas with this resource.

## Example Usage

```hcl
resource "ibm_sm_hsm_server_ca" "sm_hsm_server_ca_instance" {
  certificate = "-----BEGIN CERTIFICATE-----
MIIDXTCCAkWg...
-----END CERTIFICATE-----"
  destination_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
  instance_id = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
  name = "HSM1_Server_CA"
  path_id = "fa3d1234-0000-0000-0000-000000000001"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `certificate` - (Required, Forces new resource, String) PEM-encoded Server CA certificate. Immutable.
  * Constraints: The maximum length is `65536` characters. The minimum length is `1` character. The value must match regular expression `/^-----BEGIN CERTIFICATE-----[\\s\\S]+-----END CERTIFICATE-----\\s*$/`.
* `description` - (Optional, String) Optional description. Omitted if not set.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `name` - (Required, String) Human-readable display name.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `path_id` - (Required, Forces new resource, String) Path ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sm_hsm_server_ca.
* `created_at` - (String) Computed field.
* `created_by` - (String) Computed field.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
* `expires_at` - (String) Expiry parsed from certificate NotAfter field. System-managed. Populated on creation.
* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
* `health_state` - (String) AS OF health state of the Server CA certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation. Health check not yet performed.- `evaluating`: A health check is in progress for this Server CA.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.
  * Constraints: Allowable values are: `unknown`, `evaluating`, `valid`, `expired`.
* `server_ca_id` - (String) Server CA ID. Immutable.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `status` - (String) Only one Server CA can be active per path. Starts inactive on creation.
  * Constraints: Allowable values are: `active`, `inactive`.
* `updated_at` - (String) Computed field. Omitted until first update.
* `updated_by` - (String) Computed field. Omitted until first update.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.


## Import

You can import the `ibm_sm_hsm_server_ca` resource by using `id`.
The `id` property can be formed from `instance_id`, `destination_id`, `path_id`, and `server_ca_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;destination_id&gt;/&lt;path_id&gt;/&lt;server_ca_id&gt;
</pre>
* `instance_id`: A string in the format `fe5c5c27-3e44-474a-9484-55892e12d6d1`. Secrets Manager instance ID.
* `destination_id`: A string in the format `b2c3d4e5-f6a7-8901-bcde-f12345678901`. Destination ID.
* `path_id`: A string in the format `fa3d1234-0000-0000-0000-000000000001`. Path ID.
* `server_ca_id`: A string in the format `5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f`. Server CA ID. Immutable.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager).

# Syntax
<pre>
$ terraform import ibm_sm_hsm_server_ca.sm_hsm_server_ca &lt;instance_id&gt;/&lt;destination_id&gt;/&lt;path_id&gt;/&lt;server_ca_id&gt;
</pre>
