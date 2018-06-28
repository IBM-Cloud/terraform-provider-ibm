---
layout: "ibm"
page_title: "IBM : iam_user_policy"
sidebar_current: "docs-ibm-resource-iam-user-policy"
description: |-
  Manages IBM IAM User Policy.
---

# ibm\_iam_user_policy

Provides a resource for IAM User Policy. This allows user policy to be created, updated and deleted. To assign a policy to one user, the user must exist in the account to which you assign the policy. 

## Example Usage

### User Policy for All Identity and Access enabled services 

```hcl
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]
}

```

### User Policy using service with region

```hcl
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]

  resources = [{
    service = "kms"
    region  = "us-south"
  }]
}

```
### User Policy using resource instance 

```hcl
resource "ibm_resource_instance" "instance" {
  name     = "test"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Manager", "Viewer", "Administrator"]

  resources = [{
    service              = "kms"
    region               = "us-south"
    resource_instance_id = "${element(split(":",ibm_resource_instance.instance.id),7)}"
  }]
}

```

### User Policy using resource group 

```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]

  resources = [{
    service           = "containers-kubernetes"
    resource_group_id = "${data.ibm_resource_group.group.id}"
  }]
}

```

### User Policy using resource and resource type 

```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Administrator"]

  resources = [{
    resource_type = "resource-group"
    resource      = "${data.ibm_resource_group.group.id}"
  }]
}

```

## Argument Reference

The following arguments are supported:

* `ibm_id` - (Required, string) The ibm id or email of user.
* `roles` - (Required, list) comma separated list of roles. Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor.
* `resources` - (Optional, list) A nested block describing the resource of this policy.
Nested `resources` blocks have the following structure:
  * `service` - (Optional, string) Service name of the policy definition.  You can retrieve the value by running the `bx catalog service-marketplace` or `bx catalog search` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
  * `resource_instance_id` - (Optional, string) ID of resource instance of the policy definition.
  * `region` - (Optional, string) Region of the policy definition.
  * `resource_type` - (Optional, string) Resource type of the policy definition.
  * `resource` - (Optional, string) Resource of the policy definition.
  * `resource_group_id` - (Optional, string) The ID of the resource group. You can retrieve the value from data source `ibm_resource_group`.
* `tags` - (Optional, array of strings) Tags associated with the user policy instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the User Policy. The id is composed of \<ibm_id\>/\<user_policy_id\>

* `version` - Version of the User Policy.

## Import

ibm_iam_user_policy can be imported using IBMID and User Policy id, eg

```
$ terraform import ibm_iam_user_policy.example test@in.ibm.com/9ebf7018-3d0c-4965-9976-ef8e0c38a7e2
```