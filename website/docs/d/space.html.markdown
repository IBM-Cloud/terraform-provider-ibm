---
layout: "ibm"
page_title: "IBM: ibm_space"
sidebar_current: "docs-ibm-datasource-space"
description: |-
  Get information about an IBM Bluemix space.
---

# ibm\_space

Import the details of an existing IBM Bluemix space as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax. 

## Example Usage

```hcl
data "ibm_space" "spaceData" {
  space = "prod"
  org   = "someexample.com"
}
```

The following example shows how you can use the data source to reference the space ID in the `ibm_service_instance` resource.

```hcl
resource "ibm_service_instance" "service_instance" {
  name              = "test"
  space_guid        = "${data.ibm_space.spaceData.id}"
  service           = "cloudantNOSQLDB"
  plan              = "Lite"
  tags              = ["cluster-service", "cluster-bind"]
}

```

## Argument Reference

The following arguments are supported:

* `org` - (Required) The name of your Bluemix org. The value can be retrieved by running the `bx iam orgs` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `space` - (Required) The name of your space. The value can be retrieved by running the `bx iam spaces` command in the Bluemix CLI.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the space.  
* `managers` - The emails (associated with IBM ID) of the users who have manager role in this space
* `auditors` - The emails (associated with IBM ID) of the users who have auditor role in this space
* `developers` - The emails (associated with IBM ID) of the users who have developer role in this space
