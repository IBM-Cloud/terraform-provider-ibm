---
layout: "ibm"
page_title: "IBM: pi_instance"
sidebar_current: "docs-ibm-resource-pi-instance"
description: |-
  Manages an instance (a.k.a. VM/LPAR) in the Power Virtual Server Cloud.
---

# ibm\_pi_instance

Provides an instance resource. This allows instances to be created or updated in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create an instance:

```hcl
resource "ibm_pi_instance" "test-instance" {
    pi_memory             = "4"
    pi_processors         = "2"
    pi_instance_name      = "test-vm"
    pi_proc_type          = "shared"
    pi_migratable         = "true"
    pi_image_id           = "<id of the image to deploy - e.g., 7200-03-03>"
    pi_volume_ids         = []
    pi_network_ids        = ["<id of the VM's network IDs>"]
    pi_key_pair_name      = "<name of SSH key>"
    pi_sys_type           = "s922"
    pi_replication_policy = "none"
    pi_replicants         = "1"
    pi_cloud_instance_id  = "<value of the cloud_instance_id>"
```

## Timeouts

ibm_pi_instance provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating an instance.
* `delete` - (Default 60 minutes) Used for deleting an instance.

## Argument Reference

The following arguments are supported:

* `pi_instance_name` - (Required, string) The name of the VM.
* `pi_creation_date` - (Required, string) The date on which the VM was created.
* `pi_key_pair_name` - (Required, string) The name of the Power Virtual Server Cloud SSH key to used to login to the VM.
* `pi_image_id` - (Required, string) The name of the image to deploy (e.g., 7200-03-03).
* `pi_processors` - (Required, float) The number of vCPUs to assign to the VM (as visibile within the guest operating system).
* `pi_proc_type` - (Required, string) The type of processor mode in which the VM will run (shared/dedicated).
* `pi_memory` - (Required, float) The amount of memory (GB) to assign to the VM.
* `pi_sys_type` - (Required, string) The type of system on which to create the VM (s922/e880/any).
* `pi_volume_ids` - (Required, list(string)) The list of volume IDs to attach to the VM at creation time.
* `pi_network_ids` - (Required, list(string)) The list of network IDs assigned to the VM.
* `pi_migratable` - (Required, boolean) Dictates whether the VM can be migrated (e.g., for server maintenance).
* `pi_cloud_instance_id` - (Required, string) The cloud_instance_id for this account.
* `pi_user_data` - (Optional, string) The base64-encoded form of the user data (cloud-init) to pass to the VM at creation time.
* `pi_public_network` - (Optional, boolean) Dictates whether the VM will be attached to a public network (true/false); default is false.
* `pi_replicants` - (Optional, float) Specifies the number of VMs to deploy; default is 1.
* `pi_replication_policy` - (Optional, string) Specifies the replication policy (e.g., none).
* `pi_replication_scheme` - (Optional, string) Specifies the replicate scheme (prefix/suffix).

## Attribute Reference

The following attributes are exported:

* `pi_instance_id` - (string) The unique identifier of the instance.
* `pi_disk_size` - (int) The size (GB) of the root disk.
* `pi_instance_status` - (string) The status of the VM.
* `pi_minproc` - (float) The minimum number of processors the VM can have.
* `addresses` - The addresses associated with this instance.  Nested `addresses` blocks have the following structure:
  * `ip` - IP of the instance.
  * `macaddress` - The macaddress of the instance.
  * `networkid` - The networkID of the instance.
  * `networkname` - The network name of the instance.
  * `type` - The type of the network
  * `externalip` - The externalIP address of the instance.
* `pi_instance_progress` - (float) Specifies the overall progress of the VM deployment process.
