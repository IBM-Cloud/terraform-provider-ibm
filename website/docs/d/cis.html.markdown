---
layout: "ibm"
page_title: "IBM : Cloud Internet Services instance"
sidebar_current: "docs-ibm-datasource-cis"
description: |-
  Get information on an IBM Cloud Internet Services Instance.
---

# ibm\_cis

Imports the a read only copy of the details of an existing Internet Services resource. This allows CIS sub-resources to be added to an existing CIS instance. This includes domains, DNS records, pools, healthchecks and global load balancers. 

## Example Usage

```hcl
data "ibm_cis" "cis_instance" {
  name              = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name used to identify the Internet Services instance in the IBM Cloud UI. 

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of this CIS instance.
* `plan` - The service plan for this Internet Services' instance
* `location` - The location for this Internet Services' instance
* `status` - Status of the CIS instance.
