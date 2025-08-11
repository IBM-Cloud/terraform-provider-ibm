---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_identities"
description: |-
  Manages iam_trusted_profile_identities.
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_identities

Create, update, and delete iam_trusted_profile_identitiess with this resource.

## Example Usage

```hcl
resource "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_instance" {
  identities {
		iam_id = "iam_id"
		identifier = "identifier"
		type = "user"
		accounts = [ "accounts" ]
		description = "description"
  }
  profile_id = "profile_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `identities` - (Optional, List) List of identities.
Nested schema for **identities**:
	* `accounts` - (Optional, List) Only valid for the type user. Accounts from which a user can assume the trusted profile.
	* `description` - (Optional, String) Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.
	* `iam_id` - (Required, String) IAM ID of the identity.
	* `identifier` - (Required, String) Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.
	* `type` - (Required, String) Type of the identity.
	  * Constraints: Allowable values are: `user`, `serviceid`, `crn`.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_trusted_profile_identities.


## Import

You can import the `ibm_iam_trusted_profile_identities` resource by using `profile_id`. Profile id of the profile identities response.

# Syntax
<pre>
$ terraform import ibm_iam_trusted_profile_identities.iam_trusted_profile_identities &lt;profile-id&gt;
</pre>
