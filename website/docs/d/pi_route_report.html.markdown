---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_route_report"
description: |-
  Manages route report for workspace in the Power Virtual Server cloud.
---

# ibm_pi_route_report

Retrieve a route report. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_route_report" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `routes` - (List) List of all route report routes.

  Nested scheme for `routes`:
  - `destination` - (String) The route destination.
  - `type` - (String) The route type.
