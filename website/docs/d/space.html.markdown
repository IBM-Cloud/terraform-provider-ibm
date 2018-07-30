---
layout: "ibm"
page_title: "IBM: ibm_space"
sidebar_current: "docs-ibm-datasource-space"
description: |-
  Get information about an IBM Cloud space.
---

# ibm\_space

Import the details of an existing IBM Cloud space as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

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
  service           = "speech_to_text"
  plan              = "lite"
  tags              = ["cluster-service", "cluster-bind"]
}

```

## Argument Reference

The following arguments are supported:

* `org` - (Required) The name of your IBM Cloud organization. You can retrieve the value by running the `bx iam orgs` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `space` - (Required) The name of your space. You can retrieve the value by running the `bx iam spaces` command in the IBM Cloud CLI.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the space.  
* `managers` - The email addresses (associated with IBMid) of the users who have a manager role in this space.
* `auditors` - The email addresses (associated with IBMid) of the users who have an auditor role in this space.
* `developers` - The email addresses (associated with IBMid) of the users who have a developer role in this space.
