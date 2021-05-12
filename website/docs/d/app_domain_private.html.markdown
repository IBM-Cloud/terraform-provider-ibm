---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_app_domain_private"
description: |-
  Get information about an IBM Cloud domain private.
---

# `ibm_app_domain_private`

Retrieve information about an existing private domain for an app. For more information, about an app domain, see [getting started with app private domain,](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-getting-started).

## Example usage
The following example retrieves information about the `example.com` domain. 

```
data "ibm_app_domain_private" "private_domain" {
  name = "foo.com"
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 

- `name` - (Required, String) The name of the private domain that is assigned to the app.


## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `id` - (String) The unique identifier of the private app domain.


