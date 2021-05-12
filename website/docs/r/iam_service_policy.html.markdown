---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_policy"
description: |-
  Manages IBM IAM Service Policy.
---

# ibm\_iam_service_policy

Provides a resource for IAM Service Policy. This allows service policy  to be created, updated and deleted.

## Example Usage

### Service Policy for All Identity and Access enabled services 

```hcl
resource "ibm_iam_service_id" "serviceID" {
  name = "test"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Viewer"]
}

```

### Service Policy using service with region

```hcl
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
### Service Policy using resource instance 

```hcl
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

### Service Policy using resource group 

```hcl
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

### Service Policy using resource and resource type 

```hcl
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

### Service Policy using attributes 

```hcl
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
### Cross account Service policy using iam_id

```hcl
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

### Service Policy using resource_attributes

```hcl
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

## Argument Reference

The following arguments are supported:

* `iam_service_id` - (Optional, Forces new resource, string) UUID of the serviceID. Exactly one of `iam_service_id`, `iam_id` is required.
* `iam_id` - (Optional, Forces new resource, string) IAM ID of the serviceID. Exactly one of `iam_service_id`, `iam_id` is required. Can be used to assign cross account service ID Policy.
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
  **NOTE**: Conflicts with `resources`and `resource_attributes`.
* `tags` - (Optional, array of strings) Tags associated with the service policy instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the service policy. The id is composed of \<iam_service_id\>/\<service_policy_id\> if policy is created using <iam_service_id>. The id is composed of \<iam_id\>/\<service_policy_id\> if policy is created using <iam_id>. 

* `version` - Version of the service policy.

## Import

ibm_iam_service_policy can be imported using serviceID and service policy id or iamID and service policy id, eg

```
$ terraform import ibm_iam_service_policy.example ServiceId-d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
$ terraform import ibm_iam_service_policy.example iam-ServiceId-d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```

