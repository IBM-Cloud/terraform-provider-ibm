---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM : Cloud Internet Services instance"
description: |-
  Get information on an IBM Cloud Internet Services instance.
---

# ibm_cis

Retrieve information about an existing CIS resource. This allows CIS sub resources to be added to an existing CIS instance. This includes domains, DNS records, pools, healthchecks and Global Load Balancers. For more information, about CIS instance, see [getting started with CIS](https://cloud.ibm.com/docs/cis?topic=cis-getting-started).

## Example usage

```terraform
data "ibm_cis" "cis_instance" {
  name              = "test"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of a CIS instance.
- `resource_group_id` - (Optional, String) The ID of the resource group where you want to create the service. You can retrieve the value from data source `ibm_resource_group`. If not provided creates the service in default resource group.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `guid` -  (String) The unique identifier of the instance.
- `id` - (String) The CRN of your instance.
- `location` - (String) The location of your instance.
- `plan` - (String) The service plan for the instance.
* `service` - (String) The service type of the instance.
* `status` - (String) Status of the resource instance.