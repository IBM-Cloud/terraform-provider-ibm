---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: app_domain_private"
description: |-
  Manages IBM application private domain.
---

# `ibm_app_domain_private`

Create, update, or delete a private domain for your Cloud Foundry app. For more information, about an app domain, see [getting started with app private domain,](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-getting-started).


## Example usage
The following example creates the `example.com` private domain. 

```
data "ibm_org" "orgdata" {
  org = "example.com"
}

resource "ibm_app_domain_private" "domain" {
  name     = "example.com"
  org_guid = data.ibm_org.orgdata.id
  tags     = ["tag1", "tag2"]
}
```

## Argument reference
Review the input parameters that you can specify for your resource. 

- `name`- (Required, String) The name of the private domain.
- `org_guid` - (Required, String) The GUID of the Cloud Foundry organization where you want to create the domain. You can retrieve the value from data source `ibm_org` or by running the `ibmcloud iam orgs guid` command in the IBM Cloud CLI.
- `tags`- (Array of Strings, Optional) The tags that you want to add to your private domain.

## Attribute reference
Review the output parameters that you can access after your resource is created.

- `id` - (String) The unique identifier of the private domain.




