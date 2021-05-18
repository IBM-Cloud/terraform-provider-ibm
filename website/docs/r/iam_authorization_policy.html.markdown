---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_authorization_policy"
description: |-
  Manages IBM IAM Service Authorizations.
---

# ibm\_iam_authorization_policy

Provides a resource for IAM Service Authorizations. This allows authorization policy to be created and deleted.

## Example Usage

### Authorization policy between two services

```hcl
resource "ibm_iam_authorization_policy" "policy" {
  source_service_name = "cloud-object-storage"
  target_service_name = "kms"
  roles               = ["Reader"]
}

```

### Authorization policy between two services with Authorize dependent services enabled

```hcl
resource "ibm_iam_authorization_policy" "policy" {
  source_service_name         = "databases-for-postgresql"
  target_service_name         = "kms"
  roles                       = ["Reader", "AuthorizationDelegator"]
}
```

### Authorization policy between two services with specific resource type

```hcl
resource "ibm_iam_authorization_policy" "policy" {
  source_service_name  = "is"
  source_resource_type = "image"
  target_service_name  = "cloud-object-storage"
  roles                = ["Reader"]
}

```
### Authorization policy between two specific instances

```hcl
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

```hcl
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

## Argument Reference

The following arguments are supported:

* `source_service_name` - (Required, Forces new resource, string) The source service name.
* `target_service_name` - (Required, Forces new resource, string) The target service name.
* `roles` - (Required, list) comma separated list of roles.
* `source_resource_instance_id` - (Optional, Forces new resource, string) The Source resource instance id.
* `target_resource_instance_id` - (Optional, Forces new resource, string) The target resource instance id.
* `source_resource_type` - (Optional, Forces new resource, string) Resource type of source service.
* `target_resource_type` - (Optional, Forces new resource, string) Resource type of target service.
* `source_service_account` - (Optional, Forces new resource, string) Account GUID of source service.
* `source_resource_group_id` - (Optional, Forces new resource, string) The Source resource group id.
* `target_resource_group_id` - (Optional, Forces new resource, string) The target resource group id.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the authorization policy. 

* `version` - Version of the authorization policy.

## Import

ibm_iam_authorization_policy can be imported using authorization policy ID, eg

```
$ terraform import ibm_iam_authorization_policy.example 12fe9d62-81b1-41ee-8233-53150e38a61c
```
