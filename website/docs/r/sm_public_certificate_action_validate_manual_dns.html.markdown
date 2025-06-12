---
layout: "ibm"
page_title: "IBM : ibm_sm_public_certificate_action_validate_manual_dns"
description: |-
  Manages PublicCertificateActionValidateManualDns.
subcategory: "Secrets Manager"
---

# ibm_sm_public_certificate_action_validate_manual_dns

Provides a resource for PublicCertificateActionValidateManualDns. This allows PublicCertificateActionValidateManualDns to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_public_certificate_action_validate_manual_dns" "validate_manual_dns_action" {
  instance_id           = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region                = "us-south"
  secret_id             = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Required, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `secret_id` - (Required, String) The ID of the secret.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
  
