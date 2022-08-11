---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_cloud_connection_network_attach"
description: |-
  Manages IBM Cloud Connection Network attachment in the Power Virtual Server cloud.
---

# ibm_pi_cloud_connection_network_attach

Attach, detach Network to a Cloud Connection for a Power Systems Virtual Server. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

The following example enables you attach a network to a cloud connection:

```terraform
resource "ibm_pi_cloud_connection_network_attach" "example" {
  pi_cloud_instance_id = "<value of the service instance id>"
  pi_cloud_connection_id = "<value of the cloud connection id>"
  pi_network_id = "<value of the network id>"
}
```

**Note**

* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  
  Example usage:
  
  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Timeouts

The `ibm_pi_cloud_connection_network_attach` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The attach of network to the cloud connection is considered failed if no response is received for 30 minutes.
- **Delete** The detach of network from the cloud connection is considered failed if no response is received for 30 minutes.

## Argument reference

Review the argument references that you can specify for your resource. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_cloud_connection_id` - (Required, String) The Cloud Connection ID.
- `pi_network_id` - (Required, String) The Network ID to attach to this cloud connection.


## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of cloud connection network attachment.

## Import

The `ibm_pi_cloud_connection_network_attach` can be imported by using `pi_cloud_instance_id`, `pi_cloud_connection_id` and `pi_network_id`.

**Example**

```sh
$ terraform import ibm_pi_cloud_connection_network_attach.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb/4726d7be-c597-4438-9f8a-cea6651abc0a
```
