---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_org"
description: |-
  Get information about an IBM Cloud organization.
---

# `ibm_org`

Retrieve information about an existing Cloud Foundry organization. For more information, about organization, see [Updating orgs and spaces](https://cloud.ibm.com/docs/account?topic=account-orgupdates).

## Example usage
The following example retrieves information about the `myorg` Cloud Foundry organization. 

```
data "ibm_org" "orgdata" {
  org = "example.com"
}
```

## Argument reference
Review the input parameters that you can specify for your data source.

- `name` - (Optional, String) The name of the IBM Cloud organization.
- `org` - (Deprecated, String) The name of the IBM Cloud organization.

## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `id` - (String) The unique identifier of the Cloud Foundry organization.


