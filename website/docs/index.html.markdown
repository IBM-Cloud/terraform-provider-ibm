---
layout: "ibm"
page_title: "Provider: IBM"
sidebar_current: "docs-ibm-index"
description: |-
  The IBM Cloud provider is used to interact with IBM Cloud resources.
---

# IBM Cloud Provider

The IBM Cloud provider is used to manage IBM Cloud resources. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation menu on the left to read about the available resources.


## Example Usage


```hcl
# Configure the IBM Cloud Provider
provider "ibm" {
  bluemix_api_key    = "${var.ibm_bmx_api_key}"
  softlayer_username = "${var.ibm_sl_username}"
  softlayer_api_key  = "${var.ibm_sl_api_key}"
}

# Create an SSH key. The SSH key surfaces in the SoftLayer console under Devices > Manage > SSH Keys.
resource "ibm_compute_ssh_key" "test_key_1" {
  label      = "test_key_1"
  public_key = "${var.ssh_public_key}"
}

# Create a virtual server with the SSH key.
resource "ibm_compute_vm_instance" "my_server_2" {
  hostname          = "host-b.example.com"
  domain            = "example.com"
  ssh_key_ids       = [123456, "${ibm_compute_ssh_key.test_key_1.id}"]
  os_reference_code = "CENTOS_6_64"
  datacenter        = "ams01"
  network_speed     = 10
  cores             = 1
  memory            = 1024
}

# Read details of IBM Bluemix Space
data "ibm_space" "space" {
  space = "${var.space}"
  org   = "${var.org}"
}

# Create an instance of a IBM service.
resource "ibm_service_instance" "service" {
  name       = "${var.instance_name}"
  space_guid = "${data.ibm_space.space.id}"
  service    = "cleardb"
  plan       = "cb5"
  tags       = ["cluster-service", "cluster-bind"]
}
```

## Authentication

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

### Static credentials ###

Static credentials can be provided by adding an `bluemix_api_key`, `softlayer_username`, `softlayer_api_key` in the IBM Cloud provider block.

Usage:

```
provider "ibm" {
    bluemix_api_key = ""
    softlayer_username = ""
    softlayer_api_key = ""

}
```


### Environment variables

You can provide your credentials with the `BM_API_KEY`, `SL_USERNAME`, and `SL_API_KEY` environment variables, representing your Bluemix API key, SoftLayer user name, and SoftLayer API key, respectively.  

```
provider "ibm" {}
```

Usage:

```
$ export BM_API_KEY="bmx_api_key"
$ export SL_USERNAME="sl_username"
$ export SL_API_KEY="sl_api_key"
$ terraform plan
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `bluemix_api_key` - (Optional) The Bluemix API key. It must be provided, but it can also be sourced from the `BM_API_KEY` or `BLUEMIX_API_KEY` environment variable. The former variable has higher precedence. The key is required to provision Cloud Foundry or IBM Container Service resources, such as any resource that begins with `ibm` or `ibm_container`.

* `bluemix_timeout` - (Optional) The timeout, expressed in seconds, for the SoftLayer API key. It can also be sourced from the `BM_TIMEOUT` or `BLUEMIX_TIMEOUT` environment variable. The former variable has higher precedence. Default value: `60`.

* `softlayer_username` - (Optional) The SoftLayer user name. It must be provided, but it can also be sourced from the `SL_USERNAME` or `SOFTLAYER_USERNAME` environment variable. The former variable has higher precedence. 

* `softlayer_api_key` - (Optional) The SoftLayer user name. It must be provided, but it can also be sourced from the `SL_API_KEY` or `SOFTLAYER_API_KEY` environment variable. The former variable has higher precedence. The key is required to provision SoftLayer resources, such as any resource that begins with `ibm_compute`.

* `softlayer_timeout` - (Optional) The timeout, expressed in seconds, for the SoftLayer API key. It can also be sourced from the `SL_TIMEOUT` or `SOFTLAYER_TIMEOUT` environment variable. The former variable has higher precedence. Default value: `60`.

* `region` - (Optional) The Bluemix region. It can also be sourced from the `BM_REGION` or `BLUEMIX_REGION` environment variable. The former variable has higher precedence. Default value: `us-south`.
