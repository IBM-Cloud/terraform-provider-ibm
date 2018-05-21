---
layout: "ibm"
page_title: "IBM : service_id"
sidebar_current: "docs-ibm-resource-service-id"
description: |-
  Manages IBM IAM ServiceID.
---

# ibm\_service_id

Provides a resource for IAM ServiceID. This allows serviceID  to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_iam_service_id" "serviceID" {
  name        = "test"
  description = "New ServiceID"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the serviceID.
* `description` - (Optional, string) Description of the serviceID.
* `tags` - (Optional, array of strings) Tags associated with the IAM ServiceID.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the serviceID.
* `version` - Version of the serviceID.
* `crn` - crn of the serviceID.
