---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_policy"
description: |-
  Manages IBM IAM service policy.
---

# ibm_iam_service_policy

Retrieve information about an IAM service policy. For more information, about IAM role action, see [managing access to resources](https://cloud.ibm.com/docs/account?topic=account-assign-access-resources).

## Example usage

```terraform
resource "ibm_iam_service_policy" "policy" {
  iam_service_id = "ServiceId-d7bec597-4726-451f-8a63-e62e6f19c32c"
  roles          = ["Manager", "Viewer", "Administrator"]

  resources {
    service              = "kms"
    region               = "us-south"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
}

data "ibm_iam_service_policy" "testacc_ds_service_policy" {
  iam_service_id = ibm_iam_service_policy.policy.iam_service_id
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `iam_service_id` - (Required, String) The UUID of the service ID.
- `iam_id` - (Optional, String) IAM ID of the service ID. One of the `iam_service_id` or `iam_id` is required argument. You can use to get cross account service ID policy.
- `sort`- Optional -  (String) The single field sort query for policies.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `policies` - (List) A nested block describes IAM service policies that are assigned to a service ID.

  Nested scheme for `policies`:
  - `description`  (String) The description of the IAM Service Policy.
  - `id` - (String) The unique identifier of the IAM service policy. The ID is composed of `<iam_service_id>/<service_policy_id>`. If policy is created by using <iam_service_id>. The ID is composed of `<iam_id>/<service_policy_id>` if policy is created by using <iam_id>.
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