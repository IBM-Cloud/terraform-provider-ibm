---
layout: "ibm"
page_title: "IBM : iam_user_invite"
sidebar_current: "docs-ibm-resource-iam-user-invite"
description: |-
  Manages IBM IAM User Invite.
---

# ibm\_iam_user_invite

Provides a resource for IAM User Invite. This allows batch of users or single user to be invited, updated and deleted. User to be invited can be added to one or more access groups

## Example Usage

### Inviting batch of Users

```hcl
resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
}

```

### Inviting batch of Users with access groups
```hcl
resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    access_groups = ["accessgroup-id-9876543210"]
}

```

### User invite with access IAM policies

#### Inviting batch of Users with User Policy for All Identity and Access enabled services
```hcl
resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    iam_policy =[{
      roles  = ["Viewer"]
    }]
}

```

#### Inviting batch of Users with User Policy using service with region
```hcl
resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    iam_policy =[{
      roles  = ["Viewer"]
      resources = [{
        service = "kms"
      }]
    }]
}

```

#### Inviting batch of Users with User Policy using resource instance
```hcl
resource "ibm_resource_instance" "instance" {
  name     = "test"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    iam_policy =[{
      roles  = ["Manager", "Viewer", "Administrator"]
      resources = [{
        service              = "kms"
        resource_instance_id = "${element(split(":",ibm_resource_instance.instance.id),7)}"
      }]
    }]
}

```

#### Inviting batch of Users with User Policy using resource group
```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    iam_policy =[{
      roles  = ["Manager", "Viewer", "Administrator"]
      resources = [{
        service           = "containers-kubernetes"
        resource_group_id = "${data.ibm_resource_group.group.id}"
      }]
    }]
}

```

#### Inviting batch of Users with User Policy using resource and resource type
```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    iam_policy =[{
      roles  = ["Manager", "Viewer", "Administrator"]
      resources = [{
        resource_type = "resource-group"
        resource      = "${data.ibm_resource_group.group.id}"
      }]
    }]
}

```

#### Inviting batch of Users with User Policy using attributes
```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    iam_policy =[{
      roles  = ["Manager", "Viewer", "Administrator"]
      resources = [{
        service = "is"
        attributes = {
          "vpcId" = "*"
        }
      }]
    }]
}

```
## Argument Reference

The following arguments are supported:

* `users` - (Required, list) comma separated list of users email-id. 
* `access_groups` - (Optional, list) comma seperated list of access group ids.
* `iam_policy` - (Optional, list) comma seperated list of IAM user policies
* `roles` - (Required, list) comma separated list of roles. Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor.
* `resources` - (Optional, list) A nested block describing the resource of this policy.
Nested `resources` blocks have the following structure:
  * `service` - (Optional, string) Service name of the policy definition.  You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
  * `resource_instance_id` - (Optional, string) ID of resource instance of the policy definition.
  * `region` - (Optional, string) Region of the policy definition.
  * `resource_type` - (Optional, string) Resource type of the policy definition.
  * `resource` - (Optional, string) Resource of the policy definition.
  * `resource_group_id` - (Optional, string) The ID of the resource group. You can retrieve the value from data source `ibm_resource_group`. 
  * `attributes` - (Optional, map) Set resource attributes in the form of `'name=value,name=value...`.

## Import

Import functionality not supported for this resource.