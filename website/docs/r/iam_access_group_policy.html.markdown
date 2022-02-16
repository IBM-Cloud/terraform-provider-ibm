---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group_policy"
description: |-
  Manages IBM IAM access group policy.
---

# ibm_iam_access_group_policy

Create, update, or delete an IAM policy for an IAM access group. For more information, about IBM access group policy, see [creating policies for account management service access](https://cloud.ibm.com/docs/account?topic=account-account-services#account-management-access).

## Example usage

### Access group policy for all Identity and Access enabled services 
The following example creates an IAM policy that grants members of the access group the IAM `Viewer` platform role to all IAM-enabled services. 

```terraform
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]
  
  resource_tags {
    name = "env"
    value = "dev"
  }
}

```

### Access group policy for all Identity and Access enabled services within a resource group
The following example creates an IAM policy that grants members of the access group the IAM `Operator` platform role and the `Writer` service access role to all IAM-enabled services within a resource group. 

```terraform
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

### Access group policy using service with region
The following example creates an IAM policy that grants members of the access group the IAM `Viewer` platform role to all service instances of cloudantnosqldb in us-south region

```terraform
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]

  resources {
    service = "cloudantnosqldb"
    region  = "us-south"
  }
}
```

### Access group policy using service_type with region

```terraform
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = ibm_iam_access_group.accgrp.id
  roles           = ["Viewer"]

  resources {
    service_type = "service"
    region = "us-south"
  }
}

```

### Access group policy using resource instance 
The following example creates an IAM policy that grants members of the access group the IAM `Viewer` and `Administrator` platform role, and the `Manager` service access role to a single service instance. 

```terraform
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

### Create a policy to all instances of an IBM Cloud service within a resource group
The following example creates an IAM policy that grants members of the access group the IAM `Viewer` platform role to all instances of IBM Cloud Kubernetes Service that are created within a specific resource group. 

```terraform
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

### Access group policy by using resource and resource type 

```terraform
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

### Access group policy by using attributes

```terraform
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

### Access Group Policy by using resource_attributes

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `access_group_id` - (Required, Forces new resource, String) The ID of the access group.
- `account_management` - (Optional, Bool) Gives access to all account management services if set to **true**. Default value **false**. If you set this option, do not specify `resources` at the same time. **Note** Conflicts with `resources` and `resource_attributes`.
- `roles` - (Required, List)  A comma separated list of roles. Valid roles are `Writer`, `Reader`, `Manager`, `Administrator`, `Operator`, `Viewer`, and `Editor`. For more information, about supported service specific roles, see  [IAM roles and actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions)
- `resources`  (List , Optional) A nested block describes the resource of this policy. **Note** Conflicts with `account_management` and `resource_attributes`.

  Nested scheme for `resources`:
  - `attributes` (Optional, Map) Set resource attributes in the form of `name=value,name=value`.  If you set this option, do not specify `account_management` at the same time.
  - `resource_instance_id` - (Optional, String) The ID of resource instance of the policy definition.
  - `region`  (Optional, String) The region of the policy definition.
  - `resource_type`  (Optional, String) The resource type of the policy definition.
  - `resource`  (Optional, String) The resource of the policy definition.
  - `resources.resource_group_id` - (Optional, String) The ID of the resource group. To retrieve the ID, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.
  - `service` - (Optional, String) The service name that you want to include in your policy definition. For account management services, you can find supported values in the [documentation](https://cloud.ibm.com/docs/account?topic=account-account-services#api-acct-mgmt). For other services, run the `ibmcloud catalog service-marketplace` command and retrieve the value from the **Name** column of your command line output. Attributes service, service_type are mutually exclusive.
  - `service_type`  (Optional, String) The service type of the policy definition. **Note** Attributes service, service_type are mutually exclusive.

- `resource_attributes` - (Optional, List) A nested block describing the resource of this policy. **Note** Conflicts with `account_management` and `resources`.

  Nested scheme for `resource_attributes`:
  - `name` - (Required, String) Name of an attribute. Supported values are `serviceName`, `serviceInstance`, `region`,`resourceType`, `resource`, `resourceGroupId`, and other service specific resource attributes.
  - `value` - (Required, String) Value of an attribute.
  - `operator` - (Optional, string) Operator of an attribute. Default value is `stringEquals`. **Note** Conflicts with `account_management` and `resources`.

- `resource_tags`  (Optional, List)  A nested block describing the access management tags.  **Note** `resource_tags` are only allowed in policy with resource attribute serviceType, where value is equal to service.
  
  Nested scheme for `resource_tags`:
  - `name` - (Required, String) The key of an access management tag. 
  - `value` - (Required, String) The value of an access management tag.
  - `operator` - (Optional, String) Operator of an attribute. The default value is `stringEquals`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.


- `id` - (String) The unique identifier of the access group policy. The ID is composed of `<access_group_id>/<access_group_policy_id>`.
- `version` - (String) The version of the access group policy.

## Import

The `ibm_iam_access_group_policy` resource can be imported by using access group ID and access group policy ID.

**Syntax**

```
$ terraform import ibm_iam_access_group_policy.example <access_group_ID>/<access_group_policy_ID>
```

**Example**

```
$ terraform import ibm_iam_access_group_policy.example AccessGroupId-1148204e-6ef2-4ce1-9fd2-05e82a390fcf/bf5d6807-371e-4755-a282-64ebf575b80a
```
