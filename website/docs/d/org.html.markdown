---
layout: "ibm"
page_title: "IBM: ibm_org"
sidebar_current: "docs-ibm-datasource-org"
description: |-
  Get information about an IBM Bluemix organization.
---

# ibm\_org

Import the details of an existing IBM Bluemix org as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax. 

## Example Usage

```hcl
data "ibm_org" "orgdata" {
  org = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `org` - (Required) The name of the Bluemix org. The value can be retrieved by running the `bx iam orgs` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the org.  
