---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_client_certificate"
description: |-
  Get information about sm_hsm_client_certificate
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_client_certificate

Provides a read-only data source to retrieve information about a sm_hsm_client_certificate. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_hsm_client_certificate" "sm_hsm_client_certificate" {
	cert_id = ibm_sm_hsm_client_certificate.sm_hsm_client_certificate_instance.cert_id
	destination_id = ibm_sm_hsm_client_certificate.sm_hsm_client_certificate_instance.destination_id
	instance_id = ibm_sm_hsm_client_certificate.sm_hsm_client_certificate_instance.instance_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `cert_id` - (Required, Forces new resource, String) The identifier of the Client Certificate to operate on. Matches the `id` field in the `ClientCertificate` response schema.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the sm_hsm_client_certificate.
* `created_at` - (String) Computed field.
* `created_by` - (String) Computed field.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.
* `csr` - (String) PEM-encoded CSR. Platform-generated. Immutable once set. Populated on creation (always present once the resource exists).
  * Constraints: The maximum length is `65536` characters. The minimum length is `1` character. The value must match regular expression `/^-----BEGIN CERTIFICATE REQUEST-----[\\s\\S]+-----END CERTIFICATE REQUEST-----\\s*$/`.
* `csr_cn` - (String) Common Name embedded in the CSR. Set to the SM instance UUID at generation time. Immutable once set. Populated on creation (always present once the resource exists).
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `description` - (String) Optional description. Omitted if not set.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `expires_at` - (String) Parsed from signed_cert NotAfter field. Omitted until signed_cert is uploaded.
* `health_as_of` - (String) Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
* `health_state` - (String) AS OF health state of the client certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.- `evaluating`: A health check is in progress for this client certificate.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.
  * Constraints: Allowable values are: `unknown`, `evaluating`, `valid`, `expired`.
* `href` - (String) URL of the client certificate resource.
  * Constraints: The maximum length is `2048` characters.
* `name` - (String) Human-readable display name.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `provisioning_state` - (String) Lifecycle state of the client certificate:- `csr_generated`: CSR and private key generated successfully. `csr` and `csr_cn` are populated. Initial state on POST. Waiting for signed_cert to be uploaded via PATCH.- `inactive`: signed_cert has been uploaded but the certificate has not been activated yet.- `active`: This certificate is the active client certificate for the destination. Only one can be active at a time.
  * Constraints: Allowable values are: `csr_generated`, `inactive`, `active`.
* `signed_cert` - (String) CA-signed certificate (PEM). Omitted until uploaded via PATCH.
  * Constraints: The maximum length is `65536` characters. The minimum length is `1` character. The value must match regular expression `/^-----BEGIN CERTIFICATE-----[\\s\\S]+-----END CERTIFICATE-----\\s*$/`.
* `status` - (String) Request-only. Set to `active` in a PATCH to activate this certificate. Never returned; read the current state from `provisioning_state`.
  * Constraints: Length must be `6` characters. The value must match regular expression `/^active$/`.
* `updated_at` - (String) Computed field. Omitted until first update.
* `updated_by` - (String) Computed field. Omitted until first update.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_@\/:.]+$/`.

