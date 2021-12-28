---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_zone"
description: |-
  Manages IBM Private DNS Zone.
---


# ibm_dns_zone

Create, update, or delete a DNS zone. For more information, see [Managing DNS zones](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-dns-zones).


## Example usage

```terraform
resource "ibm_dns_zone" "pdns-1-zone" {
    name = "test.com"
    instance_id = p-dns-instance-id
    description = "testdescription"
    label = "testlabel"
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `description` - (Optional, String) The description of the DNS zone.
- `instance_id` - (Required, String) The GUID of the IBM Cloud DNS service instance where you want to create a DNS zone.
- `name` - (Required, String) The name of the DNS zone that you want to create. 
- `label` - (Optional, String) The label of the DNS zone. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time when the DNS zone was created. 
- `id` - (String) The ID of the DNS zone. The ID is composed of `<instance_id>/<zone_id>`.
- `modified_on` - (Timestamp) The time when the DNS zone was updated. 
- `state` - (String) The state of the DNS zone.
- `zone_id` - (String) The ID of the zone that is associated with the DNS zone.

## Import

The `ibm_dns_zone` resource can be imported by using private DNS instance ID and zone ID.

**Example**

```
$ terraform import ibm_dns_zone.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308
```
