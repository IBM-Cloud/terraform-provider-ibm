---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_action"
description: |-
  Manages IBM bare metal sever action.
---

# ibm\_is_bare_metal_server_action

Start/Stop/Restart a Bare Metal Server for VPC. For more information, about managing VPC Bare Metal Server, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```


## Example Usage

In the following example, you can update name of a Bare Metal Server disk:

```terraform
resource "ibm_is_bare_metal_server_action" "bms_action" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  action            = "stop"
  stop_type         = "hard"
}
```

## Argument Reference

Review the argument references that you can specify for your resource. 


- `action` - (Required, String) The type of action to perfrom on the  Bare metal server.[**start**, **stop**, **restart**]
- `bare_metal_server` - (Required, String) Bare metal server identifier. 
- `stop_type` - (Optional, String) The type of stop for the `stop` action. [**soft**, **hard**]. By default its `hard`


## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `status` - (String) The status of the bare metal server :[ failed, pending, restarting, running, starting, stopped, stopping ]
- `status_reasons` - (List) Array of reasons for the current status (if any).

  Nested `status_reasons`:
    - `code` - (String) The status reason code
    - `message` - (String) An explanation of the status reason
    - `more_info` - (String) Link to documentation about this status reason
