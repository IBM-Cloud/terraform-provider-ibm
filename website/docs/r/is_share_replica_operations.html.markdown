---
layout: "ibm"
page_title: "IBM : is_share_replica_operations"
description: |-
  Manages Share replica operations fallback and split.
subcategory: "VPC infrastructure"
---

# ibm\_is_share_target

Provides a resource for ShareTarget. This allows ShareTarget to be created, updated and deleted.

~> **NOTE**
IBM Cloud® File Storage for VPC is available for customers with special approval. Contact your IBM Sales representative if you are interested in getting access.

~> **NOTE**
This is a Beta feature and it is subject to change in the GA release 

~> **NOTE**
`ibm_is_share_replica_operations` is used for either failing over to replica share or splitting the source and replica shares. 
When a failover is performed, replica share becomes the source, and the source share becomes replica. Hence terraform configuration should be modified and adjusted accordingly.


## Example Usage

```terraform
resource "ibm_is_share" "example" {
  name = "my-share"
  size = 200
  profile = "tier-3iops"
  zone = "us-south-2"
}
```
## Example Usage (Create a replica share)

```terraform
resource "ibm_is_share" "example1" {
    zone = "us-south-3"
    source_share = ibm_is_share.example.id
    name = "my-replica1"
    profile = "tier-3iops"
    replication_cron_spec = "0 */5 * * *"
}
```

## Example Usage

```hcl
// Split source share and replica share
resource "ibm_is_share_replica_operations" "test" {
  share_replica = ibm_is_share.example1.id
  split_share = true
}
```


```hcl
// failover to replica share
resource "ibm_is_share_replica_operations" "test" {
  share_replica = ibm_is_share.example1.id
  fallback_policy = "split"
  timeout = 500
}
```

## Argument Reference

~> **Note** 
  `ibm_is_share_replica_operations` is a one time action and can be removed once done. Configuration for share and replica share should be adjusted accordingly.

The following arguments are supported:

* `share_replica` - (Required, string) The file share identifier.
* `fallback_policy` - (Optional, string) The action to take if the failover request is accepted but cannot be performed or times out. Accepted values are **split**, **fail**
* `timeout` - (Optional, string) The failover timeout in seconds. Required with `fallback_policy`
* `split_share` - (Boolean, string) If set to true the replication relationship between source share and replica will be removed.

~>**Note**

`split_share` and `fallback_policy` are mutually exclusive

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the Share.
