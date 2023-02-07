---
layout: "ibm"
page_title: "IBM : ibm_sm_public_certificate_action_validate_manual_dns" (Beta)
description: |-
  Get information about PublicCertificateActionValidateManualDNS
subcategory: "IBM Cloud Secrets Manager API"
---

# ibm_sm_public_certificate_action_validate_manual_dns

Provides a read-only data source for PublicCertificateActionValidateManualDNS. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_public_certificate_action_validate_manual_dns" {
  instance_id   = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region        = "us-south"
  id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
  secret_action_prototype = {"action_type":"public_cert_action_validate_dns_challenge"}
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The ID of the secret.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `secret_action_prototype` - (Required, List) Specify the properties for your secret action.
Nested scheme for **secret_action_prototype**:
	* `action_type` - (Optional, String) The type of secret action.
	  * Constraints: Allowable values are: `public_cert_action_validate_dns_challenge`, `private_cert_action_revoke_certificate`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the PublicCertificateActionValidateManualDNS.
* `action_type` - (String) The type of secret action.
  * Constraints: Allowable values are: `public_cert_action_validate_dns_challenge`, `private_cert_action_revoke_certificate`.

