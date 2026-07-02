---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_identity"
description: |-
  Manages iam_trusted_profile_identity.
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_identity

Create, update, and delete iam_trusted_profile_identity with this resource.

## Example Usage

```hcl
resource "ibm_iam_trusted_profile_identity" "iam_trusted_profile_identity_instance" {
  identifier = "identifier"
  identity_type = "user"
  profile_id = "profile_id"
  type = "user"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accounts` - (Optional, Forces new resource, List) Only valid for the type user. Accounts from which a user can assume the trusted profile.
* `description` - (Optional, Forces new resource, String) Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.
* `identifier` - (Required, Forces new resource, String) Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.
* `identity_type` - (Required, Forces new resource, String) Type of the identity.
  * Constraints: Allowable values are: `user`, `serviceid`, `crn`.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.
* `type` - (Required, Forces new resource, String) Type of the identity.
  * Constraints: Allowable values are: `user`, `serviceid`, `crn`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_trusted_profile_identity. Id is a combination of `profile_id`|`identity-type`|`identifier-id`.


## Import

You can import the `ibm_iam_trusted_profile_identity` resource by using `id`.
The `id` property can be formed from `profile_id`, `identity-type`, and `identifier` in the following format:

<pre>
&lt;profile_id&gt;|&lt;identity-type&gt;|&lt;identifier&gt;
</pre>
* `profile-id`: A string. ID of the trusted profile.
* `identity-type`: A string. Type of the identity.
* `identifier-id`: A string. Identifier of the identity that can assume the trusted profiles.

# Syntax
<pre>
$ terraform import ibm_iam_trusted_profile_identity.iam_trusted_profile_identity &lt;profile_id&gt;|&lt;identity-type&gt;|&lt;identifier&gt;
</pre>
