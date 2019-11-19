---
layout: "ibm"
page_title: "IBM : iam_authorization_policy"
sidebar_current: "docs-ibm-resource-iam-authorization-policy"
description: |-
  Manages IBM IAM Service Authorizations.
---

# ibm\_authorization_policy

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

## Argument Reference

The following arguments are supported:

* `source_service_name` - (Required, string) The source service name.
* `target_service_name` - (Required, string) The target service name.
* `roles` - (Required, list) comma separated list of roles.
* `source_resource_instance_id` - (Optional, string) The Source resource instance id.
* `target_resource_instance_id` - (Optional, string) The target resource instance id.
* `source_resource_type` - (Optional, string) Resource type of source service.
* `target_resource_type` - (Optional, string) Resource type of target service.
* `source_service_account` - (Optional, string) Account GUID of source service.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the authorization policy. 

* `version` - Version of the authorization policy.

## Import

ibm_iam_authorization_policy can be imported using authorization policy ID, eg

```
$ terraform import ibm_iam_authorization_policy.example 12fe9d62-81b1-41ee-8233-53150e38a61c
```