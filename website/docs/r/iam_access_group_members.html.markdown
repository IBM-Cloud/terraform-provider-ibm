---
layout: "ibm"
page_title: "IBM : iam_access_group_members"
sidebar_current: "docs-ibm-resource-iam-access-group-members"
description: |-
  Manages IBM IAM Access Group Members.
---

# ibm\_iam_access_group_members


~> **WARNING:** Multiple ibm_iam_access_group_members resources with the same group name will produce inconsistent behavior!

Provides a resource for IAM access group members. This allows access group members to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_iam_access_group" "accgroup" {
  name = "testgroup"
}

resource "ibm_iam_service_id" "serviceID" {
  name = "testserviceid"
}

resource "ibm_iam_access_group_members" "accgroupmem" {
  access_group_id = "${ibm_iam_access_group.accgroup.id}"
  ibm_ids         = ["test@in.ibm.com"]
  iam_service_ids = ["${ibm_iam_service_id.serviceID.id}"]
}

```

## Argument Reference

The following arguments are supported:

* `access_group_id` - (Required, string) ID of the access group.
* `ibm_ids` - (Optional, array of strings) List of IBMid's.
* `iam_service_ids` - (Optional, array of strings) List of serviceID's.  
  

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the access group members. The id is composed of \<iam_access_group_id\>/\<randomid\>
* `members` - List of members attached to the access group.
Nested `members` blocks have the following structure:
  * `iam_id` - The IBMid or Service Id of the member
  * `type` - The type of the member, either "user" or "service".

## Import

ibm_iam_access_group_members can be imported using access group ID and random id, eg

```
$ terraform import ibm_iam_access_group_members.example AccessGroupId-5391772e-1207-45e8-b032-2a21941c11ab/2018-10-04 06:27:40.041599641 +0000 UTC
```