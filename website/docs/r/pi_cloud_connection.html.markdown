---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_cloud_connection"
description: |-
  Manages IBM Cloud connection in the Power Virtual Server cloud.
---

# ibm_pi_cloud_connection

Create, update, or delete for a Power Systems Virtual Server cloud connection. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

~> **NOTE:** `Cloud connection` are not supported in **new** workspaces in `DAL10 data center.`

## Example usage

The following example enables you to create a cloud connection:

```terraform
resource "ibm_pi_cloud_connection" "cloud_connection" {
  pi_cloud_instance_id  = "<value of the cloud_instance_id>"
  pi_cloud_connection_name = "test_cloud_connection"
  pi_cloud_connection_speed = 50
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

The `ibm_pi_cloud_connection` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating cloud connection.
- **update** - (Default 30 minutes) Used for updating cloud connection.
- **delete** - (Default 30 minutes) Used for deleting cloud connection.

## Argument reference

Review the argument references that you can specify for your resource.

- `pi_cloud_connection_classic_enabled` - (Optional, Bool) Enable classic endpoint destination.
- `pi_cloud_connection_global_routing` - (Optional, Bool) Enable global routing for this cloud connection.
- `pi_cloud_connection_gre_cidr` - (Optional, String) The GRE network in CIDR notation.
- `pi_cloud_connection_gre_destination_address` - (Optional, String) The GRE destination IP address.
- `pi_cloud_connection_metered` - (Optional, Bool) Enable metered for this cloud connection.
- `pi_cloud_connection_name` - (Required, String) The name of the cloud connection.
- `pi_cloud_connection_networks` - (Optional, Set of String) Set of Networks to attach to this cloud connection.
- `pi_cloud_connection_speed` - (Required, String) Speed of the cloud connection (speed in megabits per second). Supported values are `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`, `10000`.
- `pi_cloud_connection_transit_enabled` - (Optional, Bool) Enable transit gateway for this cloud connection.
- `pi_cloud_connection_vpc_crns` - (Optional, Set of String) Set of VPC CRNs to attach to this cloud connection.
- `pi_cloud_connection_vpc_enabled` - (Optional, Bool) Enable VPC for this cloud connection.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `cloud_connection_id` - (String) The cloud connection ID.
- `connection_mode` - (String) Type of service the gateway is attached to.
- `gre_source_address` - (String) The GRE auto-assigned source IP address.
- `ibm_ip_address` - (String) The IBM IP address.
- `id` - (String) The unique identifier of cloud connection.
- `port` - (String) Port.
- `status` - (String) Link status.
- `user_ip_address` - (String) User IP address.

## Import

The `ibm_pi_cloud_connection` can be imported by using `power_instance_id` and `cloud_connection_id`.

### Example

```bash
terraform import ibm_pi_cloud_connection.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
