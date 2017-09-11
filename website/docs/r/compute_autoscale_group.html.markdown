---
layout: "ibm"
page_title: "IBM : compute_autoscale_group"
sidebar_current: "docs-ibm-resource-compute-autoscale-group"
description: |-
  Manages IBM Compute Auto Scale Group.
---

# ibm\_compute_autoscale_group

Provides a resource for auto scaling groups. This allows auto scaling groups to be created, updated, and deleted.

For additional details, see the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Scale_Group).

## Example Usage

```hcl
# Create an auto scaling group using a Debian image
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
      os_reference_code = "DEBIAN_7_64"
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

* `name` - (Required, string) Name of the auto scaling group.
* `regional_group` - (Required, string) Regional group for the auto scaling group.
* `minimum_member_count` - (Required, integer) The fewest number of virtual guest members allowed in the auto scaling group.
* `maximum_member_count` - (Required, integer) The greatest number of virtual guest members that are allowed in the auto scaling group.
* `cooldown` - (Required, integer) The duration, expressed in seconds, that the auto scaling group waits before performing another scaling action.
* `termination_policy` - (Required, string) The termination policy for the auto scaling group.
* `virtual_guest_member_template` - (Required, array) The template to create guest members with. Only one template can be configured. Accepted values can be found in the [ibm_compute_vm_instance](compute_vm_instance.html) resource.
* `network_vlan_ids` - (Optional, array) Collection of VLAN IDs for this auto scaling group. Accepted values can be found in the [VLAN docs](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID on the resulting URL. Or, you can [refer to a VLAN by name using a data source](../d/network_vlan.html).
* `virtual_server_id` - (Optional, integer) Specifies the ID of a virtual server in a local load balancer. The ID can be found with the following URL: `https://api.softlayer.com/rest/v3/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress/<load_balancer_ID>/getObject?objectMask=virtualServers`. Replace _[load_balancer_ID]_ with the target load balancer's ID and open the URL with a web browser. A Bluemix Infrastructure (SoftLayer) user name and API key are required.
* `port` - (Optional, integer) The port number in a local load balancer. For example, `8080`.
* `health_check` - (Optional, map) Specifies the type of health check in a local load balancer. For example, `HTTP`. You can also use this value to specify custom HTTP methods.
* `tags` - (Optional, array of strings) Set tags on the auto scaling group instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - ID of the auto scaling group.
