---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_identities"
description: |-
  Get information about iam_trusted_profile_identities
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_identities

Provides a read-only data source to retrieve information about iam_trusted_profile_identities. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities" {
	profile_id = ibm_iam_trusted_profile_identities.iam_trusted_profile_identities_instance.profile_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_trusted_profile_identities.
* `entity_tag` - (String) Entity tag of the profile identities response.
* `identities` - (List) List of identities.
Nested schema for **identities**:
	* `accounts` - (List) Only valid for the type user. Accounts from which a user can assume the trusted profile.
	* `description` - (String) Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.
	* `iam_id` - (String) IAM ID of the identity.
	* `identifier` - (String) Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.
	* `type` - (String) Type of the identity.
	  * Constraints: Allowable values are: `user`, `serviceid`, `crn`.

