---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance action"
description: |-
  Manages IBM instance action.
---

# ibm_is_instance_action

Start, stop, or reboot an instance for VPC. For more information, about managing VPC instance, see [about virtual server instances for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-advanced-virtual-servers).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

In the following example, you can perform instance action:

```terraform

resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "bc1-2x8"

  boot_volume {
    encryption = "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
  }

  primary_network_interface {
    subnet               = ibm_is_subnet.example.id
    primary_ipv4_address = "10.240.0.6"
    allow_ip_spoofing    = true
  }

  network_interfaces {
    name              = "eth1"
    subnet            = ibm_is_subnet.example.id
    allow_ip_spoofing = false
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

resource "ibm_is_instance_action" "example" {
  action       = "stop"
  force_action = true
  instance     = ibm_is_instance.example.id
}


```

## Argument reference

Review the argument references that you can specify for your resource. 

- `action` - (Required, String) The type of action to perfrom on the instance. Supported values are `stop`, `start`, or `reboot`.
- `force_action` - (Optional, Boolean)  If set to `true`, the action will be forced immediately, and all queued actions deleted. Ignored for the start action. The Default value is `false`.
- `instance` - (Required, String) Instance identifier.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `status` - (String) The status of the instance. The supported status are **failed**, **pending**, **restarting**, **running**, **starting**, **stopped**, or **stopping**.
- `status_reasons` - (List) Array of reasons for the current status (if any).

  Nested `status_reasons`:
    - `code` - (String) The status reason code.
    - `message` - (String) An explanation of the status reason.
    - `more_info` - (String) Link to documentation about this status reason
    
## Import
The `ibm_is_instance_action` resource can be imported by using instance action ID.

**Example**

```sh
$ terraform import ibm_is_instance_action.example d7bec597-4726-451f-8a63-e62e6f121c32c
```
