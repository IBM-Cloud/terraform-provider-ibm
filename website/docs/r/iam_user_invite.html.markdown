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
  users         = ["test@in.ibm.com"]
  access_groups = ["accessgroup-id-9876543210"]
}

```

### Inviting batch of Users with Classic Infrastructure roles

#### Inviting batch of Users with Classic Infrastructure permissions

```hcl
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  classic_infra_roles {
    permissions = ["PORT_CONTROL", "DATACENTER_ACCESS"]
  }
}

```

#### Inviting batch of Users with Classic Infrastructure permission set

```hcl
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  classic_infra_roles {
    permission_set = "superuser"
  }
}

```


### User invite with access IAM policies

#### Inviting batch of Users with User Policy for All Identity and Access enabled services

```hcl
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  iam_policy {
    roles = ["Viewer"]
  }
}

```

#### Inviting batch of Users with User Policy using service with region

```hcl
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  iam_policy {
    roles = ["Viewer"]
    resources {
      service = "kms"
    }
  }
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
  iam_policy {
    roles = ["Manager", "Viewer", "Administrator"]
    resources {
      service              = "kms"
      resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
    }
  }
}

```

#### Inviting batch of Users with User Policy using resource group

```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  iam_policy {
    roles = ["Manager", "Viewer", "Administrator"]
    resources {
      service           = "containers-kubernetes"
      resource_group_id = data.ibm_resource_group.group.id
    }
  }
}

```

#### Inviting batch of Users with User Policy using resource and resource type

```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  iam_policy {
    roles = ["Manager", "Viewer", "Administrator"]
    resources {
      resource_type = "resource-group"
      resource      = data.ibm_resource_group.group.id
    }
  }
}

```

#### Inviting batch of Users with User Policy using attributes

```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  iam_policy {
    roles = ["Manager", "Viewer", "Administrator"]
    resources {
      service = "is"
      attributes = {
        "vpcId" = "*"
      }
    }
  }
}

```

### User invite with access cloud foundry roles

```
provider "ibm" {
}

data "ibm_org" "org" {
  org = var.org
}

data "ibm_space" "space" {
  org   = var.org
  space = var.space
}

resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  cloud_foundry_roles {
    organization_guid = data.ibm_org.org.id
    org_roles         = ["Manager", "Auditor"]
    spaces {
      space_guid  = data.ibm_space.space.id
      space_roles = ["Manager", "Developer"]
    }
  }
}

```


## Argument Reference

The following arguments are supported:

* `users` - (Required, list) comma separated list of users email-id. 
* `access_groups` - (Optional, list) comma seperated list of access group ids.
* `classic_infra_roles` - (Optional, map) A nested block describing the classic infrastrucre roles for the inviting users. The nested classic_infra_roles block have the following structure:
  * `permissions` - (Optional, list) comma seperated list of classic infrastructure permissions. You can obtain the supported permissions from [Permissions List](https://sldn.softlayer.com/article/permission-enforcement-softlayer-api)
  * `permission_set` - (Optional, string) Permission set to be applied. The valid permission sets are noacess, viewonly, basicuser, superuser
* `iam_policy` - (Optional, list) A nested block describing the IAM Polocies for inviting users. The nested iam_policy block have the following structure:
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
  * `account_management` - (Optional, bool) Gives access to all account management services if set to `true`. Default value `false`.
* `cloud_foundry_roles` - (Optional, list) A nested block describing the cloud foundry roles of inviting user. The nested cloud_foundry_roles block have the following structure:
  * `organization_guid` - (Required, string) ID of the cloud foundry organization.
  * `org_roles` - (Required, list) The orgnization roles assigned for the inviting user. The supported org_roles are Manager, Auditor, BillingManager.
  * `spaces` - (Optional, list) A nested block describing the cloud foundry space roles and space details. The nested spaces block have the following structure:
    * `space_guid` - (Required, string) ID of the cloud foundry space.
    * `space_roles` - (Required, list) The space roles assigned for the inviting user. The supported space roles are Manager, Developer, Auditor.

**NOTE**: ibmcloud `Lite account` does not support classic infrastructure roles. For more info refer [whats available in lite account?](https://cloud.ibm.com/docs/account?topic=account-accounts#lite-account-features).

## Attribute Reference

The following attributes are exported:
* `number_of_invited_users` - Number of users invited to a particular account
* `invited_users` - List of invited users.
Nested `invited_users` block have the following structure:
  * `user_id` - Email Id of the member
  * `user_policies` - List of policies associated to a particular user.
  Nested `user_policies` block have the following structure:
    * `id` - Policy ID
    * `roles` - comma separated list of roles
    * `resources` - A nested block describes the resource of this policy.
    Nested `resources` block have the following structure: 
      * `service` - service name of the policy definition.
      * `resource_instance_id` - ID of the resource instance of the policy definition.
      * `region` - region of the policy definition.
      * `resource_type` - resource type of the policy definition.
      * `resource` - resource of the policy definition.
      * `resource_group_id` - ID of the resource group.
      * `attributes` - set of resource attributes
* `access_groups` - List of access groups
Nested `access_groups` block have the following structure: 
  * `name` - Name of the access group
  * `policies` - access group policies of invited user
  Nested `policies` block have the following structure: 
    * `id` - policy ID
    * `roles` - roles associted to policy
    * `resources` - A nested block describes the resource of this policy.
    Nested `resources` block have the following structure: 
      * `service` - service name of the policy definition.
      * `resource_instance_id` - ID of the resource instance of the policy definition.
      * `region` - region of the policy definition.
      * `resource_type` - resource type of the policy definition.
      * `resource` - resource of the policy definition.
      * `resource_group_id` - ID of the resource group.
      * `attributes` - set of resource attributes


## Import

Import functionality not supported for this resource.