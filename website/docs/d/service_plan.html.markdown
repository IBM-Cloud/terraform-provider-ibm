---
layout: "ibm"
page_title: "IBM: ibm_service_plan"
sidebar_current: "docs-ibm-datasource-service-plan"
description: |-
  Get information about a service plan from IBM Cloud.
---

# ibm\_service_plan

Import the details of an existing service plan from IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_service_plan" "service_plan" {
  service = "cloudantNoSQLDB"
  plan    = "Lite"
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required, string) The name of the service offering. You can retrieve the name of the service by running the `ibmcloud service offerings` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `plan` - (Required, string) The name of the plan type supported by the service. You can retrieve the plan type by running the `ibmcloud service offerings` command in the IBM Cloud CLI.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the service plan.  
