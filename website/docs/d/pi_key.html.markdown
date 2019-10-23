---
layout: "ibm"
page_title: "IBM: pi_key"
sidebar_current: "docs-ibm-datasources-pi-key"
description: |-
  Manages an key in the Power Virtual Server Cloud.
---

# ibm\_pi_key

Import the details of an existing IBM Power Virtual Server key as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_key" "ds_instance" {
  pi_key_name          = "terraform-test-key"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `pi_key_name` - (Required, string) The name of the key.
* `pi_cloud_instance_id` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this instance.
* `creationdate` - The creation date.
* `sshkey` - The SSH RSA key.
