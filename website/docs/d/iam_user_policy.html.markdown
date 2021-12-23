---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_user_policy"
description: |-
  Manages IBM IAM user policy.
---

# ibm_iam_user_policy

Retrieve information about an IAM user policy. For more information, about IAM role action, see [managing access to resources](https://cloud.ibm.com/docs/account?topic=account-assign-access-resources).

## Example usage

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]

  resources {
    service = "kms"
    region  = "us-south"
  }
}

data "ibm_iam_user_policy" "testacc_ds_user_policy" {
  ibm_id = ibm_iam_user_policy.policy.ibm_id
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `ibm_id` - (Required, String) The IBM ID or email address of the user.
- `sort`- (Optional, String) The single field sort query for  policies.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `policies` - (List) A nested block describes IAM Policies assigned to user.

  Nested scheme for `policies`:
  - `description`  (String) The description of the IAM User Policy.
  - `id` - (String) The unique identifier of the IAM user policy. The ID is composed of `<ibm_id>/<user_policy_id>`.
  - `roles`-  (String) The roles that are assigned to the policy.
  - `resources`- (List of objects) A nested block describes the resources in the policy.

      Nested scheme for `resources`:
      -  `service` - (String) The service name of the policy definition. 
      - `region` - (String) The region of the policy definition.
      - `resource_type` - (String) The resource type of the policy definition.
      - `resource` - (String) The resource of the policy definition.
      - `resource_group_id` - (String) The ID of the resource group.
      - `resource_instance_id`- (String) The ID of resource instance of the policy definition.
    
  - `resource_tags`- (List of objects) A nested block describes the access management tags in the policy.
  
      Nested scheme for `resource_tags`:
        - `name` - (String) The key of an access management tag. 
        - `value` - (String) The value of an access management tag.
        - `operator` - (String) Operator of an attribute.
