---
layout: "ibm"
page_title: "IBM: ibm_account"
sidebar_current: "docs-ibm-datasource-account"
description: |-
  Get information about an IBM Bluemix account.
---

# ibm\_account

Import the details of an existing IBM Bluemix account as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_org" "orgData" {
  org = "example.com"
}

data "ibm_account" "accountData" {
  org_guid = "${data.ibm_org.orgData.id}"
}
```

## Argument Reference

The following arguments are supported:

* `org_guid` - (Required, string) The GUID of the Bluemix organization. You can retrieve the value from the `ibm_org` data source or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the account.  
