---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_cloud_connection"
description: |-
  Manages IBM Cloud connection in the Power Virtual Server cloud.
---

# ibm_pi_cloud_connection

Retrieve information about an existing IBM Cloud Power Virtual Server Cloud cloud connection. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_cloud_connection" "example" {
  pi_cloud_connection_name  = "test_cloud_connection"
  pi_cloud_instance_id      = "<value of the cloud_instance_id>"
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
- `pi_cloud_connection_name` - (Required, String) The cloud connection name to be used.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `classic_enabled` - (Boolean) Enable classic endpoint destination.
- `connection_mode` - (String) Type of service the gateway is attached to.
- `global_routing` - (String) Enable global routing for this cloud connection.
- `gre_destination_address` - (String) GRE destination IP address.
- `gre_source_address` - (String) GRE auto-assigned source IP address.
- `id` - (String) The unique identifier of the cloud connection.
- `ibm_ip_address` - (String) The IBM IP address.
- `metered` - (String) Enable metering for this cloud connection.
- `networks` - (Set) Set of Networks attached to this cloud connection.
- `port` - (String) Port.
- `speed` - (Integer) Speed of the cloud connection (speed in megabits per second).
- `status` - (String) Link status.
- `user_ip_address` - (String) User IP address.
- `vpc_crns` - (Set) Set of VPCs attached to this cloud connection.
- `vpc_enabled` - (Boolean) Enable VPC for this cloud connection.
