---
layout: "ibm"
page_title: "IBM: pi_instance"
sidebar_current: "docs-ibm-datasources-pi-instance"
description: |-
  Manages an instance in the Power Virtual Server Cloud.
---

# ibm\_pi_instance

Import the details of an existing IBM Power Virtual Server instance as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_instance" "ds_instance" {
  pi_instance_name     = "terraform-test-instance"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `pi_instance_name` - (Required, string) The name of the instance.
* `pi_cloud_instance_id` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this instance.
* `addresses` - The addresses associated with this instance.  Nested `addresses` blocks have the following structure:
	* `ip` - IP of the instance.
  * `macaddress` - The macaddress of the instance.
  * `networkid` - The networkID of the instance.
  * `networkname` - The network name of the instance.
  * `type` - The type of the network
  * `externalip` - The externalIP address of the instance.
* `state` - The state of the instance.
