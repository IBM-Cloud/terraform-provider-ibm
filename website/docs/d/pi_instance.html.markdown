---
layout: "ibm"
page_title: "IBM : Instance"
sidebar_current: "docs-ibm-datasources-pi-instance"
description: |-
  Manages IBM Cloud Infrastructure Instance for IBM Power
---

# ibm\_pi_instance

Import the details of an existing IBM Cloud Infrastructure Instance for IBM Power as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_pi_instance" "ds_instance" {
    name = "terraform-test-instance"
    powerinstanceid="49fba6c9-23f8-40bc-9899-aca322ee7d5b"

}

```

## Argument Reference

The following arguments are supported:

* `instancename` - (Required, string) The name of the instance.
* `powerinstanceid` - (Required, string) The service instance associated with the account


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this instance.
* `address` - The address associated with this instance.
* `state` - The state of the instance.




