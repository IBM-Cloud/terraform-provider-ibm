---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: network_private_ip"
description: |-
  Manages IBM Cloud network public IP.
---

# ibm_network_public_ip
Create, delete, and update a public IP resource to route between servers. Public IPs are not restricted to routing within the same data center. For more information, about IBM Cloud network public IP, see [networking services](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-network).

For more information, see [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/services/SoftLayer_Network_Subnet_IpAddress_Global)

## Example usage

```terraform
resource "ibm_network_public_ip" "test_public_ip " {
    routes_to = "119.81.82.163"
    notes     = "public ip notes"
}
```

## Timeouts
The `ibm_network_public_ip` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating instance.

## Argument reference 
Review the argument references that you can specify for your resource.

- `notes`- (Optional, String) Descriptive text to associate with the public IP instance.
- `routes_to` - (Required, String) The destination IP address that the public IP routes traffic through. The destination IP address can be a public IP address of IBM resources in the same account, such as a public IP address of a VM or public virtual IP addresses of **NetscalerVPXs**.
- `tags`- (Optional, Array of string)  Tags associated with the public IP instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the public IP.
- `ip_address` - (String) The address of the public IP.
