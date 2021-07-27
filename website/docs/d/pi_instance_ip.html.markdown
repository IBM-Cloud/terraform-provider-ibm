---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_ip"
description: |-
  Obtains an information about the IP address for a specific subnet on an instance.
---

# ibm_pi_instance_ip
Retrieve information about a Power Systems Virtual Server instance IP address. For more information, about Power Systems Virtual Server instance IP address, see [configuring and adding a private network subnet](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-configuring-subnet).

## Example Usage

```terraform
data "ibm_pi_instance_ip" "ds_instance_ip" {
  pi_instance_name     = "terraform-test-instance"
  pi_network_name = "APP"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```
**Notes**
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  Example Usage:
  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument reference
Review the argument references that you can specify for your data source. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_name` - (Required, String) The name of the instance.
- `pi_network_name` - (Required, String) The subnet that the instance belongs to. 


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `external_ip` - (String) The external IP of the network that is attached to this instance.
- `ip` - (String) The IP address that is attached to this instance from the subnet.
- `id` - (String) The unique identifier of the network.
- `ipoctet` - (String) The IP octet of the network that is attached to this instance.
- `macaddress` - (String) The MAC address of the network that is attached to this instance.
- `type` - (String) The type of the network that is attached to this instance.
