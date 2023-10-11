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
resource "ibm_resource_instance" "test-pdns-instance" {
  name              = "test-pdns"
  resource_group_id = data.ibm_resource_group.rg.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}

resource "ibm_dns_zone" "test-pdns-zone" {
  name        = "test.com"
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  description = "testdescription"
  label       = "testlabel-updated"
}

resource "ibm_dns_linked_zone" "test" {
  name          = "test_dns_linked_zone"
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  description   = "seczone terraform plugin test"
  owner_instance_id = "OWNER Instance ID"
  owner_zone_id = "OWNER ZONE ID"
  label         = "test"
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `instance_id` - (Required, String) The unique identifier of a DNS Linked zone.
- `name`        - (Required, String) The name of the DNS Linked zone.
- `description` - (Optional, String) Descriptive text of the DNS Linked zone.
- `owner_instance_id` - (Required, String) The unique identifier of the owner DNS instance.
- `owner_zone_id`     - (Required, String) The unique identifier of the owner DNS zone.
- `label`             - (Optional, String) The label of the DNS Linked zone.
- `approval_required_before` - (Optional, String) DNS Linked Approval required before.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `state`      - (String) The state of the DNS Linked zone.
- `created_on` - (Timestamp) The time (created On) of the Linked Zone. 
- `modified_on` - (Timestamp) The time (modified On) of the Linked Zone.

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
