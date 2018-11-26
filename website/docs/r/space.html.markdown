---
layout: "ibm"
page_title: "IBM : Space"
sidebar_current: "docs-ibm-resource-space"
description: |-
  Manages IBM Space.
---

# ibm\_space

Provides a space resource. This allows spaces to be created, updated, and deleted.

## Example Usage

```hcl
resource "ibm_space" "space" {
  name        = "myspace"
  org         = "myorg"
  space_quota = "myspacequota"
  managers    = ["manager@example.com"]
  auditors    = ["auditor@example.com"]
  developers  = ["developer@example.com"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify a space.
* `org` - (Required, string) The name of the organization to which this space belongs.
* `space_quota` - (Optional, string) The name of the Space Quota Definition associated with the space.
* `managers` - (Optional, set) The email addresses (associated with IBMids) of the users to whom you want to give a manager role in this space. Users with the manager role can invite users, manage users, and enable features for the given space.
* `developers` - (Optional, set) The email addresses (associated with IBMids) of the users to whom you want to give a developer role in this space. Users with the developer role can create apps and services, manage apps and services, and see logs and reports in the given space.
* `auditors` - (Optional, set) The email addresses (associated with IBMids) of the users to whom you want to give an auditor role in this space. Users with the auditor role can view logs, reports, and settings in the given space.  

**NOTE**: By default the newly created space has no user associated with it. Add your own email address to the `managers` or `developers` field in order to be able to use the space correctly for the first time.

* `tags` - (Optional, array of strings) Tags associated with the space instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new space.
