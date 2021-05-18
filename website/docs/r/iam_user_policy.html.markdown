---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_user_policy"
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

  resources {
    service = "kms"
  }
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

  resources {
    service              = "kms"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
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

  resources {
    service           = "containers-kubernetes"
    resource_group_id = data.ibm_resource_group.group.id
  }
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

  resources {
    resource_type = "resource-group"
    resource      = data.ibm_resource_group.group.id
  }
}

```

### User Policy using attributes 

```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Administrator"]

  resources {
    service = "is"

    attributes = {
      "vpcId" = "*"
    }
  }
}

```

### User Policy using resource_attributes

```hcl
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles           = ["Viewer"]
  resource_attributes {
    name  = "resource"
    value = "test123*"
    operator = "stringMatch"
  }
  resource_attributes {
    name  = "serviceName"
    value = "messagehub"
  }
}
```


## Argument Reference

The following arguments are supported:

* `ibm_id` - (Required, Forces new resource, string) The ibm id or email of user.
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
 **NOTE**: Conflicts with `account_management` and `resource_attributes`.
* `resource_attributes` - (Optional, list) A nested block describing the resource of this policy.
Nested `resource_attributes` blocks have the following structure:
  * `name` - (Required, string) Name of the Attribute. Supported values are`serviceName` , `serviceInstance` , `region` ,`resourceType` , `resource` , `resourceGroupId` and other service specific resource attributes.
  * `value` - (Required, string) Value of the Attribute.
  * `operator` - (Optional, string) Operator of the Attribute. Default Value: `stringEquals`
 **NOTE**: Conflicts with `account_management` and `resources`.
* `account_management` - (Optional, bool) Gives access to all account management services if set to `true`. Default value `false`. 
  **NOTE**: Conflicts with `resources`and `resource_attributes`.
* `tags` - (Optional, array of strings) Tags associated with the user policy instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the User Policy. The id is composed of \<ibm_id\>/\<user_policy_id\>

* `version` - Version of the User Policy.

## Import

ibm_iam_user_policy can be imported using IBMID and User Policy id, eg

```
$ terraform import ibm_iam_user_policy.example test@in.ibm.com/9ebf7018-3d0c-4965-9976-ef8e0c38a7e2
```