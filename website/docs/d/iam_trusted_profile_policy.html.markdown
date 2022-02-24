---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_trusted_profile_policy"
description: |-
  Manages IBM IAM trusted profile policy.
---

# ibm_iam_trusted_profile_policy

Retrieve information about an IAM trusted profile policy. For more information, about IAM role action, see [managing access to resources](https://cloud.ibm.com/docs/account?topic=account-assign-access-resources).

## Example usage

```terraform
resource "ibm_iam_trusted_profile_policy" "policy" {
  profile_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.profile_id
  roles      = ["Manager", "Viewer", "Administrator"]

  resources {
    service              = "kms"
    region               = "us-south"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
}

data "ibm_iam_trusted_profile_policy" "policy" {
  profile_id = ibm_iam_trusted_profile_policy.policy.profile_id
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `profile_id` - (Required, String) The UUID of the trusted profile. Either `profile_id` or `iam_id` is required.
- `iam_id` - (Optional, String) IAM ID of the trusted profile. Either `profile_id` or `iam_id` is required.
- `sort`- Optional -  (String) The single field sort query for policies.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `policies` - (List) A nested block describes IAM trusted profile policies that are assigned to a profile ID.

  Nested scheme for `policies`:
  - `description`  (String) The description of the IAM trusted profile policy.
  - `id` - (String) The unique identifier of the IAM trusted profile policy. The ID is composed of `<profile_id>/<profile_policy_id>`. If policy is created by using <profile_id>. The ID is composed of `<iam_id>/<profile_policy_id>` if policy is created by using <iam_id>.
  - `roles`-  (String) The roles that are assigned to the policy.
  - `resources`- (List of objects) A nested block describes the resources in the policy.
  

    Nested scheme for `resources`:
      - `service`- (String) The service name of the policy definition.
      - `resource_instance_id`- (String) The ID of resource instance of the policy definition.
      - `region`-  (String) The region of the policy definition.
      - `resource_type`- (String) The resource type of the policy definition.
      - `resource`- (String) The resource of the policy definition.
      - `resource_group_id`- (String) The ID of the resource group.
 
  - `resource_tags`- (List of objects) A nested block describes the access management tags in the policy.

    Nested scheme for `resource_tags`:
      - `name` - (String) The key of an access management tag. 
      - `value` - (String) The value of an access management tag.
      - `operator` - (String) Operator of an attribute.