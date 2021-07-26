---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : compute_autoscale_group"
description: |-
  Manages IBM Cloud compute auto scale group.
---

# ibm_compute_autoscale_group
Create, update, or delete an auto scaling group. For more information about compute auto scale group, see [enabling auto scale for better capacity and resiliency](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-ha-auto-scale).

**Note**

For more information, about SoftLayer auto scale APIs, see [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Scale_Group).

## Example usage

In the following example, you can create an auto scaling group that uses a Debian image:

```terraform
/* Deprecated in terraform v0.12 hence not updated */

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


## Argument reference
Review the argument references that you can specify for your resource. 

- `cooldown`- (Required, Integer) The duration, expressed in seconds, that the autoscaling group waits before performing another scaling action.
- `health_check` (Optional, Map) The type of health check in a local load balancer. For example, `HTTP`. You can also use this value to specify custom HTTP methods.
- `minimum_member_count`- (Required, Integer) The minimum number of virtual guest members that are allowed in the autoscaling group.
- `maximum_member_count`- (Required, Integer) The maximum number of virtual guest members that are allowed in the autoscaling group.
- `name` - (Required, String) The name of the autoscaling group.
- `network_vlan_ids` - (Optional, Array) The collection of VLAN IDs for the autoscaling group. You can find accepted values in the [VLAN console](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want and notes the ID in the resulting URL. You can also refer to a VLAN name by using a data source.
- `port` - (Optional, Integer) The port number in a local load balancer. For example, `8080`.
- `regional_group` - (Required, Forces new resource, String) The regional group for the autoscaling group.
- `tags` (Optional, Array of Strings) A list of tags that you want to add to the autoscaling group. Tags are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `termination_policy` - (Required, String) The termination policy for the autoscaling group.
- `virtual_guest_member_template` (Required, Array of Strings) The template with which to create guest members. Only one template can be configured. You can find accepted values in the ibm_compute_vm_instance resource.
- `virtual_server_id` - (Optional, Integer) The ID of a virtual server in a local load balancer. You can find the ID with the following URL `https://api.softlayer.com/rest/v3/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress/<load_balancer_ID>/getObject?objectMask=virtualServers`. Replace `<load_balancer_ID>` with the ID of the target load balancer. **Note** To view the load balancer ID. Log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3/SoftLayer_Hardware/getCreateObjectOptions.json) that uses your API key as the password. For more information, about creating classic infrastructure keys and locating your VPN username and password, refer [Managing classic infrastructure API keys](https://cloud.ibm.com/docs/account?topic=account-classic_keys).


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the autoscaling group.
