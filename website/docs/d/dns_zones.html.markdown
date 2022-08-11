---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM Cloud infrastructure private domain name service zones.
---

# ibm_dns_zones

Retrieve details about a zone that you added to your private DNS service instance. For more information, see [Managing DNS zones](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-dns-zones).


## Example usage

```terraform
data "ibm_resource_instance" "dns" { }

data "ibm_dns_zones" "ds_pdnszones" {
  instance_id = data.ibm_resource_instance.dns.guid
}
```


## Argument reference
Review the argument reference that you can specify for your data source. 

- `instance_id` - (Required, String) The GUID of the private DNS service instance.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `dns_zones`- (List) A List of zones that you added to your private DNS service instance. 
   
   Nested scheme for `dns_zones`:
   - `created_on` - (Timestamp) The date and time when the zone was added to the private DNS service instance.
   - `description` - (String) The description of the zone.
   - `instance_id` - (String) The ID of the private DNS service instance where you added the zone.
   - `label` - (String) The label of the zone.
   - `modified_on` - (Timestamp) The date and time when the zone was updated.
   - `name` - (String) The name of the zone.
   - `state` - (String) The state of the zone.
   - `zone_id` - (String) The ID of the zone.

