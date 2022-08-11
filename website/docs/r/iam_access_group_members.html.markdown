---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group_members"
description: |-
  Manages IBM IAM access group members.
---

# ibm_iam_access_group_members


~> **WARNING:** Multiple `ibm_iam_access_group_members` resources with the same group name produce inconsistent behavior!

Add, update, or remove users from an IAM access group members. For more information, about IAM access group members, see [managing public access to resources](https://cloud.ibm.com/docs/account?topic=account-public).

## Example usage
The following example creates an IAM access group, a service ID and a trusted profile ID. Then, the service ID, profile ID and a user with the ID `user@ibm.com` is added to the access group.

```terraform
resource "ibm_iam_access_group" "accgroup" {
  name = "testgroup"
}

resource "ibm_iam_service_id" "serviceID" {
  name = "testserviceid"
}

resource "ibm_iam_trusted_profile" "profileID" {
  name = "testprofileid"
}

resource "ibm_iam_access_group_members" "accgroupmem" {
  access_group_id = ibm_iam_access_group.accgroup.id
  ibm_ids         = ["test@in.ibm.com"]
  iam_service_ids = [ibm_iam_service_id.serviceID.id]
  iam_profile_ids = [ibm_iam_trusted_profile.profileID.id]
}

```

## Argument reference

Review the argument references that you can specify for your resource. 

- `access_group_id` - (Required, String) The ID of the access group. 
- `ibm_ids` - (Optional, Array of string)  A list of IBM IDs that you want to add to or remove from the access group. 
- `iam_service_ids` - (Optional, Array of string)  A list of service IDS that you want to add to or remove from the access group.
- `iam_profile_ids` - (Optional, Array of string)  A list of trusted profile IDS that you want to add to or remove from the access group.
  

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created. 

- `id` - (String) The unique identifier of the access group members. The ID is returned in the format `<iam_access_group_ID>/<random_ID>`. 
- `members` - (Array of objects) A list of members that are included in the access group.

  Nested scheme for `members`:
	- `iam_id` - (String) The IBM ID or service ID or profile ID of the member.
	- `type` - (String) The type of member. Supported values are `user` or `service` or `profile`.


## Import

The `ibm_iam_access_group_members` can be imported by using access group ID and random ID. 

**Syntax**

```
$ terraform import ibm_iam_access_group_members.example <accessgroupID>/<random_ID>
```

**Example**

```
$ terraform import ibm_iam_access_group_members.example AccessGroupId-5391772e-1207-45e8-b032-2a21941c11ab/2018-10-04 06:27:40.041599641 +0000 UTC
```