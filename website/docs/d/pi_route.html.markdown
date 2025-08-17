---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_route"
description: |-
  Manages route in a routing table in the Power Virtual Server cloud.
---

# ibm_pi_route

Retrieve information about a route. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_route" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_route_id          = "<route id>
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
- `pi_route_id` - (Required, String) ID of route.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `action` - (String) The route action.
- `advertise` - (String) Indicates if the route is advertised.
- `crn` - (String) The CRN of this resource.
- `destination` - (String) The route destination.
- `destination_type` - (String) The destination type.
- `enabled` - (Boolean) Indicates if the route should be enabled in the fabric.
- `name` - (String) Name of the route.
- `next_hop` - (String) The next hop in the route.
- `next_hop_type` - (String) The next hop type.
- `state` - (String) The state of the route.
- `user_tags` - (Set of String) List of user tags attached to the resource.
