---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_authorization_policy"
description: |-
  Manages IBM IAM service authorizations.
---

# ibm_iam_authorization_policy

Create or delete an IAM service authorization policy. For more information, about IAM service authorizations, see [using authorizations to grant access between services](https://cloud.ibm.com/docs/account?topic=account-serviceauth).

## Example usage

### Authorization policy between two services

```terraform
resource "ibm_iam_authorization_policy" "policy" {
  source_service_name = "cloud-object-storage"
  target_service_name = "kms"
  roles               = ["Reader"]
  description         = "Authorization Policy"
}

```

### Authorization policy between two services with authorize dependent services enabled

```terraform
resource "ibm_iam_authorization_policy" "policy" {
  source_service_name         = "databases-for-postgresql"
  target_service_name         = "kms"
  roles                       = ["Reader", "Authorization Delegator"]
}
```

### Authorization policy between two services with specific resource type

```terraform
resource "ibm_iam_authorization_policy" "policy" {
  source_service_name  = "is"
  source_resource_type = "image"
  target_service_name  = "cloud-object-storage"
  roles                = ["Reader"]
}

```
### Authorization policy between two specific instances

```terraform
resource "ibm_resource_instance" "instance1" {
  name     = "mycos"
  service  = "cloud-object-storage"
  plan     = "lite"
  location = "global"
}

resource "ibm_resource_instance" "instance2" {
  name     = "mykms"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_authorization_policy" "policy" {
  source_service_name         = "cloud-object-storage"
  source_resource_instance_id = ibm_resource_instance.instance1.id
  target_service_name         = "kms"
  target_resource_instance_id = ibm_resource_instance.instance2.id
  roles                       = ["Reader"]
}

```
### Authorization policy between two specific resource group

```terraform
resource "ibm_resource_group" "source_resource_group" {
  name     = "123123"
}
	  
resource "ibm_resource_group" "target_resource_group" {
  name     = "456456"
}

resource "ibm_iam_authorization_policy" "policy" {
  source_service_name         = "cloud-object-storage"
  source_resource_group_id    = ibm_resource_group.source_resource_group.id
  target_service_name         = "kms"
  target_resource_group_id    = ibm_resource_group.target_resource_group.id
  roles                       = ["Reader"]
}

```

## Argument reference
Review the argument references that you can specify for your resource.

- `description`  (Optional, String) The description of the Authorization Policy.
- `roles` - (Required, list) The comma separated list of roles. For more information, about supported service specific roles, see  [IAM roles and actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions)
- `source_service_account` - (Optional, Forces new resource, string) The account GUID of source service.
- `source_service_name` - (Required, Forces new resource, string) The source service name.
- `target_service_name` - (Required, Forces new resource, string) The target service name.
- `source_resource_instance_id` - (Optional, Forces new resource, string) The source resource instance id.
- `target_resource_instance_id` - (Optional, Forces new resource, string) The target resource instance id.
- `source_resource_type` - (Optional, Forces new resource, string) The resource type of source service.
- `target_resource_type` - (Optional, Forces new resource, string) The resource type of target service.
- `source_resource_group_id` - (Optional, Forces new resource, string) The source resource group id.
- `target_resource_group_id` - (Optional, Forces new resource, string) The target resource group id.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the authorization policy.
- `version` - (String) The version of the authorization policy.

## Import

The `ibm_iam_authorization_policy` resource can be imported by using authorization policy ID.

**Syntax**

```
$ terraform import ibm_iam_authorization_policy.example <authorization policy ID>
```

**Example**

```
$ terraform import ibm_iam_authorization_policy.example 12fe9d62-81b1-41ee-8233-53150e38a61c
```
