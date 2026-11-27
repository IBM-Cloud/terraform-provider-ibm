---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_client_certificate"
description: |-
  Manages sm_hsm_client_certificate.
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_client_certificate

Create, update, and delete sm_hsm_client_certificates with this resource.

## Example Usage

```hcl
resource "ibm_sm_hsm_client_certificate" "sm_hsm_client_certificate_instance" {
  destination_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
  instance_id = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
  name = "hsm-client-cert-prod"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) Optional description. Omitted if not set.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `name` - (Required, String) Human-readable display name.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sm_hsm_client_certificate.
* `cert_id` - (String) Certificate ID. Immutable.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `created_at` - (String) Computed field.
* `created_by` - (String) Computed field.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
* `csr` - (String) PEM-encoded CSR. Platform-generated. Immutable once set. Populated on creation (always present once the resource exists).
  * Constraints: The maximum length is `65536` characters. The minimum length is `1` character. The value must match regular expression `/^-----BEGIN CERTIFICATE REQUEST-----[\\s\\S]+-----END CERTIFICATE REQUEST-----\\s*$/`.
* `csr_cn` - (String) Common Name embedded in the CSR. Set to the SM instance UUID at generation time. Immutable once set. Populated on creation (always present once the resource exists).
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `expires_at` - (String) Parsed from signed_cert NotAfter field. Omitted until signed_cert is uploaded.
* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
* `health_state` - (String) AS OF health state of the client certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.- `evaluating`: A health check is in progress for this client certificate.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.
  * Constraints: Allowable values are: `unknown`, `evaluating`, `valid`, `expired`.
* `href` - (String) URL of the client certificate resource.
  * Constraints: The maximum length is `2048` characters.
* `provisioning_state` - (String) Lifecycle state of the client certificate:- `csr_generated`: CSR and private key generated successfully. `csr` and `csr_cn` are populated. Initial state on POST. Waiting for signed_cert to be uploaded via PATCH.- `inactive`: signed_cert has been uploaded but the certificate has not been activated yet.- `active`: This certificate is the active client certificate for the destination. Only one can be active at a time.
  * Constraints: Allowable values are: `csr_generated`, `inactive`, `active`.
* `signed_cert` - (String) CA-signed certificate (PEM). Omitted until uploaded via PATCH.
  * Constraints: The maximum length is `65536` characters. The minimum length is `1` character. The value must match regular expression `/^-----BEGIN CERTIFICATE-----[\\s\\S]+-----END CERTIFICATE-----\\s*$/`.
* `status` - (String) Request-only. Set to `active` in a PATCH to activate this certificate. Never returned; read the current state from `provisioning_state`.
  * Constraints: Length must be `6` characters. The value must match regular expression `/^active$/`.
* `updated_at` - (String) Computed field. Omitted until first update.
* `updated_by` - (String) Computed field. Omitted until first update.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.


## Import

You can import the `ibm_sm_hsm_client_certificate` resource by using `id`.
The `id` property can be formed from `instance_id`, `destination_id`, and `cert_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;destination_id&gt;/&lt;cert_id&gt;
</pre>
* `instance_id`: A string in the format `fe5c5c27-3e44-474a-9484-55892e12d6d1`. Secrets Manager instance ID.
* `destination_id`: A string in the format `b2c3d4e5-f6a7-8901-bcde-f12345678901`. Destination ID.
* `cert_id`: A string in the format `7c9e6679-7425-40de-944b-e07fc1f90ae7`. Certificate ID. Immutable.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager).

# Syntax
<pre>
$ terraform import ibm_sm_hsm_client_certificate.sm_hsm_client_certificate &lt;instance_id&gt;/&lt;destination_id&gt;/&lt;cert_id&gt;
</pre>
