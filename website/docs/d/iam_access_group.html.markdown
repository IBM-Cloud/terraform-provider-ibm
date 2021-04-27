---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group"
description: |-
  Get information about IBM IAM Access Group and all the members and dynamic rules associated with the group.
---

# ibm\_iam_access_group

Import the details of an existing [IAM Access Group](https://cloud.ibm.com/iam/groups) as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_iam_access_group" "accgroup" {
  access_group_name = ibm_iam_access_group.accgroup.name
}
```

## Argument Reference

The following arguments are supported:

* `access_group_name` - (Optional, string) The name of the Access Group. If not specified all access group present in the account will be fetched.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `groups` - List of Access Groups attached to the Account.
Nested `groups` blocks have the following structure:
  * `id` - The id of the access group
  * `name` - Name of the access group.
  * `description` - Description of the access group.
  * `ibm_ids` - List of IBMid of the member
  * `iam_service_ids` - List of Service Id of the member.
  * `rules` - List of Access Groups attached to the Account.
  Nested `rules` blocks have the following structure:
    * `name` -  Name of the dynamic rule.
    * `expiration` -  The number of hours that the rule lives for (Must be between 1 and 24).
    * `identity_provider` - (Required, string) The url of the identity provider.  
    * `conditions` -  A list of conditions the rule must satisfy:
      * `claim` - The claim to evaluate against. This will be found in the ext claims of a user's login request. 
      * `operator` -  The operation to perform on the claim. Valid operators are EQUALS, EQUALS_IGNORE_CASE, IN, NOT_EQUALS_IGNORE_CASE, NOT_EQUALS, and CONTAINS.
      * `value` - The stringified JSON value that the claim is compared to using the operator.
    * `rule_id` -  ID of the dynamic rule.


    