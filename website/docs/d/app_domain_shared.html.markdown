---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_app_domain_shared"
description: |-
  Get information about an IBM Cloud shared domain.
---

# `ibm_app_domain_shared`

Retrieve information about an existing shared domain for an app. For more information, about an app domain shared, see [Managing your domains](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-custom-domains).


## Example usage
The following example retrieves information about the `example.com` domain. 


```
data "ibm_app_domain_shared" "shared_domain" {
  name = "foo.com"
}
```


## Argument reference
Review the input parameters that you can specify for your data source. 

- `name` - (Required, String)  The name of the shared domain.


## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `id` - (String) The unique identifier of the shared domain.


