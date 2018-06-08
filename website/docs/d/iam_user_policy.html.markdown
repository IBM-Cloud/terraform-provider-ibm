---
layout: "ibm"
page_title: "IBM : iam_user_policy"
sidebar_current: "docs-ibm-datasource-iam-user-policy"
description: |-
  Manages IBM IAM User Policy.
---

# ibm\_iam_user_policy

Import the details of an IAM (Identity and Access Management) user policy on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]

  resources = [{
    service = "kms"
    region  = "us-south"
  }]
}

data "ibm_iam_user_policy" "testacc_ds_user_policy" {
  ibm_id = "${ibm_iam_user_policy.policy.ibm_id}"
}

```

## Argument Reference

The following arguments are supported:

* `ibm_id` - (Required, string) The ibm id or email of user.

## Attribute Reference

The following attributes are exported:

* `policies` - A nested block describing IAM Policies assigned to user. Nested `policies` blocks have the following structure:
  * `id` - The unique identifier of the IAM user policy.The id is composed of \<ibm_id\>/\<user_policy_id\>
  * `roles` -  Roles assigned to the policy.
	* `resources` -  A nested block describing the resources in the policy.
		* `service` - Service name of the policy definition. 
		* `resource_instance_id` - ID of resource instance of the policy definition.
		* `region` - Region of the policy definition.
		* `resource_type` - Resource type of the policy definition.
		* `resource` - Resource of the policy definition.
		* `resource_group_id` - The ID of the resource group. 


  