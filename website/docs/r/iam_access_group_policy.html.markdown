---
layout: "ibm"
page_title: "IBM : iam_access_group_policy"
sidebar_current: "docs-ibm-resource-iam-access-group-policy"
description: |-
  Manages IBM IAM Access Group Policy.
---

# ibm\_access_group_id

Provides a resource for IAM Access Group Policy. This allows access group policy to be created, updated and deleted.

## Example Usage

### Access Group Policy for All Identity and Access enabled services 

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = "${iam_access_group.accgrp.id}"
  roles        = ["Viewer"]
}

```

### Access Group Policy using service with region

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name = "test"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = "${iam_access_group.accgrp.id}"
  roles        = ["Viewer"]

  resources = [{
    service = "cloud-object-storage"
  }]
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
  access_group_id = "${iam_access_group.accgrp.id}"
  roles        = ["Manager", "Viewer", "Administrator"]

  resources = [{
    service              = "kms"
    region               = "us-south"
    resource_instance_id = "${element(split(":",ibm_resource_instance.instance.id),7)}"
  }]
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
  access_group_id = "${iam_access_group.accgrp.id}"
  roles        = ["Viewer"]

  resources = [{
    service           = "containers-kubernetes"
    resource_group_id = "${data.ibm_resource_group.group.id}"
  }]
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
  access_group_id = "${iam_access_group.accgrp.id}"
  roles        = ["Administrator"]

  resources = [{
    resource_type = "resource-group"
    resource      = "${data.ibm_resource_group.group.id}"
  }]
}

```

## Argument Reference

The following arguments are supported:

* `access_group_id` - (Required, string) ID of the access group.
* `roles` - (Required, list) comma separated list of roles. Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor.
* `resources` - (Optional, list) A nested block describing the resource of this policy.
Nested `resources` blocks have the following structure:
  * `service` - (Optional, string) Service name of the policy definition.  You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
  * `resource_instance_id` - (Optional, string) ID of resource instance of the policy definition.
  * `region` - (Optional, string) Region of the policy definition.
  * `resource_type` - (Optional, string) Resource type of the policy definition.
  * `resource` - (Optional, string) Resource of the policy definition.
  * `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`.
* `tags` - (Optional, array of strings) Tags associated with the access group Policy instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the access group policy. The id is composed of \<access_group_id\>/\<access_group_policy_id\>

* `version` - Version of the access group policy.

## Import

ibm_iam_access_group_policy can be imported using access group ID and access group policy ID, eg

```
$ terraform import ibm_iam_access_group_policy.example AccessGroupId-1148204e-6ef2-4ce1-9fd2-05e82a390fcf/bf5d6807-371e-4755-a282-64ebf575b80a
```