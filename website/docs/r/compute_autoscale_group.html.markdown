---
layout: "ibm"
page_title: "IBM : compute_autoscale_group"
sidebar_current: "docs-ibm-resource-compute-autoscale-group"
description: |-
  Manages IBM Compute Auto Scale Group.
---

# ibm\_compute_autoscale_group

Provides a resource for auto scaling groups. This allows auto scaling groups to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Scale_Group).

## Example Usage

In the following example, you can create an auto scaling group using a Debian image:

```hcl
resource "ibm_compute_autoscale_group" "test_scale_group" {
    name = "test_scale_group_name"
    regional_group = "as-sgp-central-1"
    minimum_member_count = 1
    maximum_member_count = 10
    cooldown = 30
    termination_policy = "CLOSEST_TO_NEXT_CHARGE"
    virtual_server_id = 267513
    port = 8080
    health_check = {
      type = "HTTP"
    }
    virtual_guest_member_template = {
      hostname = "test_virtual_guest_name"
      domain = "example.com"
      cores = 1
      memory = 1024
      network_speed = 100
      hourly_billing = true
      os_reference_code = "DEBIAN_8_64"
# Optional fields for virtual guest template (SoftLayer defaults apply):
      local_disk = false
      disks = [25]
      datacenter = "sng01"
      post_install_script_uri = ""
      ssh_key_ids = [383111]
      user_metadata = "#!/bin/bash ..."
    }
# Optional fields for scale_group:
    network_vlan_ids = [1234567, 7654321]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the auto scaling group.
* `regional_group` - (Required, string) The regional group for the auto scaling group.
* `minimum_member_count` - (Required, integer) The fewest number of virtual guest members that are allowed in the auto scaling group.
* `maximum_member_count` - (Required, integer) The greatest number of virtual guest members that are allowed in the auto scaling group.
* `cooldown` - (Required, integer) The duration, expressed in seconds, that the auto scaling group waits before performing another scaling action.
* `termination_policy` - (Required, string) The termination policy for the auto scaling group.
* `virtual_guest_member_template` - (Required, array) The template with which to create guest members. Only one template can be configured. You can find accepted values in the [ibm_compute_vm_instance](compute_vm_instance.html) resource.
* `network_vlan_ids` - (Optional, array) The collection of VLAN IDs for the auto scaling group. You can find accepted values in the [VLAN docs](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID in the resulting URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).
* `virtual_server_id` - (Optional, integer) The ID of a virtual server in a local load balancer. You can find the ID with the following URL: `https://api.softlayer.com/rest/v3/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress/<load_balancer_ID>/getObject?objectMask=virtualServers`. Replace _<load_balancer_ID>_ with the ID of the target load balancer. An IBM Cloud Infrastructure (SoftLayer) user name and API key are required.
* `port` - (Optional, integer) The port number in a local load balancer. For example, `8080`.
* `health_check` - (Optional, map) The type of health check in a local load balancer. For example, `HTTP`. You can also use this value to specify custom HTTP methods.
* `tags` - (Optional, array of strings) Tags associated with the auto scaling group instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the auto scaling group.
