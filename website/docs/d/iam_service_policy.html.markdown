---
subcategory: "Identity & Access (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_policy"
description: |-
  Manages IBM IAM Service Policy.
---

# ibm\_iam_service_policy

Import the details of an IAM (Identity and Access Management) service policy on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
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
  policy_id = ibm_iam_service_policy.policy.id
}

```

## Argument Reference

The following arguments are supported:

* `policy_id` - (Required, string) The unique identifier of the IAM service policy. The id is composed of \<iam_service_id\>/\<service_policy_id\> if policy is created using <iam_service_id>. The id is composed of \<iam_id\>/\<service_policy_id\> if policy is created using <iam_id>. 

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of this service policy data source. It is equivalent to <service_policy_id\> 
* `roles` -  Roles assigned to the policy.
* `resources` -  A nested block describing the resources in the policy.
  * `service` - Service name of the policy definition. 
  * `resource_instance_id` - ID of resource instance of the policy definition.
  * `region` - Region of the policy definition.
  * `resource_type` - Resource type of the policy definition.
  * `resource` - Resource of the policy definition.
  * `resource_group_id` - The ID of the resource group.
* `version` - Version of the service policy.