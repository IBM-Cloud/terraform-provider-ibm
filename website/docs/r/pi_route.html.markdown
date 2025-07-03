---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_route"
description: |-
  Manages a route in a routing table in the Power Virtual Server cloud.
---

# ibm_pi_route

Create, update or delete a route.

## Example usage

The following example enables you to create a route:

```terraform
resource "ibm_pi_route" "route" {
  pi_cloud_instance_id = "<cloud-instance-id>"
  pi_name              = "test-route"
  pi_next_hop          = "<next-hop-ip>"
  pi_destination       = "<destination-cidr>"
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

## Timeouts

ibm_pi_route provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating a route.
- **update** - (Default 10 minutes) Used for updating a route.
- **delete** - (Default 10 minutes) Used for deleting a route.

## Argument reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_action` - (Required, String) Specifies action for route. Valid values are `deliver`. Default value is `deliver`.
- `pi_advertise` - (Optional, String) Indicates if the route is advertised. Default is `enable`.
- `pi_destination` - (Required, String) Destination of route.
- `pi_destination_type` - (Required, String) The destination type. valid values are `ipv4-address`. Default value is `ipv4-address`.
- `pi_enabled` - (Optional, Boolean) Indicates if the route should be enabled in the fabric. Default value is `false`.
- `pi_name` - (Required, String) Name of the route.
- `pi_next_hop` - (Required, String) The next hop.
- `pi_next_hop_type` - (Required, String) The next hop type. Valid values are `ipv4-address`. Default value is `ipv4-address`.
- `pi_user_tags` - (Optional, Set of String) The user tags attached to this resource.

## Attribute reference

 In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of this resource.
- `id` - (String) The unique identifier of the route in the terraform state.
- `route_id` - (String) The unique route ID.
- `state` - (String) The state of the route.

## Import

The `ibm_pi_route` resource can be imported by using `power_instance_id` and `route_id`.

### Example

```bash
terraform import ibm_pi_route.example d7bec597-4726-451f-8a63-e62e6f19c32c/b17a2b7f-77ab-491c-811e-495f8d4c8947
```
