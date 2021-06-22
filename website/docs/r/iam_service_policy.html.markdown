---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_policy"
description: |-
  Manages IBM IAM service policy.
---

# ibm_iam_service_policy

Create, update, or delete an IAM service policy. For more information, about IAM role action, see [managing access to resources](https://cloud.ibm.com/docs/account?topic=account-assign-access-resources).

## Example usage

### Service policy for all Identity and Access enabled services 

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Viewer"]
}

```

### Service Policy using service with region

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Viewer"]

  resources {
    service = "cloud-object-storage"
  }
}

```
### Service policy by using resource instance 

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}

resource "ibm_resource_instance" "instance" {
  name     = "test"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Manager", "Viewer", "Administrator"]

  resources {
    service              = "kms"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
}

```

### Service policy by using resource group 

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Viewer"]

  resources {
    service           = "containers-kubernetes"
    resource_group_id = data.ibm_resource_group.group.id
  }
}

```

### Service policy by using resource and resource type 

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Administrator"]

  resources {
    resource_type = "resource-group"
    resource      = data.ibm_resource_group.group.id
  }
}

```

### Service policy by using attributes 

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Administrator"]

  resources {
    service = "is"

    attributes = {
      "vpcId" = "*"
    }
  }
}

```
### Cross account service policy by using `iam_id`

```terraform
provider "ibm" {
    alias             = "accA"
    ibmcloud_api_key  = "Account A Api Key"
}
resource "ibm_iam_service_id" "serviceID" {
  provider = ibm.accA
  name     = "test"
}

provider "ibm" {
    alias             = "accB"
    ibmcloud_api_key  = "Account B Api Key"
}
resource "ibm_iam_service_policy" "policy" {
  provider       =  ibm.accB
  iam_id         =  ibm_iam_service_id.serviceID.iam_id
  roles          =  ["Reader"]
  resources {
    service = "cloud-object-storage"
  }
}

```

### Service policy by using resource_attributes

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}
resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
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

- `account_management` - (Optional, Bool) Gives access to all account management services if set to **true**. Default value is **false**. If you set this option, do not set `resources` at the same time.
- `iam_service_id` - (Required, Forces new resource, String) The UUID of the service ID.
- `iam_id` - (Optional,  Forces new resource, String) IAM ID of the service ID. Used to assign cross account service ID policy. Either `iam_service_id` or `iam_id` is required.
- `resources` - (List of Objects) Optional- A nested block describes the resource of this policy.

  Nested scheme for `resources`:
  - `service`  (Optional, String) The service name of the policy definition. You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search`.
  - `resource_instance_id` - (Optional, String) The ID of the resource instance of the policy definition.
  - `region` - (Optional, String) The region of the policy definition.
  - `resource_type` - (Optional, String) The resource type of the policy definition.
  - `resource` - (Optional, String) The resource of the policy definition.
  - `resource_group_id` - (Optional, String) The ID of the resource group. To retrieve the value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.
  - `attributes` (Optional, Map)  A set of resource attributes in the format `name=value,name=value`. If you set this option, do not specify `account_management` and `resource_attributes` at the same time.
- `resource_attributes` - (Optional, list) A nested block describing the resource of this policy.

  Nested scheme for `resource_attributes`:
  - `name` - (Required, String) The name of an attribute. Supported values are `serviceName` , `serviceInstance` , `region` ,`resourceType` , `resource` , `resourceGroupId` and other service specific resource attributes.
  - `value` - (Required, String) The value of an attribute.
  - `operator` - (Optional, String) Operator of an attribute. The default value is `stringEquals`. **Note** Conflicts with `account_management` and `resources`.
- `roles` - (Required, List) A comma separated list of roles. Valid roles are `Writer`, `Reader`, `Manager`, `Administrator`, `Operator`, `Viewer`, and `Editor`. For more information, about supported service specific roles, see  [IAM roles and actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions)
- `tags`  - (Optional, List of Strings) A list of tags with the service policy instance. **Note** Tags are managed locally and not stored in the IBM Cloud service endpoint at this moment.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`  - (String) The unique identifier of the service policy. The ID is composed of `<iam_service_id>/<service_policy_id>`. If policy is created by using `<iam_service_id>`. The ID is composed of `<iam_id>/<service_policy_id>` if policy is created by using `<iam_id>`.
- `version`  - (String) The version of the service policy.

## Import

The  `ibm_iam_service_policy` resource can be imported by using service ID and service policy ID or IAM ID and service policy ID.

**Syntax**

```
$ terraform import ibm_iam_service_policy.example <service_ID>/<service_policy_ID>
```

**Example**

```
$ terraform import ibm_iam_service_policy.example ServiceId-d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb

```

