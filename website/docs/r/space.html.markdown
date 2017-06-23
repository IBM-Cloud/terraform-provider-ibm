---
layout: "ibm"
page_title: "IBM : Space"
sidebar_current: "docs-ibm-resource-space"
description: |-
  Manages IBM Space.
---

# ibm\_space

Create, update, or delete spaces for IBM Bluemix.

## Example Usage

```hcl
resource "ibm_space" "space" {
  name        = "myspace"
  org         = "myorg"
  space_quota = "myspacequota"
  managers = ["manager@example.com"]
  auditors = ["auditor@example.com"]
  developers = ["developer@example.com"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify a space.
* `org` - (Required, string) Name of the org this space belongs to.
* `space_quota` - (Optional, string) The name of the Space Quota Definition associated with the space.
* `managers` - (Optional, set) The emails (associated with IBM ID) of the users who will be given manager role in this space. They can invite and manage users, and enable features for a given space.
* `developers` - (Optional, set) The emails (associated with IBM ID) of the users who will be given developer role in this space. They can create and manage apps and services, and see logs and reports.
* `auditors` - (Optional, set) The emails (associated with IBM ID) of the users who will be given auditor role in this space. They can view logs, reports, and settings on this space.

**Note**: By default the newly created space doesn't have any user associated with it. You should add your email to one of the `managers` or `developers` field in order to be able to use the space correctly for the first time.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new space.
