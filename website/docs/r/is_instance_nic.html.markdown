---
layout: "ibm"
page_title: "IBM : instance_nic"
sidebar_current: "docs-ibm-resource-is-instance-nic"
description: |-
  Manages IBM IS Instance Nic.
---

# ibm\_is_instance_nic

Provides a instance network interface resource. This allows instance network interface to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet"
  vpc             = "${ibm_is_vpc.testacc_vpc.id}"
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "b-2x8"

  primary_network_interface = {
    port_speed = "100"
    subnet     = "${ibm_is_subnet.testacc_subnet.id}"
  }

  vpc  = "${ibm_is_vpc.testacc_vpc.id}"
  zone = "us-south-1"
  keys = ["${ibm_is_ssh_key.testacc_sshkey.id}"]

  //User can configure timeouts
  	timeouts {
      	create = "90m"
      	delete = "30m"
    }
}

resource "ibm_is_instance_nic" "testacc_instance_nic" {
	name        = "test-instance-nic"
	instance_id = "${ibm_is_instance.testacc_instance.id}"
	port_speed = "100"
	subnet = "${ibm_is_subnet.testacc_subnet.id}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the network interface.
* `port_speed` - (Required, int) Speed of the network interface.
* `subnet` -  (Required, string) ID of the subnet.
* `instance_id` - (Required, string) The instance identifier.
* `primary_ipv4_address` -(Optional, string) The primary IPv4 address.
* `primary_ipv6_address` - (Optional, string) The primary IPv6 address in any valid notation as specified by RFC 4291.
* `secondary_addresses` - (Optional, array of strings) A secondary IP address.
* `security_groups` - (Optional, array of strings) Comma separated IDs of security groups.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the instance nic.
* `status` - Status of the instance.

## Import

ibm_is_instance_nic can be imported using instanceID and nicID, eg

```
$ terraform import ibm_is_instance_nic.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
