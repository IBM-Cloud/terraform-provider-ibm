---
layout: "ibm"
page_title: "IBM : is_share_replica_operations"
description: |-
  Manages Share replica operations fallback and split.
subcategory: "Virtual Private Cloud API"
---

# ibm\_is_share_target

Provides a resource for ShareTarget. This allows ShareTarget to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_is_share_replica_operations" "test" {
  share_replica = ibm_is_share.replica.id
  split_share = true
}
```

```hcl
resource "ibm_is_share_replica_operations" "test" {
  share_replica = ibm_is_share.replica.id
  fallback_policy = "split"
  timeout = 500
}
```

## Argument Reference

The following arguments are supported:

* `share_replica` - (Required, string) The file share identifier.
* `fallback_policy` - (Optional, string) The action to take if the failover request is accepted but cannot be performed or times out. Accepted values are **split**, **fail**
* `timeout` - (Optional, string) The failover timeout in seconds. Required with `fallback_policy`
* `split_share` - (Boolean, string) If set to true the replication relationship between source share and replica will be removed.

~>**Note**

`split_share` and `fallback_policy` are mutually exclusive

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the Share.
