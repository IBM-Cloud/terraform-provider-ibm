---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_org"
description: |-
  Get information about an IBM Cloud organization.
---

# ibm_org

Retrieve information about an existing Cloud Foundry organization. For more information, about organization, see [updating orgs and spaces](https://cloud.ibm.com/docs/account?topic=account-orgupdates).

## Example usage
The following example retrieves information about the `myorg` Cloud Foundry organization. 

```terraform
data "ibm_org" "orgdata" {
  org = "example.com"
}
```

## Argument reference
Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) The name of the IBM Cloud organization.
- `org` - (Deprecated, String) The name of the IBM Cloud organization.

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the Cloud Foundry organization.


