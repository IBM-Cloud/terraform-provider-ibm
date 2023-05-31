---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_linked_zone"
description: |-
  Manages IBM Private DNS linked zone.
---

# ibm_dns_linked_zone

The DNS linked zone resource allows users to request and manage linked zones. 


## Example usage

```
resource "ibm_dns_linked_zone" "test" {
  name          = "test_dns_linked_zone"
  instance_id   = "****************************"
  description   = "seczone terraform plugin test"
  owner_instance_id = "**************************"
  owner_zone_id = "************************"
  label         = "test"
}

data "ibm_dns_linked_zone" "test-lz" {
  depends_on  = [ibm_dns_linked_zone.test]
  instance_id   = "***********************"
}

output "ibm_dns_linked_zone_id" {
  value = data.ibm_dns_linked_zone.test-lz.instance_id
}

```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `instance_id` - (Required, String) The unique identifier of a service instance.
- `description` - (Optional, String) Descriptive text of the linked zone.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time (created On) of the Linked Zone. 
- `modified_on` - (Timestamp) The time (modified On) of the Linked Zone.
- `linked_zone_id` - (String) The unique ID of the DNS Services linked zone.

## Import
The `ibm_dns_linked_zone` can be imported by using DNS Services instance ID, Linked Zone ID.
The `id` property can be formed from `instance_id`, `linked_zone_id` in the following format:

```
<instance_id>/<linked_zone_id>
```

**Example**

```
terraform import ibm_dns_linked_zone.sample "d10e6956-377a-43fb-a5a6-54763a6b1dc2/63481bef-3759-4b5e-99df-73be7ba40a8a"
```
