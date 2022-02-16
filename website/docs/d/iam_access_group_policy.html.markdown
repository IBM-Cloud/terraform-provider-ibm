---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group_policy"
description: |-
  Manages IBM IAM access group policy.
---

# ibm_iam_access_group_policy

Retrieve information about an IAM access group policy. For more information, about IAM role action, see [managing access to resources](https://cloud.ibm.com/docs/account?topic=account-assign-access-resources).

## Example usage

```terraform
resource "ibm_iam_access_group" "accgrp" {
  name = "test123"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]
  
  resources {
    service = "cloud-object-storage"
  }
}

data "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group_policy.policy.access_group_id
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `access_group_id` - (Required, Forces new resource, String) The ID of the access group.
- `sort`- (Optional, String) The single field sort query for policies. Allowed values are `id`, `type`, `href`, `created_at`, `created_by_id`, `last_modified_at`,`last_modified_by_id`, `state`

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `policies` - (List) A nested block describes IAM Policies assigned to access group.

  Nested scheme for `policies`:
  - `description`  (String) The description of the IAM access group Policy.
  - `id` - (String) The unique identifier of the IAM access group policy. The ID is composed of `<ibm_id>/<access_group_policy_id>`.
  - `roles`-  (String) The roles that are assigned to the policy.
  - `resources`- (List of objects) A nested block describes the resources in the policy.

      Nested scheme for `resources`:
      - `service` - (String) The service name of the policy definition. 
      - `region` - (String) The region of the policy definition.
      - `resource_type` - (String) The resource type of the policy definition.
      - `resource` - (String) The resource of the policy definition.
      - `resource_group_id` - (String) The ID of the resource group.
      - `resource_instance_id`- (String) The ID of resource instance of the policy definition.
      - `service_type`- (String) The service type of the policy definition.

  - `resource_tags`- (List of objects) A nested block describes the access management tags in the policy.
    
    Nested scheme for `resource_tags`:
    - `name` - (String) The key of an access management tag. 
    - `value` - (String) The value of an access management tag.
    - `operator` - (String) Operator of an attribute.
