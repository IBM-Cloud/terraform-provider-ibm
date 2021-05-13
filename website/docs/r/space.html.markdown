---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM : Space"
description: |-
  Manages IBM space.
---

# `ibm_space`

Create, update, or delete a Cloud Foundry space. For more information, about Cloud Foundry space, see [Getting started with IBM Cloud Foundry Enterprise Environment](https://cloud.ibm.com/docs/cloud-foundry?topic=cloud-foundry-getting-started).


## Example usage
The following example creates the `myspace` Cloud Foundry space. 


```
resource "ibm_space" "space" {
  name        = "myspace"
  org         = "myorg"
  space_quota = "myspacequota"
  managers    = ["manager@example.com"]
  auditors    = ["auditor@example.com"]
  developers  = ["developer@example.com"]
}
```

## Argument reference
Review the input parameters that you can specify for your resource. 

- `auditors`(Optional, Sets)  The email addresses (associated with IBM IDs) of the users to whom you want to give an auditor role in this space. Users with the auditor role can view logs, reports, and settings in the given space.
- `developers`(Optional, Sets) The email addresses (associated with IBM IDs) of the users to whom you want to give a developer role in this space. Users with the developer role can create apps and services, manage apps and services, and see logs and reports in the given space.
- `managers`(Optional, Sets) The email addresses (associated with IBM IDs) of the users to whom you want to give a manager role in this space. Users with the manager role can invite users, manage users, and enable features for the given space.
- `name` - (Required, String) The descriptive name for the space.
- `org` - (Required, String) The name of the Cloud Foundry organization to which this space belongs.
- `space_quota` - (Optional, String) The name of the space quota definition that is associated with the space.
- `tags` (Optional, Array of Strings) The tags that you want to add to the space. Tags can help you find the space more easily later. **Note** Currently, `Tags` that are managed locally and not sored on the IBM Cloud service endpoint.

**Note**
 By default, the newly created space has no user associated with it. Add your own email address to the `managers` or `developers` field in order to be able to use the space correctly for the first time.

## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `id` - (String) The unique identifier of the new space.
