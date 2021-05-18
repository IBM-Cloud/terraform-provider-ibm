---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group_policy"
description: |-
  Manages IBM IAM Access Group Policy.
---

# ibm\_access_group_policy

Provides a resource for IAM Access Group Policy. This allows access group policy to be created, updated and deleted.

## Example Usage

### Access Group Policy for All Identity and Access enabled services 

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]
}

```

### Access Group Policy for All Identity and Access enabled services within a resource group

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Operator", "Writer"]

  resources {
    resource_group_id = data.ibm_resource_group.group.id
  }
}
```

### Access Group Policy using service with region

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]

  resources {
    service = "cloud-object-storage"
  }
}

```
### Access Group Policy using resource instance 

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_resource_instance" "instance" {
  name     = "test"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Manager", "Viewer", "Administrator"]

  resources {
    service              = "kms"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
}

```

### Access Group Policy using resource group 

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]

  resources {
    service           = "containers-kubernetes"
    resource_group_id = data.ibm_resource_group.group.id
  }
}


```

### Access Group Policy using resource and resource type 

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Administrator"]

  resources {
    resource_type = "resource-group"
    resource      = data.ibm_resource_group.group.id
  }
}

```

### Access Group Policy using attributes

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]

  resources {
    service = "is"

    attributes = {
      "vpcId" = "*"
    }

    resource_group_id = data.ibm_resource_group.group.id
  }
}

```

### Access Group Policy using resource_attributes

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "access_group"
}
resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
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

* `access_group_id` - (Required, Forces new resource, string) ID of the access group.
* `roles` - (Required, list) comma separated list of roles. Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor.
* `resources` - (Optional, list) A nested block describing the resource of this policy.
Nested `resources` blocks have the following structure:
  * `service` - (Optional, string) Service name of the policy definition.  You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
  * `resource_instance_id` - (Optional, string) ID of resource instance of the policy definition.
  * `region` - (Optional, string) Region of the policy definition.
  * `resource_type` - (Optional, string) Resource type of the policy definition.
  * `resource` - (Optional, string) Resource of the policy definition.
  * `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. 
  * `attributes` - (Optional, map) Set resource attributes in the form of `'name=value,name=value...`.
 **NOTE**: Conflicts with `account_management` and `resource_attributes`.
* `resource_attributes` - (Optional, list) A nested block describing the resource of this policy.
Nested `resource_attributes` blocks have the following structure:
  * `name` - (Required, string) Name of the Attribute. Supported values are`serviceName` , `serviceInstance` , `region` ,`resourceType` , `resource` , `resourceGroupId` and other service specific resource attributes.
  * `value` - (Required, string) Value of the Attribute.
  * `operator` - (Optional, string) Operator of the Attribute. Default Value: `stringEquals`
 **NOTE**: Conflicts with `account_management` and `resources`.
* `account_management` - (Optional, bool) Gives access to all account management services if set to `true`. Default value `false`. 
 **NOTE**: **NOTE**: Conflicts with `resources` and `resource_attributes`.
* `tags` - (Optional, array of strings) Tags associated with the access group Policy instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the access group policy. The id is composed of \<access_group_id\>/\<access_group_policy_id\>

* `version` - Version of the access group policy.

## Import

ibm_iam_access_group_policy can be imported using access group ID and access group policy ID, eg

```
$ terraform import ibm_iam_access_group_policy.example AccessGroupId-1148204e-6ef2-4ce1-9fd2-05e82a390fcf/bf5d6807-371e-4755-a282-64ebf575b80a
```