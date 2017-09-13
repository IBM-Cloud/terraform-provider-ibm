---
layout: "ibm"
page_title: "Provider: IBM"
sidebar_current: "docs-ibm-index"
description: |-
  The IBM Cloud provider is used to interact with IBM Cloud resources.
---

# IBM Cloud Provider

The IBM Cloud provider is used to manage IBM Cloud resources. The provider must be configured with the proper credentials before it can be used.

Use the navigation menu on the left to read about the available data sources and resources.

## Example Usage


```hcl
# Configure the IBM Cloud Provider
provider "ibm" {
  bluemix_api_key    = "${var.ibm_bmx_api_key}"
  softlayer_username = "${var.ibm_sl_username}"
  softlayer_api_key  = "${var.ibm_sl_api_key}"
}

# Create an SSH key. You can find the SSH key surfaces in the SoftLayer console under Devices > Manage > SSH Keys
resource "ibm_compute_ssh_key" "test_key_1" {
  label      = "test_key_1"
  public_key = "${file("~/.ssh/id_rsa_test_key_1.pub")}"

  # Windows example:
  # public_key = "${file("C:\ssh\keys\path\id_rsa_test_key_1.pub")}"
}

# Create a virtual server with the SSH key
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

# Reference details of the IBM Bluemix Space
data "ibm_space" "space" {
  space = "${var.space}"
  org   = "${var.org}"
}

# Create an instance of an IBM service
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

You can provide your static credentials by adding the `bluemix_api_key`, `softlayer_username`, and `softlayer_api_key` arguments in the IBM Cloud provider block.

Usage:

```
provider "ibm" {
    bluemix_api_key = ""
    softlayer_username = ""
    softlayer_api_key = ""

}
```


### Environment variables

You can provide your credentials by exporting the `BM_API_KEY`, `SL_USERNAME`, and `SL_API_KEY` environment variables, representing your Bluemix API key, SoftLayer user name, and SoftLayer API key, respectively.  

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

* `bluemix_api_key` - (optional) The Bluemix API key. You must either add it as a credential in the provider block or source it from the `BM_API_KEY` (higher precedence) or `BLUEMIX_API_KEY` environment variable. The key is required to provision Cloud Foundry or IBM Container Service resources, such as any resource that begins with `ibm` or `ibm_container`.

* `bluemix_timeout` - (optional) The timeout, expressed in seconds, for interacting with Bluemix APIs. You can also source the timeout from the `BM_TIMEOUT` (higher precedence) or `BLUEMIX_TIMEOUT` environment variable. The default value is `60`.

* `softlayer_username` - (optional) The SoftLayer user name. You must either add it as a credential in the provider block or source it from the `SL_USERNAME` (higher precedence) or `SOFTLAYER_USERNAME` environment variable.

* `softlayer_api_key` - (optional) The SoftLayer API key. You must either add it as a credential in the provider block or source it from the `SL_API_KEY` (higher precedence) or `SOFTLAYER_API_KEY` environment variable. The key is required to provision SoftLayer resources, such as any resource that begins with `ibm_compute`.

* `softlayer_timeout` - (optional) The timeout, expressed in seconds, for the SoftLayer API key. You can also source the timeout from the `SL_TIMEOUT` (higher precedence) or `SOFTLAYER_TIMEOUT` environment variable. The default value is `60`.

* `region` - (optional) The Bluemix region. You can also source it from the `BM_REGION` (higher precedence) or `BLUEMIX_REGION` environment variable. The default value is `us-south`.
