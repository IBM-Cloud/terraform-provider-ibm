---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_user_invite"
description: |-
  Manages IBM IAM user invite.
---

# ibm_iam_user_invite

Invite, update, or delete IAM users to your IBM Cloud account. User to be invited can be added to one or more access groups. For more information, see [inviting users](https://cloud.ibm.com/docs/account?topic=account-access-getstarted).

## Example usage

### Inviting batch of users

```terraform
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
}

```

### Inviting batch of users with access groups

```terraform
resource "ibm_iam_user_invite" "invite_user" {
  users         = ["test@in.ibm.com"]
  access_groups = ["accessgroup-id-9876543210"]
}

```

### Inviting batch of users with Classic Infrastructure roles
The following example provides the Classic Infrastructure permissions, and permission set.

#### Inviting batch of users with Classic Infrastructure permissions

```terraform
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  classic_infra_roles {
    permissions = ["PORT_CONTROL", "DATACENTER_ACCESS"]
  }
}

```

#### Inviting batch of users with Classic Infrastructure permission set

```terraform
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  classic_infra_roles {
    permission_set = "superuser"
  }
}

```


### User invite with access IAM policies
The following sample provides about user invite with access to IAM policies for all identity and access enables services, and user policy by using service with region and resourse instance, resource group, resource type, and attributes.

#### Inviting batch of users with user policy for All Identity and Access enabled services

```terraform
resource "ibm_iam_user_invite" "invite_user" {
  users = ["test@in.ibm.com"]
  iam_policy {
    roles = ["Viewer"]
  }
}

```

#### Inviting batch of users with user policy using service with region

```terraform
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

#### Inviting batch of users with user policy using resource instance

```terraform
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

#### Inviting batch of users with user policy using resource group

```terraform
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

#### Inviting batch of users with user policy using resource and resource type

```terraform
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

#### Inviting batch of users with user policy using attributes

```terraform
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

### User invite with access Cloud Foundry roles

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `account_management` - (Optional, Bool) Gives access to all account management services if set to **true**. Default value is **false**. If you set this option, do not set `resources` at the same time.
- `access_groups` (Optional, List) A comma separated list of access group IDs.
- `classic_infra_roles` (Optional, Map) A nested block describes the classic infrastructure roles for the inviting users. </br></br>**Note** If you have an IBM Cloud Lite account, you cannot set classic infrastructure roles. For more information, about Lite accounts, see [What's available?](https://cloud.ibm.com/docs/account?topic=account-accounts#lite-account-features).

  Nested scheme for `classic_infra_roles`:
  - `permissions`  (Optional, List) A comma separated list of classic infrastructure permissions.
  - `permission_set` - (Optional, String) The permission set to be applied. The valid permission sets are `noacess`, `viewonly`, `basicuser`, and `superuser`.
- `cloud_foundry_roles` -  (Optional, List) A nested block describes the cloud foundry roles of inviting user.

  Nested scheme for `cloud_foundry_roles`:
  - `organization_guid` - (Required, String) The ID of the Cloud Foundry organization.
  - `org_roles` - (Required, List) The organization roles that are assigned to invited user. The supported roles are `Manager`, `Auditor`, `BillingManager`.
  - `spaces`  (Optional, List) A nested block describes the Cloud Foundry space roles and space details.

    Nested scheme for `spaces`:
    - `space_guid` - (Required, String) The ID of the Cloud Foundry space.
    - `space_roles` - (Required, List) The space roles that you want to assign to the invited user. The supported space roles are `Manager`, `Developer`, `Auditor`.
- `iam_policy` (Optional, List) A nested block describes the IAM policies for invited users.

  Nested scheme for `iam_policy`:
  - `roles` - (Required, List) A comma separated list of roles. Valid roles are `Writer`, `Reader`, `Manager`, `Administrator`, `Operator`, `Viewer`, and `Editor`.
  - `resources` - (List of Objects) Optional- A nested block describes the resource of this policy. For more information, about supported service specific roles, see  [IAM roles and actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions)

    Nested scheme for `resources`:
    - `attributes` (Optional, Map)  A set of resource attributes in the format `name=value, name=value`. If you set this option, do not specify `account_management` at the same time.
    - `resource_instance_id` - (Optional, String) The ID of the resource instance of the policy definition.
    - `region`  (Optional, String) The region of the policy definition.
    - `resource_type` - (Optional, String) The resource type of the policy definition.
    - `resource` - (Optional, String) The resource of the policy definition.
    - `resource_group_id` - (Optional, String) The ID of the resource group. To retrieve the value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.
    - `service` - (Optional, String) The service name of the policy definition. You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `users` - (Required, List) A comma separated list of user Email IDs.
 
 **Note** 
 
 IBM Cloud `Lite account` does not support classic infrastructure roles. For more information, see [What's available in lite account?](https://cloud.ibm.com/docs/account?topic=account-accounts#lite-account-features).

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `access_groups` - (String) The lock down ID.

  Nested scheme for `access_group`:
  - `name` - (String) The name of the access group.
  - `policies` - (String) The access group policies of invited user.

    Nested scheme for `policies`:
    - `id` - (String) The policy ID.
    - `roles` - (String) The roles associated to the policy.
    - `resources` - (String)  A nested block describes the resource of the policy.

      Nested scheme for `resources`:
      - `attributes` - (String) The set of resource attributes.
      - `resource_instance_id` - (String) The resource instance ID of the policy definition.
      - `region` - (String) The region of the policy definition.
      - `resource_type` - (String) The resource type of the policy definition.
      - `resource` - (String) The resource of the policy definition.
      - `resource_group_id` - (String) The ID of the resource group.
      - `service` - (String)  Service name of the policy definition.
- `invited_users` - (String) List of invited users. 

  Nested scheme for `invited_users`:
  - `user_id` - (String) The Email ID of the member.
  - `user_policies` - (String)  List of policies associated to a particular user.

    Nested scheme for `user_policies`:
    - `id` - (String) The policy ID.
    - `roles` - (String) Comma separated list of the roles.
    - `resources` - (String)  A nested block describes the resource of the policy.

      Nested scheme for `resources`:
      - `attributes` - (String) The set of resource attributes.
      - `resource_instance_id` - (String) The resource instance ID of the policy definition.
      - `region` - (String) The region of the policy definition.
      - `resource_type` - (String) The resource type of the policy definition.
      - `resource` - (String) The resource of the policy definition.
      - `resource_group_id` - (String) The ID of the resource group.
      - `service` - (String)  Service name of the policy definition.
- `number_of_invited_users` - (String) Number of users invited to a particular account.

## Import
The import functionality is not supported for this resource.
