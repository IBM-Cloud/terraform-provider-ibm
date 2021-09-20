---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : floating_ip"
description: |-
  Fetches floating IP information.
---

# ibm_floating_ip
Retrieve an information of VPC floating IP on IBM Cloud as a read-only data source. For more information, about floating IP, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

## Example usage

```terraform

    data "ibm_is_floating_ip" "test" {
        name   = "test-fp"
    }

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the floating IP.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `address` - (String) The floating IP address that is created.
- `crn` - (String) The CRN for this floating IP.
- `id` - (String) The unique identifier of the floating IP.
- `status` - (String) Provisioning status of the floating IP address.
- `tags` - (String) The tags associated with VPC.
- `target` - (String) The ID of the network interface used to allocate the floating IP address.
- `zone` - (String) The zone name where to create the floating IP address.
