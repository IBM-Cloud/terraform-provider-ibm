---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_identity"
description: |-
  Get information about iam_trusted_profile_identity
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_identity

Provides a read-only data source for iam_trusted_profile_identity. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
  name = "test"
}

data "ibm_iam_trusted_profile_identity" "iam_trusted_profile_identity" {
	identifier_id = "IBMid-1234567898"
	identity_type = "user"
	profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `identifier_id` - (Required, String) Identifier of the identity that can assume the trusted profiles.
* `identity_type` - (Required, String) Type of the identity.
  * Constraints: Allowable values are: `user`, `serviceid`, `crn`.
* `profile_id` - (Required, String) ID of the trusted profile.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the iam_trusted_profile_identity.

* `accounts` - (List) Only valid for the type user. Accounts from which a user can assume the trusted profile.

* `description` - (String) Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.

* `iam_id` - (String) IAM ID of the identity.

* `identifier` - (String) Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.

* `type` - (String) Type of the identity.
  * Constraints: Allowable values are: `user`, `serviceid`, `crn`.

