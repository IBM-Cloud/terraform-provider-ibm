---
layout: "ibm"
page_title: "IBM: ibm_org"
sidebar_current: "docs-ibm-datasource-org"
description: |-
  Get information about an IBM Cloud organization.
---

# ibm\_org

Import the details of an existing IBM Cloud org as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_org" "orgdata" {
  org = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `org` - (Required, string) The name of the IBM Cloud organization. You can retrieve the value by running the `ibmcloud iam orgs` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the organization.  
